package gofakeit

import (
	"bytes"
	"errors"
	"strings"
	"unicode"
)

// Comment will generate a random statement or remark expressing an opinion, observation, or reaction
func Comment() string { return comment(GlobalFaker) }

// Comment will generate a random statement or remark expressing an opinion, observation, or reaction
func (f *Faker) Comment() string { return comment(f) }

func comment(f *Faker) string {
	template := getRandValue(f, []string{"text", "comment"})

	// Generate using your faker
	comment, err := generate(f, template)
	if err != nil {
		return template
	}

	// ~16% chance to capitalize first letter
	if prob, _ := weighted(f, []any{0, 1}, []float32{5, 1}); prob == 1 {
		if len(comment) > 0 {
			comment = strings.ToUpper(comment[:1]) + comment[1:]
		}
	}

	// ~15% chance to add punctuation
	if prob, _ := weighted(f, []any{0, 1}, []float32{17, 3}); prob == 1 {
		// Within punctuation: ? (4), ! (3), . (1)
		switch p, _ := weighted(f, []any{"?", "!", "."}, []float32{4, 3, 1}); p {
		case "?":
			comment += "?"
		case "!":
			comment += "!"
		case ".":
			comment += "."
		}
	}

	return comment
}

// Phrase will return a random phrase
func Phrase() string { return phrase(GlobalFaker) }

// Phrase will return a random phrase
func (f *Faker) Phrase() string { return phrase(f) }

func phrase(f *Faker) string { return getRandValue(f, []string{"text", "phrase"}) }

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
	switch prob {
	case 1:
		str = getArticle(str) + " " + str
	case 2:
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

// Sentence will generate a random sentence
func Sentence() string { return sentence(GlobalFaker) }

// Sentence will generate a random sentence
func (f *Faker) Sentence() string { return sentence(f) }

func sentence(f *Faker) string {
	sentence, err := generate(f, getRandValue(f, []string{"text", "sentence"}))
	if err != nil {
		return ""
	}

	// Capitalize the first letter
	return strings.ToUpper(sentence[:1]) + sentence[1:]
}

// Paragraph will generate a random paragraph
func Paragraph() string {
	return paragraph(GlobalFaker)
}

// Paragraph will generate a random paragraph
func (f *Faker) Paragraph() string {
	return paragraph(f)
}

func paragraph(f *Faker) string {
	// generate 2-5 sentences
	sentenceCount := f.Number(2, 5)
	sentences := make([]string, sentenceCount)
	for i := 0; i < sentenceCount; i++ {
		sentences[i] = sentence(f)
	}

	return strings.Join(sentences, " ")
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
	question, err := generate(f, getRandValue(f, []string{"text", "question"}))
	if err != nil {
		return ""
	}

	// Capitalize the first letter and add a question mark
	return strings.ToUpper(question[:1]) + question[1:] + "?"
}

// Quote will return a random quote from a random person
func Quote() string { return quote(GlobalFaker) }

// Quote will return a random quote from a random person
func (f *Faker) Quote() string { return quote(f) }

func quote(f *Faker) string {
	quote, err := generate(f, getRandValue(f, []string{"text", "quote"}))
	if err != nil {
		return ""
	}

	// Capitalize the first letter after the opening quote
	return quote[:1] + strings.ToUpper(quote[1:2]) + quote[2:]
}

// LoremIpsumSentence will generate a random sentence
func LoremIpsumSentence(wordCount int) string {
	return loremIpsumSentence(GlobalFaker, wordCount)
}

// LoremIpsumSentence will generate a random sentence
func (f *Faker) LoremIpsumSentence(wordCount int) string {
	return loremIpsumSentence(f, wordCount)
}

func loremIpsumSentence(f *Faker, wordCount int) string {
	if wordCount <= 0 {
		return ""
	}

	sentence := bytes.Buffer{}
	sentence.Grow(wordCount * 6) // estimate 6 bytes per word

	for i := 0; i < wordCount; i++ {
		word := loremIpsumWord(f)
		if i == 0 {
			runes := []rune(word)
			runes[0] = unicode.ToTitle(runes[0])
			word = string(runes)
		}
		sentence.WriteString(word)
		if i < wordCount-1 {
			sentence.WriteRune(' ')
		}
	}
	sentence.WriteRune('.')
	return sentence.String()
}

// LoremIpsumParagraph will generate a random paragraphGenerator
func LoremIpsumParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return loremIpsumParagraph(GlobalFaker, paragraphCount, sentenceCount, wordCount, separator)
}

