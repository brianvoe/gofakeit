package gofakeit

import (
	"strings"

	"github.com/digitalmint/gofakeit/data"
)

// Word will generate a random word
func Word() string { return word(GlobalFaker) }

// Word will generate a random word
func (f *Faker) Word() string { return word(f) }

func word(f *Faker) string {
	word := getRandValue(f, []string{"word", randomString(f, data.WordKeys)})

	// Word may return a couple of words, if so we will split on space and return a random word
	if strings.Contains(word, " ") {
		return randomString(f, strings.Split(word, " "))
	}

	return word
}

func addWordGeneralLookup() {
	AddFuncLookup(
		"word", Info{
			Display:     "Word",
			Category:    "word",
			Description: "Basic unit of language representing a concept or thing, consisting of letters and having meaning",
			Example:     "man",
			Output:      "string",
			Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
				return word(f), nil
			},
		},
	)
}
