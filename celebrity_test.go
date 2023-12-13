package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCelebrityActor() {
	Seed(11)
	fmt.Println(CelebrityActor())

	// Output: Owen Wilson
}

func ExampleFaker_CelebrityActor() {
	f := New(11)
	fmt.Println(f.CelebrityActor())

	// Output: Owen Wilson
}

func BenchmarkCelebrityActor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CelebrityActor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CelebrityActor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CelebrityActor()
		}
	})
}

func ExampleCelebrityBusiness() {
	Seed(11)
	fmt.Println(CelebrityBusiness())

	// Output: Cameron Diaz
}

func ExampleFaker_CelebrityBusiness() {
	f := New(11)
	fmt.Println(f.CelebrityBusiness())

	// Output: Cameron Diaz
}

func BenchmarkCelebrityBusiness(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CelebrityBusiness()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CelebrityBusiness()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CelebrityBusiness()
		}
	})
}

func ExampleCelebritySport() {
	Seed(11)
	fmt.Println(CelebritySport())

	// Output: Hicham El Guerrouj
}

func ExampleFaker_CelebritySport() {
	f := New(11)
	fmt.Println(f.CelebritySport())

	// Output: Hicham El Guerrouj
}

func BenchmarkCelebritySport(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CelebritySport()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CelebritySport()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CelebritySport()
		}
	})
}
