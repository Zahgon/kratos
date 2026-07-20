package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	errorsPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v3/errors")
	fmtPackage    = protogen.GoImportPath("fmt")
)

var enCases = cases.Title(language.AmericanEnglish, cases.NoLower)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	_ = "STUB: not implemented"
	return nil
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	_ = "STUB: not implemented"
	return
}

func genErrorsReason(_ *protogen.Plugin, _ *protogen.File, g *protogen.GeneratedFile, enum *protogen.Enum) bool {
	_ = "STUB: not implemented"
	return false
}

func case2Camel(name string) string { _ = "STUB: not implemented"; return "" }
