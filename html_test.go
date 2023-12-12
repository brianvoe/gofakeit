package gofakeit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6/data"
)

func ExampleInputName() {
	Seed(11)
	fmt.Println(InputName())

	// Output: message
}

func ExampleFaker_InputName() {
	f := New(11)
	fmt.Println(f.InputName())

	// Output: message
}

func BenchmarkInputName(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InputName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.InputName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.InputName()
		}
	})
}

func TestSvg(t *testing.T) {
	Seed(11)

	// Loop through SvgTypes and set that type for each test
	var SvgTypes = data.GetSubData("html", "svg")
	for _, svgType := range SvgTypes {
		// Run the test
		t.Run(svgType, func(t *testing.T) {

			// Get the image
			img := Svg(&SVGOptions{
				Type: svgType,
			})

			// Check the image
			if img == "" {
				t.Error("Svg returned an empty string")
			}

			// Check the image
			if !strings.Contains(img, svgType) {
				t.Errorf("Svg returned an image of type %s, but should have been %s", svgType, svgType)
			}
		})
	}
}

func ExampleSvg() {
	Seed(11)
	fmt.Println(Svg(nil))

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 420 496" width="420" height="496"><rect x="0" y="0" width="100%" height="100%" fill="#4f2958" /><polygon points="382,87 418,212 415,110" fill="#fffbb7" /><polygon points="283,270 225,264 411,352" fill="#5b7c8d" /><polygon points="147,283 388,2 117,263" fill="#4f2958" /><polygon points="419,123 71,282 56,55" fill="#fffbb7" /><polygon points="54,451 377,89 52,351" fill="#66b6ab" /><polygon points="395,169 397,256 110,208" fill="#5b7c8d" /><polygon points="222,52 96,147 107,296" fill="#66b6ab" /><polygon points="126,422 57,93 43,221" fill="#fffbb7" /><polygon points="126,125 61,130 348,57" fill="#5b7c8d" /><polygon points="26,235 97,182 58,37" fill="#a6f6af" /><polygon points="190,495 407,44 53,79" fill="#66b6ab" /></svg>
}

func ExampleFaker_Svg() {
	f := New(11)
	fmt.Println(f.Svg(nil))

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 420 496" width="420" height="496"><rect x="0" y="0" width="100%" height="100%" fill="#4f2958" /><polygon points="382,87 418,212 415,110" fill="#fffbb7" /><polygon points="283,270 225,264 411,352" fill="#5b7c8d" /><polygon points="147,283 388,2 117,263" fill="#4f2958" /><polygon points="419,123 71,282 56,55" fill="#fffbb7" /><polygon points="54,451 377,89 52,351" fill="#66b6ab" /><polygon points="395,169 397,256 110,208" fill="#5b7c8d" /><polygon points="222,52 96,147 107,296" fill="#66b6ab" /><polygon points="126,422 57,93 43,221" fill="#fffbb7" /><polygon points="126,125 61,130 348,57" fill="#5b7c8d" /><polygon points="26,235 97,182 58,37" fill="#a6f6af" /><polygon points="190,495 407,44 53,79" fill="#66b6ab" /></svg>
}

func BenchmarkSvg(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Svg(nil)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Svg(nil)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Svg(nil)
		}
	})
}
