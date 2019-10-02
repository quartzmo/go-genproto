// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/text_label.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A type of label displaying text on a colored background.
type TextLabel struct {
	// Background color of the label in RGB format. This string must match the
	// regular expression '^\#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$'.
	// Note: The background color may not be visible for manager accounts.
	BackgroundColor *wrappers.StringValue `protobuf:"bytes,1,opt,name=background_color,json=backgroundColor,proto3" json:"background_color,omitempty"`
	// A short description of the label. The length must be no more than 200
	// characters.
	Description          *wrappers.StringValue `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TextLabel) Reset()         { *m = TextLabel{} }
func (m *TextLabel) String() string { return proto.CompactTextString(m) }
func (*TextLabel) ProtoMessage()    {}
func (*TextLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e124a9c58395a2ff, []int{0}
}

func (m *TextLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextLabel.Unmarshal(m, b)
}
func (m *TextLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextLabel.Marshal(b, m, deterministic)
}
func (m *TextLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextLabel.Merge(m, src)
}
func (m *TextLabel) XXX_Size() int {
	return xxx_messageInfo_TextLabel.Size(m)
}
func (m *TextLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_TextLabel.DiscardUnknown(m)
}

var xxx_messageInfo_TextLabel proto.InternalMessageInfo

func (m *TextLabel) GetBackgroundColor() *wrappers.StringValue {
	if m != nil {
		return m.BackgroundColor
	}
	return nil
}

func (m *TextLabel) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func init() {
	proto.RegisterType((*TextLabel)(nil), "google.ads.googleads.v2.common.TextLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/text_label.proto", fileDescriptor_e124a9c58395a2ff)
}

var fileDescriptor_e124a9c58395a2ff = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x69, 0x05, 0xc1, 0x0e, 0x54, 0x76, 0x1a, 0x63, 0x0c, 0xd9, 0xc9, 0x53, 0x02, 0xf5,
	0x16, 0x41, 0xe8, 0x26, 0xec, 0xe2, 0x61, 0xa8, 0xf4, 0x20, 0x85, 0x91, 0x36, 0x31, 0x14, 0xb3,
	0x7c, 0x21, 0xc9, 0xe6, 0x1e, 0xc4, 0x27, 0xf0, 0xe8, 0xa3, 0xf8, 0x26, 0xfa, 0x14, 0x92, 0xa6,
	0xad, 0x5e, 0x14, 0x4f, 0xfd, 0xd3, 0xfc, 0x7e, 0xdf, 0x3f, 0xf9, 0x12, 0x2c, 0x00, 0x84, 0xe4,
	0x98, 0x32, 0xdb, 0x46, 0x9f, 0x76, 0x29, 0xae, 0x60, 0xb3, 0x01, 0x85, 0x1d, 0xdf, 0xbb, 0xb5,
	0xa4, 0x25, 0x97, 0x48, 0x1b, 0x70, 0x30, 0x9c, 0x06, 0x0a, 0x51, 0x66, 0x51, 0x2f, 0xa0, 0x5d,
	0x8a, 0x82, 0x30, 0x6e, 0xcf, 0x71, 0x43, 0x97, 0xdb, 0x47, 0xfc, 0x6c, 0xa8, 0xd6, 0xdc, 0xd8,
	0xe0, 0x8f, 0x27, 0x5d, 0xa1, 0xae, 0x31, 0x55, 0x0a, 0x1c, 0x75, 0x35, 0xa8, 0xf6, 0x74, 0xf6,
	0x12, 0x25, 0x47, 0xf7, 0x7c, 0xef, 0x6e, 0x7c, 0xe3, 0x70, 0x99, 0x9c, 0x96, 0xb4, 0x7a, 0x12,
	0x06, 0xb6, 0x8a, 0xad, 0x2b, 0x90, 0x60, 0x46, 0xd1, 0x59, 0x74, 0x3e, 0x48, 0x27, 0x6d, 0x37,
	0xea, 0x6a, 0xd0, 0x9d, 0x33, 0xb5, 0x12, 0x39, 0x95, 0x5b, 0x7e, 0x7b, 0xf2, 0x6d, 0x2d, 0xbc,
	0x34, 0xbc, 0x4a, 0x06, 0x8c, 0xdb, 0xca, 0xd4, 0xda, 0x97, 0x8d, 0xe2, 0x7f, 0xcc, 0xf8, 0x29,
	0xcc, 0x3f, 0xa2, 0x64, 0x56, 0xc1, 0x06, 0xfd, 0xfd, 0xf6, 0xf9, 0x71, 0x7f, 0xf5, 0x95, 0x1f,
	0xb9, 0x8a, 0x1e, 0xae, 0x5b, 0x43, 0x80, 0xa4, 0x4a, 0x20, 0x30, 0x02, 0x0b, 0xae, 0x9a, 0xc2,
	0x6e, 0xdd, 0xba, 0xb6, 0xbf, 0x6d, 0xff, 0x32, 0x7c, 0x5e, 0xe3, 0x83, 0x65, 0x96, 0xbd, 0xc5,
	0xd3, 0x65, 0x18, 0x96, 0x31, 0x8b, 0x42, 0xf4, 0x29, 0x4f, 0xd1, 0xa2, 0xc1, 0xde, 0x3b, 0xa0,
	0xc8, 0x98, 0x2d, 0x7a, 0xa0, 0xc8, 0xd3, 0x22, 0x00, 0x9f, 0xf1, 0x2c, 0xfc, 0x25, 0x24, 0x63,
	0x96, 0x90, 0x1e, 0x21, 0x24, 0x4f, 0x09, 0x09, 0x50, 0x79, 0xd8, 0xdc, 0xee, 0xe2, 0x2b, 0x00,
	0x00, 0xff, 0xff, 0x72, 0xf9, 0xd6, 0xc2, 0x1a, 0x02, 0x00, 0x00,
}