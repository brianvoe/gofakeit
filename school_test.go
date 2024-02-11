package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSchool() {
	Seed(11)
	fmt.Println(School())
	// Output:
	// Harborview State Academy
}

func ExampleFaker_School() {
	f := New(11)
	fmt.Println(f.School())
	// Output:
	// Harborview State Academy
}

func BenchmarkSchool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		School()
	}
}
