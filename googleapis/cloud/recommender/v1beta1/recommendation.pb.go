// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/recommender/v1beta1/recommendation.proto

package recommender

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	money "google.golang.org/genproto/googleapis/type/money"
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

// The category of the impact.
type Impact_Category int32

const (
	// Default unspecified category. Don't use directly.
	Impact_CATEGORY_UNSPECIFIED Impact_Category = 0
	// Indicates a potential increase or decrease in cost.
	Impact_COST Impact_Category = 1
	// Indicates a potential increase or decrease in security.
	Impact_SECURITY Impact_Category = 2
	// Indicates a potential increase or decrease in performance.
	Impact_PERFORMANCE Impact_Category = 3
)

var Impact_Category_name = map[int32]string{
	0: "CATEGORY_UNSPECIFIED",
	1: "COST",
	2: "SECURITY",
	3: "PERFORMANCE",
}

var Impact_Category_value = map[string]int32{
	"CATEGORY_UNSPECIFIED": 0,
	"COST":                 1,
	"SECURITY":             2,
	"PERFORMANCE":          3,
}

func (x Impact_Category) String() string {
	return proto.EnumName(Impact_Category_name, int32(x))
}

func (Impact_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{5, 0}
}

// Represents Recommendation State
type RecommendationStateInfo_State int32

const (
	// Default state. Don't use directly.
	RecommendationStateInfo_STATE_UNSPECIFIED RecommendationStateInfo_State = 0
	// Recommendation is active and can be applied. Recommendations content can
	// be updated by Google.
	//
	// ACTIVE recommendations can be marked as CLAIMED, SUCCEEDED, or FAILED.
	RecommendationStateInfo_ACTIVE RecommendationStateInfo_State = 1
	// Recommendation is in claimed state. Recommendations content is
	// immutable and cannot be updated by Google.
	//
	// CLAIMED recommendations can be marked as CLAIMED, SUCCEEDED, or FAILED.
	RecommendationStateInfo_CLAIMED RecommendationStateInfo_State = 6
	// Recommendation is in succeeded state. Recommendations content is
	// immutable and cannot be updated by Google.
	//
	// SUCCEEDED recommendations can be marked as SUCCEEDED, or FAILED.
	RecommendationStateInfo_SUCCEEDED RecommendationStateInfo_State = 3
	// Recommendation is in failed state. Recommendations content is immutable
	// and cannot be updated by Google.
	//
	// FAILED recommendations can be marked as SUCCEEDED, or FAILED.
	RecommendationStateInfo_FAILED RecommendationStateInfo_State = 4
	// Recommendation is in dismissed state. Recommendation content can be
	// updated by Google.
	//
	// DISMISSED recommendations can be marked as ACTIVE.
	RecommendationStateInfo_DISMISSED RecommendationStateInfo_State = 5
)

var RecommendationStateInfo_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	6: "CLAIMED",
	3: "SUCCEEDED",
	4: "FAILED",
	5: "DISMISSED",
}

var RecommendationStateInfo_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"CLAIMED":           6,
	"SUCCEEDED":         3,
	"FAILED":            4,
	"DISMISSED":         5,
}

func (x RecommendationStateInfo_State) String() string {
	return proto.EnumName(RecommendationStateInfo_State_name, int32(x))
}

func (RecommendationStateInfo_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{6, 0}
}

