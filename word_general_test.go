package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleWord() {
	Seed(11)
	fmt.Println(Word())

	// Output: bathe
}

func ExampleFaker_Word() {
	f := New(11)
	fmt.Println(f.Word())

	// Output: bathe
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func TestWord(t *testing.T) {
	for i := 0; i < 10000; i++ {
		if Word() == "" {
			t.Errorf("result should not be blank")
		}
	}
}
