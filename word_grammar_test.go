package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSentenceSimple() {
	Seed(11)
	fmt.Println(SentenceSimple())
	// Output: A tribe fly the lemony kitchen.
}

func ExampleFaker_SentenceSimple() {
	f := New(11)
	fmt.Println(f.SentenceSimple())
	// Output: A tribe fly the lemony kitchen.
}

func BenchmarkSentenceSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SentenceSimple()
	}
}
