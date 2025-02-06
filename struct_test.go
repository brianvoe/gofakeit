package gofakeit

import (
	"fmt"
	"net"
	"regexp"
	"strings"
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

type StrArray []string
type IntArray []int

type StructArray struct {
	Bars      []*Basic
	Builds    []BuiltIn
	Skips     []string  `fake:"skip"`
	SkipsAlt  []string  `fake:"-"`
	Strings   []string  `fake:"{firstname}" fakesize:"3"`
	SetLen    [5]string `fake:"{firstname}"`
	SubStr    [][]string
	SubStrLen [][2]string
	Empty     []*Basic    `fakesize:"0"`
	Multy     []*Function `fakesize:"3"`
	StrArray  StrArray    `fake:"{firstname}" fakesize:"3"`
	IntArray  *IntArray   `fake:"{number:1,10}" fakesize:"3"`
}

type NestedArray struct {
	NA []StructArray `fakesize:"2"`
}

func ExampleStruct() {
	Seed(11)

	type Bar struct {
		Name   string
		Number int
		Float  float32
	}

	type Foo struct {
		Str        string
		Int        int
		Pointer    *int
		Name       string            `fake:"{firstname}"`
		Number     string            `fake:"{number:1,10}"`
		Skip       *string           `fake:"skip"`
		SkipAlt    *string           `fake:"-"`
		Array      []string          `fakesize:"2"`
		ArrayRange []string          `fakesize:"2,6"`
		Map        map[string]string `fakesize:"2"`
		MapRange   map[string]string `fakesize:"2,4"`
		Bar        Bar
	}

	var f Foo
	Struct(&f)

	fmt.Printf("%s\n", f.Str)
	fmt.Printf("%d\n", f.Int)
	fmt.Printf("%d\n", *f.Pointer)
	fmt.Printf("%v\n", f.Name)
	fmt.Printf("%v\n", f.Number)
	fmt.Printf("%v\n", f.Skip)
	fmt.Printf("%v\n", f.SkipAlt)
	fmt.Printf("%v\n", f.Array)
	fmt.Printf("%v\n", f.ArrayRange)
	fmt.Printf("%v\n", f.Map)
	fmt.Printf("%v\n", f.MapRange)
	fmt.Printf("%+v\n", f.Bar)

	// Output: sKKbMlbxqu
	// 1365845625386394310
	// 4415419407664574903
	// Tomasa
	// 10
	// <nil>
	// <nil>
	// [RWVlPmwRAe wVHB]
	// [rxXDcxFe kfvChMC KfmBwRtn]
	// map[gyJrZ:tfZAk oOEImWl:NXYnuKiGab]
	// map[WigEnVzBy:JJXMkG aHqIpyVs:UxvheGVE pYyDuWC:yzjqsX]
	// {Name:LfLRuhFRHQ Number:5737279083624806933 Float:0.59020466}
}

func ExampleFaker_Struct() {
	fake := New(11)

	type Bar struct {
		Name   string
		Number int
		Float  float32
	}

	type Foo struct {
		Str        string
		Int        int
		Pointer    *int
		Name       string            `fake:"{firstname}"`
		Number     string            `fake:"{number:1,10}"`
		Skip       *string           `fake:"skip"`
		Array      []string          `fakesize:"2"`
		ArrayRange []string          `fakesize:"2,6"`
		Map        map[string]string `fakesize:"2"`
		MapRange   map[string]string `fakesize:"2,4"`
		Bar        Bar
	}

	var f Foo
	fake.Struct(&f)

	fmt.Printf("%s\n", f.Str)
	fmt.Printf("%d\n", f.Int)
	fmt.Printf("%d\n", *f.Pointer)
	fmt.Printf("%v\n", f.Name)
	fmt.Printf("%v\n", f.Number)
	fmt.Printf("%v\n", f.Skip)
	fmt.Printf("%v\n", f.Array)
	fmt.Printf("%v\n", f.Map)
	fmt.Printf("%v\n", f.MapRange)
	fmt.Printf("%+v\n", f.Bar)

	// Output: sKKbMlbxqu
	// 1365845625386394310
	// 4415419407664574903
	// Tomasa
	// 10
	// <nil>
	// [RWVlPmwRAe wVHB]
	// map[gyJrZ:tfZAk oOEImWl:NXYnuKiGab]
	// map[WigEnVzBy:JJXMkG aHqIpyVs:UxvheGVE pYyDuWC:yzjqsX]
	// {Name:LfLRuhFRHQ Number:5737279083624806933 Float:0.59020466}
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
		Foos       []Foo    `fakesize:"1"`
		Names      []string `fake:"{firstname}" fakesize:"3"`
		NamesRange []string `fake:"{firstname}" fakesize:"3,6"`
	}

	var fm FooMany
	Struct(&fm)

	fmt.Printf("%v\n", fm.Foos)
	fmt.Printf("%v\n", fm.Names)
	fmt.Printf("%v\n", fm.NamesRange)

	// Output: [{sKKbMlbxqu 1365845625386394310 Myrtis 10 <nil>}]
	// [Tomasa Sunny Ernest]
	// [Hassan Missouri Elisabeth Nedra Tiana Evan]
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
		Foos       []Foo    `fakesize:"1"`
		Names      []string `fake:"{firstname}" fakesize:"3"`
		NamesRange []string `fake:"{firstname}" fakesize:"3,6"`
	}

	var fm FooMany
	f.Struct(&fm)

	fmt.Printf("%v\n", fm.Foos)
	fmt.Printf("%v\n", fm.Names)
	fmt.Printf("%v\n", fm.NamesRange)

	// Output: [{sKKbMlbxqu 1365845625386394310 Myrtis 10 <nil>}]
	// [Tomasa Sunny Ernest]
	// [Hassan Missouri Elisabeth Nedra Tiana Evan]
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

