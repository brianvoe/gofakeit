package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMap() {
	Seed(11)
	fmt.Println(Map())
	// Output: map[eaque:784604.3 excepturi:Christy Ratke quo:99536 North Streamville, Rossieview, Hawaii 42591 repellat:776037 voluptatem:57704.402]
}

func TestMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		Map()
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map()
	}
}
