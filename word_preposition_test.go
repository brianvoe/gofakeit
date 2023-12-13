package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePreposition() {
	Seed(11)
	fmt.Println(Preposition())

	// Output: other than
}

func ExampleFaker_Preposition() {
	f := New(11)
	fmt.Println(f.Preposition())

	// Output: other than
}

func BenchmarkPreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Preposition()
	}
}

func ExamplePrepositionSimple() {
	Seed(11)
	fmt.Println(PrepositionSimple())

	// Output: out
}

func ExampleFaker_PrepositionSimple() {
	f := New(11)
	fmt.Println(f.PrepositionSimple())

	// Output: out
}

func BenchmarkPrepositionSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionSimple()
	}
}

func ExamplePrepositionDouble() {
	Seed(11)
	fmt.Println(PrepositionDouble())

	// Output: before
}

func ExampleFaker_PrepositionDouble() {
	f := New(11)
	fmt.Println(f.PrepositionDouble())

	// Output: before
}

func BenchmarkPrepositionDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionDouble()
	}
}

func ExamplePrepositionCompound() {
	Seed(11)
	fmt.Println(PrepositionCompound())

	// Output: according to
}

func ExampleFaker_PrepositionCompound() {
	f := New(11)
	fmt.Println(f.PrepositionCompound())

	// Output: according to
}

func BenchmarkPrepositionCompound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionCompound()
	}
}
