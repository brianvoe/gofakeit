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

func ExampleBreakfast() {
	Seed(11)
	fmt.Println(Breakfast())
	// Output: Blueberry banana happy face pancakes
}

func BenchmarkBreakfast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Breakfast()
	}
}

func ExampleLunch() {
	Seed(11)
	fmt.Println(Lunch())
	// Output: No bake hersheys bar pie
}

func BenchmarkLunch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lunch()
	}
}

func ExampleDinner() {
	Seed(11)
	fmt.Println(Dinner())
	// Output: Wild addicting dip
}

func BenchmarkDinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dinner()
	}
}

func ExampleSnack() {
	Seed(11)
	fmt.Println(Snack())
	// Output: Hoisin marinated wing pieces
}

func BenchmarkSnack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Snack()
	}
}

func ExampleDessert() {
	Seed(11)
	fmt.Println(Dessert())
	// Output: French napoleons
}

func BenchmarkDessert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dessert()
	}
}
