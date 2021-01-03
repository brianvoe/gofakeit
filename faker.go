package gofakeit

import (
	crand "crypto/rand"
	"encoding/binary"
	rand "math/rand"
	"sync"
)

// Create global variable to deal with global function call.
var globalFaker *Faker = New(0)

// Faker struct is the primary struct for using localized.
type Faker struct {
	Rand *rand.Rand
}

type cryptoRand struct {
	sync.Mutex
	buf []byte
}

func (c *cryptoRand) Seed(seed int64) {}

func (c *cryptoRand) Uint64() uint64 {
	// Lock to make reading thread safe
	c.Lock()
	defer c.Unlock()

	crand.Read(c.buf)
	return binary.BigEndian.Uint64(c.buf)
}

func (c *cryptoRand) Int63() int64 {
	return int64(c.Uint64() & ^uint64(1<<63))
}

// NewCrypto will utilize crypto/rand for concurrent pseudo random usage.
func NewCrypto() *Faker {
	return &Faker{Rand: rand.New(&cryptoRand{
		buf: make([]byte, 8),
	})}
}

// New will utilize math/rand for concurrent random usage.
// Setting seed to 0 will use crypto/rand for the initial seed number.
func New(seed int64) *Faker {
	// If passing 0 create crypto safe int64 for initial seed number
	if seed == 0 {
		binary.Read(crand.Reader, binary.BigEndian, &seed)
	}

	return &Faker{Rand: rand.New(rand.NewSource(seed))}
}

// NewCustom will utilize a custom rand.Source64 for concurrent random usage
// See https://golang.org/src/math/rand/rand.go for required interface methods
func NewCustom(source rand.Source64) *Faker {
	return &Faker{Rand: rand.New(source)}
}

// SetGlobalFaker will allow you to set what type of faker is globally used. Defailt is math/rand
func SetGlobalFaker(faker *Faker) {
	globalFaker = faker
}
