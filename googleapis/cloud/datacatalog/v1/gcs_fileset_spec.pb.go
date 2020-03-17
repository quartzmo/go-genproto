// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1/gcs_fileset_spec.proto

package datacatalog

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Describes a Cloud Storage fileset entry.
type GcsFilesetSpec struct {
	// Required. Patterns to identify a set of files in Google Cloud Storage. See [Cloud
	// Storage documentation](/storage/docs/gsutil/addlhelp/WildcardNames) for
	// more information. Note that bucket wildcards are currently not supported.
	//
	// Examples of valid file_patterns:
	//
	//  * `gs://bucket_name/dir/*`: matches all files within `bucket_name/dir`
	//                              directory.
	//  * `gs://bucket_name/dir/**`: matches all files in `bucket_name/dir`
	//                               spanning all subdirectories.
	//  * `gs://bucket_name/file*`: matches files prefixed by `file` in
	//                              `bucket_name`
	//  * `gs://bucket_name/??.txt`: matches files with two characters followed by
	//                               `.txt` in `bucket_name`
	//  * `gs://bucket_name/[aeiou].txt`: matches files that contain a single
	//                                    vowel character followed by `.txt` in
	//                                    `bucket_name`
	//  * `gs://bucket_name/[a-m].txt`: matches files that contain `a`, `b`, ...
	//                                  or `m` followed by `.txt` in `bucket_name`
	//  * `gs://bucket_name/a/*/b`: matches all files in `bucket_name` that match
	//                              `a/*/b` pattern, such as `a/c/b`, `a/d/b`
	//  * `gs://another_bucket/a.txt`: matches `gs://another_bucket/a.txt`
	//
	// You can combine wildcards to provide more powerful matches, for example:
	//
	//  * `gs://bucket_name/[a-m]??.j*g`
	FilePatterns []string `protobuf:"bytes,1,rep,name=file_patterns,json=filePatterns,proto3" json:"file_patterns,omitempty"`
	// Output only. Sample files contained in this fileset, not all files
	// contained in this fileset are represented here.
	SampleGcsFileSpecs   []*GcsFileSpec `protobuf:"bytes,2,rep,name=sample_gcs_file_specs,json=sampleGcsFileSpecs,proto3" json:"sample_gcs_file_specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GcsFilesetSpec) Reset()         { *m = GcsFilesetSpec{} }
func (m *GcsFilesetSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFilesetSpec) ProtoMessage()    {}
func (*GcsFilesetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef0234cfb0655d3, []int{0}
}

func (m *GcsFilesetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFilesetSpec.Unmarshal(m, b)
}
func (m *GcsFilesetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFilesetSpec.Marshal(b, m, deterministic)
}
func (m *GcsFilesetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFilesetSpec.Merge(m, src)
}
func (m *GcsFilesetSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFilesetSpec.Size(m)
}
func (m *GcsFilesetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFilesetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFilesetSpec proto.InternalMessageInfo

func (m *GcsFilesetSpec) GetFilePatterns() []string {
	if m != nil {
		return m.FilePatterns
	}
	return nil
}

func (m *GcsFilesetSpec) GetSampleGcsFileSpecs() []*GcsFileSpec {
	if m != nil {
		return m.SampleGcsFileSpecs
	}
	return nil
}

// Specifications of a single file in Cloud Storage.
type GcsFileSpec struct {
	// Required. The full file path. Example: `gs://bucket_name/a/b.txt`.
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	// Output only. Timestamps about the Cloud Storage file.
	GcsTimestamps *SystemTimestamps `protobuf:"bytes,2,opt,name=gcs_timestamps,json=gcsTimestamps,proto3" json:"gcs_timestamps,omitempty"`
	// Output only. The size of the file, in bytes.
	SizeBytes            int64    `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GcsFileSpec) Reset()         { *m = GcsFileSpec{} }
func (m *GcsFileSpec) String() string { return proto.CompactTextString(m) }
func (*GcsFileSpec) ProtoMessage()    {}
func (*GcsFileSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ef0234cfb0655d3, []int{1}
}

func (m *GcsFileSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GcsFileSpec.Unmarshal(m, b)
}
func (m *GcsFileSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GcsFileSpec.Marshal(b, m, deterministic)
}
func (m *GcsFileSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GcsFileSpec.Merge(m, src)
}
func (m *GcsFileSpec) XXX_Size() int {
	return xxx_messageInfo_GcsFileSpec.Size(m)
}
func (m *GcsFileSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_GcsFileSpec.DiscardUnknown(m)
}

var xxx_messageInfo_GcsFileSpec proto.InternalMessageInfo

func (m *GcsFileSpec) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *GcsFileSpec) GetGcsTimestamps() *SystemTimestamps {
	if m != nil {
		return m.GcsTimestamps
	}
	return nil
}

func (m *GcsFileSpec) GetSizeBytes() int64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*GcsFilesetSpec)(nil), "google.cloud.datacatalog.v1.GcsFilesetSpec")
	proto.RegisterType((*GcsFileSpec)(nil), "google.cloud.datacatalog.v1.GcsFileSpec")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1/gcs_fileset_spec.proto", fileDescriptor_5ef0234cfb0655d3)
}

var fileDescriptor_5ef0234cfb0655d3 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xb1, 0x6e, 0xea, 0x30,
	0x14, 0x86, 0x65, 0x72, 0x75, 0x75, 0x31, 0x17, 0x86, 0x48, 0x57, 0x8a, 0xb8, 0x03, 0x11, 0x53,
	0x86, 0x36, 0x11, 0x74, 0xec, 0x54, 0x06, 0xba, 0x22, 0xe8, 0xd2, 0x2e, 0xa9, 0x31, 0x07, 0x63,
	0xc9, 0x89, 0xad, 0x1c, 0x17, 0x89, 0x3e, 0x4b, 0x9f, 0xa1, 0xcf, 0xd4, 0xc7, 0xe8, 0x58, 0x39,
	0x0e, 0x25, 0x43, 0x95, 0xd1, 0xff, 0xf1, 0x7f, 0xfe, 0xef, 0xd7, 0xa1, 0x73, 0xa1, 0xb5, 0x50,
	0x90, 0x71, 0xa5, 0x5f, 0x76, 0xd9, 0x8e, 0x59, 0xc6, 0x99, 0x65, 0x4a, 0x8b, 0xec, 0x38, 0xcb,
	0x04, 0xc7, 0x7c, 0x2f, 0x15, 0x20, 0xd8, 0x1c, 0x0d, 0xf0, 0xd4, 0x54, 0xda, 0xea, 0xf0, 0xbf,
	0xf7, 0xa4, 0xb5, 0x27, 0x6d, 0x79, 0xd2, 0xe3, 0x6c, 0x3c, 0x69, 0x16, 0x32, 0x23, 0xb3, 0xbd,
	0x04, 0xb5, 0xcb, 0xb7, 0x70, 0x60, 0x47, 0xa9, 0x2b, 0xef, 0x1e, 0x5f, 0x75, 0x25, 0x5a, 0x59,
	0x00, 0x5a, 0x56, 0x18, 0xf4, 0xbf, 0xa7, 0x6f, 0x84, 0x8e, 0xee, 0x39, 0x2e, 0x3d, 0xc5, 0xc6,
	0x00, 0x0f, 0x13, 0x3a, 0x74, 0x50, 0xb9, 0x61, 0xd6, 0x42, 0x55, 0x62, 0x44, 0xe2, 0x20, 0xe9,
	0x2f, 0x82, 0x8f, 0xbb, 0xde, 0xfa, 0xaf, 0x9b, 0xac, 0x9a, 0x41, 0xf8, 0x4c, 0xff, 0x21, 0x2b,
	0x8c, 0x82, 0xfc, 0xdc, 0xa4, 0xae, 0x81, 0x51, 0x2f, 0x0e, 0x92, 0xc1, 0x3c, 0x49, 0x3b, 0x8a,
	0xa4, 0x4d, 0xaa, 0x8b, 0x74, 0xbb, 0x83, 0x75, 0xe8, 0x77, 0xb5, 0x74, 0x9c, 0xbe, 0x13, 0x3a,
	0x68, 0x09, 0x61, 0x4c, 0xfb, 0x67, 0xb6, 0x43, 0x44, 0x62, 0x72, 0xe6, 0xfa, 0xd3, 0x70, 0x1d,
	0xc2, 0x47, 0x3a, 0x72, 0x30, 0x97, 0xa2, 0x51, 0x2f, 0x26, 0xc9, 0x60, 0x7e, 0xdd, 0x09, 0xb3,
	0x39, 0xa1, 0x85, 0xe2, 0xe1, 0xdb, 0xe4, 0x89, 0x86, 0x82, 0xe3, 0x45, 0x0b, 0xa7, 0x94, 0xa2,
	0x7c, 0x85, 0x7c, 0x7b, 0xb2, 0x80, 0xd1, 0xaf, 0x98, 0x24, 0x81, 0xff, 0xd7, 0x77, 0xf2, 0xc2,
	0xa9, 0x8b, 0x92, 0x4e, 0xb8, 0x2e, 0xba, 0xb2, 0x56, 0xe4, 0x69, 0xd9, 0x8c, 0x85, 0x56, 0xac,
	0x14, 0xa9, 0xae, 0x44, 0x26, 0xa0, 0xac, 0x0f, 0x92, 0xf9, 0x11, 0x33, 0x12, 0x7f, 0xbc, 0xe0,
	0x6d, 0xeb, 0xf9, 0x49, 0xc8, 0xf6, 0x77, 0xed, 0xba, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x15,
	0xfe, 0xbb, 0x50, 0x68, 0x02, 0x00, 0x00,
}
