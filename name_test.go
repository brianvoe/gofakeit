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

func ExampleNamePrefix() {
	Seed(11)
	fmt.Println(NamePrefix())
	// Output: Mr.
}

func BenchmarkNamePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NamePrefix()
	}
}

func ExampleNameSuffix() {
	Seed(11)
	fmt.Println(NameSuffix())
	// Output: Jr.
}

func BenchmarkNameSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NameSuffix()
	}
}
