package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCusip() {
	Seed(11)
	fmt.Println(Cusip())
	// Output: CBHG2P1N5
}

func ExampleFaker_Cusip() {
	f := New(11)
	fmt.Println(f.Cusip())
	// Output: CBHG2P1N5
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
	if CusipCheckDigit(cusip[:8]) != string(cusip[8]) {
		t.Error("Cusip has invalid checksum")
	}

}

func BenchmarkCusip(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Cusip()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Cusip()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Cusip()
		}
	})
}

func ExamplePpnCusip() {
	Seed(11)
	fmt.Println(PpnCusip())
	// Output: 6EHPQ4AK9
}

func ExampleFaker_PpnCusip() {
	f := New(11)
	fmt.Println(f.PpnCusip())
	// Output: 6EHPQ4AK9
}

func BenchmarkPpnCusip(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PpnCusip()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.PpnCusip()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.PpnCusip()
		}
	})
}
