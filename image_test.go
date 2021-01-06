package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleImageURL() {
	Seed(11)
	fmt.Println(ImageURL(640, 480))
	// Output: https://picsum.photos/640/480
}

func ExampleFaker_ImageURL() {
	f := New(11)
	fmt.Println(f.ImageURL(640, 480))
	// Output: https://picsum.photos/640/480
}

func BenchmarkImageURL(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ImageURL(640, 480)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ImageURL(640, 480)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ImageURL(640, 480)
		}
	})
}
