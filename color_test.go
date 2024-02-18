package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleColor() {
	Seed(11)
	fmt.Println(Color())

	// Output: SlateGray
}

func ExampleFaker_Color() {
	f := New(11)
	fmt.Println(f.Color())

	// Output: SlateGray
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color()
	}
}

func ExampleNiceColors() {
	Seed(11)
	fmt.Println(NiceColors())

	// Output: [#fffbb7 #a6f6af #66b6ab #5b7c8d #4f2958]
}

func ExampleFaker_NiceColors() {
	f := New(11)
	fmt.Println(f.NiceColors())

	// Output: [#fffbb7 #a6f6af #66b6ab #5b7c8d #4f2958]
}

func BenchmarkNiceColors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NiceColors()
	}
}

func ExampleSafeColor() {
	Seed(11)
	fmt.Println(SafeColor())

	// Output: aqua
}

func ExampleFaker_SafeColor() {
	f := New(11)
	fmt.Println(f.SafeColor())

	// Output: aqua
}

func BenchmarkSafeColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeColor()
	}
}

func ExampleHexColor() {
	Seed(11)
	fmt.Println(HexColor())

	// Output: #ef759a
}

func ExampleFaker_HexColor() {
	f := New(11)
	fmt.Println(f.HexColor())

	// Output: #ef759a
}

func BenchmarkHexColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexColor()
	}
}

func ExampleRGBColor() {
	Seed(11)
	fmt.Println(RGBColor())

	// Output: [180 18 181]
}

func ExampleFaker_RGBColor() {
	f := New(11)
	fmt.Println(f.RGBColor())

	// Output: [180 18 181]
}

func BenchmarkRGBColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGBColor()
	}
}
