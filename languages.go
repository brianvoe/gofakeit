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
