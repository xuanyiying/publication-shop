// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: api/catalog/v1/catalog.proto

package v1

import (
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

type CreateCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCatalogRequest) Reset() {
	*x = CreateCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogRequest) ProtoMessage() {}

func (x *CreateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogRequest.ProtoReflect.Descriptor instead.
func (*CreateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{0}
}

type CreateCatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCatalogReply) Reset() {
	*x = CreateCatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogReply) ProtoMessage() {}

func (x *CreateCatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogReply.ProtoReflect.Descriptor instead.
func (*CreateCatalogReply) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{1}
}

type UpdateCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCatalogRequest) Reset() {
	*x = UpdateCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogRequest) ProtoMessage() {}

func (x *UpdateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{2}
}

type UpdateCatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCatalogReply) Reset() {
	*x = UpdateCatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogReply) ProtoMessage() {}

func (x *UpdateCatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogReply.ProtoReflect.Descriptor instead.
func (*UpdateCatalogReply) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{3}
}

type DeleteCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCatalogRequest) Reset() {
	*x = DeleteCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogRequest) ProtoMessage() {}

func (x *DeleteCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogRequest.ProtoReflect.Descriptor instead.
func (*DeleteCatalogRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{4}
}

type DeleteCatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCatalogReply) Reset() {
	*x = DeleteCatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogReply) ProtoMessage() {}

func (x *DeleteCatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogReply.ProtoReflect.Descriptor instead.
func (*DeleteCatalogReply) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{5}
}

type GetCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{6}
}

type GetCatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCatalogReply) Reset() {
	*x = GetCatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogReply) ProtoMessage() {}

func (x *GetCatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogReply.ProtoReflect.Descriptor instead.
func (*GetCatalogReply) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{7}
}

type ListCatalogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCatalogRequest) Reset() {
	*x = ListCatalogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogRequest) ProtoMessage() {}

func (x *ListCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogRequest.ProtoReflect.Descriptor instead.
func (*ListCatalogRequest) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{8}
}

type ListCatalogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCatalogReply) Reset() {
	*x = ListCatalogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_catalog_v1_catalog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCatalogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogReply) ProtoMessage() {}

func (x *ListCatalogReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_catalog_v1_catalog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogReply.ProtoReflect.Descriptor instead.
func (*ListCatalogReply) Descriptor() ([]byte, []int) {
	return file_api_catalog_v1_catalog_proto_rawDescGZIP(), []int{9}
}

var File_api_catalog_v1_catalog_proto protoreflect.FileDescriptor

var file_api_catalog_v1_catalog_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x16,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xc1, 0x03, 0x0a, 0x07, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x59, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x59, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x59, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x53, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x41,
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x68, 0x6f, 0x70, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_catalog_v1_catalog_proto_rawDescOnce sync.Once
	file_api_catalog_v1_catalog_proto_rawDescData = file_api_catalog_v1_catalog_proto_rawDesc
)

func file_api_catalog_v1_catalog_proto_rawDescGZIP() []byte {
	file_api_catalog_v1_catalog_proto_rawDescOnce.Do(func() {
		file_api_catalog_v1_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_catalog_v1_catalog_proto_rawDescData)
	})
	return file_api_catalog_v1_catalog_proto_rawDescData
}

var file_api_catalog_v1_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_catalog_v1_catalog_proto_goTypes = []interface{}{
	(*CreateCatalogRequest)(nil), // 0: api.catalog.v1.CreateCatalogRequest
	(*CreateCatalogReply)(nil),   // 1: api.catalog.v1.CreateCatalogReply
	(*UpdateCatalogRequest)(nil), // 2: api.catalog.v1.UpdateCatalogRequest
	(*UpdateCatalogReply)(nil),   // 3: api.catalog.v1.UpdateCatalogReply
	(*DeleteCatalogRequest)(nil), // 4: api.catalog.v1.DeleteCatalogRequest
	(*DeleteCatalogReply)(nil),   // 5: api.catalog.v1.DeleteCatalogReply
	(*GetCatalogRequest)(nil),    // 6: api.catalog.v1.GetCatalogRequest
	(*GetCatalogReply)(nil),      // 7: api.catalog.v1.GetCatalogReply
	(*ListCatalogRequest)(nil),   // 8: api.catalog.v1.ListCatalogRequest
	(*ListCatalogReply)(nil),     // 9: api.catalog.v1.ListCatalogReply
}
var file_api_catalog_v1_catalog_proto_depIdxs = []int32{
	0, // 0: api.catalog.v1.Catalog.CreateCatalog:input_type -> api.catalog.v1.CreateCatalogRequest
	2, // 1: api.catalog.v1.Catalog.UpdateCatalog:input_type -> api.catalog.v1.UpdateCatalogRequest
	4, // 2: api.catalog.v1.Catalog.DeleteCatalog:input_type -> api.catalog.v1.DeleteCatalogRequest
	6, // 3: api.catalog.v1.Catalog.GetCatalog:input_type -> api.catalog.v1.GetCatalogRequest
	8, // 4: api.catalog.v1.Catalog.ListCatalog:input_type -> api.catalog.v1.ListCatalogRequest
	1, // 5: api.catalog.v1.Catalog.CreateCatalog:output_type -> api.catalog.v1.CreateCatalogReply
	3, // 6: api.catalog.v1.Catalog.UpdateCatalog:output_type -> api.catalog.v1.UpdateCatalogReply
	5, // 7: api.catalog.v1.Catalog.DeleteCatalog:output_type -> api.catalog.v1.DeleteCatalogReply
	7, // 8: api.catalog.v1.Catalog.GetCatalog:output_type -> api.catalog.v1.GetCatalogReply
	9, // 9: api.catalog.v1.Catalog.ListCatalog:output_type -> api.catalog.v1.ListCatalogReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_catalog_v1_catalog_proto_init() }
func file_api_catalog_v1_catalog_proto_init() {
	if File_api_catalog_v1_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_catalog_v1_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatalogRequest); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCatalogReply); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatalogRequest); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCatalogReply); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatalogRequest); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCatalogReply); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogRequest); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCatalogReply); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogRequest); i {
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
		file_api_catalog_v1_catalog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCatalogReply); i {
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
			RawDescriptor: file_api_catalog_v1_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_catalog_v1_catalog_proto_goTypes,
		DependencyIndexes: file_api_catalog_v1_catalog_proto_depIdxs,
		MessageInfos:      file_api_catalog_v1_catalog_proto_msgTypes,
	}.Build()
	File_api_catalog_v1_catalog_proto = out.File
	file_api_catalog_v1_catalog_proto_rawDesc = nil
	file_api_catalog_v1_catalog_proto_goTypes = nil
	file_api_catalog_v1_catalog_proto_depIdxs = nil
}
