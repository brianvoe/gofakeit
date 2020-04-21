package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleGamertag() {
	Seed(11)
	fmt.Println(Gamertag())
	// Output: footinterpret63
}

func BenchmarkGamertag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gamertag()
	}
}