func TestStructPointer(t *testing.T) {
	type Info struct {
		Person        PersonInfo  `fake:"{person}"`
		PersonPointer *PersonInfo `fake:"{person}"`
	}

	var info Info
	err := Struct(&info)
	if err != nil {
		t.Fatal(err)
	}

	if info.Person.FirstName == "" {
		t.Error("Person wasnt properly set")
	}

	if info.PersonPointer == nil {
		t.Error("PersonPointer wasnt properly set")
	}
	if info.PersonPointer.FirstName == "" {
		t.Error("PersonPointer firstname wasnt properly set")
	}
}

func TestCustomArrayType(t *testing.T) {
	Seed(11)
	AddFuncLookup("customType", Info{
		Category:    "custom",
		Description: "Random int array",
		Example:     "[1]",
		Output:      "CustomType",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			data := make([]int, 1)
			data[0] = 42
			return data, nil
		},
	})
	defer RemoveFuncLookup("customType")

	AddFuncLookup("customByte", Info{
		Category:    "custom",
		Description: "Random byte",
		Example:     "[1]",
		Output:      "byte",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			data := byte(42)
			return data, nil
		},
	})

	type customType []int
	var wct struct {
		WithTag         customType `fake:"{customType}"`
		WithOutTag      customType
		ArrayWithTag    []customType `fake:"{customType}"`
		ArrayWithOutTag []customType
		IntSlice        []int
		Char            byte
		CharWithTag     byte `fake:"{customByte}"`
	}
	err := Struct(&wct)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(wct.WithTag) != 1 {
		t.Error("wct slice WithTag is not populated from custom function")
	}

	if wct.WithTag[0] != 42 {
		t.Errorf("wct slice WithTag: want 42, got %d", wct.WithTag[0])
	}

	if len(wct.WithOutTag) == 0 {
		t.Error("wct slice WithoutTag is not populated")
	}

	if wct.IntSlice[0] == wct.WithOutTag[0] {
		t.Error("two different slices should have different values")
	}

	if wct.ArrayWithTag[0][0] != 42 {
		t.Error("wct ArrayWithTag is not populated from custom function")
	}

	if wct.ArrayWithOutTag[0][0] == 0 {
		t.Errorf("wct ArrayWithoutTag did not set value got:%+v", wct.ArrayWithOutTag)
	}

	if len(wct.IntSlice) == 0 {
		t.Error("wct slice IntSlice is not populated")
	}

	if wct.Char == byte(0) {
		t.Error("wct Char is not populated")
	}

	if wct.CharWithTag != byte(42) {
		t.Error("wct CharWithTag is not populated from custom function")
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
	if len(sa.StrArray) != 3 {
		t.Error("sa slice StrArray is not fully populated")
	}
	if sa.IntArray == nil || len(*sa.IntArray) != 3 {
		t.Error("sa slice IntArray is not fully populated")
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
	if sf.FloatGenerate != 87.473328 {
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
		t.Error("Bool should be true got true")
	}
	if sf.BoolConst != true {
		t.Errorf("BoolConst should be true got false")
	}
	if sf.BoolGenerate != false {
		t.Errorf("Bool Generate should be %t but got %t", false, sf.BoolGenerate)
	}
}

func TestStructToDateTime(t *testing.T) {
	Seed(11)

	AddFuncLookup("datetimestatic", Info{
		Description: "A static date time",
		Example:     "2021-11-26 15:22:00",
		Output:      "time.Time",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			// Create new static date time
			return time.Date(2021, 11, 26, 15, 22, 0, 0, time.UTC), nil
		},
	})

	var datetime struct {
		Simple               time.Time
		Tag                  time.Time `fake:"{date}"`
		TagCustom            time.Time `fake:"{datetimestatic}"`
		TagFormat            time.Time `fake:"{number:1900,1950}-12-05" format:"2006-01-02"`
		TagJavaFormat        time.Time `fake:"{number:1900,1950}-12-05" format:"yyyy-MM-dd"`
		Range                time.Time `fake:"{daterange:1970-01-01,2000-12-31,2006-01-02}" format:"yyyy-MM-dd"`
		PointerSimple        *time.Time
		PointerTag           *time.Time `fake:"{date}"`
		PointerTagCustom     *time.Time `fake:"{datetimestatic}"`
		PointerTagFormat     *time.Time `fake:"{number:1900,1950}-12-05" format:"2006-01-02"`
		PointerTagJavaFormat *time.Time `fake:"{number:1900,1950}-12-05" format:"yyyy-MM-dd"`
		PointerRange         *time.Time `fake:"{daterange:1970-01-01,2000-12-31,2006-01-02}" format:"yyyy-MM-dd"`
	}
	err := Struct(&datetime)
	if err != nil {
		t.Fatal(err)
	}

	type testCmp struct {
		name     string
		observed string
		expected string
	}

	testComparisons := []testCmp{
		{
			name:     "Simple",
			observed: datetime.Simple.String(),
			expected: "2012-11-07 04:31:13.726582492 +0000 UTC",
		},
		{
			name:     "Tag",
			observed: datetime.Tag.String(),
			expected: "1967-12-26 21:04:44 +0000 UTC",
		},
		{
			name:     "TagCustom",
			observed: datetime.TagCustom.String(),
			expected: "2021-11-26 15:22:00 +0000 UTC",
		},
		{
			name:     "TagFormat",
			observed: datetime.TagFormat.String(),
			expected: "1947-12-05 00:00:00 +0000 UTC",
		},
		{
			name:     "TagJavaFormat",
			observed: datetime.TagJavaFormat.String(),
			expected: "1946-12-05 00:00:00 +0000 UTC",
		},
		{
			name:     "Range",
			observed: datetime.Range.String(),
			expected: "1980-03-10 00:00:00 +0000 UTC",
		},
		{
			name:     "PointerSimple",
			observed: datetime.PointerSimple.String(),
			expected: "1954-05-23 07:44:55.341122669 +0000 UTC",
		},
		{
			name:     "PointerTag",
			observed: datetime.PointerTag.String(),
			expected: "1900-07-01 22:25:09 +0000 UTC",
		},
		{
			name:     "PointerTagCustom",
			observed: datetime.PointerTagCustom.String(),
			expected: "2021-11-26 15:22:00 +0000 UTC",
		},
		{
			name:     "PointerTagFormat",
			observed: datetime.PointerTagFormat.String(),
			expected: "1910-12-05 00:00:00 +0000 UTC",
		},
		{
			name:     "PointerTagJavaFormat",
			observed: datetime.PointerTagJavaFormat.String(),
			expected: "1932-12-05 00:00:00 +0000 UTC",
		},
		{
			name:     "PointerRange",
			observed: datetime.PointerRange.String(),
			expected: "1995-12-07 00:00:00 +0000 UTC",
		},
	}
	for _, c := range testComparisons {
		if c.expected != c.observed {
			t.Errorf("%s should be %s and instead got %s", c.name, c.expected, c.observed)
		}
	}

	RemoveFuncLookup("datetimestatic")
}

