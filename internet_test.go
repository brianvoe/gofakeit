package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())
	// Output: com
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix()
	}
}
