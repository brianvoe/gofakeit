package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleGamertag() {
	Seed(11)
	fmt.Println(Gamertag())
	// Output: footinterpret63
}

func ExampleFaker_Gamertag() {
	f := New(11)
	fmt.Println(f.Gamertag())
	// Output: footinterpret63
}

func BenchmarkGamertag(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Gamertag()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Gamertag()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Gamertag()
		}
	})
}
