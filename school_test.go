package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSchool() {
	Seed(11)
	fmt.Println(School())
	// Output:
	// Harborview State Academy
}

func ExampleFaker_School() {
	f := New(11)
	fmt.Println(f.School())
	// Output:
	// Harborview State Academy
}

func BenchmarkExampleFaker_SchoolGen(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			School()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.School()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.School()
		}
	})
}