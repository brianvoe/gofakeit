package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFruit() {
	Seed(11)
	fmt.Println(Fruit())

	// Output: Peach
}

func ExampleFaker_Fruit() {
	f := New(11)
	fmt.Println(f.Fruit())

	// Output: Peach
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

func ExampleFaker_Vegetable() {
	f := New(11)
	fmt.Println(f.Vegetable())

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

func ExampleFaker_Breakfast() {
	f := New(11)
	fmt.Println(f.Breakfast())

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

func ExampleFaker_Lunch() {
	f := New(11)
	fmt.Println(f.Lunch())

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

func ExampleFaker_Dinner() {
	f := New(11)
	fmt.Println(f.Dinner())

	// Output: Wild addicting dip
}

func BenchmarkDinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dinner()
	}
}

func ExampleDrink() {
	Seed(11)
	fmt.Println(Drink())

	// Output: Juice
}

func ExampleFaker_Drink() {
	f := New(11)
	fmt.Println(f.Drink())

	// Output: Juice
}

func BenchmarkDrink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Drink()
	}
}

func ExampleSnack() {
	Seed(11)
	fmt.Println(Snack())

	// Output: Hoisin marinated wing pieces
}

func ExampleFaker_Snack() {
	f := New(11)
	fmt.Println(f.Snack())

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

func ExampleFaker_Dessert() {
	f := New(11)
	fmt.Println(f.Dessert())

	// Output: French napoleons
}

func BenchmarkDessert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dessert()
	}
}
