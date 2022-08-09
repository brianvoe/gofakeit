package gofakeit

import (
	"math/rand"
	"strings"

	"github.com/simon0-o/gofakeit/v6/data"
)

// Word will generate a random word
func Word() string { return word(globalFaker.Rand) }

// Word will generate a random word
func (f *Faker) Word() string { return word(f.Rand) }

func word(r *rand.Rand) string {
	word := getRandValue(r, []string{"word", randomString(r, data.WordKeys)})

	// Word may return a couple of words, if so we will split on space and return a random word
	if strings.Contains(word, " ") {
		return randomString(r, strings.Split(word, " "))
	}

	return word
}

func addWordGeneralLookup() {
	AddFuncLookup("word", Info{
		Display:     "Word",
		Category:    "word",
		Description: "Random word",
		Example:     "man",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return word(r), nil
		},
	})
}
