// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	prototype "google.golang.org/protobuf/reflect/prototype"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

type AnnotationsTestEnum int32

const (
	AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE AnnotationsTestEnum = 0
)

var AnnotationsTestEnum_name = map[int32]string{
	0: "ANNOTATIONS_TEST_ENUM_VALUE",
}

var AnnotationsTestEnum_value = map[string]int32{
	"ANNOTATIONS_TEST_ENUM_VALUE": 0,
}

func (x AnnotationsTestEnum) Enum() *AnnotationsTestEnum {
	p := new(AnnotationsTestEnum)
	*p = x
	return p
}

func (x AnnotationsTestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnnotationsTestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_annotations_annotations_proto_enumTypes[0].EnumDescriptor
}

func (AnnotationsTestEnum) Type() protoreflect.EnumType {
	return &file_annotations_annotations_proto_enumTypes[0]
}

func (x AnnotationsTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AnnotationsTestEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AnnotationsTestEnum(num)
	return nil
}

// Deprecated: Use AnnotationsTestEnum.Type instead.
func (AnnotationsTestEnum) EnumDescriptor() ([]byte, []int) {
	return file_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

type AnnotationsTestMessage struct {
	state                protoimpl.MessageState
	AnnotationsTestField *string `protobuf:"bytes,1,opt,name=AnnotationsTestField" json:"AnnotationsTestField,omitempty"`
	sizeCache            protoimpl.SizeCache
	unknownFields        protoimpl.UnknownFields
}

func (x *AnnotationsTestMessage) Reset() {
	*x = AnnotationsTestMessage{}
}

func (x *AnnotationsTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsTestMessage) ProtoMessage() {}

func (x *AnnotationsTestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_annotations_annotations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationsTestMessage.ProtoReflect.Type instead.
func (*AnnotationsTestMessage) Descriptor() ([]byte, []int) {
	return file_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationsTestMessage) GetAnnotationsTestField() string {
	if x != nil && x.AnnotationsTestField != nil {
		return *x.AnnotationsTestField
	}
	return ""
}

var File_annotations_annotations_proto protoreflect.FileDescriptor

var file_annotations_annotations_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x36, 0x0a, 0x13, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
}

var (
	file_annotations_annotations_proto_rawDescOnce sync.Once
	file_annotations_annotations_proto_rawDescData = file_annotations_annotations_proto_rawDesc
)

func file_annotations_annotations_proto_rawDescGZIP() []byte {
	file_annotations_annotations_proto_rawDescOnce.Do(func() {
		file_annotations_annotations_proto_rawDescData = protoimpl.X.CompressGZIP(file_annotations_annotations_proto_rawDescData)
	})
	return file_annotations_annotations_proto_rawDescData
}

var file_annotations_annotations_proto_enumTypes = make([]prototype.Enum, 1)
var file_annotations_annotations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_annotations_annotations_proto_goTypes = []interface{}{
	(AnnotationsTestEnum)(0),       // 0: goproto.protoc.annotations.AnnotationsTestEnum
	(*AnnotationsTestMessage)(nil), // 1: goproto.protoc.annotations.AnnotationsTestMessage
}
var file_annotations_annotations_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_annotations_annotations_proto_init() }
func file_annotations_annotations_proto_init() {
	if File_annotations_annotations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_annotations_annotations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationsTestMessage); i {
			case 0:
				return &v.state
			case 2:
				return &v.sizeCache
			case 3:
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
			RawDescriptor: file_annotations_annotations_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_annotations_annotations_proto_goTypes,
		DependencyIndexes: file_annotations_annotations_proto_depIdxs,
		MessageInfos:      file_annotations_annotations_proto_msgTypes,
	}.Build()
	File_annotations_annotations_proto = out.File
	file_annotations_annotations_proto_enumTypes = out.Enums
	file_annotations_annotations_proto_rawDesc = nil
	file_annotations_annotations_proto_goTypes = nil
	file_annotations_annotations_proto_depIdxs = nil
}
