package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMimeType() {
	Seed(11)
	fmt.Println(MimeType())
	// Output: video/3gpp2
}

func BenchmarkMimeType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MimeType()
	}
}
