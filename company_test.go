package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCompany() {
	fmt.Println(Company())
	// Output: Bauch-Ritchie
}

func BenchmarkCompany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Company()
	}
}

func ExampleCompanySuffix() {
	fmt.Println(CompanySuffix())
	// Output: Inc
}

func BenchmarkCompanySuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompanySuffix()
	}
}

func ExampleBuzzWord() {
	fmt.Println(BuzzWord())
	// Output: Triple-buffered
}

func BenchmarkBuzzWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuzzWord()
	}
}

func ExampleBS() {
	fmt.Println(BS())
	// Output: distributed
}

func BenchmarkBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BS()
	}
}
