// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/extensions/filters/network/thrift_proxy/filters/payload_to_metadata/v3/payload_to_metadata.proto

package payload_to_metadatav3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PayloadToMetadata_ValueType int32

const (
	PayloadToMetadata_STRING PayloadToMetadata_ValueType = 0
	PayloadToMetadata_NUMBER PayloadToMetadata_ValueType = 1
)

// Enum value maps for PayloadToMetadata_ValueType.
var (
	PayloadToMetadata_ValueType_name = map[int32]string{
		0: "STRING",
		1: "NUMBER",
	}
	PayloadToMetadata_ValueType_value = map[string]int32{
		"STRING": 0,
		"NUMBER": 1,
	}
)

func (x PayloadToMetadata_ValueType) Enum() *PayloadToMetadata_ValueType {
	p := new(PayloadToMetadata_ValueType)
	*p = x
	return p
}

func (x PayloadToMetadata_ValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadToMetadata_ValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_enumTypes[0].Descriptor()
}

func (PayloadToMetadata_ValueType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_enumTypes[0]
}

func (x PayloadToMetadata_ValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadToMetadata_ValueType.Descriptor instead.
func (PayloadToMetadata_ValueType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type PayloadToMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rules to apply to requests.
	RequestRules []*PayloadToMetadata_Rule `protobuf:"bytes,1,rep,name=request_rules,json=requestRules,proto3" json:"request_rules,omitempty"`
}

func (x *PayloadToMetadata) Reset() {
	*x = PayloadToMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadToMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadToMetadata) ProtoMessage() {}

func (x *PayloadToMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadToMetadata.ProtoReflect.Descriptor instead.
func (*PayloadToMetadata) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *PayloadToMetadata) GetRequestRules() []*PayloadToMetadata_Rule {
	if x != nil {
		return x.RequestRules
	}
	return nil
}

// [#next-free-field: 6]
type PayloadToMetadata_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to ValueType:
	//
	//	*PayloadToMetadata_KeyValuePair_Value
	//	*PayloadToMetadata_KeyValuePair_RegexValueRewrite
	ValueType isPayloadToMetadata_KeyValuePair_ValueType `protobuf_oneof:"value_type"`
	// The value's type — defaults to string.
	Type PayloadToMetadata_ValueType `protobuf:"varint,5,opt,name=type,proto3,enum=envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata_ValueType" json:"type,omitempty"`
}

func (x *PayloadToMetadata_KeyValuePair) Reset() {
	*x = PayloadToMetadata_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadToMetadata_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadToMetadata_KeyValuePair) ProtoMessage() {}

