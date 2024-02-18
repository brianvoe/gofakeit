package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))

	// Output: Bathe regularly quiver these sprint.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence(5))

	// Output: Bathe regularly quiver these sprint.
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

	// Output: Bathe regularly quiver these sprint fight something am elsewhere since previously was. That there a tennis occur why depend heels out can fire anyone. Sometimes that leap whom troop now scarcely since dance huh somebody constantly. Italian outfit certain weekly murder who cackle because her break write nobody. Could according whoa bevy we them summation to oxygen Greek how previously.
	// A what answer handle daily ouch upon encourage that after his every. She wearily still he herself huge failure for yay sparse thoughtfully box. Anyway Jungian regularly world great daily recently whose Elizabethan outside to theirs. Block while French closely sit you to out fight bored yours those. Bale Parisian host which so today Ecuadorian hourly decidedly his everybody this.
	// Play this little them on recently eager usually bunch daily whose somebody. Them few tonight dishonesty spell battery less patience whose inside hair your. Answer hand tonight these she team which today Einsteinian country yesterday for. Phew for now your him hourly infrequently lots drink absolutely when east. Conclude little annually hers nobody since troop it it whoever the park.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph(3, 5, 12, "\n"))

	// Output: Bathe regularly quiver these sprint fight something am elsewhere since previously was. That there a tennis occur why depend heels out can fire anyone. Sometimes that leap whom troop now scarcely since dance huh somebody constantly. Italian outfit certain weekly murder who cackle because her break write nobody. Could according whoa bevy we them summation to oxygen Greek how previously.
	// A what answer handle daily ouch upon encourage that after his every. She wearily still he herself huge failure for yay sparse thoughtfully box. Anyway Jungian regularly world great daily recently whose Elizabethan outside to theirs. Block while French closely sit you to out fight bored yours those. Bale Parisian host which so today Ecuadorian hourly decidedly his everybody this.
	// Play this little them on recently eager usually bunch daily whose somebody. Them few tonight dishonesty spell battery less patience whose inside hair your. Answer hand tonight these she team which today Einsteinian country yesterday for. Phew for now your him hourly infrequently lots drink absolutely when east. Conclude little annually hers nobody since troop it it whoever the park.
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

	// Output: Everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro?
}

func ExampleFaker_Question() {
	f := New(11)
	fmt.Println(f.Question())

	// Output: Everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro?
}

func BenchmarkQuestion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Question()
	}
}

func ExampleQuote() {
	Seed(11)
	fmt.Println(Quote())

	// Output: "Everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro." - Vance Schroeder
}

func ExampleFaker_Quote() {
	f := New(11)
	fmt.Println(f.Quote())

	// Output: "Everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro." - Vance Schroeder
}

func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quote()
	}
}
