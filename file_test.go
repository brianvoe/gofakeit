package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFileMimeType() {
	Seed(11)
	fmt.Println(FileMimeType())
	// Output: application/dsptype
}

func BenchmarkFileMimeType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileMimeType()
	}
}

func ExampleFileExtension() {
	Seed(11)
	fmt.Println(FileExtension())
	// Output: nes
}

func BenchmarkFileExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FileExtension()
	}
}
