// Code generated by protoc-gen-go. DO NOT EDIT.
// comments/deprecated.proto is a deprecated file.

package comments

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

type xxx_DeprecatedEnum DeprecatedEnum

func (e DeprecatedEnum) ProtoReflect() protoreflect.Enum {
	return (xxx_DeprecatedEnum)(e)
}
func (e xxx_DeprecatedEnum) Type() protoreflect.EnumType {
	return xxx_Deprecated_ProtoFile_EnumTypes[0]
}
func (e xxx_DeprecatedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return proto.EnumName(DeprecatedEnum_name, int32(x))
}

func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0336e614ee2de5f7, []int{0}
}

// Deprecated: Do not use.
type DeprecatedMessage struct {
	DeprecatedField      string   `protobuf:"bytes,1,opt,name=deprecated_field,json=deprecatedField,proto3" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_DeprecatedMessage struct{ m *DeprecatedMessage }

func (m *DeprecatedMessage) ProtoReflect() protoreflect.Message {
	return xxx_DeprecatedMessage{m}
}
func (m xxx_DeprecatedMessage) Type() protoreflect.MessageType {
	return xxx_Deprecated_ProtoFile_MessageTypes[0].Type
}
func (m xxx_DeprecatedMessage) KnownFields() protoreflect.KnownFields {
	return xxx_Deprecated_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_DeprecatedMessage) UnknownFields() protoreflect.UnknownFields {
	return xxx_Deprecated_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_DeprecatedMessage) Interface() protoreflect.ProtoMessage {
	return m.m
}

func (m *DeprecatedMessage) Reset()         { *m = DeprecatedMessage{} }
func (m *DeprecatedMessage) String() string { return proto.CompactTextString(m) }
func (*DeprecatedMessage) ProtoMessage()    {}
func (*DeprecatedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0336e614ee2de5f7, []int{0}
}

func (m *DeprecatedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedMessage.Unmarshal(m, b)
}
func (m *DeprecatedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedMessage.Marshal(b, m, deterministic)
}
func (m *DeprecatedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedMessage.Merge(m, src)
}
func (m *DeprecatedMessage) XXX_Size() int {
	return xxx_messageInfo_DeprecatedMessage.Size(m)
}
func (m *DeprecatedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedMessage proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *DeprecatedMessage) GetDeprecatedField() string {
	if m != nil {
		return m.DeprecatedField
	}
	return ""
}

func init() {
	proto.RegisterFile("comments/deprecated.proto", fileDescriptor_0336e614ee2de5f7)
	proto.RegisterEnum("goproto.protoc.comments.DeprecatedEnum", DeprecatedEnum_name, DeprecatedEnum_value)
	proto.RegisterType((*DeprecatedMessage)(nil), "goproto.protoc.comments.DeprecatedMessage")
}

var fileDescriptor_0336e614ee2de5f7 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x49, 0x2d, 0x28, 0x4a, 0x4d, 0x4e, 0x2c, 0x49, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0xcf, 0x07, 0x33, 0x20, 0xdc, 0x64, 0x3d, 0x98,
	0x4a, 0x25, 0x37, 0x2e, 0x41, 0x17, 0xb8, 0x62, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21,
	0x5d, 0x2e, 0x01, 0x84, 0x09, 0xf1, 0x69, 0x99, 0xa9, 0x39, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x4e, 0x4c, 0x12, 0x8c, 0x41, 0xfc, 0x08, 0x39, 0x37, 0x90, 0x94, 0x15, 0x93, 0x04, 0xa3,
	0x96, 0x06, 0x17, 0x1f, 0xc2, 0x1c, 0xd7, 0xbc, 0xd2, 0x5c, 0x21, 0x21, 0x2e, 0x2e, 0x17, 0xd7,
	0x80, 0x20, 0x57, 0x67, 0xc7, 0x10, 0x57, 0x17, 0x01, 0x06, 0x29, 0x26, 0x0e, 0x46, 0x29, 0x26,
	0x09, 0x46, 0x27, 0xb7, 0x28, 0xc7, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0x90, 0x23, 0xf4, 0xd3,
	0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xc1, 0xce, 0x4a, 0x2a, 0x4d, 0xd3, 0x2f, 0x33, 0xd2, 0x4f,
	0xce, 0x4d, 0x81, 0xf0, 0x93, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd3, 0xf3, 0xf5, 0x4b, 0x52, 0x8b,
	0x4b, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0x61, 0xce, 0xde, 0xc1, 0xc8, 0x98, 0xc4, 0x06, 0x56, 0x63,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0xab, 0x43, 0x93, 0xf6, 0x00, 0x00, 0x00,
}

func init() {
	xxx_Deprecated_ProtoFile_FileDesc.Enums = xxx_Deprecated_ProtoFile_EnumDescs[0:1]
	xxx_Deprecated_ProtoFile_FileDesc.Messages = xxx_Deprecated_ProtoFile_MessageDescs[0:1]
	var err error
	Deprecated_ProtoFile, err = prototype.NewFile(&xxx_Deprecated_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Deprecated_ProtoFile protoreflect.FileDescriptor

var xxx_Deprecated_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto3,
	Path:    "comments/deprecated.proto",
	Package: "goproto.protoc.comments",
}
var xxx_Deprecated_ProtoFile_EnumTypes = [1]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_Deprecated_ProtoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return DeprecatedEnum(n)
		},
	),
}
var xxx_Deprecated_ProtoFile_EnumDescs = [1]prototype.Enum{
	{
		Name: "DeprecatedEnum",
		Values: []prototype.EnumValue{
			{Name: "DEPRECATED", Number: 0},
		},
	},
}
var xxx_Deprecated_ProtoFile_MessageTypes = [1]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Deprecated_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(DeprecatedMessage)
		},
	)},
}
var xxx_Deprecated_ProtoFile_MessageDescs = [1]prototype.Message{
	{
		Name: "DeprecatedMessage",
		Fields: []prototype.Field{
			{
				Name:        "deprecated_field",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "deprecatedField",
			},
		},
	},
}
