// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/targeting_setting.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// Settings for the targeting-related features, at the campaign and ad group
// levels. For more details about the targeting setting, visit
// https://support.google.com/google-ads/answer/7365594
type TargetingSetting struct {
	// The per-targeting-dimension setting to restrict the reach of your campaign
	// or ad group.
	TargetRestrictions   []*TargetRestriction `protobuf:"bytes,1,rep,name=target_restrictions,json=targetRestrictions,proto3" json:"target_restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetingSetting) Reset()         { *m = TargetingSetting{} }
func (m *TargetingSetting) String() string { return proto.CompactTextString(m) }
func (*TargetingSetting) ProtoMessage()    {}
func (*TargetingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ec7b50ceb43d68, []int{0}
}

func (m *TargetingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetingSetting.Unmarshal(m, b)
}
func (m *TargetingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetingSetting.Marshal(b, m, deterministic)
}
func (m *TargetingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetingSetting.Merge(m, src)
}
func (m *TargetingSetting) XXX_Size() int {
	return xxx_messageInfo_TargetingSetting.Size(m)
}
func (m *TargetingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_TargetingSetting proto.InternalMessageInfo

func (m *TargetingSetting) GetTargetRestrictions() []*TargetRestriction {
	if m != nil {
		return m.TargetRestrictions
	}
	return nil
}

// The list of per-targeting-dimension targeting settings.
type TargetRestriction struct {
	// The targeting dimension that these settings apply to.
	TargetingDimension enums.TargetingDimensionEnum_TargetingDimension `protobuf:"varint,1,opt,name=targeting_dimension,json=targetingDimension,proto3,enum=google.ads.googleads.v2.enums.TargetingDimensionEnum_TargetingDimension" json:"targeting_dimension,omitempty"`
	// Indicates whether to restrict your ads to show only for the criteria you
	// have selected for this targeting_dimension, or to target all values for
	// this targeting_dimension and show ads based on your targeting in other
	// TargetingDimensions. A value of `true` means that these criteria will only
	// apply bid modifiers, and not affect targeting. A value of `false` means
	// that these criteria will restrict targeting as well as applying bid
	// modifiers.
	BidOnly              *wrappers.BoolValue `protobuf:"bytes,2,opt,name=bid_only,json=bidOnly,proto3" json:"bid_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TargetRestriction) Reset()         { *m = TargetRestriction{} }
func (m *TargetRestriction) String() string { return proto.CompactTextString(m) }
func (*TargetRestriction) ProtoMessage()    {}
func (*TargetRestriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_73ec7b50ceb43d68, []int{1}
}

func (m *TargetRestriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetRestriction.Unmarshal(m, b)
}
func (m *TargetRestriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetRestriction.Marshal(b, m, deterministic)
}
func (m *TargetRestriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetRestriction.Merge(m, src)
}
func (m *TargetRestriction) XXX_Size() int {
	return xxx_messageInfo_TargetRestriction.Size(m)
}
func (m *TargetRestriction) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetRestriction.DiscardUnknown(m)
}

var xxx_messageInfo_TargetRestriction proto.InternalMessageInfo

func (m *TargetRestriction) GetTargetingDimension() enums.TargetingDimensionEnum_TargetingDimension {
	if m != nil {
		return m.TargetingDimension
	}
	return enums.TargetingDimensionEnum_UNSPECIFIED
}

func (m *TargetRestriction) GetBidOnly() *wrappers.BoolValue {
	if m != nil {
		return m.BidOnly
	}
	return nil
}

func init() {
	proto.RegisterType((*TargetingSetting)(nil), "google.ads.googleads.v2.common.TargetingSetting")
	proto.RegisterType((*TargetRestriction)(nil), "google.ads.googleads.v2.common.TargetRestriction")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/targeting_setting.proto", fileDescriptor_73ec7b50ceb43d68)
}

