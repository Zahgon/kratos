package form

import (
	"errors"
	"net/url"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const fieldSeparator = "."

var errInvalidFormatMapKey = errors.New("invalid formatting for map key")

func DecodeValues(msg proto.Message, values url.Values) error {
	_ = "STUB: not implemented"
	return nil
}

func populateFieldValues(v protoreflect.Message, fieldPath []string, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func getFieldDescriptor(v protoreflect.Message, fieldName string) protoreflect.FieldDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.FieldDescriptor)
}

func getDescriptorByFieldAndName(fields protoreflect.FieldDescriptors, fieldName string) protoreflect.FieldDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.FieldDescriptor)
}

func populateField(fd protoreflect.FieldDescriptor, v protoreflect.Message, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func populateRepeatedField(fd protoreflect.FieldDescriptor, list protoreflect.List, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func populateMapField(fd protoreflect.FieldDescriptor, mp protoreflect.Map, fieldPath []string, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseField(fd protoreflect.FieldDescriptor, value string) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

func parseMessage(md protoreflect.MessageDescriptor, value string) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

//nolint:mnd

func jsonSnakeCase(s string) string { _ = "STUB: not implemented"; return "" }

func isASCIIUpper(c byte) bool { _ = "STUB: not implemented"; return false }

func parseURLQueryMapKey(key string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
