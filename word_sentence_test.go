package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))

	// Output: There had result everything niche.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))

	// Output: There had result everything niche.
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

	// Output: There had result everything niche backwards caused quarterly without week it hungry. Thing someone him regularly today whomever this revolt hence from his timing. As quantity us these yours live these frantic not may another how. This ours his them those whose them batch its Iraqi most that. Few abroad cheese this whereas next how there gorgeous genetics time choir.
	// Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow. Where several his orchard to frequently hence victorious boxers each does auspicious. Yourselves first soup tomorrow this that must conclude anyway some yearly who. Cough laugh himself both yet rarely me dolphin intensely block would leap. Plane us first then down them eager would hundred super throughout animal.
	// Yet themselves been group flock shake part purchase up usually it her. None it hers boat what their there Turkmen moreover one Lebanese to. Brace these shower in it everybody should whatever those uninterested several nobody. Yours thankful innocent power to any from its few luxury none boy. Religion themselves it there might must near theirs mine thing tonight outside.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))

	// Output: There had result everything niche backwards caused quarterly without week it hungry. Thing someone him regularly today whomever this revolt hence from his timing. As quantity us these yours live these frantic not may another how. This ours his them those whose them batch its Iraqi most that. Few abroad cheese this whereas next how there gorgeous genetics time choir.
	// Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow. Where several his orchard to frequently hence victorious boxers each does auspicious. Yourselves first soup tomorrow this that must conclude anyway some yearly who. Cough laugh himself both yet rarely me dolphin intensely block would leap. Plane us first then down them eager would hundred super throughout animal.
	// Yet themselves been group flock shake part purchase up usually it her. None it hers boat what their there Turkmen moreover one Lebanese to. Brace these shower in it everybody should whatever those uninterested several nobody. Yours thankful innocent power to any from its few luxury none boy. Religion themselves it there might must near theirs mine thing tonight outside.
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
