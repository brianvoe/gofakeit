package gofakeit

import (
	"math"
	"reflect"
	"testing"
)

type weightedFixedSource uint64

func (s weightedFixedSource) Uint64() uint64 { return uint64(s) }

func TestWeightedOverflowingFiniteSum(t *testing.T) {
	options := []any{"first", "second"}
	for name, draw := range map[string]uint64{"zero": 0, "quarter": 1 << 22, "half": 1 << 23, "three quarters": 3 << 22, "last": (1 << 24) - 1} {
		t.Run(name, func(t *testing.T) {
			// Float32 uses the low 24 bits of Uint32, which uses Uint64's high 32 bits.
			newFaker := func() *Faker { return NewFaker(weightedFixedSource(draw<<32), false) }
			got, err := newFaker().Weighted(options, []float32{math.MaxFloat32, math.MaxFloat32})
			want, werr := newFaker().Weighted(options, []float32{1, 1})
			expected := "first"
			if draw > 1<<23 {
				expected = "second"
			}
			if want != expected {
				t.Fatalf("scaled control gave %v, want explicit %s", want, expected)
			}
			if err != nil || werr != nil || got != want {
				t.Fatalf("large equal finite weights gave %v,%v; scaled control gave %v,%v", got, err, want, werr)
			}
		})
	}
}

func TestWeightedOverflowMatchesScaledDistribution(t *testing.T) {
	cases := []struct {
		name          string
		large, scaled []float32
	}{
		{"equal", []float32{math.MaxFloat32, math.MaxFloat32}, []float32{1, 1}},
		{"unequal", []float32{math.MaxFloat32, math.MaxFloat32 / 2, math.MaxFloat32 / 4}, []float32{4, 2, 1}},
		{"zeros", []float32{0, math.MaxFloat32, 0, math.MaxFloat32, 0}, []float32{0, 1, 0, 1, 0}},
		{"last-heavy", []float32{math.MaxFloat32 / 4, math.MaxFloat32}, []float32{1, 4}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			options := make([]any, len(tc.large))
			for i := range options {
				options[i] = i
			}
			original := append([]float32(nil), tc.large...)
			for i := uint64(0); i <= 256; i++ {
				draw := i * ((1 << 24) - 1) / 256
				f := func() *Faker { return NewFaker(weightedFixedSource(draw<<32), false) }
				got, err := f().Weighted(options, tc.large)
				want, werr := f().Weighted(options, tc.scaled)
				if err != nil || werr != nil || got != want {
					t.Fatalf("draw %d: got %v,%v, scaled want %v,%v", draw, got, err, want, werr)
				}
			}
			if !reflect.DeepEqual(tc.large, original) {
				t.Fatal("Weighted mutated weights")
			}
		})
	}
}

type weightedCountingSource struct {
	draw  uint64
	calls int
}

func (s *weightedCountingSource) Uint64() uint64 { s.calls++; return s.draw }
func TestWeightedOverflowConsumesOneDraw(t *testing.T) {
	for _, draw := range []uint64{0, 1 << 54, (1 << 56) - (1 << 32)} {
		src := &weightedCountingSource{draw: draw}
		f := NewFaker(src, true)
		if _, err := f.Weighted([]any{1, 2}, []float32{math.MaxFloat32, math.MaxFloat32}); err != nil {
			t.Fatal(err)
		}
		if src.calls != 1 {
			t.Fatalf("draws=%d,want1", src.calls)
		}
	}
}

func TestWeightedOverflowLookupUsesProvidedFaker(t *testing.T) {
	info := GetFuncLookup("weighted")
	params := MapParams{"options": []string{"first", "second"}, "weights": []string{"3.4028235e38", "3.4028235e38"}}
	f := NewFaker(weightedFixedSource(0), true)
	got, err := info.Generate(f, &params, info)
	if err != nil || got != "first" {
		t.Fatalf("lookup got %v,%v,want first,nil", got, err)
	}
}