// A recommendation along with a suggested action. E.g., a rightsizing
// recommendation for an underutilized VM, IAM role recommendations, etc
type Recommendation struct {
	// Name of recommendation.
	//
	// A project recommendation is represented as
	//   projects/[PROJECT_NUMBER]/locations/[LOCATION]/recommenders/[RECOMMENDER_ID]/recommendations/[RECOMMENDATION_ID]
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Free-form human readable summary in English. The maximum length is 500
	// characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Contains an identifier for a subtype of recommendations produced for the
	// same recommender. Subtype is a function of content and impact, meaning a
	// new subtype will be added when either content or primary impact category
	// changes.
	//
	// Examples:
	//   For recommender = "google.iam.policy.RoleRecommender",
	//   recommender_subtype can be one of "REMOVE_ROLE"/"REPLACE_ROLE"
	RecommenderSubtype string `protobuf:"bytes,12,opt,name=recommender_subtype,json=recommenderSubtype,proto3" json:"recommender_subtype,omitempty"`
	// Last time this recommendation was refreshed by the system that created it
	// in the first place.
	LastRefreshTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_refresh_time,json=lastRefreshTime,proto3" json:"last_refresh_time,omitempty"`
	// The primary impact that this recommendation can have while trying to
	// optimize for one category.
	PrimaryImpact *Impact `protobuf:"bytes,5,opt,name=primary_impact,json=primaryImpact,proto3" json:"primary_impact,omitempty"`
	// Optional set of additional impact that this recommendation may have when
	// trying to optimize for the primary category. These may be positive
	// or negative.
	AdditionalImpact []*Impact `protobuf:"bytes,6,rep,name=additional_impact,json=additionalImpact,proto3" json:"additional_impact,omitempty"`
	// Content of the recommendation describing recommended changes to resources.
	Content *RecommendationContent `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	// Information for state. Contains state and metadata.
	StateInfo *RecommendationStateInfo `protobuf:"bytes,10,opt,name=state_info,json=stateInfo,proto3" json:"state_info,omitempty"`
	// Fingerprint of the Recommendation. Provides optimistic locking when
	// updating states.
	Etag                 string   `protobuf:"bytes,11,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recommendation) Reset()         { *m = Recommendation{} }
func (m *Recommendation) String() string { return proto.CompactTextString(m) }
func (*Recommendation) ProtoMessage()    {}
func (*Recommendation) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{0}
}

func (m *Recommendation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recommendation.Unmarshal(m, b)
}
func (m *Recommendation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recommendation.Marshal(b, m, deterministic)
}
func (m *Recommendation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recommendation.Merge(m, src)
}
func (m *Recommendation) XXX_Size() int {
	return xxx_messageInfo_Recommendation.Size(m)
}
func (m *Recommendation) XXX_DiscardUnknown() {
	xxx_messageInfo_Recommendation.DiscardUnknown(m)
}

var xxx_messageInfo_Recommendation proto.InternalMessageInfo

func (m *Recommendation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Recommendation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Recommendation) GetRecommenderSubtype() string {
	if m != nil {
		return m.RecommenderSubtype
	}
	return ""
}

func (m *Recommendation) GetLastRefreshTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastRefreshTime
	}
	return nil
}

func (m *Recommendation) GetPrimaryImpact() *Impact {
	if m != nil {
		return m.PrimaryImpact
	}
	return nil
}

func (m *Recommendation) GetAdditionalImpact() []*Impact {
	if m != nil {
		return m.AdditionalImpact
	}
	return nil
}

func (m *Recommendation) GetContent() *RecommendationContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Recommendation) GetStateInfo() *RecommendationStateInfo {
	if m != nil {
		return m.StateInfo
	}
	return nil
}

