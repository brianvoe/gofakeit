package gofakeit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

var (
	testTimeValue = time.Now()
)

type CustomString string

func (c CustomString) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomString("hello test"), nil
}

type CustomBool bool

func (c CustomBool) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomBool(true), nil
}

type CustomInt int

func (c CustomInt) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomInt(-42), nil
}

type CustomInt8 int8

func (c CustomInt8) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomInt8(-42), nil
}

type CustomInt16 int16

func (c CustomInt16) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomInt16(-42), nil
}

type CustomInt32 int32

func (c CustomInt32) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomInt32(-42), nil
}

type CustomInt64 int64

func (c CustomInt64) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomInt64(-42), nil
}

type CustomUint uint

func (c CustomUint) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomUint(42), nil
}

type CustomUint8 uint8

func (c CustomUint8) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomUint8(42), nil
}

type CustomUint16 uint16

func (c CustomUint16) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomUint16(42), nil
}

type CustomUint32 uint32

func (c CustomUint32) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomUint32(42), nil
}

type CustomUint64 uint64

func (c CustomUint64) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomUint64(42), nil
}

type CustomFloat32 float32

func (c CustomFloat32) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomFloat32(42.123), nil
}

type CustomFloat64 float64

func (c CustomFloat64) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomFloat64(42.123), nil
}

type CustomTime time.Time

func (c *CustomTime) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomTime(testTimeValue), nil
}

func (c CustomTime) String() string {
	return time.Time(c).String()
}

type CustomSlice []string

func (c CustomSlice) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomSlice([]string{"hello", "test"}), nil
}

type CustomMap map[string]string

func (c CustomMap) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomMap(map[string]string{"hello": "1", "test": "2"}), nil
}

type CustomArray [2]int

func (c CustomArray) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomArray([2]int{1, 2}), nil
}

type CustomStruct struct {
	Str string
	Int int
}

func (c CustomStruct) Fake(faker *gofakeit.Faker) (any, error) {
	return CustomStruct{
		Str: "hello test",
		Int: 42,
	}, nil
}

type NestedCustom struct {
	Str          CustomString
	PtrStr       *CustomString
	Bool         CustomBool
	Int          CustomInt
	Int8         CustomInt8
	Int16        CustomInt16
	Int32        CustomInt32
	Int64        CustomInt64
	Uint         CustomUint
	Uint8        CustomUint8
	Uint16       CustomUint16
	Uint32       CustomUint32
	Uint64       CustomUint64
	Float32      CustomFloat32
	Float64      CustomFloat64
	Timestamp    CustomTime
	PtrTimestamp *CustomTime
	SliceStr     CustomSlice
	Array        CustomArray
	MapStr       CustomMap
	Struct       CustomStruct
	PtrStruct    *CustomStruct
}

type NestedOverrideCustom struct {
	Str          CustomString  `fake:"{name}"`
	PtrStr       *CustomString `fake:"{name}"`
	Bool         CustomBool    `fake:"false"`
	Int          CustomInt     `fake:"{number:-10,1000}"`
	Int8         CustomInt8    `fake:"{number:-10,1000}"`
	Int16        CustomInt16   `fake:"{number:-10,1000}"`
	Int32        CustomInt32   `fake:"{number:-10,1000}"`
	Int64        CustomInt64   `fake:"{number:-10,1000}"`
	Uint         CustomUint    `fake:"{number:100,1000}"`
	Uint8        CustomUint8   `fake:"{number:100,1000}"`
	Uint16       CustomUint16  `fake:"{number:100,1000}"`
	Uint32       CustomUint32  `fake:"{number:100,1000}"`
	Uint64       CustomUint64  `fake:"{number:100,1000}"`
	Float32      CustomFloat32 `fake:"{number:100,1000}"`
	Float64      CustomFloat64 `fake:"{number:100,1000}"`
	Timestamp    CustomTime    `fake:"{raw_test_date}"`
	PtrTimestamp *CustomTime   `fake:"{raw_test_date}"`
	SliceStr     CustomSlice   `fake:"{word}"`
	Array        CustomArray   `fake:"{number:100,1000}"`
	MapStr       CustomMap     `fakesize:"2"`
}

func TestCustomString(t *testing.T) {
	var d CustomString
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := "hello test"
	if d != CustomString(expected) {
		t.Errorf("expected %q, got %q", expected, d)
	}
}

func TestCustomBool(t *testing.T) {
	var d CustomBool
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := true
	if d != CustomBool(expected) {
		t.Errorf("expected %t, got %t", expected, d)
	}
}

