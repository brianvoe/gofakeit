package gofakeit

import (
	"github.com/brianvoe/gofakeit/v7/data"
)

// Color will generate a random color string
func Color() string { return color(GlobalFaker) }

// Color will generate a random color string
func (f *Faker) Color() string { return color(f) }

func color(f *Faker) string { return getRandValue(f, []string{"color", "full"}) }

// NiceColor will generate a random safe color string
func NiceColors() []string { return niceColors(GlobalFaker) }

// NiceColor will generate a random safe color string
func (f *Faker) NiceColors() []string { return niceColors(f) }

func niceColors(f *Faker) []string {
	return data.ColorsNice[randIntRange(f, 0, len(data.ColorsNice)-1)]
}

// SafeColor will generate a random safe color string
func SafeColor() string { return safeColor(GlobalFaker) }

// SafeColor will generate a random safe color string
func (f *Faker) SafeColor() string { return safeColor(f) }

func safeColor(f *Faker) string { return getRandValue(f, []string{"color", "safe"}) }

// HexColor will generate a random hexadecimal color string
func HexColor() string { return hexColor(GlobalFaker) }

// HexColor will generate a random hexadecimal color string
func (f *Faker) HexColor() string { return hexColor(f) }

func hexColor(f *Faker) string {
	color := make([]byte, 6)
	hashQuestion := []byte("?#")
	for i := 0; i < 6; i++ {
		color[i] = hashQuestion[f.IntN(2)]
	}

	return "#" + replaceWithHexLetters(f, replaceWithNumbers(f, string(color)))
}

// RGBColor will generate a random int slice color
func RGBColor() []int { return rgbColor(GlobalFaker) }

// RGBColor will generate a random int slice color
func (f *Faker) RGBColor() []int { return rgbColor(f) }

func rgbColor(f *Faker) []int {
	return []int{randIntRange(f, 0, 255), randIntRange(f, 0, 255), randIntRange(f, 0, 255)}
}

func addColorLookup() {
	AddFuncLookup("color", Info{
		Display:     "Color",
		Category:    "color",
		Description: "Hue seen by the eye, returns the name of the color like red or blue",
		Example:     "MediumOrchid",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return color(f), nil
		},
	})

	AddFuncLookup("nicecolors", Info{
		Display:     "Nice Colors",
		Category:    "color",
		Description: "Attractive and appealing combinations of colors, returns an list of color hex codes",
		Example:     `["#cfffdd","#b4dec1","#5c5863","#a85163","#ff1f4c"]`,
		Output:      "[]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return color(f), nil
		},
	})

	AddFuncLookup("safecolor", Info{
		Display:     "Safe Color",
		Category:    "color",
		Description: "Colors displayed consistently on different web browsers and devices",
		Example:     "black",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return safeColor(f), nil
		},
	})

	AddFuncLookup("hexcolor", Info{
		Display:     "Hex Color",
		Category:    "color",
		Description: "Six-digit code representing a color in the color model",
		Example:     "#a99fb4",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexColor(f), nil
		},
	})

	AddFuncLookup("rgbcolor", Info{
		Display:     "RGB Color",
		Category:    "color",
		Description: "Color defined by red, green, and blue light values",
		Example:     "[85, 224, 195]",
		Output:      "[]int",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return rgbColor(f), nil
		},
	})
}
