// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_2/m4.proto

package test_a_2

import (
	proto "github.com/golang/protobuf/proto"
	protoapi "github.com/golang/protobuf/protoapi"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type M4 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M4) ProtoReflect() protoreflect.Message {
	return xxx_File_imports_test_a_2_m4_proto_messageTypes[0].MessageOf(m)
}
func (m *M4) Reset()         { *m = M4{} }
func (m *M4) String() string { return proto.CompactTextString(m) }
func (*M4) ProtoMessage()    {}
func (*M4) Descriptor() ([]byte, []int) {
	return xxx_File_imports_test_a_2_m4_proto_rawdesc_gzipped, []int{0}
}

func (m *M4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M4.Unmarshal(m, b)
}
func (m *M4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M4.Marshal(b, m, deterministic)
}
func (m *M4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M4.Merge(m, src)
}
func (m *M4) XXX_Size() int {
	return xxx_messageInfo_M4.Size(m)
}
func (m *M4) XXX_DiscardUnknown() {
	xxx_messageInfo_M4.DiscardUnknown(m)
}

var xxx_messageInfo_M4 proto.InternalMessageInfo

func init() {
	proto.RegisterFile("imports/test_a_2/m4.proto", xxx_File_imports_test_a_2_m4_proto_rawdesc_gzipped)
	proto.RegisterType((*M4)(nil), "test.a.M4")
}

var xxx_File_imports_test_a_2_m4_proto_rawdesc = []byte{
	// 119 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x5f, 0x32, 0x2f, 0x6d, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x61, 0x22, 0x04, 0x0a, 0x02, 0x4d, 0x34, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x5f, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var xxx_File_imports_test_a_2_m4_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_imports_test_a_2_m4_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_imports_test_a_2_m4_proto protoreflect.FileDescriptor

var xxx_File_imports_test_a_2_m4_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_imports_test_a_2_m4_proto_goTypes = []interface{}{
	(*M4)(nil), // 0: test.a.M4
}
var xxx_File_imports_test_a_2_m4_proto_depIdxs = []int32{}

func init() { xxx_File_imports_test_a_2_m4_proto_init() }
func xxx_File_imports_test_a_2_m4_proto_init() {
	if File_imports_test_a_2_m4_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_imports_test_a_2_m4_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_imports_test_a_2_m4_proto_rawdesc,
		GoTypes:            xxx_File_imports_test_a_2_m4_proto_goTypes,
		DependencyIndexes:  xxx_File_imports_test_a_2_m4_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_imports_test_a_2_m4_proto_goTypes[0:][:1]
	for i, mt := range messageTypes {
		xxx_File_imports_test_a_2_m4_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_imports_test_a_2_m4_proto_messageTypes[i].PBType = mt
	}
	xxx_File_imports_test_a_2_m4_proto_goTypes = nil
	xxx_File_imports_test_a_2_m4_proto_depIdxs = nil
}
