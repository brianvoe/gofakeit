package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleImageURL() {
	Seed(11)
	fmt.Println(ImageURL(640, 480))
	// Output: http://lorempixel.com/640/480
}

func BenchmarkImageURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImageURL(640, 480)
	}
}
