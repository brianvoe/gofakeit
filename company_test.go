package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCompany() {
	Seed(11)
	fmt.Println(Company())
	// Output: Moen, Pagac and Wuckert
}

func BenchmarkCompany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Company()
	}
}

func TestCompany(t *testing.T) {
	for i := 0; i < 100; i++ {
		Company()
	}
}

func ExampleCompanySuffix() {
	Seed(11)
	fmt.Println(CompanySuffix())
	// Output: Inc
}

func BenchmarkCompanySuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompanySuffix()
	}
}

func ExampleBuzzWord() {
	Seed(11)
	fmt.Println(BuzzWord())
	// Output: disintermediate
}

func BenchmarkBuzzWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuzzWord()
	}
}

func ExampleBS() {
	Seed(11)
	fmt.Println(BS())
	// Output: front-end
}

func BenchmarkBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BS()
	}
}
