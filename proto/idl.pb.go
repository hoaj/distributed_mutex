// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: idl.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

var File_idl_proto protoreflect.FileDescriptor

var file_idl_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22, 0x19, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x32, 0x61, 0x0a, 0x07, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x12, 0x2b, 0x0a,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_proto_rawDescOnce sync.Once
	file_idl_proto_rawDescData = file_idl_proto_rawDesc
)

func file_idl_proto_rawDescGZIP() []byte {
	file_idl_proto_rawDescOnce.Do(func() {
		file_idl_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_proto_rawDescData)
	})
	return file_idl_proto_rawDescData
}

var file_idl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_idl_proto_goTypes = []interface{}{
	(*Node)(nil),  // 0: proto.Node
	(*Token)(nil), // 1: proto.Token
	(*Ack)(nil),   // 2: proto.Ack
}
var file_idl_proto_depIdxs = []int32{
	0, // 0: proto.Central.requestToken:input_type -> proto.Node
	1, // 1: proto.Central.returnToken:input_type -> proto.Token
	1, // 2: proto.Central.requestToken:output_type -> proto.Token
	2, // 3: proto.Central.returnToken:output_type -> proto.Ack
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_proto_init() }
func file_idl_proto_init() {
	if File_idl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_idl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_idl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
			RawDescriptor: file_idl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_proto_goTypes,
		DependencyIndexes: file_idl_proto_depIdxs,
		MessageInfos:      file_idl_proto_msgTypes,
	}.Build()
	File_idl_proto = out.File
	file_idl_proto_rawDesc = nil
	file_idl_proto_goTypes = nil
	file_idl_proto_depIdxs = nil
}
