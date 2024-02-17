package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCelebrityActor() {
	Seed(11)
	fmt.Println(CelebrityActor())

	// Output: Shah Rukh Khan
}

func ExampleFaker_CelebrityActor() {
	f := New(11)
	fmt.Println(f.CelebrityActor())

	// Output: Shah Rukh Khan
}

func BenchmarkCelebrityActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebrityActor()
	}
}

func ExampleCelebrityBusiness() {
	Seed(11)
	fmt.Println(CelebrityBusiness())

	// Output: Prescott Bush
}

func ExampleFaker_CelebrityBusiness() {
	f := New(11)
	fmt.Println(f.CelebrityBusiness())

	// Output: Prescott Bush
}

func BenchmarkCelebrityBusiness(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebrityBusiness()
	}
}

func ExampleCelebritySport() {
	Seed(11)
	fmt.Println(CelebritySport())

	// Output: Grete Waitz
}

func ExampleFaker_CelebritySport() {
	f := New(11)
	fmt.Println(f.CelebritySport())

	// Output: Grete Waitz
}

func BenchmarkCelebritySport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CelebritySport()
	}
}
