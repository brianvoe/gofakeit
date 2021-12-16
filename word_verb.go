package gofakeit

import rand "math/rand"

// Verb will generate a random verb
func Verb() string { return verb(globalFaker.Rand) }

// Verb will generate a random verb
func (f *Faker) Verb() string { return verb(f.Rand) }

func verb(r *rand.Rand) string {
	var nounType = map[int]string{
		0: "verb_action",
		1: "verb_linking",
		2: "verb_helping",
	}
	return getRandValue(r, []string{"word", nounType[number(r, 0, 2)]})
}

// VerbAction will generate a random action verb
func VerbAction() string { return verbAction(globalFaker.Rand) }

// VerbAction will generate a random action verb
func (f *Faker) VerbAction() string { return verbAction(f.Rand) }

func verbAction(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_action"}) }

// VerbLinking will generate a random linking verb
func VerbLinking() string { return verbLinking(globalFaker.Rand) }

// VerbLinking will generate a random linking verb
func (f *Faker) VerbLinking() string { return verbLinking(f.Rand) }

func verbLinking(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_linking"}) }

// VerbHelping will generate a random helping verb
func VerbHelping() string { return verbHelping(globalFaker.Rand) }

// VerbHelping will generate a random helping verb
func (f *Faker) VerbHelping() string { return verbHelping(f.Rand) }

func verbHelping(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_helping"}) }

func addWordVerbLookup() {
	AddFuncLookup("verb", Info{
		Display:     "Verb",
		Category:    "word",
		Description: "Random verb",
		Example:     "release",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return verb(r), nil
		},
	})

	AddFuncLookup("verbaction", Info{
		Display:     "Action Verb",
		Category:    "word",
		Description: "Random action verb",
		Example:     "close",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return verbAction(r), nil
		},
	})

	AddFuncLookup("verblinking", Info{
		Display:     "Linking Verb",
		Category:    "word",
		Description: "Random linking verb",
		Example:     "was",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return verbLinking(r), nil
		},
	})

	AddFuncLookup("verbhelping", Info{
		Display:     "Helping Verb",
		Category:    "word",
		Description: "Random helping verb",
		Example:     "be",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return verbHelping(r), nil
		},
	})
}
