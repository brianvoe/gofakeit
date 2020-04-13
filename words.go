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

type sentenceGenerator func(wordCount int) string
type wordGenerator func() string

// Word will generate a random word
func Word() string {
	return getRandValue([]string{"lorem", "word"})
}

// Sentence will generate a random sentence
func Sentence(wordCount int) string {
	return sentence(wordCount, Word)
}

// Paragraph will generate a random paragraphGenerator
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGenerator(paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, Sentence)
}

func sentence(wordCount int, word wordGenerator) string {
	if wordCount <= 0 {
		return ""
	}

	wordSeparator := ' '
	sentence := bytes.Buffer{}
	sentence.Grow(wordCount * bytesPerWordEstimation)

	for i := 0; i < wordCount; i++ {
		word := word()
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

func paragraphGenerator(opts paragrapOptions, sentecer sentenceGenerator) string {
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
			paragraphs.WriteString(sentecer(opts.wordCount))
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
	return strings.Replace(HipsterSentence(Number(3, 10)), ".", "?", 1)
}

// Quote will return a random quote from a random person
func Quote() string {
	return `"` + HipsterSentence(Number(3, 10)) + `" - ` + FirstName() + " " + LastName()
}

func addWordLookup() {
	AddLookupData("word", Info{
		Category:    "word",
		Description: "Random lorem ipsum word",
		Example:     "quia",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Word(), nil
		},
	})

	AddLookupData("sentence", Info{
		Category:    "word",
		Description: "Random sentence",
		Example:     "Quia quae repellat consequatur quidem.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Type: "int", Default: "5", Description: "Number of words in a sentence"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			return Sentence(wordCount), nil
		},
	})

	AddLookupData("paragraph", Info{
		Category:    "word",
		Description: "Random paragraph",
		Example:     "Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.<br />Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.<br />Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseperator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			paragraphCount, err := info.GetInt(m, "paragraphcount")
			if err != nil {
				return nil, err
			}
			if paragraphCount <= 0 || paragraphCount > 20 {
				return nil, errors.New("Invalid paragraph count, must be greater than 0, less than 20")
			}

			sentenceCount, err := info.GetInt(m, "sentencecount")
			if err != nil {
				return nil, err
			}
			if sentenceCount <= 0 || sentenceCount > 20 {
				return nil, errors.New("Invalid sentence count, must be greater than 0, less than 20")
			}

			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			paragraphSeperator, err := info.GetString(m, "paragraphseperator")
			if err != nil {
				return nil, err
			}

			return Paragraph(paragraphCount, sentenceCount, wordCount, paragraphSeperator), nil
		},
	})

	AddLookupData("question", Info{
		Category:    "word",
		Description: "Random question",
		Example:     "Roof chia echo?",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Question(), nil
		},
	})

	AddLookupData("quote", Info{
		Category:    "word",
		Description: "Random quote",
		Example:     `"Roof chia echo." - Lura Lockman`,
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Quote(), nil
		},
	})
}
