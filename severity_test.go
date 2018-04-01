package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSeverity() {
	Seed(11)
	fmt.Println(Severity())
	// Output: EMERGE
}

func BenchmarkSeverity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Severity()
	}
}
