// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: proto/transport.proto

package proto

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

type AuthorBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64 `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
}

func (x *AuthorBooksRequest) Reset() {
	*x = AuthorBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorBooksRequest) ProtoMessage() {}

func (x *AuthorBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorBooksRequest.ProtoReflect.Descriptor instead.
func (*AuthorBooksRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorBooksRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

type AuthorBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *AuthorBooksResponse) Reset() {
	*x = AuthorBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorBooksResponse) ProtoMessage() {}

func (x *AuthorBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorBooksResponse.ProtoReflect.Descriptor instead.
func (*AuthorBooksResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type BookAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId int64 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *BookAuthorsRequest) Reset() {
	*x = BookAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAuthorsRequest) ProtoMessage() {}

func (x *BookAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookAuthorsRequest.ProtoReflect.Descriptor instead.
func (*BookAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{2}
}

func (x *BookAuthorsRequest) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type BookAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *BookAuthorsResponse) Reset() {
	*x = BookAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookAuthorsResponse) ProtoMessage() {}

func (x *BookAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookAuthorsResponse.ProtoReflect.Descriptor instead.
func (*BookAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_proto_transport_proto_rawDescGZIP(), []int{3}
}

func (x *BookAuthorsResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

var File_proto_transport_proto protoreflect.FileDescriptor

var file_proto_transport_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x12, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x13,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x2d, 0x0a,
	0x12, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x13,
	0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x74, 0x61, 0x6b,
	0x65, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transport_proto_rawDescOnce sync.Once
	file_proto_transport_proto_rawDescData = file_proto_transport_proto_rawDesc
)

func file_proto_transport_proto_rawDescGZIP() []byte {
	file_proto_transport_proto_rawDescOnce.Do(func() {
		file_proto_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transport_proto_rawDescData)
	})
	return file_proto_transport_proto_rawDescData
}

var file_proto_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_transport_proto_goTypes = []interface{}{
	(*AuthorBooksRequest)(nil),  // 0: book_service.AuthorBooksRequest
	(*AuthorBooksResponse)(nil), // 1: book_service.AuthorBooksResponse
	(*BookAuthorsRequest)(nil),  // 2: book_service.BookAuthorsRequest
	(*BookAuthorsResponse)(nil), // 3: book_service.BookAuthorsResponse
	(*Book)(nil),                // 4: book_service.Book
	(*Author)(nil),              // 5: book_service.Author
}
var file_proto_transport_proto_depIdxs = []int32{
	4, // 0: book_service.AuthorBooksResponse.books:type_name -> book_service.Book
	5, // 1: book_service.BookAuthorsResponse.authors:type_name -> book_service.Author
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_transport_proto_init() }
func file_proto_transport_proto_init() {
	if File_proto_transport_proto != nil {
		return
	}
	file_proto_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorBooksRequest); i {
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
		file_proto_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorBooksResponse); i {
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
		file_proto_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookAuthorsRequest); i {
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
		file_proto_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookAuthorsResponse); i {
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
			RawDescriptor: file_proto_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_transport_proto_goTypes,
		DependencyIndexes: file_proto_transport_proto_depIdxs,
		MessageInfos:      file_proto_transport_proto_msgTypes,
	}.Build()
	File_proto_transport_proto = out.File
	file_proto_transport_proto_rawDesc = nil
	file_proto_transport_proto_goTypes = nil
	file_proto_transport_proto_depIdxs = nil
}
