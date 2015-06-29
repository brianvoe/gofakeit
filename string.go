package gofakeit

import "math/rand"

// Letter will generate a single random lower case letter
func Letter() string {
	return randLetter()
}

// Lexify will replace ? will random generated letters
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// ShuffleStrings will randomize a slice of strings
func ShuffleStrings(str []string) []string {
	final := make([]string, len(str))
	perm := rand.Perm(len(str))
	for i, v := range perm {
		final[v] = str[i]
	}
	return final
}
