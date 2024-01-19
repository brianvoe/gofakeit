package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))

	// Output: None how these keep trip.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))

	// Output: None how these keep trip.
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

	// Output: None how these keep trip Congolese choir computer still far unless army. Party riches theirs instead here mine whichever that those instance growth has. Ouch enough Swiss us since down he she aha us you to. Upon how this this furniture way no play towel that us to. Accordingly theirs purse enough so though election often until eek weekly yet.
	// Oops until conclude his stay elsewhere such that galaxy clean that last. Each stack reluctantly theirs wisp aid firstly highly butter accordingly should already. Calm shake according fade neither kuban upon this he fortnightly occasionally bunch. On who elsewhere lastly hourly right there honesty we is how result. Out Shakespearean have whom yearly another packet are behind late lot finally.
	// Time themselves goodness quizzical our therefore could fact cackle yourselves zebra for. Whose enormously all bowl out wandering secondly yellow another your hourly spit. Since tomorrow hers words little think will our by Polynesian write much. Of herself uptight these composer these any firstly stack you much terribly. Over pose place sprint it child is joyously that I whom mango.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))

	// Output: None how these keep trip Congolese choir computer still far unless army. Party riches theirs instead here mine whichever that those instance growth has. Ouch enough Swiss us since down he she aha us you to. Upon how this this furniture way no play towel that us to. Accordingly theirs purse enough so though election often until eek weekly yet.
	// Oops until conclude his stay elsewhere such that galaxy clean that last. Each stack reluctantly theirs wisp aid firstly highly butter accordingly should already. Calm shake according fade neither kuban upon this he fortnightly occasionally bunch. On who elsewhere lastly hourly right there honesty we is how result. Out Shakespearean have whom yearly another packet are behind late lot finally.
	// Time themselves goodness quizzical our therefore could fact cackle yourselves zebra for. Whose enormously all bowl out wandering secondly yellow another your hourly spit. Since tomorrow hers words little think will our by Polynesian write much. Of herself uptight these composer these any firstly stack you much terribly. Over pose place sprint it child is joyously that I whom mango.
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
