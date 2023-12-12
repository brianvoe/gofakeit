package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSentenceSimple() {
	Seed(11)
	fmt.Println(SentenceSimple())

	// Output: The purple tribe indeed swiftly laugh.
}

func ExampleFaker_SentenceSimple() {
	f := New(11)
	fmt.Println(f.SentenceSimple())

	// Output: The purple tribe indeed swiftly laugh.
}

func BenchmarkSentenceSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SentenceSimple()
	}
}
