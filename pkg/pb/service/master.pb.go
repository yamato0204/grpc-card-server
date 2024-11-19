// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: service/master.proto

package service

import (
	master "github.com/yamato0204/grpc-card-server/pkg/pb/master"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_service_master_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_service_master_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_service_master_proto_rawDescGZIP(), []int{0}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards      []*master.Card      `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	Characters []*master.Character `protobuf:"bytes,2,rep,name=characters,proto3" json:"characters,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	mi := &file_service_master_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_master_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_service_master_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllResponse) GetCards() []*master.Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

func (x *GetAllResponse) GetCharacters() []*master.Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

type GetCharacterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Characters []*master.Character `protobuf:"bytes,1,rep,name=characters,proto3" json:"characters,omitempty"`
}

func (x *GetCharacterResponse) Reset() {
	*x = GetCharacterResponse{}
	mi := &file_service_master_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCharacterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterResponse) ProtoMessage() {}

func (x *GetCharacterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_master_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterResponse.ProtoReflect.Descriptor instead.
func (*GetCharacterResponse) Descriptor() ([]byte, []int) {
	return file_service_master_proto_rawDescGZIP(), []int{2}
}

func (x *GetCharacterResponse) GetCharacters() []*master.Character {
	if x != nil {
		return x.Characters
	}
	return nil
}

type GetCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*master.Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *GetCardResponse) Reset() {
	*x = GetCardResponse{}
	mi := &file_service_master_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardResponse) ProtoMessage() {}

func (x *GetCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_master_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardResponse.ProtoReflect.Descriptor instead.
func (*GetCardResponse) Descriptor() ([]byte, []int) {
	return file_service_master_proto_rawDescGZIP(), []int{3}
}

func (x *GetCardResponse) GetCards() []*master.Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

var File_service_master_proto protoreflect.FileDescriptor

var file_service_master_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x11, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22, 0x49, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x22, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x32, 0xbc,
	0x01, 0x0a, 0x0d, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6d, 0x61,
	0x74, 0x6f, 0x30, 0x32, 0x30, 0x34, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x64,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_master_proto_rawDescOnce sync.Once
	file_service_master_proto_rawDescData = file_service_master_proto_rawDesc
)

func file_service_master_proto_rawDescGZIP() []byte {
	file_service_master_proto_rawDescOnce.Do(func() {
		file_service_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_master_proto_rawDescData)
	})
	return file_service_master_proto_rawDescData
}

var file_service_master_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_master_proto_goTypes = []any{
	(*Empty)(nil),                // 0: service.Empty
	(*GetAllResponse)(nil),       // 1: service.GetAllResponse
	(*GetCharacterResponse)(nil), // 2: service.GetCharacterResponse
	(*GetCardResponse)(nil),      // 3: service.GetCardResponse
	(*master.Card)(nil),          // 4: master.Card
	(*master.Character)(nil),     // 5: master.Character
}
var file_service_master_proto_depIdxs = []int32{
	4, // 0: service.GetAllResponse.cards:type_name -> master.Card
	5, // 1: service.GetAllResponse.characters:type_name -> master.Character
	5, // 2: service.GetCharacterResponse.characters:type_name -> master.Character
	4, // 3: service.GetCardResponse.cards:type_name -> master.Card
	0, // 4: service.MasterService.GetAll:input_type -> service.Empty
	0, // 5: service.MasterService.GetCard:input_type -> service.Empty
	0, // 6: service.MasterService.GetCharacter:input_type -> service.Empty
	1, // 7: service.MasterService.GetAll:output_type -> service.GetAllResponse
	3, // 8: service.MasterService.GetCard:output_type -> service.GetCardResponse
	2, // 9: service.MasterService.GetCharacter:output_type -> service.GetCharacterResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_master_proto_init() }
func file_service_master_proto_init() {
	if File_service_master_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_master_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_master_proto_goTypes,
		DependencyIndexes: file_service_master_proto_depIdxs,
		MessageInfos:      file_service_master_proto_msgTypes,
	}.Build()
	File_service_master_proto = out.File
	file_service_master_proto_rawDesc = nil
	file_service_master_proto_goTypes = nil
	file_service_master_proto_depIdxs = nil
}