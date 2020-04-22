package gofakeit

import "strings"

// HackerPhrase will return a random hacker sentence
func HackerPhrase() string {
	words := strings.Split(Generate(getRandValue([]string{"hacker", "phrase"})), " ")
	words[0] = strings.Title(words[0])
	return strings.Join(words, " ")
}

// HackerAbbreviation will return a random hacker abbreviation
func HackerAbbreviation() string {
	return getRandValue([]string{"hacker", "abbreviation"})
}

// HackerAdjective will return a random hacker adjective
func HackerAdjective() string {
	return getRandValue([]string{"hacker", "adjective"})
}

// HackerNoun will return a random hacker noun
func HackerNoun() string {
	return getRandValue([]string{"hacker", "noun"})
}

// HackerVerb will return a random hacker verb
func HackerVerb() string {
	return getRandValue([]string{"hacker", "verb"})
}

// HackeringVerb will return a random hacker ingverb
func HackeringVerb() string {
	return getRandValue([]string{"hacker", "ingverb"})
}

func addHackerLookup() {
	AddFuncLookup("hackerphrase", Info{
		Display:     "Hacker Phrase",
		Category:    "hacker",
		Description: "Random hacker phrase",
		Example:     "If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerPhrase(), nil
		},
	})

	AddFuncLookup("hackerabbreviation", Info{
		Display:     "Hacker Abbreviation",
		Category:    "hacker",
		Description: "Random hacker abbreviation",
		Example:     "ADP",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAbbreviation(), nil
		},
	})

	AddFuncLookup("hackeradjective", Info{
		Display:     "Hacker Adjective",
		Category:    "hacker",
		Description: "Random hacker adjective",
		Example:     "wireless",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAdjective(), nil
		},
	})

	AddFuncLookup("hackernoun", Info{
		Display:     "Hacker Noun",
		Category:    "hacker",
		Description: "Random hacker noun",
		Example:     "driver",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerNoun(), nil
		},
	})

	AddFuncLookup("hackerverb", Info{
		Display:     "Hacker Verb",
		Category:    "hacker",
		Description: "Random hacker verb",
		Example:     "synthesize",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerVerb(), nil
		},
	})

	AddFuncLookup("hackeringverb", Info{
		Display:     "Hackering Verb",
		Category:    "hacker",
		Description: "Random hackering verb",
		Example:     "connecting",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackeringVerb(), nil
		},
	})
}
