package gofakeit

import "math/rand"

// Generate Random Number Between Min And Max Number
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Replace # With Random Numerical Values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}

// Randomize []int
func ShuffleInts(ints []int) []int {
	final := make([]int, len(ints))
	perm := rand.Perm(len(ints))
	for i, v := range perm {
		final[v] = ints[i]
	}
	return final
}
