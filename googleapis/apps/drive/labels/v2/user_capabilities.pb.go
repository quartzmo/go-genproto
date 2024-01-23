// Copyright 2023 Google LLC
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
// source: google/apps/drive/labels/v2/user_capabilities.proto

package labels

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

// The capabilities of a user.
type UserCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name for the user capabilities.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Whether the user is allowed access to the label manager.
	CanAccessLabelManager bool `protobuf:"varint,2,opt,name=can_access_label_manager,json=canAccessLabelManager,proto3" json:"can_access_label_manager,omitempty"`
	// Output only. Whether the user is an administrator for the shared labels
	// feature.
	CanAdministrateLabels bool `protobuf:"varint,3,opt,name=can_administrate_labels,json=canAdministrateLabels,proto3" json:"can_administrate_labels,omitempty"`
	// Output only. Whether the user is allowed to create new shared labels.
	CanCreateSharedLabels bool `protobuf:"varint,4,opt,name=can_create_shared_labels,json=canCreateSharedLabels,proto3" json:"can_create_shared_labels,omitempty"`
	// Output only. Whether the user is allowed to create new admin labels.
	CanCreateAdminLabels bool `protobuf:"varint,5,opt,name=can_create_admin_labels,json=canCreateAdminLabels,proto3" json:"can_create_admin_labels,omitempty"`
}

func (x *UserCapabilities) Reset() {
	*x = UserCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2_user_capabilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCapabilities) ProtoMessage() {}

func (x *UserCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2_user_capabilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCapabilities.ProtoReflect.Descriptor instead.
func (*UserCapabilities) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescGZIP(), []int{0}
}

func (x *UserCapabilities) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserCapabilities) GetCanAccessLabelManager() bool {
	if x != nil {
		return x.CanAccessLabelManager
	}
	return false
}

func (x *UserCapabilities) GetCanAdministrateLabels() bool {
	if x != nil {
		return x.CanAdministrateLabels
	}
	return false
}

func (x *UserCapabilities) GetCanCreateSharedLabels() bool {
	if x != nil {
		return x.CanCreateSharedLabels
	}
	return false
}

func (x *UserCapabilities) GetCanCreateAdminLabels() bool {
	if x != nil {
		return x.CanCreateAdminLabels
	}
	return false
}

var File_google_apps_drive_labels_v2_user_capabilities_proto protoreflect.FileDescriptor

var file_google_apps_drive_labels_v2_user_capabilities_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0,
	0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a,
	0x18, 0x63, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x03, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x17,
	0x63, 0x61, 0x6e, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2,
	0x41, 0x01, 0x03, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x18, 0x63, 0x61,
	0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x03, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x17, 0x63, 0x61, 0x6e,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03,
	0x52, 0x14, 0x63, 0x61, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x49, 0xea, 0x41, 0x46, 0x0a, 0x2b, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x17, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x42, 0x84, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x15, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0xa2, 0x02, 0x04, 0x44, 0x4c, 0x42, 0x4c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescOnce sync.Once
	file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescData = file_google_apps_drive_labels_v2_user_capabilities_proto_rawDesc
)

func file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescGZIP() []byte {
	file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescData)
	})
	return file_google_apps_drive_labels_v2_user_capabilities_proto_rawDescData
}

var file_google_apps_drive_labels_v2_user_capabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_apps_drive_labels_v2_user_capabilities_proto_goTypes = []interface{}{
	(*UserCapabilities)(nil), // 0: google.apps.drive.labels.v2.UserCapabilities
}
var file_google_apps_drive_labels_v2_user_capabilities_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_apps_drive_labels_v2_user_capabilities_proto_init() }
func file_google_apps_drive_labels_v2_user_capabilities_proto_init() {
	if File_google_apps_drive_labels_v2_user_capabilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_labels_v2_user_capabilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCapabilities); i {
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
			RawDescriptor: file_google_apps_drive_labels_v2_user_capabilities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_labels_v2_user_capabilities_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_labels_v2_user_capabilities_proto_depIdxs,
		MessageInfos:      file_google_apps_drive_labels_v2_user_capabilities_proto_msgTypes,
	}.Build()
	File_google_apps_drive_labels_v2_user_capabilities_proto = out.File
	file_google_apps_drive_labels_v2_user_capabilities_proto_rawDesc = nil
	file_google_apps_drive_labels_v2_user_capabilities_proto_goTypes = nil
	file_google_apps_drive_labels_v2_user_capabilities_proto_depIdxs = nil
}
