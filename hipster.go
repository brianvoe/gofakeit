package gofakeit

import (
	"bytes"
	"unicode"
)

// HipsterWord will return a single hipster word
func HipsterWord() string {
	return getRandValue([]string{"hipster", "word"})
}

// HipsterSentence will generate a random sentence
func HipsterSentence(wordCount int) string {
	if wordCount <= 0 {
		return ""
	}

	wordSeparator := ' '
	sentence := bytes.Buffer{}
	//we try to guess the size
	sentence.Grow(wordCount * 6)

	for i := 0; i < wordCount; i++ {
		word := getRandValue([]string{"hipster", "word"})
		if i == 0 { //First letter should be capital
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

// HipsterParagraph will generate a random paragraph
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Separator
func HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	if paragraphCount <= 0 || sentenceCount <= 0 || wordCount <= 0 {
		return ""
	}

	//to avoid making Go 1.10 dependency, we cannot use strings.Builder
	paragraphs := bytes.Buffer{}
	//we presume the length
	paragraphs.Grow(paragraphCount * sentenceCount * wordCount * 6)
	wordSeparator := ' '

	for i := 0; i < paragraphCount; i++ {
		for e := 0; e < sentenceCount; e++ {
			paragraphs.WriteString(HipsterSentence(wordCount))
			if e < sentenceCount-1 {
				paragraphs.WriteRune(wordSeparator)
			}
		}

		if i < paragraphCount-1 {
			paragraphs.WriteString(separator)
		}
	}

	return paragraphs.String()
}
