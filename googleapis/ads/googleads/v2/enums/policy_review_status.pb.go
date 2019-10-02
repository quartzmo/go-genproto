// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/policy_review_status.proto

package enums

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

// The possible policy review statuses.
type PolicyReviewStatusEnum_PolicyReviewStatus int32

const (
	// No value has been specified.
	PolicyReviewStatusEnum_UNSPECIFIED PolicyReviewStatusEnum_PolicyReviewStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyReviewStatusEnum_UNKNOWN PolicyReviewStatusEnum_PolicyReviewStatus = 1
	// Currently under review.
	PolicyReviewStatusEnum_REVIEW_IN_PROGRESS PolicyReviewStatusEnum_PolicyReviewStatus = 2
	// Primary review complete. Other reviews may be continuing.
	PolicyReviewStatusEnum_REVIEWED PolicyReviewStatusEnum_PolicyReviewStatus = 3
	// The resource has been resubmitted for approval or its policy decision has
	// been appealed.
	PolicyReviewStatusEnum_UNDER_APPEAL PolicyReviewStatusEnum_PolicyReviewStatus = 4
	// The resource is eligible and may be serving but could still undergo
	// further review.
	PolicyReviewStatusEnum_ELIGIBLE_MAY_SERVE PolicyReviewStatusEnum_PolicyReviewStatus = 5
)

var PolicyReviewStatusEnum_PolicyReviewStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "REVIEW_IN_PROGRESS",
	3: "REVIEWED",
	4: "UNDER_APPEAL",
	5: "ELIGIBLE_MAY_SERVE",
}

var PolicyReviewStatusEnum_PolicyReviewStatus_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"REVIEW_IN_PROGRESS": 2,
	"REVIEWED":           3,
	"UNDER_APPEAL":       4,
	"ELIGIBLE_MAY_SERVE": 5,
}

func (x PolicyReviewStatusEnum_PolicyReviewStatus) String() string {
	return proto.EnumName(PolicyReviewStatusEnum_PolicyReviewStatus_name, int32(x))
}

func (PolicyReviewStatusEnum_PolicyReviewStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0a6b7274bed48e79, []int{0, 0}
}

// Container for enum describing possible policy review statuses.
type PolicyReviewStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyReviewStatusEnum) Reset()         { *m = PolicyReviewStatusEnum{} }
func (m *PolicyReviewStatusEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyReviewStatusEnum) ProtoMessage()    {}
func (*PolicyReviewStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6b7274bed48e79, []int{0}
}

func (m *PolicyReviewStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyReviewStatusEnum.Unmarshal(m, b)
}
func (m *PolicyReviewStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyReviewStatusEnum.Marshal(b, m, deterministic)
}
func (m *PolicyReviewStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyReviewStatusEnum.Merge(m, src)
}
func (m *PolicyReviewStatusEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyReviewStatusEnum.Size(m)
}
func (m *PolicyReviewStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyReviewStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyReviewStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.PolicyReviewStatusEnum_PolicyReviewStatus", PolicyReviewStatusEnum_PolicyReviewStatus_name, PolicyReviewStatusEnum_PolicyReviewStatus_value)
	proto.RegisterType((*PolicyReviewStatusEnum)(nil), "google.ads.googleads.v2.enums.PolicyReviewStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/policy_review_status.proto", fileDescriptor_0a6b7274bed48e79)
}

