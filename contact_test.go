package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePhone() {
	Seed(11)
	fmt.Println(Phone())
	// Output: 287-271-5702
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExampleEmail() {
	Seed(11)
	fmt.Println(Email())
	// Output: markusmoen@pagac.net
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}
