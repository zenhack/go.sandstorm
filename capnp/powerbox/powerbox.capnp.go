// Code generated by capnpc-go. DO NOT EDIT.

package powerbox

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	util "zenhack.net/go/sandstorm/capnp/util"
)

type PowerboxDescriptor capnp.Struct

// PowerboxDescriptor_TypeID is the unique identifier for the type PowerboxDescriptor.
const PowerboxDescriptor_TypeID = 0xedcf0fa3bfc71c58

func NewPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor(st), err
}

func NewRootPowerboxDescriptor(s *capnp.Segment) (PowerboxDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor(st), err
}

func ReadRootPowerboxDescriptor(msg *capnp.Message) (PowerboxDescriptor, error) {
	root, err := msg.Root()
	return PowerboxDescriptor(root.Struct()), err
}

func (s PowerboxDescriptor) String() string {
	str, _ := text.Marshal(0xedcf0fa3bfc71c58, capnp.Struct(s))
	return str
}

func (s PowerboxDescriptor) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (PowerboxDescriptor) DecodeFromPtr(p capnp.Ptr) PowerboxDescriptor {
	return PowerboxDescriptor(capnp.Struct{}.DecodeFromPtr(p))
}

func (s PowerboxDescriptor) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s PowerboxDescriptor) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s PowerboxDescriptor) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s PowerboxDescriptor) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s PowerboxDescriptor) Tags() (PowerboxDescriptor_Tag_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return PowerboxDescriptor_Tag_List(p.List()), err
}

func (s PowerboxDescriptor) HasTags() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s PowerboxDescriptor) SetTags(v PowerboxDescriptor_Tag_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewTags sets the tags field to a newly
// allocated PowerboxDescriptor_Tag_List, preferring placement in s's segment.
func (s PowerboxDescriptor) NewTags(n int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := NewPowerboxDescriptor_Tag_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return PowerboxDescriptor_Tag_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}
func (s PowerboxDescriptor) Quality() PowerboxDescriptor_MatchQuality {
	return PowerboxDescriptor_MatchQuality(capnp.Struct(s).Uint16(0))
}

func (s PowerboxDescriptor) SetQuality(v PowerboxDescriptor_MatchQuality) {
	capnp.Struct(s).SetUint16(0, uint16(v))
}

// PowerboxDescriptor_List is a list of PowerboxDescriptor.
type PowerboxDescriptor_List = capnp.StructList[PowerboxDescriptor]

// NewPowerboxDescriptor creates a new list of PowerboxDescriptor.
func NewPowerboxDescriptor_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[PowerboxDescriptor](l), err
}

// PowerboxDescriptor_Future is a wrapper for a PowerboxDescriptor promised by a client call.
type PowerboxDescriptor_Future struct{ *capnp.Future }

func (f PowerboxDescriptor_Future) Struct() (PowerboxDescriptor, error) {
	p, err := f.Future.Ptr()
	return PowerboxDescriptor(p.Struct()), err
}

type PowerboxDescriptor_Tag capnp.Struct

// PowerboxDescriptor_Tag_TypeID is the unique identifier for the type PowerboxDescriptor_Tag.
const PowerboxDescriptor_Tag_TypeID = 0xbe3d16a8df03c418

func NewPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag(st), err
}

func NewRootPowerboxDescriptor_Tag(s *capnp.Segment) (PowerboxDescriptor_Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PowerboxDescriptor_Tag(st), err
}

func ReadRootPowerboxDescriptor_Tag(msg *capnp.Message) (PowerboxDescriptor_Tag, error) {
	root, err := msg.Root()
	return PowerboxDescriptor_Tag(root.Struct()), err
}

func (s PowerboxDescriptor_Tag) String() string {
	str, _ := text.Marshal(0xbe3d16a8df03c418, capnp.Struct(s))
	return str
}

