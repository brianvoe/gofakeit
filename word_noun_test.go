package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())

	// Output: aunt
}

func ExampleFaker_Noun() {
	f := New(11)
	fmt.Println(f.Noun())

	// Output: aunt
}

func BenchmarkNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Noun()
	}
}

func ExampleNounCommon() {
	Seed(11)
	fmt.Println(NounCommon())

	// Output: part
}

func ExampleFaker_NounCommon() {
	f := New(11)
	fmt.Println(f.NounCommon())

	// Output: part
}

func BenchmarkNounCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCommon()
	}
}

func ExampleNounConcrete() {
	Seed(11)
	fmt.Println(NounConcrete())

	// Output: snowman
}

func ExampleFaker_NounConcrete() {
	f := New(11)
	fmt.Println(f.NounConcrete())

	// Output: snowman
}

func BenchmarkNounConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounConcrete()
	}
}

func ExampleNounAbstract() {
	Seed(11)
	fmt.Println(NounAbstract())

	// Output: confusion
}

func ExampleFaker_NounAbstract() {
	f := New(11)
	fmt.Println(f.NounAbstract())

	// Output: confusion
}

func BenchmarkNounAbstract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounAbstract()
	}
}

func ExampleNounCollectivePeople() {
	Seed(11)
	fmt.Println(NounCollectivePeople())

	// Output: body
}

func ExampleFaker_NounCollectivePeople() {
	f := New(11)
	fmt.Println(f.NounCollectivePeople())

	// Output: body
}

func BenchmarkNounCollectivePeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectivePeople()
	}
}

func ExampleNounCollectiveAnimal() {
	Seed(11)
	fmt.Println(NounCollectiveAnimal())

	// Output: party
}

func ExampleFaker_NounCollectiveAnimal() {
	f := New(11)
	fmt.Println(f.NounCollectiveAnimal())

	// Output: party
}

func BenchmarkNounCollectiveAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveAnimal()
	}
}

func ExampleNounCollectiveThing() {
	Seed(11)
	fmt.Println(NounCollectiveThing())

	// Output: hand
}

func ExampleFaker_NounCollectiveThing() {
	f := New(11)
	fmt.Println(f.NounCollectiveThing())

	// Output: hand
}

func BenchmarkNounCollectiveThing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveThing()
	}
}

func ExampleNounCountable() {
	Seed(11)
	fmt.Println(NounCountable())

	// Output: neck
}

func ExampleFaker_NounCountable() {
	f := New(11)
	fmt.Println(f.NounCountable())

	// Output: neck
}

func BenchmarkNounCountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCountable()
	}
}

func ExampleNounUncountable() {
	Seed(11)
	fmt.Println(NounUncountable())

	// Output: seafood
}

func ExampleFaker_NounUncountable() {
	f := New(11)
	fmt.Println(f.NounUncountable())

	// Output: seafood
}

func BenchmarkNounUncountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounUncountable()
	}
}

func ExampleNounProper() {
	Seed(11)
	fmt.Println(NounProper())

	// Output: Marcel
}

func ExampleFaker_NounProper() {
	f := New(11)
	fmt.Println(f.NounProper())

	// Output: Marcel
}

func TestNounProper(t *testing.T) {
	f := New(11)
	for i := 0; i < 100; i++ {
		if f.NounProper() == "" {
			t.Errorf("Expected a non-empty string, got nothing")
		}
	}
}

func BenchmarkNounProper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounProper()
	}
}
