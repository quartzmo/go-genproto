// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_parameter.proto

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

// An ad parameter that is used to update numeric values (such as prices or
// inventory levels) in any text line of an ad (including URLs). There can
// be a maximum of two AdParameters per ad group criterion. (One with
// parameter_index = 1 and one with parameter_index = 2.)
// In the ad the parameters are referenced by a placeholder of the form
// "{param#:value}". E.g. "{param1:$17}"
type AdParameter struct {
	// The resource name of the ad parameter.
	// Ad parameter resource names have the form:
	//
	// `customers/{customer_id}/adParameters/{ad_group_id}~{criterion_id}~{parameter_index}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group criterion that this ad parameter belongs to.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The unique index of this ad parameter. Must be either 1 or 2.
	ParameterIndex *wrappers.Int64Value `protobuf:"bytes,3,opt,name=parameter_index,json=parameterIndex,proto3" json:"parameter_index,omitempty"`
	// Numeric value to insert into the ad text. The following restrictions
	//  apply:
	//  - Can use comma or period as a separator, with an optional period or
	//    comma (respectively) for fractional values. For example, 1,000,000.00
	//    and 2.000.000,10 are valid.
	//  - Can be prepended or appended with a currency symbol. For example,
	//    $99.99 is valid.
	//  - Can be prepended or appended with a currency code. For example, 99.99USD
	//    and EUR200 are valid.
	//  - Can use '%'. For example, 1.0% and 1,0% are valid.
	//  - Can use plus or minus. For example, -10.99 and 25+ are valid.
	//  - Can use '/' between two numbers. For example 4/1 and 0.95/0.45 are
	//    valid.
	InsertionText        *wrappers.StringValue `protobuf:"bytes,4,opt,name=insertion_text,json=insertionText,proto3" json:"insertion_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdParameter) Reset()         { *m = AdParameter{} }
func (m *AdParameter) String() string { return proto.CompactTextString(m) }
func (*AdParameter) ProtoMessage()    {}
func (*AdParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba62669d0fe2a117, []int{0}
}

func (m *AdParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdParameter.Unmarshal(m, b)
}
func (m *AdParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdParameter.Marshal(b, m, deterministic)
}
func (m *AdParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdParameter.Merge(m, src)
}
func (m *AdParameter) XXX_Size() int {
	return xxx_messageInfo_AdParameter.Size(m)
}
func (m *AdParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_AdParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AdParameter proto.InternalMessageInfo

func (m *AdParameter) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdParameter) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *AdParameter) GetParameterIndex() *wrappers.Int64Value {
	if m != nil {
		return m.ParameterIndex
	}
	return nil
}

func (m *AdParameter) GetInsertionText() *wrappers.StringValue {
	if m != nil {
		return m.InsertionText
	}
	return nil
}

func init() {
	proto.RegisterType((*AdParameter)(nil), "google.ads.googleads.v2.resources.AdParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_parameter.proto", fileDescriptor_ba62669d0fe2a117)
}

var fileDescriptor_ba62669d0fe2a117 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0xb1, 0x33, 0x06, 0x73, 0x96, 0x2c, 0xf8, 0xca, 0x64, 0x61, 0x24, 0x1b, 0x81, 0x5c,
	0xc9, 0xe0, 0x85, 0x5d, 0x68, 0x57, 0x4e, 0x06, 0x21, 0xb9, 0x18, 0x21, 0x2d, 0xbe, 0x28, 0x06,
	0xa3, 0x44, 0xaa, 0x10, 0xc4, 0x92, 0x91, 0xe4, 0x34, 0xaf, 0xd0, 0xd7, 0xe8, 0x65, 0x1f, 0xa5,
	0x8f, 0xd2, 0x77, 0x28, 0x14, 0xff, 0x91, 0x28, 0x14, 0xda, 0xde, 0x1d, 0xa4, 0xf3, 0x3b, 0xdf,
	0x91, 0x3e, 0x6f, 0x4e, 0x85, 0xa0, 0x47, 0x12, 0x22, 0xac, 0xc2, 0x46, 0x56, 0xea, 0x14, 0x85,
	0x92, 0x28, 0x51, 0xca, 0x03, 0x51, 0x21, 0xc2, 0x59, 0x81, 0x24, 0xca, 0x89, 0x26, 0x12, 0x14,
	0x52, 0x68, 0xe1, 0x4f, 0x1a, 0x2b, 0x40, 0x58, 0x01, 0x4b, 0x81, 0x53, 0x04, 0x2c, 0x35, 0xfc,
	0xd1, 0x06, 0xd7, 0xc0, 0xbe, 0xbc, 0x0e, 0x6f, 0x24, 0x2a, 0x0a, 0x22, 0x55, 0x13, 0x31, 0x1c,
	0x99, 0xc1, 0x05, 0x0b, 0x11, 0xe7, 0x42, 0x23, 0xcd, 0x04, 0x6f, 0x6f, 0x7f, 0xde, 0xba, 0x5e,
	0x37, 0xc6, 0x5b, 0x33, 0xd6, 0xff, 0xe5, 0xf5, 0x4c, 0x74, 0xc6, 0x51, 0x4e, 0x02, 0x67, 0xec,
	0xcc, 0xbe, 0xec, 0xbe, 0x9a, 0xc3, 0xff, 0x28, 0x27, 0xfe, 0xc6, 0xf3, 0x11, 0xce, 0xa8, 0x14,
	0x65, 0x91, 0x1d, 0x24, 0xd3, 0x44, 0x32, 0xc1, 0x03, 0x77, 0xec, 0xcc, 0xba, 0xd1, 0xa8, 0xed,
	0x09, 0x4c, 0x1f, 0x70, 0xa1, 0x25, 0xe3, 0x34, 0x41, 0xc7, 0x92, 0xec, 0x06, 0x08, 0xaf, 0x2a,
	0x6c, 0x69, 0x28, 0xff, 0x9f, 0xf7, 0xcd, 0x3e, 0x3a, 0x63, 0x1c, 0x93, 0x73, 0xd0, 0xa9, 0x83,
	0xbe, 0xbf, 0x0a, 0x5a, 0x73, 0xfd, 0x67, 0xde, 0xe4, 0xf4, 0x2d, 0xb3, 0xae, 0x10, 0x7f, 0xe9,
	0xf5, 0x19, 0x57, 0x44, 0x56, 0x4f, 0xcb, 0x34, 0x39, 0xeb, 0xe0, 0xd3, 0x07, 0xda, 0xf4, 0x2c,
	0x73, 0x49, 0xce, 0x7a, 0xf1, 0xe4, 0x78, 0xd3, 0x83, 0xc8, 0xc1, 0xbb, 0x7f, 0xbe, 0x18, 0xbc,
	0xf8, 0xb2, 0x6d, 0x95, 0xbc, 0x75, 0xae, 0x36, 0x2d, 0x46, 0xc5, 0x11, 0x71, 0x0a, 0x84, 0xa4,
	0x21, 0x25, 0xbc, 0x9e, 0x6b, 0x16, 0x5e, 0x30, 0xf5, 0xc6, 0xfe, 0xff, 0x5a, 0x75, 0xe7, 0x76,
	0x56, 0x71, 0x7c, 0xef, 0x4e, 0x56, 0x4d, 0x64, 0x8c, 0x15, 0x68, 0x64, 0xa5, 0x92, 0x08, 0xec,
	0x8c, 0xf3, 0xc1, 0x78, 0xd2, 0x18, 0xab, 0xd4, 0x7a, 0xd2, 0x24, 0x4a, 0xad, 0xe7, 0xd1, 0x9d,
	0x36, 0x17, 0x10, 0xc6, 0x58, 0x41, 0x68, 0x5d, 0x10, 0x26, 0x11, 0x84, 0xd6, 0xb7, 0xff, 0x5c,
	0x97, 0xfd, 0xfd, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x49, 0x21, 0x5b, 0x0c, 0xab, 0x02, 0x00, 0x00,
}