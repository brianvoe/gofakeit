package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSimpleStatusCode() {
	Seed(11)
	fmt.Println(SimpleStatusCode())
	// Output: 200
}

func BenchmarkSimpleStatusCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleStatusCode()
	}
}

func ExampleStatusCode() {
	Seed(11)
	fmt.Println(StatusCode())
	// Output: 404
}

func BenchmarkStatusCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StatusCode()
	}
}
