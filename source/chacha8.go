package source

import (
	"math/rand/v2"
	"sync"
)

type ChaCha8 struct {
	rand.ChaCha8

	// Lock to make reading thread safe
	sync.Mutex
	lock bool
}

// NewChaCha8 creates and returns a new ChaCha8 pseudo-random number generator seeded with a given seed.
func NewChaCha8(seed [32]byte, lock bool) *ChaCha8 {
	return &ChaCha8{
		ChaCha8: *rand.NewChaCha8(seed),
		lock:    lock,
	}
}

// Uint64 generates a pseudo-random 64-bit value using the ChaCha8 algorithm.
func (c *ChaCha8) Uint64() uint64 {
	// If locking is enabled, lock the generator to make it thread-safe
	if c.lock {
		c.Lock()
		defer c.Unlock()
	}

	return c.ChaCha8.Uint64()
}

// Seed sets the seed of the ChaCha8. This implementation can be enhanced to
// provide a more distributed seeding process across the state variables.
func (c *ChaCha8) Seed(seed [32]byte) {
	// If locking is enabled, lock the generator to make it thread-safe
	if c.lock {
		c.Lock()
		defer c.Unlock()
	}

	c.ChaCha8.Seed(seed)
}
