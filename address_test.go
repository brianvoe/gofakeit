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

	// Output: 125 East Routemouth, North Las Vegas, South Dakota 17999
	// 125 East Routemouth
	// North Las Vegas
	// South Dakota
	// 17999
	// Iran (Islamic Republic of)
	// 83.558542
	// -159.896615
}

func ExampleFaker_Address() {
	f := New(11)
	address := f.Address()
	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)

	// Output: 125 East Routemouth, North Las Vegas, South Dakota 17999
	// 125 East Routemouth
	// North Las Vegas
	// South Dakota
	// 17999
	// Iran (Islamic Republic of)
	// 83.558542
	// -159.896615
}

func BenchmarkAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Address()
	}
}

func ExampleStreet() {
	Seed(11)
	fmt.Println(Street())

	// Output: 125 East Routemouth
}

func ExampleFaker_Street() {
	f := New(11)
	fmt.Println(f.Street())

	// Output: 125 East Routemouth
}

func BenchmarkStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Street()
	}
}

func ExampleStreetNumber() {
	Seed(11)
	fmt.Println(StreetNumber())

	// Output: 812
}

func ExampleFaker_StreetNumber() {
	f := New(11)
	fmt.Println(f.StreetNumber())

	// Output: 812
}

func BenchmarkStreetNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetNumber()
	}
}

func ExampleStreetPrefix() {
	Seed(11)
	fmt.Println(StreetPrefix())

	// Output: Port
}

func ExampleFaker_StreetPrefix() {
	f := New(11)
	fmt.Println(f.StreetPrefix())

	// Output: Port
}

func BenchmarkStreetPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetPrefix()
	}
}

func ExampleStreetName() {
	Seed(11)
	fmt.Println(StreetName())

	// Output: Turnpike
}

func ExampleFaker_StreetName() {
	f := New(11)
	fmt.Println(f.StreetName())

	// Output: Turnpike
}

func BenchmarkStreetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetName()
	}
}

func ExampleStreetSuffix() {
	Seed(11)
	fmt.Println(StreetSuffix())

	// Output: side
}

func ExampleFaker_StreetSuffix() {
	f := New(11)
	fmt.Println(f.StreetSuffix())

	// Output: side
}

func BenchmarkStreetSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetSuffix()
	}
}

func ExampleCity() {
	Seed(11)
	fmt.Println(City())

	// Output: Reno
}

func ExampleFaker_City() {
	f := New(11)
	fmt.Println(f.City())

	// Output: Reno
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

	// Output: Vermont
}

func ExampleFaker_State() {
	f := New(11)
	fmt.Println(f.State())

	// Output: Vermont
}

func BenchmarkState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		State()
	}
}

func ExampleStateAbr() {
	Seed(11)
	fmt.Println(StateAbr())

	// Output: WV
}

func ExampleFaker_StateAbr() {
	f := New(11)
	fmt.Println(f.StateAbr())

	// Output: WV
}

func BenchmarkStateAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StateAbr()
	}
}

func ExampleZip() {
	Seed(11)
	fmt.Println(Zip())

	// Output: 81252
}

func ExampleFaker_Zip() {
	f := New(11)
	fmt.Println(f.Zip())

	// Output: 81252
}

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Zip()
	}
}

func ExampleCountry() {
	Seed(11)
	fmt.Println(Country())

	// Output: Tonga
}

func ExampleFaker_Country() {
	f := New(11)
	fmt.Println(f.Country())

	// Output: Tonga
}

func BenchmarkCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Country()
	}
}

func ExampleCountryAbr() {
	Seed(11)
	fmt.Println(CountryAbr())

	// Output: TO
}

func ExampleFaker_CountryAbr() {
	f := New(11)
	fmt.Println(f.CountryAbr())

	// Output: TO
}

func BenchmarkCountryAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountryAbr()
	}
}

func ExampleLatitude() {
	Seed(11)
	fmt.Println(Latitude())

	// Output: 48.654167
}

func ExampleFaker_Latitude() {
	f := New(11)
	fmt.Println(f.Latitude())

	// Output: 48.654167
}

func BenchmarkLatitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Latitude()
	}
}

func ExampleLongitude() {
	Seed(11)
	fmt.Println(Longitude())

	// Output: 97.308335
}

func ExampleFaker_Longitude() {
	f := New(11)
	fmt.Println(f.Longitude())

	// Output: 97.308335
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

	// Output: 37.176319
}

func ExampleFaker_LatitudeInRange() {
	f := New(11)
	lat, _ := f.LatitudeInRange(21, 42)
	fmt.Println(lat)

	// Output: 37.176319
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

	// Output: 5.406018
}

func ExampleFaker_LongitudeInRange() {
	f := New(11)
	long, _ := f.LongitudeInRange(-10, 10)
	fmt.Println(long)

	// Output: 5.406018
}

func BenchmarkLongitudeInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongitudeInRange(-180, 180)
	}
}