func (s PowerboxDescriptor_Tag) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (PowerboxDescriptor_Tag) DecodeFromPtr(p capnp.Ptr) PowerboxDescriptor_Tag {
	return PowerboxDescriptor_Tag(capnp.Struct{}.DecodeFromPtr(p))
}

func (s PowerboxDescriptor_Tag) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s PowerboxDescriptor_Tag) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s PowerboxDescriptor_Tag) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s PowerboxDescriptor_Tag) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s PowerboxDescriptor_Tag) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s PowerboxDescriptor_Tag) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s PowerboxDescriptor_Tag) Value() (capnp.Ptr, error) {
	return capnp.Struct(s).Ptr(0)
}

func (s PowerboxDescriptor_Tag) HasValue() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s PowerboxDescriptor_Tag) SetValue(v capnp.Ptr) error {
	return capnp.Struct(s).SetPtr(0, v)
}

// PowerboxDescriptor_Tag_List is a list of PowerboxDescriptor_Tag.
type PowerboxDescriptor_Tag_List = capnp.StructList[PowerboxDescriptor_Tag]

// NewPowerboxDescriptor_Tag creates a new list of PowerboxDescriptor_Tag.
func NewPowerboxDescriptor_Tag_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[PowerboxDescriptor_Tag](l), err
}

// PowerboxDescriptor_Tag_Future is a wrapper for a PowerboxDescriptor_Tag promised by a client call.
type PowerboxDescriptor_Tag_Future struct{ *capnp.Future }

func (f PowerboxDescriptor_Tag_Future) Struct() (PowerboxDescriptor_Tag, error) {
	p, err := f.Future.Ptr()
	return PowerboxDescriptor_Tag(p.Struct()), err
}
func (p PowerboxDescriptor_Tag_Future) Value() *capnp.Future {
	return p.Future.Field(0, nil)
}

type PowerboxDescriptor_MatchQuality uint16

// PowerboxDescriptor_MatchQuality_TypeID is the unique identifier for the type PowerboxDescriptor_MatchQuality.
const PowerboxDescriptor_MatchQuality_TypeID = 0xbb1afa45d25ce8de

// Values of PowerboxDescriptor_MatchQuality.
const (
	PowerboxDescriptor_MatchQuality_preferred    PowerboxDescriptor_MatchQuality = 1
	PowerboxDescriptor_MatchQuality_acceptable   PowerboxDescriptor_MatchQuality = 0
	PowerboxDescriptor_MatchQuality_unacceptable PowerboxDescriptor_MatchQuality = 2
)

// String returns the enum's constant name.
func (c PowerboxDescriptor_MatchQuality) String() string {
	switch c {
	case PowerboxDescriptor_MatchQuality_preferred:
		return "preferred"
	case PowerboxDescriptor_MatchQuality_acceptable:
		return "acceptable"
	case PowerboxDescriptor_MatchQuality_unacceptable:
		return "unacceptable"

	default:
		return ""
	}
}

// PowerboxDescriptor_MatchQualityFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PowerboxDescriptor_MatchQualityFromString(c string) PowerboxDescriptor_MatchQuality {
	switch c {
	case "preferred":
		return PowerboxDescriptor_MatchQuality_preferred
	case "acceptable":
		return PowerboxDescriptor_MatchQuality_acceptable
	case "unacceptable":
		return PowerboxDescriptor_MatchQuality_unacceptable

	default:
		return 0
	}
}

type PowerboxDescriptor_MatchQuality_List = capnp.EnumList[PowerboxDescriptor_MatchQuality]

func NewPowerboxDescriptor_MatchQuality_List(s *capnp.Segment, sz int32) (PowerboxDescriptor_MatchQuality_List, error) {
	return capnp.NewEnumList[PowerboxDescriptor_MatchQuality](s, sz)
}

type PowerboxDisplayInfo capnp.Struct

// PowerboxDisplayInfo_TypeID is the unique identifier for the type PowerboxDisplayInfo.
const PowerboxDisplayInfo_TypeID = 0xa553a209bee32bec

func NewPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo(st), err
}

func NewRootPowerboxDisplayInfo(s *capnp.Segment) (PowerboxDisplayInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return PowerboxDisplayInfo(st), err
}

func ReadRootPowerboxDisplayInfo(msg *capnp.Message) (PowerboxDisplayInfo, error) {
	root, err := msg.Root()
	return PowerboxDisplayInfo(root.Struct()), err
}

func (s PowerboxDisplayInfo) String() string {
	str, _ := text.Marshal(0xa553a209bee32bec, capnp.Struct(s))
	return str
}

func (s PowerboxDisplayInfo) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (PowerboxDisplayInfo) DecodeFromPtr(p capnp.Ptr) PowerboxDisplayInfo {
	return PowerboxDisplayInfo(capnp.Struct{}.DecodeFromPtr(p))
}

func (s PowerboxDisplayInfo) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s PowerboxDisplayInfo) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s PowerboxDisplayInfo) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s PowerboxDisplayInfo) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s PowerboxDisplayInfo) Title() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return util.LocalizedText(p.Struct()), err
}

func (s PowerboxDisplayInfo) HasTitle() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s PowerboxDisplayInfo) SetTitle(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewTitle sets the title field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewTitle() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) VerbPhrase() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return util.LocalizedText(p.Struct()), err
}

func (s PowerboxDisplayInfo) HasVerbPhrase() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s PowerboxDisplayInfo) SetVerbPhrase(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(1, capnp.Struct(v).ToPtr())
}

// NewVerbPhrase sets the verbPhrase field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewVerbPhrase() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(1, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s PowerboxDisplayInfo) Description() (util.LocalizedText, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return util.LocalizedText(p.Struct()), err
}

func (s PowerboxDisplayInfo) HasDescription() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s PowerboxDisplayInfo) SetDescription(v util.LocalizedText) error {
	return capnp.Struct(s).SetPtr(2, capnp.Struct(v).ToPtr())
}

// NewDescription sets the description field to a newly
// allocated util.LocalizedText struct, preferring placement in s's segment.
func (s PowerboxDisplayInfo) NewDescription() (util.LocalizedText, error) {
	ss, err := util.NewLocalizedText(capnp.Struct(s).Segment())
	if err != nil {
		return util.LocalizedText{}, err
	}
	err = capnp.Struct(s).SetPtr(2, capnp.Struct(ss).ToPtr())
	return ss, err
}

// PowerboxDisplayInfo_List is a list of PowerboxDisplayInfo.
type PowerboxDisplayInfo_List = capnp.StructList[PowerboxDisplayInfo]

// NewPowerboxDisplayInfo creates a new list of PowerboxDisplayInfo.
func NewPowerboxDisplayInfo_List(s *capnp.Segment, sz int32) (PowerboxDisplayInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return capnp.StructList[PowerboxDisplayInfo](l), err
}

// PowerboxDisplayInfo_Future is a wrapper for a PowerboxDisplayInfo promised by a client call.
type PowerboxDisplayInfo_Future struct{ *capnp.Future }

func (f PowerboxDisplayInfo_Future) Struct() (PowerboxDisplayInfo, error) {
	p, err := f.Future.Ptr()
	return PowerboxDisplayInfo(p.Struct()), err
}
func (p PowerboxDisplayInfo_Future) Title() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(0, nil)}
}
func (p PowerboxDisplayInfo_Future) VerbPhrase() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(1, nil)}
}
func (p PowerboxDisplayInfo_Future) Description() util.LocalizedText_Future {
	return util.LocalizedText_Future{Future: p.Future.Field(2, nil)}
}

