package gofakeit

import (
	"math/rand"
	"time"
)

// Seed random. Setting seed to 0 will use time.Now().UnixNano()
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}

type Faker struct {
	Rand *rand.Rand
}

func New(seed ...int64) *Faker {
	var s rand.Source

	if len(seed) > 0 {
		s = rand.NewSource(seed[0])
	} else {
		s = rand.NewSource(time.Now().UTC().UnixNano())
	}

	return &Faker{
		rand.New(s),
	}
}