func TestStructSetSubStruct(t *testing.T) {
	Seed(11)

	type Sub struct {
		Str string
		Num int
		Flo float32
	}

	type Foo struct {
		Sub Sub `fake:"{setstruct}"`
	}

	AddFuncLookup("setstruct", Info{
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return Sub{Str: "string", Num: 123, Flo: 123.456}, nil
		},
	})

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

func TestStructMap(t *testing.T) {
	Seed(11)

	type Bar struct {
		Name string
	}

	type MapCustom map[string]Bar

	type Foo struct {
		MapStr       map[string]string
		MapInt       map[int]int
		MapFloat     map[float32]float32
		MapStrPtr    map[string]*string
		MapPtr       *map[string]any
		MapStruct    map[string]Bar
		MapArray     map[string][]Bar
		MapSize      map[string]string `fakesize:"20"`
		MapCustom    MapCustom
		MapCustomMap map[string]MapCustom
	}

	var f Foo
	err := Struct(&f)
	if err != nil {
		t.Fatal(err)
	}

	if len(f.MapStr) == 0 {
		t.Errorf("Map Str %+v", f.MapStr)
	}

	if len(f.MapInt) == 0 {
		t.Errorf("Map Int %+v", f.MapInt)
	}

	if len(f.MapFloat) == 0 {
		t.Errorf("Map Float %+v", f.MapFloat)
	}

	if len(f.MapStrPtr) == 0 {
		t.Errorf("Map Str Ptr %+v", f.MapStrPtr)
	}

	if len(*f.MapPtr) == 0 {
		t.Errorf("Map Ptr %+v", *f.MapPtr)
	}

	if len(f.MapStruct) == 0 {
		t.Errorf("Map Struct %+v", f.MapStruct)
	}

	if len(f.MapArray) == 0 {
		t.Errorf("Map Array %+v", f.MapArray)
	} else {
		for _, v := range f.MapArray {
			if len(v) == 0 {
				t.Errorf("Map Array Values %+v", f.MapArray)
				break
			}
		}
	}

	if len(f.MapSize) != 20 {
		t.Errorf("Map size %+v", f.MapSize)
	}

	if len(f.MapCustom) == 0 {
		t.Errorf("Map Custom %+v", f.MapCustom)
	}

	if len(f.MapCustomMap) == 0 {
		t.Errorf("Map Custom Map %+v", f.MapCustomMap)
	} else {
		for _, v := range f.MapCustomMap {
			if len(v) == 0 {
				t.Errorf("Map Custom Map Values %+v", f.MapCustomMap)
				break
			}
		}
	}
}

