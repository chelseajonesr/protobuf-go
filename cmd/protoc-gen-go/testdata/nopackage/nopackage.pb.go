// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nopackage/nopackage.proto

package nopackage

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

type Enum int32

const (
	Enum_ZERO Enum = 0
)

func (e Enum) Type() protoreflect.EnumType {
	return xxx_File_nopackage_nopackage_proto_enumTypes[0]
}
func (e Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var Enum_name = map[int32]string{
	0: "ZERO",
}

var Enum_value = map[string]int32{
	"ZERO": 0,
}

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return proto.EnumName(Enum_name, int32(x))
}

func (x *Enum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Enum_value, data, "Enum")
	if err != nil {
		return err
	}
	*x = Enum(value)
	return nil
}

func (Enum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_nopackage_nopackage_proto_rawdesc_gzipped, []int{0}
}

type Message struct {
	StringField          *string  `protobuf:"bytes,1,opt,name=string_field,json=stringField" json:"string_field,omitempty"`
	EnumField            *Enum    `protobuf:"varint,2,opt,name=enum_field,json=enumField,enum=Enum,def=0" json:"enum_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) ProtoReflect() protoreflect.Message {
	return xxx_File_nopackage_nopackage_proto_messageTypes[0].MessageOf(m)
}
func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return xxx_File_nopackage_nopackage_proto_rawdesc_gzipped, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

const Default_Message_EnumField Enum = Enum_ZERO

func (m *Message) GetStringField() string {
	if m != nil && m.StringField != nil {
		return *m.StringField
	}
	return ""
}

func (m *Message) GetEnumField() Enum {
	if m != nil && m.EnumField != nil {
		return *m.EnumField
	}
	return Default_Message_EnumField
}

func init() {
	proto.RegisterFile("nopackage/nopackage.proto", xxx_File_nopackage_nopackage_proto_rawdesc_gzipped)
	proto.RegisterEnum("Enum", Enum_name, Enum_value)
	proto.RegisterType((*Message)(nil), "Message")
}

var xxx_File_nopackage_nopackage_proto_rawdesc = []byte{
	// 135 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x6f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2a, 0x0a, 0x0a, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x10, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a,
	0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00,
}

var xxx_File_nopackage_nopackage_proto_rawdesc_gzipped = protoapi.CompressGZIP(xxx_File_nopackage_nopackage_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_nopackage_nopackage_proto protoreflect.FileDescriptor

var xxx_File_nopackage_nopackage_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_nopackage_nopackage_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_nopackage_nopackage_proto_goTypes = []interface{}{
	(Enum)(0),       // 0: Enum
	(*Message)(nil), // 1: Message
}
var xxx_File_nopackage_nopackage_proto_depIdxs = []int32{
	0, // Message.enum_field:type_name -> Enum
}

func init() { xxx_File_nopackage_nopackage_proto_init() }
func xxx_File_nopackage_nopackage_proto_init() {
	if File_nopackage_nopackage_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 1)
	File_nopackage_nopackage_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_nopackage_nopackage_proto_rawdesc,
		GoTypes:            xxx_File_nopackage_nopackage_proto_goTypes,
		DependencyIndexes:  xxx_File_nopackage_nopackage_proto_depIdxs,
		EnumOutputTypes:    xxx_File_nopackage_nopackage_proto_enumTypes,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_nopackage_nopackage_proto_goTypes[1:][:1]
	for i, mt := range messageTypes {
		xxx_File_nopackage_nopackage_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_nopackage_nopackage_proto_messageTypes[i].PBType = mt
	}
	xxx_File_nopackage_nopackage_proto_goTypes = nil
	xxx_File_nopackage_nopackage_proto_depIdxs = nil
}
