// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: services/datastore/protocols/banned.proto

package protocols

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

type BanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*BanResponse_Ban
	//	*BanResponse_Error
	Result isBanResponse_Result `protobuf_oneof:"result"`
}

func (x *BanResponse) Reset() {
	*x = BanResponse{}
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BanResponse) ProtoMessage() {}

func (x *BanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BanResponse.ProtoReflect.Descriptor instead.
func (*BanResponse) Descriptor() ([]byte, []int) {
	return file_services_datastore_protocols_banned_proto_rawDescGZIP(), []int{0}
}

func (m *BanResponse) GetResult() isBanResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *BanResponse) GetBan() *Ban {
	if x, ok := x.GetResult().(*BanResponse_Ban); ok {
		return x.Ban
	}
	return nil
}

func (x *BanResponse) GetError() string {
	if x, ok := x.GetResult().(*BanResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isBanResponse_Result interface {
	isBanResponse_Result()
}

type BanResponse_Ban struct {
	Ban *Ban `protobuf:"bytes,1,opt,name=ban,proto3,oneof"`
}

type BanResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*BanResponse_Ban) isBanResponse_Result() {}

func (*BanResponse_Error) isBanResponse_Result() {}

type Ban struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP        string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Active    bool   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Ban) Reset() {
	*x = Ban{}
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ban) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ban) ProtoMessage() {}

func (x *Ban) ProtoReflect() protoreflect.Message {
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ban.ProtoReflect.Descriptor instead.
func (*Ban) Descriptor() ([]byte, []int) {
	return file_services_datastore_protocols_banned_proto_rawDescGZIP(), []int{1}
}

func (x *Ban) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Ban) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Ban) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type ListBan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ban []*Ban `protobuf:"bytes,1,rep,name=ban,proto3" json:"ban,omitempty"`
}

func (x *ListBan) Reset() {
	*x = ListBan{}
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBan) ProtoMessage() {}

func (x *ListBan) ProtoReflect() protoreflect.Message {
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBan.ProtoReflect.Descriptor instead.
func (*ListBan) Descriptor() ([]byte, []int) {
	return file_services_datastore_protocols_banned_proto_rawDescGZIP(), []int{2}
}

func (x *ListBan) GetBan() []*Ban {
	if x != nil {
		return x.Ban
	}
	return nil
}

type SearchBan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Active *bool  `protobuf:"varint,2,opt,name=active,proto3,oneof" json:"active,omitempty"`
}

func (x *SearchBan) Reset() {
	*x = SearchBan{}
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBan) ProtoMessage() {}

func (x *SearchBan) ProtoReflect() protoreflect.Message {
	mi := &file_services_datastore_protocols_banned_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBan.ProtoReflect.Descriptor instead.
func (*SearchBan) Descriptor() ([]byte, []int) {
	return file_services_datastore_protocols_banned_proto_rawDescGZIP(), []int{3}
}

func (x *SearchBan) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchBan) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

var File_services_datastore_protocols_banned_proto protoreflect.FileDescriptor

var file_services_datastore_protocols_banned_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x62, 0x61, 0x6e,
	0x6e, 0x65, 0x64, 0x22, 0x50, 0x0a, 0x0b, 0x42, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x62, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x2e, 0x42, 0x61, 0x6e, 0x48, 0x00, 0x52, 0x03,
	0x62, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4c, 0x0a, 0x03, 0x42, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x22, 0x28, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x61, 0x6e, 0x12, 0x1d,
	0x0a, 0x03, 0x62, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x2e, 0x42, 0x61, 0x6e, 0x52, 0x03, 0x62, 0x61, 0x6e, 0x22, 0x49, 0x0a,
	0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1b, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x6d, 0x6e, 0x69, 0x65, 0x6c, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_datastore_protocols_banned_proto_rawDescOnce sync.Once
	file_services_datastore_protocols_banned_proto_rawDescData = file_services_datastore_protocols_banned_proto_rawDesc
)

func file_services_datastore_protocols_banned_proto_rawDescGZIP() []byte {
	file_services_datastore_protocols_banned_proto_rawDescOnce.Do(func() {
		file_services_datastore_protocols_banned_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_datastore_protocols_banned_proto_rawDescData)
	})
	return file_services_datastore_protocols_banned_proto_rawDescData
}

var file_services_datastore_protocols_banned_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_datastore_protocols_banned_proto_goTypes = []any{
	(*BanResponse)(nil), // 0: banned.BanResponse
	(*Ban)(nil),         // 1: banned.Ban
	(*ListBan)(nil),     // 2: banned.ListBan
	(*SearchBan)(nil),   // 3: banned.SearchBan
}
var file_services_datastore_protocols_banned_proto_depIdxs = []int32{
	1, // 0: banned.BanResponse.ban:type_name -> banned.Ban
	1, // 1: banned.ListBan.ban:type_name -> banned.Ban
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_datastore_protocols_banned_proto_init() }
func file_services_datastore_protocols_banned_proto_init() {
	if File_services_datastore_protocols_banned_proto != nil {
		return
	}
	file_services_datastore_protocols_banned_proto_msgTypes[0].OneofWrappers = []any{
		(*BanResponse_Ban)(nil),
		(*BanResponse_Error)(nil),
	}
	file_services_datastore_protocols_banned_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_datastore_protocols_banned_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_datastore_protocols_banned_proto_goTypes,
		DependencyIndexes: file_services_datastore_protocols_banned_proto_depIdxs,
		MessageInfos:      file_services_datastore_protocols_banned_proto_msgTypes,
	}.Build()
	File_services_datastore_protocols_banned_proto = out.File
	file_services_datastore_protocols_banned_proto_rawDesc = nil
	file_services_datastore_protocols_banned_proto_goTypes = nil
	file_services_datastore_protocols_banned_proto_depIdxs = nil
}