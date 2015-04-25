package gofakeit

import (
	"strings"
)

// Generate Random Word
func Word() string {
	return getRandValue([]string{"lorem", "word"})
}

// Generate Random Sentence.
// Set word count.
func Sentence(wordCount int) string {
	words := []string{}
	sentence := ""
	for i := 0; i < wordCount; i++ {
		words = append(words, getRandValue([]string{"lorem", "word"}))
	}

	sentence = strings.Join(words, " ")
	return strings.ToUpper(sentence[:1]) + sentence[1:] + "."
}

// Generate Random Paragraph.
// Set Paragraph Count.
// Set Sentence Count.
// Set Word Count.
// Set Paragraph Seperator.
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
