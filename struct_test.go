package gofakeit

import (
	"fmt"
	rand "math/rand"
	"testing"
	"time"
)

type Basic struct {
	s string
	S string
}

type embeder struct {
	A  string
	AA []string
	embedee
}

type embedee struct {
	B  string   `fake:"{firstname}"`
	BA []string `fake:"{firstname}" fakesize:"5"`
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
	embeder
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

type Function struct {
	Number *string `fake:"#"`
	Name   *string `fake:"{firstname}"`
	Const  *string `fake:"ABC"`
}

type StructArray struct {
	Bars      []*Basic
	Builds    []BuiltIn
	Skips     []string  `fake:"skip"`
	Strings   []string  `fake:"{firstname}" fakesize:"3"`
	SetLen    [5]string `fake:"{firstname}"`
	SubStr    [][]string
	SubStrLen [][2]string
	Empty     []*Basic    `fakesize:"0"`
	Multy     []*Function `fakesize:"3"`
}

type NestedArray struct {
	NA []StructArray `fakesize:"2"`
}

func ExampleStruct() {
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

func ExampleFaker_Struct() {
	f := New(11)

	type Foo struct {
		Bar     string
		Int     int
		Pointer *int
		Name    string  `fake:"{firstname}"`
		Number  string  `fake:"{number:1,10}"`
		Skip    *string `fake:"skip"`
	}

	var foo Foo
	f.Struct(&foo)

	fmt.Printf("%s\n", foo.Bar)
	fmt.Printf("%d\n", foo.Int)
	fmt.Printf("%d\n", *foo.Pointer)
	fmt.Printf("%v\n", foo.Name)
	fmt.Printf("%v\n", foo.Number)
	fmt.Printf("%v\n", foo.Skip)

	// Output: bRMaRxHki
	// -8576773003117070818
	// -7054675846543980602
	// Enrique
	// 4
	// <nil>
}

func ExampleStruct_array() {
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
	// [{bRMaRxHki -8576773003117070818 Carole 6 <nil>}]
	// [Dawn Zachery Amie]
}

func ExampleFaker_Struct_array() {
	f := New(11)

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
	f.Struct(&fm)

	fmt.Printf("%v\n", fm.Foos)
	fmt.Printf("%v\n", fm.Names)

	// Output:
	// [{bRMaRxHki -8576773003117070818 Carole 6 <nil>}]
	// [Dawn Zachery Amie]
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
		t.Error("nested struct bar should not be populated due to unexported variable")
	}
	if nested.embeder.A == "" || nested.embeder.AA == nil {
		t.Error("nested embeder fields should be populated")
	}
	if nested.embeder.embedee.B == "" || nested.embeder.embedee.BA == nil || len(nested.embeder.embedee.BA) != 5 {
		t.Error("nested embeder fields should be populated")
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

func TestStructWithFunction(t *testing.T) {
	Seed(11)

	var function Function
	Struct(&function)
	if *function.Number == "" {
		t.Error("function number should set to number value")
	}
	if *function.Name == "" {
		t.Error("function name should set to name value")
	}
	if *function.Const != "ABC" {
		t.Error("function const should set to fixed value")
	}
}

func TestStructArray(t *testing.T) {
	Seed(11)

	var sa StructArray
	Struct(&sa)
	if len(sa.Bars) == 0 {
		t.Error("sa slice Bars is not populated")
	}
	if len(sa.Builds) == 0 {
		t.Error("sa slice Builds is not populated")
	}
	if len(sa.Strings) != 3 {
		t.Error("sa strings should have a length of 3")
	}
	if len(sa.Strings) == 3 {
		for i, s := range sa.Strings {
			if s == "" {
				t.Errorf("sa Strings index %d was empty", i)
			}
		}
	}
	if len(sa.SetLen) != 5 {
		for i, s := range sa.SetLen {
			if s == "" {
				t.Errorf("sa SetLen index %d was empty", i)
			}
		}
	}
	if len(sa.SubStr) == 0 {
		for i, s := range sa.SubStr {
			for ii, ss := range s {
				if ss == "" {
					t.Errorf("sa SubStr index %d index %d was empty", i, ii)
				}
			}
		}
	}
	if len(sa.SubStrLen) == 0 {
		for _, s := range sa.SubStrLen {
			for _, ss := range s {
				if len(ss) != 2 {
					t.Errorf("sa SubStrLen sub length should have been 2")
				}
			}
		}
	}
	if len(sa.Empty) != 0 {
		t.Error("sa slice Empty is populated")
	}
	if len(sa.Skips) != 0 {
		t.Error("sa slice Skips is populated")
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

func TestStructToInt(t *testing.T) {
	Seed(11)

	var si struct {
		Int         int
		IntConst    int8  `fake:"-123"`
		IntGenerate int64 `fake:"{number:1,10}"`
	}
	Struct(&si)
	if si.Int == 0 {
		t.Error("Int should be an actual number greater than 0")
	}
	if si.IntConst != -123 {
		t.Errorf("IntConst should be -123 and instead got %d", si.IntConst)
	}
	if si.IntGenerate == 0 || si.IntGenerate > 10 {
		t.Errorf("IntGenerate should be between 1-10 but got %d", si.IntGenerate)
	}
}

func TestStructToUint(t *testing.T) {
	Seed(11)

	var su struct {
		Uint         uint
		UintConst    uint8  `fake:"123"`
		UintGenerate uint64 `fake:"{number:1,10}"`
	}
	Struct(&su)
	if su.Uint == 0 {
		t.Error("Uint should be an actual number greater than 0")
	}
	if su.UintConst != 123 {
		t.Errorf("UintConst should be 123 and instead got %d", su.UintConst)
	}
	if su.UintGenerate == 0 || su.UintGenerate > 10 {
		t.Errorf("UintGenerate should be between 1-10 but got %d", su.UintGenerate)
	}
}

func TestStructToFloat(t *testing.T) {
	Seed(11)

	var sf struct {
		Float         float32
		FloatConst    float64 `fake:"123.456789"`
		FloatGenerate float32 `fake:"{latitude}"`
	}
	Struct(&sf)
	if sf.Float == 0 {
		t.Error("Float should be an actual number greater than 0")
	}
	if sf.FloatConst != 123.456789 {
		t.Errorf("FloatConst should be 123.456.789 and instead got %f", sf.FloatConst)
	}
	if sf.FloatGenerate != 54.887310 {
		t.Errorf("FloatGenerate should be 31.477726 but got %f", sf.FloatGenerate)
	}
}

func TestStructToBool(t *testing.T) {
	Seed(11)

	var sf struct {
		Bool         bool
		BoolConst    bool `fake:"true"`
		BoolGenerate bool `fake:"{bool}"`
	}
	Struct(&sf)
	if sf.Bool == true {
		t.Error("Bool should be true got false")
	}
	if sf.BoolConst != true {
		t.Errorf("BoolConst should be true got false")
	}
	if sf.BoolGenerate != true {
		t.Errorf("BoolGenerate should be true got false")
	}
}

func TestStructToDateTime(t *testing.T) {
	Seed(11)

	var datetime struct {
		Simple    time.Time
		Tag       time.Time `fake:"{date}"`
		TagFormat time.Time `fake:"{number:1900,1950}-12-05" format:"2006-01-02"`
	}
	Struct(&datetime)
	if datetime.Simple.String() != "1908-12-07 04:14:25.685339029 +0000 UTC" {
		t.Errorf("Simple should be 1908-12-07 04:14:25.685339029 +0000 UTC and instead got %s", datetime.Simple.String())
	}
	if datetime.Tag.String() != "1979-05-21 05:49:35 +0000 UTC" {
		t.Errorf("Tag should be 1979-05-21 05:49:35 +0000 UTC and instead got %s", datetime.Tag.String())
	}
	if datetime.TagFormat.String() != "1943-12-05 00:00:00 +0000 UTC" {
		t.Errorf("TagFormat should be 1943-12-05 00:00:00 +0000 UTC and instead got %s", datetime.TagFormat.String())
	}
}

func TestStructSetSubStruct(t *testing.T) {
	type Sub struct {
		Str string
		Num int
		Flo float32
	}

	type Foo struct {
		Sub Sub `fake:"{setstruct}"`
	}

	AddFuncLookup("setstruct", Info{
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return Sub{Str: "string", Num: 123, Flo: 123.456}, nil
		},
	})

	Seed(11)

	var f Foo
	Struct(&f)

	if f.Sub.Str != "string" {
		t.Errorf("str didnt equal string, got %s", f.Sub.Str)
	}
	if f.Sub.Num != 123 {
		t.Errorf("num didnt equal 123, got %d", f.Sub.Num)
	}
	if f.Sub.Flo != 123.456 {
		t.Errorf("flo didnt equal 123.456, got %v", f.Sub.Flo)
	}

	RemoveFuncLookup("setstruct")
}
