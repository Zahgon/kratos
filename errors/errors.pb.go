package errors

import (
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason        string                 `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata      map[string]string      `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() { _ = "STUB: not implemented"; return }

func (x *Status) String() string { _ = "STUB: not implemented"; return "" }

func (*Status) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *Status) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*Status) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *Status) GetCode() int32 { _ = "STUB: not implemented"; return 0 }

func (x *Status) GetReason() string { _ = "STUB: not implemented"; return "" }

func (x *Status) GetMessage() string { _ = "STUB: not implemented"; return "" }

func (x *Status) GetMetadata() map[string]string { _ = "STUB: not implemented"; return nil }

var file_errors_errors_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1108,
		Name:          "errors.default_code",
		Tag:           "varint,1108,opt,name=default_code",
		Filename:      "errors/errors.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1109,
		Name:          "errors.code",
		Tag:           "varint,1109,opt,name=code",
		Filename:      "errors/errors.proto",
	},
}

var (
	E_DefaultCode = &file_errors_errors_proto_extTypes[0]
)

var (
	E_Code = &file_errors_errors_proto_extTypes[1]
)

var File_errors_errors_proto protoreflect.FileDescriptor

const file_errors_errors_proto_rawDesc = "" +
	"\n" +
	"\x13errors/errors.proto\x12\x06errors\x1a google/protobuf/descriptor.proto\"\xc5\x01\n" +
	"\x06Status\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x16\n" +
	"\x06reason\x18\x02 \x01(\tR\x06reason\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x128\n" +
	"\bmetadata\x18\x04 \x03(\v2\x1c.errors.Status.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01:@\n" +
	"\fdefault_code\x12\x1c.google.protobuf.EnumOptions\x18\xd4\b \x01(\x05R\vdefaultCode:6\n" +
	"\x04code\x12!.google.protobuf.EnumValueOptions\x18\xd5\b \x01(\x05R\x04codeBY\n" +
	"\x18com.github.kratos.errorsP\x01Z,github.com/go-kratos/kratos/v3/errors;errors\xa2\x02\fKratosErrorsb\x06proto3"

var (
	file_errors_errors_proto_rawDescOnce sync.Once
	file_errors_errors_proto_rawDescData []byte
)

func file_errors_errors_proto_rawDescGZIP() []byte { _ = "STUB: not implemented"; return nil }

var file_errors_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_errors_errors_proto_goTypes = []any{
	(*Status)(nil),
	nil,
	(*descriptorpb.EnumOptions)(nil),
	(*descriptorpb.EnumValueOptions)(nil),
}
var file_errors_errors_proto_depIdxs = []int32{
	1,
	2,
	3,
	3,
	3,
	3,
	1,
	0,
}

func init()                          { file_errors_errors_proto_init() }
func file_errors_errors_proto_init() { _ = "STUB: not implemented"; return }
