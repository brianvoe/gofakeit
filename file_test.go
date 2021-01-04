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

func ExampleFaker_FileMimeType() {
	f := New(11)
	fmt.Println(f.FileMimeType())
	// Output: application/dsptype
}

func BenchmarkFileMimeType(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FileMimeType()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.FileMimeType()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.FileMimeType()
		}
	})
}

func ExampleFileExtension() {
	Seed(11)
	fmt.Println(FileExtension())
	// Output: nes
}

func ExampleFaker_FileExtension() {
	f := New(11)
	fmt.Println(f.FileExtension())
	// Output: nes
}

func BenchmarkFileExtension(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FileExtension()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.FileExtension()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.FileExtension()
		}
	})
}
