package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleColor() {
	Seed(11)
	fmt.Println(Color())
	// Output: MediumOrchid
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color()
	}
}

func ExampleSafeColor() {
	Seed(11)
	fmt.Println(SafeColor())
	// Output: black
}

func BenchmarkSafeColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafeColor()
	}
}

func ExampleHexColor() {
	Seed(11)
	fmt.Println(HexColor())
	// Output: #i99jb4
}

func BenchmarkHexColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HexColor()
	}
}

func ExampleRGBColor() {
	Seed(11)
	fmt.Println(RGBColor())
	// Output: [152 23 53]
}

func BenchmarkRGBColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGBColor()
	}
}
