// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/bookstore/v1/bookstore.proto

package v1

import (
	context "context"
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

type BookEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Authors []string `protobuf:"bytes,3,rep,name=authors,proto3" json:"authors,omitempty"`
	Press   string   `protobuf:"bytes,4,opt,name=press,proto3" json:"press,omitempty"`
}

func (x *BookEntity) Reset() {
	*x = BookEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookEntity) ProtoMessage() {}

func (x *BookEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookEntity.ProtoReflect.Descriptor instead.
func (*BookEntity) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *BookEntity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BookEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BookEntity) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *BookEntity) GetPress() string {
	if x != nil {
		return x.Press
	}
	return ""
}

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookReq *BookEntity `protobuf:"bytes,1,opt,name=bookReq,proto3" json:"bookReq,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{1}
}

func (x *BookRequest) GetBookReq() *BookEntity {
	if x != nil {
		return x.BookReq
	}
	return nil
}

type BookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *BookResp) Reset() {
	*x = BookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookResp) ProtoMessage() {}

func (x *BookResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookResp.ProtoReflect.Descriptor instead.
func (*BookResp) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *BookResp) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *BookResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetABookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetABookReq) Reset() {
	*x = GetABookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABookReq) ProtoMessage() {}

func (x *GetABookReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABookReq.ProtoReflect.Descriptor instead.
func (*GetABookReq) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *GetABookReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetABookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err  int32       `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Book *BookEntity `protobuf:"bytes,3,opt,name=book,proto3" json:"book,omitempty"` // ???????????????
}

func (x *GetABookResp) Reset() {
	*x = GetABookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetABookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetABookResp) ProtoMessage() {}

func (x *GetABookResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetABookResp.ProtoReflect.Descriptor instead.
func (*GetABookResp) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *GetABookResp) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *GetABookResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetABookResp) GetBook() *BookEntity {
	if x != nil {
		return x.Book
	}
	return nil
}

type GetAllBookReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllBookReq) Reset() {
	*x = GetAllBookReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBookReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBookReq) ProtoMessage() {}

func (x *GetAllBookReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBookReq.ProtoReflect.Descriptor instead.
func (*GetAllBookReq) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{5}
}

type GetAllBookResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err   int32         `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg   string        `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Books []*BookEntity `protobuf:"bytes,3,rep,name=books,proto3" json:"books,omitempty"` // ??????????????????
}

func (x *GetAllBookResp) Reset() {
	*x = GetAllBookResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllBookResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllBookResp) ProtoMessage() {}

func (x *GetAllBookResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_bookstore_v1_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllBookResp.ProtoReflect.Descriptor instead.
func (*GetAllBookResp) Descriptor() ([]byte, []int) {
	return file_api_bookstore_v1_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllBookResp) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *GetAllBookResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetAllBookResp) GetBooks() []*BookEntity {
	if x != nil {
		return x.Books
	}
	return nil
}

var File_api_bookstore_v1_bookstore_proto protoreflect.FileDescriptor

var file_api_bookstore_v1_bookstore_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x60, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x41, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x07, 0x62, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x22, 0x2e, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x22, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x32, 0xdd, 0x01,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x12, 0x5a,
	0x10, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_bookstore_v1_bookstore_proto_rawDescOnce sync.Once
	file_api_bookstore_v1_bookstore_proto_rawDescData = file_api_bookstore_v1_bookstore_proto_rawDesc
)

func file_api_bookstore_v1_bookstore_proto_rawDescGZIP() []byte {
	file_api_bookstore_v1_bookstore_proto_rawDescOnce.Do(func() {
		file_api_bookstore_v1_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_bookstore_v1_bookstore_proto_rawDescData)
	})
	return file_api_bookstore_v1_bookstore_proto_rawDescData
}

