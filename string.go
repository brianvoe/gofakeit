package gofakeit

import (
	"math/rand"
)

// Letter will generate a single random lower case ASCII letter
func Letter() string {
	return string(randLetter())
}

// Digit will generate a single ASCII digit
func Digit() string {
	return string(randDigit())
}

// Numerify will replace # with random numerical values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}

// Lexify will replace ? will random generated letters
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// ShuffleStrings will randomize a slice of strings
func ShuffleStrings(a []string) {
	swap := func(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
	//to avoid upgrading to 1.10 I copied the algorithm
	n := len(a)
	if n <= 1 {
		return
	}

	//if size is > int32 probably it will never finish, or ran out of entropy
	i := n - 1
	for ; i > 0; i-- {
		j := int(rand.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

// RandString will take in a slice of string and return a randomly selected value
func RandString(a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	return a[rand.Intn(size)]
}

func addStringLookup() {
	AddLookupData("letter", Info{
		Category:    "string",
		Description: "Generate a single random lower case ASCII letter",
		Example:     "g",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Letter(), nil
		},
	})

	AddLookupData("digit", Info{
		Category:    "string",
		Description: "Generate a single random lower case ASCII letter",
		Example:     "g",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Letter(), nil
		},
	})

	AddLookupData("numerify", Info{
		Category:    "string",
		Description: "Replace # with random numerical values",
		Example:     "(###)###-#### => (555)867-5309",
		Output:      "string",
		Params: []Param{
			{Field: "str", Type: "string", Description: "String value to replace #'s"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			return Numerify(str), nil
		},
	})

	AddLookupData("lexify", Info{
		Category:    "string",
		Description: "Replace ? will random generated letters",
		Example:     "?????@??????.com => billy@mister.com",
		Output:      "string",
		Params: []Param{
			{Field: "str", Type: "string", Description: "String value to replace #'s"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			return Lexify(str), nil
		},
	})

	AddLookupData("shufflestrings", Info{
		Category:    "string",
		Description: "Shuffle an array of strings",
		Example:     "hello,world,whats,up => whats,world,hello,up",
		Output:      "[]string",
		Params: []Param{
			{Field: "strs", Type: "[]string", Description: "Delimited seperated strings"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			strs, err := info.GetStringArray(m, "strs")
			if err != nil {
				return nil, err
			}

			ShuffleStrings(strs)

			return strs, nil
		},
	})
}
