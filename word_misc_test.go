package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleInterjection() {
	Seed(11)
	fmt.Println(Interjection())

	// Output: alas
}

func ExampleFaker_Interjection() {
	f := New(11)
	fmt.Println(f.Interjection())

	// Output: alas
}

func BenchmarkInterjection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Interjection()
	}
}
