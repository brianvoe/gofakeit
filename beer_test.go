package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBeerName() {
	Seed(11)
	fmt.Println(BeerName())

	// Output: Sierra Nevada Bigfoot Barleywine Style Ale
}

func ExampleFaker_BeerName() {
	f := New(11)
	fmt.Println(f.BeerName())

	// Output: Sierra Nevada Bigfoot Barleywine Style Ale
}

func BenchmarkBeerName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerName()
	}
}

func ExampleBeerStyle() {
	Seed(11)
	fmt.Println(BeerStyle())

	// Output: Vegetable Beer
}

func ExampleFaker_BeerStyle() {
	f := New(11)
	fmt.Println(f.BeerStyle())

	// Output: Vegetable Beer
}

func BenchmarkBeerStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerStyle()
	}
}

func ExampleBeerHop() {
	Seed(11)
	fmt.Println(BeerHop())

	// Output: TriplePearl
}

func ExampleFaker_BeerHop() {
	f := New(11)
	fmt.Println(f.BeerHop())

	// Output: TriplePearl
}

func BenchmarkBeerHop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerHop()
	}
}

func ExampleBeerYeast() {
	Seed(11)
	fmt.Println(BeerYeast())

	// Output: 2308 - Munich Lager
}

func ExampleFaker_BeerYeast() {
	f := New(11)
	fmt.Println(f.BeerYeast())

	// Output: 2308 - Munich Lager
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

func ExampleFaker_BeerMalt() {
	f := New(11)
	fmt.Println(f.BeerMalt())

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

	// Output: 91 IBU
}

func ExampleFaker_BeerIbu() {
	f := New(11)
	fmt.Println(f.BeerIbu())

	// Output: 91 IBU
}

func BenchmarkBeerIbu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerIbu()
	}
}

func ExampleBeerAlcohol() {
	Seed(11)
	fmt.Println(BeerAlcohol())

	// Output: 8.2%
}

func ExampleFaker_BeerAlcohol() {
	f := New(11)
	fmt.Println(f.BeerAlcohol())

	// Output: 8.2%
}

func BenchmarkBeerAlcohol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerAlcohol()
	}
}

func ExampleBeerBlg() {
	Seed(11)
	fmt.Println(BeerBlg())

	// Output: 16.6°Blg
}

func ExampleFaker_BeerBlg() {
	f := New(11)
	fmt.Println(f.BeerBlg())

	// Output: 16.6°Blg
}

func BenchmarkBeerBlg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BeerBlg()
	}
}
