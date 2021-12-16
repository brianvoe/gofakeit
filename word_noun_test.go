package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())
	// Output: foot
}

func ExampleFaker_Noun() {
	f := New(11)
	fmt.Println(f.Noun())
	// Output: foot
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
	// Output: part
}

func ExampleFaker_NounConcrete() {
	f := New(11)
	fmt.Println(f.NounConcrete())
	// Output: part
}

func BenchmarkNounConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounConcrete()
	}
}

func ExampleNounAbstract() {
	Seed(11)
	fmt.Println(NounAbstract())
	// Output: part
}

func ExampleFaker_NounAbstract() {
	f := New(11)
	fmt.Println(f.NounAbstract())
	// Output: part
}

func BenchmarkNounAbstract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounAbstract()
	}
}

func ExampleNounCollectivePeople() {
	Seed(11)
	fmt.Println(NounCollectivePeople())
	// Output: part
}

func ExampleFaker_NounCollectivePeople() {
	f := New(11)
	fmt.Println(f.NounCollectivePeople())
	// Output: part
}

func BenchmarkNounCollectivePeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectivePeople()
	}
}

func ExampleNounCollectiveAnimal() {
	Seed(11)
	fmt.Println(NounCollectiveAnimal())
	// Output: part
}

func ExampleFaker_NounCollectiveAnimal() {
	f := New(11)
	fmt.Println(f.NounCollectiveAnimal())
	// Output: part
}

func BenchmarkNounCollectiveAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveAnimal()
	}
}

func ExampleNounCollectiveThing() {
	Seed(11)
	fmt.Println(NounCollectiveThing())
	// Output: part
}

func ExampleFaker_NounCollectiveThing() {
	f := New(11)
	fmt.Println(f.NounCollectiveThing())
	// Output: part
}

func BenchmarkNounCollectiveThing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveThing()
	}
}

func ExampleNounCountable() {
	Seed(11)
	fmt.Println(NounCountable())
	// Output: part
}

func ExampleFaker_NounCountable() {
	f := New(11)
	fmt.Println(f.NounCountable())
	// Output: part
}

func BenchmarkNounCountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCountable()
	}
}

func ExampleNounUncountable() {
	Seed(11)
	fmt.Println(NounUncountable())
	// Output: part
}

func ExampleFaker_NounUncountable() {
	f := New(11)
	fmt.Println(f.NounUncountable())
	// Output: part
}

func BenchmarkNounUncountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounUncountable()
	}
}

func ExampleNounProper() {
	Seed(11)
	fmt.Println(NounProper())
	// Output: Arlington
}

func ExampleFaker_NounProper() {
	f := New(11)
	fmt.Println(f.NounProper())
	// Output: Arlington
}

func BenchmarkNounProper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounProper()
	}
}
