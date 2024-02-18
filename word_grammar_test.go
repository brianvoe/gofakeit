package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSentenceSimple() {
	Seed(11)
	fmt.Println(SentenceSimple())

	// Output: A fear selfishly cook a tough doctor hardly innocently to realistic project utterly ingeniously.
}

func ExampleFaker_SentenceSimple() {
	f := New(11)
	fmt.Println(f.SentenceSimple())

	// Output: A fear selfishly cook a tough doctor hardly innocently to realistic project utterly ingeniously.
}

func BenchmarkSentenceSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SentenceSimple()
	}
}
