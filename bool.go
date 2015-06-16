package gofakeit

import "math/rand"

func Bool() bool {
	if rand.Intn(2) == 1 {
		return true
	}

	return false
}
