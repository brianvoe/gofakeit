package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())

	// Output: how many siblings do you have
}

func ExampleFaker_Phrase() {
	f := New(11)
	fmt.Println(f.Phrase())

	// Output: how many siblings do you have
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}

func ExamplePhraseNoun() {
	Seed(11)
	fmt.Println(PhraseNoun())

	// Output: a fear
}

func ExampleFaker_PhraseNoun() {
	f := New(11)
	fmt.Println(f.PhraseNoun())

	// Output: a fear
}

func BenchmarkPhraseNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseNoun()
	}
}

func ExamplePhraseVerb() {
	Seed(11)
	fmt.Println(PhraseVerb())

	// Output: bathe the jittery trip totally brightly under a troubling part scarcely unexpectedly
}

func ExampleFaker_PhraseVerb() {
	f := New(11)
	fmt.Println(f.PhraseVerb())

	// Output: bathe the jittery trip totally brightly under a troubling part scarcely unexpectedly
}

func BenchmarkPhraseVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseVerb()
	}
}

func ExamplePhraseAdverb() {
	Seed(11)
	fmt.Println(PhraseAdverb())

	// Output: successfully
}

func ExampleFaker_PhraseAdverb() {
	f := New(11)
	fmt.Println(f.PhraseAdverb())

	// Output: successfully
}

func BenchmarkPhraseAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseAdverb()
	}
}

func ExamplePhrasePreposition() {
	Seed(11)
	fmt.Println(PhrasePreposition())

	// Output: with an archipelago
}

func ExampleFaker_PhrasePreposition() {
	f := New(11)
	fmt.Println(f.PhrasePreposition())

	// Output: with an archipelago
}

func BenchmarkPhrasePreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhrasePreposition()
	}
}
