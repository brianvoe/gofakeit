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
	// Output: failed to bundle method
}

func ExampleFaker_Error() {
	f := New(11)
	fmt.Println(f.Error())
	// Output: failed to bundle method
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
func ExampleHTTPError() {
	Seed(11)
	fmt.Println(HTTPError())
	// Output: http error
}

func ExampleFaker_HTTPError() {
	f := New(11)
	fmt.Println(f.HTTPError())
	// Output: http error
}

func BenchmarkHTTPError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPError()
		}
	})
}

func ExampleHTTPClientError() {
	Seed(11)
	fmt.Println(HTTPClientError())
	// Output: not found
}

func ExampleFaker_HTTPClientError() {
	f := New(11)
	fmt.Println(f.HTTPClientError())
	// Output: not found
}

func BenchmarkHTTPClientError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPClientError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPClientError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPClientError()
		}
	})
}

func ExampleHTTPServerError() {
	Seed(11)
	fmt.Println(HTTPServerError())
	// Output: internal server error
}

func ExampleFaker_HTTPServerError() {
	f := New(11)
	fmt.Println(f.HTTPServerError())
	// Output: internal server error
}

func BenchmarkHTTPServerError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPServerError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPServerError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPServerError()
		}
	})
}

func ExampleDatabaseError() {
	Seed(11)
	fmt.Println(DatabaseError())
	// Output: bad connection
}

func ExampleFaker_DatabaseError() {
	f := New(11)
	fmt.Println(f.DatabaseError())
	// Output: bad connection
}

func BenchmarkDatabaseError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DatabaseError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.DatabaseError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.DatabaseError()
		}
	})
}
func ExampleGRPCError() {
	Seed(11)
	fmt.Println(GRPCError())
	// Output: rpc error
}

func ExampleFaker_GRPCError() {
	f := New(11)
	fmt.Println(f.GRPCError())
	// Output: rpc error
}

func BenchmarkGRPCError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GRPCError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.GRPCError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.GRPCError()
		}
	})
}
func ExampleRuntimeError() {
	Seed(11)
	fmt.Println(RuntimeError())
	// Output: panic: runtime error: invalid memory address or nil pointer dereference
}

func ExampleFaker_RuntimeError() {
	f := New(11)
	fmt.Println(f.RuntimeError())
	// Output: panic: runtime error: invalid memory address or nil pointer dereference
}

func BenchmarkRuntimeError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RuntimeError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.RuntimeError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.RuntimeError()
		}
	})
}
func ExampleInputError() {
	Seed(11)
	fmt.Println(InputError())
	// Output: input error
}

func ExampleFaker_InputError() {
	f := New(11)
	fmt.Println(f.InputError())
	// Output: input error
}

func BenchmarkInputError(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			InputError()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.InputError()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.InputError()
		}
	})
}
