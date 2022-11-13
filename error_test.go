package gofakeit

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := Error()

	// return value satisfies the error interface
	if fmt.Sprintf("%T", err.Error()) != "string" {
		t.Error("return type must satisfy the error interface")
	}
}

func ExampleError() {
	Seed(11)
	fmt.Println(Error())
	// Output: Validation Error: First name is required
}

func ExampleFaker_Error() {
	f := New(11)
	fmt.Println(f.Error())
	// Output: Validation Error: First name is required
}

func BenchmarkError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Error()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Error()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Error()
		}
	})
}
