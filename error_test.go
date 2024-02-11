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
	for i := 0; i < b.N; i++ {
		Error()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorObject()
	}
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

	// Output: payment required
}

func ExampleFaker_ErrorHTTPClient() {
	f := New(11)
	fmt.Println(f.ErrorHTTPClient())

	// Output: payment required
}

func BenchmarkErrorHTTPClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ErrorHTTPClient()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorHTTPServer()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorRuntime()
	}
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
	for i := 0; i < b.N; i++ {
		ErrorValidation()
	}
}
