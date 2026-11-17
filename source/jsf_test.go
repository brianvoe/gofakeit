package source

import "testing"

func TestJSF(t *testing.T) {
	jsf := NewJSF(0)
	jsf.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for range 10000 {
		v := jsf.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

// The expected values come from the 64 bit ranval and raninit on
// https://burtleburtle.net/bob/rand/smallprng.html
func TestJSFReference(t *testing.T) {
	tests := []struct {
		seed uint64
		want [5]uint64
	}{
		{0, [5]uint64{0x4b39c42db38fcdf5, 0xaee2c9e919833f29, 0x30611cd75d0254ce, 0x7fcfd4f0c54692bb, 0xb58f7ae8edf72d7b}},
		{11, [5]uint64{0x7e8a368271f63808, 0x53a9ff3820bdca66, 0x091d5697351f7c28, 0x4dda72f2a323122c, 0xcd81a1bd6beb1c0d}},
		{0xdeadbeef12345678, [5]uint64{0x5941d34c3090014b, 0xb25c5f93c9808ebd, 0xcede5387e31f17d5, 0x0750e1c9b8ce4d07, 0xd7f3dd6305e270a0}},
	}

	jsf := NewJSF(1)
	for _, tt := range tests {
		// Seed also has to restart a generator that is already in use.
		jsf.Seed(tt.seed)
		for i, want := range tt.want {
			if got := jsf.Uint64(); got != want {
				t.Errorf("seed %d output %d: got 0x%016x, want 0x%016x", tt.seed, i, got, want)
			}
		}
	}
}

func BenchmarkJSF(b *testing.B) {
	jsf := NewJSF(0)

	for i := 0; i < b.N; i++ {
		jsf.Uint64()
	}
}
