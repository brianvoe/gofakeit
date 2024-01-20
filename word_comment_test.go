package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleComment() {
	Seed(11)
	fmt.Println(Comment())

	// Output: Phew Substantial Thing Had Regularly.
}

func ExampleFaker_Comment() {
	f := New(11)
	fmt.Println(f.Comment())

	// Output: Phew Substantial Thing Had Regularly.
}

func TestComment(t *testing.T) {
	f := New(11)
	for i := 0; i < 1000; i++ {
		f.Comment()
	}
}

func BenchmarkComment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comment()
	}
}
