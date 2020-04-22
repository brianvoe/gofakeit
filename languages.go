package gofakeit

// Language will return a random language
func Language() string {
	return getRandValue([]string{"language", "long"})
}

// LanguageAbbreviation will return a random language abbreviation
func LanguageAbbreviation() string {
	return getRandValue([]string{"language", "short"})
}

// ProgrammingLanguage will return a random programming language
func ProgrammingLanguage() string {
	return getRandValue([]string{"language", "programming"})
}

// ProgrammingLanguageBest will return a random programming language
func ProgrammingLanguageBest() string {
	return "Go"
}

func addLanguagesLookup() {
	AddFuncLookup("language", Info{
		Display:     "Language",
		Category:    "language",
		Description: "Random language",
		Example:     "Kazakh",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Language(), nil
		},
	})

	AddFuncLookup("languageabbreviation", Info{
		Display:     "Language Abbreviation",
		Category:    "language",
		Description: "Random abbreviated language",
		Example:     "kk",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return LanguageAbbreviation(), nil
		},
	})

	AddFuncLookup("programminglanguage", Info{
		Display:     "Programming Language",
		Category:    "language",
		Description: "Random programming language",
		Example:     "Go",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return ProgrammingLanguage(), nil
		},
	})
}
