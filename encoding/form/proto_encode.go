package form

import (
	"net/url"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func EncodeValues(msg any) (url.Values, error) {
	_ = "STUB: not implemented"
	return *new(url.Values), nil
}

func encodeByField(u url.Values, path string, m protoreflect.Message) (finalErr error) {
	_ = "STUB: not implemented"
	return nil
}

func encodeRepeatedField(fieldDescriptor protoreflect.FieldDescriptor, list protoreflect.List) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeMapField(fieldDescriptor protoreflect.FieldDescriptor, mp protoreflect.Map) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeField(fieldDescriptor protoreflect.FieldDescriptor, value protoreflect.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func encodeMessage(msgDescriptor protoreflect.MessageDescriptor, value protoreflect.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func EncodeFieldMask(m protoreflect.Message) (query string) { _ = "STUB: not implemented"; return "" }

func jsonCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

func isASCIILower(c byte) bool { _ = "STUB: not implemented"; return false }
