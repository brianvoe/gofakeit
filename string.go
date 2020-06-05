package gofakeit

import (
	"math/rand"
)

// Stores language UTF-8 bondary limits
type langRuneBoundary struct {
	start rune
	end   rune
}

// Language rune boundaries here
var (
	// LangENG is for english language
	LangENG = langRuneBoundary{65, 122}
	// LangCHI is for chinese language
	LangCHI = langRuneBoundary{19968, 40869}
	// LangRUS is for russian language
	LangRUS = langRuneBoundary{1025, 1105}
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

// RandomChars generate random sequence of chars of selected length and language (LangENG, LangCHI, LangRUS)
func RandomChars(len int, lang *langRuneBoundary) string {
	b := make([]rune, 0)
	for i := 0; i < len; {
		randRune := rune(rand.Intn(int(lang.end-lang.start)) + int(lang.start))
		b = append(b, randRune)
		i++
	}
	return string(b)
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

	AddFuncLookup("digit", Info{
		Display:     "Digit",
		Category:    "string",
		Description: "Generate a single random lower case ASCII letter",
		Example:     "g",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Letter(), nil
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
