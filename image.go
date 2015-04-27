package gofakeit

import "strconv"

// Generate Random Image Based Upon Height And Width. Images Provided by LoremPixel (http://lorempixel.com/)
func ImageUrl(width int, height int) string {
	return "http://lorempixel.com/" + strconv.Itoa(width) + "/" + strconv.Itoa(height)
}
