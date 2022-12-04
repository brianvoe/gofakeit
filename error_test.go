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
func ExampleErrorHTTP() {
	Seed(11)
	fmt.Println(ErrorHTTP())
	// Output: http error
}

func ExampleFaker_ErrorHTTP() {
	f := New(11)
	fmt.Println(f.ErrorHTTP())
	// Output: http error
}

func BenchmarkErrorHTTP(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorHTTP()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorHTTP()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorHTTP()
		}
	})
}

func ExampleErrorHTTPClient() {
	Seed(11)
	fmt.Println(ErrorHTTPClient())
	// Output: not found
}

func ExampleFaker_ErrorHTTPClient() {
	f := New(11)
	fmt.Println(f.ErrorHTTPClient())
	// Output: not found
}

func BenchmarkErrorHTTPClient(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorHTTPClient()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorHTTPClient()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorHTTPClient()
		}
	})
}

func ExampleErrorHTTPServer() {
	Seed(11)
	fmt.Println(ErrorHTTPServer())
	// Output: internal server error
}

func ExampleFaker_ErrorHTTPServer() {
	f := New(11)
	fmt.Println(f.ErrorHTTPServer())
	// Output: internal server error
}

func BenchmarkErrorHTTPServer(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorHTTPServer()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorHTTPServer()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorHTTPServer()
		}
	})
}

func ExampleErrorDatabase() {
	Seed(11)
	fmt.Println(ErrorDatabase())
	// Output: bad connection
}

func ExampleFaker_ErrorDatabase() {
	f := New(11)
	fmt.Println(f.ErrorDatabase())
	// Output: bad connection
}

func BenchmarkErrorDatabase(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorDatabase()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorDatabase()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorDatabase()
		}
	})
}
func ExampleErrorGRPC() {
	Seed(11)
	fmt.Println(ErrorGRPC())
	// Output: rpc error
}

func ExampleFaker_ErrorGRPC() {
	f := New(11)
	fmt.Println(f.ErrorGRPC())
	// Output: rpc error
}

func BenchmarkErrorGRPC(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorGRPC()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorGRPC()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorGRPC()
		}
	})
}
func ExampleErrorRuntime() {
	Seed(11)
	fmt.Println(ErrorRuntime())
	// Output: panic: runtime error: invalid memory address or nil pointer dereference
}

func ExampleFaker_ErrorRuntime() {
	f := New(11)
	fmt.Println(f.ErrorRuntime())
	// Output: panic: runtime error: invalid memory address or nil pointer dereference
}

func BenchmarkErrorRuntime(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorRuntime()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorRuntime()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorRuntime()
		}
	})
}
func ExampleErrorInput() {
	Seed(11)
	fmt.Println(ErrorInput())
	// Output: input error
}

func ExampleFaker_ErrorInput() {
	f := New(11)
	fmt.Println(f.ErrorInput())
	// Output: input error
}

func BenchmarkErrorInput(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorInput()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorInput()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorInput()
		}
	})
}
