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

func ExampleNounProper() {
	Seed(11)
	fmt.Println(NounProper())
	// Output: part
}

func ExampleFaker_NounProper() {
	f := New(11)
	fmt.Println(f.NounProper())
	// Output: part
}

func BenchmarkNounProper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NounProper()
	}
}

func 