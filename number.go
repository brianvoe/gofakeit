package gofakeit

import "math/rand"

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Numerify will replace # with random numerical values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
