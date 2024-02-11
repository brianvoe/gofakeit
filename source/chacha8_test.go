package source

import "testing"

func TestChaCha8(t *testing.T) {
	// Unlocked
	fast := NewChaCha8([32]byte{0, 1, 2, 3, 4, 5}, false)
	fast.Seed([32]byte{5, 4, 3, 2, 1, 0})

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
	fast = NewChaCha8([32]byte{0, 1, 2, 3, 4, 5}, true)
	fast.Seed([32]byte{5, 4, 3, 2, 1, 0})

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

func BenchmarkChaCha8(b *testing.B) {
	b.Run("unlocked", func(b *testing.B) {
		fast := NewChaCha8([32]byte{0, 1, 2, 3, 4, 5}, false)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})

	b.Run("locked", func(b *testing.B) {
		fast := NewChaCha8([32]byte{0, 1, 2, 3, 4, 5}, true)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})
}
