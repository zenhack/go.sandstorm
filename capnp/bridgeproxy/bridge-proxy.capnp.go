// Code generated by capnpc-go. DO NOT EDIT.

package bridgeproxy

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	util "zenhack.net/go/sandstorm/capnp/util"
)

type ProxyClaimRequestRequest capnp.Struct

// ProxyClaimRequestRequest_TypeID is the unique identifier for the type ProxyClaimRequestRequest.
const ProxyClaimRequestRequest_TypeID = 0xe85d79ab9bde4783

func NewProxyClaimRequestRequest(s *capnp.Segment) (ProxyClaimRequestRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ProxyClaimRequestRequest(st), err
}

func NewRootProxyClaimRequestRequest(s *capnp.Segment) (ProxyClaimRequestRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return ProxyClaimRequestRequest(st), err
}

func ReadRootProxyClaimRequestRequest(msg *capnp.Message) (ProxyClaimRequestRequest, error) {
	root, err := msg.Root()
	return ProxyClaimRequestRequest(root.Struct()), err
}

func (s ProxyClaimRequestRequest) String() string {
	str, _ := text.Marshal(0xe85d79ab9bde4783, capnp.Struct(s))
	return str
}

func (s ProxyClaimRequestRequest) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ProxyClaimRequestRequest) DecodeFromPtr(p capnp.Ptr) ProxyClaimRequestRequest {
	return ProxyClaimRequestRequest(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ProxyClaimRequestRequest) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ProxyClaimRequestRequest) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ProxyClaimRequestRequest) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ProxyClaimRequestRequest) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s ProxyClaimRequestRequest) RequestToken() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s ProxyClaimRequestRequest) HasRequestToken() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s ProxyClaimRequestRequest) RequestTokenBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s ProxyClaimRequestRequest) SetRequestToken(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s ProxyClaimRequestRequest) RequiredPermissions() (capnp.TextList, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return capnp.TextList(p.List()), err
}

func (s ProxyClaimRequestRequest) HasRequiredPermissions() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s ProxyClaimRequestRequest) SetRequiredPermissions(v capnp.TextList) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewRequiredPermissions sets the requiredPermissions field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s ProxyClaimRequestRequest) NewRequiredPermissions(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s ProxyClaimRequestRequest) Label() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return util.LocalizedText(p.Struct()), err
}

func (s ProxyClaimRequestRequest) HasLabel() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s ProxyClaimRequestRequest) SetLabel(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(2, capnp.Struct(v).ToPtr())
}

// NewLabel sets the label field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s ProxyClaimRequestRequest) NewLabel() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(2, capnp.Struct(ss).ToPtr())
	return ss, err
}

// ProxyClaimRequestRequest_List is a list of ProxyClaimRequestRequest.
type ProxyClaimRequestRequest_List = capnp.StructList[ProxyClaimRequestRequest]

// NewProxyClaimRequestRequest creates a new list of ProxyClaimRequestRequest.
func NewProxyClaimRequestRequest_List(s *capnp.Segment, sz int32) (ProxyClaimRequestRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return capnp.StructList[ProxyClaimRequestRequest](l), err
}

// ProxyClaimRequestRequest_Future is a wrapper for a ProxyClaimRequestRequest promised by a client call.
type ProxyClaimRequestRequest_Future struct{ *capnp.Future }

func (f ProxyClaimRequestRequest_Future) Struct() (ProxyClaimRequestRequest, error) {
	p, err := f.Future.Ptr()
	return ProxyClaimRequestRequest(p.Struct()), err
}
func (p ProxyClaimRequestRequest_Future) Label() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(2, nil)}
}

type ProxyClaimRequestResponse capnp.Struct

// ProxyClaimRequestResponse_TypeID is the unique identifier for the type ProxyClaimRequestResponse.
const ProxyClaimRequestResponse_TypeID = 0x88b48ab0daf5d8d0

func NewProxyClaimRequestResponse(s *capnp.Segment) (ProxyClaimRequestResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ProxyClaimRequestResponse(st), err
}

func NewRootProxyClaimRequestResponse(s *capnp.Segment) (ProxyClaimRequestResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ProxyClaimRequestResponse(st), err
}

