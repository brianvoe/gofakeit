package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePreposition() {
	Seed(11)
	fmt.Println(Preposition())

	// Output: instead of
}

func ExampleFaker_Preposition() {
	f := New(11)
	fmt.Println(f.Preposition())

	// Output: instead of
}

func BenchmarkPreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Preposition()
	}
}

func ExamplePrepositionSimple() {
	Seed(11)
	fmt.Println(PrepositionSimple())

	// Output: with
}

func ExampleFaker_PrepositionSimple() {
	f := New(11)
	fmt.Println(f.PrepositionSimple())

	// Output: with
}

func BenchmarkPrepositionSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionSimple()
	}
}

func ExamplePrepositionDouble() {
	Seed(11)
	fmt.Println(PrepositionDouble())

	// Output: next to
}

func ExampleFaker_PrepositionDouble() {
	f := New(11)
	fmt.Println(f.PrepositionDouble())

	// Output: next to
}

func BenchmarkPrepositionDouble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionDouble()
	}
}

func ExamplePrepositionCompound() {
	Seed(11)
	fmt.Println(PrepositionCompound())

	// Output: other than
}

func ExampleFaker_PrepositionCompound() {
	f := New(11)
	fmt.Println(f.PrepositionCompound())

	// Output: other than
}

func BenchmarkPrepositionCompound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrepositionCompound()
	}
}
