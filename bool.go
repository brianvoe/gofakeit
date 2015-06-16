package gofakeit

import "math/rand"

// Generate Random Boolean value
func Bool() bool {
	if rand.Intn(2) == 1 {
		return true
	}

	return false
}
