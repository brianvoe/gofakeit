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

// The number of shapes (10 to 20) is drawn once, it used to be drawn again on each pass of the loop,
// so 18 or more shapes were rare
func TestSvgShapeCount(t *testing.T) {
	f := New(11)

	var counts []int
	for range 20000 {
		svg := f.Svg(&SVGOptions{Type: "rect", Width: 100, Height: 100, Colors: []string{"#000000"}})

		// The background is a rect too
		counts = append(counts, strings.Count(svg, "<rect")-1)
	}
	checkUniformCounts(t, "svg shapes", 10, 20, counts)
}

func ExampleSvg() {
	Seed(11)
	fmt.Println(Svg(nil))

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 459 443" width="459" height="443"><rect x="0" y="0" width="100%" height="100%" fill="#64908a" /><circle cx="334" cy="238" r="437" fill="#cc2a41" /><circle cx="415" cy="32" r="340" fill="#cc2a41" /><circle cx="428" cy="401" r="151" fill="#64908a" /><circle cx="191" cy="320" r="138" fill="#e8caa4" /><circle cx="425" cy="151" r="3" fill="#64908a" /><circle cx="5" cy="410" r="192" fill="#351330" /><circle cx="16" cy="91" r="296" fill="#cc2a41" /><circle cx="437" cy="197" r="34" fill="#64908a" /><circle cx="439" cy="46" r="269" fill="#64908a" /><circle cx="321" cy="268" r="422" fill="#351330" /><circle cx="293" cy="106" r="20" fill="#64908a" /><circle cx="94" cy="265" r="339" fill="#351330" /></svg>
}

func ExampleFaker_Svg() {
	f := New(11)
	fmt.Println(f.Svg(nil))

	// Output: <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 459 443" width="459" height="443"><rect x="0" y="0" width="100%" height="100%" fill="#64908a" /><circle cx="334" cy="238" r="437" fill="#cc2a41" /><circle cx="415" cy="32" r="340" fill="#cc2a41" /><circle cx="428" cy="401" r="151" fill="#64908a" /><circle cx="191" cy="320" r="138" fill="#e8caa4" /><circle cx="425" cy="151" r="3" fill="#64908a" /><circle cx="5" cy="410" r="192" fill="#351330" /><circle cx="16" cy="91" r="296" fill="#cc2a41" /><circle cx="437" cy="197" r="34" fill="#64908a" /><circle cx="439" cy="46" r="269" fill="#64908a" /><circle cx="321" cy="268" r="422" fill="#351330" /><circle cx="293" cy="106" r="20" fill="#64908a" /><circle cx="94" cy="265" r="339" fill="#351330" /></svg>
}

func BenchmarkSvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Svg(nil)
	}
}
