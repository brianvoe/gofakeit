package source

import "testing"

func TestSFC(t *testing.T) {
	sfc := NewSFC(0)
	sfc.Seed(0)

	// test for duplicates
	m := make(map[uint64]bool)
	for range 10000 {
		v := sfc.Uint64()
		if m[v] {
			t.Errorf("Duplicate value: %v", v)
		}
		m[v] = true
	}
}

// The expected values come from numpy's SFC64 with its state set to
// seed, seed, seed, 1 and 12 outputs thrown away, as its sfc64_set_seed does.
func TestSFCReference(t *testing.T) {
	tests := []struct {
		seed uint64
		want [5]uint64
	}{
		{0, [5]uint64{0x3acfa029e3cc6041, 0xf5b6515bf2ee419c, 0x1259635894a29b61, 0x0b6ae75395f8ebd6, 0x225622285ce302e2}},
		{11, [5]uint64{0xc355446eaea52a10, 0x7d234f60e0fc58ae, 0x56e3ffe0e9426d06, 0xf228f51ddc7e3577, 0xe2aa20032f65c7fc}},
		{0xdeadbeef12345678, [5]uint64{0xe14a7e556a88e40b, 0x440514e4f466531b, 0x4a00d644902fa04c, 0x2c1bdcbd79bf2409, 0xaf8a0db8edbf25a3}},
	}

	sfc := NewSFC(1)
	for _, tt := range tests {
		// Seed also has to restart a generator that is already in use.
		sfc.Seed(tt.seed)
		for i, want := range tt.want {
			if got := sfc.Uint64(); got != want {
				t.Errorf("seed %d output %d: got 0x%016x, want 0x%016x", tt.seed, i, got, want)
			}
		}
	}
}

func BenchmarkSFC(b *testing.B) {
	sfc := NewSFC(0)

	for i := 0; i < b.N; i++ {
		sfc.Uint64()
	}
}
