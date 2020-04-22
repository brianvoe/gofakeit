package gofakeit

import (
	"math/rand"
)

// Color will generate a random color string
func Color() string {
	return getRandValue([]string{"color", "full"})
}

// SafeColor will generate a random safe color string
func SafeColor() string {
	return getRandValue([]string{"color", "safe"})
}

// HexColor will generate a random hexadecimal color string
func HexColor() string {
	color := make([]byte, 6)
	hashQuestion := []byte("?#")
	for i := 0; i < 6; i++ {
		color[i] = hashQuestion[rand.Intn(2)]
	}

	return "#" + replaceWithHexLetters(replaceWithNumbers(string(color)))
}

// RGBColor will generate a random int slice color
func RGBColor() []int {
	return []int{randIntRange(0, 255), randIntRange(0, 255), randIntRange(0, 255)}
}

func addColorLookup() {
	AddFuncLookup("color", Info{
		Display:     "Color",
		Category:    "color",
		Description: "Random color",
		Example:     "MediumOrchid",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Color(), nil
		},
	})

	AddFuncLookup("safecolor", Info{
		Display:     "Safe Color",
		Category:    "color",
		Description: "Random safe color",
		Example:     "black",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return SafeColor(), nil
		},
	})

	AddFuncLookup("hexcolor", Info{
		Display:     "Hex Color",
		Category:    "color",
		Description: "Random hex color",
		Example:     "#a99fb4",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HexColor(), nil
		},
	})

	AddFuncLookup("rgbcolor", Info{
		Display:     "RGB Color",
		Category:    "color",
		Description: "Random rgb color",
		Example:     "[152 23 53]",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return RGBColor(), nil
		},
	})
}
