package gofakeit

import rand "math/rand"

// Phrase will return a random dictionary phrase
func Phrase() string { return phrase(globalFaker.Rand) }

// Phrase will return a random dictionary phrase
func (f *Faker) Phrase() string { return phrase(f.Rand) }

func phrase(r *rand.Rand) string { return getRandValue(r, []string{"sentence", "phrase"}) }

func addWordPhraseLookup() {
	AddFuncLookup("phrase", Info{
		Display:     "Phrase",
		Category:    "word",
		Description: "Random phrase",
		Example:     "time will tell",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return phrase(r), nil
		},
	})
}
