package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBool() {
	Seed(11)
	fmt.Println(Bool())
	// Output: false
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool()
	}
}
