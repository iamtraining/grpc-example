// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/crud.proto

package crudapi

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *CreateBlogRequest) Reset() {
	*x = CreateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogRequest) ProtoMessage() {}

func (x *CreateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogRequest.ProtoReflect.Descriptor instead.
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlogRequest) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *CreateBlogResponse) Reset() {
	*x = CreateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlogResponse) ProtoMessage() {}

func (x *CreateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlogResponse.ProtoReflect.Descriptor instead.
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId string `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *ReadBlogRequest) Reset() {
	*x = ReadBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogRequest) ProtoMessage() {}

func (x *ReadBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogRequest.ProtoReflect.Descriptor instead.
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{3}
}

func (x *ReadBlogRequest) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ReadBlogResponse) Reset() {
	*x = ReadBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBlogResponse) ProtoMessage() {}

func (x *ReadBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBlogResponse.ProtoReflect.Descriptor instead.
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{4}
}

func (x *ReadBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogRequest) Reset() {
	*x = UpdateBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogRequest) ProtoMessage() {}

func (x *UpdateBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBlogRequest) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type UpdateBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *UpdateBlogResponse) Reset() {
	*x = UpdateBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlogResponse) ProtoMessage() {}

func (x *UpdateBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlogResponse.ProtoReflect.Descriptor instead.
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId string `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *DeleteBlogRequest) Reset() {
	*x = DeleteBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogRequest) ProtoMessage() {}

func (x *DeleteBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBlogRequest) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type DeleteBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlogId string `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
}

func (x *DeleteBlogResponse) Reset() {
	*x = DeleteBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlogResponse) ProtoMessage() {}

func (x *DeleteBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlogResponse.ProtoReflect.Descriptor instead.
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBlogResponse) GetBlogId() string {
	if x != nil {
		return x.BlogId
	}
	return ""
}

type ListBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBlogRequest) Reset() {
	*x = ListBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogRequest) ProtoMessage() {}

func (x *ListBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogRequest.ProtoReflect.Descriptor instead.
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{9}
}

type ListBlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *Blog `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
}

func (x *ListBlogResponse) Reset() {
	*x = ListBlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_crud_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogResponse) ProtoMessage() {}

func (x *ListBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_crud_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogResponse.ProtoReflect.Descriptor instead.
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return file_proto_crud_proto_rawDescGZIP(), []int{10}
}

func (x *ListBlogResponse) GetBlog() *Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

var File_proto_crud_proto protoreflect.FileDescriptor

var file_proto_crud_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x22, 0x63, 0x0a, 0x04, 0x42,
	0x6c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x36, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x37, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x22, 0x2a, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x35, 0x0a,
	0x10, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04,
	0x62, 0x6c, 0x6f, 0x67, 0x22, 0x36, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70,
	0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x37, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c,
	0x6f, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f,
	0x67, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6c, 0x6f,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x6c, 0x6f,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70,
	0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x32, 0xf0, 0x02, 0x0a,
	0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x6f, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x63, 0x72, 0x75,
	0x64, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f,
	0x67, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12,
	0x1a, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x72,
	0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x72, 0x75, 0x64, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_crud_proto_rawDescOnce sync.Once
	file_proto_crud_proto_rawDescData = file_proto_crud_proto_rawDesc
)

func file_proto_crud_proto_rawDescGZIP() []byte {
	file_proto_crud_proto_rawDescOnce.Do(func() {
		file_proto_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_crud_proto_rawDescData)
	})
	return file_proto_crud_proto_rawDescData
}

