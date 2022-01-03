package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSentenceSimple() {
	Seed(11)
	fmt.Println(SentenceSimple())
	// Output: A purple tribe fly a buckles.
}

func ExampleFaker_SentenceSimple() {
	f := New(11)
	fmt.Println(f.SentenceSimple())
	// Output: A purple tribe fly a buckles.
}

func BenchmarkSentenceSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SentenceSimple()
	}
}
