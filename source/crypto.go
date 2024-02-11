package source

import (
	crand "crypto/rand"
	"encoding/binary"
	"sync"
)

// Package source implements a cryptographically secure pseudo-random number generator (CSPRNG)
// using Go's crypto/rand. The Crypto type is designed for generating high-quality random
// uint64 values, suitable for cryptographic applications like secure token generation,
// cryptographic key creation, and other security-sensitive operations. It offers optional
// thread safety through a locking mechanism, making it suitable for concurrent usage.

// Pros:
// - Provides cryptographically secure randomness, suitable for security-sensitive applications.
// - Optional thread safety with locking, enabling safe concurrent access.

// Cons:
// - Locking mechanism, when enabled, may introduce performance overhead.
// - Does not utilize a seed, as it leverages the system's cryptographic RNG, which may be a
//   limitation in scenarios where deterministic pseudo-randomness is desired.

type Crypto struct {
	buf []byte

	// Lock to make reading thread safe
	sync.Mutex
	lock bool
}

// NewCrypto will utilize crypto/rand for concurrent random usage.
func NewCrypto(lock bool) *Crypto {
	return &Crypto{
		lock: lock,
		buf:  make([]byte, 8),
	}
}

func (c *Crypto) Seed() {}

func (c *Crypto) Uint64() uint64 {
	// Lock to make reading thread safe
	if c.lock {
		c.Lock()
		defer c.Unlock()
	}

	crand.Read(c.buf)
	return binary.BigEndian.Uint64(c.buf)
}
