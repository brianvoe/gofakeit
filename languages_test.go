package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLanguage() {
	Seed(11)
	fmt.Println(Language())
	// Output: Kazakh
}

func BenchmarkLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Language()
	}
}

func ExampleLanguageAbbreviation() {
	Seed(11)
	fmt.Println(LanguageAbbreviation())
	// Output: kk
}

func BenchmarkLanguageAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageAbbreviation()
	}
}

func ExampleProgrammingLanguage() {
	Seed(464)
	fmt.Println(ProgrammingLanguage())
	// Output: Go
}

func BenchmarkProgrammingLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProgrammingLanguage()
	}
}

func ExampleProgrammingLanguageBest() {
	Seed(11)
	fmt.Println(ProgrammingLanguageBest())
	// Output: Go
}

func BenchmarkProgrammingLanguageBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProgrammingLanguageBest()
	}
}
