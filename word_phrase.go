package gofakeit

import (
	"strings"
)

// Phrase will return a random phrase
func Phrase() string { return phrase(GlobalFaker) }

// Phrase will return a random phrase
func (f *Faker) Phrase() string { return phrase(f) }

func phrase(f *Faker) string { return getRandValue(f, []string{"sentence", "phrase"}) }

// PhraseNoun will return a random noun phrase
func PhraseNoun() string { return phraseNoun(GlobalFaker) }

// PhraseNoun will return a random noun phrase
func (f *Faker) PhraseNoun() string { return phraseNoun(f) }

func phraseNoun(f *Faker) string {
	str := ""

	// You may also want to add an adjective to describe the noun
	if boolFunc(f) {
		str = adjectiveDescriptive(f) + " " + noun(f)
	} else {
		str = noun(f)
	}

	// Add determiner from weighted list
	prob, _ := weighted(f, []any{1, 2, 3}, []float32{2, 1.5, 1})
	if prob == 1 {
		str = getArticle(str) + " " + str
	} else if prob == 2 {
		str = "the " + str
	}

	return str
}

// PhraseVerb will return a random preposition phrase
func PhraseVerb() string { return phraseVerb(GlobalFaker) }

// PhraseVerb will return a random preposition phrase
func (f *Faker) PhraseVerb() string { return phraseVerb(f) }

func phraseVerb(f *Faker) string {
	// Put together a string builder
	sb := []string{}

	// You may have an adverb phrase
	if boolFunc(f) {
		sb = append(sb, phraseAdverb(f))
	}

	// Lets add the primary verb
	sb = append(sb, verbAction(f))

	// You may have a noun phrase
	if boolFunc(f) {
		sb = append(sb, phraseNoun(f))
	}

	// You may have an adverb phrase
	if boolFunc(f) {
		sb = append(sb, phraseAdverb(f))

		// You may also have a preposition phrase
		if boolFunc(f) {
			sb = append(sb, phrasePreposition(f))
		}

		// You may also hae an adverb phrase
		if boolFunc(f) {
			sb = append(sb, phraseAdverb(f))
		}
	}

	return strings.Join(sb, " ")
}

// PhraseAdverb will return a random adverb phrase
func PhraseAdverb() string { return phraseAdverb(GlobalFaker) }

// PhraseAdverb will return a random adverb phrase
func (f *Faker) PhraseAdverb() string { return phraseAdverb(f) }

func phraseAdverb(f *Faker) string {
	if boolFunc(f) {
		return adverbDegree(f) + " " + adverbManner(f)
	}

	return adverbManner(f)
}

// PhrasePreposition will return a random preposition phrase
func PhrasePreposition() string { return phrasePreposition(GlobalFaker) }

// PhrasePreposition will return a random preposition phrase
func (f *Faker) PhrasePreposition() string { return phrasePreposition(f) }

func phrasePreposition(f *Faker) string {
	return prepositionSimple(f) + " " + phraseNoun(f)
}

func addWordPhraseLookup() {
	AddFuncLookup("phrase", Info{
		Display:     "Phrase",
		Category:    "word",
		Description: "A small group of words standing together",
		Example:     "time will tell",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phrase(f), nil
		},
	})

	AddFuncLookup("phrasenoun", Info{
		Display:     "Noun Phrase",
		Category:    "word",
		Description: "Phrase with a noun as its head, functions within sentence like a noun",
		Example:     "a tribe",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseNoun(f), nil
		},
	})

	AddFuncLookup("phraseverb", Info{
		Display:     "Verb Phrase",
		Category:    "word",
		Description: "Phrase that Consists of a verb and its modifiers, expressing an action or state",
		Example:     "a tribe",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseVerb(f), nil
		},
	})

	AddFuncLookup("phraseadverb", Info{
		Display:     "Adverb Phrase",
		Category:    "word",
		Description: "Phrase that modifies a verb, adjective, or another adverb, providing additional information.",
		Example:     "fully gladly",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseAdverb(f), nil
		},
	})

	AddFuncLookup("phrasepreposition", Info{
		Display:     "Preposition Phrase",
		Category:    "word",
		Description: "Phrase starting with a preposition, showing relation between elements in a sentence.",
		Example:     "out the black thing",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phrasePreposition(f), nil
		},
	})
}
