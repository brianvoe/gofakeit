package source

import "testing"

func TestSimpleFastCounter(t *testing.T) {
	// Unlocked
	fast := NewSimpleFastCounter(0, false)
	fast.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		v := fast.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}

	// Locked
	fast = NewSimpleFastCounter(0, true)
	fast.Seed(0)

	// test for duplicates
	m = make(map[uint64]bool)
	for i := 0; i < 10000; i++ {
		v := fast.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

func BenchmarkSimpleFastCounter(b *testing.B) {
	b.Run("unlocked", func(b *testing.B) {
		fast := NewSimpleFastCounter(0, false)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})

	b.Run("locked", func(b *testing.B) {
		fast := NewSimpleFastCounter(0, true)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})
}
