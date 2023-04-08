package gofakeit

import (
	"fmt"
	"testing"
)

// CUSIP Tests
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
		t.Error("Generated Cusip has invalid checksum")
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

// PPN CUSIP Tests
func ExamplePpnCusip() {
	Seed(11)
	fmt.Println(PpnCusip())
	// Output: CBHG2P*N7
}

func ExampleFaker_PpnCusip() {
	f := New(11)
	fmt.Println(f.PpnCusip())
	// Output: CBHG2P*N7
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

// ISIN Tests
func ExampleIsin() {
	Seed(11)
	fmt.Println(Isin())
	// Output: CVBHG2P1NG14
}

func ExampleFaker_Isin() {
	f := New(11)
	fmt.Println(f.Isin())
	// Output: AMCBHG2P1N52
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
	if IsinCheckDigit(isin[:11]) != string(isin[11]) {
		t.Error("Generated ISIN has invalid check digit")
	}
}

func BenchmarkIsin(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Isin()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Isin()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Isin()
		}
	})
}
