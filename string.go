package gofakeit

import "math/rand"

// Letter will generate a single random lower case ASCII letter
func Letter() string {
	return string(randLetter())
}

// Digit will generate a single ASCII digit
func Digit() string {
	return string(randDigit())
}

// Lexify will replace ? will random generated letters
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// ShuffleStrings will randomize a slice of strings
func ShuffleStrings(a []string) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

// RandString will take in a slice of string and return a randomly selected value
func RandString(a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	return a[rand.Intn(size)]
}
