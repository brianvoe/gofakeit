package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleGamertag() {
	Seed(11)
	fmt.Println(Gamertag())
	// Output: PurpleSheep5
}

func ExampleFaker_Gamertag() {
	f := New(11)
	fmt.Println(f.Gamertag())
	// Output: PurpleSheep5
}

func TestGamertag(t *testing.T) {
	for i := 0; i < 100; i++ {
		g := Gamertag()
		if g == "" {
			t.Errorf("Gamertag() returned empty string")
		}
	}
}

func BenchmarkGamertag(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Gamertag()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Gamertag()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Gamertag()
		}
	})
}

func ExampleDice() {
	Seed(11)
	fmt.Println(Dice(1, []uint{6}))
	// Output: [6]
}

func ExampleFaker_Dice() {
	f := New(11)
	fmt.Println(f.Dice(1, []uint{6}))
	// Output: [6]
}

func TestDice(t *testing.T) {
	for i := 0; i < 100; i++ {
		// put together random number of dice and sides
		numDice := uint(Number(1, 10))
		sides := make([]uint, numDice)
		for i := 0; i < int(numDice); i++ {
			sides[i] = uint(Number(1, 10))
		}

		g := Dice(numDice, sides)
		if len(g) == 0 {
			t.Errorf("Dice() returned empty uint array")
		}
	}
}

func BenchmarkDice(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Dice(1, []uint{6})
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Dice(1, []uint{6})
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Dice(1, []uint{6})
		}
	})
}
