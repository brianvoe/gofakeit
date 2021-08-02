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
	// Output: 364 East Rapids borough, Rutherfordstad, New Jersey 36906
	// 364 East Rapids borough
	// Rutherfordstad
	// New Jersey
	// 36906
	// Tokelau
	// 23.058758
	// 89.022594
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
	// Output: 364 East Rapids borough, Rutherfordstad, New Jersey 36906
	// 364 East Rapids borough
	// Rutherfordstad
	// New Jersey
	// 36906
	// Tokelau
	// 23.058758
	// 89.022594
}

func BenchmarkAddress(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Address()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Address()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Address()
		}
	})
}

func ExampleStreet() {
	Seed(11)
	fmt.Println(Street())
	// Output: 364 East Rapids borough
}

func ExampleFaker_Street() {
	f := New(11)
	fmt.Println(f.Street())
	// Output: 364 East Rapids borough
}

func BenchmarkStreet(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Street()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Street()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Street()
		}
	})
}

func ExampleStreetNumber() {
	Seed(11)
	fmt.Println(StreetNumber())
	// Output: 13645
}

func ExampleFaker_StreetNumber() {
	f := New(11)
	fmt.Println(f.StreetNumber())
	// Output: 13645
}

func BenchmarkStreetNumber(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StreetNumber()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.StreetNumber()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.StreetNumber()
		}
	})
}

func ExampleStreetPrefix() {
	Seed(11)
	fmt.Println(StreetPrefix())
	// Output: Lake
}

func ExampleFaker_StreetPrefix() {
	f := New(11)
	fmt.Println(f.StreetPrefix())
	// Output: Lake
}

func BenchmarkStreetPrefix(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StreetPrefix()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.StreetPrefix()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.StreetPrefix()
		}
	})
}

func ExampleStreetName() {
	Seed(11)
	fmt.Println(StreetName())
	// Output: View
}

func ExampleFaker_StreetName() {
	f := New(11)
	fmt.Println(f.StreetName())
	// Output: View
}

func BenchmarkStreetName(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StreetName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.StreetName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.StreetName()
		}
	})
}

func ExampleStreetSuffix() {
	Seed(11)
	fmt.Println(StreetSuffix())
	// Output: land
}

func ExampleFaker_StreetSuffix() {
	f := New(11)
	fmt.Println(f.StreetSuffix())
	// Output: land
}

func BenchmarkStreetSuffix(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StreetSuffix()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.StreetSuffix()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.StreetSuffix()
		}
	})
}

func ExampleCity() {
	Seed(11)
	fmt.Println(City())
	// Output: Marcelside
}

func ExampleFaker_City() {
	f := New(11)
	fmt.Println(f.City())
	// Output: Marcelside
}

func TestCity(t *testing.T) {
	for i := 0; i < 100; i++ {
		City()
	}
}

func BenchmarkCity(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			City()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.City()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.City()
		}
	})
}

func ExampleState() {
	Seed(11)
	fmt.Println(State())
	// Output: Hawaii
}

func ExampleFaker_State() {
	f := New(11)
	fmt.Println(f.State())
	// Output: Hawaii
}

func BenchmarkState(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			State()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.State()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.State()
		}
	})
}

func ExampleStateAbr() {
	Seed(11)
	fmt.Println(StateAbr())
	// Output: CO
}

func ExampleFaker_StateAbr() {
	f := New(11)
	fmt.Println(f.StateAbr())
	// Output: CO
}

func BenchmarkStateAbr(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StateAbr()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.StateAbr()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.StateAbr()
		}
	})
}

func ExampleZip() {
	Seed(11)
	fmt.Println(Zip())
	// Output: 13645
}

func ExampleFaker_Zip() {
	f := New(11)
	fmt.Println(f.Zip())
	// Output: 13645
}

func BenchmarkZip(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Zip()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Zip()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Zip()
		}
	})
}

func ExampleCountry() {
	Seed(11)
	fmt.Println(Country())
	// Output: Cabo Verde
}

func ExampleFaker_Country() {
	f := New(11)
	fmt.Println(f.Country())
	// Output: Cabo Verde
}

func BenchmarkCountry(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Country()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Country()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Country()
		}
	})
}

func ExampleCountryAbr() {
	Seed(11)
	fmt.Println(CountryAbr())
	// Output: CV
}

func ExampleFaker_CountryAbr() {
	f := New(11)
	fmt.Println(f.CountryAbr())
	// Output: CV
}

func BenchmarkCountryAbr(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountryAbr()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CountryAbr()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CountryAbr()
		}
	})
}

func ExampleLatitude() {
	Seed(11)
	fmt.Println(Latitude())
	// Output: -73.534057
}

func ExampleFaker_Latitude() {
	f := New(11)
	fmt.Println(f.Latitude())
	// Output: -73.534057
}

func BenchmarkLatitude(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Latitude()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Latitude()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Latitude()
		}
	})
}

func ExampleLongitude() {
	Seed(11)
	fmt.Println(Longitude())
	// Output: -147.068113
}

func ExampleFaker_Longitude() {
	f := New(11)
	fmt.Println(f.Longitude())
	// Output: -147.068113
}

func BenchmarkLongitude(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Longitude()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Longitude()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Longitude()
		}
	})
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
	// Output: 22.921026
}

func ExampleFaker_LatitudeInRange() {
	f := New(11)
	lat, _ := f.LatitudeInRange(21, 42)
	fmt.Println(lat)
	// Output: 22.921026
}

func BenchmarkLatitudeInRange(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LatitudeInRange(-90, 90)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.LatitudeInRange(-90, 90)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.LatitudeInRange(-90, 90)
		}
	})
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
	// Output: -8.170451
}

func ExampleFaker_LongitudeInRange() {
	f := New(11)
	long, _ := f.LongitudeInRange(-10, 10)
	fmt.Println(long)
	// Output: -8.170451
}

func BenchmarkLongitudeInRange(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LongitudeInRange(-180, 180)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.LongitudeInRange(-180, 180)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.LongitudeInRange(-180, 180)
		}
	})
}
