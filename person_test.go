package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSSN() {
	Seed(11)
	fmt.Println(SSN())
	// Output: 328-727-1570
}

func BenchmarkSSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SSN()
	}
}