func (x *PayloadToMetadata_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadToMetadata_KeyValuePair.ProtoReflect.Descriptor instead.
func (*PayloadToMetadata_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PayloadToMetadata_KeyValuePair) GetMetadataNamespace() string {
	if x != nil {
		return x.MetadataNamespace
	}
	return ""
}

func (x *PayloadToMetadata_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *PayloadToMetadata_KeyValuePair) GetValueType() isPayloadToMetadata_KeyValuePair_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (x *PayloadToMetadata_KeyValuePair) GetValue() string {
	if x, ok := x.GetValueType().(*PayloadToMetadata_KeyValuePair_Value); ok {
		return x.Value
	}
	return ""
}

func (x *PayloadToMetadata_KeyValuePair) GetRegexValueRewrite() *v3.RegexMatchAndSubstitute {
	if x, ok := x.GetValueType().(*PayloadToMetadata_KeyValuePair_RegexValueRewrite); ok {
		return x.RegexValueRewrite
	}
	return nil
}

func (x *PayloadToMetadata_KeyValuePair) GetType() PayloadToMetadata_ValueType {
	if x != nil {
		return x.Type
	}
	return PayloadToMetadata_STRING
}

type isPayloadToMetadata_KeyValuePair_ValueType interface {
	isPayloadToMetadata_KeyValuePair_ValueType()
}

type PayloadToMetadata_KeyValuePair_Value struct {
	// The value to pair with the given key.
	//
	// When used for on_present case, if value is non-empty it'll be used instead
	// of the field value. If both are empty, the field value is used as-is.
	//
	// When used for on_missing case, a non-empty value must be provided.
	Value string `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

type PayloadToMetadata_KeyValuePair_RegexValueRewrite struct {
	// If present, the header's value will be matched and substituted with this.
	// If there is no match or substitution, the field value is used as-is.
	//
	// This is only used for on_present.
	RegexValueRewrite *v3.RegexMatchAndSubstitute `protobuf:"bytes,4,opt,name=regex_value_rewrite,json=regexValueRewrite,proto3,oneof"`
}

func (*PayloadToMetadata_KeyValuePair_Value) isPayloadToMetadata_KeyValuePair_ValueType() {}

func (*PayloadToMetadata_KeyValuePair_RegexValueRewrite) isPayloadToMetadata_KeyValuePair_ValueType() {
}

// A Rule defines what metadata to apply when a field is present or missing.
// [#next-free-field: 6]
type PayloadToMetadata_Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MatchSpecifier:
	//
	//	*PayloadToMetadata_Rule_MethodName
	//	*PayloadToMetadata_Rule_ServiceName
	MatchSpecifier isPayloadToMetadata_Rule_MatchSpecifier `protobuf_oneof:"match_specifier"`
	// Specifies that a match will be performed on the value of a field.
	FieldSelector *PayloadToMetadata_FieldSelector `protobuf:"bytes,3,opt,name=field_selector,json=fieldSelector,proto3" json:"field_selector,omitempty"`
	// If the field is present, apply this metadata KeyValuePair.
	OnPresent *PayloadToMetadata_KeyValuePair `protobuf:"bytes,4,opt,name=on_present,json=onPresent,proto3" json:"on_present,omitempty"`
	// If the field is missing, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set, since it'll be used in lieu
	// of the missing field value.
	OnMissing *PayloadToMetadata_KeyValuePair `protobuf:"bytes,5,opt,name=on_missing,json=onMissing,proto3" json:"on_missing,omitempty"`
}

func (x *PayloadToMetadata_Rule) Reset() {
	*x = PayloadToMetadata_Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadToMetadata_Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadToMetadata_Rule) ProtoMessage() {}

func (x *PayloadToMetadata_Rule) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadToMetadata_Rule.ProtoReflect.Descriptor instead.
func (*PayloadToMetadata_Rule) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP(), []int{0, 1}
}

func (m *PayloadToMetadata_Rule) GetMatchSpecifier() isPayloadToMetadata_Rule_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (x *PayloadToMetadata_Rule) GetMethodName() string {
	if x, ok := x.GetMatchSpecifier().(*PayloadToMetadata_Rule_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (x *PayloadToMetadata_Rule) GetServiceName() string {
	if x, ok := x.GetMatchSpecifier().(*PayloadToMetadata_Rule_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (x *PayloadToMetadata_Rule) GetFieldSelector() *PayloadToMetadata_FieldSelector {
	if x != nil {
		return x.FieldSelector
	}
	return nil
}

func (x *PayloadToMetadata_Rule) GetOnPresent() *PayloadToMetadata_KeyValuePair {
	if x != nil {
		return x.OnPresent
	}
	return nil
}

func (x *PayloadToMetadata_Rule) GetOnMissing() *PayloadToMetadata_KeyValuePair {
	if x != nil {
		return x.OnMissing
	}
	return nil
}

type isPayloadToMetadata_Rule_MatchSpecifier interface {
	isPayloadToMetadata_Rule_MatchSpecifier()
}

type PayloadToMetadata_Rule_MethodName struct {
	// If specified, the route must exactly match the request method name. As a special case,
	// an empty string matches any request method name.
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type PayloadToMetadata_Rule_ServiceName struct {
	// If specified, the route must have the service name as the request method name prefix.
	// As a special case, an empty string matches any service name. Only relevant when service
	// multiplexing.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*PayloadToMetadata_Rule_MethodName) isPayloadToMetadata_Rule_MatchSpecifier() {}

func (*PayloadToMetadata_Rule_ServiceName) isPayloadToMetadata_Rule_MatchSpecifier() {}

type PayloadToMetadata_FieldSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// field name to log
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// field id to match
	Id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// next node of the field selector
	Child *PayloadToMetadata_FieldSelector `protobuf:"bytes,3,opt,name=child,proto3" json:"child,omitempty"`
}

func (x *PayloadToMetadata_FieldSelector) Reset() {
	*x = PayloadToMetadata_FieldSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadToMetadata_FieldSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadToMetadata_FieldSelector) ProtoMessage() {}

func (x *PayloadToMetadata_FieldSelector) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadToMetadata_FieldSelector.ProtoReflect.Descriptor instead.
func (*PayloadToMetadata_FieldSelector) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP(), []int{0, 2}
}

func (x *PayloadToMetadata_FieldSelector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PayloadToMetadata_FieldSelector) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayloadToMetadata_FieldSelector) GetChild() *PayloadToMetadata_FieldSelector {
	if x != nil {
		return x.Child
	}
	return nil
}

var File_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDesc = []byte{
	0x0a, 0x66, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66,
	0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65,
	0x67, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbc, 0x0a, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x93, 0x01, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x64, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0xea, 0x02,
	0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x2d,
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x60, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x65, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f,
	0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x67, 0x65, 0x78, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x41, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x11, 0x72, 0x65, 0x67, 0x65, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x69, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0c, 0x0a, 0x0a,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x1a, 0xa3, 0x04, 0x0a, 0x04, 0x52,
	0x75, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x0e,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0d, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x8b, 0x01, 0x0a,
	0x0a, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x6c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x09, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x6f,
	0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x6c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x09, 0x6f,
	0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x42, 0x16, 0x0a, 0x0f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x1a, 0xd8, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0xfa, 0x42, 0x11,
	0x1a, 0x0f, 0x18, 0xff, 0xff, 0x01, 0x28, 0x80, 0x80, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x6d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x22, 0x23, 0x0a, 0x09, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01,
	0x42, 0x8a, 0x02, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x0a, 0x5a, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x54, 0x6f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x89, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69,
	0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x33, 0x3b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescData = file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDesc
)

func file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDescData
}

var file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_goTypes = []interface{}{
	(PayloadToMetadata_ValueType)(0),        // 0: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.ValueType
	(*PayloadToMetadata)(nil),               // 1: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata
	(*PayloadToMetadata_KeyValuePair)(nil),  // 2: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.KeyValuePair
	(*PayloadToMetadata_Rule)(nil),          // 3: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.Rule
	(*PayloadToMetadata_FieldSelector)(nil), // 4: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.FieldSelector
	(*v3.RegexMatchAndSubstitute)(nil),      // 5: envoy.type.matcher.v3.RegexMatchAndSubstitute
}
var file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.request_rules:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.Rule
	5, // 1: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.KeyValuePair.regex_value_rewrite:type_name -> envoy.type.matcher.v3.RegexMatchAndSubstitute
	0, // 2: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.KeyValuePair.type:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.ValueType
	4, // 3: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.Rule.field_selector:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.FieldSelector
	2, // 4: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.Rule.on_present:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.KeyValuePair
	2, // 5: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.Rule.on_missing:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.KeyValuePair
	4, // 6: envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.FieldSelector.child:type_name -> envoy.extensions.filters.network.thrift_proxy.filters.payload_to_metadata.v3.PayloadToMetadata.FieldSelector
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_init()
}
func file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_init() {
	if File_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadToMetadata); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadToMetadata_KeyValuePair); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadToMetadata_Rule); i {
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
		file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadToMetadata_FieldSelector); i {
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
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PayloadToMetadata_KeyValuePair_Value)(nil),
		(*PayloadToMetadata_KeyValuePair_RegexValueRewrite)(nil),
	}
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PayloadToMetadata_Rule_MethodName)(nil),
		(*PayloadToMetadata_Rule_ServiceName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto = out.File
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_rawDesc = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_goTypes = nil
	file_envoy_extensions_filters_network_thrift_proxy_filters_payload_to_metadata_v3_payload_to_metadata_proto_depIdxs = nil
}
