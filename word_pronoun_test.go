package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePronoun() {
	Seed(11)
	fmt.Println(Pronoun())

	// Output: me
}

func ExampleFaker_Pronoun() {
	f := New(11)
	fmt.Println(f.Pronoun())

	// Output: me
}

func BenchmarkPronoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pronoun()
	}
}

func ExamplePronounPersonal() {
	Seed(11)
	fmt.Println(PronounPersonal())

	// Output: it
}

func ExampleFaker_PronounPersonal() {
	f := New(11)
	fmt.Println(f.PronounPersonal())

	// Output: it
}

func BenchmarkPronounPersonal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounPersonal()
	}
}

func ExamplePronounObject() {
	Seed(11)
	fmt.Println(PronounObject())

	// Output: it
}

func ExampleFaker_PronounObject() {
	f := New(11)
	fmt.Println(f.PronounObject())

	// Output: it
}

func BenchmarkPronounObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounObject()
	}
}

func ExamplePronounPossessive() {
	Seed(11)
	fmt.Println(PronounPossessive())

	// Output: mine
}

func ExampleFaker_PronounPossessive() {
	f := New(11)
	fmt.Println(f.PronounPossessive())

	// Output: mine
}

func BenchmarkPronounPossessive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounPossessive()
	}
}

func ExamplePronounReflective() {
	Seed(11)
	fmt.Println(PronounReflective())

	// Output: myself
}

func ExampleFaker_PronounReflective() {
	f := New(11)
	fmt.Println(f.PronounReflective())

	// Output: myself
}

func BenchmarkPronounReflective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounReflective()
	}
}

func ExamplePronounIndefinite() {
	Seed(11)
	fmt.Println(PronounIndefinite())

	// Output: few
}

func ExampleFaker_PronounIndefinite() {
	f := New(11)
	fmt.Println(f.PronounIndefinite())

	// Output: few
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

	// Output: what
}

func ExampleFaker_PronounInterrogative() {
	f := New(11)
	fmt.Println(f.PronounInterrogative())

	// Output: what
}

func BenchmarkPronounInterrogative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounInterrogative()
	}
}

func ExamplePronounRelative() {
	Seed(11)
	fmt.Println(PronounRelative())

	// Output: as
}

func ExampleFaker_PronounRelative() {
	f := New(11)
	fmt.Println(f.PronounRelative())

	// Output: as
}

func BenchmarkPronounRelative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PronounRelative()
	}
}
