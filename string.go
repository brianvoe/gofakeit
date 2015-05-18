package gofakeit

import "math/rand"

// Generate Random Lower Case Letter
func Letter() string {
	return randLetter()
}

// Generate Random Letters Replacing ? With Letters
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// Randomize []string
func ShuffleStrings(str []string) []string {
	final := make([]string, len(str))
	perm := rand.Perm(len(str))
	for i, v := range perm {
		final[v] = str[i]
	}
	return final
}