func (m *Recommendation) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// Contains what resources are changing and how they are changing.
type RecommendationContent struct {
	// Operations to one or more Google Cloud resources grouped in such a way
	// that, all operations within one group are expected to be performed
	// atomically and in an order.
	OperationGroups      []*OperationGroup `protobuf:"bytes,2,rep,name=operation_groups,json=operationGroups,proto3" json:"operation_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RecommendationContent) Reset()         { *m = RecommendationContent{} }
func (m *RecommendationContent) String() string { return proto.CompactTextString(m) }
func (*RecommendationContent) ProtoMessage()    {}
func (*RecommendationContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{1}
}

func (m *RecommendationContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationContent.Unmarshal(m, b)
}
func (m *RecommendationContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationContent.Marshal(b, m, deterministic)
}
func (m *RecommendationContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationContent.Merge(m, src)
}
func (m *RecommendationContent) XXX_Size() int {
	return xxx_messageInfo_RecommendationContent.Size(m)
}
func (m *RecommendationContent) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationContent.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationContent proto.InternalMessageInfo

func (m *RecommendationContent) GetOperationGroups() []*OperationGroup {
	if m != nil {
		return m.OperationGroups
	}
	return nil
}

// Group of operations that need to be performed atomically.
type OperationGroup struct {
	// List of operations across one or more resources that belong to this group.
	// Loosely based on RFC6902 and should be performed in the order they appear.
	Operations           []*Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *OperationGroup) Reset()         { *m = OperationGroup{} }
func (m *OperationGroup) String() string { return proto.CompactTextString(m) }
func (*OperationGroup) ProtoMessage()    {}
func (*OperationGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{2}
}

func (m *OperationGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationGroup.Unmarshal(m, b)
}
func (m *OperationGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationGroup.Marshal(b, m, deterministic)
}
func (m *OperationGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationGroup.Merge(m, src)
}
func (m *OperationGroup) XXX_Size() int {
	return xxx_messageInfo_OperationGroup.Size(m)
}
func (m *OperationGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationGroup.DiscardUnknown(m)
}

var xxx_messageInfo_OperationGroup proto.InternalMessageInfo

func (m *OperationGroup) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// Contains an operation for a resource inspired by the JSON-PATCH format with
// support for:
// * Custom filters for describing partial array patch.
// * Extended path values for describing nested arrays.
// * Custom fields for describing the resource for which the operation is being
//   described.
// * Allows extension to custom operations not natively supported by RFC6902.
// See https://tools.ietf.org/html/rfc6902 for details on the original RFC.
type Operation struct {
	// Type of this operation. Contains one of 'and', 'remove', 'replace', 'move',
	// 'copy', 'test' and custom operations. This field is case-insensitive and
	// always populated.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Type of GCP resource being modified/tested. This field is always populated.
	// Example: cloudresourcemanager.googleapis.com/Project,
	// compute.googleapis.com/Instance
	ResourceType string `protobuf:"bytes,2,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Contains the fully qualified resource name. This field is always populated.
	// ex: //cloudresourcemanager.googleapis.com/projects/foo.
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Path to the target field being operated on. If the operation is at the
	// resource level, then path should be "/". This field is always populated.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Can be set with action 'copy' to copy resource configuration across
	// different resources of the same type. Example: A resource clone can be
	// done via action = 'copy', path = "/", from = "/",
	// source_resource = <source> and resource_name = <target>.
	// This field is empty for all other values of `action`.
	SourceResource string `protobuf:"bytes,5,opt,name=source_resource,json=sourceResource,proto3" json:"source_resource,omitempty"`
	// Can be set with action 'copy' or 'move' to indicate the source field within
	// resource or source_resource, ignored if provided for other operation types.
	SourcePath string `protobuf:"bytes,6,opt,name=source_path,json=sourcePath,proto3" json:"source_path,omitempty"`
	// Value for the `path` field. Set if action is 'add'/'replace'/'test'.
	Value *_struct.Value `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	// Set of filters to apply if `path` refers to array elements or nested array
	// elements in order to narrow down to a single unique element that is being
	// tested/modified.
	// Note that this is intended to be an exact match per filter.
	// Example: {
	//   "/versions/*/name" : "it-123"
	//   "/versions/*/targetSize/percent": 20
	// }
	// Example: {
	//   "/bindings/*/role": "roles/admin"
	//   "/bindings/*/condition" : null
	// }
	// Example: {
	//   "/bindings/*/role": "roles/admin"
	//   "/bindings/*/members/*" : ["x@google.com", "y@google.com"]
	// }
	PathFilters          map[string]*_struct.Value `protobuf:"bytes,8,rep,name=path_filters,json=pathFilters,proto3" json:"path_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{3}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Operation) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *Operation) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Operation) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Operation) GetSourceResource() string {
	if m != nil {
		return m.SourceResource
	}
	return ""
}

