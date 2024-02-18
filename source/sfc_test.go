package source

import "testing"

func TestSFC(t *testing.T) {
	sfc := NewSFC(0)
	sfc.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		v := sfc.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

func BenchmarkSFC(b *testing.B) {
	sfc := NewSFC(0)

	for i := 0; i < b.N; i++ {
		sfc.Uint64()
	}
}
