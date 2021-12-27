package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))
	// Output: Yesterday next worrisome she one.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))
	// Output: Yesterday next worrisome she one.
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
	// Output: Yesterday next worrisome she one must as woman niche our eye paint. Money throughout hair neither journey yet theirs e.g. of calm i.e. her. Tonight busy loosely theirs could delay that some forest not to theirs. Casino our then straightaway herself no cluster herself insufficient us give silly. By below by quarterly his behind enchanted help may that next anybody.
	// Several page gossip necklace am nobody bored accordingly they should can these. Microscope being to reluctantly thing sometimes oxygen than above daily mine a. Brace hug smell vivaciously annually had besides accordingly there these speed lot. Woman loneliness elsewhere park finally wad at most wad conclude thoughtfully stupid. Over one quarterly it growth their is tonight has super terribly now.
	// He recently which troop Bangladeshi ourselves behind youth summation Darwinian being what. Then when this why their stealthily here theirs lot was without crawl. Mine why finally yours how think should as as great annually none. Sammarinese where very bouquet Danish hair ears in no with eye other. Early Swiss buy mob religion later you tonight occasionally in class here.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))
	// Output: Yesterday next worrisome she one must as woman niche our eye paint. Money throughout hair neither journey yet theirs e.g. of calm i.e. her. Tonight busy loosely theirs could delay that some forest not to theirs. Casino our then straightaway herself no cluster herself insufficient us give silly. By below by quarterly his behind enchanted help may that next anybody.
	// Several page gossip necklace am nobody bored accordingly they should can these. Microscope being to reluctantly thing sometimes oxygen than above daily mine a. Brace hug smell vivaciously annually had besides accordingly there these speed lot. Woman loneliness elsewhere park finally wad at most wad conclude thoughtfully stupid. Over one quarterly it growth their is tonight has super terribly now.
	// He recently which troop Bangladeshi ourselves behind youth summation Darwinian being what. Then when this why their stealthily here theirs lot was without crawl. Mine why finally yours how think should as as great annually none. Sammarinese where very bouquet Danish hair ears in no with eye other. Early Swiss buy mob religion later you tonight occasionally in class here.
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