func (m *Operation) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *Operation) GetValue() *_struct.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Operation) GetPathFilters() map[string]*_struct.Value {
	if m != nil {
		return m.PathFilters
	}
	return nil
}

// Contains metadata about how much money a recommendation can save or incur.
type CostProjection struct {
	// An approximate projection on amount saved or amount incurred. Negative cost
	// units indicate cost savings and positive cost units indicate increase.
	// See google.type.Money documentation for positive/negative units.
	Cost *money.Money `protobuf:"bytes,1,opt,name=cost,proto3" json:"cost,omitempty"`
	// Duration for which this cost applies.
	Duration             *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CostProjection) Reset()         { *m = CostProjection{} }
func (m *CostProjection) String() string { return proto.CompactTextString(m) }
func (*CostProjection) ProtoMessage()    {}
func (*CostProjection) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{4}
}

func (m *CostProjection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostProjection.Unmarshal(m, b)
}
func (m *CostProjection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostProjection.Marshal(b, m, deterministic)
}
func (m *CostProjection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostProjection.Merge(m, src)
}
func (m *CostProjection) XXX_Size() int {
	return xxx_messageInfo_CostProjection.Size(m)
}
func (m *CostProjection) XXX_DiscardUnknown() {
	xxx_messageInfo_CostProjection.DiscardUnknown(m)
}

var xxx_messageInfo_CostProjection proto.InternalMessageInfo

func (m *CostProjection) GetCost() *money.Money {
	if m != nil {
		return m.Cost
	}
	return nil
}

