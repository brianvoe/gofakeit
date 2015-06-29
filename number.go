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
func ShuffleInts(ints []int) []int {
	final := make([]int, len(ints))
	perm := rand.Perm(len(ints))
	for i, v := range perm {
		final[v] = ints[i]
	}
	return final
}