func TestCustomInt(t *testing.T) {
	var d CustomInt
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := -42
	if d != CustomInt(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomInt8(t *testing.T) {
	var d CustomInt8
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := -42
	if d != CustomInt8(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomInt16(t *testing.T) {
	var d CustomInt16
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := -42
	if d != CustomInt16(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomInt32(t *testing.T) {
	var d CustomInt32
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := -42
	if d != CustomInt32(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomInt64(t *testing.T) {
	var d CustomInt64
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := -42
	if d != CustomInt64(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomUint(t *testing.T) {
	var d CustomUint
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42
	if d != CustomUint(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomUint8(t *testing.T) {
	var d CustomUint8
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42
	if d != CustomUint8(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomUint16(t *testing.T) {
	var d CustomUint16
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42
	if d != CustomUint16(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomUint32(t *testing.T) {
	var d CustomUint32
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42
	if d != CustomUint32(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}

func TestCustomUint64(t *testing.T) {
	var d CustomUint64
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42
	if d != CustomUint64(expected) {
		t.Errorf("expected %d, got %d", expected, d)
	}
}
func TestCustomFloat32(t *testing.T) {
	var d CustomFloat32
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42.123
	if d != CustomFloat32(expected) {
		t.Errorf("expected %f, got %f", expected, d)
	}
}

func TestCustomFloat64(t *testing.T) {
	var d CustomFloat64
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := 42.123
	if d != CustomFloat64(expected) {
		t.Errorf("expected %f, got %f", expected, d)
	}
}

func TestCustomTime(t *testing.T) {
	var d CustomTime
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := testTimeValue
	if d != CustomTime(expected) {
		t.Errorf("expected %q, got %q", expected.String(), d.String())
	}
}

func TestCustomTimePtr(t *testing.T) {
	var d *CustomTime
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := testTimeValue
	if d == nil {
		t.Fatal("expected a pointer to a CustomTime, got nil")
	}
	if *d != CustomTime(expected) {
		t.Errorf("expected %q, got %q", expected.String(), d.String())
	}
}

func TestCustomSlice(t *testing.T) {
	var d CustomSlice
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"hello", "test"}
	if len(d) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
	for i, v := range expected {
		if d[i] != v {
			t.Errorf("expected item %d of the slice to be: %v, got %v", i, expected[i], d[i])
		}
	}
}

func TestCustomMap(t *testing.T) {
	var d CustomMap
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{"hello": "1", "test": "2"}
	if len(d) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, d)
	}
	for k, v := range expected {
		if d[k] != v {
			t.Errorf("expected item %v of the slice to be: %v, got %v", k, expected[k], d[k])
		}
	}
}

func TestCustomArray(t *testing.T) {
	var d CustomArray
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expected := [2]int{1, 2}
	for i, v := range expected {
		if d[i] != v {
			t.Errorf("expected item %d of the array to be: %v, got %v", i, expected[i], d[i])
		}
	}
}

func TestCustomStruct(t *testing.T) {
	var d CustomStruct
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	if d.Str != "hello test" {
		t.Errorf("expected %q, got %q", "hello test", d.Str)
	}
	if d.Int != 42 {
		t.Errorf("expected %d, got %d", 42, d.Int)
	}
}

func TestNestedCustom(t *testing.T) {
	var d NestedCustom
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	expectedStr := "hello test"
	if d.Str != CustomString(expectedStr) {
		t.Errorf("Str: expected %q, got %q", expectedStr, d.Str)
	}

	if *d.PtrStr != CustomString(expectedStr) {
		t.Errorf("Str: expected %q, got %q", expectedStr, *d.PtrStr)
	}

	if !d.Bool {
		t.Errorf("Bool: expected true, got false")
	}

	expectedInt := -42
	if d.Int != CustomInt(expectedInt) {
		t.Errorf("Int: expected %d, got %d", expectedInt, d.Int)
	}
	if d.Int8 != CustomInt8(expectedInt) {
		t.Errorf("Int: expected %d, got %d", expectedInt, d.Int8)
	}
	if d.Int16 != CustomInt16(expectedInt) {
		t.Errorf("Int: expected %d, got %d", expectedInt, d.Int16)
	}
	if d.Int32 != CustomInt32(expectedInt) {
		t.Errorf("Int: expected %d, got %d", expectedInt, d.Int32)
	}
	if d.Int64 != CustomInt64(expectedInt) {
		t.Errorf("Int: expected %d, got %d", expectedInt, d.Int64)
	}

	expectedUint := uint(42)
	if d.Uint != CustomUint(expectedUint) {
		t.Errorf("Uint: expected %d, got %d", expectedUint, d.Uint)
	}
	if d.Uint8 != CustomUint8(expectedUint) {
		t.Errorf("Uint: expected %d, got %d", expectedUint, d.Uint8)
	}
	if d.Uint16 != CustomUint16(expectedUint) {
		t.Errorf("Uint: expected %d, got %d", expectedUint, d.Uint16)
	}
	if d.Uint32 != CustomUint32(expectedUint) {
		t.Errorf("Uint: expected %d, got %d", expectedUint, d.Uint32)
	}
	if d.Uint64 != CustomUint64(expectedUint) {
		t.Errorf("Uint: expected %d, got %d", expectedUint, d.Uint64)
	}

	expectedFloat := 42.123
	if d.Float32 != CustomFloat32(expectedFloat) {
		t.Errorf("Float: expected %f, got %f", expectedFloat, d.Float32)
	}
	if d.Float64 != CustomFloat64(expectedFloat) {
		t.Errorf("Float: expected %f, got %f", expectedFloat, d.Float64)
	}

	expectedSlice := []string{"hello", "test"}
	if len(d.SliceStr) != len(expectedSlice) {
		t.Fatalf("expected %v, got %v", expectedSlice, d.SliceStr)
	}
	for i, v := range expectedSlice {
		if d.SliceStr[i] != v {
			t.Errorf("expected item %d of the slice to be: %v, got %v", i, expectedSlice[i], d.SliceStr[i])
		}
	}

	expectedArray := [2]int{1, 2}
	for i, v := range expectedArray {
		if d.Array[i] != v {
			t.Errorf("expected item %d of the slice to be: %v, got %v", i, expectedArray[i], d.Array[i])
		}
	}

	expectedMap := map[string]string{"hello": "1", "test": "2"}
	if len(d.MapStr) != len(expectedMap) {
		t.Fatalf("expected %v, got %v", expectedMap, d)
	}
	for k, v := range expectedMap {
		if d.MapStr[k] != v {
			t.Errorf("expected item %v of the map to be: %v, got %v", k, expectedMap[k], d.MapStr[k])
		}
	}

	if d.Struct.Str != "hello test" {
		t.Errorf("expected %q, got %q", "hello test", d.Struct.Str)
	}
	if d.Struct.Int != 42 {
		t.Errorf("expected %d, got %d", 42, d.Struct.Int)
	}

	if d.PtrStruct == nil {
		t.Fatal("expected PtrStruct to not be nil")
	}

	if d.PtrStruct.Str != "hello test" {
		t.Errorf("expected %q, got %q", "hello test", d.PtrStruct.Str)
	}
	if d.PtrStruct.Int != 42 {
		t.Errorf("expected %d, got %d", 42, d.PtrStruct.Int)
	}

	expectedTimestamp := testTimeValue
	if d.Timestamp != CustomTime(expectedTimestamp) {
		t.Errorf("expected %q, got %q", expectedTimestamp.String(), d.Timestamp.String())
	}

	if d.PtrTimestamp == nil {
		t.Fatal("expected a pointer to a CustomTime, got nil")
	}
	if *d.PtrTimestamp != CustomTime(expectedTimestamp) {
		t.Errorf("expected %q, got %q", expectedTimestamp.String(), d.PtrTimestamp.String())
	}
}

func TestNestedOverrideCustom(t *testing.T) {
	gofakeit.AddFuncLookup("raw_test_date", gofakeit.Info{
		Display:     "Date",
		Category:    "time",
		Description: "Random date",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "time.Time",
		Params: []gofakeit.Param{
			{
				Field:       "format",
				Display:     "Format",
				Type:        "time.Time",
				Description: "Raw date time.Time object",
			},
		},
		Generate: func(f *gofakeit.Faker, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return gofakeit.Date(), nil
		},
	})

	defer gofakeit.RemoveFuncLookup("raw_test_date")

	var d NestedOverrideCustom
	err := gofakeit.Struct(&d)
	if err != nil {
		t.Fatal(err)
	}

	nonOverrideStr := "hello test"
	if d.Str == CustomString(nonOverrideStr) {
		t.Errorf("Str: expected a random string but got the non-overriden value")
	}

	if *d.PtrStr == CustomString(nonOverrideStr) {
		t.Errorf("PtrStr: expected a random string but got the non-overriden value")
	}

	if d.Bool {
		t.Errorf("Bool: expected false, got true")
	}

	nonOverrideInt := -42
	if d.Int == CustomInt(nonOverrideInt) {
		t.Errorf("Int: expected a random integer but got the non-overriden value")
	}
	if d.Int8 == CustomInt8(nonOverrideInt) {
		t.Errorf("Int: expected a random integer but got the non-overriden value")
	}
	if d.Int16 == CustomInt16(nonOverrideInt) {
		t.Errorf("Int: expected a random integer but got the non-overriden value")
	}
	if d.Int32 == CustomInt32(nonOverrideInt) {
		t.Errorf("Int: expected a random integer but got the non-overriden value")
	}
	if d.Int64 == CustomInt64(nonOverrideInt) {
		t.Errorf("Int: expected a random integer but got the non-overriden value")
	}

	nonOverrideUint := uint(42)
	if d.Uint == CustomUint(nonOverrideUint) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}
	if d.Uint8 == CustomUint8(nonOverrideUint) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}
	if d.Uint16 == CustomUint16(nonOverrideUint) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}
	if d.Uint32 == CustomUint32(nonOverrideUint) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}
	if d.Uint64 == CustomUint64(nonOverrideUint) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}

	nonOverrideFloat := 42.123
	if d.Float32 == CustomFloat32(nonOverrideFloat) {
		t.Errorf("Float: expected a random unsigned integer but got the non-overriden value")
	}
	if d.Float64 == CustomFloat64(nonOverrideFloat) {
		t.Errorf("Uint: expected a random unsigned integer but got the non-overriden value")
	}

	nonOverrideSlice := []string{"hello", "test"}
	if len(d.SliceStr) == len(nonOverrideSlice) {
		t.Logf("Slice: Got the same length as the non-overriden slice: %v vs %v", nonOverrideSlice, d.SliceStr)
		for i, v := range nonOverrideSlice {
			if d.SliceStr[i] == v {
				t.Errorf("Slice: Got non-overriden item %d in the slice", i)
			}
		}
	}

	nonOverrideArray := [2]int{1, 2}
	for i, v := range nonOverrideArray {
		if d.Array[i] == v {
			t.Errorf("Array: Got non-overriden item %d in the array", i)
		}
	}

	nonOverrideMap := map[string]string{"hello": "1", "test": "2"}
	if len(d.MapStr) == len(nonOverrideMap) {
		t.Logf("Map: Got the same length as the non-overriden map: %v vs %v", nonOverrideMap, d.MapStr)

		for k, v := range nonOverrideMap {
			if d.MapStr[k] == v {
				t.Errorf("Map: Got non-overriden item %v in the slice", k)
			}
		}
	}
}

func TestSliceCustom(t *testing.T) {
	var B []CustomString
	gofakeit.Slice(&B)

	if len(B) == 0 {
		t.Errorf("expected slice to not be empty")
	}

	expected := CustomString("hello test")
	for _, v := range B {
		if v != expected {
			t.Errorf("expected all items to be %q, got %q", expected, v)
		}
	}
}

func TestSliceNestedCustom(t *testing.T) {
	var B []NestedCustom
	gofakeit.Slice(&B)

	if len(B) == 0 {
		t.Errorf("expected slice to not be empty")
	}

	expected := CustomString("hello test")
	for _, v := range B {
		if v.Str != expected {
			t.Fatalf("expected all items to be %q, got %q", expected, v.Str)
		}
	}
}

func ExampleCustomInt() {
	f1 := gofakeit.New(10)
	f2 := gofakeit.New(100)

	var A1 CustomInt
	var A2 CustomInt
	// CustomInt always returns -42 independently of the seed
	f1.Struct(&A1)
	f2.Struct(&A2)

	fmt.Println(A1)
	fmt.Println(A2)

	// Output:
	// -42
	// -42
}

type EvenInt int

func (e EvenInt) Fake(faker *gofakeit.Faker) (any, error) {
	return EvenInt(faker.Int8() * 2), nil
}

func ExampleEvenInt() {
	f1 := gofakeit.New(10)
	f2 := gofakeit.New(100)

	var E1 EvenInt
	var E2 EvenInt
	// EventInt always returns an even number
	f1.Struct(&E1)
	f2.Struct(&E2)

	fmt.Println(E1)
	fmt.Println(E2)

	// Output: -2
	// 122
}