var file_api_bookstore_v1_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_bookstore_v1_bookstore_proto_goTypes = []interface{}{
	(*BookEntity)(nil),     // 0: bookstore.v1.BookEntity
	(*BookRequest)(nil),    // 1: bookstore.v1.BookRequest
	(*BookResp)(nil),       // 2: bookstore.v1.BookResp
	(*GetABookReq)(nil),    // 3: bookstore.v1.GetABookReq
	(*GetABookResp)(nil),   // 4: bookstore.v1.GetABookResp
	(*GetAllBookReq)(nil),  // 5: bookstore.v1.GetAllBookReq
	(*GetAllBookResp)(nil), // 6: bookstore.v1.GetAllBookResp
}
var file_api_bookstore_v1_bookstore_proto_depIdxs = []int32{
	0, // 0: bookstore.v1.BookRequest.bookReq:type_name -> bookstore.v1.BookEntity
	0, // 1: bookstore.v1.GetABookResp.book:type_name -> bookstore.v1.BookEntity
	0, // 2: bookstore.v1.GetAllBookResp.books:type_name -> bookstore.v1.BookEntity
	1, // 3: bookstore.v1.Bookstore.CreateBook:input_type -> bookstore.v1.BookRequest
	3, // 4: bookstore.v1.Bookstore.GetBook:input_type -> bookstore.v1.GetABookReq
	5, // 5: bookstore.v1.Bookstore.GetAllBook:input_type -> bookstore.v1.GetAllBookReq
	2, // 6: bookstore.v1.Bookstore.CreateBook:output_type -> bookstore.v1.BookResp
	4, // 7: bookstore.v1.Bookstore.GetBook:output_type -> bookstore.v1.GetABookResp
	6, // 8: bookstore.v1.Bookstore.GetAllBook:output_type -> bookstore.v1.GetAllBookResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_bookstore_v1_bookstore_proto_init() }
func file_api_bookstore_v1_bookstore_proto_init() {
	if File_api_bookstore_v1_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_bookstore_v1_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookEntity); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookResp); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABookReq); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetABookResp); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBookReq); i {
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
		file_api_bookstore_v1_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllBookResp); i {
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
			RawDescriptor: file_api_bookstore_v1_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_bookstore_v1_bookstore_proto_goTypes,
		DependencyIndexes: file_api_bookstore_v1_bookstore_proto_depIdxs,
		MessageInfos:      file_api_bookstore_v1_bookstore_proto_msgTypes,
	}.Build()
	File_api_bookstore_v1_bookstore_proto = out.File
	file_api_bookstore_v1_bookstore_proto_rawDesc = nil
	file_api_bookstore_v1_bookstore_proto_goTypes = nil
	file_api_bookstore_v1_bookstore_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookstoreClient is the client API for Bookstore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookstoreClient interface {
	CreateBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResp, error)
	GetBook(ctx context.Context, in *GetABookReq, opts ...grpc.CallOption) (*GetABookResp, error)
	GetAllBook(ctx context.Context, in *GetAllBookReq, opts ...grpc.CallOption) (*GetAllBookResp, error)
}

type bookstoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreClient(cc grpc.ClientConnInterface) BookstoreClient {
	return &bookstoreClient{cc}
}

func (c *bookstoreClient) CreateBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := c.cc.Invoke(ctx, "/bookstore.v1.Bookstore/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetBook(ctx context.Context, in *GetABookReq, opts ...grpc.CallOption) (*GetABookResp, error) {
	out := new(GetABookResp)
	err := c.cc.Invoke(ctx, "/bookstore.v1.Bookstore/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreClient) GetAllBook(ctx context.Context, in *GetAllBookReq, opts ...grpc.CallOption) (*GetAllBookResp, error) {
	out := new(GetAllBookResp)
	err := c.cc.Invoke(ctx, "/bookstore.v1.Bookstore/GetAllBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServer is the server API for Bookstore service.
type BookstoreServer interface {
	CreateBook(context.Context, *BookRequest) (*BookResp, error)
	GetBook(context.Context, *GetABookReq) (*GetABookResp, error)
	GetAllBook(context.Context, *GetAllBookReq) (*GetAllBookResp, error)
}

// UnimplementedBookstoreServer can be embedded to have forward compatible implementations.
type UnimplementedBookstoreServer struct {
}

func (*UnimplementedBookstoreServer) CreateBook(context.Context, *BookRequest) (*BookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedBookstoreServer) GetBook(context.Context, *GetABookReq) (*GetABookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookstoreServer) GetAllBook(context.Context, *GetAllBookReq) (*GetAllBookResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBook not implemented")
}

func RegisterBookstoreServer(s *grpc.Server, srv BookstoreServer) {
	s.RegisterService(&_Bookstore_serviceDesc, srv)
}

func _Bookstore_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.v1.Bookstore/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).CreateBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetABookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.v1.Bookstore/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetBook(ctx, req.(*GetABookReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bookstore_GetAllBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBookReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServer).GetAllBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bookstore.v1.Bookstore/GetAllBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServer).GetAllBook(ctx, req.(*GetAllBookReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bookstore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bookstore.v1.Bookstore",
	HandlerType: (*BookstoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _Bookstore_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _Bookstore_GetBook_Handler,
		},
		{
			MethodName: "GetAllBook",
			Handler:    _Bookstore_GetAllBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/bookstore/v1/bookstore.proto",
}
