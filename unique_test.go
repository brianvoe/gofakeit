package gofakeit

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	id := UUID()

	if len(id) != 36 {
		t.Error("unique length does not equal requested length")
	}
}

func ExampleUUID() {
	Seed(11)
	fmt.Println(UUID())
	// Output: 590c1440-9888-45b0-bd51-a817ee07c3f2
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}
