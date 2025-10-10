package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHipsterWord() {
	Seed(11)
	fmt.Println(HipsterWord())

	// Output: semiotics
}

func ExampleFaker_HipsterWord() {
	f := New(11)
	fmt.Println(f.HipsterWord())

	// Output: semiotics
}

func BenchmarkHipsterWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterWord()
	}
}

func ExampleHipsterSentence() {
	Seed(11)
	fmt.Println(HipsterSentence())

	// Output: Soul loops with you probably haven't heard of them undertones.
}

func ExampleFaker_HipsterSentence() {
	f := New(11)
	fmt.Println(f.HipsterSentence())

	// Output: Soul loops with you probably haven't heard of them undertones.
}

func BenchmarkHipsterSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterSentence()
	}
}

func ExampleHipsterParagraph() {
	Seed(11)
	fmt.Println(HipsterParagraph())

	// Output: Compared Care meets you probably haven't heard of them ethos. Single-origin austin, double why.
}

func ExampleFaker_HipsterParagraph() {
	f := New(11)
	fmt.Println(f.HipsterParagraph())

	// Output: Compared Care meets you probably haven't heard of them ethos. Single-origin austin, double why.
}

func BenchmarkHipsterParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterParagraph()
	}
}
