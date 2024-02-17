package gofakeit

import (
	"fmt"
	"testing"
)

// CUSIP Tests
func ExampleCusip() {
	Seed(11)
	fmt.Println(Cusip())

	// Output: 64HHTI0T8
}

func ExampleFaker_Cusip() {
	f := New(11)
	fmt.Println(f.Cusip())

	// Output: 64HHTI0T8
}

func TestCusip(t *testing.T) {
	Seed(11)
	cusip := Cusip()
	if cusip == "" {
		t.Error("Valid Cusips are not blank")
	}
	if len(cusip) != 9 {
		t.Error("Valid Cusips are 9 characters in length")
	}
	if cusipChecksumDigit(cusip[:8]) != string(cusip[8]) {
		t.Error("Generated Cusip has invalid checksum")
	}
}

func TestCusipCheckDigit(t *testing.T) {
	type test struct {
		base string
		want string
	}

	tests := []test{
		{base: "03783310", want: "0"},
		{base: "17275R10", want: "2"},
		{base: "38259P50", want: "8"},
	}
	for _, tc := range tests {
		digit := cusipChecksumDigit(tc.base)
		if digit != tc.want {
			t.Errorf("Expected check digit of %s, got %s", tc.want, digit)
		}
	}
}

func BenchmarkCusip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cusip()
	}
}

// ISIN Tests
func ExampleIsin() {
	Seed(11)
	fmt.Println(Isin())

	// Output: TO4HHTI0T819
}

func ExampleFaker_Isin() {
	f := New(11)
	fmt.Println(f.Isin())

	// Output: TO4HHTI0T819
}

func TestIsin(t *testing.T) {
	Seed(11)
	isin := Isin()
	if isin == "" {
		t.Error("Valid ISINs are not blank")
	}
	if len(isin) != 12 {
		t.Error("Valid ISINs are 12 characters in length")
	}
	if isinChecksumDigit(isin[:11]) != string(isin[11]) {
		t.Error("Generated ISIN has invalid check digit")
	}
}

func TestIsinCheckDigit(t *testing.T) {
	type test struct {
		base string
		want string
	}

	tests := []test{
		{base: "US037833100", want: "5"},
		{base: "GB000263494", want: "6"},
		{base: "US000402625", want: "0"},
	}
	for _, tc := range tests {
		digit := isinChecksumDigit(tc.base)
		if digit != tc.want {
			t.Errorf("Expected check digit of %s, got %s", tc.want, digit)
		}
	}
}

func BenchmarkIsin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Isin()
	}
}
