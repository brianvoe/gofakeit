package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAdverb() {
	Seed(11)
	fmt.Println(Adverb())

	// Output: ever
}

func ExampleFaker_Adverb() {
	f := New(11)
	fmt.Println(f.Adverb())

	// Output: ever
}

func BenchmarkAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adverb()
	}
}

func ExampleAdverbManner() {
	Seed(11)
	fmt.Println(AdverbManner())

	// Output: tensely
}

func ExampleFaker_AdverbManner() {
	f := New(11)
	fmt.Println(f.AdverbManner())

	// Output: tensely
}

func BenchmarkAdverbManner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbManner()
	}
}

func ExampleAdverbDegree() {
	Seed(11)
	fmt.Println(AdverbDegree())

	// Output: too
}

func ExampleFaker_AdverbDegree() {
	f := New(11)
	fmt.Println(f.AdverbDegree())

	// Output: too
}

func BenchmarkAdverbDegree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbDegree()
	}
}

func ExampleAdverbPlace() {
	Seed(11)
	fmt.Println(AdverbPlace())

	// Output: under
}

func ExampleFaker_AdverbPlace() {
	f := New(11)
	fmt.Println(f.AdverbPlace())

	// Output: under
}

func BenchmarkAdverbPlace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbPlace()
	}
}

func ExampleAdverbTimeDefinite() {
	Seed(11)
	fmt.Println(AdverbTimeDefinite())

	// Output: yesterday
}

func ExampleFaker_AdverbTimeDefinite() {
	f := New(11)
	fmt.Println(f.AdverbTimeDefinite())

	// Output: yesterday
}

func BenchmarkAdverbTimeDefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbTimeDefinite()
	}
}

func ExampleAdverbTimeIndefinite() {
	Seed(11)
	fmt.Println(AdverbTimeIndefinite())

	// Output: soon
}

func ExampleFaker_AdverbTimeIndefinite() {
	f := New(11)
	fmt.Println(f.AdverbTimeIndefinite())

	// Output: soon
}

func BenchmarkAdverbTimeIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbTimeIndefinite()
	}
}

func ExampleAdverbFrequencyDefinite() {
	Seed(11)
	fmt.Println(AdverbFrequencyDefinite())

	// Output: yearly
}

func ExampleFaker_AdverbFrequencyDefinite() {
	f := New(11)
	fmt.Println(f.AdverbFrequencyDefinite())

	// Output: yearly
}

func BenchmarkAdverbFrequencyDefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbFrequencyDefinite()
	}
}

func ExampleAdverbFrequencyIndefinite() {
	Seed(11)
	fmt.Println(AdverbFrequencyIndefinite())

	// Output: generally
}

func ExampleFaker_AdverbFrequencyIndefinite() {
	f := New(11)
	fmt.Println(f.AdverbFrequencyIndefinite())

	// Output: generally
}

func BenchmarkAdverbFrequencyIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AdverbFrequencyIndefinite()
	}
}
