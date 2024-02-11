package source

import (
	"math/rand/v2"
	"sync"
)

type PCG struct {
	rand.PCG

	// Lock to make reading thread safe
	sync.Mutex
	lock bool
}

// NewPCG creates and returns a new PCG pseudo-random number generator seeded with a given seed.
func NewPCG(seed1, seed2 uint64, lock bool) *PCG {
	return &PCG{
		PCG:  *rand.NewPCG(seed1, seed2),
		lock: lock,
	}
}

// Uint64 generates a pseudo-random 64-bit value using the PCG algorithm.
func (p *PCG) Uint64() uint64 {
	// If locking is enabled, lock the generator to make it thread-safe
	if p.lock {
		p.Lock()
		defer p.Unlock()
	}

	return p.PCG.Uint64()
}

// Seed sets the seed of the PCG. This implementation can be enhanced to
// provide a more distributed seeding process across the state variables.
func (p *PCG) Seed(seed1, seed2 uint64) {
	// If locking is enabled, lock the generator to make it thread-safe
	if p.lock {
		p.Lock()
		defer p.Unlock()
	}

	p.PCG.Seed(seed1, seed2)
}
