package gofakeit

import (
	"fmt"
	"testing"
)

func TestUniqueId(t *testing.T) {
	id := UniqueId()

	if len(id) != 32 {
		t.Error("unique length does not equal requested length")
	}
}

func ExampleUniqueId() {
	Seed(11)
	fmt.Println(UniqueId())
	// Output: Q3LewXJfQuPrnLiduMnTnQcsGJ1NqBmv
}

func BenchmarkUniqueId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueId()
	}
}
