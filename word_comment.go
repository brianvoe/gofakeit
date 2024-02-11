package gofakeit

import (
	"strings"
)

// Comment will generate a random statement or remark expressing an opinion, observation, or reaction
func Comment() string { return comment(GlobalFaker) }

// Comment will generate a random statement or remark expressing an opinion, observation, or reaction
func (f *Faker) Comment() string { return comment(f) }

func comment(f *Faker) string {
	structures := [][]string{
		{"interjection", "adjective", "noun", "verb", "adverb"},
		{"noun", "verb", "preposition", "determiner", "adjective", "noun"},
		{"noun", "verb", "adverb"},
		{"adjective", "noun", "verb"},
		{"noun", "verb", "preposition", "noun"},
	}

	// Randomly select a structure
	structure := structures[number(f, 0, len(structures)-1)]

	// Build the sentence
	var commentParts []string
	for _, wordType := range structure {
		switch wordType {
		case "noun":
			commentParts = append(commentParts, noun(f))
		case "verb":
			commentParts = append(commentParts, verb(f))
		case "adjective":
			commentParts = append(commentParts, adjective(f))
		case "adverb":
			commentParts = append(commentParts, adverb(f))
		case "interjection":
			commentParts = append(commentParts, interjection(f))
		case "preposition":
			commentParts = append(commentParts, preposition(f))
		case "determiner":
			commentParts = append(commentParts, nounDeterminer(f))
		default:
			// Should never hit
			panic("Invalid word type")
		}
	}

	// Combine the words into a sentence
	sentence := strings.Join(commentParts, " ")

	// Capitalize the first letter
	sentence = title(sentence)

	// Add a period to the end of the sentence
	sentence = sentence + "."

	return sentence
}

func addWordCommentLookup() {
	AddFuncLookup("comment", Info{
		Display:     "Comment",
		Category:    "word",
		Description: "Statement or remark expressing an opinion, observation, or reaction",
		Example:     "wow",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return interjection(f), nil
		},
	})
}
