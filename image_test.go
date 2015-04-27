package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleImageUrl() {
	Seed(11)
	fmt.Println(ImageUrl(640, 480))
	// Output: http://lorempixel.com/640/480
}

func BenchmarkImageUrl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ImageUrl(640, 480)
	}
}