var fileDescriptor_73ec7b50ceb43d68 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xb1, 0x03, 0xdb, 0x70, 0x60, 0x6c, 0x1e, 0x83, 0x10, 0x46, 0x08, 0x3e, 0xe5, 0x24,
	0x31, 0x8d, 0x6d, 0xa0, 0x9d, 0x9c, 0x65, 0x64, 0xb7, 0x85, 0x6c, 0xf8, 0x30, 0x0c, 0x41, 0x8e,
	0x54, 0x21, 0xb0, 0x25, 0x63, 0xc9, 0x29, 0xf9, 0x3a, 0x3d, 0xf6, 0x7b, 0xf4, 0xd2, 0x8f, 0xd2,
	0x53, 0x3f, 0x42, 0xb1, 0x65, 0xb9, 0x69, 0xda, 0xb4, 0x27, 0xbd, 0x7a, 0xf5, 0x3c, 0xbf, 0xf7,
	0x0f, 0x0a, 0xbe, 0x71, 0xa5, 0x78, 0xce, 0x20, 0xa1, 0x1a, 0xda, 0xb0, 0x89, 0x76, 0x08, 0x6e,
	0x55, 0x51, 0x28, 0x09, 0x0d, 0xa9, 0x38, 0x33, 0x42, 0xf2, 0x8d, 0x66, 0xa6, 0x39, 0x41, 0x59,
	0x29, 0xa3, 0xc2, 0x89, 0x15, 0x03, 0x42, 0x35, 0xe8, 0x7d, 0x60, 0x87, 0x80, 0xf5, 0x8d, 0xbf,
	0x9f, 0xe2, 0x32, 0x59, 0x17, 0xfa, 0x00, 0x4b, 0x45, 0xc1, 0xa4, 0x16, 0x4a, 0x5a, 0xf0, 0xb8,
	0x03, 0xc3, 0xf6, 0x96, 0xd5, 0x67, 0xf0, 0xbc, 0x22, 0x65, 0xc9, 0x2a, 0xdd, 0xbd, 0x7f, 0x72,
	0xe0, 0x52, 0x40, 0x22, 0xa5, 0x32, 0xc4, 0x08, 0x25, 0xbb, 0xd7, 0x68, 0x17, 0xbc, 0xfb, 0xe7,
	0xd0, 0x7f, 0x6d, 0xc3, 0x61, 0x16, 0x7c, 0xb0, 0xe5, 0x36, 0x15, 0xd3, 0xa6, 0x12, 0xdb, 0xd6,
	0x30, 0xf2, 0xa6, 0x83, 0xd9, 0x10, 0x7d, 0x06, 0xcf, 0x0f, 0x02, 0x2c, 0x6e, 0x7d, 0xef, 0x5c,
	0x87, 0xe6, 0x38, 0xa5, 0xa3, 0x2b, 0x2f, 0x78, 0xff, 0x48, 0x19, 0xee, 0x5d, 0xe5, 0x07, 0x83,
	0x8e, 0xbc, 0xa9, 0x37, 0x7b, 0x8b, 0x7e, 0x9f, 0xac, 0xdc, 0xae, 0x08, 0xf4, 0x73, 0x2c, 0x9c,
	0xf1, 0x97, 0xac, 0x8b, 0x27, 0xd2, 0xae, 0xa1, 0xc3, 0x5c, 0xf8, 0x35, 0x78, 0x93, 0x09, 0xba,
	0x51, 0x32, 0xdf, 0x8f, 0xfc, 0xa9, 0x37, 0x1b, 0xa2, 0xb1, 0xab, 0xe7, 0x36, 0x0b, 0xe6, 0x4a,
	0xe5, 0x09, 0xc9, 0x6b, 0xb6, 0x7e, 0x9d, 0x09, 0xfa, 0x47, 0xe6, 0xfb, 0xf9, 0xad, 0x17, 0x44,
	0x5b, 0x55, 0xbc, 0xb0, 0x94, 0xf9, 0xc7, 0xe3, 0x25, 0xaf, 0x1a, 0xe6, 0xca, 0xfb, 0xbf, 0xe8,
	0x8c, 0x5c, 0xe5, 0x44, 0x72, 0xa0, 0x2a, 0x0e, 0x39, 0x93, 0x6d, 0x45, 0xf7, 0x0d, 0x4a, 0xa1,
	0x4f, 0xfd, 0xb6, 0x1f, 0xf6, 0xb8, 0xf0, 0x07, 0xcb, 0x38, 0xbe, 0xf4, 0x27, 0x4b, 0x0b, 0x8b,
	0xa9, 0x06, 0x36, 0x6c, 0xa2, 0x04, 0x81, 0x9f, 0xad, 0xec, 0xda, 0x09, 0xd2, 0x98, 0xea, 0xb4,
	0x17, 0xa4, 0x09, 0x4a, 0xad, 0xe0, 0xc6, 0x8f, 0x6c, 0x16, 0xe3, 0x98, 0x6a, 0x8c, 0x7b, 0x09,
	0xc6, 0x09, 0xc2, 0xd8, 0x8a, 0xb2, 0x57, 0x6d, 0x77, 0x5f, 0xee, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x9c, 0x9b, 0x31, 0x6d, 0x0a, 0x03, 0x00, 0x00,
}