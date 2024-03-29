package websession

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	"zenhack.net/go/sandstorm/capnp/util"
	"zenhack.net/go/sandstorm/capnp/websession"
)

// A Handler implements http.Handler on top of a WebSession. NOTE: this is work in progress
// and does not handle all requests gracefully (sometime codepaths simply panic("TODO")).
type Handler struct {
	Session websession.WebSession
}

// ServeHTTP implements http.Handler.ServeHTTP
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.doGet(w, req, false)
	case "HEAD":
		h.doGet(w, req, true)
	default:
		panic("TODO")
	}
}

// doGet makes a request using the WebSession.get() method.
func (h Handler) doGet(w http.ResponseWriter, req *http.Request, ignoreBody bool) {
	// TODO: block sending the body when ignoreBody = true
	responseStreamServer := newResponseStreamImpl(w)
	responseStreamClient := util.ByteStream_ServerToClient(responseStreamServer)

	respFut, rel := h.Session.Get(req.Context(), func(p websession.WebSession_get_Params) error {
		if !strings.HasPrefix(req.RequestURI, "/") {
			return fmt.Errorf("Error: malformed RequestURI (no leading slash): %q", req.RequestURI)
		}
		path := req.RequestURI[1:]
		if err := p.SetPath(path); err != nil {
			return err
		}
		wsCtx, err := p.NewContext()
		if err != nil {
			return err
		}

		if err = populateContext(wsCtx, req, responseStreamClient); err != nil {
			return err
		}

		p.SetIgnoreBody(ignoreBody)
		return nil
	})
	defer rel()
	resp, err := respFut.Struct()
	if err != nil {
		replyErr(w, err)
		return
	}
	relayResponse(
		req.Context(),
		w,
		req,
		resp,
		responseStreamServer,
	)
}

