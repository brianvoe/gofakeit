package source

import (
	"encoding/binary"
	"math"
	"math/bits"
	"math/rand/v2"
	"testing"
)

// PCG and ChaCha8 come from the standard library and show that the checks
// below accept a good generator. JSF and SFC have to pass them as well.
var qualitySources = []struct {
	name string
	new  func(seed uint64) rand.Source
}{
	{"PCG", func(seed uint64) rand.Source { return rand.NewPCG(seed, seed) }},
	{"ChaCha8", func(seed uint64) rand.Source {
		var key [32]byte
		binary.LittleEndian.PutUint64(key[:], seed)
		return rand.NewChaCha8(key)
	}},
	{"JSF", func(seed uint64) rand.Source { return NewJSF(seed) }},
	{"SFC", func(seed uint64) rand.Source { return NewSFC(seed) }},
}

var qualitySeeds = []uint64{0, 1, 2, 42, 123456789, 1 << 40}

// Every bit of the output is set in about half of the outputs.
func TestBitBalance(t *testing.T) {
	const n = 1 << 16
	sigma := math.Sqrt(n) / 2

	for _, tc := range qualitySources {
		for _, seed := range qualitySeeds {
			src := tc.new(seed)
			var ones [64]int
			for range n {
				v := src.Uint64()
				for bit := range 64 {
					ones[bit] += int(v >> bit & 1)
				}
			}

			worstBit, worst := 0, 0.0
			for bit, count := range ones {
				if dev := (float64(count) - n/2) / sigma; math.Abs(dev) > math.Abs(worst) {
					worstBit, worst = bit, dev
				}
			}
			if math.Abs(worst) > 5 {
				t.Errorf("%s seed %d: bit %d is set in %.2f%% of %d outputs (%.0f standard deviations from 50%%)",
					tc.name, seed, worstBit, 100*float64(ones[worstBit])/n, n, worst)
			}
		}
	}
}

// math/rand/v2 uses the low bits of the output when a range is a power of two
// (Bool, IntN(256) and so on), so they must not repeat after a short period.
func TestLowBitsDoNotRepeat(t *testing.T) {
	const window = 256
	const maxLag = 1 << 14

	for _, tc := range qualitySources {
		for _, seed := range qualitySeeds {
			src := tc.new(seed)
			out := make([]uint64, window+maxLag)
			for i := range out {
				out[i] = src.Uint64()
			}

			for _, mask := range []uint64{0x1, 0xff, 0xffff} {
				for lag := 1; lag <= maxLag; lag++ {
					repeats := true
					for i := range window {
						if (out[i]^out[i+lag])&mask != 0 {
							repeats = false
							break
						}
					}
					if repeats {
						t.Errorf("%s seed %d: the bits in mask 0x%x repeat every %d outputs (%d outputs in a row)",
							tc.name, seed, mask, lag, window)
						break
					}
				}
			}
		}
	}
}

// Seeds are often small neighbouring numbers, so the first outputs of seed s
// and seed s+1 must not look alike.
func TestNearbySeedsDiffer(t *testing.T) {
	const seeds = 2000

	for _, tc := range qualitySources {
		total := 0
		for seed := range uint64(seeds) {
			total += bits.OnesCount64(tc.new(seed).Uint64() ^ tc.new(seed+1).Uint64())
		}

		// Unrelated outputs differ in 32 bits on average. One pair has a
		// standard deviation of 4 bits, so the mean of 2000 pairs is within 0.1.
		if mean := float64(total) / seeds; mean < 30 || mean > 34 {
			t.Errorf("%s: the first outputs of seeds s and s+1 differ in %.2f bits on average, want about 32", tc.name, mean)
		}
	}
}

// Every value of a small power of two range comes up equally often.
func TestPowerOfTwoRanges(t *testing.T) {
	const draws = 1 << 16

	for _, tc := range qualitySources {
		for _, n := range []int{2, 8, 256} {
			r := rand.New(tc.new(7))
			counts := make([]int, n)
			for range draws {
				counts[r.IntN(n)]++
			}

			want := float64(draws) / float64(n)
			sigma := math.Sqrt(want * (1 - 1/float64(n)))
			for value, count := range counts {
				if math.Abs(float64(count)-want) > 5*sigma {
					t.Errorf("%s IntN(%d): %d came up %d times in %d draws, want about %.0f", tc.name, n, value, count, draws, want)
					break
				}
			}
		}
	}
}
