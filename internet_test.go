package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleDomainSuffix() {
	fmt.Println(DomainSuffix())
	// Output: net
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix()
	}
}
