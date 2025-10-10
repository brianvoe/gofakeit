package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleAirlineAircraftType() {
	Seed(11)
	fmt.Println(AirlineAircraftType())

	// Output: regional
}

func ExampleFaker_AirlineAircraftType() {
	f := New(11)
	fmt.Println(f.AirlineAircraftType())

	// Output: regional
}

func BenchmarkAirlineAircraftType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineAircraftType()
	}
}

func ExampleAirlineAirplane() {
	Seed(11)
	fmt.Println(AirlineAirplane())

	// Output: De Havilland Dash 8
}

func ExampleFaker_AirlineAirplane() {
	f := New(11)
	fmt.Println(f.AirlineAirplane())

	// Output: De Havilland Dash 8
}

func BenchmarkAirlineAirplane(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineAirplane()
	}
}

func ExampleAirlineAirport() {
	Seed(11)
	fmt.Println(AirlineAirport())

	// Output: Cairo International Airport
}

func ExampleFaker_AirlineAirport() {
	f := New(11)
	fmt.Println(f.AirlineAirport())

	// Output: Cairo International Airport
}

func BenchmarkAirlineAirport(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineAirport()
	}
}

func ExampleAirlineAirportIATA() {
	Seed(11)
	fmt.Println(AirlineAirportIATA())

	// Output: CAI
}

func ExampleFaker_AirlineAirportIATA() {
	f := New(11)
	fmt.Println(f.AirlineAirportIATA())

	// Output: CAI
}

func BenchmarkAirlineAirportIATA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineAirportIATA()
	}
}

func ExampleAirlineFlightNumber() {
	Seed(11)
	fmt.Println(AirlineFlightNumber())

	// Output: US1996
}

func ExampleFaker_AirlineFlightNumber() {
	f := New(11)
	fmt.Println(f.AirlineFlightNumber())

	// Output: US1996
}

func BenchmarkAirlineFlightNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineFlightNumber()
	}
}

func TestAirlineFlightNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		flightNum := AirlineFlightNumber()

		// Should be 3-6 characters (2 letters + 1-4 digits)
		if len(flightNum) < 3 || len(flightNum) > 6 {
			t.Errorf("Flight number length should be 3-6, got %d: %s", len(flightNum), flightNum)
		}

		// First 2 characters should be letters
		if !isLetter(rune(flightNum[0])) || !isLetter(rune(flightNum[1])) {
			t.Errorf("First 2 characters should be letters: %s", flightNum)
		}
	}
}

func ExampleAirlineRecordLocator() {
	Seed(11)
	fmt.Println(AirlineRecordLocator())

	// Output: USKKBM
}

func ExampleFaker_AirlineRecordLocator() {
	f := New(11)
	fmt.Println(f.AirlineRecordLocator())

	// Output: USKKBM
}

func BenchmarkAirlineRecordLocator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineRecordLocator()
	}
}

func TestAirlineRecordLocator(t *testing.T) {
	for i := 0; i < 100; i++ {
		locator := AirlineRecordLocator()

		// Should be exactly 6 characters
		if len(locator) != 6 {
			t.Errorf("Record locator should be 6 characters, got %d: %s", len(locator), locator)
		}

		// All characters should be uppercase letters
		for _, c := range locator {
			if !isLetter(c) || c < 'A' || c > 'Z' {
				t.Errorf("Record locator should only contain uppercase letters: %s", locator)
				break
			}
		}
	}
}

func ExampleAirlineSeat() {
	Seed(11)
	fmt.Println(AirlineSeat())

	// Output: 54J
}

func ExampleFaker_AirlineSeat() {
	f := New(11)
	fmt.Println(f.AirlineSeat())

	// Output: 54J
}

func BenchmarkAirlineSeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AirlineSeat()
	}
}

func TestAirlineSeat(t *testing.T) {
	for i := 0; i < 100; i++ {
		seat := AirlineSeat()

		// Should be 2-3 characters (row number + letter)
		if len(seat) < 2 || len(seat) > 3 {
			t.Errorf("Seat should be 2-3 characters, got %d: %s", len(seat), seat)
		}

		// Last character should be a letter (A-K, excluding I)
		lastChar := seat[len(seat)-1]
		validSeats := "ABCDEFGHJK"
		if !strings.ContainsRune(validSeats, rune(lastChar)) {
			t.Errorf("Seat letter should be A-K (excluding I), got: %c in %s", lastChar, seat)
		}
	}
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
