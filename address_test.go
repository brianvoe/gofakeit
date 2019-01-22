package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAddress() {
	Seed(11)
	address := Address()
	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)
	// 872 East Rapidsborough, Rutherfordstad, New Jersey 74853
	// 872 East Rapidsborough
	// Rutherfordstad
	// New Jersey
	// 74853
	// South Africa
	// 23.05875828427908
	// 89.02259415693374
}

func BenchmarkAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Address()
	}
}

func ExampleStreet() {
	Seed(11)
	fmt.Println(Street())
	// Output: 364 East Rapidsborough
}

func BenchmarkStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Street()
	}
}

func ExampleStreetNumber() {
	Seed(11)
	fmt.Println(StreetNumber())
	// Output: 13645
}

func BenchmarkStreetNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetNumber()
	}
}

func ExampleStreetPrefix() {
	Seed(11)
	fmt.Println(StreetPrefix())
	// Output: Lake
}

func BenchmarkStreetPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetPrefix()
	}
}

func ExampleStreetName() {
	Seed(11)
	fmt.Println(StreetName())
	// Output: View
}

func BenchmarkStreetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetName()
	}
}

func ExampleStreetSuffix() {
	Seed(11)
	fmt.Println(StreetSuffix())
	// Output: land
}

func BenchmarkStreetSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetSuffix()
	}
}

func ExampleCity() {
	Seed(11)
	fmt.Println(City())
	// Output: Marcelside
}

func TestCity(t *testing.T) {
	for i := 0; i < 100; i++ {
		City()
	}
}

func BenchmarkCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		City()
	}
}

func ExampleState() {
	Seed(11)
	fmt.Println(State())
	// Output: Hawaii
}

func BenchmarkState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		State()
	}
}

func ExampleStateAbr() {
	Seed(11)
	fmt.Println(StateAbr())
	// Output: OR
}

func BenchmarkStateAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StateAbr()
	}
}

func ExampleZip() {
	Seed(11)
	fmt.Println(Zip())
	// Output: 13645
}

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Zip()
	}
}

func ExampleCountry() {
	Seed(11)
	fmt.Println(Country())
	// Output: Tajikistan
}

func BenchmarkCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Country()
	}
}

func ExampleCountryAbr() {
	Seed(11)
	fmt.Println(CountryAbr())
	// Output: FI
}

func BenchmarkCountryAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountryAbr()
	}
}

func ExampleLatitude() {
	Seed(11)
	fmt.Println(Latitude())
	// Output: -73.53405629980608
}

func BenchmarkLatitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Latitude()
	}
}

func ExampleLongitude() {
	Seed(11)
	fmt.Println(Longitude())
	// Output: -147.06811259961216
}

func BenchmarkLongitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Longitude()
	}
}

func TestLatitudeInRange(t *testing.T) {
	Seed(11)
	lat, err := LatitudeInRange(21, 42)
	if err != nil {
		t.Error("error should be nil")
	}

	if lat == 0 {
		t.Error("lat should be not be zero")
	}

	_, err = LatitudeInRange(50, 42)
	if err == nil {
		t.Error("error should be not be nil")
	}

	_, err = LatitudeInRange(-100, 42)
	if err == nil {
		t.Error("error should be not be nil")
	}
}

func ExampleLatitudeInRange() {
	Seed(11)
	lat, _ := LatitudeInRange(21, 42)
	fmt.Println(lat)
	// Output: 22.921026765022624
}

func BenchmarkLatitudeInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LatitudeInRange(-90, 90)
	}
}

func TestLongitudeInRange(t *testing.T) {
	Seed(11)
	long, err := LongitudeInRange(21, 42)
	if err != nil {
		t.Error("error should be nil")
	}

	if long == 0 {
		t.Error("long should be not be zero")
	}

	_, err = LongitudeInRange(-32, -42)
	if err == nil {
		t.Error("error should be not be nil")
	}

	_, err = LongitudeInRange(190, 192)
	if err == nil {
		t.Error("error should be not be nil")
	}
}

func ExampleLongitudeInRange() {
	Seed(11)
	long, _ := LongitudeInRange(-10, 10)
	fmt.Println(long)
	// Output: -8.170450699978453
}

func BenchmarkLongitudeInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongitudeInRange(-180, 180)
	}
}
