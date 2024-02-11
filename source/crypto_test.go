package source

import "testing"

func TestCrypto(t *testing.T) {
	// Unlocked
	fast := NewCrypto(false)
	fast.Seed()

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
	fast = NewCrypto(true)
	fast.Seed()

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

func BenchmarkCrypto(b *testing.B) {
	b.Run("unlocked", func(b *testing.B) {
		fast := NewCrypto(false)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})

	b.Run("locked", func(b *testing.B) {
		fast := NewCrypto(true)

		for i := 0; i < b.N; i++ {
			fast.Uint64()
		}
	})
}
