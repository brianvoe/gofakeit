package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())

	// Output: horses for courses
}

func ExampleFaker_Phrase() {
	f := New(11)
	fmt.Println(f.Phrase())

	// Output: horses for courses
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}

func ExamplePhraseNoun() {
	Seed(11)
	fmt.Println(PhraseNoun())

	// Output: the purple tribe
}

func ExampleFaker_PhraseNoun() {
	f := New(11)
	fmt.Println(f.PhraseNoun())

	// Output: the purple tribe
}

func BenchmarkPhraseNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseNoun()
	}
}

func ExamplePhraseVerb() {
	Seed(11)
	fmt.Println(PhraseVerb())

	// Output: gladly dream indeed swiftly till a problem poorly
}

func ExampleFaker_PhraseVerb() {
	f := New(11)
	fmt.Println(f.PhraseVerb())

	// Output: gladly dream indeed swiftly till a problem poorly
}

func BenchmarkPhraseVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseVerb()
	}
}

func ExamplePhraseAdverb() {
	Seed(11)
	fmt.Println(PhraseAdverb())

	// Output: fully gladly
}

func ExampleFaker_PhraseAdverb() {
	f := New(11)
	fmt.Println(f.PhraseAdverb())

	// Output: fully gladly
}

func BenchmarkPhraseAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseAdverb()
	}
}

func ExamplePhrasePreposition() {
	Seed(11)
	fmt.Println(PhrasePreposition())

	// Output: out the tribe
}

func ExampleFaker_PhrasePreposition() {
	f := New(11)
	fmt.Println(f.PhrasePreposition())

	// Output: out the tribe
}

func BenchmarkPhrasePreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhrasePreposition()
	}
}
