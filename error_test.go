package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleError() {
	Seed(11)
	fmt.Println(Error())

	// Output: variable assigned before declaration
}

func ExampleFaker_Error() {
	f := New(11)
	fmt.Println(f.Error())

	// Output: variable assigned before declaration
}

func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Error()
	}
}

func ExampleErrorObject() {
	Seed(11)
	fmt.Println(ErrorObject())

	// Output: url
}

func ExampleFaker_ErrorObject() {
	f := New(11)
	fmt.Println(f.ErrorObject())

	// Output: url
}

func BenchmarkErrorObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorObject()
	}
}

func ExampleErrorDatabase() {
	Seed(11)
	fmt.Println(ErrorDatabase())

	// Output: destination pointer is nil
}

func ExampleFaker_ErrorDatabase() {
	f := New(11)
	fmt.Println(f.ErrorDatabase())

	// Output: destination pointer is nil
}

func BenchmarkErrorDatabase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorDatabase()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorGRPC()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorHTTP()
	}
}

func ExampleErrorHTTPClient() {
	Seed(11)
	fmt.Println(ErrorHTTPClient())

	// Output: expectation failed
}

func ExampleFaker_ErrorHTTPClient() {
	f := New(11)
	fmt.Println(f.ErrorHTTPClient())

	// Output: expectation failed
}

func BenchmarkErrorHTTPClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorHTTPClient()
	}
}

func ExampleErrorHTTPServer() {
	Seed(11)
	fmt.Println(ErrorHTTPServer())

	// Output: not extended
}

func ExampleFaker_ErrorHTTPServer() {
	f := New(11)
	fmt.Println(f.ErrorHTTPServer())

	// Output: not extended
}

func BenchmarkErrorHTTPServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorHTTPServer()
	}
}

func ExampleErrorRuntime() {
	Seed(11)
	fmt.Println(ErrorRuntime())

	// Output: expected 2 arguments, got 3
}

func ExampleFaker_ErrorRuntime() {
	f := New(11)
	fmt.Println(f.ErrorRuntime())

	// Output: expected 2 arguments, got 3
}

func BenchmarkErrorRuntime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorRuntime()
	}
}

func ExampleErrorValidation() {
	Seed(11)
	fmt.Println(ErrorValidation())

	// Output: payment details cannot be verified
}

func ExampleFaker_ErrorValidation() {
	f := New(11)
	fmt.Println(f.ErrorValidation())

	// Output: payment details cannot be verified
}

func BenchmarkErrorValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorValidation()
	}
}
