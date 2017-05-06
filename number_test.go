package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNumber() {
	Seed(11)
	fmt.Println(Number(50, 23456))
	// Output: 14866
}

func BenchmarkNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number(10, 999999)
	}
}

func ExampleNumerify() {
	Seed(11)
	fmt.Println(Numerify("###-###-####"))
	// Output: 328-727-1570
}

func BenchmarkNumerify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Numerify("###-###-####")
	}
}

func ExampleShuffleInts() {
	Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	ShuffleInts(ints)
	fmt.Println(ints)
	// Output: [74125 777 941 89416 8413 854 52 841657]
}

func BenchmarkShuffleInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}
