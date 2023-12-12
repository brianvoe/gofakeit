package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAdjective() {
	Seed(11)
	fmt.Println(Adjective())

	// Output: Dutch
}

func ExampleFaker_Adjective() {
	f := New(11)
	fmt.Println(f.Adjective())

	// Output: Dutch
}

func BenchmarkAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adjective()
	}
}

func ExampleAdjectiveDescriptive() {
	Seed(11)
	fmt.Println(AdjectiveDescriptive())

	// Output: brave
}

func ExampleFaker_AdjectiveDescriptive() {
	f := New(11)
	fmt.Println(f.AdjectiveDescriptive())

	// Output: brave
}

func BenchmarkAdjectiveDescriptive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveDescriptive()
	}
}

func ExampleAdjectiveQuantitative() {
	Seed(11)
	fmt.Println(AdjectiveQuantitative())

	// Output: a little
}

func ExampleFaker_AdjectiveQuantitative() {
	f := New(11)
	fmt.Println(f.AdjectiveQuantitative())

	// Output: a little
}

func BenchmarkAdjectiveQuantitative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveQuantitative()
	}
}

func ExampleAdjectiveProper() {
	Seed(11)
	fmt.Println(AdjectiveProper())

	// Output: Afghan
}

func ExampleFaker_AdjectiveProper() {
	f := New(11)
	fmt.Println(f.AdjectiveProper())

	// Output: Afghan
}

func BenchmarkAdjectiveProper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveProper()
	}
}

func ExampleAdjectiveDemonstrative() {
	Seed(11)
	fmt.Println(AdjectiveDemonstrative())

	// Output: this
}

func ExampleFaker_AdjectiveDemonstrative() {
	f := New(11)
	fmt.Println(f.AdjectiveDemonstrative())

	// Output: this
}

func BenchmarkAdjectiveDemonstrative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveDemonstrative()
	}
}

func ExampleAdjectivePossessive() {
	Seed(11)
	fmt.Println(AdjectivePossessive())

	// Output: our
}

func ExampleFaker_AdjectivePossessive() {
	f := New(11)
	fmt.Println(f.AdjectivePossessive())

	// Output: our
}

func BenchmarkAdjectivePossessive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectivePossessive()
	}
}

func ExampleAdjectiveInterrogative() {
	Seed(11)
	fmt.Println(AdjectiveInterrogative())

	// Output: what
}

func ExampleFaker_AdjectiveInterrogative() {
	f := New(11)
	fmt.Println(f.AdjectiveInterrogative())

	// Output: what
}

func BenchmarkAdjectiveInterrogative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveInterrogative()
	}
}

func ExampleAdjectiveIndefinite() {
	Seed(11)
	fmt.Println(AdjectiveIndefinite())

	// Output: few
}

func ExampleFaker_AdjectiveIndefinite() {
	f := New(11)
	fmt.Println(f.AdjectiveIndefinite())

	// Output: few
}

func BenchmarkAdjectiveIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveIndefinite()
	}
}
