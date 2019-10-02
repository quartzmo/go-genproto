// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_label.proto

package resources

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

// A relationship between an ad group and a label.
type AdGroupLabel struct {
	// The resource name of the ad group label.
	// Ad group label resource names have the form:
	// `customers/{customer_id}/adGroupLabels/{ad_group_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group to which the label is attached.
	AdGroup *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The label assigned to the ad group.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdGroupLabel) Reset()         { *m = AdGroupLabel{} }
func (m *AdGroupLabel) String() string { return proto.CompactTextString(m) }
func (*AdGroupLabel) ProtoMessage()    {}
func (*AdGroupLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c126c30fd91d977, []int{0}
}

func (m *AdGroupLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupLabel.Unmarshal(m, b)
}
func (m *AdGroupLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupLabel.Marshal(b, m, deterministic)
}
func (m *AdGroupLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupLabel.Merge(m, src)
}
func (m *AdGroupLabel) XXX_Size() int {
	return xxx_messageInfo_AdGroupLabel.Size(m)
}
func (m *AdGroupLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupLabel.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupLabel proto.InternalMessageInfo

func (m *AdGroupLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupLabel) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupLabel)(nil), "google.ads.googleads.v2.resources.AdGroupLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_label.proto", fileDescriptor_4c126c30fd91d977)
}

var fileDescriptor_4c126c30fd91d977 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0xc7, 0x69, 0xc7, 0xf3, 0xa8, 0x75, 0x1e, 0xec, 0x69, 0x8c, 0x21, 0x9b, 0x32, 0xd8, 0x29,
	0x81, 0x08, 0x0a, 0xf1, 0xd4, 0x5d, 0x06, 0x22, 0x32, 0x26, 0xf4, 0x20, 0x85, 0xf1, 0xdb, 0x12,
	0x43, 0xa1, 0x4b, 0x4a, 0xd2, 0xce, 0xb7, 0x23, 0x1e, 0x7d, 0x29, 0xbe, 0x14, 0x5f, 0x84, 0x48,
	0x9b, 0x26, 0x78, 0x52, 0x6f, 0x5f, 0x9a, 0xcf, 0xf7, 0x4f, 0x9a, 0xe8, 0x4a, 0x28, 0x25, 0x0a,
	0x8e, 0x81, 0x19, 0x6c, 0x65, 0xa3, 0xf6, 0x04, 0x6b, 0x6e, 0x54, 0xad, 0xb7, 0xdc, 0x60, 0x60,
	0x6b, 0xa1, 0x55, 0x5d, 0xae, 0x0b, 0xd8, 0xf0, 0x02, 0x95, 0x5a, 0x55, 0x2a, 0x9e, 0x58, 0x18,
	0x01, 0x33, 0xc8, 0xfb, 0xd0, 0x9e, 0x20, 0xef, 0x1b, 0x9e, 0x75, 0xd1, 0xad, 0x61, 0x53, 0x3f,
	0xe1, 0x67, 0x0d, 0x65, 0xc9, 0xb5, 0xb1, 0x11, 0xc3, 0x91, 0xab, 0x2e, 0x73, 0x0c, 0x52, 0xaa,
	0x0a, 0xaa, 0x5c, 0xc9, 0xee, 0xf4, 0xfc, 0x25, 0x88, 0xfa, 0x09, 0x5b, 0x34, 0xc5, 0x77, 0x4d,
	0x6f, 0x7c, 0x11, 0x9d, 0xb8, 0xec, 0xb5, 0x84, 0x1d, 0x1f, 0x04, 0xe3, 0x60, 0x76, 0xb4, 0xea,
	0xbb, 0x8f, 0xf7, 0xb0, 0xe3, 0xf1, 0x75, 0x74, 0xe8, 0xe6, 0x0e, 0xc2, 0x71, 0x30, 0x3b, 0x26,
	0xa3, 0x6e, 0x1e, 0x72, 0x33, 0xd0, 0x43, 0xa5, 0x73, 0x29, 0x52, 0x28, 0x6a, 0xbe, 0x3a, 0x00,
	0x5b, 0x11, 0x93, 0xe8, 0x5f, 0x7b, 0xbd, 0x41, 0xef, 0x0f, 0x2e, 0x8b, 0xce, 0x3f, 0x83, 0x68,
	0xba, 0x55, 0x3b, 0xf4, 0xeb, 0xaf, 0x98, 0x9f, 0x7e, 0xbf, 0xc9, 0xb2, 0x89, 0x5c, 0x06, 0x8f,
	0xb7, 0x9d, 0x4f, 0xa8, 0x02, 0xa4, 0x40, 0x4a, 0x0b, 0x2c, 0xb8, 0x6c, 0x0b, 0xdd, 0x53, 0x94,
	0xb9, 0xf9, 0xe1, 0x65, 0x6e, 0xbc, 0x7a, 0x0d, 0x7b, 0x8b, 0x24, 0x79, 0x0b, 0x27, 0x0b, 0x1b,
	0x99, 0x30, 0x83, 0xac, 0x6c, 0x54, 0x4a, 0xd0, 0xca, 0x91, 0xef, 0x8e, 0xc9, 0x12, 0x66, 0x32,
	0xcf, 0x64, 0x29, 0xc9, 0x3c, 0xf3, 0x11, 0x4e, 0xed, 0x01, 0xa5, 0x09, 0x33, 0x94, 0x7a, 0x8a,
	0xd2, 0x94, 0x50, 0xea, 0xb9, 0xcd, 0xff, 0x76, 0xec, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x43, 0x31, 0x9a, 0x45, 0x02, 0x00, 0x00,
}