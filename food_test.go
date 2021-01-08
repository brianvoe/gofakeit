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

func ExampleFaker_Fruit() {
	f := New(11)
	fmt.Println(f.Fruit())
	// Output: Date
}

func BenchmarkFruit(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fruit()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Fruit()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Fruit()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Vegetable()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Vegetable()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Vegetable()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Breakfast()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Breakfast()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Breakfast()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Lunch()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Lunch()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Lunch()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Dinner()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Dinner()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Dinner()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Snack()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Snack()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Snack()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Dessert()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Dessert()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Dessert()
		}
	})
}
