package gofakeit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7/data"
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
	for i := 0; i < b.N; i++ {
		InputName()
	}
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

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 459 443" width="459" height="443"><rect x="0" y="0" width="100%" height="100%" fill="#64908a" /><circle cx="334" cy="238" r="437" fill="#cc2a41" /><circle cx="34" cy="328" r="428" fill="#cc2a41" /><circle cx="151" cy="190" r="191" fill="#e8caa4" /><circle cx="343" cy="410" r="156" fill="#351330" /><circle cx="5" cy="410" r="192" fill="#351330" /><circle cx="94" cy="286" r="384" fill="#cc2a41" /><circle cx="34" cy="239" r="439" fill="#351330" /><circle cx="199" cy="310" r="278" fill="#cc2a41" /><circle cx="293" cy="106" r="20" fill="#64908a" /><circle cx="275" cy="327" r="11" fill="#cc2a41" /><circle cx="406" cy="338" r="239" fill="#e8caa4" /><circle cx="41" cy="71" r="344" fill="#64908a" /><circle cx="411" cy="116" r="208" fill="#64908a" /><circle cx="410" cy="88" r="303" fill="#351330" /><circle cx="241" cy="69" r="285" fill="#cc2a41" /></svg>
}

func ExampleFaker_Svg() {
	f := New(11)
	fmt.Println(f.Svg(nil))

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 459 443" width="459" height="443"><rect x="0" y="0" width="100%" height="100%" fill="#64908a" /><circle cx="334" cy="238" r="437" fill="#cc2a41" /><circle cx="34" cy="328" r="428" fill="#cc2a41" /><circle cx="151" cy="190" r="191" fill="#e8caa4" /><circle cx="343" cy="410" r="156" fill="#351330" /><circle cx="5" cy="410" r="192" fill="#351330" /><circle cx="94" cy="286" r="384" fill="#cc2a41" /><circle cx="34" cy="239" r="439" fill="#351330" /><circle cx="199" cy="310" r="278" fill="#cc2a41" /><circle cx="293" cy="106" r="20" fill="#64908a" /><circle cx="275" cy="327" r="11" fill="#cc2a41" /><circle cx="406" cy="338" r="239" fill="#e8caa4" /><circle cx="41" cy="71" r="344" fill="#64908a" /><circle cx="411" cy="116" r="208" fill="#64908a" /><circle cx="410" cy="88" r="303" fill="#351330" /><circle cx="241" cy="69" r="285" fill="#cc2a41" /></svg>
}

func BenchmarkSvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Svg(nil)
	}
}
