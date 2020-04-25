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

// Noun will generate a random noun
func Noun() string {
	return getRandValue([]string{"word", "noun"})
}

// Verb will generate a random verb
func Verb() string {
	return getRandValue([]string{"word", "verb"})
}

// Adverb will generate a random adverb
func Adverb() string {
	return getRandValue([]string{"word", "adverb"})
}

// Preposition will generate a random preposition
func Preposition() string {
	return getRandValue([]string{"word", "preposition"})
}

// Adjective will generate a random adjective
func Adjective() string {
	return getRandValue([]string{"word", "adjective"})
}

// Word will generate a random word
func Word() string {
	if Bool() {
		return getRandValue([]string{"word", "noun"})
	}

	return getRandValue([]string{"word", "verb"})
}

// Sentence will generate a random sentence
func Sentence(wordCount int) string {
	return sentence(wordCount, Word)
}

// Paragraph will generate a random paragraphGenerator
func Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGenerator(paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, Sentence)
}

// LoremIpsumWord will generate a random word
func LoremIpsumWord() string {
	return getRandValue([]string{"lorem", "word"})
}

// LoremIpsumSentence will generate a random sentence
func LoremIpsumSentence(wordCount int) string {
	return sentence(wordCount, LoremIpsumWord)
}

// LoremIpsumParagraph will generate a random paragraphGenerator
func LoremIpsumParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return paragraphGenerator(paragrapOptions{paragraphCount, sentenceCount, wordCount, separator}, LoremIpsumSentence)
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

// Phrase will return a random dictionary phrase
func Phrase() string {
	return getRandValue([]string{"word", "phrase"})
}

func addWordLookup() {
	AddFuncLookup("noun", Info{
		Display:     "Noun",
		Category:    "word",
		Description: "Random noun",
		Example:     "foot",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Noun(), nil
		},
	})

	AddFuncLookup("verb", Info{
		Display:     "Verb",
		Category:    "word",
		Description: "Random verb",
		Example:     "release",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Verb(), nil
		},
	})

	AddFuncLookup("adverb", Info{
		Display:     "Adverb",
		Category:    "word",
		Description: "Random adverb",
		Example:     "smoothly",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Adverb(), nil
		},
	})

	AddFuncLookup("preposition", Info{
		Display:     "Preposition",
		Category:    "word",
		Description: "Random preposition",
		Example:     "down",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Preposition(), nil
		},
	})

	AddFuncLookup("adjective", Info{
		Display:     "Adjective",
		Category:    "word",
		Description: "Random adjective",
		Example:     "genuine",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Adjective(), nil
		},
	})

	AddFuncLookup("word", Info{
		Display:     "Word",
		Category:    "word",
		Description: "Random word",
		Example:     "man",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Word(), nil
		},
	})

	AddFuncLookup("sentence", Info{
		Display:     "Sentence",
		Category:    "word",
		Description: "Random sentence",
		Example:     "Interpret context record river mind.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
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

	AddFuncLookup("paragraph", Info{
		Display:     "Paragraph",
		Category:    "word",
		Description: "Random paragraph",
		Example:     "Interpret context record river mind press self should compare property outcome divide. Combine approach sustain consult discover explanation direct address church husband seek army. Begin own act welfare replace press suspect stay link place manchester specialist. Arrive price satisfy sign force application hair train provide basis right pay. Close mark teacher strengthen information attempt head touch aim iron tv take. Handle wait begin look speech trust cancer visit capacity disease chancellor clean. Race aim function gain couple push faith enjoy admit ring attitude develop. Edge game prevent cast mill favour father star live search aim guess. West heart item adopt compete equipment miss output report communicate model cabinet. Seek worker variety step argue air improve give succeed relief artist suffer. Hide finish insist knowledge thatcher make research chance structure proportion husband implement. Town crown restaurant cost material compete lady climb football region discussion order. Place lee market ice like display mind stress compete weather station raise. Democracy college major recall struggle use cut intention accept period generation strike. Benefit defend recommend conclude justify result depend succeed address owner fill interpret.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Display: "Paragraph Count", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Display: "Sentence Count", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseparator", Display: "Paragraph Separator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
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

			paragraphSeparator, err := info.GetString(m, "paragraphseparator")
			if err != nil {
				return nil, err
			}

			return Paragraph(paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})

	AddFuncLookup("loremipsumword", Info{
		Display:     "Lorem Ipsum Word",
		Category:    "word",
		Description: "Random lorem ipsum word",
		Example:     "quia",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return LoremIpsumWord(), nil
		},
	})

	AddFuncLookup("loremipsumsentence", Info{
		Display:     "Lorem Ipsum Sentence",
		Category:    "word",
		Description: "Random lorem ipsum sentence",
		Example:     "Quia quae repellat consequatur quidem.",
		Output:      "string",
		Params: []Param{
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			wordCount, err := info.GetInt(m, "wordcount")
			if err != nil {
				return nil, err
			}
			if wordCount <= 0 || wordCount > 50 {
				return nil, errors.New("Invalid word count, must be greater than 0, less than 50")
			}

			return LoremIpsumSentence(wordCount), nil
		},
	})

	AddFuncLookup("loremipsumparagraph", Info{
		Display:     "Lorem Ipsum Paragraph",
		Category:    "word",
		Description: "Random lorem ipsum paragraph",
		Example:     "Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.<br />Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.<br />Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.",
		Output:      "string",
		Params: []Param{
			{Field: "paragraphcount", Display: "Paragraph Count", Type: "int", Default: "2", Description: "Number of paragraphs"},
			{Field: "sentencecount", Display: "Sentence Count", Type: "int", Default: "2", Description: "Number of sentences in a paragraph"},
			{Field: "wordcount", Display: "Word Count", Type: "int", Default: "5", Description: "Number of words in a sentence"},
			{Field: "paragraphseparator", Display: "Paragraph Separator", Type: "string", Default: "<br />", Description: "String value to add between paragraphs"},
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

			paragraphSeparator, err := info.GetString(m, "paragraphseparator")
			if err != nil {
				return nil, err
			}

			return LoremIpsumParagraph(paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})

	AddFuncLookup("question", Info{
		Display:     "Question",
		Category:    "word",
		Description: "Random question",
		Example:     "Roof chia echo?",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Question(), nil
		},
	})

	AddFuncLookup("quote", Info{
		Display:     "Qoute",
		Category:    "word",
		Description: "Random quote",
		Example:     `"Roof chia echo." - Lura Lockman`,
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Quote(), nil
		},
	})

	AddFuncLookup("phrase", Info{
		Display:     "Phrase",
		Category:    "word",
		Description: "Random phrase",
		Example:     "time will tell",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Phrase(), nil
		},
	})
}
