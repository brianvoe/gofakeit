package gofakeit

import (
	"fmt"
	"strings"
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

	// Output: #f6ddb4
}

func ExampleFaker_HexColor() {
	f := New(11)
	fmt.Println(f.HexColor())

	// Output: #f6ddb4
}

// Every color can come out: each hex digit of each position shows up about equally often.
// It never returned a 0 as the first digit, and a-f twice as often as it should
func TestHexColor(t *testing.T) {
	f := New(11)
	const samples = 160000
	const digits = "0123456789abcdef"

	var first, all [16]int
	for i := 0; i < samples; i++ {
		color := f.HexColor()
		if len(color) != 7 || color[0] != '#' {
			t.Fatalf("expected # and six hex digits, got %q", color)
		}

		for pos := 1; pos < 7; pos++ {
			d := strings.IndexByte(digits, color[pos])
			if d < 0 {
				t.Fatalf("%q has %q, which is not a lower case hex digit", color, color[pos])
			}

			all[d]++
			if pos == 1 {
				first[d]++
			}
		}
	}

	// 10,000 and 60,000 expected, the standard deviation is 97 and 237
	for d := range 16 {
		if first[d] < 9500 || first[d] > 10500 {
			t.Errorf("the first digit is %c %d times of %d, expected about 10000", digits[d], first[d], samples)
		}

		if all[d] < 58200 || all[d] > 61800 {
			t.Errorf("the digit %c shows up %d times of %d, expected about 60000", digits[d], all[d], 6*samples)
		}
	}
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

func ExampleHSLColor() {
	Seed(11)
	fmt.Println(HSLColor())
	// Output: [323 86 20]
}

func ExampleFaker_HSLColor() {
	f := New(11)
	fmt.Println(f.HSLColor())
	// Output: [323 86 20]
}

func BenchmarkHSLColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HSLColor()
	}
}

func BenchmarkRGBColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RGBColor()
	}
}
