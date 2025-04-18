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
// source: google/apps/script/type/gmail/gmail_addon_manifest.proto

package gmail

import (
	reflect "reflect"
	sync "sync"

	_type "google.golang.org/genproto/googleapis/apps/script/type"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// An enum defining the level of data access this compose trigger requires.
type ComposeTrigger_DraftAccess int32

const (
	// Default value when nothing is set for DraftAccess.
	ComposeTrigger_UNSPECIFIED ComposeTrigger_DraftAccess = 0
	// NONE means compose trigger won't be able to access any data of the draft
	// when a compose addon is triggered.
	ComposeTrigger_NONE ComposeTrigger_DraftAccess = 1
	// METADATA gives compose trigger the permission to access the metadata of
	// the draft when a compose addon is triggered. This includes the audience
	// list (To/cc list) of a draft message.
	ComposeTrigger_METADATA ComposeTrigger_DraftAccess = 2
)

// Enum value maps for ComposeTrigger_DraftAccess.
var (
	ComposeTrigger_DraftAccess_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "NONE",
		2: "METADATA",
	}
	ComposeTrigger_DraftAccess_value = map[string]int32{
		"UNSPECIFIED": 0,
		"NONE":        1,
		"METADATA":    2,
	}
)

func (x ComposeTrigger_DraftAccess) Enum() *ComposeTrigger_DraftAccess {
	p := new(ComposeTrigger_DraftAccess)
	*p = x
	return p
}

func (x ComposeTrigger_DraftAccess) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComposeTrigger_DraftAccess) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_enumTypes[0].Descriptor()
}

func (ComposeTrigger_DraftAccess) Type() protoreflect.EnumType {
	return &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_enumTypes[0]
}

func (x ComposeTrigger_DraftAccess) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComposeTrigger_DraftAccess.Descriptor instead.
func (ComposeTrigger_DraftAccess) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{2, 0}
}

// Properties customizing the appearance and execution of a Gmail add-on.
type GmailAddOnManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines an endpoint that will be executed in contexts that don't
	// match a declared contextual trigger. Any cards generated by this function
	// will always be available to the user, but may be eclipsed by contextual
	// content when this add-on declares more targeted triggers.
	//
	// If present, this overrides the configuration from
	// `addOns.common.homepageTrigger`.
	HomepageTrigger *_type.HomepageExtensionPoint `protobuf:"bytes,14,opt,name=homepage_trigger,json=homepageTrigger,proto3" json:"homepage_trigger,omitempty"`
	// Defines the set of conditions that trigger the add-on.
	ContextualTriggers []*ContextualTrigger `protobuf:"bytes,3,rep,name=contextual_triggers,json=contextualTriggers,proto3" json:"contextual_triggers,omitempty"`
	// Defines set of [universal
	// actions](/gmail/add-ons/how-tos/universal-actions) for the add-on. The user
	// triggers universal actions from the add-on toolbar menu.
	UniversalActions []*UniversalAction `protobuf:"bytes,4,rep,name=universal_actions,json=universalActions,proto3" json:"universal_actions,omitempty"`
	// Defines the compose time trigger for a compose time add-on. This is the
	// trigger that causes an add-on to take action when the user is composing an
	// email.
	// All compose time addons are required to have the
	// gmail.addons.current.action.compose scope even though it might not edit the
	// draft.
	ComposeTrigger *ComposeTrigger `protobuf:"bytes,12,opt,name=compose_trigger,json=composeTrigger,proto3" json:"compose_trigger,omitempty"`
	// The name of an endpoint that verifies that the add-on has
	// all the required third-party authorizations, by probing the third-party
	// APIs. If the probe fails, the function should throw an exception to
	// initiate the authorization flow. This function is called before each
	// invocation of the add-on, in order to ensure a smooth user experience.
	AuthorizationCheckFunction string `protobuf:"bytes,7,opt,name=authorization_check_function,json=authorizationCheckFunction,proto3" json:"authorization_check_function,omitempty"`
}

func (x *GmailAddOnManifest) Reset() {
	*x = GmailAddOnManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GmailAddOnManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GmailAddOnManifest) ProtoMessage() {}

func (x *GmailAddOnManifest) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GmailAddOnManifest.ProtoReflect.Descriptor instead.
func (*GmailAddOnManifest) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *GmailAddOnManifest) GetHomepageTrigger() *_type.HomepageExtensionPoint {
	if x != nil {
		return x.HomepageTrigger
	}
	return nil
}

func (x *GmailAddOnManifest) GetContextualTriggers() []*ContextualTrigger {
	if x != nil {
		return x.ContextualTriggers
	}
	return nil
}

func (x *GmailAddOnManifest) GetUniversalActions() []*UniversalAction {
	if x != nil {
		return x.UniversalActions
	}
	return nil
}

func (x *GmailAddOnManifest) GetComposeTrigger() *ComposeTrigger {
	if x != nil {
		return x.ComposeTrigger
	}
	return nil
}

func (x *GmailAddOnManifest) GetAuthorizationCheckFunction() string {
	if x != nil {
		return x.AuthorizationCheckFunction
	}
	return ""
}

