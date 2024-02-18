package source

import "testing"

func TestCrypto(t *testing.T) {
	crypto := NewCrypto()

	// test for duplicates
	m := make(map[uint64]bool)
	for i := 0; i < 100000; i++ {
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