// LoremIpsumParagraph will generate a random paragraphGenerator
func (f *Faker) LoremIpsumParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	return loremIpsumParagraph(f, paragraphCount, sentenceCount, wordCount, separator)
}

func loremIpsumParagraph(f *Faker, paragraphCount int, sentenceCount int, wordCount int, separator string) string {
	if paragraphCount <= 0 || sentenceCount <= 0 || wordCount <= 0 {
		return ""
	}

	paragraphs := bytes.Buffer{}
	paragraphs.Grow(paragraphCount * sentenceCount * wordCount * 6) // estimate 6 bytes per word

	for i := 0; i < paragraphCount; i++ {
		for e := 0; e < sentenceCount; e++ {
			paragraphs.WriteString(loremIpsumSentence(f, wordCount))
			if e < sentenceCount-1 {
				paragraphs.WriteRune(' ')
			}
		}

		if i < paragraphCount-1 {
			paragraphs.WriteString(separator)
		}
	}

	return paragraphs.String()
}

func addTextLookup() {
	AddFuncLookup("comment", Info{
		Display:     "Comment",
		Category:    "text",
		Description: "Statement or remark expressing an opinion, observation, or reaction",
		Example:     "add some a little bit team",
		Output:      "string",
		Aliases: []string{
			"verbal statement", "expressed thought", "spoken remark", "communication element", "casual note",
		},
		Keywords: []string{
			"opinion", "observation", "reaction", "response", "feedback", "critique", "interpretation", "perspective", "reflection", "discussion",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return comment(f), nil
		},
	})

	AddFuncLookup("phrase", Info{
		Display:     "Phrase",
		Category:    "text",
		Description: "A small group of words standing together",
		Example:     "time will tell",
		Output:      "string",
		Aliases:     []string{"word group", "language unit", "text element", "expression block"},
		Keywords:    []string{"phrase", "words", "group", "sentence", "text", "language", "grammar", "expression", "unit", "collection"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phrase(f), nil
		},
	})

	AddFuncLookup("phrasenoun", Info{
		Display:     "Noun Phrase",
		Category:    "text",
		Description: "Phrase with a noun as its head, functions within sentence like a noun",
		Example:     "a tribe",
		Output:      "string",
		Aliases:     []string{"nominal phrase", "substantive element", "subject phrase", "object phrase"},
		Keywords:    []string{"phrase", "noun", "grammar", "subject", "object", "head", "sentence", "nominal", "substantive", "entity"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseNoun(f), nil
		},
	})

	AddFuncLookup("phraseverb", Info{
		Display:     "Verb Phrase",
		Category:    "text",
		Description: "Phrase that Consists of a verb and its modifiers, expressing an action or state",
		Example:     "a tribe",
		Output:      "string",
		Aliases:     []string{"predicate phrase", "verbal element", "action phrase", "state phrase"},
		Keywords:    []string{"phrase", "verb", "grammar", "action", "state", "modifiers", "sentence", "predicate", "verbal", "behavior"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseVerb(f), nil
		},
	})

	AddFuncLookup("phraseadverb", Info{
		Display:     "Adverb Phrase",
		Category:    "text",
		Description: "Phrase that modifies a verb, adjective, or another adverb, providing additional information.",
		Example:     "fully gladly",
		Output:      "string",
		Aliases:     []string{"adverbial phrase", "qualifier element", "modifier phrase", "description phrase"},
		Keywords:    []string{"phrase", "adverb", "grammar", "modifier", "description", "information", "adverbial", "qualifier", "modification"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phraseAdverb(f), nil
		},
	})

	AddFuncLookup("phrasepreposition", Info{
		Display:     "Preposition Phrase",
		Category:    "text",
		Description: "Phrase starting with a preposition, showing relation between elements in a sentence.",
		Example:     "out the black thing",
		Output:      "string",
		Aliases:     []string{"prepositional phrase", "relational element", "connection phrase", "grammar bridge"},
		Keywords:    []string{"phrase", "preposition", "grammar", "relation", "connection", "sentence", "prepositional", "relational", "linking"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phrasePreposition(f), nil
		},
	})

	AddFuncLookup("sentence", Info{
		Display:     "Sentence",
		Category:    "text",
		Description: "Set of words expressing a statement, question, exclamation, or command",
		Example:     "Guide person with kind affordances.",
		Output:      "string",
		Aliases:     []string{"complete thought", "grammatical unit", "word group", "linguistic element"},
		Keywords:    []string{"sentence", "complete", "thought", "grammatical", "unit", "word", "group", "expression", "clause", "utterance"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return sentence(f), nil
		},
	})

	AddFuncLookup("paragraph", Info{
		Display:     "Paragraph",
		Category:    "text",
		Description: "Distinct section of writing covering a single theme, composed of multiple sentences",
		Example:     "Protect the place under grumpy load. Decompose work into smaller group. Ruthlessly remove dead work.",
		Output:      "string",
		Aliases:     []string{"text block", "writing section", "thematic unit", "content block"},
		Keywords:    []string{"paragraph", "text", "block", "writing", "section", "theme", "sentences", "composition", "distinct", "passage", "content"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return paragraph(f), nil
		},
	})

	AddFuncLookup("question", Info{
		Display:     "Question",
		Category:    "text",
		Description: "Statement formulated to inquire or seek clarification",
		Example:     "Roof chia echo?",
		Output:      "string",
		Aliases: []string{
			"interrogative sentence",
			"information request",
			"asking phrase",
			"query prompt",
			"clarifying ask",
		},
		Keywords: []string{
			"question", "inquiry", "clarification", "interrogative",
			"ask", "who", "what", "when", "where", "why", "how", "mark",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return question(f), nil
		},
	})

	AddFuncLookup("quote", Info{
		Display:     "Quote",
		Category:    "text",
		Description: "Direct repetition of someone else's words",
		Example:     `"Does orange help the tissue or distract it"`,
		Output:      "string",
		Aliases: []string{
			"direct speech",
			"verbatim line",
			"cited passage",
			"attributed text",
			"pulled excerpt",
		},
		Keywords: []string{
			"quote", "quotation", "citation", "reference", "excerpt",
			"epigraph", "saying", "maxim", "attribution", "blockquote",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return quote(f), nil
		},
	})

	AddFuncLookup("loremipsumsentence", Info{
		Display:     "Lorem Ipsum Sentence",
		Category:    "text",
		Description: "Sentence of the Lorem Ipsum placeholder text used in design and publishing",
		Example:     "Quia quae repellat consequatur quidem.",
		Output:      "string",
		Aliases: []string{
			"lorem sentence",
			"ipsum sentence",
			"placeholder sentence",
			"latin sentence",
		},
		Keywords: []string{
			"lorem", "ipsum", "sentence", "placeholder",
			"latin", "dummy", "filler", "text",
			"typography", "mockup",
		},
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

			return loremIpsumSentence(f, wordCount), nil
		},
	})

	AddFuncLookup("loremipsumparagraph", Info{
		Display:     "Lorem Ipsum Paragraph",
		Category:    "text",
		Description: "Paragraph of the Lorem Ipsum placeholder text used in design and publishing",
		Example: `Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.

Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.

Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.`,
		Output: "string",
		Aliases: []string{
			"lorem paragraph",
			"ipsum paragraph",
			"placeholder paragraph",
			"latin paragraph",
		},
		Keywords: []string{
			"lorem", "ipsum", "paragraph", "placeholder",
			"latin", "dummy", "filler", "text",
			"typography", "mockup",
		},
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

			return loremIpsumParagraph(f, paragraphCount, sentenceCount, wordCount, paragraphSeparator), nil
		},
	})
}
