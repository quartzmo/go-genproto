// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/custom_interest.proto

package resources

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

// A custom interest. This is a list of users by interest.
type CustomInterest struct {
	// The resource name of the custom interest.
	// Custom interest resource names have the form:
	//
	// `customers/{customer_id}/customInterests/{custom_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the custom interest.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Status of this custom interest. Indicates whether the custom interest is
	// enabled or removed.
	Status enums.CustomInterestStatusEnum_CustomInterestStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CustomInterestStatusEnum_CustomInterestStatus" json:"status,omitempty"`
	// Name of the custom interest. It should be unique across the same custom
	// affinity audience.
	// This field is required for create operations.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the custom interest, CUSTOM_AFFINITY or CUSTOM_INTENT.
	// By default the type is set to CUSTOM_AFFINITY.
	Type enums.CustomInterestTypeEnum_CustomInterestType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CustomInterestTypeEnum_CustomInterestType" json:"type,omitempty"`
	// Description of this custom interest audience.
	Description *wrappers.StringValue `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// List of custom interest members that this custom interest is composed of.
	// Members can be added during CustomInterest creation. If members are
	// presented in UPDATE operation, existing members will be overridden.
	Members              []*CustomInterestMember `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CustomInterest) Reset()         { *m = CustomInterest{} }
func (m *CustomInterest) String() string { return proto.CompactTextString(m) }
func (*CustomInterest) ProtoMessage()    {}
func (*CustomInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b9ef39246e5a8dc, []int{0}
}

func (m *CustomInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterest.Unmarshal(m, b)
}
func (m *CustomInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterest.Marshal(b, m, deterministic)
}
func (m *CustomInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterest.Merge(m, src)
}
func (m *CustomInterest) XXX_Size() int {
	return xxx_messageInfo_CustomInterest.Size(m)
}
func (m *CustomInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterest proto.InternalMessageInfo

func (m *CustomInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomInterest) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CustomInterest) GetStatus() enums.CustomInterestStatusEnum_CustomInterestStatus {
	if m != nil {
		return m.Status
	}
	return enums.CustomInterestStatusEnum_UNSPECIFIED
}

func (m *CustomInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CustomInterest) GetType() enums.CustomInterestTypeEnum_CustomInterestType {
	if m != nil {
		return m.Type
	}
	return enums.CustomInterestTypeEnum_UNSPECIFIED
}

func (m *CustomInterest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CustomInterest) GetMembers() []*CustomInterestMember {
	if m != nil {
		return m.Members
	}
	return nil
}

