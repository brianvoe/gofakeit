package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFileMimeType() {
	Seed(11)
	fmt.Println(FileMimeType())

	// Output: application/x-wri
}

func ExampleFaker_FileMimeType() {
	f := New(11)
	fmt.Println(f.FileMimeType())

	// Output: application/x-wri
}

func BenchmarkFileMimeType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileMimeType()
	}
}

func ExampleFileExtension() {
	Seed(11)
	fmt.Println(FileExtension())

	// Output: dtd
}

func ExampleFaker_FileExtension() {
	f := New(11)
	fmt.Println(f.FileExtension())

	// Output: dtd
}

func BenchmarkFileExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileExtension()
	}
}
