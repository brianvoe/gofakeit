package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleName() {
	Seed(11)
	fmt.Println(Name())
	// Output: Markus Moen
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}

func ExampleFirstName() {
	Seed(11)
	fmt.Println(FirstName())
	// Output: Markus
}

func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

func ExampleLastName() {
	Seed(11)
	fmt.Println(LastName())
	// Output: Daniel
}

func BenchmarkLastName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastName()
	}
}

func ExamplePrefixName() {
	Seed(11)
	fmt.Println(PrefixName())
	// Output: Mr.
}

func BenchmarkPrefixName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrefixName()
	}
}

func ExampleSuffixName() {
	Seed(11)
	fmt.Println(SuffixName())
	// Output: Jr.
}

func BenchmarkSuffixName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SuffixName()
	}
}
