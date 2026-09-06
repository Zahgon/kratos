package main

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	contextPackage       = protogen.GoImportPath("context")
	transportHTTPPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v3/transport/http")
	httpBodyFullName     = protoreflect.FullName("google.api.HttpBody")
)

var methodSets = make(map[string]int)

func generateFile(gen *protogen.Plugin, file *protogen.File, omitempty bool, omitemptyPrefix string) *protogen.GeneratedFile {
	_ = "STUB: not implemented"
	return nil
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, omitempty bool, omitemptyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func genService(_ *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, omitempty bool, omitemptyPrefix string) {
	_ = "STUB: not implemented"
	return
}

func hasHTTPRule(services []*protogen.Service) bool { _ = "STUB: not implemented"; return false }

func buildHTTPRule(g *protogen.GeneratedFile, service *protogen.Service, m *protogen.Method, rule *annotations.HttpRule, omitemptyPrefix string) *methodDesc {
	_ = "STUB: not implemented"
	return nil
}

func buildMethodDesc(g *protogen.GeneratedFile, m *protogen.Method, method, path string) *methodDesc {
	_ = "STUB: not implemented"
	return nil
}

func isHTTPBodyField(fd protoreflect.FieldDescriptor) bool { _ = "STUB: not implemented"; return false }

func isHTTPBodyMessage(md protoreflect.MessageDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

func buildPathVars(path string) (res map[string]*string) { _ = "STUB: not implemented"; return nil }

func replacePath(name string, value string, path string) string {
	_ = "STUB: not implemented"
	return ""
}

func pathTemplateRegex(value string) string { _ = "STUB: not implemented"; return "" }

func camelCaseVars(s string) string { _ = "STUB: not implemented"; return "" }

func camelCase(s string) string { _ = "STUB: not implemented"; return "" }

func isASCIILower(c byte) bool { _ = "STUB: not implemented"; return false }

func isASCIIDigit(c byte) bool { _ = "STUB: not implemented"; return false }

func protocVersion(gen *protogen.Plugin) string { _ = "STUB: not implemented"; return "" }

const deprecationComment = "// Deprecated: Do not use."