func ReadRootProxyClaimRequestResponse(msg *capnp.Message) (ProxyClaimRequestResponse, error) {
	root, err := msg.Root()
	return ProxyClaimRequestResponse(root.Struct()), err
}

func (s ProxyClaimRequestResponse) String() string {
	str, _ := text.Marshal(0x88b48ab0daf5d8d0, capnp.Struct(s))
	return str
}

func (s ProxyClaimRequestResponse) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (ProxyClaimRequestResponse) DecodeFromPtr(p capnp.Ptr) ProxyClaimRequestResponse {
	return ProxyClaimRequestResponse(capnp.Struct{}.DecodeFromPtr(p))
}

func (s ProxyClaimRequestResponse) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s ProxyClaimRequestResponse) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s ProxyClaimRequestResponse) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s ProxyClaimRequestResponse) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s ProxyClaimRequestResponse) Cap() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s ProxyClaimRequestResponse) HasCap() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s ProxyClaimRequestResponse) CapBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s ProxyClaimRequestResponse) SetCap(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// ProxyClaimRequestResponse_List is a list of ProxyClaimRequestResponse.
type ProxyClaimRequestResponse_List = capnp.StructList[ProxyClaimRequestResponse]

// NewProxyClaimRequestResponse creates a new list of ProxyClaimRequestResponse.
func NewProxyClaimRequestResponse_List(s *capnp.Segment, sz int32) (ProxyClaimRequestResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[ProxyClaimRequestResponse](l), err
}

// ProxyClaimRequestResponse_Future is a wrapper for a ProxyClaimRequestResponse promised by a client call.
type ProxyClaimRequestResponse_Future struct{ *capnp.Future }

func (f ProxyClaimRequestResponse_Future) Struct() (ProxyClaimRequestResponse, error) {
	p, err := f.Future.Ptr()
	return ProxyClaimRequestResponse(p.Struct()), err
}

const schema_bf72526e76ecd73b = "x\xda\x8c\x8e\xbfK\xebP\x00\x85\xcf\xb97}yC" +
	"\x7f\x85t\x7f\xf0\xe6\xf7\x04u\x10tPP\x11\x07!" +
	"\xb7\xb8vH\x9b\x8b\xa4\xa6I\x9a\xb4bqt\x11\xfd" +
	"\x13\xdc\\\x1dt\x10\x1c\xfd\x0bDqQ\x04\x05\x9d\x1c" +
	"\\\x1d%r\x15:8\xb9\x1d>>>N\xfdf\xc1" +
	"\x9a\xac\\I\x08\xf5\xaf\xf4\xab\xb8\xbe{\xbb?\xdd?" +
	"\xdb\x83\xe3\xb2\x98\xbb}\xdd\x8a\x9b\xd9\x05J\xb4\x81\xe9" +
	"Kv\xe9>\x99\xe9>p\x1e,vW\x1e\x0f\x8fG" +
	"\xad\x97o\xb24\xf2;\xdbt\x1da\xe4\x8a8AP" +
	"\xb4\xb30\xd8\xd0\xffS+K\xb6G\x13\x1d?\x8d\xd3" +
	"Y\xcf\xec\xc5\xc8\x0f{M\xdd\x1f\xea|\xd0\xd4yZ" +
	"K\xe2\\{\xa4\xb2\xa4\x05X\x04\x9c\xca_@\xfd\x96" +
	"T\x0dA\xbb\xe3\xa7,C\xb0\x0c\xfe4\xda\x1f\xda:" +
	"\x1f\x98fy\xdc\\\xee\x02jIRy\x82\x0e\xd9\xa0" +
	"\x81kG\x80\xf2$U$\xe8\x08\xd1\xa0\x00\x9cp\x0a" +
	"P\x81\xa4J\x05\x8b\xec\xab\xb9\x8eZ\xb2\xa9\xe3\xf1\x11" +
	"\x83\xc3L\x07\xf4t\xd6\x0b\xf3<\xb4\x938g\x15\xf4" +
	"$?\xa5*\xf8'\xf2\xdb:b\xbdX}\xde\x99\xc9" +
	"\xce[\x07\x00Y\x07?\x02\x00\x00\xff\xff\xc1Pg\x8e"

func init() {
	schemas.Register(schema_bf72526e76ecd73b,
		0x88b48ab0daf5d8d0,
		0xe85d79ab9bde4783)
}
