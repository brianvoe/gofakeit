package gofakeit

import "strings"

// HipsterWord will return a single hipster word
func HipsterWord() string { return hipsterWord(GlobalFaker) }

// HipsterWord will return a single hipster word
func (f *Faker) HipsterWord() string { return hipsterWord(f) }

func hipsterWord(f *Faker) string { return getRandValue(f, []string{"hipster", "word"}) }

// HipsterSentence will generate a random hipster sentence
func HipsterSentence() string { return hipsterSentence(GlobalFaker) }

// HipsterSentence will generate a random hipster sentence
func (f *Faker) HipsterSentence() string { return hipsterSentence(f) }

func hipsterSentence(f *Faker) string {
	sentence, err := generate(f, getRandValue(f, []string{"hipster", "sentence"}))
	if err != nil {
		return ""
	}

	// Capitalize the first letter
	sentence = strings.ToUpper(sentence[:1]) + sentence[1:]

	return sentence
}

// HipsterParagraph will generate a random hipster paragraph
func HipsterParagraph() string {
	return hipsterParagraph(GlobalFaker)
}

// HipsterParagraph will generate a random hipster paragraph
func (f *Faker) HipsterParagraph() string {
	return hipsterParagraph(f)
}

func hipsterParagraph(f *Faker) string {
	// generate 2-5 sentences
	sentenceCount := f.Number(2, 5)
	sentences := make([]string, sentenceCount)
	for i := 0; i < sentenceCount; i++ {
		sentences[i] = hipsterSentence(f)
	}

	return strings.Join(sentences, " ")
}

func addHipsterLookup() {
	AddFuncLookup("hipsterword", Info{
		Display:     "Hipster Word",
		Category:    "hipster",
		Description: "Trendy and unconventional vocabulary used by hipsters to express unique cultural preferences",
		Example:     "microdosing",
		Output:      "string",
		Aliases:     []string{"word", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords:    []string{"hipster", "preferences", "microdosing", "artisanal", "craft", "organic", "sustainable", "authentic"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterWord(f), nil
		},
	})

	AddFuncLookup("hipstersentence", Info{
		Display:     "Hipster Sentence",
		Category:    "hipster",
		Description: "Sentence showcasing the use of trendy and unconventional vocabulary associated with hipster culture",
		Example:     "Soul loops with you probably haven't heard of them undertones.",
		Output:      "string",
		Aliases:     []string{"sentence", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords:    []string{"hipster", "showcasing", "microdosing", "roof", "chia", "echo", "pickled", "artisanal"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterSentence(f), nil
		},
	})

	AddFuncLookup("hipsterparagraph", Info{
		Display:     "Hipster Paragraph",
		Category:    "hipster",
		Description: "Paragraph showcasing the use of trendy and unconventional vocabulary associated with hipster culture",
		Example:     "Single-origin austin, double why. Tag it Yuccie, keep it any. Ironically pug, sincerely several. Roof > helvetica, discuss. From France to Jersey, chasing ennui.",
		Output:      "string",
		Aliases:     []string{"paragraph", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords:    []string{"hipster", "showcasing", "meditation", "cold-pressed", "raw", "denim", "fingerstache", "normcore", "sriracha"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterParagraph(f), nil
		},
	})
}
