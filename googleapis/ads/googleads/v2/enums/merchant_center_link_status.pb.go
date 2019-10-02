// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/merchant_center_link_status.proto

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

// Describes the possible statuses for a link between a Google Ads customer
// and a Google Merchant Center account.
type MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus int32

const (
	// Not specified.
	MerchantCenterLinkStatusEnum_UNSPECIFIED MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus = 0
	// Used for return value only. Represents value unknown in this version.
	MerchantCenterLinkStatusEnum_UNKNOWN MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus = 1
	// The link is enabled.
	MerchantCenterLinkStatusEnum_ENABLED MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus = 2
	// The link has no effect. It was proposed by the Merchant Center Account
	// owner and hasn't been confirmed by the customer.
	MerchantCenterLinkStatusEnum_PENDING MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus = 3
)

var MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PENDING",
}

var MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PENDING":     3,
}

func (x MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus) String() string {
	return proto.EnumName(MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus_name, int32(x))
}

func (MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0640355d09949a0f, []int{0, 0}
}

// Container for enum describing possible statuses of a Google Merchant Center
// link.
type MerchantCenterLinkStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MerchantCenterLinkStatusEnum) Reset()         { *m = MerchantCenterLinkStatusEnum{} }
func (m *MerchantCenterLinkStatusEnum) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLinkStatusEnum) ProtoMessage()    {}
func (*MerchantCenterLinkStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0640355d09949a0f, []int{0}
}

func (m *MerchantCenterLinkStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLinkStatusEnum.Unmarshal(m, b)
}
func (m *MerchantCenterLinkStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLinkStatusEnum.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLinkStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLinkStatusEnum.Merge(m, src)
}
func (m *MerchantCenterLinkStatusEnum) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLinkStatusEnum.Size(m)
}
func (m *MerchantCenterLinkStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLinkStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLinkStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus", MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus_name, MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus_value)
	proto.RegisterType((*MerchantCenterLinkStatusEnum)(nil), "google.ads.googleads.v2.enums.MerchantCenterLinkStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/merchant_center_link_status.proto", fileDescriptor_0640355d09949a0f)
}

var fileDescriptor_0640355d09949a0f = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x6a, 0xeb, 0x30,
	0x14, 0xfc, 0x71, 0xe0, 0x17, 0x94, 0x45, 0x4d, 0x56, 0xa5, 0x24, 0x8b, 0xe4, 0x00, 0x12, 0xb8,
	0x3b, 0x75, 0x51, 0xe4, 0x44, 0x0d, 0xa1, 0xa9, 0x6a, 0x1a, 0x92, 0x42, 0x31, 0x04, 0x35, 0x16,
	0xaa, 0x49, 0x2c, 0x05, 0x4b, 0xc9, 0x81, 0xba, 0xec, 0x51, 0x7a, 0x94, 0x2e, 0x7b, 0x82, 0x22,
	0x39, 0xf6, 0xce, 0xdd, 0x98, 0x79, 0x9e, 0xf7, 0x66, 0x46, 0x03, 0xee, 0xa4, 0xd6, 0x72, 0x2f,
	0x10, 0xcf, 0x0c, 0xaa, 0xa0, 0x43, 0xa7, 0x08, 0x09, 0x75, 0x2c, 0x0c, 0x2a, 0x44, 0xb9, 0x7d,
	0xe7, 0xca, 0x6e, 0xb6, 0x42, 0x59, 0x51, 0x6e, 0xf6, 0xb9, 0xda, 0x6d, 0x8c, 0xe5, 0xf6, 0x68,
	0xe0, 0xa1, 0xd4, 0x56, 0xf7, 0x87, 0xd5, 0x15, 0xe4, 0x99, 0x81, 0x8d, 0x00, 0x3c, 0x45, 0xd0,
	0x0b, 0x5c, 0x0f, 0x6a, 0xfd, 0x43, 0x8e, 0xb8, 0x52, 0xda, 0x72, 0x9b, 0x6b, 0x75, 0x3e, 0x1e,
	0x97, 0x60, 0xf0, 0x78, 0x76, 0x98, 0x78, 0x83, 0x45, 0xae, 0x76, 0x4b, 0x2f, 0x4f, 0xd5, 0xb1,
	0x18, 0x3f, 0x83, 0xab, 0x36, 0xbe, 0x7f, 0x09, 0x7a, 0x2b, 0xb6, 0x4c, 0xe8, 0x64, 0x7e, 0x3f,
	0xa7, 0xd3, 0xf0, 0x5f, 0xbf, 0x07, 0x2e, 0x56, 0xec, 0x81, 0x3d, 0xbd, 0xb0, 0xb0, 0xe3, 0x06,
	0xca, 0x48, 0xbc, 0xa0, 0xd3, 0x30, 0x70, 0x43, 0x42, 0xd9, 0x74, 0xce, 0x66, 0x61, 0x37, 0xfe,
	0xe9, 0x80, 0xd1, 0x56, 0x17, 0xf0, 0xcf, 0xdc, 0xf1, 0xb0, 0xcd, 0x37, 0x71, 0xc1, 0x93, 0xce,
	0x6b, 0x7c, 0xbe, 0x97, 0x7a, 0xcf, 0x95, 0x84, 0xba, 0x94, 0x48, 0x0a, 0xe5, 0x9f, 0x55, 0x17,
	0x79, 0xc8, 0x4d, 0x4b, 0xaf, 0xb7, 0xfe, 0xfb, 0x11, 0x74, 0x67, 0x84, 0x7c, 0x06, 0xc3, 0x59,
	0x25, 0x45, 0x32, 0x03, 0x2b, 0xe8, 0xd0, 0x3a, 0x82, 0xae, 0x03, 0xf3, 0x55, 0xf3, 0x29, 0xc9,
	0x4c, 0xda, 0xf0, 0xe9, 0x3a, 0x4a, 0x3d, 0xff, 0x1d, 0x8c, 0xaa, 0x9f, 0x18, 0x93, 0xcc, 0x60,
	0xdc, 0x6c, 0x60, 0xbc, 0x8e, 0x30, 0xf6, 0x3b, 0x6f, 0xff, 0x7d, 0xb0, 0x9b, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xee, 0xcf, 0x07, 0xa2, 0xef, 0x01, 0x00, 0x00,
}