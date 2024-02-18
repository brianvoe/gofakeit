package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLanguage() {
	Seed(11)
	fmt.Println(Language())

	// Output: Turkish
}

func ExampleFaker_Language() {
	f := New(11)
	fmt.Println(f.Language())

	// Output: Turkish
}

func BenchmarkLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Language()
	}
}

func ExampleLanguageAbbreviation() {
	Seed(11)
	fmt.Println(LanguageAbbreviation())

	// Output: tr
}

func ExampleFaker_LanguageAbbreviation() {
	f := New(11)
	fmt.Println(f.LanguageAbbreviation())

	// Output: tr
}

func BenchmarkLanguageAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageAbbreviation()
	}
}

func ExampleLanguageBCP() {
	Seed(11)
	fmt.Println(LanguageBCP())

	// Output: tr-TR
}

func ExampleFaker_LanguageBCP() {
	f := New(11)
	fmt.Println(f.LanguageBCP())

	// Output: tr-TR
}

func BenchmarkLanguageBCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageBCP()
	}
}

func ExampleProgrammingLanguage() {
	Seed(11)
	fmt.Println(ProgrammingLanguage())

	// Output: TELCOMP
}

func ExampleFaker_ProgrammingLanguage() {
	f := New(11)
	fmt.Println(f.ProgrammingLanguage())

	// Output: TELCOMP
}

func BenchmarkProgrammingLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProgrammingLanguage()
	}
}
