package gofakeit

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

type paragrapOptions struct {
	paragraphCount int
	sentenceCount  int
	wordCount      int
	separator      string
}

const bytesPerWordEstimation = 6

type sentenceGenerator func(f *Faker, wordCount int) string
type wordGenerator func(f *Faker) string

// Sentence will generate a random sentence
func Sentence(wordCount int) string { return sentence(GlobalFaker, wordCount) }

// Sentence will generate a random sentence
func (f *Faker) Sentence(wordCount int) string { return sentence(f, wordCount) }

func sentence(f *Faker, wordCount int) string {
	return sentenceGen(f, wordCount, word)
}

// Paragraph will generate a random paragraphGenerator
func Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraph(GlobalFaker, paragraphCount, sentenceCount, wordCount, separator)
}

// Paragraph will generate a random paragraphGenerator
func (f *Faker) Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraph(f, paragraphCount, sentenceCount, wordCount, separator)
}

func paragraph(f *Faker, paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGen(f, paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, sentence)
}

func sentenceGen(f *Faker, wordCount int, word wordGenerator) string {
	if wordCount <= 0 {
		return ""
	}

	wordSeparator := ' '
	sentence := bytes.Buffer{}
	sentence.Grow(wordCount * bytesPerWordEstimation)

	for i := 0; i < wordCount; i++ {
		word := word(f)
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			word = string(runes)
		}
		sentence.WriteString(word)
		if i < wordCount-1 {
			sentence.WriteRune(wordSeparator)
		}
	}
	sentence.WriteRune('.')
	return sentence.String()
}

func paragraphGen(f *Faker, opts paragrapOptions, sentecer sentenceGenerator) string {
	if opts.paragraphCount <= 0 || opts.sentenceCount <= 0 || opts.wordCount <= 0 {
		return ""
	}

	//to avoid making Go 1.10 dependency, we cannot use strings.Builder
	paragraphs := bytes.Buffer{}
	//we presume the length
	paragraphs.Grow(opts.paragraphCount * opts.sentenceCount * opts.wordCount * bytesPerWordEstimation)
	wordSeparator := ' '

	for i := 0; i < opts.paragraphCount; i++ {
		for e := 0; e < opts.sentenceCount; e++ {
			paragraphs.WriteString(sentecer(f, opts.wordCount))
			if e < opts.sentenceCount-1 {
				paragraphs.WriteRune(wordSeparator)
			}
		}

		if i < opts.paragraphCount-1 {
			paragraphs.WriteString(opts.separator)
		}
	}

	return paragraphs.String()
}

// Question will return a random question
func Question() string {
	return question(GlobalFaker)
}

// Question will return a random question
func (f *Faker) Question() string {
	return question(f)
}

func question(f *Faker) string {
	return strings.Replace(hipsterSentence(f, number(f, 3, 10)), ".", "?", 1)
}

// Quote will return a random quote from a random person
func Quote() string { return quote(GlobalFaker) }

// Quote will return a random quote from a random person
func (f *Faker) Quote() string { return quote(f) }

func quote(f *Faker) string {
	return `"` + hipsterSentence(f, number(f, 3, 10)) + `" - ` + firstName(f) + " " + lastName(f)
}

func addWordSentenceLookup() {
	AddFuncLookup("sentence", Info{
		Display:     "Sentence",
		Category:    "word",
		Description: "Set of words expressing a statement, question, exclamation, or command",
		Example:     "Interpret context record river mind.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("invalid word count, must be greater than 0, less than 50")
			}

			return sentence(f, wordCount), nil
		},
	})

	AddFuncLookup("paragraph", Info{
		Display:     "Paragraph",
		Category:    "word",
		Description: "Distinct section of writing covering a single theme, composed of multiple sentences",
		Example:     "Interpret context record river mind press self should compare property outcome divide. Combine approach sustain consult discover explanation direct address church husband seek army. Begin own act welfare replace press suspect stay link place manchester specialist. Arrive price satisfy sign force application hair train provide basis right pay. Close mark teacher strengthen information attempt head touch aim iron tv take.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Display: "Paragraph Count", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Display: "Sentence Count", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseparator", Display: "Paragraph Separator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			paragraphCount, err := info.GetInt(m, "paragraphcount")
			if err != nil {
				return nil, err
			}
			if paragraphCount <= 0 || paragraphCount > 20 {
				return nil, errors.New("invalid paragraph count, must be greater than 0, less than 20")
			}

			sentenceCount, err := info.GetInt(m, "sentencecount")
			if err != nil {
				return nil, err
			}
			if sentenceCount <= 0 || sentenceCount > 20 {
				return nil, errors.New("invalid sentence count, must be greater than 0, less than 20")
			}

			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("invalid word count, must be greater than 0, less than 50")
			}

			paragraphSeparator, err := info.GetString(m, "paragraphseparator")
			if err != nil {
				return nil, err
			}

			return paragraph(f, paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})

	AddFuncLookup("question", Info{
		Display:     "Question",
		Category:    "word",
		Description: "Statement formulated to inquire or seek clarification",
		Example:     "Roof chia echo?",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return question(f), nil
		},
	})

	AddFuncLookup("quote", Info{
		Display:     "Quote",
		Category:    "word",
		Description: "Direct repetition of someone else's words",
		Example:     `"Roof chia echo." - Lura Lockman`,
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return quote(f), nil
		},
	})
}
