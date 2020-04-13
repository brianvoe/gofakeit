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

// HackerIngverb will return a random hacker ingverb
func HackerIngverb() string {
	return getRandValue([]string{"hacker", "ingverb"})
}

func addHackerLookup() {
	AddLookupData("hackerabbreviation", Info{
		Category:    "hacker",
		Description: "Random hacker abbreviation",
		Example:     "ADP",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAbbreviation(), nil
		},
	})

	AddLookupData("hackeradjective", Info{
		Category:    "hacker",
		Description: "Random hacker adjective",
		Example:     "wireless",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerAdjective(), nil
		},
	})

	AddLookupData("hackerverb", Info{
		Category:    "hacker",
		Description: "Random hacker verb",
		Example:     "synthesize",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerVerb(), nil
		},
	})

	AddLookupData("hackernoun", Info{
		Category:    "hacker",
		Description: "Random hacker noun",
		Example:     "driver",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerNoun(), nil
		},
	})

	AddLookupData("hackerphrase", Info{
		Category:    "hacker",
		Description: "Random hacker phrase",
		Example:     "If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HackerPhrase(), nil
		},
	})
}
