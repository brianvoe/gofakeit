package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleFruit() {
	Seed(11)
	fmt.Println(Fruit())

	// Output: Redcurrant
}

func ExampleFaker_Fruit() {
	f := New(11)
	fmt.Println(f.Fruit())

	// Output: Redcurrant
}

func BenchmarkFruit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fruit()
	}
}

func ExampleVegetable() {
	Seed(11)
	fmt.Println(Vegetable())

	// Output: Sweet Potato
}

func ExampleFaker_Vegetable() {
	f := New(11)
	fmt.Println(f.Vegetable())

	// Output: Sweet Potato
}

func BenchmarkVegetable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Vegetable()
	}
}

func ExampleBreakfast() {
	Seed(11)
	fmt.Println(Breakfast())

	// Output: Purple cow
}

func ExampleFaker_Breakfast() {
	f := New(11)
	fmt.Println(f.Breakfast())

	// Output: Purple cow
}

func BenchmarkBreakfast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Breakfast()
	}
}

func ExampleLunch() {
	Seed(11)
	fmt.Println(Lunch())

	// Output: Quick chile relleno casserole
}

func ExampleFaker_Lunch() {
	f := New(11)
	fmt.Println(f.Lunch())

	// Output: Quick chile relleno casserole
}

func BenchmarkLunch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lunch()
	}
}

func ExampleDinner() {
	Seed(11)
	fmt.Println(Dinner())

	// Output: German apple cake with cream cheese frosting
}

func ExampleFaker_Dinner() {
	f := New(11)
	fmt.Println(f.Dinner())

	// Output: German apple cake with cream cheese frosting
}

func BenchmarkDinner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dinner()
	}
}

func ExampleDrink() {
	Seed(11)
	fmt.Println(Drink())

	// Output: Wine
}

func ExampleFaker_Drink() {
	f := New(11)
	fmt.Println(f.Drink())

	// Output: Wine
}

func BenchmarkDrink(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Drink()
	}
}

func ExampleSnack() {
	Seed(11)
	fmt.Println(Snack())

	// Output: Fantastic banana bran muffins
}

func ExampleFaker_Snack() {
	f := New(11)
	fmt.Println(f.Snack())

	// Output: Fantastic banana bran muffins
}

func BenchmarkSnack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Snack()
	}
}

func ExampleDessert() {
	Seed(11)
	fmt.Println(Dessert())

	// Output: Amish cream pie
}

func ExampleFaker_Dessert() {
	f := New(11)
	fmt.Println(f.Dessert())

	// Output: Amish cream pie
}

func BenchmarkDessert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dessert()
	}
}