func (m *CostProjection) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// Contains the impact a recommendation can have for a given category.
type Impact struct {
	// Category that is being targeted.
	Category Impact_Category `protobuf:"varint,1,opt,name=category,proto3,enum=google.cloud.recommender.v1beta1.Impact_Category" json:"category,omitempty"`
	// Contains projections (if any) for this category.
	//
	// Types that are valid to be assigned to Projection:
	//	*Impact_CostProjection
	Projection           isImpact_Projection `protobuf_oneof:"projection"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Impact) Reset()         { *m = Impact{} }
func (m *Impact) String() string { return proto.CompactTextString(m) }
func (*Impact) ProtoMessage()    {}
func (*Impact) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{5}
}

func (m *Impact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Impact.Unmarshal(m, b)
}
func (m *Impact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Impact.Marshal(b, m, deterministic)
}
func (m *Impact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Impact.Merge(m, src)
}
func (m *Impact) XXX_Size() int {
	return xxx_messageInfo_Impact.Size(m)
}
func (m *Impact) XXX_DiscardUnknown() {
	xxx_messageInfo_Impact.DiscardUnknown(m)
}

var xxx_messageInfo_Impact proto.InternalMessageInfo

func (m *Impact) GetCategory() Impact_Category {
	if m != nil {
		return m.Category
	}
	return Impact_CATEGORY_UNSPECIFIED
}

type isImpact_Projection interface {
	isImpact_Projection()
}

type Impact_CostProjection struct {
	CostProjection *CostProjection `protobuf:"bytes,100,opt,name=cost_projection,json=costProjection,proto3,oneof"`
}

func (*Impact_CostProjection) isImpact_Projection() {}

func (m *Impact) GetProjection() isImpact_Projection {
	if m != nil {
		return m.Projection
	}
	return nil
}

func (m *Impact) GetCostProjection() *CostProjection {
	if x, ok := m.GetProjection().(*Impact_CostProjection); ok {
		return x.CostProjection
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Impact) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Impact_CostProjection)(nil),
	}
}

// Information for state. Contains state and metadata.
type RecommendationStateInfo struct {
	// The state of the recommendation, Eg ACTIVE, SUCCEEDED, FAILED.
	State RecommendationStateInfo_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.recommender.v1beta1.RecommendationStateInfo_State" json:"state,omitempty"`
	// A map of metadata for the state, provided by user or automations systems.
	StateMetadata        map[string]string `protobuf:"bytes,2,rep,name=state_metadata,json=stateMetadata,proto3" json:"state_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RecommendationStateInfo) Reset()         { *m = RecommendationStateInfo{} }
func (m *RecommendationStateInfo) String() string { return proto.CompactTextString(m) }
func (*RecommendationStateInfo) ProtoMessage()    {}
func (*RecommendationStateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_79be8d5bd206286c, []int{6}
}

func (m *RecommendationStateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationStateInfo.Unmarshal(m, b)
}
func (m *RecommendationStateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationStateInfo.Marshal(b, m, deterministic)
}
func (m *RecommendationStateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationStateInfo.Merge(m, src)
}
func (m *RecommendationStateInfo) XXX_Size() int {
	return xxx_messageInfo_RecommendationStateInfo.Size(m)
}
func (m *RecommendationStateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationStateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationStateInfo proto.InternalMessageInfo

func (m *RecommendationStateInfo) GetState() RecommendationStateInfo_State {
	if m != nil {
		return m.State
	}
	return RecommendationStateInfo_STATE_UNSPECIFIED
}

func (m *RecommendationStateInfo) GetStateMetadata() map[string]string {
	if m != nil {
		return m.StateMetadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.recommender.v1beta1.Impact_Category", Impact_Category_name, Impact_Category_value)
	proto.RegisterEnum("google.cloud.recommender.v1beta1.RecommendationStateInfo_State", RecommendationStateInfo_State_name, RecommendationStateInfo_State_value)
	proto.RegisterType((*Recommendation)(nil), "google.cloud.recommender.v1beta1.Recommendation")
	proto.RegisterType((*RecommendationContent)(nil), "google.cloud.recommender.v1beta1.RecommendationContent")
	proto.RegisterType((*OperationGroup)(nil), "google.cloud.recommender.v1beta1.OperationGroup")
	proto.RegisterType((*Operation)(nil), "google.cloud.recommender.v1beta1.Operation")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.recommender.v1beta1.Operation.PathFiltersEntry")
	proto.RegisterType((*CostProjection)(nil), "google.cloud.recommender.v1beta1.CostProjection")
	proto.RegisterType((*Impact)(nil), "google.cloud.recommender.v1beta1.Impact")
	proto.RegisterType((*RecommendationStateInfo)(nil), "google.cloud.recommender.v1beta1.RecommendationStateInfo")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.recommender.v1beta1.RecommendationStateInfo.StateMetadataEntry")
}

func init() {
	proto.RegisterFile("google/cloud/recommender/v1beta1/recommendation.proto", fileDescriptor_79be8d5bd206286c)
}

var fileDescriptor_79be8d5bd206286c = []byte{
	// 956 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x6b, 0x6f, 0xe3, 0x44,
	0x14, 0xdd, 0x3c, 0x9b, 0xdc, 0xb4, 0xa9, 0x3b, 0xec, 0xc3, 0x44, 0x88, 0x8d, 0x02, 0x82, 0x4a,
	0x20, 0x87, 0x16, 0xad, 0x78, 0x4a, 0x90, 0x3a, 0x6e, 0xb1, 0xb6, 0xd9, 0x96, 0x71, 0x52, 0xb1,
	0xac, 0x90, 0x35, 0x75, 0x26, 0xd9, 0x40, 0xec, 0xb1, 0xec, 0xc9, 0x4a, 0xf9, 0x0b, 0x7c, 0xe3,
	0x17, 0x20, 0xf1, 0x09, 0xf8, 0x95, 0x68, 0x1e, 0x4e, 0x9c, 0x76, 0x57, 0x6d, 0xd9, 0x6f, 0x33,
	0xf7, 0x9e, 0x73, 0xee, 0x9d, 0xeb, 0x33, 0x23, 0xc3, 0x93, 0x29, 0x63, 0xd3, 0x39, 0xed, 0x06,
	0x73, 0xb6, 0x18, 0x77, 0x13, 0x1a, 0xb0, 0x30, 0xa4, 0xd1, 0x98, 0x26, 0xdd, 0x57, 0x07, 0x97,
	0x94, 0x93, 0x83, 0x75, 0x8c, 0xf0, 0x19, 0x8b, 0xac, 0x38, 0x61, 0x9c, 0xa1, 0xb6, 0xa2, 0x59,
	0x92, 0x66, 0xe5, 0x68, 0x96, 0xa6, 0xb5, 0xde, 0xd7, 0xc2, 0x12, 0x7f, 0xb9, 0x98, 0x74, 0xc7,
	0x8b, 0x24, 0xa7, 0xd0, 0x7a, 0xef, 0x6a, 0x3e, 0xe5, 0xc9, 0x22, 0xe0, 0x3a, 0xfb, 0xf8, 0x6a,
	0x96, 0xcf, 0x42, 0x9a, 0x72, 0x12, 0xc6, 0x1a, 0xf0, 0x48, 0x03, 0xf8, 0x32, 0xa6, 0xdd, 0x90,
	0x45, 0x74, 0xa9, 0x12, 0x9d, 0xbf, 0xcb, 0xd0, 0xc4, 0x1b, 0x2d, 0x23, 0x04, 0xe5, 0x88, 0x84,
	0xd4, 0x2c, 0xb4, 0x0b, 0xfb, 0x75, 0x2c, 0xd7, 0xa8, 0x0d, 0x8d, 0x31, 0x4d, 0x83, 0x64, 0x16,
	0x0b, 0x88, 0x59, 0x94, 0xa9, 0x7c, 0x08, 0x75, 0xe1, 0x9d, 0xdc, 0xb9, 0xfc, 0x74, 0x71, 0x29,
	0x6a, 0x99, 0xdb, 0x12, 0x89, 0x72, 0x29, 0x4f, 0x65, 0xd0, 0x31, 0xec, 0xcd, 0x49, 0xca, 0xfd,
	0x84, 0x4e, 0x12, 0x9a, 0xbe, 0xf4, 0x45, 0xcb, 0x66, 0xb9, 0x5d, 0xd8, 0x6f, 0x1c, 0xb6, 0x2c,
	0x3d, 0xaf, 0xec, 0x3c, 0xd6, 0x30, 0x3b, 0x0f, 0xde, 0x15, 0x24, 0xac, 0x38, 0x22, 0x8a, 0xce,
	0xa0, 0x19, 0x27, 0xb3, 0x90, 0x24, 0x4b, 0x7f, 0x16, 0xc6, 0x24, 0xe0, 0x66, 0x45, 0x8a, 0xec,
	0x5b, 0x37, 0x0d, 0xdd, 0x72, 0x25, 0x1e, 0xef, 0x68, 0xbe, 0xda, 0xa2, 0x11, 0xec, 0x91, 0xf1,
	0x78, 0x26, 0x4e, 0x45, 0xe6, 0x99, 0x66, 0xb5, 0x5d, 0xba, 0x93, 0xa6, 0xb1, 0x96, 0xd0, 0xb2,
	0x3f, 0xc2, 0x56, 0xc0, 0x22, 0x4e, 0x23, 0x6e, 0x6e, 0xc9, 0x06, 0xbf, 0xb8, 0x59, 0x6c, 0xf3,
	0xcb, 0xd8, 0x8a, 0x8e, 0x33, 0x1d, 0xf4, 0x13, 0x40, 0xca, 0x09, 0xa7, 0xfe, 0x2c, 0x9a, 0x30,
	0x13, 0xa4, 0xea, 0x57, 0x77, 0x55, 0xf5, 0x84, 0x82, 0x1b, 0x4d, 0x18, 0xae, 0xa7, 0xd9, 0x52,
	0x78, 0x80, 0x72, 0x32, 0x35, 0x1b, 0xca, 0x03, 0x62, 0xdd, 0xe1, 0xf0, 0xe0, 0xb5, 0xfd, 0xa0,
	0x17, 0x60, 0xb0, 0x98, 0x2a, 0xbb, 0xfa, 0xd3, 0x84, 0x2d, 0xe2, 0xd4, 0x2c, 0xca, 0x79, 0x7d,
	0x76, 0x73, 0x33, 0x67, 0x19, 0xf3, 0x44, 0x10, 0xf1, 0x2e, 0xdb, 0xd8, 0xa7, 0x9d, 0x5f, 0xa0,
	0xb9, 0x09, 0x41, 0x4f, 0x01, 0x56, 0xa0, 0xd4, 0x2c, 0xc8, 0x42, 0x9f, 0xdc, 0xa1, 0x10, 0xce,
	0xd1, 0x3b, 0xff, 0x94, 0xa0, 0xbe, 0xca, 0xa0, 0x87, 0x50, 0x25, 0x81, 0x74, 0xb8, 0x32, 0xbf,
	0xde, 0xa1, 0x0f, 0x60, 0x27, 0xa1, 0x29, 0x5b, 0x24, 0x01, 0xf5, 0xa5, 0xad, 0xd5, 0x05, 0xd8,
	0xce, 0x82, 0x43, 0x61, 0xe8, 0x16, 0xd4, 0xb2, 0xbd, 0x59, 0x92, 0xf9, 0xd5, 0x5e, 0xcc, 0x33,
	0x26, 0xfc, 0xa5, 0xf4, 0x77, 0x1d, 0xcb, 0x35, 0xfa, 0x18, 0x76, 0xb5, 0xe4, 0x8a, 0x56, 0x91,
	0xe9, 0xa6, 0xda, 0xe1, 0x8c, 0xfc, 0x18, 0x1a, 0x1a, 0x28, 0x35, 0xaa, 0x12, 0x04, 0x2a, 0x74,
	0x2e, 0x94, 0x3e, 0x85, 0xca, 0x2b, 0x32, 0x5f, 0x50, 0x6d, 0xac, 0x87, 0xd7, 0xae, 0xcf, 0x85,
	0xc8, 0x62, 0x05, 0x42, 0x3e, 0x6c, 0x0b, 0x1d, 0x7f, 0x32, 0x9b, 0x73, 0x9a, 0xa4, 0x66, 0x4d,
	0x4e, 0xf0, 0xdb, 0x3b, 0x4c, 0xd0, 0x12, 0x55, 0x8f, 0x15, 0xdd, 0x89, 0x78, 0xb2, 0xc4, 0x8d,
	0x78, 0x1d, 0x69, 0x5d, 0x80, 0x71, 0x15, 0x80, 0x0c, 0x28, 0xfd, 0x46, 0x97, 0x7a, 0xac, 0x62,
	0xb9, 0x6e, 0xba, 0x78, 0x8b, 0xa6, 0xbf, 0x2e, 0x7e, 0x59, 0xe8, 0x30, 0x68, 0xda, 0x2c, 0xe5,
	0xe7, 0x09, 0xfb, 0x95, 0xaa, 0xef, 0xf2, 0x11, 0x94, 0x03, 0x96, 0x72, 0x29, 0xdb, 0x38, 0x44,
	0x99, 0x84, 0xf8, 0x44, 0xd6, 0x40, 0xbc, 0x72, 0x58, 0xe6, 0xd1, 0x13, 0xa8, 0x65, 0xef, 0xa9,
	0x2e, 0xf7, 0xee, 0xb5, 0x72, 0x7d, 0x0d, 0xc0, 0x2b, 0x68, 0xe7, 0x8f, 0x22, 0x54, 0xf5, 0xed,
	0x1d, 0x40, 0x2d, 0x20, 0x9c, 0x4e, 0x59, 0xa2, 0x0e, 0xd1, 0x3c, 0x3c, 0xb8, 0xed, 0x5b, 0x60,
	0xd9, 0x9a, 0x88, 0x57, 0x12, 0xe8, 0x05, 0xec, 0x8a, 0xc6, 0xfc, 0x78, 0x75, 0x16, 0x73, 0x2c,
	0xfb, 0xba, 0xc5, 0x8d, 0xd9, 0x9c, 0xc1, 0x0f, 0xf7, 0x70, 0x33, 0xd8, 0x88, 0x74, 0x06, 0x50,
	0xcb, 0x4a, 0x22, 0x13, 0xee, 0xdb, 0xbd, 0xa1, 0x73, 0x72, 0x86, 0x9f, 0xfb, 0xa3, 0x67, 0xde,
	0xb9, 0x63, 0xbb, 0xc7, 0xae, 0xd3, 0x37, 0xee, 0xa1, 0x1a, 0x94, 0xed, 0x33, 0x6f, 0x68, 0x14,
	0xd0, 0x36, 0xd4, 0x3c, 0xc7, 0x1e, 0x61, 0x77, 0xf8, 0xdc, 0x28, 0xa2, 0x5d, 0x68, 0x9c, 0x3b,
	0xf8, 0xf8, 0x0c, 0x0f, 0x7a, 0xcf, 0x6c, 0xc7, 0x28, 0x1d, 0x6d, 0x03, 0xac, 0xdb, 0xec, 0xfc,
	0x5e, 0x82, 0x47, 0x6f, 0x78, 0x40, 0xd0, 0x08, 0x2a, 0xf2, 0x09, 0xd1, 0x13, 0xfa, 0xee, 0x7f,
	0x3f, 0x45, 0x96, 0x5c, 0x61, 0xa5, 0x86, 0x52, 0x68, 0xaa, 0x67, 0x2e, 0xa4, 0x9c, 0x8c, 0x09,
	0x27, 0xfa, 0x75, 0x39, 0x7d, 0x4b, 0xfd, 0x81, 0x96, 0x53, 0x16, 0xde, 0x49, 0xf3, 0xb1, 0xd6,
	0xf7, 0x80, 0xae, 0x83, 0x5e, 0x63, 0xe3, 0xfb, 0x79, 0x1b, 0xd7, 0xf3, 0x76, 0x25, 0x50, 0x91,
	0x0a, 0xe8, 0x01, 0xec, 0x79, 0xc3, 0xde, 0xd0, 0xb9, 0xf2, 0x01, 0x00, 0xaa, 0x3d, 0x7b, 0xe8,
	0x5e, 0x38, 0x46, 0x01, 0x35, 0x60, 0xcb, 0x3e, 0xed, 0xb9, 0x03, 0xa7, 0x6f, 0x54, 0xd1, 0x0e,
	0xd4, 0xbd, 0x91, 0x6d, 0x3b, 0x4e, 0xdf, 0xe9, 0x1b, 0x25, 0x81, 0x3b, 0xee, 0xb9, 0xa7, 0x4e,
	0xdf, 0x28, 0x8b, 0x54, 0xdf, 0xf5, 0x06, 0xae, 0xe7, 0x39, 0x7d, 0xa3, 0x72, 0xf4, 0x67, 0x01,
	0x3e, 0x0c, 0x58, 0x78, 0xe3, 0x1c, 0xce, 0x0b, 0x3f, 0x3f, 0xd5, 0x98, 0x29, 0x9b, 0x93, 0x68,
	0x6a, 0xb1, 0x64, 0xda, 0x9d, 0xd2, 0x48, 0xba, 0xbf, 0xab, 0x52, 0x24, 0x9e, 0xa5, 0x6f, 0xfe,
	0xb1, 0xf9, 0x26, 0x17, 0xfb, 0xab, 0x58, 0xb6, 0xb1, 0x63, 0xff, 0x5b, 0x6c, 0x9f, 0x28, 0x51,
	0x5b, 0x16, 0xc6, 0xb9, 0xc2, 0x17, 0x07, 0x47, 0x82, 0x75, 0x59, 0x95, 0x15, 0x3e, 0xff, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0xef, 0x07, 0xdd, 0x2d, 0x39, 0x09, 0x00, 0x00,
}