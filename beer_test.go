package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBeerName() {
	Seed(11)
	fmt.Println(BeerName())
	// Output: Duvel
}

func ExampleFaker_BeerName() {
	f := New(11)
	fmt.Println(f.BeerName())
	// Output: Duvel
}

func BenchmarkBeerName(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerName()
		}
	})
}

func ExampleBeerStyle() {
	Seed(11)
	fmt.Println(BeerStyle())
	// Output: European Amber Lager
}

func ExampleFaker_BeerStyle() {
	f := New(11)
	fmt.Println(f.BeerStyle())
	// Output: European Amber Lager
}

func BenchmarkBeerStyle(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerStyle()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerStyle()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerStyle()
		}
	})
}

func ExampleBeerHop() {
	Seed(11)
	fmt.Println(BeerHop())
	// Output: Glacier
}

func ExampleFaker_BeerHop() {
	f := New(11)
	fmt.Println(f.BeerHop())
	// Output: Glacier
}

func BenchmarkBeerHop(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerHop()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerHop()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerHop()
		}
	})
}

func ExampleBeerYeast() {
	Seed(11)
	fmt.Println(BeerYeast())
	// Output: 1388 - Belgian Strong Ale
}

func ExampleFaker_BeerYeast() {
	f := New(11)
	fmt.Println(f.BeerYeast())
	// Output: 1388 - Belgian Strong Ale
}

func BenchmarkBeerYeast(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerYeast()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerYeast()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerYeast()
		}
	})
}

func ExampleBeerMalt() {
	Seed(11)
	fmt.Println(BeerMalt())
	// Output: Munich
}

func ExampleFaker_BeerMalt() {
	f := New(11)
	fmt.Println(f.BeerMalt())
	// Output: Munich
}

func BenchmarkBeerMalt(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerMalt()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerMalt()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerMalt()
		}
	})
}

func ExampleBeerIbu() {
	Seed(11)
	fmt.Println(BeerIbu())
	// Output: 29 IBU
}

func ExampleFaker_BeerIbu() {
	f := New(11)
	fmt.Println(f.BeerIbu())
	// Output: 29 IBU
}

func BenchmarkBeerIbu(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerIbu()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerIbu()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerIbu()
		}
	})
}

func ExampleBeerAlcohol() {
	Seed(11)
	fmt.Println(BeerAlcohol())
	// Output: 2.7%
}

func ExampleFaker_BeerAlcohol() {
	f := New(11)
	fmt.Println(f.BeerAlcohol())
	// Output: 2.7%
}

func BenchmarkBeerAlcohol(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerAlcohol()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerAlcohol()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerAlcohol()
		}
	})
}

func ExampleBeerBlg() {
	Seed(11)
	fmt.Println(BeerBlg())
	// Output: 6.4°Blg
}

func ExampleFaker_BeerBlg() {
	f := New(11)
	fmt.Println(f.BeerBlg())
	// Output: 6.4°Blg
}

func BenchmarkBeerBlg(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BeerBlg()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BeerBlg()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BeerBlg()
		}
	})
}
