package source

import "testing"

func TestDumb(t *testing.T) {
	dumb := NewDumb(0)
	dumb.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		v := dumb.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

func BenchmarkDumb(b *testing.B) {
	dumb := NewDumb(0)

	for i := 0; i < b.N; i++ {
		dumb.Uint64()
	}
}
