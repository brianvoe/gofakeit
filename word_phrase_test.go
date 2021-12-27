package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())
	// Output: horses for courses
}

func ExampleFaker_Phrase() {
	f := New(11)
	fmt.Println(f.Phrase())
	// Output: horses for courses
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}
