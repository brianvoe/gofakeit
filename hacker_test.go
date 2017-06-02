package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHackerPhrase() {
	Seed(20)
	fmt.Println(HackerPhrase())
	// Output: Connecting the array won't do anything, we need to generate the haptic COM driver!
}

func BenchmarkHackerPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerPhrase()
	}
}

func ExampleHackerAbbreviation() {
	Seed(20)
	fmt.Println(HackerAbbreviation())
	// Output: AGP
}

func BenchmarkHackerAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAbbreviation()
	}
}

func ExampleHackerAdjective() {
	Seed(20)
	fmt.Println(HackerAdjective())
	// Output: online
}

func BenchmarkHackerAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAdjective()
	}
}

func ExampleHackerNoun() {
	Seed(20)
	fmt.Println(HackerNoun())
	// Output: pixel
}

func BenchmarkHackerNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerNoun()
	}
}

func ExampleHackerVerb() {
	Seed(20)
	fmt.Println(HackerVerb())
	// Output: connect
}

func BenchmarkHackerVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerVerb()
	}
}

func ExampleHackerIngverb() {
	Seed(20)
	fmt.Println(HackerIngverb())
	// Output: navigating
}

func BenchmarkHackerIngverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerIngverb()
	}
}
