package source

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkChaCha8(b *testing.B) {
	chacha := rand.NewChaCha8([32]byte{0, 1, 2, 3, 4, 5})

	for i := 0; i < b.N; i++ {
		chacha.Uint64()
	}
}

func BenchmarkPCG(b *testing.B) {
	pcg := rand.NewPCG(0, 0)

	for i := 0; i < b.N; i++ {
		pcg.Uint64()
	}
}
