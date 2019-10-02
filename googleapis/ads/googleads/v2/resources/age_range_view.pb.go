// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/age_range_view.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// An age range view.
type AgeRangeView struct {
	// The resource name of the age range view.
	// Age range view resource names have the form:
	//
	// `customers/{customer_id}/ageRangeViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRangeView) Reset()         { *m = AgeRangeView{} }
func (m *AgeRangeView) String() string { return proto.CompactTextString(m) }
func (*AgeRangeView) ProtoMessage()    {}
func (*AgeRangeView) Descriptor() ([]byte, []int) {
	return fileDescriptor_122920b378c098b3, []int{0}
}

func (m *AgeRangeView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRangeView.Unmarshal(m, b)
}
func (m *AgeRangeView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRangeView.Marshal(b, m, deterministic)
}
func (m *AgeRangeView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRangeView.Merge(m, src)
}
func (m *AgeRangeView) XXX_Size() int {
	return xxx_messageInfo_AgeRangeView.Size(m)
}
func (m *AgeRangeView) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRangeView.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRangeView proto.InternalMessageInfo

func (m *AgeRangeView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*AgeRangeView)(nil), "google.ads.googleads.v2.resources.AgeRangeView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/age_range_view.proto", fileDescriptor_122920b378c098b3)
}

var fileDescriptor_122920b378c098b3 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4a, 0xc4, 0x30,
	0x18, 0xc7, 0x69, 0x05, 0xc1, 0x72, 0x0e, 0xde, 0x24, 0xe2, 0xe0, 0x29, 0x07, 0x4e, 0x09, 0xe4,
	0xc0, 0x21, 0x4e, 0xb9, 0xe5, 0xc0, 0x41, 0x8e, 0x0e, 0x1d, 0xa4, 0x50, 0xe2, 0xe5, 0x23, 0x04,
	0xae, 0xf9, 0x4a, 0x52, 0x7b, 0xef, 0xe3, 0xe8, 0xa3, 0xf8, 0x28, 0x3e, 0x84, 0x48, 0x1b, 0x13,
	0x9c, 0xbc, 0xed, 0x4f, 0xf2, 0xfb, 0xff, 0xbe, 0x2f, 0x29, 0x1e, 0x34, 0xa2, 0xde, 0x03, 0x95,
	0xca, 0xd3, 0x10, 0xc7, 0x34, 0x30, 0xea, 0xc0, 0xe3, 0x9b, 0xdb, 0x81, 0xa7, 0x52, 0x43, 0xe3,
	0xa4, 0xd5, 0xd0, 0x0c, 0x06, 0x0e, 0xa4, 0x73, 0xd8, 0xe3, 0x7c, 0x11, 0x60, 0x22, 0x95, 0x27,
	0xa9, 0x47, 0x06, 0x46, 0x52, 0xef, 0xea, 0x3a, 0xaa, 0x3b, 0x43, 0xa5, 0xb5, 0xd8, 0xcb, 0xde,
	0xa0, 0xf5, 0x41, 0x70, 0xbb, 0x2a, 0x66, 0x42, 0x43, 0x39, 0x7a, 0x2b, 0x03, 0x87, 0xf9, 0x5d,
	0x71, 0x1e, 0xab, 0x8d, 0x95, 0x2d, 0x5c, 0x66, 0x37, 0xd9, 0xfd, 0x59, 0x39, 0x8b, 0x87, 0xcf,
	0xb2, 0x85, 0xf5, 0x77, 0x56, 0x2c, 0x77, 0xd8, 0x92, 0xa3, 0xc3, 0xd7, 0x17, 0x7f, 0xe5, 0xdb,
	0x71, 0xe2, 0x36, 0x7b, 0x79, 0xfa, 0xed, 0x69, 0xdc, 0x4b, 0xab, 0x09, 0x3a, 0x4d, 0x35, 0xd8,
	0x69, 0x9f, 0xf8, 0xf8, 0xce, 0xf8, 0x7f, 0xfe, 0xe2, 0x31, 0xa5, 0xf7, 0xfc, 0x64, 0x23, 0xc4,
	0x47, 0xbe, 0xd8, 0x04, 0xa5, 0x50, 0x9e, 0x84, 0x38, 0xa6, 0x8a, 0x91, 0x32, 0x92, 0x9f, 0x91,
	0xa9, 0x85, 0xf2, 0x75, 0x62, 0xea, 0x8a, 0xd5, 0x89, 0xf9, 0xca, 0x97, 0xe1, 0x82, 0x73, 0xa1,
	0x3c, 0xe7, 0x89, 0xe2, 0xbc, 0x62, 0x9c, 0x27, 0xee, 0xf5, 0x74, 0x5a, 0x76, 0xf5, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x03, 0x61, 0xc0, 0xb7, 0x01, 0x00, 0x00,
}