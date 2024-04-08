// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: policy/namespaces/namespaces.proto

package namespaces

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	common "github.com/opentdf/platform/protocol/go/common"
	policy "github.com/opentdf/platform/protocol/go/policy"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNamespaceRequest) Reset() {
	*x = GetNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespaceRequest) ProtoMessage() {}

func (x *GetNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespaceRequest.ProtoReflect.Descriptor instead.
func (*GetNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{0}
}

func (x *GetNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *policy.Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetNamespaceResponse) Reset() {
	*x = GetNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNamespaceResponse) ProtoMessage() {}

func (x *GetNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNamespaceResponse.ProtoReflect.Descriptor instead.
func (*GetNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{1}
}

func (x *GetNamespaceResponse) GetNamespace() *policy.Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type ListNamespacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ACTIVE by default when not specified
	State common.ActiveStateEnum `protobuf:"varint,1,opt,name=state,proto3,enum=common.ActiveStateEnum" json:"state,omitempty"`
}

func (x *ListNamespacesRequest) Reset() {
	*x = ListNamespacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesRequest) ProtoMessage() {}

func (x *ListNamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesRequest.ProtoReflect.Descriptor instead.
func (*ListNamespacesRequest) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{2}
}

func (x *ListNamespacesRequest) GetState() common.ActiveStateEnum {
	if x != nil {
		return x.State
	}
	return common.ActiveStateEnum(0)
}

type ListNamespacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespaces []*policy.Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *ListNamespacesResponse) Reset() {
	*x = ListNamespacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesResponse) ProtoMessage() {}

func (x *ListNamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesResponse.ProtoReflect.Descriptor instead.
func (*ListNamespacesResponse) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{3}
}

func (x *ListNamespacesResponse) GetNamespaces() []*policy.Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional
	Metadata *common.MetadataMutable `protobuf:"bytes,100,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetMetadata() *common.MetadataMutable {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CreateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *policy.Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *CreateNamespaceResponse) Reset() {
	*x = CreateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceResponse) ProtoMessage() {}

func (x *CreateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*CreateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{5}
}

func (x *CreateNamespaceResponse) GetNamespace() *policy.Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type UpdateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional
	Metadata               *common.MetadataMutable   `protobuf:"bytes,100,opt,name=metadata,proto3" json:"metadata,omitempty"`
	MetadataUpdateBehavior common.MetadataUpdateEnum `protobuf:"varint,101,opt,name=metadata_update_behavior,json=metadataUpdateBehavior,proto3,enum=common.MetadataUpdateEnum" json:"metadata_update_behavior,omitempty"`
}

func (x *UpdateNamespaceRequest) Reset() {
	*x = UpdateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceRequest) ProtoMessage() {}

func (x *UpdateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateNamespaceRequest) GetMetadata() *common.MetadataMutable {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateNamespaceRequest) GetMetadataUpdateBehavior() common.MetadataUpdateEnum {
	if x != nil {
		return x.MetadataUpdateBehavior
	}
	return common.MetadataUpdateEnum(0)
}

type UpdateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *policy.Namespace `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *UpdateNamespaceResponse) Reset() {
	*x = UpdateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceResponse) ProtoMessage() {}

func (x *UpdateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNamespaceResponse) GetNamespace() *policy.Namespace {
	if x != nil {
		return x.Namespace
	}
	return nil
}

type DeactivateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeactivateNamespaceRequest) Reset() {
	*x = DeactivateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateNamespaceRequest) ProtoMessage() {}

func (x *DeactivateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeactivateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{8}
}

func (x *DeactivateNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeactivateNamespaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeactivateNamespaceResponse) Reset() {
	*x = DeactivateNamespaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_namespaces_namespaces_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateNamespaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateNamespaceResponse) ProtoMessage() {}

