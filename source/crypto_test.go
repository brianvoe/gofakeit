package source

import "testing"

func TestCrypto(t *testing.T) {
	crypto := NewCrypto()

	// test for duplicates
	m := make(map[uint64]bool)
	for range 100000 {
		v := crypto.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

func BenchmarkCrypto(b *testing.B) {
	crypto := NewCrypto()

	for i := 0; i < b.N; i++ {
		crypto.Uint64()
	}
}
