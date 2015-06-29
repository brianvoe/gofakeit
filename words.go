package gofakeit

import (
	"strings"
)

// Word will generate a random word
func Word() string {
	return getRandValue([]string{"lorem", "word"})
}

// Sentence will generate a random sentence
func Sentence(wordCount int) string {
	words := []string{}
	sentence := ""
	for i := 0; i < wordCount; i++ {
		words = append(words, getRandValue([]string{"lorem", "word"}))
	}

	sentence = strings.Join(words, " ")
	return strings.ToUpper(sentence[:1]) + sentence[1:] + "."
}

// Paragraph will generate a random paragraph
// Set Paragraph Count
// Set Sentence Count
// Set Word Count
// Set Paragraph Seperator
func Paragraph(paragraphCount int, sentenceCount int, wordCount int, seperator string) string {
	sentences := []string{}
	paragraphs := []string{}
	for i := 0; i < paragraphCount; i++ {
		sentences = []string{}
		for e := 0; e < sentenceCount; e++ {
			sentences = append(sentences, Sentence(wordCount))
		}
		paragraphs = append(paragraphs, strings.Join(sentences, " "))
	}

	return strings.Join(paragraphs, seperator)
}
