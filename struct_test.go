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
	Name   *string `fake:"{firstname}"`
	Const  *string `fake:"ABC"`
}

type StructArray struct {
	Bars   []*Basic
	Builds []BuiltIn
	Skips  []string    `fake:"skip"`
	Empty  []*Basic    `fakesize:"0"`
	Multy  []*Template `fakesize:"3"`
}

type NestedArray struct {
	NA []StructArray `fake:"size=2"`
}

func TestStructBasic(t *testing.T) {
	Seed(11)

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
	Seed(11)

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
	Seed(11)

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
	Seed(11)

	var template Template
	Struct(&template)
	if *template.Number == "" {
		t.Error("template number should set to number value")
	}
	if *template.Name == "" {
		t.Error("template name should set to name value")
	}
	if *template.Const != "ABC" {
		t.Error("template const should set to fixed value")
	}
}

func TestStructArray(t *testing.T) {
	Seed(11)

	var sa StructArray
	Struct(&sa)
	if len(sa.Bars) != 1 {
		t.Error("sa slice Bars is not populated")
	}
	if len(sa.Builds) != 10 {
		t.Error("sa slice Builds is not populated")
	}
	if len(sa.Empty) != 0 {
		t.Error("sa slice Empty is populated")
	}
	if len(sa.Skips) != 0 {
		t.Error("sa slice Empty is populated")
	}
	if len(sa.Multy) != 3 {
		t.Error("sa slice Multy is not fully populated")
	}
}

func TestStructNestedArray(t *testing.T) {
	Seed(11)

	var na NestedArray
	Struct(&na)
	if len(na.NA) != 2 {
		t.Error("na slice NA is not populated")
	}
	for _, elem := range na.NA {
		if len(elem.Builds) == 0 {
			t.Error("a slice Builds is not populated")
		}
		if len(elem.Empty) != 0 {
			t.Error("a slice Empty is populated")
		}
		if len(elem.Skips) != 0 {
			t.Error("a slice Empty is populated")
		}
		if len(elem.Multy) != 3 {
			t.Error("a slice Multy is not fully populated")
		}
	}
}

func Example_struct() {
	Seed(11)

	type Foo struct {
		Bar     string
		Int     int
		Pointer *int
		Name    string  `fake:"{firstname}"`
		Number  string  `fake:"{number:1,10}"`
		Skip    *string `fake:"skip"`
	}

	var f Foo
	Struct(&f)

	fmt.Printf("%s\n", f.Bar)
	fmt.Printf("%d\n", f.Int)
	fmt.Printf("%d\n", *f.Pointer)
	fmt.Printf("%v\n", f.Name)
	fmt.Printf("%v\n", f.Number)
	fmt.Printf("%v\n", f.Skip)

	// Output: bRMaRxHki
	// -8576773003117070818
	// -7054675846543980602
	// Enrique
	// 4
	// <nil>
}

func Example_array() {
	Seed(11)

	type Foo struct {
		Bar    string
		Int    int
		Name   string  `fake:"{firstname}"`
		Number string  `fake:"{number:1,10}"`
		Skip   *string `fake:"skip"`
	}

	type FooMany struct {
		Foos  []Foo    `fakesize:"1"`
		Names []string `fake:"{firstname}" fakesize:"3"`
	}

	var fm FooMany
	Struct(&fm)

	fmt.Printf("%v\n", fm.Foos)
	fmt.Printf("%v\n", fm.Names)

	// Output:
	// [{MaRxH -6396943744266753635 0xc000016320 Dawn 10 <nil>}]
	// [Bertha Brannon Zackary]
}
