// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/actions/sdk/v2/surface.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible set of surface capabilities.
type CapabilityRequirement_SurfaceCapability int32

const (
	// Unknown / Unspecified.
	CapabilityRequirement_SURFACE_CAPABILITY_UNSPECIFIED CapabilityRequirement_SurfaceCapability = 0
	// Surface supports audio output.
	CapabilityRequirement_AUDIO_OUTPUT CapabilityRequirement_SurfaceCapability = 1
	// Surface supports screen/visual output.
	CapabilityRequirement_SCREEN_OUTPUT CapabilityRequirement_SurfaceCapability = 2
	// Surface supports media response audio.
	CapabilityRequirement_MEDIA_RESPONSE_AUDIO CapabilityRequirement_SurfaceCapability = 3
	// Surface supports web browsers.
	CapabilityRequirement_WEB_BROWSER CapabilityRequirement_SurfaceCapability = 4
	// Surface supports account linking.
	CapabilityRequirement_ACCOUNT_LINKING CapabilityRequirement_SurfaceCapability = 7
	// Surface supports Interactive Canvas.
	CapabilityRequirement_INTERACTIVE_CANVAS CapabilityRequirement_SurfaceCapability = 8
	// Surface supports home storage.
	CapabilityRequirement_HOME_STORAGE CapabilityRequirement_SurfaceCapability = 9
)

// Enum value maps for CapabilityRequirement_SurfaceCapability.
var (
	CapabilityRequirement_SurfaceCapability_name = map[int32]string{
		0: "SURFACE_CAPABILITY_UNSPECIFIED",
		1: "AUDIO_OUTPUT",
		2: "SCREEN_OUTPUT",
		3: "MEDIA_RESPONSE_AUDIO",
		4: "WEB_BROWSER",
		7: "ACCOUNT_LINKING",
		8: "INTERACTIVE_CANVAS",
		9: "HOME_STORAGE",
	}
	CapabilityRequirement_SurfaceCapability_value = map[string]int32{
		"SURFACE_CAPABILITY_UNSPECIFIED": 0,
		"AUDIO_OUTPUT":                   1,
		"SCREEN_OUTPUT":                  2,
		"MEDIA_RESPONSE_AUDIO":           3,
		"WEB_BROWSER":                    4,
		"ACCOUNT_LINKING":                7,
		"INTERACTIVE_CANVAS":             8,
		"HOME_STORAGE":                   9,
	}
)

func (x CapabilityRequirement_SurfaceCapability) Enum() *CapabilityRequirement_SurfaceCapability {
	p := new(CapabilityRequirement_SurfaceCapability)
	*p = x
	return p
}

func (x CapabilityRequirement_SurfaceCapability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapabilityRequirement_SurfaceCapability) Descriptor() protoreflect.EnumDescriptor {
	return file_google_actions_sdk_v2_surface_proto_enumTypes[0].Descriptor()
}

func (CapabilityRequirement_SurfaceCapability) Type() protoreflect.EnumType {
	return &file_google_actions_sdk_v2_surface_proto_enumTypes[0]
}

func (x CapabilityRequirement_SurfaceCapability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapabilityRequirement_SurfaceCapability.Descriptor instead.
func (CapabilityRequirement_SurfaceCapability) EnumDescriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_surface_proto_rawDescGZIP(), []int{1, 0}
}

// Contains a set of requirements that the client surface must support to invoke
// Actions in your project.
type SurfaceRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The minimum set of capabilities needed to invoke the Actions in your
	// project. If the surface is missing any of these, the Action will not be
	// triggered.
	MinimumRequirements []*CapabilityRequirement `protobuf:"bytes,1,rep,name=minimum_requirements,json=minimumRequirements,proto3" json:"minimum_requirements,omitempty"`
}

func (x *SurfaceRequirements) Reset() {
	*x = SurfaceRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_surface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SurfaceRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SurfaceRequirements) ProtoMessage() {}

