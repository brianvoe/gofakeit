package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLetter() {
	Seed(11)
	fmt.Println(Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Letter()
	}
}