var file_proto_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_crud_proto_goTypes = []interface{}{
	(*Blog)(nil),               // 0: crudapi.Blog
	(*CreateBlogRequest)(nil),  // 1: crudapi.CreateBlogRequest
	(*CreateBlogResponse)(nil), // 2: crudapi.CreateBlogResponse
	(*ReadBlogRequest)(nil),    // 3: crudapi.ReadBlogRequest
	(*ReadBlogResponse)(nil),   // 4: crudapi.ReadBlogResponse
	(*UpdateBlogRequest)(nil),  // 5: crudapi.UpdateBlogRequest
	(*UpdateBlogResponse)(nil), // 6: crudapi.UpdateBlogResponse
	(*DeleteBlogRequest)(nil),  // 7: crudapi.DeleteBlogRequest
	(*DeleteBlogResponse)(nil), // 8: crudapi.DeleteBlogResponse
	(*ListBlogRequest)(nil),    // 9: crudapi.ListBlogRequest
	(*ListBlogResponse)(nil),   // 10: crudapi.ListBlogResponse
}
var file_proto_crud_proto_depIdxs = []int32{
	0,  // 0: crudapi.CreateBlogRequest.blog:type_name -> crudapi.Blog
	0,  // 1: crudapi.CreateBlogResponse.blog:type_name -> crudapi.Blog
	0,  // 2: crudapi.ReadBlogResponse.blog:type_name -> crudapi.Blog
	0,  // 3: crudapi.UpdateBlogRequest.blog:type_name -> crudapi.Blog
	0,  // 4: crudapi.UpdateBlogResponse.blog:type_name -> crudapi.Blog
	0,  // 5: crudapi.ListBlogResponse.blog:type_name -> crudapi.Blog
	1,  // 6: crudapi.BlogServoce.CreateBlog:input_type -> crudapi.CreateBlogRequest
	3,  // 7: crudapi.BlogServoce.ReadBlog:input_type -> crudapi.ReadBlogRequest
	5,  // 8: crudapi.BlogServoce.UpdateBlog:input_type -> crudapi.UpdateBlogRequest
	7,  // 9: crudapi.BlogServoce.DeleteBlog:input_type -> crudapi.DeleteBlogRequest
	9,  // 10: crudapi.BlogServoce.ListBlog:input_type -> crudapi.ListBlogRequest
	2,  // 11: crudapi.BlogServoce.CreateBlog:output_type -> crudapi.CreateBlogResponse
	4,  // 12: crudapi.BlogServoce.ReadBlog:output_type -> crudapi.ReadBlogResponse
	6,  // 13: crudapi.BlogServoce.UpdateBlog:output_type -> crudapi.UpdateBlogResponse
	8,  // 14: crudapi.BlogServoce.DeleteBlog:output_type -> crudapi.DeleteBlogResponse
	10, // 15: crudapi.BlogServoce.ListBlog:output_type -> crudapi.ListBlogResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_crud_proto_init() }
func file_proto_crud_proto_init() {
	if File_proto_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_proto_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogRequest); i {
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
		file_proto_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlogResponse); i {
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
		file_proto_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogRequest); i {
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
		file_proto_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBlogResponse); i {
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
		file_proto_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogRequest); i {
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
		file_proto_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlogResponse); i {
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
		file_proto_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogRequest); i {
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
		file_proto_crud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlogResponse); i {
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
		file_proto_crud_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogRequest); i {
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
		file_proto_crud_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogResponse); i {
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
			RawDescriptor: file_proto_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_crud_proto_goTypes,
		DependencyIndexes: file_proto_crud_proto_depIdxs,
		MessageInfos:      file_proto_crud_proto_msgTypes,
	}.Build()
	File_proto_crud_proto = out.File
	file_proto_crud_proto_rawDesc = nil
	file_proto_crud_proto_goTypes = nil
	file_proto_crud_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogServoceClient is the client API for BlogServoce service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServoceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogServoce_ListBlogClient, error)
}

type blogServoceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServoceClient(cc grpc.ClientConnInterface) BlogServoceClient {
	return &blogServoceClient{cc}
}

func (c *blogServoceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/crudapi.BlogServoce/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServoceClient) ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/crudapi.BlogServoce/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServoceClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error) {
	out := new(UpdateBlogResponse)
	err := c.cc.Invoke(ctx, "/crudapi.BlogServoce/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServoceClient) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error) {
	out := new(DeleteBlogResponse)
	err := c.cc.Invoke(ctx, "/crudapi.BlogServoce/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServoceClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogServoce_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogServoce_serviceDesc.Streams[0], "/crudapi.BlogServoce/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServoceListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogServoce_ListBlogClient interface {
	Recv() (*ListBlogResponse, error)
	grpc.ClientStream
}

type blogServoceListBlogClient struct {
	grpc.ClientStream
}

func (x *blogServoceListBlogClient) Recv() (*ListBlogResponse, error) {
	m := new(ListBlogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServoceServer is the server API for BlogServoce service.
type BlogServoceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogResponse, error)
	ListBlog(*ListBlogRequest, BlogServoce_ListBlogServer) error
}

// UnimplementedBlogServoceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServoceServer struct {
}

func (*UnimplementedBlogServoceServer) CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServoceServer) ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServoceServer) UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServoceServer) DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (*UnimplementedBlogServoceServer) ListBlog(*ListBlogRequest, BlogServoce_ListBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}

func RegisterBlogServoceServer(s *grpc.Server, srv BlogServoceServer) {
	s.RegisterService(&_BlogServoce_serviceDesc, srv)
}

func _BlogServoce_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServoceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crudapi.BlogServoce/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServoceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogServoce_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServoceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crudapi.BlogServoce/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServoceServer).ReadBlog(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogServoce_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServoceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crudapi.BlogServoce/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServoceServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogServoce_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServoceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crudapi.BlogServoce/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServoceServer).DeleteBlog(ctx, req.(*DeleteBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogServoce_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServoceServer).ListBlog(m, &blogServoceListBlogServer{stream})
}

type BlogServoce_ListBlogServer interface {
	Send(*ListBlogResponse) error
	grpc.ServerStream
}

type blogServoceListBlogServer struct {
	grpc.ServerStream
}

func (x *blogServoceListBlogServer) Send(m *ListBlogResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogServoce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crudapi.BlogServoce",
	HandlerType: (*BlogServoceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogServoce_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogServoce_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogServoce_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogServoce_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _BlogServoce_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/crud.proto",
}
