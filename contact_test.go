package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePhone() {
	fmt.Println(Phone())
	// Output: 1-138-021-2711
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExampleEmail() {
	fmt.Println(Email())
	// Output: antonettalangosh@abshire.info
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}
