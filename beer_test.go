package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBeerName() {
	Seed(11)
	fmt.Println(BeerName())
	// Output: Duvel
}

func BenchmarkBeerName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerName()
	}
}

func ExampleBeerStyle() {
	Seed(11)
	fmt.Println(BeerStyle())
	// Output: European Amber Lager
}

func BenchmarkBeerStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerStyle()
	}
}

func ExampleBeerHop() {
	Seed(11)
	fmt.Println(BeerHop())
	// Output: Glacier
}

func BenchmarkBeerHop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerHop()
	}
}

func ExampleBeerYeast() {
	Seed(11)
	fmt.Println(BeerYeast())
	// Output: 1388 - Belgian Strong Ale
}

func BenchmarkBeerYeast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerYeast()
	}
}

func ExampleBeerMalt() {
	Seed(11)
	fmt.Println(BeerMalt())
	// Output: Munich
}

func BenchmarkBeerMalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerMalt()
	}
}

func ExampleBeerIbu() {
	Seed(11)
	fmt.Println(BeerIbu())
	// Output: 29 IBU
}

func BenchmarkBeerIbu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerIbu()
	}
}

func ExampleBeerAlcohol() {
	Seed(11)
	fmt.Println(BeerAlcohol())
	// Output: 2.7%
}

func BenchmarkBeerAlcohol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerAlcohol()
	}
}

func ExampleBeerBlg() {
	Seed(11)
	fmt.Println(BeerBlg())
	// Output: 6.4°Blg
}

func BenchmarkBeerBlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerBlg()
	}
}
