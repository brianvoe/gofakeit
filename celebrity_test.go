package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCelebrityActor() {
	Seed(11)
	fmt.Println(CelebrityActor())

	// Output: Owen Wilson
}

func ExampleFaker_CelebrityActor() {
	f := New(11)
	fmt.Println(f.CelebrityActor())

	// Output: Owen Wilson
}

func BenchmarkCelebrityActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebrityActor()
	}
}

func ExampleCelebrityBusiness() {
	Seed(11)
	fmt.Println(CelebrityBusiness())

	// Output: Cameron Diaz
}

func ExampleFaker_CelebrityBusiness() {
	f := New(11)
	fmt.Println(f.CelebrityBusiness())

	// Output: Cameron Diaz
}

func BenchmarkCelebrityBusiness(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebrityBusiness()
	}
}

func ExampleCelebritySport() {
	Seed(11)
	fmt.Println(CelebritySport())

	// Output: Hicham El Guerrouj
}

func ExampleFaker_CelebritySport() {
	f := New(11)
	fmt.Println(f.CelebritySport())

	// Output: Hicham El Guerrouj
}

func BenchmarkCelebritySport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebritySport()
	}
}
