package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMimeType() {
	Seed(11)
	fmt.Println(MimeType())
	// Output: application/dsptype
}

func BenchmarkMimeType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MimeType()
	}
}

func ExampleExtension() {
	Seed(11)
	fmt.Println(Extension())
	// Output: nes
}

func BenchmarkExtension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Extension()
	}
}