var fileDescriptor_0a6b7274bed48e79 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xb6, 0x9d, 0xff, 0xc8, 0x06, 0x96, 0x1c, 0x26, 0x88, 0x3b, 0x6c, 0x0f, 0x90, 0x42, 0xbd,
	0x48, 0x3c, 0xa5, 0x2e, 0x96, 0xe2, 0xec, 0x4a, 0xcb, 0x3a, 0x94, 0x42, 0xa8, 0x6b, 0x29, 0x85,
	0x2d, 0x29, 0x4b, 0x37, 0xf1, 0xea, 0x7b, 0x78, 0xf1, 0xe8, 0xa3, 0xf8, 0x28, 0xe2, 0x43, 0x48,
	0xd3, 0x6d, 0x97, 0xa1, 0x97, 0xf0, 0x91, 0xef, 0x0f, 0xbf, 0xef, 0x03, 0xd7, 0xb9, 0x10, 0xf9,
	0x3c, 0x33, 0x93, 0x54, 0x9a, 0x0d, 0xac, 0xd1, 0xda, 0x32, 0x33, 0xbe, 0x5a, 0x48, 0xb3, 0x14,
	0xf3, 0x62, 0xf6, 0xca, 0x96, 0xd9, 0xba, 0xc8, 0x5e, 0x98, 0xac, 0x92, 0x6a, 0x25, 0x51, 0xb9,
	0x14, 0x95, 0x80, 0xbd, 0x46, 0x8e, 0x92, 0x54, 0xa2, 0x9d, 0x13, 0xad, 0x2d, 0xa4, 0x9c, 0x17,
	0x97, 0xdb, 0xe0, 0xb2, 0x30, 0x13, 0xce, 0x45, 0x95, 0x54, 0x85, 0xe0, 0x1b, 0xf3, 0xe0, 0x5d,
	0x03, 0x5d, 0x5f, 0x65, 0x07, 0x2a, 0x3a, 0x54, 0xc9, 0x94, 0xaf, 0x16, 0x83, 0x37, 0x0d, 0xc0,
	0x7d, 0x0a, 0x9e, 0x81, 0xf6, 0xc4, 0x0b, 0x7d, 0x7a, 0xeb, 0xde, 0xb9, 0x74, 0x68, 0x1c, 0xc0,
	0x36, 0x38, 0x99, 0x78, 0xf7, 0xde, 0x78, 0xea, 0x19, 0x1a, 0xec, 0x02, 0x18, 0xd0, 0xc8, 0xa5,
	0x53, 0xe6, 0x7a, 0xcc, 0x0f, 0xc6, 0x4e, 0x40, 0xc3, 0xd0, 0xd0, 0x61, 0x07, 0x9c, 0x36, 0xff,
	0x74, 0x68, 0xb4, 0xa0, 0x01, 0x3a, 0x13, 0x6f, 0x48, 0x03, 0x46, 0x7c, 0x9f, 0x92, 0x91, 0x71,
	0x58, 0xfb, 0xe8, 0xc8, 0x75, 0x5c, 0x7b, 0x44, 0xd9, 0x03, 0x79, 0x64, 0x21, 0x0d, 0x22, 0x6a,
	0x1c, 0xd9, 0x3f, 0x1a, 0xe8, 0xcf, 0xc4, 0x02, 0xfd, 0xdb, 0xd1, 0x3e, 0xdf, 0xbf, 0xd3, 0xaf,
	0xeb, 0xf9, 0xda, 0x93, 0xbd, 0x71, 0xe6, 0x62, 0x9e, 0xf0, 0x1c, 0x89, 0x65, 0x6e, 0xe6, 0x19,
	0x57, 0xe5, 0xb7, 0x3b, 0x97, 0x85, 0xfc, 0x63, 0xf6, 0x1b, 0xf5, 0x7e, 0xe8, 0x2d, 0x87, 0x90,
	0x4f, 0xbd, 0xe7, 0x34, 0x51, 0x24, 0x95, 0xa8, 0x81, 0x35, 0x8a, 0x2c, 0x54, 0xcf, 0x25, 0xbf,
	0xb6, 0x7c, 0x4c, 0x52, 0x19, 0xef, 0xf8, 0x38, 0xb2, 0x62, 0xc5, 0x7f, 0xeb, 0xfd, 0xe6, 0x13,
	0x63, 0x92, 0x4a, 0x8c, 0x77, 0x0a, 0x8c, 0x23, 0x0b, 0x63, 0xa5, 0x79, 0x3e, 0x56, 0x87, 0x5d,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x97, 0x75, 0x88, 0x0e, 0x02, 0x00, 0x00,
}