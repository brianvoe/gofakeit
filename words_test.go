package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())
	// Output: foot
}

func ExampleFaker_Noun() {
	f := New(11)
	fmt.Println(f.Noun())
	// Output: foot
}

func BenchmarkNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Noun()
	}
}

func ExampleVerb() {
	Seed(11)
	fmt.Println(Verb())
	// Output: release
}

func ExampleFaker_Verb() {
	f := New(11)
	fmt.Println(f.Verb())
	// Output: release
}

func BenchmarkVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Verb()
	}
}

func ExampleAdverb() {
	Seed(11)
	fmt.Println(Adverb())
	// Output: smoothly
}

func ExampleFaker_Adverb() {
	f := New(11)
	fmt.Println(f.Adverb())
	// Output: smoothly
}

func BenchmarkAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adverb()
	}
}

func ExamplePreposition() {
	Seed(11)
	fmt.Println(Preposition())
	// Output: down
}

func ExampleFaker_Preposition() {
	f := New(11)
	fmt.Println(f.Preposition())
	// Output: down
}

func BenchmarkPreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Preposition()
	}
}

func ExampleAdjective() {
	Seed(11)
	fmt.Println(Adjective())
	// Output: genuine
}

func ExampleFaker_Adjective() {
	f := New(11)
	fmt.Println(f.Adjective())
	// Output: genuine
}

func BenchmarkAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adjective()
	}
}

func ExampleWord() {
	Seed(11)
	fmt.Println(Word())
	// Output: interpret
}

func ExampleFaker_Word() {
	f := New(11)
	fmt.Println(f.Word())
	// Output: interpret
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))
	// Output: Interpret context record river mind.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))
	// Output: Interpret context record river mind.
}

func TestSentence(t *testing.T) {
	for i := 0; i < 1000; i++ {
		length := len(strings.Split(Sentence(10), " "))
		if length != 10 {
			t.Errorf("result should have a length of 10 got %d", length)
		}
	}

	for _, count := range []int{-100, -1, 0} {
		if Sentence(count) != "" {
			t.Errorf("result should be blank for %d words", count)
		}
	}
}

func BenchmarkSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sentence(10)
	}
}

func ExampleParagraph() {
	Seed(11)
	fmt.Println(Paragraph(3, 5, 12, "\n"))
	// Output: Interpret context record river mind press self should compare property outcome divide. Combine approach sustain consult discover explanation direct address church husband seek army. Begin own act welfare replace press suspect stay link place manchester specialist. Arrive price satisfy sign force application hair train provide basis right pay. Close mark teacher strengthen information attempt head touch aim iron tv take.
	// Handle wait begin look speech trust cancer visit capacity disease chancellor clean. Race aim function gain couple push faith enjoy admit ring attitude develop. Edge game prevent cast mill favour father star live search aim guess. West heart item adopt compete equipment miss output report communicate model cabinet. Seek worker variety step argue air improve give succeed relief artist suffer.
	// Hide finish insist knowledge thatcher make research chance structure proportion husband implement. Town crown restaurant cost material compete lady climb football region discussion order. Place lee market ice like display mind stress compete weather station raise. Democracy college major recall struggle use cut intention accept period generation strike. Benefit defend recommend conclude justify result depend succeed address owner fill interpret.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))
	// Output: Interpret context record river mind press self should compare property outcome divide. Combine approach sustain consult discover explanation direct address church husband seek army. Begin own act welfare replace press suspect stay link place manchester specialist. Arrive price satisfy sign force application hair train provide basis right pay. Close mark teacher strengthen information attempt head touch aim iron tv take.
	// Handle wait begin look speech trust cancer visit capacity disease chancellor clean. Race aim function gain couple push faith enjoy admit ring attitude develop. Edge game prevent cast mill favour father star live search aim guess. West heart item adopt compete equipment miss output report communicate model cabinet. Seek worker variety step argue air improve give succeed relief artist suffer.
	// Hide finish insist knowledge thatcher make research chance structure proportion husband implement. Town crown restaurant cost material compete lady climb football region discussion order. Place lee market ice like display mind stress compete weather station raise. Democracy college major recall struggle use cut intention accept period generation strike. Benefit defend recommend conclude justify result depend succeed address owner fill interpret.
}

func TestParagraph(t *testing.T) {
	for _, count := range []struct{ parag, sent, words int }{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, -100},
		{1, -100, 1},
		{-100, 1, 1},
		{0, 0, 0},
	} {
		if Paragraph(count.parag, count.sent, count.words, " ") != "" {
			t.Errorf("result should be blank for %v input", count)
		}
	}
}

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Paragraph(3, 5, 12, "\n")
	}
}

func ExampleQuestion() {
	Seed(11)
	fmt.Println(Question())
	// Output: Roof chia echo?
}

func ExampleFaker_Question() {
	f := New(11)
	fmt.Println(f.Question())
	// Output: Roof chia echo?
}

func BenchmarkQuestion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Question()
	}
}

func ExampleQuote() {
	Seed(11)
	fmt.Println(Quote())
	// Output: "Roof chia echo." - Lura Lockman
}

func ExampleFaker_Quote() {
	f := New(11)
	fmt.Println(f.Quote())
	// Output: "Roof chia echo." - Lura Lockman
}

func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quote()
	}
}

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())
	// Output: shut the front door
}

func ExampleFaker_Phrase() {
	f := New(11)
	fmt.Println(f.Phrase())
	// Output: shut the front door
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}
