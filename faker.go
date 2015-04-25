package gofakeit

import (
	"math/rand"
	"time"
)

// Seed random
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}
