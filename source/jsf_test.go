package source

import "testing"

func TestJSF(t *testing.T) {
	jsf := NewJSF(0)
	jsf.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		v := jsf.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

func BenchmarkJSF(b *testing.B) {
	jsf := NewJSF(0)

	for i := 0; i < b.N; i++ {
		jsf.Uint64()
	}
}
