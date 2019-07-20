package gofakeit

import "math/rand"

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
