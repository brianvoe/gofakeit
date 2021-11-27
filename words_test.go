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
	// Output: park
}

func ExampleFaker_Word() {
	f := New(11)
	fmt.Println(f.Word())
	// Output: park
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))
	// Output: Park context record river college.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))
	// Output: Park context record river college.
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
	// Output: Park context record river college purchase appreciate should compare like preserve deal. Combine approach sustain north session explanation race deputy church view efficiency fix. While own act fight roof care god bath link place form specialist. Period sign awareness spread force remember hair wonder provide link right specialist. Cabinet mark teacher gas information attempt flat representation staff iron lift take.
	// Handle food begin look speech trust accept visit go consider chancellor accident. Race intention function sugar exclude person confirm enjoy criticism ring attitude crisis. Complete appear prevent cast mill stop father invite live border aim row. West heart communicate advance compete equipment remain demand report life model cabinet. Chapter use variety realise argue air influence assistance succeed cross artist suffer.
	// Hide forest insist knowledge thatcher executive fashion knock sign proportion expect implement. Work crown restaurant victim material map follow climb attack answer plan order. Place explore telephone operate like display star teach compete weather recommend raise. Democracy college dance recall struggle use cut intention aircraft period generation variety. Benefit defend inquiry conclude justify result depend driver address owner conversation potential.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))
	// Output: Park context record river college purchase appreciate should compare like preserve deal. Combine approach sustain north session explanation race deputy church view efficiency fix. While own act fight roof care god bath link place form specialist. Period sign awareness spread force remember hair wonder provide link right specialist. Cabinet mark teacher gas information attempt flat representation staff iron lift take.
	// Handle food begin look speech trust accept visit go consider chancellor accident. Race intention function sugar exclude person confirm enjoy criticism ring attitude crisis. Complete appear prevent cast mill stop father invite live border aim row. West heart communicate advance compete equipment remain demand report life model cabinet. Chapter use variety realise argue air influence assistance succeed cross artist suffer.
	// Hide forest insist knowledge thatcher executive fashion knock sign proportion expect implement. Work crown restaurant victim material map follow climb attack answer plan order. Place explore telephone operate like display star teach compete weather recommend raise. Democracy college dance recall struggle use cut intention aircraft period generation variety. Benefit defend inquiry conclude justify result depend driver address owner conversation potential.
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
	// Output: Roof chia echo pickled?
}

func ExampleFaker_Question() {
	f := New(11)
	fmt.Println(f.Question())
	// Output: Roof chia echo pickled?
}

func BenchmarkQuestion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Question()
	}
}

func ExampleQuote() {
	Seed(11)
	fmt.Println(Quote())
	// Output: "Roof chia echo pickled." - Marques Jakubowski
}

func ExampleFaker_Quote() {
	f := New(11)
	fmt.Println(f.Quote())
	// Output: "Roof chia echo pickled." - Marques Jakubowski
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
