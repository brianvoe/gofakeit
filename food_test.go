package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFruit() {
	Seed(11)
	fmt.Println(Fruit())
	// Output: Date
}

func BenchmarkFruit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fruit()
	}
}

func ExampleVegetable() {
	Seed(11)
	fmt.Println(Vegetable())
	// Output: Amaranth Leaves
}

func BenchmarkVegetable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Vegetable()
	}
}
