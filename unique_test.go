package gofakeit

import (
	"fmt"
	"testing"
)

func TestUniqueID(t *testing.T) {
	id := UniqueID()

	if len(id) != 32 {
		t.Error("unique length does not equal requested length")
	}
}

func ExampleUniqueID() {
	Seed(11)
	fmt.Println(UniqueID())
	// Output: Q3LewXJfQuPrnLiduMnTnQcsGJ1NqBmv
}

func BenchmarkUniqueID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniqueID()
	}
}