// An action that is always available in the add-on toolbar menu regardless of
// message context.
type UniversalAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. User-visible text describing the action, for example, "Add a new
	// contact."
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The type of the action determines the behavior of Gmail when the user
	// invokes the action.
	//
	// Types that are assignable to ActionType:
	//
	//	*UniversalAction_OpenLink
	//	*UniversalAction_RunFunction
	ActionType isUniversalAction_ActionType `protobuf_oneof:"action_type"`
}

func (x *UniversalAction) Reset() {
	*x = UniversalAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UniversalAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniversalAction) ProtoMessage() {}

func (x *UniversalAction) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniversalAction.ProtoReflect.Descriptor instead.
func (*UniversalAction) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *UniversalAction) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (m *UniversalAction) GetActionType() isUniversalAction_ActionType {
	if m != nil {
		return m.ActionType
	}
	return nil
}

func (x *UniversalAction) GetOpenLink() string {
	if x, ok := x.GetActionType().(*UniversalAction_OpenLink); ok {
		return x.OpenLink
	}
	return ""
}

func (x *UniversalAction) GetRunFunction() string {
	if x, ok := x.GetActionType().(*UniversalAction_RunFunction); ok {
		return x.RunFunction
	}
	return ""
}

type isUniversalAction_ActionType interface {
	isUniversalAction_ActionType()
}

type UniversalAction_OpenLink struct {
	// A link that is opened by Gmail when the user triggers the action.
	OpenLink string `protobuf:"bytes,2,opt,name=open_link,json=openLink,proto3,oneof"`
}

type UniversalAction_RunFunction struct {
	// An endpoint that is called when the user triggers the
	// action. See the [universal actions
	// guide](/gmail/add-ons/how-tos/universal-actions) for details.
	RunFunction string `protobuf:"bytes,3,opt,name=run_function,json=runFunction,proto3,oneof"`
}

func (*UniversalAction_OpenLink) isUniversalAction_ActionType() {}

func (*UniversalAction_RunFunction) isUniversalAction_ActionType() {}

// A trigger that activates when user is composing an email.
type ComposeTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the set of actions for compose time add-on. These are actions
	// that user can trigger on a compose time addon.
	Actions []*_type.MenuItemExtensionPoint `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions,omitempty"`
	// Define the level of data access when a compose time addon is triggered.
	DraftAccess ComposeTrigger_DraftAccess `protobuf:"varint,4,opt,name=draft_access,json=draftAccess,proto3,enum=google.apps.script.type.gmail.ComposeTrigger_DraftAccess" json:"draft_access,omitempty"`
}

func (x *ComposeTrigger) Reset() {
	*x = ComposeTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeTrigger) ProtoMessage() {}

func (x *ComposeTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeTrigger.ProtoReflect.Descriptor instead.
func (*ComposeTrigger) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *ComposeTrigger) GetActions() []*_type.MenuItemExtensionPoint {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *ComposeTrigger) GetDraftAccess() ComposeTrigger_DraftAccess {
	if x != nil {
		return x.DraftAccess
	}
	return ComposeTrigger_UNSPECIFIED
}

// Defines a trigger that fires when the open email meets a specific criteria.
// When the trigger fires, it executes a specific endpoint, usually
// in order to create new cards and update the UI.
type ContextualTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of trigger determines the conditions Gmail uses to show the
	// add-on.
	//
	// Types that are assignable to Trigger:
	//
	//	*ContextualTrigger_Unconditional
	Trigger isContextualTrigger_Trigger `protobuf_oneof:"trigger"`
	// Required. The name of the endpoint to call when a message matches the
	// trigger.
	OnTriggerFunction string `protobuf:"bytes,4,opt,name=on_trigger_function,json=onTriggerFunction,proto3" json:"on_trigger_function,omitempty"`
}

func (x *ContextualTrigger) Reset() {
	*x = ContextualTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextualTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextualTrigger) ProtoMessage() {}

func (x *ContextualTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextualTrigger.ProtoReflect.Descriptor instead.
func (*ContextualTrigger) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{3}
}

func (m *ContextualTrigger) GetTrigger() isContextualTrigger_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (x *ContextualTrigger) GetUnconditional() *UnconditionalTrigger {
	if x, ok := x.GetTrigger().(*ContextualTrigger_Unconditional); ok {
		return x.Unconditional
	}
	return nil
}

func (x *ContextualTrigger) GetOnTriggerFunction() string {
	if x != nil {
		return x.OnTriggerFunction
	}
	return ""
}

type isContextualTrigger_Trigger interface {
	isContextualTrigger_Trigger()
}

type ContextualTrigger_Unconditional struct {
	// UnconditionalTriggers are executed when any mail message is opened.
	Unconditional *UnconditionalTrigger `protobuf:"bytes,1,opt,name=unconditional,proto3,oneof"`
}

func (*ContextualTrigger_Unconditional) isContextualTrigger_Trigger() {}

// A trigger that fires when any email message is opened.
type UnconditionalTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnconditionalTrigger) Reset() {
	*x = UnconditionalTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnconditionalTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnconditionalTrigger) ProtoMessage() {}

