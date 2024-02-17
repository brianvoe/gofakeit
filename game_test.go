package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleGamertag() {
	Seed(11)
	fmt.Println(Gamertag())

	// Output: TurkeyThinker
}

func ExampleFaker_Gamertag() {
	f := New(11)
	fmt.Println(f.Gamertag())

	// Output: TurkeyThinker
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
	for i := 0; i < b.N; i++ {
		Gamertag()
	}
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

		// Make sure the length of the array is the same as the number of dice
		if len(g) != int(numDice) {
			t.Errorf("Dice() returned wrong length array")
		}
	}
}

func TestDiceNoSides(t *testing.T) {
	for i := 0; i < 100; i++ {
		g := Dice(1, []uint{})
		if len(g) != 1 {
			t.Errorf("Dice() returned non-empty array")
		}

		// Make sure g[1] is betwwen 1 and 6
		if g[0] < 1 || g[0] > 6 {
			t.Errorf("Dice() returned wrong number")
		}
	}
}

func TestDiceOneSide(t *testing.T) {
	for i := 0; i < 100; i++ {
		g := Dice(10, []uint{1})
		if len(g) != 10 {
			t.Errorf("Dice() returned non 10 value array")
		}

		// Make sure all g values are 1
		for _, v := range g {
			if v != 1 {
				t.Errorf("Dice() returned wrong number")
			}
		}
	}
}

func BenchmarkDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dice(1, []uint{6})
	}
}
