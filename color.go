package gofakeit

import "strings"

// Generate Random Color
func Color() string {
	return getRandValue([]string{"color", "full"})
}

// Generate Random Safe Color
func SafeColor() string {
	return getRandValue([]string{"color", "safe"})
}

// Generate Random Hex Color
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

// Generate Random RGB Color
func RGBColor() []int {
	return []int{randIntRange(0, 255), randIntRange(0, 255), randIntRange(0, 255)}
}
