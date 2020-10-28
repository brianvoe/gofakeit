package gofakeit

import (
	"math/rand"
)

// Letter will generate a single random lower case ASCII letter
func Letter() string {
	return string(randLetter())
}

// LetterN will generate a random ASCII string with length N
func LetterN(n uint) string {
	// Make sure we dont use 0
	if n == 0 {
		n = 1
	}
	out := make([]rune, n)
	for i := 0; i < int(n); i++ {
		out[i] = randLetter()
	}
	return string(out)
}

// Digit will generate a single ASCII digit
func Digit() string {
	return string(randDigit())
}

// DigitN will generate a random string of length N consists of ASCII digits (note it can start with 0).
func DigitN(n uint) string {
	// Make sure we dont use 0
	if n == 0 {
		n = 1
	}
	out := make([]rune, n)
	for i := 0; i < int(n); i++ {
		out[i] = randDigit()
	}
	return string(out)
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

// RandomString will take in a slice of string and return a randomly selected value
func RandomString(a []string) string {
	size := len(a)
	if size == 0 {
		return ""
	}
	if size == 1 {
		return a[0]
	}
	return a[rand.Intn(size)]
}

func addStringLookup() {
	AddFuncLookup("letter", Info{
		Display:     "Letter",
		Category:    "string",
		Description: "Generate a single random lower case ASCII letter",
		Example:     "g",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Letter(), nil
		},
	})

	AddFuncLookup("lettern", Info{
		Display:     "LetterN",
		Category:    "string",
		Description: "Generate a random ASCII string with length N",
		Example:     "gbRMaRxHki",
		Output:      "string",
		Params: []Param{
			{Field: "count", Display: "Count", Type: "uint", Description: "Number of digits to generate"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			ui, err := info.GetUint(m, "count")
			if err != nil {
				return nil, err
			}

			return LetterN(ui), nil
		},
	})

	AddFuncLookup("digit", Info{
		Display:     "Digit",
		Category:    "string",
		Description: "Generate a single ASCII digit",
		Example:     "0",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Digit(), nil
		},
	})

	AddFuncLookup("digitn", Info{
		Display:     "DigitN",
		Category:    "string",
		Description: "Generate a random string of length N consists of ASCII digits",
		Example:     "0136459948",
		Output:      "string",
		Params: []Param{
			{Field: "count", Display: "Count", Type: "uint", Description: "Number of digits to generate"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			ui, err := info.GetUint(m, "count")
			if err != nil {
				return nil, err
			}

			return DigitN(ui), nil
		},
	})

	AddFuncLookup("numerify", Info{
		Display:     "Numerify",
		Category:    "string",
		Description: "Replace # with random numerical values",
		Example:     "(###)###-#### => (555)867-5309",
		Output:      "string",
		Params: []Param{
			{Field: "str", Display: "String", Type: "string", Description: "String value to replace #'s"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			return Numerify(str), nil
		},
	})

	AddFuncLookup("lexify", Info{
		Display:     "Lexify",
		Category:    "string",
		Description: "Replace ? will random generated letters",
		Example:     "?????@??????.com => billy@mister.com",
		Output:      "string",
		Params: []Param{
			{Field: "str", Display: "String", Type: "string", Description: "String value to replace #'s"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			return Lexify(str), nil
		},
	})

	AddFuncLookup("shufflestrings", Info{
		Display:     "Shuffle Strings",
		Category:    "string",
		Description: "Shuffle an array of strings",
		Example:     "hello,world,whats,up => whats,world,hello,up",
		Output:      "[]string",
		Params: []Param{
			{Field: "strs", Display: "Strings", Type: "[]string", Description: "Delimited separated strings"},
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

	AddFuncLookup("randomstring", Info{
		Display:     "Random String",
		Category:    "string",
		Description: "Randomly grab one string from array",
		Example:     "hello,world,whats,up => world",
		Output:      "[]string",
		Params: []Param{
			{Field: "strs", Display: "Strings", Type: "[]string", Description: "Delimited separated strings"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			strs, err := info.GetStringArray(m, "strs")
			if err != nil {
				return nil, err
			}

			return RandomString(strs), nil
		},
	})
}
