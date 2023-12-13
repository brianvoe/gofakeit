package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAdverb() {
	Seed(11)
	fmt.Println(Adverb())

	// Output: over
}

func ExampleFaker_Adverb() {
	f := New(11)
	fmt.Println(f.Adverb())

	// Output: over
}

func BenchmarkAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adverb()
	}
}

func ExampleAdverbManner() {
	Seed(11)
	fmt.Println(AdverbManner())

	// Output: stupidly
}

func ExampleFaker_AdverbManner() {
	f := New(11)
	fmt.Println(f.AdverbManner())

	// Output: stupidly
}

func BenchmarkAdverbManner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbManner()
	}
}

func ExampleAdverbDegree() {
	Seed(11)
	fmt.Println(AdverbDegree())

	// Output: intensely
}

func ExampleFaker_AdverbDegree() {
	f := New(11)
	fmt.Println(f.AdverbDegree())

	// Output: intensely
}

func BenchmarkAdverbDegree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbDegree()
	}
}

func ExampleAdverbPlace() {
	Seed(11)
	fmt.Println(AdverbPlace())

	// Output: east
}

func ExampleFaker_AdverbPlace() {
	f := New(11)
	fmt.Println(f.AdverbPlace())

	// Output: east
}

func BenchmarkAdverbPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbPlace()
	}
}

func ExampleAdverbTimeDefinite() {
	Seed(11)
	fmt.Println(AdverbTimeDefinite())

	// Output: now
}

func ExampleFaker_AdverbTimeDefinite() {
	f := New(11)
	fmt.Println(f.AdverbTimeDefinite())

	// Output: now
}

func BenchmarkAdverbTimeDefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbTimeDefinite()
	}
}

func ExampleAdverbTimeIndefinite() {
	Seed(11)
	fmt.Println(AdverbTimeIndefinite())

	// Output: already
}

func ExampleFaker_AdverbTimeIndefinite() {
	f := New(11)
	fmt.Println(f.AdverbTimeIndefinite())

	// Output: already
}

func BenchmarkAdverbTimeIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbTimeIndefinite()
	}
}

func ExampleAdverbFrequencyDefinite() {
	Seed(11)
	fmt.Println(AdverbFrequencyDefinite())

	// Output: hourly
}

func ExampleFaker_AdverbFrequencyDefinite() {
	f := New(11)
	fmt.Println(f.AdverbFrequencyDefinite())

	// Output: hourly
}

func BenchmarkAdverbFrequencyDefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbFrequencyDefinite()
	}
}

func ExampleAdverbFrequencyIndefinite() {
	Seed(11)
	fmt.Println(AdverbFrequencyIndefinite())

	// Output: occasionally
}

func ExampleFaker_AdverbFrequencyIndefinite() {
	f := New(11)
	fmt.Println(f.AdverbFrequencyIndefinite())

	// Output: occasionally
}

func BenchmarkAdverbFrequencyIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbFrequencyIndefinite()
	}
}
