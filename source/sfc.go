package source

import "sync"

// The SimpleFastCounter algorithm is designed for fast and efficient generation of pseudo-random numbers,
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

type SimpleFastCounter struct {
	a, b, c, counter uint64

	// Lock to make reading thread safe
	sync.Mutex
	lock bool
}

// NewSimpleFastCounter creates and returns a new SimpleFastCounter pseudo-random number generator seeded with a given seed.
func NewSimpleFastCounter(seed uint64, lock bool) *SimpleFastCounter {
	// Initialize the state with a seed. Here, we use a simple seeding method,
	// but you might consider a more sophisticated approach.
	return &SimpleFastCounter{
		a:       seed,
		b:       seed,
		c:       seed,
		counter: 1, // Start the counter at 1 to avoid an all-zero state
		lock:    lock,
	}
}

// Uint64 generates a pseudo-random 64-bit value using the SimpleFastCounter algorithm.
func (s *SimpleFastCounter) Uint64() uint64 {
	// If locking is enabled, lock the generator to make it thread-safe
	if s.lock {
		s.Lock()
		defer s.Unlock()
	}

	s.a += s.b + s.counter
	s.b ^= s.c
	s.c -= s.a
	s.counter++
	return s.c + s.b
}

// Seed sets the seed of the SimpleFastCounter. This implementation can be enhanced to
// provide a more distributed seeding process across the state variables.
func (s *SimpleFastCounter) Seed(seed int64) {
	// If locking is enabled, lock the generator to make it thread-safe
	if s.lock {
		s.Lock()
		defer s.Unlock()
	}

	s.a = uint64(seed)
	s.b = uint64(seed)
	s.c = uint64(seed)
	s.counter = 1 // Reset counter with new seed
}
