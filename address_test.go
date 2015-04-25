package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAddress() {
	fmt.Println(Address())
	// Output: 2503 East Manor land, Milesfurt, New Jersey 03746-2050
}

func BenchmarkAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Address()
	}
}

func ExampleStreet() {
	fmt.Println(Street())
	// Output: 1573 East Courts fort
}

func BenchmarkStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Street()
	}
}

func ExampleAddressNumber() {
	fmt.Println(AddressNumber())
	// Output: 14314
}

func BenchmarkAddressNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddressNumber()
	}
}

func ExampleAddressStreetPrefix() {
	fmt.Println(AddressStreetPrefix())
	// Output: New
}

func BenchmarkAddressStreetPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddressStreetPrefix()
	}
}

func ExampleAddressStreetName() {
	fmt.Println(AddressStreetName())
	// Output: Isle
}

func BenchmarkAddressStreetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddressStreetName()
	}
}

func ExampleAddressStreetSuffix() {
	fmt.Println(AddressStreetSuffix())
	// Output: burgh
}

func BenchmarkAddressStreetSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddressStreetSuffix()
	}
}

func ExampleCity() {
	fmt.Println(City())
	// Output: Gusikowskiview
}

func BenchmarkCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		City()
	}
}

func ExampleState() {
	fmt.Println(State())
	// Output: California
}

func BenchmarkState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		State()
	}
}

func ExampleStateAbr() {
	fmt.Println(StateAbr())
	// Output: FL
}

func BenchmarkStateAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StateAbr()
	}
}

func ExampleZip() {
	fmt.Println(Zip())
	// Output: 51063
}

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Zip()
	}
}

func ExampleCountry() {
	fmt.Println(Country())
	// Output: Macao
}

func BenchmarkCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Country()
	}
}
