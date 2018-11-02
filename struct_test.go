package gofakeit

import (
	"fmt"
	"testing"
)

type Basic struct {
	s string
	S string
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
}

type BuiltIn struct {
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	Float32 *float32
	Float64 *float64
	Bool    *bool
}

type Template struct {
	Number *string `fake:"#"`
	Name   *string `fake:"{person.first}"`
}

func TestStructBasic(t *testing.T) {
	var basic Basic
	Struct(&basic)
	if basic.s != "" {
		t.Error("unexported field is not populated")
	}
	if basic.S == "" {
		t.Error("exported field is populated")
	}
}

func TestStructNested(t *testing.T) {
	var nested Nested
	Struct(&nested)
	if nested.A == "" {
		t.Error("exported string field is populated")
	}
	if nested.B == nil {
		t.Error("exported struct field is populated")
	}
	if nested.B.S == "" {
		t.Error("nested struct string field is not populated")
	}
	if nested.bar != nil {
		t.Error("nested struct bar should be be populated due to unexported variable")
	}
}

func TestStructBuiltInTypes(t *testing.T) {
	var builtIn BuiltIn
	Struct(&builtIn)
	if builtIn.Uint8 == nil {
		t.Error("builtIn Uint8 was not set")
	}
	if builtIn.Uint16 == nil {
		t.Error("builtIn Uint16 was not set")
	}
	if builtIn.Uint32 == nil {
		t.Error("builtIn Uint32 was not set")
	}
	if builtIn.Uint64 == nil {
		t.Error("builtIn Uint64 was not set")
	}
	if builtIn.Int == nil {
		t.Error("builtIn int was not set")
	}
	if builtIn.Int8 == nil {
		t.Error("builtIn int8 was not set")
	}
	if builtIn.Int16 == nil {
		t.Error("builtIn int16 was not set")
	}
	if builtIn.Int32 == nil {
		t.Error("builtIn int32 was not set")
	}
	if builtIn.Int64 == nil {
		t.Error("builtIn int64 was not set")
	}
	if builtIn.Float32 == nil {
		t.Error("builtIn float32 was not set")
	}
	if builtIn.Float64 == nil {
		t.Error("builtIn float64 was not set")
	}
	if builtIn.Bool == nil {
		t.Error("builtIn bool was not set")
	}
}

func TestStructWithTemplate(t *testing.T) {
	var template Template
	Struct(&template)
	if *template.Number == "" {
		t.Error("template number should set to number value")
	}
	if *template.Name == "" {
		t.Error("template number should set to number value")
	}
}

func ExampleStruct() {
	Seed(11)
	type Foo struct {
		Bar     string
		Baz     string
		Int     int
		Pointer *int
		Skip    *string `fake:"skip"`
	}
	var f Foo
	Struct(&f)
	fmt.Printf("%s\n", f.Bar)
	fmt.Printf("%s\n", f.Baz)
	fmt.Printf("%d\n", f.Int)
	fmt.Printf("%d\n", *f.Pointer)
	fmt.Printf("%v\n", f.Skip)
	// Output: gbrmarxhkijbptapwyj
	// dnsmkgtlxwnqhqclayk
	// -5858358572185296359
	// -8038678955577270446
	// <nil>
}