// A member of custom interest audience. A member can be a keyword or url.
// It is immutable, that is, it can only be created or removed but not changed.
type CustomInterestMember struct {
	// The type of custom interest member, KEYWORD or URL.
	MemberType enums.CustomInterestMemberTypeEnum_CustomInterestMemberType `protobuf:"varint,1,opt,name=member_type,json=memberType,proto3,enum=google.ads.googleads.v2.enums.CustomInterestMemberTypeEnum_CustomInterestMemberType" json:"member_type,omitempty"`
	// Keyword text when member_type is KEYWORD or URL string when
	// member_type is URL.
	Parameter            *wrappers.StringValue `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomInterestMember) Reset()         { *m = CustomInterestMember{} }
func (m *CustomInterestMember) String() string { return proto.CompactTextString(m) }
func (*CustomInterestMember) ProtoMessage()    {}
func (*CustomInterestMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b9ef39246e5a8dc, []int{1}
}

func (m *CustomInterestMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestMember.Unmarshal(m, b)
}
func (m *CustomInterestMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestMember.Marshal(b, m, deterministic)
}
func (m *CustomInterestMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestMember.Merge(m, src)
}
func (m *CustomInterestMember) XXX_Size() int {
	return xxx_messageInfo_CustomInterestMember.Size(m)
}
func (m *CustomInterestMember) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestMember.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestMember proto.InternalMessageInfo

func (m *CustomInterestMember) GetMemberType() enums.CustomInterestMemberTypeEnum_CustomInterestMemberType {
	if m != nil {
		return m.MemberType
	}
	return enums.CustomInterestMemberTypeEnum_UNSPECIFIED
}

func (m *CustomInterestMember) GetParameter() *wrappers.StringValue {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomInterest)(nil), "google.ads.googleads.v2.resources.CustomInterest")
	proto.RegisterType((*CustomInterestMember)(nil), "google.ads.googleads.v2.resources.CustomInterestMember")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/custom_interest.proto", fileDescriptor_5b9ef39246e5a8dc)
}

var fileDescriptor_5b9ef39246e5a8dc = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xd2, 0xd2, 0x69, 0x2e, 0xf4, 0xc1, 0xf0, 0x10, 0x8d, 0x09, 0x75, 0x43, 0x93, 0x2a,
	0x21, 0x39, 0x28, 0x20, 0x86, 0x8c, 0x04, 0xca, 0x10, 0x1a, 0x43, 0x80, 0x46, 0x36, 0xf5, 0x01,
	0x45, 0xaa, 0xdc, 0xc6, 0x44, 0x91, 0x6a, 0x3b, 0xb2, 0x9d, 0xa1, 0xbd, 0xf1, 0x2d, 0x3c, 0xf2,
	0x27, 0xf0, 0x29, 0xf0, 0x13, 0x28, 0x76, 0x9c, 0x75, 0xac, 0xdd, 0xe8, 0xdb, 0xad, 0xef, 0x39,
	0xe7, 0x1e, 0x1f, 0xdf, 0x06, 0xec, 0xe7, 0x42, 0xe4, 0x73, 0x1a, 0x92, 0x4c, 0x85, 0xb6, 0xac,
	0xab, 0xb3, 0x28, 0x94, 0x54, 0x89, 0x4a, 0xce, 0xa8, 0x0a, 0x67, 0x95, 0xd2, 0x82, 0x4d, 0x0a,
	0xae, 0xa9, 0xa4, 0x4a, 0xa3, 0x52, 0x0a, 0x2d, 0xe0, 0x8e, 0x45, 0x23, 0x92, 0x29, 0xd4, 0x12,
	0xd1, 0x59, 0x84, 0x5a, 0xe2, 0xd6, 0xab, 0x55, 0xda, 0x94, 0x57, 0xec, 0x8a, 0xee, 0x84, 0x51,
	0x36, 0xa5, 0x72, 0xa2, 0xcf, 0x4b, 0x6a, 0x67, 0x6c, 0xe1, 0xf5, 0x04, 0x94, 0x26, 0xba, 0x52,
	0x0d, 0xf7, 0xf9, 0x7a, 0xdc, 0x85, 0xa9, 0x0f, 0x1a, 0xa6, 0xf9, 0x35, 0xad, 0xbe, 0x84, 0x5f,
	0x25, 0x29, 0x4b, 0x2a, 0x9d, 0xf2, 0xb6, 0x53, 0x2e, 0x8b, 0x90, 0x70, 0x2e, 0x34, 0xd1, 0x85,
	0xe0, 0x4d, 0x77, 0xf7, 0x4f, 0x07, 0x0c, 0x5e, 0x1b, 0xf1, 0xa3, 0x46, 0x1b, 0x3e, 0x04, 0x77,
	0x5c, 0x28, 0x13, 0x4e, 0x18, 0x0d, 0xbc, 0xa1, 0x37, 0xda, 0x4c, 0x6e, 0xbb, 0xc3, 0x8f, 0x84,
	0x51, 0xf8, 0x08, 0xf8, 0x45, 0x16, 0xf8, 0x43, 0x6f, 0xd4, 0x8f, 0xee, 0x37, 0x89, 0x22, 0x67,
	0x01, 0x1d, 0x71, 0xfd, 0xec, 0xe9, 0x98, 0xcc, 0x2b, 0x9a, 0xf8, 0x45, 0x06, 0x33, 0xd0, 0xb3,
	0x97, 0x0d, 0x3a, 0x43, 0x6f, 0x34, 0x88, 0xde, 0xa3, 0x55, 0xaf, 0x61, 0x6e, 0x8b, 0x2e, 0x1b,
	0x3a, 0x31, 0xd4, 0x37, 0xbc, 0x62, 0x4b, 0x1b, 0x49, 0xa3, 0x0d, 0x1f, 0x83, 0xae, 0xb1, 0xdb,
	0x35, 0xa6, 0xb6, 0xaf, 0x98, 0x3a, 0xd1, 0xb2, 0xe0, 0xb9, 0x75, 0x65, 0x90, 0x30, 0x05, 0xdd,
	0x3a, 0xc8, 0xe0, 0x96, 0x71, 0xf5, 0x76, 0x2d, 0x57, 0xa7, 0xe7, 0x25, 0x5d, 0xe2, 0xa9, 0x3e,
	0x4e, 0x8c, 0x2a, 0x7c, 0x09, 0xfa, 0x19, 0x55, 0x33, 0x59, 0x94, 0x75, 0xe0, 0x41, 0xef, 0x3f,
	0x6c, 0x2d, 0x12, 0xe0, 0x27, 0xb0, 0x61, 0x77, 0x4c, 0x05, 0x1b, 0xc3, 0xce, 0xa8, 0x1f, 0xed,
	0xa3, 0x1b, 0x97, 0xf8, 0x1f, 0x37, 0x1f, 0x0c, 0x3f, 0x71, 0x3a, 0xbb, 0x3f, 0x3d, 0x70, 0x6f,
	0x19, 0x02, 0x56, 0xa0, 0xbf, 0xb0, 0xcf, 0xe6, 0xc5, 0x07, 0xd1, 0xe9, 0x5a, 0x81, 0x58, 0xa5,
	0x15, 0xb1, 0x5c, 0x34, 0x13, 0xc0, 0xda, 0x1a, 0x62, 0xb0, 0x59, 0x12, 0x49, 0x18, 0xd5, 0x54,
	0x36, 0xcb, 0x74, 0x7d, 0x40, 0x17, 0xf0, 0x83, 0x6f, 0x3e, 0xd8, 0x9b, 0x09, 0x76, 0x73, 0x26,
	0x07, 0x77, 0x2f, 0x7b, 0x39, 0xae, 0x85, 0x8f, 0xbd, 0xcf, 0xef, 0x1a, 0x66, 0x2e, 0xe6, 0x84,
	0xe7, 0x48, 0xc8, 0x3c, 0xcc, 0x29, 0x37, 0x63, 0xdd, 0x5f, 0xb0, 0x2c, 0xd4, 0x35, 0x9f, 0x9a,
	0x17, 0x6d, 0xf5, 0xdd, 0xef, 0x1c, 0xc6, 0xf1, 0x0f, 0x7f, 0xe7, 0xd0, 0x4a, 0xc6, 0x99, 0x42,
	0xb6, 0xac, 0xab, 0x71, 0x84, 0x12, 0x87, 0xfc, 0xe5, 0x30, 0x69, 0x9c, 0xa9, 0xb4, 0xc5, 0xa4,
	0xe3, 0x28, 0x6d, 0x31, 0xbf, 0xfd, 0x3d, 0xdb, 0xc0, 0x38, 0xce, 0x14, 0xc6, 0x2d, 0x0a, 0xe3,
	0x71, 0x84, 0x71, 0x8b, 0x9b, 0xf6, 0x8c, 0xd9, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0,
	0x47, 0x77, 0x27, 0x16, 0x05, 0x00, 0x00,
}