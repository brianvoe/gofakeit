package source

import "math/bits"

// The SFC(Simple Fast Counter) algorithm is designed for fast and efficient generation of pseudo-random numbers,
// utilizing arithmetic and bitwise operations across state variables and a counter to ensure
// good randomness quality. It is particularly well-suited for applications requiring rapid
// number generation without the need for cryptographic security.

// Pros:
// - High efficiency and speed, ideal for performance-sensitive applications.
// - Simple to implement and maintain, with minimal computational overhead.
// - Offers a balance between speed and randomness quality, suitable for a wide range of uses.

// Cons:
// - Not designed for cryptographic applications due to its level of randomness.
// - Initial seeding mechanism is basic; may require enhancement for more complex use cases.

type SFC struct {
	a, b, c, counter uint64
}

// NewSFC creates and returns a new SFC pseudo-random number generator seeded with a given seed.
func NewSFC(seed uint64) *SFC {
	s := &SFC{}
	s.Seed(seed)
	return s
}

// Seed sets the seed of the SFC like the original does: a, b and c are the seed, the counter
// starts at 1 and the first 12 outputs are thrown away, so that neighbouring seeds do not give neighbouring sequences.
func (s *SFC) Seed(seed uint64) {
	s.a = seed
	s.b = seed
	s.c = seed
	s.counter = 1 // Reset counter with new seed
	for range 12 {
		s.Uint64()
	}
}

// Uint64 generates a pseudo-random 64-bit value using the SFC64 algorithm.
func (s *SFC) Uint64() uint64 {
	result := s.a + s.b + s.counter
	s.counter++
	s.a = s.b ^ (s.b >> 11)
	s.b = s.c + (s.c << 3)
	s.c = bits.RotateLeft64(s.c, 24) + result
	return result
}
