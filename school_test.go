package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSchoolGen() {
	Seed(11)
	fmt.Println(SchoolGen())
	// Output:
	// Harborview State Middle School
}

func ExampleFaker_SchoolGen() {
	f := New(11)
	fmt.Println(f.SchoolGen())
	// Output:
	// Harborview State Middle School
}

func BenchmarkExampleFaker_SchoolGen(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SchoolGen()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.SchoolGen()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.SchoolGen()
		}
	})
}