// relayResponse relays a response received from a WebSession back to an http.ResponseWriter.
//
// responseStream should be the value of WebSession.Context.responseStream that was passed in
// to the request.
func relayResponse(
	ctx context.Context,
	w http.ResponseWriter,
	req *http.Request,
	resp websession.WebSession_Response,
	responseStream *responseStreamImpl,
) {
	status, err := responseStatus(req, resp)
	if err != nil {
		replyErr(w, err)
		close(responseStream.ready)
		return
	}

	if resp.Which() == websession.WebSession_Response_Which_content {
		content := resp.Content()
		body := content.Body()
		if body.Which() == websession.WebSession_Response_content_body_Which_stream {
			defer body.Stream().Release()
			responseStream.used = true

			select {
			case <-ctx.Done():
				return
			case size, ok := <-responseStream.size:
				if ok {
					w.Header().Set("Content-Length", strconv.FormatUint(size, 10))
				}
			}
			if err := populateResponseHeaders(w, req, resp); err != nil {
				replyErr(w, err)
				return
			}
			w.WriteHeader(status)
			close(responseStream.ready)
			select {
			case <-ctx.Done():
			case <-responseStream.done:
			case <-responseStream.shutdown:
			}
			return
		}
	}

	close(responseStream.ready)
	if err := populateResponseHeaders(w, req, resp); err != nil {
		replyErr(w, err)
		return
	}
	data, err := responseBodyBytes(resp)
	if err != nil {
		replyErr(w, err)
		return
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

// responseStatus returns the correct HTTP status code for the the response.
// req is the original request that this is a response to.
func responseStatus(req *http.Request, resp websession.WebSession_Response) (int, error) {
	switch resp.Which() {
	case websession.WebSession_Response_Which_content:
		successCode := resp.Content().StatusCode()
		status, ok := successCodeStatuses[successCode]
		if ok {
			return status, nil
		}
		return 0, fmt.Errorf("Unknown success code enumerant: %v", successCode)
	case websession.WebSession_Response_Which_noContent:
		if resp.NoContent().ShouldResetForm() {
			return http.StatusResetContent, nil
		} else {
			return http.StatusNoContent, nil
		}
	case websession.WebSession_Response_Which_preconditionFailed:
		if (req.Method == "GET" || req.Method == "HEAD") &&
			req.Header.Get("If-None-Match") != "" {

			return http.StatusNotModified, nil
		}
		return http.StatusPreconditionFailed, nil
	case websession.WebSession_Response_Which_redirect:
		r := resp.Redirect()
		switch {
		case r.IsPermanent() && r.SwitchToGet():
			return http.StatusMovedPermanently, nil
		case !r.IsPermanent() && r.SwitchToGet():
			return http.StatusSeeOther, nil
		case r.IsPermanent() && !r.SwitchToGet():
			return http.StatusPermanentRedirect, nil
		default: // !r.IsPermanent() && !r.SwitchToGet():
			return http.StatusTemporaryRedirect, nil
		}
	case websession.WebSession_Response_Which_clientError:
		errorCode := resp.ClientError().StatusCode()
		status, ok := clientErrorCodeStatuses[errorCode]
		if ok {
			return status, nil
		}
		return 0, fmt.Errorf("Unknown error code enumerant: %v", errorCode)
	case websession.WebSession_Response_Which_serverError:
		return http.StatusInternalServerError, nil
	default:
		return 0, fmt.Errorf("Unknown response variant: %v", resp.Which())
	}
}

// Return a []byte with the body of the response. This will return an error for
// streaming responses.
func responseBodyBytes(resp websession.WebSession_Response) ([]byte, error) {
	switch resp.Which() {
	case websession.WebSession_Response_Which_content:
		body := resp.Content().Body()
		switch body.Which() {
		case websession.WebSession_Response_content_body_Which_bytes:
			return body.Bytes()
		case websession.WebSession_Response_content_body_Which_stream:
			return nil, fmt.Errorf("Can't get []byte for streaming body")
		default:
			return nil, fmt.Errorf("Unknown body variant: %v", body.Which())
		}
	case websession.WebSession_Response_Which_noContent:
		return nil, nil
	case websession.WebSession_Response_Which_preconditionFailed:
		return nil, nil
	case websession.WebSession_Response_Which_redirect:
		return nil, nil
	case websession.WebSession_Response_Which_clientError:
		return errorBodyBytes(resp.ClientError())
	case websession.WebSession_Response_Which_serverError:
		return errorBodyBytes(resp.ServerError())
	default:
		return nil, fmt.Errorf("Unknown response variant: %v", resp.Which())
	}
}

// errorBodyBytes is a helper for responseBodyBytes; it handles the clientError and serverError
// cases.
func errorBodyBytes(r hasErrorBody) ([]byte, error) {
	if !r.HasNonHtmlBody() {
		str, err := r.DescriptionHtml()
		return []byte(str), err
	}
	errBody, err := r.NonHtmlBody()
	if err != nil {
		return nil, err
	}
	return errBody.Data()
}

func populateErrorBodyHeaders(dst http.Header, src hasErrorBody) error {
	if !src.HasNonHtmlBody() {
		dst.Set("Content-Type", "text/html")
		return nil
	}
	errBody, err := src.NonHtmlBody()
	if err != nil {
		return err
	}
	return populateHasContentHeaders(dst, errBody)
}

type hasErrorBody interface {
	DescriptionHtml() (string, error)
	NonHtmlBody() (websession.WebSession_Response_ErrorBody, error)
	HasNonHtmlBody() bool
}

// populateResponseHeaders fills in the response headers based on the contents of the response.
func populateResponseHeaders(w http.ResponseWriter, req *http.Request, resp websession.WebSession_Response) error {
	isHttps := req.TLS != nil

	setCookies, err := resp.SetCookies()
	if err != nil {
		return err
	}
	for i := 0; i < setCookies.Len(); i++ {
		setCookie := setCookies.At(i)
		name, err := setCookie.Name()
		if err != nil {
			return err
		}
		value, err := setCookie.Value()
		if err != nil {
			return err
		}
		path, err := setCookie.Path()
		if err != nil {
			return err
		}
		cookie := &http.Cookie{
			Name:   name,
			Value:  value,
			Secure: isHttps,
			Path:   path,
		}
		expires := setCookie.Expires()
		switch expires.Which() {
		case websession.WebSession_Cookie_expires_Which_none:
		case websession.WebSession_Cookie_expires_Which_absolute:
			cookie.Expires = time.Unix(expires.Absolute(), 0)
		case websession.WebSession_Cookie_expires_Which_relative:
			cookie.Expires = time.Now().Add(time.Duration(expires.Relative()) * time.Second)
		}
		http.SetCookie(w, cookie)
	}
	// TODO: cachePolicy

	additionalHeaders, err := resp.AdditionalHeaders()
	if err != nil {
		return err
	}
	wHeaders := w.Header()
	for i := 0; i < additionalHeaders.Len(); i++ {
		item := additionalHeaders.At(i)
		name, err := item.Name()
		if err != nil {
			return err
		}
		v, err := item.Value()
		if err != nil {
			return err
		}
		k := http.CanonicalHeaderKey(name)
		if ResponseHeaderFilter.Allows(k) {
			wHeaders[k] = append(wHeaders[k], v)
		}
	}

	switch resp.Which() {
	case websession.WebSession_Response_Which_content:
		return populateContentResponseHeaders(wHeaders, resp.Content())
	case websession.WebSession_Response_Which_noContent:
		nc := resp.NoContent()
		if nc.HasETag() {
			etag, err := nc.ETag()
			if err != nil {
				return err
			}
			return setETag(wHeaders, etag)
		}
		return nil
	case websession.WebSession_Response_Which_preconditionFailed:
		pf := resp.PreconditionFailed()
		if pf.HasMatchingETag() {
			etag, err := pf.MatchingETag()
			if err != nil {
				return err
			}
			return setETag(wHeaders, etag)
		}
		return nil
	case websession.WebSession_Response_Which_redirect:
		loc, err := resp.Redirect().Location()
		if err != nil {
			return err
		}
		wHeaders.Set("Location", loc)
		return nil
	case websession.WebSession_Response_Which_clientError:
		return populateErrorBodyHeaders(w.Header(), resp.ClientError())
	case websession.WebSession_Response_Which_serverError:
		return populateErrorBodyHeaders(w.Header(), resp.ServerError())
	default:
		return fmt.Errorf("Unknown response variant: %v", resp.Which())
	}
}

func setETag(h http.Header, etag websession.WebSession_ETag) error {
	s, err := eTagStr(etag)
	if err != nil {
		return err
	}
	h.Set("ETag", s)
	return nil
}

func eTagStr(etag websession.WebSession_ETag) (string, error) {
	value, err := etag.Value()
	if err != nil {
		return "", err
	}
	if etag.Weak() {
		// FIXME: do we need to escape this?
		return "W/\"" + value + "\"", nil
	}
	return "\"" + value + "\"", nil
}

func populateContentResponseHeaders(h http.Header, r websession.WebSession_Response_content) error {
	if err := populateHasContentHeaders(h, r); err != nil {
		return err
	}
	disposition := r.Disposition()
	switch disposition.Which() {
	case websession.WebSession_Response_content_disposition_Which_normal:
		// Default
	case websession.WebSession_Response_content_disposition_Which_download:
		filename, err := disposition.Download()
		if err != nil {
			return err
		}
		h.Set("Content-Disposition", mime.FormatMediaType(
			"attachment",
			map[string]string{"filename": filename},
		))
	}
	if r.HasETag() {
		etag, err := r.ETag()
		if err != nil {
			return err
		}
		return setETag(h, etag)
	}
	return nil
}

func populateHasContentHeaders(dst http.Header, src hasContent) error {
	if src.HasEncoding() {
		v, err := src.Encoding()
		if err != nil {
			return err
		}
		dst.Set("Content-Encoding", v)
	}
	if src.HasLanguage() {
		v, err := src.Language()
		if err != nil {
			return err
		}
		dst.Set("Content-Language", v)
	}
	contentType, err := src.MimeType()
	if err != nil {
		return err
	}
	// parse & re-format to ensure well-formedness:
	mimeType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	contentType = mime.FormatMediaType(mimeType, params)
	dst.Set("Content-Type", contentType)
	return nil
}

type hasContent interface {
	HasEncoding() bool
	Encoding() (string, error)
	HasLanguage() bool
	Language() (string, error)
	MimeType() (string, error)
}

// replyErr responds to the request with a 500 status and the given error. If any headers had been
// set in the response, they will be cleared.
func replyErr(w http.ResponseWriter, err error) {
	hdr := w.Header()

	// Delete any headers we may have set when trying to build a normal response.
	for k, _ := range hdr {
		delete(hdr, k)
	}

	body := []byte(err.Error())

	hdr.Set("Content-Type", "text/plain")
	hdr.Set("Content-Length", strconv.Itoa(len(body)))

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

// populateContext populates a websession context based on the request, using the supplied
// value for the responseStream field. The reference to responseStream is stolen.
func populateContext(wsCtx websession.WebSession_Context, req *http.Request, responseStream util.ByteStream) error {
	// Copy in cookies
	reqCookies := req.Cookies()
	wsCookies, err := wsCtx.NewCookies(int32(len(reqCookies)))
	if err != nil {
		return err
	}
	for i, c := range reqCookies {
		wsC := wsCookies.At(i)
		wsC.SetKey(c.Name)
		wsC.SetValue(c.Value)
	}

	// Rig up the response body
	wsCtx.SetResponseStream(responseStream)

	// Process the Accept header
	if accept := req.Header.Get("Accept"); accept != "" {
		types := strings.Split(accept, ",")
		wsTypes, err := wsCtx.NewAccept(int32(len(types)))
		if err != nil {
			return err
		}
		for i, t := range types {
			mimeType, params, err := mime.ParseMediaType(t)
			if err != nil {
				return fmt.Errorf("Error parsing media type at index %v (%q): %w", i, t, err)
			}
			acceptedType := wsTypes.At(i)
			acceptedType.SetMimeType(mimeType)
			qStr, ok := params["q"]
			if ok {
				q, err := strconv.ParseFloat(qStr, 32)
				if err != nil {
					return fmt.Errorf(
						"Error parsing qValue %q in media type %q: %w",
						qStr, t, err)
				}
				acceptedType.SetQValue(float32(q))
			}
		}
	}

	// TODO: acceptEncoding
	// TODO: eTagPrecondition

	additionalHeaders := make(http.Header)
	ContextHeaderFilter.Copy(additionalHeaders, req.Header)
	var numHeaders int32
	for _, vs := range additionalHeaders {
		numHeaders += int32(len(vs))
	}

	dstAdditionalHeaders, err := wsCtx.NewAdditionalHeaders(numHeaders)
	if err != nil {
		return err
	}
	i := 0
	for k, vs := range additionalHeaders {
		for _, v := range vs {
			h := dstAdditionalHeaders.At(i)
			if err := h.SetName(k); err != nil {
				return err
			}
			if err := h.SetValue(v); err != nil {
				return err
			}
			i++
		}
	}

	//panic("TODO")
	return nil
}
