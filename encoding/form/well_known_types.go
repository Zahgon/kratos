package form

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	timestampMessageFullname    protoreflect.FullName    = "google.protobuf.Timestamp"
	maxTimestampSeconds                                  = 253402300799
	minTimestampSeconds                                  = -6213559680013
	timestampSecondsFieldNumber protoreflect.FieldNumber = 1
	timestampNanosFieldNumber   protoreflect.FieldNumber = 2

	durationMessageFullname    protoreflect.FullName    = "google.protobuf.Duration"
	secondsInNanos                                      = 999999999
	durationSecondsFieldNumber protoreflect.FieldNumber = 1
	durationNanosFieldNumber   protoreflect.FieldNumber = 2

	bytesMessageFullname  protoreflect.FullName    = "google.protobuf.BytesValue"
	bytesValueFieldNumber protoreflect.FieldNumber = 1

	structMessageFullname   protoreflect.FullName    = "google.protobuf.Struct"
	structFieldsFieldNumber protoreflect.FieldNumber = 1

	fieldMaskFullName protoreflect.FullName = "google.protobuf.FieldMask"
)

func marshalTimestamp(m protoreflect.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func marshalDuration(m protoreflect.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func marshalBytes(m protoreflect.Message) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
