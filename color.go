package gofakeit

import "strings"

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
	color := ""
	for i := 1; i <= 6; i++ {
		color += ShuffleStrings([]string{"?", "#"})[0]
	}

	// Replace # with number
	color = replaceWithNumbers(color)

	// Replace ? with letter
	for strings.Count(color, "?") > 0 {
		color = strings.Replace(color, "?", ShuffleStrings([]string{"a", "b", "c", "d", "e", "f"})[0], 1)
	}

	return "#" + color
}

// RGBColor will generate a random int slice color
func RGBColor() []int {
	return []int{randIntRange(0, 255), randIntRange(0, 255), randIntRange(0, 255)}
}
