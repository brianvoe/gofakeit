package gofakeit

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func ExampleLetter() {
	Seed(11)
	fmt.Println(Letter())
	// Output: g
}

func BenchmarkLetter(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		Letter()
	}
}

func ExampleDigit() {
	Seed(11)
	fmt.Println(Digit())
	// Output: 0
}

func BenchmarkDigit(b *testing.B) {
	Seed(11)
	for i := 0; i < b.N; i++ {
		Digit()
	}
}

func ExampleNumerify() {
	Seed(11)
	fmt.Println(Numerify("###-###-####"))
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

func ExampleRandomChars() {
	Seed(1)
	fmt.Println(RandomChars(5, &LangENG))
	fmt.Println(RandomChars(5, &LangCHI))
	fmt.Println(RandomChars(5, &LangRUS))
	// Output:
	// FbpXH
	// 鷣頦苷峑榫
	// зРЃвЉ
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

func TestRandomChars(t *testing.T) {
	var err error
	var engStr = RandomChars(5, &LangENG)
	var chiStr = RandomChars(5, &LangCHI)
	var rusStr = RandomChars(5, &LangRUS)
	err = isStringLangCorrect(engStr, LangENG)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(chiStr, LangCHI)
	if err != nil {
		t.Error(err.Error())
	}
	err = isStringLangCorrect(rusStr, LangRUS)
	if err != nil {
		t.Error(err.Error())
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString([]string{"hello", "world"})
	}
}

func isStringLangCorrect(value string, lang langRuneBoundary) error {
	for i := 0; i < len(value); {
		r, size := utf8.DecodeLastRuneInString(value[i:])
		if r < lang.start || r > lang.end {
			return fmt.Errorf("Symbol is not in selected alphabet: start: %d, end: %d", lang.start, lang.end)
		}
		i += size
	}
	return nil
}
