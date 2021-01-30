package gofakeit

import (
	"fmt"
	"testing"
	"unicode"
)

func ExampleLetter() {
	Seed(11)
	fmt.Println(Letter())
	// Output: g
}

func ExampleFaker_Letter() {
	f := New(11)
	fmt.Println(f.Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		Letter()
	}
}

func TestLetterN(t *testing.T) {
	if len(LetterN(0)) != 1 {
		t.Errorf("expected length %d but did not get that number", 1)
	}

	var dataSize uint = 10
	data := LetterN(dataSize)
	if len(data) != int(dataSize) {
		t.Errorf("expected length %d but got %d", dataSize, len(data))
	}
	for i := 0; i < len(data); i++ {
		if data[i] > unicode.MaxASCII {
			t.Errorf("non ascii char %s found at index %d", string(data[i]), i)
		}
	}
}

func ExampleLetterN() {
	Seed(11)
	fmt.Println(LetterN(10))
	// Output: gbRMaRxHki
}

func ExampleFaker_LetterN() {
	f := New(11)
	fmt.Println(f.LetterN(10))
	// Output: gbRMaRxHki
}

func BenchmarkLetterN(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		LetterN(10)
	}
}

func ExampleDigit() {
	Seed(11)
	fmt.Println(Digit())
	// Output: 0
}

func ExampleFaker_Digit() {
	f := New(11)
	fmt.Println(f.Digit())
	// Output: 0
}

func BenchmarkDigit(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		Digit()
	}
}

func TestDigitN(t *testing.T) {
	if len(DigitN(0)) != 1 {
		t.Errorf("expected length %d but did not get that number", 1)
	}

	var dataSize uint = 10
	data := DigitN(dataSize)
	if len(data) != int(dataSize) {
		t.Errorf("expected length %d but got %d", dataSize, len(data))
	}
	for i := 0; i < len(data); i++ {
		val := int(data[i] - '0')
		if val < 0 || 9 < val {
			t.Errorf("non digit %s found at index %d", string(data[i]), i)
		}
	}
}

func ExampleDigitN() {
	Seed(11)
	fmt.Println(DigitN(10))
	// Output: 0136459948
}

func ExampleFaker_DigitN() {
	f := New(11)
	fmt.Println(f.DigitN(10))
	// Output: 0136459948
}

func BenchmarkDigitN(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		DigitN(10)
	}
}

func ExampleNumerify() {
	Seed(11)
	fmt.Println(Numerify("###-###-####"))
	// Output: 613-645-9948
}

func ExampleFaker_Numerify() {
	f := New(11)
	fmt.Println(f.Numerify("###-###-####"))
	// Output: 613-645-9948
}

func BenchmarkNumerify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Numerify("###-###-####")
	}
}

func ExampleLexify() {
	Seed(11)
	fmt.Println(Lexify("?????"))
	// Output: gbRMa
}

func ExampleFaker_Lexify() {
	f := New(11)
	fmt.Println(f.Lexify("?????"))
	// Output: gbRMa
}

func BenchmarkLexify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lexify("??????")
	}
}

func ExampleShuffleStrings() {
	Seed(11)
	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	ShuffleStrings(strings)
	fmt.Println(strings)
	// Output: [good everyone have for times a day happy]
}

func ExampleFaker_ShuffleStrings() {
	f := New(11)
	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	f.ShuffleStrings(strings)
	fmt.Println(strings)
	// Output: [good everyone have for times a day happy]
}

func TestShuffleStrings(t *testing.T) {
	ShuffleStrings([]string{"a"})
	ShuffleStrings(nil)

	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := make([]string, len(a))
	copy(b, a)
	ShuffleStrings(a)
	if equalSliceString(a, b) {
		t.Errorf("shuffle resulted in the same permutation, the odds are slim")
	}
}

func BenchmarkShuffleStrings(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		ShuffleStrings([]string{"happy", "times", "for", "everyone", "have", "a", "good", "day"})
	}
}

func ExampleRandomString() {
	Seed(11)
	fmt.Println(RandomString([]string{"hello", "world"}))
	// Output: hello
}

func ExampleFaker_RandomString() {
	f := New(11)
	fmt.Println(f.RandomString([]string{"hello", "world"}))
	// Output: hello
}

func TestRandomString(t *testing.T) {
	for _, test := range []struct {
		in     []string
		should string
	}{
		{[]string{}, ""},
		{nil, ""},
		{[]string{"a"}, "a"},
		{[]string{"a", "b", "c", "d", "e", "f"}, "f"},
	} {
		Seed(44)
		got := RandomString(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%v' should '%s' got '%s'",
			test.in, test.should, got)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString([]string{"hello", "world"})
	}
}
