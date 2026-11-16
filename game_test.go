package gofakeit

import (
	"fmt"
	"regexp"
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
	for range 100 {
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
	for range 100 {
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
	for range 100 {
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
	for range 100 {
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

func TestDiceLookupDefaultSides(t *testing.T) {
	// Only the number of dice is given, the sides fall back to their default of [6]
	info := GetFuncLookup("dice")
	m := NewMapParams()
	m.Add("numdice", "3")

	for range 100 {
		value, err := info.Generate(GlobalFaker, m, info)
		if err != nil {
			t.Fatalf("dice with only numdice returned an error: %v", err)
		}

		g := value.([]uint)
		if len(g) != 3 {
			t.Fatalf("expected 3 dice, got %v", g)
		}
		for _, v := range g {
			if v < 1 || v > 6 {
				t.Fatalf("expected dice with 6 sides, got %v", g)
			}
		}
	}

	// The same through a template and a struct tag, the list is printed like [4 5 5]
	rolls := regexp.MustCompile(`^\[[1-6] [1-6] [1-6]\]$`)

	got, err := Generate("{dice:3}")
	if err != nil {
		t.Fatalf("{dice:3} returned an error: %v", err)
	}
	if !rolls.MatchString(got) {
		t.Fatalf("expected 3 dice with 6 sides, got %q", got)
	}

	type roll struct {
		Dice string `fake:"{dice:3}"`
	}
	var r roll
	if err := Struct(&r); err != nil {
		t.Fatalf("struct tag {dice:3} returned an error: %v", err)
	}
	if !rolls.MatchString(r.Dice) {
		t.Fatalf("expected 3 dice with 6 sides, got %q", r.Dice)
	}
}

func BenchmarkDice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dice(1, []uint{6})
	}
}