func (x *UnconditionalTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnconditionalTrigger.ProtoReflect.Descriptor instead.
func (*UnconditionalTrigger) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP(), []int{4}
}

var File_google_apps_script_type_gmail_gmail_addon_manifest_proto protoreflect.FileDescriptor

var file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2f,
	0x67, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f,
	0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x03, 0x0a, 0x12, 0x47, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12,
	0x5a, 0x0a, 0x10, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0f, 0x68, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x13, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x75, 0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x12, 0x5b,
	0x0a, 0x11, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x1c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a, 0x0f, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0c, 0x72,
	0x75, 0x6e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22,
	0xf1, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x12, 0x49, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5c, 0x0a,
	0x0c, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0b,
	0x64, 0x72, 0x61, 0x66, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x0b, 0x44,
	0x72, 0x61, 0x66, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54,
	0x41, 0x10, 0x02, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12, 0x5b, 0x0a, 0x0d, 0x75, 0x6e, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6d, 0x61, 0x69, 0x6c,
	0x2e, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x6e, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x6e, 0x5f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x6e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x6e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x42, 0xe0, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x42,
	0x17, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5c, 0x54, 0x79,
	0x70, 0x65, 0x5c, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3a,
	0x3a, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x3a, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescOnce sync.Once
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescData = file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDesc
)

func file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescGZIP() []byte {
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescOnce.Do(func() {
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescData)
	})
	return file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDescData
}

var file_google_apps_script_type_gmail_gmail_addon_manifest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_apps_script_type_gmail_gmail_addon_manifest_proto_goTypes = []interface{}{
	(ComposeTrigger_DraftAccess)(0),      // 0: google.apps.script.type.gmail.ComposeTrigger.DraftAccess
	(*GmailAddOnManifest)(nil),           // 1: google.apps.script.type.gmail.GmailAddOnManifest
	(*UniversalAction)(nil),              // 2: google.apps.script.type.gmail.UniversalAction
	(*ComposeTrigger)(nil),               // 3: google.apps.script.type.gmail.ComposeTrigger
	(*ContextualTrigger)(nil),            // 4: google.apps.script.type.gmail.ContextualTrigger
	(*UnconditionalTrigger)(nil),         // 5: google.apps.script.type.gmail.UnconditionalTrigger
	(*_type.HomepageExtensionPoint)(nil), // 6: google.apps.script.type.HomepageExtensionPoint
	(*_type.MenuItemExtensionPoint)(nil), // 7: google.apps.script.type.MenuItemExtensionPoint
}
var file_google_apps_script_type_gmail_gmail_addon_manifest_proto_depIdxs = []int32{
	6, // 0: google.apps.script.type.gmail.GmailAddOnManifest.homepage_trigger:type_name -> google.apps.script.type.HomepageExtensionPoint
	4, // 1: google.apps.script.type.gmail.GmailAddOnManifest.contextual_triggers:type_name -> google.apps.script.type.gmail.ContextualTrigger
	2, // 2: google.apps.script.type.gmail.GmailAddOnManifest.universal_actions:type_name -> google.apps.script.type.gmail.UniversalAction
	3, // 3: google.apps.script.type.gmail.GmailAddOnManifest.compose_trigger:type_name -> google.apps.script.type.gmail.ComposeTrigger
	7, // 4: google.apps.script.type.gmail.ComposeTrigger.actions:type_name -> google.apps.script.type.MenuItemExtensionPoint
	0, // 5: google.apps.script.type.gmail.ComposeTrigger.draft_access:type_name -> google.apps.script.type.gmail.ComposeTrigger.DraftAccess
	5, // 6: google.apps.script.type.gmail.ContextualTrigger.unconditional:type_name -> google.apps.script.type.gmail.UnconditionalTrigger
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_apps_script_type_gmail_gmail_addon_manifest_proto_init() }
func file_google_apps_script_type_gmail_gmail_addon_manifest_proto_init() {
	if File_google_apps_script_type_gmail_gmail_addon_manifest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GmailAddOnManifest); i {
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
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UniversalAction); i {
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
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeTrigger); i {
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
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextualTrigger); i {
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
		file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnconditionalTrigger); i {
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
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UniversalAction_OpenLink)(nil),
		(*UniversalAction_RunFunction)(nil),
	}
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ContextualTrigger_Unconditional)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_script_type_gmail_gmail_addon_manifest_proto_goTypes,
		DependencyIndexes: file_google_apps_script_type_gmail_gmail_addon_manifest_proto_depIdxs,
		EnumInfos:         file_google_apps_script_type_gmail_gmail_addon_manifest_proto_enumTypes,
		MessageInfos:      file_google_apps_script_type_gmail_gmail_addon_manifest_proto_msgTypes,
	}.Build()
	File_google_apps_script_type_gmail_gmail_addon_manifest_proto = out.File
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_rawDesc = nil
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_goTypes = nil
	file_google_apps_script_type_gmail_gmail_addon_manifest_proto_depIdxs = nil
}
