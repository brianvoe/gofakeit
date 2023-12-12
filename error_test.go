package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleError() {
	Seed(11)
	fmt.Println(Error())

	// Output: failed to calculate pointer
}

func ExampleFaker_Error() {
	f := New(11)
	fmt.Println(f.Error())

	// Output: failed to calculate pointer
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

func ExampleErrorObject() {
	Seed(11)
	fmt.Println(ErrorObject())

	// Output: argument
}

func ExampleFaker_ErrorObject() {
	f := New(11)
	fmt.Println(f.ErrorObject())

	// Output: argument
}

func BenchmarkErrorObject(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorObject()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorObject()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorObject()
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

	// Output: connection refused
}

func ExampleFaker_ErrorGRPC() {
	f := New(11)
	fmt.Println(f.ErrorGRPC())

	// Output: connection refused
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

func ExampleErrorHTTP() {
	Seed(11)
	fmt.Println(ErrorHTTP())

	// Output: wrote more than the declared Content-Length
}

func ExampleFaker_ErrorHTTP() {
	f := New(11)
	fmt.Println(f.ErrorHTTP())

	// Output: wrote more than the declared Content-Length
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

	// Output: payment required
}

func ExampleFaker_ErrorHTTPClient() {
	f := New(11)
	fmt.Println(f.ErrorHTTPClient())

	// Output: payment required
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

func ExampleErrorValidation() {
	Seed(11)
	fmt.Println(ErrorValidation())

	// Output: state max length exceeded
}

func ExampleFaker_ErrorValidation() {
	f := New(11)
	fmt.Println(f.ErrorValidation())

	// Output: state max length exceeded
}

func BenchmarkErrorValidation(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ErrorValidation()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ErrorValidation()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ErrorValidation()
		}
	})
}
