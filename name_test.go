package gofakeit

import (
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}
