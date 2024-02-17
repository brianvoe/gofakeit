package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNumber() {
	Seed(11)
	fmt.Println(Number(50, 23456))

	// Output: 21019
}

func ExampleFaker_Number() {
	f := New(11)
	fmt.Println(f.Number(50, 23456))

	// Output: 21019
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number(10, 999999)
	}
}

func ExampleUint8() {
	Seed(11)
	fmt.Println(Uint8())

	// Output: 180
}

func ExampleFaker_Uint8() {
	f := New(11)
	fmt.Println(f.Uint8())

	// Output: 180
}

func BenchmarkUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint8()
	}
}

func ExampleUint16() {
	Seed(11)
	fmt.Println(Uint16())

	// Output: 56756
}

func ExampleFaker_Uint16() {
	f := New(11)
	fmt.Println(f.Uint16())

	// Output: 56756
}

func BenchmarkUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint16()
	}
}

func ExampleUint32() {
	Seed(11)
	fmt.Println(Uint32())

	// Output: 3847792206
}

func ExampleFaker_Uint32() {
	f := New(11)
	fmt.Println(f.Uint32())

	// Output: 3847792206
}

func BenchmarkUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint32()
	}
}

func ExampleUint64() {
	Seed(11)
	fmt.Println(Uint64())

	// Output: 16526141687177076148
}

func ExampleFaker_Uint64() {
	f := New(11)
	fmt.Println(f.Uint64())

	// Output: 16526141687177076148
}

func BenchmarkUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Uint64()
	}
}

func ExampleUintRange() {
	Seed(11)
	fmt.Println(UintRange(1, 10))

	// Output: 9
}

func ExampleFaker_UintRange() {
	f := New(11)
	fmt.Println(f.UintRange(1, 10))

	// Output: 9
}

func BenchmarkUintRange(b *testing.B) {
	min := uint(1)
	max := uint(10)

	for i := 0; i < b.N; i++ {
		UintRange(min, max)
	}
}

func ExampleInt8() {
	Seed(11)
	fmt.Println(Int8())

	// Output: 52
}

func ExampleFaker_Int8() {
	f := New(11)
	fmt.Println(f.Int8())

	// Output: 52
}

func BenchmarkInt8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int8()
	}
}

func ExampleInt16() {
	Seed(11)
	fmt.Println(Int16())

	// Output: 23988
}

func ExampleFaker_Int16() {
	f := New(11)
	fmt.Println(f.Int16())

	// Output: 23988
}

func BenchmarkInt16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int16()
	}
}

func ExampleInt32() {
	Seed(11)
	fmt.Println(Int32())

	// Output: 1923896103
}

func ExampleFaker_Int32() {
	f := New(11)
	fmt.Println(f.Int32())

	// Output: 1923896103
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32()
	}
}

func ExampleInt64() {
	Seed(11)
	fmt.Println(Int64())

	// Output: 7302769650322300340
}

func ExampleFaker_Int64() {
	f := New(11)
	fmt.Println(f.Int64())

	// Output: 7302769650322300340
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64()
	}
}

func ExampleIntRange() {
	Seed(11)
	fmt.Println(IntRange(1, 10))

	// Output: 9
}

func ExampleFaker_IntRange() {
	f := New(11)
	fmt.Println(f.IntRange(1, 10))

	// Output: 9
}

func BenchmarkIntRange(b *testing.B) {
	min := int(1)
	max := int(10)

	for i := 0; i < b.N; i++ {
		IntRange(min, max)
	}
}

func ExampleFloat32() {
	Seed(11)
	fmt.Println(Float32())

	// Output: 0.3462876
}

func ExampleFaker_Float32() {
	f := New(11)
	fmt.Println(f.Float32())

	// Output: 0.3462876
}

func BenchmarkFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32()
	}
}

func ExampleFloat32Range() {
	Seed(11)
	fmt.Println(Float32Range(0, 9999999))

	// Output: 3.4628758e+06
}

func ExampleFaker_Float32Range() {
	f := New(11)
	fmt.Println(f.Float32Range(0, 9999999))

	// Output: 3.4628758e+06
}

func BenchmarkFloat32Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float32Range(10, 999999)
	}
}

func TestFloat32RangeSame(t *testing.T) {
	if float32Range(GlobalFaker, 5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func ExampleFloat64() {
	Seed(11)
	fmt.Println(Float64())

	// Output: 0.7703009321621068
}

func ExampleFaker_Float64() {
	f := New(11)
	fmt.Println(f.Float64())

	// Output: 0.7703009321621068
}

func BenchmarkFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64()
	}
}

func ExampleFloat64Range() {
	Seed(11)
	fmt.Println(Float64Range(0, 9999999))

	// Output: 7.703008551320136e+06
}

func ExampleFaker_Float64Range() {
	f := New(11)
	fmt.Println(f.Float64Range(0, 9999999))

	// Output: 7.703008551320136e+06
}

func BenchmarkFloat64Range(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Float64Range(0, 999999)
	}
}

func TestRandFloat64RangeSame(t *testing.T) {
	if float64Range(GlobalFaker, 5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func ExampleShuffleInts() {
	Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	ShuffleInts(ints)
	fmt.Println(ints)

	// Output: [941 777 8413 74125 854 89416 841657 52]
}

func ExampleFaker_ShuffleInts() {
	f := New(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	f.ShuffleInts(ints)
	fmt.Println(ints)

	// Output: [941 777 8413 74125 854 89416 841657 52]
}

func BenchmarkShuffleInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}

func ExampleRandomInt() {
	Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(RandomInt(ints))

	// Output: 8413
}

func ExampleFaker_RandomInt() {
	f := New(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(f.RandomInt(ints))

	// Output: 8413
}

func TestRandomInt(t *testing.T) {
	ints := []int{}
	RandomInt(ints)

	ints = []int{1}
	RandomInt(ints)
}

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInt([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}

func ExampleRandomUint() {
	Seed(11)

	ints := []uint{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(RandomUint(ints))

	// Output: 8413
}

func ExampleFaker_RandomUint() {
	f := New(11)

	ints := []uint{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(f.RandomUint(ints))

	// Output: 8413
}

func TestRandomUint(t *testing.T) {
	ints := []uint{}
	RandomUint(ints)

	ints = []uint{1}
	RandomUint(ints)
}

func BenchmarkRandomUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomUint([]uint{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}

func ExampleHexUint() {
	Seed(11)
	fmt.Println(HexUint(16))

	// Output: 0x425b
}

func ExampleFaker_HexUint() {
	f := New(11)
	fmt.Println(f.HexUint(16))

	// Output: 0x425b
}

func BenchmarkHexUint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexUint(16)
	}
}