func TestStructSliceLoopGeneration(t *testing.T) {
	type S struct {
		A []string
	}

	s := &S{}
	for i := 0; i < 2; i++ {
		if err := Struct(s); err != nil {
			t.Fatal(err)
		}
	}
}

func TestExternalCustomType(t *testing.T) {
	type Foo struct {
		IP net.IP `fake:"{netip}"`
	}

	var f Foo

	Seed(11)
	AddFuncLookup("netip", Info{
		Category:    "custom",
		Description: "Random IPv4 Address of type net.IP",
		Example:     "1.1.1.1",
		Output:      "net.IP",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			data := net.IPv4(byte(f.IntN(256)), byte(f.IntN(256)), byte(f.IntN(256)), byte(f.IntN(256)))
			return data, nil
		},
	})
	defer RemoveFuncLookup("netip")
	Struct(&f)

	// Regex check its an IP
	reg := `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`
	matched, err := regexp.MatchString(reg, f.IP.String())
	if err != nil {
		t.Error(err)
	}

	if !matched {
		t.Errorf("IP address is not valid: %s", f.IP.String())
	}
}

func TestStructArrayWithInvalidCustomFunc(t *testing.T) {
	AddFuncLookup("customType", Info{
		Category:    "custom",
		Description: "Random int array",
		Example:     "[1]",
		Output:      "CustomType",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			data := make([]int, 1)
			data[0] = 42
			return data, nil
		},
	})
	defer RemoveFuncLookup("customType")

	var invalidCustomTag struct {
		InvalidTag []int `fake:"{customType}"`
	}
	err := Struct(&invalidCustomTag)
	if err.Error() != `strconv.ParseInt: parsing "[42]": invalid syntax` {
		t.Error(err)
	}
}

func TestStructMapWithCustomFunction(t *testing.T) {
	AddFuncLookup("custom_map", Info{
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return map[string]string{"abc": "123"}, nil
		},
	})
	defer RemoveFuncLookup("custom_map")

	type StructMap struct {
		MapCustomFun map[string]string `fake:"{custom_map}"`
	}
	var f StructMap

	err := Struct(&f)
	if err != nil {
		t.Fatal(err)
	}

	if v, ok := f.MapCustomFun["abc"]; ok {
		if v != "123" {
			t.Errorf("value didnt equal 123, got %v", v)
		}
	} else {
		t.Errorf("map didnt contain 'abc'.")
	}
}

func TestStructBrokenFakeTag(t *testing.T) {
	type StructMap struct {
		BrokenTag map[string]string `fake:"{randomstring:[hello,world}"`
	}
	var f StructMap

	err := Struct(&f)
	if err == nil {
		t.Error("expected error")
	}

	// error should contain "missing ending ]"
	if !strings.Contains(err.Error(), "missing ending ]") {
		t.Fatal(err)
	}
}

func TestStructRecursion(t *testing.T) {
	type Foo struct {
		Bar        *Foo
		OtherField string
	}

	var f Foo
	err := Struct(&f)
	if err != nil {
		t.Fatal(err)
	}

	if f.Bar == nil {
		t.Error("Bar should not be nil")
	}
	if f.Bar.Bar != nil {
		t.Error("Bar.Bar should be nil")
	}
	if f.OtherField == "" {
		t.Error("Bar.OtherField should not be empty")
	}
}
