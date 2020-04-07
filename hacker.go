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
	AddLookupData("hacker.abr", Info{
		Description: "Random hacker abbreviation",
		Example:     "ADP",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HackerAbbreviation(), nil
		},
	})

	AddLookupData("hacker.adj", Info{
		Description: "Random hacker adjective",
		Example:     "wireless",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HackerAdjective(), nil
		},
	})

	AddLookupData("hacker.verb", Info{
		Description: "Random hacker verb",
		Example:     "synthesize",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HackerVerb(), nil
		},
	})

	AddLookupData("hacker.noun", Info{
		Description: "Random hacker noun",
		Example:     "driver",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HackerNoun(), nil
		},
	})

	AddLookupData("hacker.phrase", Info{
		Description: "Random hacker phrase",
		Example:     "If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HackerPhrase(), nil
		},
	})
}