func (x *DeactivateNamespaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_policy_namespaces_namespaces_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateNamespaceResponse.ProtoReflect.Descriptor instead.
func (*DeactivateNamespaceResponse) Descriptor() ([]byte, []int) {
	return file_policy_namespaces_namespaces_proto_rawDescGZIP(), []int{9}
}

var File_policy_namespaces_namespaces_proto protoreflect.FileDescriptor

var file_policy_namespaces_namespaces_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4b, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x0a,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0xfe, 0x04, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0xae, 0x04, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x99, 0x04, 0xba, 0x48, 0x95, 0x04, 0xba, 0x01, 0x89, 0x04, 0x0a,
	0x10, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0xa1, 0x03, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x20, 0x6d, 0x75,
	0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x61, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x20, 0x49, 0x74, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c,
	0x64, 0x20, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65, 0x61,
	0x73, 0x74, 0x20, 0x6f, 0x6e, 0x65, 0x20, 0x64, 0x6f, 0x74, 0x2c, 0x20, 0x77, 0x69, 0x74, 0x68,
	0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x28, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x29, 0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x61,
	0x6e, 0x64, 0x20, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x61,
	0x6e, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x20, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x20, 0x45, 0x61, 0x63, 0x68, 0x20, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x31, 0x20, 0x74,
	0x6f, 0x20, 0x36, 0x33, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x20,
	0x6c, 0x6f, 0x6e, 0x67, 0x2c, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x20, 0x68,
	0x79, 0x70, 0x68, 0x65, 0x6e, 0x73, 0x20, 0x62, 0x75, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61,
	0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x72, 0x73, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x6c,
	0x61, 0x73, 0x74, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x20, 0x54,
	0x68, 0x65, 0x20, 0x74, 0x6f, 0x70, 0x2d, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x20, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x20, 0x28, 0x74, 0x68, 0x65, 0x20, 0x6c, 0x61, 0x73, 0x74, 0x20, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x61, 0x66, 0x74, 0x65, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x20, 0x64, 0x6f, 0x74, 0x29, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20,
	0x63, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x74, 0x20, 0x6c, 0x65,
	0x61, 0x73, 0x74, 0x20, 0x74, 0x77, 0x6f, 0x20, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x65, 0x74,
	0x69, 0x63, 0x20, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x20, 0x54,
	0x68, 0x65, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20, 0x62, 0x65, 0x20, 0x6e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x20,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x1a, 0x51, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5d, 0x28, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5c, 0x5c, 0x2d,
	0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5d, 0x29, 0x3f, 0x5c, 0x5c, 0x2e, 0x29, 0x2b, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a,
	0x5d, 0x7b, 0x32, 0x2c, 0x7d, 0x24, 0x27, 0x29, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x18, 0xfd, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x54, 0x0a, 0x18, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x16, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x34, 0x0a, 0x1a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xde, 0x05, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x12, 0x16, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x0f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x29,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a,
	0x22, 0x16, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x32, 0x1b,
	0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x99, 0x01, 0x0a, 0x13,
	0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x2d, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xc8, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x42, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x64, 0x66, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x50, 0x4e, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xca, 0x02, 0x11, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_namespaces_namespaces_proto_rawDescOnce sync.Once
	file_policy_namespaces_namespaces_proto_rawDescData = file_policy_namespaces_namespaces_proto_rawDesc
)

func file_policy_namespaces_namespaces_proto_rawDescGZIP() []byte {
	file_policy_namespaces_namespaces_proto_rawDescOnce.Do(func() {
		file_policy_namespaces_namespaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_namespaces_namespaces_proto_rawDescData)
	})
	return file_policy_namespaces_namespaces_proto_rawDescData
}

