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
// source: google/actions/sdk/v2/interactionmodel/type/class_reference.proto

package _type

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A reference to a class which is used to declare the type of a field or return
// value. Enums are also a type of class that can be referenced using
// ClassReference.
type ClassReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of a built-in type or custom type of the parameter. Examples:
	// `PizzaToppings`, `actions.type.Number`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Indicates whether the data type represents a list of values.
	List bool `protobuf:"varint,2,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *ClassReference) Reset() {
	*x = ClassReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassReference) ProtoMessage() {}

func (x *ClassReference) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassReference.ProtoReflect.Descriptor instead.
func (*ClassReference) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescGZIP(), []int{0}
}

func (x *ClassReference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClassReference) GetList() bool {
	if x != nil {
		return x.List
	}
	return false
}

var File_google_actions_sdk_v2_interactionmodel_type_class_reference_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x0e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x94, 0x01, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x13, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_goTypes = []interface{}{
	(*ClassReference)(nil), // 0: google.actions.sdk.v2.interactionmodel.type.ClassReference
}
var file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_type_class_reference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassReference); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_type_class_reference_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_type_class_reference_proto_depIdxs = nil
}