const schema_f6c200ab14cd53e4 = "x\xdat\x92Mk\x13_\x18\xc5\xcf\xb97\xf9\xa7\x81" +
	"\x7f\x9d\xdeN\x15\x0dJ (\xbea\xe8\x9bT\x84b" +
	"\x90\x16\x0c(\xe6\xda,\xa4\xd4\xc5M2M\xa7\x84d" +
	"\x9cL\x1a\xa3\x8b~\x00W~\x05E\x10\\\xb8t!" +
	"T\x10\xc1\x95\x88~\x00\x8bo\xe0B\xb7.\xdc\x8cL" +
	"^L\x90v9\xcf9s\xcfy~<\xd3\x0b\xcc\xc5" +
	"f\xc6\xb7c\x10z>\xfe_\xf8\xe3\xec\x97\x9d\xe4\xc3" +
	"\x95\xc7P\x16w\xbf\xae\xbc\x9dz\xfa\xeaW\\&\x80" +
	"\xb9\x13\"E\xfb\xbcH\x00\xf6\x8cx\x06\x86\x1f\xbf\xaf" +
	"}X\xfe\x9dz\x01\x95\x11\xe1\xcd\xa3o^>\xb2\xde" +
	"\xfd\x048\xf7^\xdc\xa5\xfd\xad\xeb\xfc$\xae\x83\xe1\xe1" +
	"\xd7r\xf7\xc9\xa1\xc5\x1d\xe8\x0c9\xb4\xc6\x19\xbdK9" +
	"K[E\x11\xf6\xb8lcD\xd7\x16\x87\x1d\xba^W" +
	"N\xd2\xee\xc8\x93\x80\xfd@\xb6Q\x09\xbdF\xdb\xf1K" +
	"\x8d;2[6^\xdd\xbbX\xe8\x7f/\xb9M\xaff" +
	":\xf9\xba\\o\x14H\xfd\xbf\x8c\x011\x02jy\x16" +
	"\xd09I}UP\x91S\x8c\x86\xf9U@_\x91\xd4" +
	"EA%\xc4\x14\x05\xa0t\x09\xd0\x05I\xbd&\x98\x0e" +
	"\xdc\xa0\xe6p\"\xcc\x7f\xbe\xb7\xe0?\xbfu\x1f '" +
	"\xc0p\xcb\xf1K\x85\x0d\xdf@6\xf7\x92+N\xb3\xec" +
	"\xbb^\x80\x84\xdb\xa8\xef\xa1\x0f\xfa\xc7\xfe\xed\xdf\xff\xaf" +
	"\xe1g\xaf\x99\xa0\xbc\xa1/\xb5L\xcd\x0d:\xdd](" +
	"Hul\x15TGn\x00\x14\xea\xe0&\x10\x9ar\xd9" +
	"\xf1\x02S\x82\xac9\xa1\xe7;\xeb\x8e\xef;`%l" +
	"\xd5{\x12,S\x8a\xa4\xfd\x90\x0d\"\xd3~\xb6h\xaa" +
	"Q\xd2\xd8_j\xa7S\x80>.\xa9\xa7\x05\x07\xd0\xce" +
	"E$OI\xeayA\xe9V\x98\x84`\x12Lo\x99" +
	"Z\xcb\xe1$\x04'G6\xdc7\x8e\xbe\x1e\xe3\xc8\x8d" +
	"\xa8df\xe4\xb4\xe2\x9b\x89\xa2\xa9\x86=\x08-\x03+" +
	"\xa20Z\xecL\xbfX\xae_\x8cT\x8b\x97\x01}A" +
	"R/\x09Z\x81\xa96y\x00,Hrb\x98\x02F" +
	"\xc3\xed\xdb=\xac\xb4\x86\x91 -\xf0O\x00\x00\x00\xff" +
	"\xff\xd9(\xda\xeb"

func init() {
	schemas.Register(schema_f6c200ab14cd53e4,
		0xa553a209bee32bec,
		0xbb1afa45d25ce8de,
		0xbe3d16a8df03c418,
		0xedcf0fa3bfc71c58)
}
