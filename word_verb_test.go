package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleVerb() {
	Seed(11)
	fmt.Println(Verb())

	// Output: would
}

func ExampleFaker_Verb() {
	f := New(11)
	fmt.Println(f.Verb())

	// Output: would
}

func BenchmarkVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Verb()
	}
}

func ExampleVerbAction() {
	Seed(11)
	fmt.Println(VerbAction())

	// Output: paint
}

func ExampleFaker_VerbAction() {
	f := New(11)
	fmt.Println(f.VerbAction())

	// Output: paint
}

func BenchmarkVerbAction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbAction()
	}
}

func ExampleVerbTransitive() {
	Seed(11)
	fmt.Println(VerbTransitive())

	// Output: upgrade
}

func ExampleFaker_VerbTransitive() {
	f := New(11)
	fmt.Println(f.VerbTransitive())

	// Output: upgrade
}

func BenchmarkVerbTransitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbTransitive()
	}
}

func ExampleVerbIntransitive() {
	Seed(11)
	fmt.Println(VerbIntransitive())

	// Output: vomit
}

func ExampleFaker_VerbIntransitive() {
	f := New(11)
	fmt.Println(f.VerbIntransitive())

	// Output: vomit
}

func BenchmarkVerbIntransitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbIntransitive()
	}
}

func ExampleVerbLinking() {
	Seed(11)
	fmt.Println(VerbLinking())

	// Output: must
}

func ExampleFaker_VerbLinking() {
	f := New(11)
	fmt.Println(f.VerbLinking())

	// Output: must
}

func BenchmarkVerbLinking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbLinking()
	}
}

func ExampleVerbHelping() {
	Seed(11)
	fmt.Println(VerbHelping())

	// Output: am
}

func ExampleFaker_VerbHelping() {
	f := New(11)
	fmt.Println(f.VerbHelping())

	// Output: am
}

func BenchmarkVerbHelping(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VerbHelping()
	}
}
