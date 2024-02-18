package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())

	// Output: nest
}

func ExampleFaker_Noun() {
	f := New(11)
	fmt.Println(f.Noun())

	// Output: nest
}

func BenchmarkNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Noun()
	}
}

func ExampleNounCommon() {
	Seed(11)
	fmt.Println(NounCommon())

	// Output: group
}

func ExampleFaker_NounCommon() {
	f := New(11)
	fmt.Println(f.NounCommon())

	// Output: group
}

func BenchmarkNounCommon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCommon()
	}
}

func ExampleNounConcrete() {
	Seed(11)
	fmt.Println(NounConcrete())

	// Output: theater
}

func ExampleFaker_NounConcrete() {
	f := New(11)
	fmt.Println(f.NounConcrete())

	// Output: theater
}

func BenchmarkNounConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounConcrete()
	}
}

func ExampleNounAbstract() {
	Seed(11)
	fmt.Println(NounAbstract())

	// Output: speed
}

func ExampleFaker_NounAbstract() {
	f := New(11)
	fmt.Println(f.NounAbstract())

	// Output: speed
}

func BenchmarkNounAbstract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounAbstract()
	}
}

func ExampleNounCollectivePeople() {
	Seed(11)
	fmt.Println(NounCollectivePeople())

	// Output: posse
}

func ExampleFaker_NounCollectivePeople() {
	f := New(11)
	fmt.Println(f.NounCollectivePeople())

	// Output: posse
}

func BenchmarkNounCollectivePeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectivePeople()
	}
}

func ExampleNounCollectiveAnimal() {
	Seed(11)
	fmt.Println(NounCollectiveAnimal())

	// Output: leap
}

func ExampleFaker_NounCollectiveAnimal() {
	f := New(11)
	fmt.Println(f.NounCollectiveAnimal())

	// Output: leap
}

func BenchmarkNounCollectiveAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveAnimal()
	}
}

func ExampleNounCollectiveThing() {
	Seed(11)
	fmt.Println(NounCollectiveThing())

	// Output: hail
}

func ExampleFaker_NounCollectiveThing() {
	f := New(11)
	fmt.Println(f.NounCollectiveThing())

	// Output: hail
}

func BenchmarkNounCollectiveThing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCollectiveThing()
	}
}

func ExampleNounCountable() {
	Seed(11)
	fmt.Println(NounCountable())

	// Output: smile
}

func ExampleFaker_NounCountable() {
	f := New(11)
	fmt.Println(f.NounCountable())

	// Output: smile
}

func BenchmarkNounCountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounCountable()
	}
}

func ExampleNounUncountable() {
	Seed(11)
	fmt.Println(NounUncountable())

	// Output: usage
}

func ExampleFaker_NounUncountable() {
	f := New(11)
	fmt.Println(f.NounUncountable())

	// Output: usage
}

func BenchmarkNounUncountable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounUncountable()
	}
}

func ExampleNounProper() {
	Seed(11)
	fmt.Println(NounProper())

	// Output: Russ
}

func ExampleFaker_NounProper() {
	f := New(11)
	fmt.Println(f.NounProper())

	// Output: Russ
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

func ExampleNounDeterminer() {
	Seed(11)
	fmt.Println(NounDeterminer())

	// Output: every
}

func ExampleFaker_NounDeterminer() {
	f := New(11)
	fmt.Println(f.NounDeterminer())

	// Output: every
}

func TestNounDeterminer(t *testing.T) {
	f := New(11)
	for i := 0; i < 100; i++ {
		if f.NounDeterminer() == "" {
			t.Errorf("Expected a non-empty string, got nothing")
		}
	}
}

func BenchmarkNounDeterminer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounDeterminer()
	}
}
