package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAdjective() {
	Seed(11)
	fmt.Println(Adjective())

	// Output: none
}

func ExampleFaker_Adjective() {
	f := New(11)
	fmt.Println(f.Adjective())

	// Output: none
}

func BenchmarkAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adjective()
	}
}

func ExampleAdjectiveDescriptive() {
	Seed(11)
	fmt.Println(AdjectiveDescriptive())

	// Output: tired
}

func ExampleFaker_AdjectiveDescriptive() {
	f := New(11)
	fmt.Println(f.AdjectiveDescriptive())

	// Output: tired
}

func BenchmarkAdjectiveDescriptive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveDescriptive()
	}
}

func ExampleAdjectiveQuantitative() {
	Seed(11)
	fmt.Println(AdjectiveQuantitative())

	// Output: sparse
}

func ExampleFaker_AdjectiveQuantitative() {
	f := New(11)
	fmt.Println(f.AdjectiveQuantitative())

	// Output: sparse
}

func BenchmarkAdjectiveQuantitative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveQuantitative()
	}
}

func ExampleAdjectiveProper() {
	Seed(11)
	fmt.Println(AdjectiveProper())

	// Output: Swiss
}

func ExampleFaker_AdjectiveProper() {
	f := New(11)
	fmt.Println(f.AdjectiveProper())

	// Output: Swiss
}

func BenchmarkAdjectiveProper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveProper()
	}
}

func ExampleAdjectiveDemonstrative() {
	Seed(11)
	fmt.Println(AdjectiveDemonstrative())

	// Output: it
}

func ExampleFaker_AdjectiveDemonstrative() {
	f := New(11)
	fmt.Println(f.AdjectiveDemonstrative())

	// Output: it
}

func BenchmarkAdjectiveDemonstrative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveDemonstrative()
	}
}

func ExampleAdjectivePossessive() {
	Seed(11)
	fmt.Println(AdjectivePossessive())

	// Output: their
}

func ExampleFaker_AdjectivePossessive() {
	f := New(11)
	fmt.Println(f.AdjectivePossessive())

	// Output: their
}

func BenchmarkAdjectivePossessive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectivePossessive()
	}
}

func ExampleAdjectiveInterrogative() {
	Seed(11)
	fmt.Println(AdjectiveInterrogative())

	// Output: which
}

func ExampleFaker_AdjectiveInterrogative() {
	f := New(11)
	fmt.Println(f.AdjectiveInterrogative())

	// Output: which
}

func BenchmarkAdjectiveInterrogative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveInterrogative()
	}
}

func ExampleAdjectiveIndefinite() {
	Seed(11)
	fmt.Println(AdjectiveIndefinite())

	// Output: several
}

func ExampleFaker_AdjectiveIndefinite() {
	f := New(11)
	fmt.Println(f.AdjectiveIndefinite())

	// Output: several
}

func BenchmarkAdjectiveIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdjectiveIndefinite()
	}
}
