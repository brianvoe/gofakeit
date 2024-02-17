package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHackerPhrase() {
	Seed(11)
	fmt.Println(HackerPhrase())

	// Output: Use the optical CSS microchip, then you can write the open-source monitor!
}

func ExampleFaker_HackerPhrase() {
	f := New(11)
	fmt.Println(f.HackerPhrase())

	// Output: Use the optical CSS microchip, then you can write the open-source monitor!
}

func BenchmarkHackerPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerPhrase()
	}
}

func ExampleHackerAbbreviation() {
	Seed(11)
	fmt.Println(HackerAbbreviation())

	// Output: SCSI
}

func ExampleFaker_HackerAbbreviation() {
	f := New(11)
	fmt.Println(f.HackerAbbreviation())

	// Output: SCSI
}

func BenchmarkHackerAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAbbreviation()
	}
}

func ExampleHackerAdjective() {
	Seed(11)
	fmt.Println(HackerAdjective())

	// Output: solid state
}

func ExampleFaker_HackerAdjective() {
	f := New(11)
	fmt.Println(f.HackerAdjective())

	// Output: solid state
}

func BenchmarkHackerAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerAdjective()
	}
}

func ExampleHackerNoun() {
	Seed(11)
	fmt.Println(HackerNoun())

	// Output: circuit
}

func ExampleFaker_HackerNoun() {
	f := New(11)
	fmt.Println(f.HackerNoun())

	// Output: circuit
}

func BenchmarkHackerNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerNoun()
	}
}

func ExampleHackerVerb() {
	Seed(11)
	fmt.Println(HackerVerb())

	// Output: lock
}

func ExampleFaker_HackerVerb() {
	f := New(11)
	fmt.Println(f.HackerVerb())

	// Output: lock
}

func BenchmarkHackerVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackerVerb()
	}
}

func ExampleHackeringVerb() {
	Seed(11)
	fmt.Println(HackeringVerb())

	// Output: compressing
}

func ExampleFaker_HackeringVerb() {
	f := New(11)
	fmt.Println(f.HackeringVerb())

	// Output: compressing
}

func BenchmarkHackeringVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HackeringVerb()
	}
}