var file_policy_namespaces_namespaces_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_policy_namespaces_namespaces_proto_goTypes = []interface{}{
	(*GetNamespaceRequest)(nil),         // 0: policy.namespaces.GetNamespaceRequest
	(*GetNamespaceResponse)(nil),        // 1: policy.namespaces.GetNamespaceResponse
	(*ListNamespacesRequest)(nil),       // 2: policy.namespaces.ListNamespacesRequest
	(*ListNamespacesResponse)(nil),      // 3: policy.namespaces.ListNamespacesResponse
	(*CreateNamespaceRequest)(nil),      // 4: policy.namespaces.CreateNamespaceRequest
	(*CreateNamespaceResponse)(nil),     // 5: policy.namespaces.CreateNamespaceResponse
	(*UpdateNamespaceRequest)(nil),      // 6: policy.namespaces.UpdateNamespaceRequest
	(*UpdateNamespaceResponse)(nil),     // 7: policy.namespaces.UpdateNamespaceResponse
	(*DeactivateNamespaceRequest)(nil),  // 8: policy.namespaces.DeactivateNamespaceRequest
	(*DeactivateNamespaceResponse)(nil), // 9: policy.namespaces.DeactivateNamespaceResponse
	(*policy.Namespace)(nil),            // 10: policy.Namespace
	(common.ActiveStateEnum)(0),         // 11: common.ActiveStateEnum
	(*common.MetadataMutable)(nil),      // 12: common.MetadataMutable
	(common.MetadataUpdateEnum)(0),      // 13: common.MetadataUpdateEnum
}
var file_policy_namespaces_namespaces_proto_depIdxs = []int32{
	10, // 0: policy.namespaces.GetNamespaceResponse.namespace:type_name -> policy.Namespace
	11, // 1: policy.namespaces.ListNamespacesRequest.state:type_name -> common.ActiveStateEnum
	10, // 2: policy.namespaces.ListNamespacesResponse.namespaces:type_name -> policy.Namespace
	12, // 3: policy.namespaces.CreateNamespaceRequest.metadata:type_name -> common.MetadataMutable
	10, // 4: policy.namespaces.CreateNamespaceResponse.namespace:type_name -> policy.Namespace
	12, // 5: policy.namespaces.UpdateNamespaceRequest.metadata:type_name -> common.MetadataMutable
	13, // 6: policy.namespaces.UpdateNamespaceRequest.metadata_update_behavior:type_name -> common.MetadataUpdateEnum
	10, // 7: policy.namespaces.UpdateNamespaceResponse.namespace:type_name -> policy.Namespace
	0,  // 8: policy.namespaces.NamespaceService.GetNamespace:input_type -> policy.namespaces.GetNamespaceRequest
	2,  // 9: policy.namespaces.NamespaceService.ListNamespaces:input_type -> policy.namespaces.ListNamespacesRequest
	4,  // 10: policy.namespaces.NamespaceService.CreateNamespace:input_type -> policy.namespaces.CreateNamespaceRequest
	6,  // 11: policy.namespaces.NamespaceService.UpdateNamespace:input_type -> policy.namespaces.UpdateNamespaceRequest
	8,  // 12: policy.namespaces.NamespaceService.DeactivateNamespace:input_type -> policy.namespaces.DeactivateNamespaceRequest
	1,  // 13: policy.namespaces.NamespaceService.GetNamespace:output_type -> policy.namespaces.GetNamespaceResponse
	3,  // 14: policy.namespaces.NamespaceService.ListNamespaces:output_type -> policy.namespaces.ListNamespacesResponse
	5,  // 15: policy.namespaces.NamespaceService.CreateNamespace:output_type -> policy.namespaces.CreateNamespaceResponse
	7,  // 16: policy.namespaces.NamespaceService.UpdateNamespace:output_type -> policy.namespaces.UpdateNamespaceResponse
	9,  // 17: policy.namespaces.NamespaceService.DeactivateNamespace:output_type -> policy.namespaces.DeactivateNamespaceResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_policy_namespaces_namespaces_proto_init() }
func file_policy_namespaces_namespaces_proto_init() {
	if File_policy_namespaces_namespaces_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_namespaces_namespaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespaceRequest); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNamespaceResponse); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespacesRequest); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNamespacesResponse); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceResponse); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceRequest); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceResponse); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateNamespaceRequest); i {
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
		file_policy_namespaces_namespaces_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactivateNamespaceResponse); i {
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
			RawDescriptor: file_policy_namespaces_namespaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_policy_namespaces_namespaces_proto_goTypes,
		DependencyIndexes: file_policy_namespaces_namespaces_proto_depIdxs,
		MessageInfos:      file_policy_namespaces_namespaces_proto_msgTypes,
	}.Build()
	File_policy_namespaces_namespaces_proto = out.File
	file_policy_namespaces_namespaces_proto_rawDesc = nil
	file_policy_namespaces_namespaces_proto_goTypes = nil
	file_policy_namespaces_namespaces_proto_depIdxs = nil
}