func (x *SurfaceRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_surface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SurfaceRequirements.ProtoReflect.Descriptor instead.
func (*SurfaceRequirements) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_surface_proto_rawDescGZIP(), []int{0}
}

func (x *SurfaceRequirements) GetMinimumRequirements() []*CapabilityRequirement {
	if x != nil {
		return x.MinimumRequirements
	}
	return nil
}

// Represents a requirement about the availability of a given capability.
type CapabilityRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of capability.
	Capability CapabilityRequirement_SurfaceCapability `protobuf:"varint,1,opt,name=capability,proto3,enum=google.actions.sdk.v2.CapabilityRequirement_SurfaceCapability" json:"capability,omitempty"`
}

func (x *CapabilityRequirement) Reset() {
	*x = CapabilityRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_surface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityRequirement) ProtoMessage() {}

func (x *CapabilityRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_surface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityRequirement.ProtoReflect.Descriptor instead.
func (*CapabilityRequirement) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_surface_proto_rawDescGZIP(), []int{1}
}

func (x *CapabilityRequirement) GetCapability() CapabilityRequirement_SurfaceCapability {
	if x != nil {
		return x.Capability
	}
	return CapabilityRequirement_SURFACE_CAPABILITY_UNSPECIFIED
}

var File_google_actions_sdk_v2_surface_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_surface_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x22, 0x76, 0x0a, 0x13,
	0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x5f, 0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x13, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0xc0, 0x02, 0x0a, 0x15, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5e,
	0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xc6,
	0x01, 0x0a, 0x11, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f,
	0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x55, 0x44, 0x49,
	0x4f, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43,
	0x52, 0x45, 0x45, 0x4e, 0x5f, 0x4f, 0x55, 0x54, 0x50, 0x55, 0x54, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f,
	0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x57, 0x45, 0x42, 0x5f, 0x42,
	0x52, 0x4f, 0x57, 0x53, 0x45, 0x52, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x43, 0x41, 0x4e,
	0x56, 0x41, 0x53, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x53, 0x54,
	0x4f, 0x52, 0x41, 0x47, 0x45, 0x10, 0x09, 0x42, 0x65, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x42, 0x0c, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_surface_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_surface_proto_rawDescData = file_google_actions_sdk_v2_surface_proto_rawDesc
)

func file_google_actions_sdk_v2_surface_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_surface_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_surface_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_surface_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_surface_proto_rawDescData
}

var file_google_actions_sdk_v2_surface_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_actions_sdk_v2_surface_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_actions_sdk_v2_surface_proto_goTypes = []interface{}{
	(CapabilityRequirement_SurfaceCapability)(0), // 0: google.actions.sdk.v2.CapabilityRequirement.SurfaceCapability
	(*SurfaceRequirements)(nil),                  // 1: google.actions.sdk.v2.SurfaceRequirements
	(*CapabilityRequirement)(nil),                // 2: google.actions.sdk.v2.CapabilityRequirement
}
var file_google_actions_sdk_v2_surface_proto_depIdxs = []int32{
	2, // 0: google.actions.sdk.v2.SurfaceRequirements.minimum_requirements:type_name -> google.actions.sdk.v2.CapabilityRequirement
	0, // 1: google.actions.sdk.v2.CapabilityRequirement.capability:type_name -> google.actions.sdk.v2.CapabilityRequirement.SurfaceCapability
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_surface_proto_init() }
func file_google_actions_sdk_v2_surface_proto_init() {
	if File_google_actions_sdk_v2_surface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_surface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SurfaceRequirements); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_actions_sdk_v2_surface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityRequirement); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_actions_sdk_v2_surface_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_surface_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_surface_proto_depIdxs,
		EnumInfos:         file_google_actions_sdk_v2_surface_proto_enumTypes,
		MessageInfos:      file_google_actions_sdk_v2_surface_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_surface_proto = out.File
	file_google_actions_sdk_v2_surface_proto_rawDesc = nil
	file_google_actions_sdk_v2_surface_proto_goTypes = nil
	file_google_actions_sdk_v2_surface_proto_depIdxs = nil
}
