package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePronoun() {
	Seed(11)
	fmt.Println(Pronoun())

	// Output: some
}

func ExampleFaker_Pronoun() {
	f := New(11)
	fmt.Println(f.Pronoun())

	// Output: some
}

func BenchmarkPronoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pronoun()
	}
}

func ExamplePronounPersonal() {
	Seed(11)
	fmt.Println(PronounPersonal())

	// Output: they
}

func ExampleFaker_PronounPersonal() {
	f := New(11)
	fmt.Println(f.PronounPersonal())

	// Output: they
}

func BenchmarkPronounPersonal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounPersonal()
	}
}

func ExamplePronounObject() {
	Seed(11)
	fmt.Println(PronounObject())

	// Output: them
}

func ExampleFaker_PronounObject() {
	f := New(11)
	fmt.Println(f.PronounObject())

	// Output: them
}

func BenchmarkPronounObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounObject()
	}
}

func ExamplePronounPossessive() {
	Seed(11)
	fmt.Println(PronounPossessive())

	// Output: theirs
}

func ExampleFaker_PronounPossessive() {
	f := New(11)
	fmt.Println(f.PronounPossessive())

	// Output: theirs
}

func BenchmarkPronounPossessive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounPossessive()
	}
}

func ExamplePronounReflective() {
	Seed(11)
	fmt.Println(PronounReflective())

	// Output: itself
}

func ExampleFaker_PronounReflective() {
	f := New(11)
	fmt.Println(f.PronounReflective())

	// Output: itself
}

func BenchmarkPronounReflective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounReflective()
	}
}

func ExamplePronounIndefinite() {
	Seed(11)
	fmt.Println(PronounIndefinite())

	// Output: somebody
}

func ExampleFaker_PronounIndefinite() {
	f := New(11)
	fmt.Println(f.PronounIndefinite())

	// Output: somebody
}

func BenchmarkPronounIndefinite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounIndefinite()
	}
}

func ExamplePronounDemonstrative() {
	Seed(11)
	fmt.Println(PronounDemonstrative())

	// Output: this
}

func ExampleFaker_PronounDemonstrative() {
	f := New(11)
	fmt.Println(f.PronounDemonstrative())

	// Output: this
}

func BenchmarkPronounDemonstrative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounDemonstrative()
	}
}

func ExamplePronounInterrogative() {
	Seed(11)
	fmt.Println(PronounInterrogative())

	// Output: how
}

func ExampleFaker_PronounInterrogative() {
	f := New(11)
	fmt.Println(f.PronounInterrogative())

	// Output: how
}

func BenchmarkPronounInterrogative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounInterrogative()
	}
}

func ExamplePronounRelative() {
	Seed(11)
	fmt.Println(PronounRelative())

	// Output: whomever
}

func ExampleFaker_PronounRelative() {
	f := New(11)
	fmt.Println(f.PronounRelative())

	// Output: whomever
}

func BenchmarkPronounRelative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounRelative()
	}
}
