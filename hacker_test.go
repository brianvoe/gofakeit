package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHackerPhrase() {
	Seed(11)
	fmt.Println(HackerPhrase())
	// Output: If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!
}

func BenchmarkHackerPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerPhrase()
	}
}

func ExampleHackerAbbreviation() {
	Seed(11)
	fmt.Println(HackerAbbreviation())
	// Output: ADP
}

func BenchmarkHackerAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAbbreviation()
	}
}

func ExampleHackerAdjective() {
	Seed(11)
	fmt.Println(HackerAdjective())
	// Output: wireless
}

func BenchmarkHackerAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAdjective()
	}
}

func ExampleHackerNoun() {
	Seed(11)
	fmt.Println(HackerNoun())
	// Output: driver
}

func BenchmarkHackerNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerNoun()
	}
}

func ExampleHackerVerb() {
	Seed(11)
	fmt.Println(HackerVerb())
	// Output: synthesize
}

func BenchmarkHackerVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerVerb()
	}
}

func ExampleHackeringVerb() {
	Seed(11)
	fmt.Println(HackeringVerb())
	// Output: connecting
}

func BenchmarkHackeringVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackeringVerb()
	}
}
