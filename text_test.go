package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleComment() {
	Seed(11)
	fmt.Println(Comment())

	// Output: Fear Drink To Heart.
}

func ExampleFaker_Comment() {
	f := New(11)
	fmt.Println(f.Comment())

	// Output: Fear Drink To Heart.
}

func TestComment(t *testing.T) {
	f := New(11)
	for i := 0; i < 1000; i++ {
		f.Comment()
	}
}

func BenchmarkComment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comment()
	}
}

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())

	// Output: how many siblings do you have
}

func ExampleFaker_Phrase() {
	f := New(11)
	fmt.Println(f.Phrase())

	// Output: how many siblings do you have
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}

func ExamplePhraseNoun() {
	Seed(11)
	fmt.Println(PhraseNoun())

	// Output: a fear
}

func ExampleFaker_PhraseNoun() {
	f := New(11)
	fmt.Println(f.PhraseNoun())

	// Output: a fear
}

func BenchmarkPhraseNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseNoun()
	}
}

func ExamplePhraseVerb() {
	Seed(11)
	fmt.Println(PhraseVerb())

	// Output: bathe the jittery trip totally brightly under a troubling part scarcely unexpectedly
}

func ExampleFaker_PhraseVerb() {
	f := New(11)
	fmt.Println(f.PhraseVerb())

	// Output: bathe the jittery trip totally brightly under a troubling part scarcely unexpectedly
}

func BenchmarkPhraseVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseVerb()
	}
}

func ExamplePhraseAdverb() {
	Seed(11)
	fmt.Println(PhraseAdverb())

	// Output: successfully
}

func ExampleFaker_PhraseAdverb() {
	f := New(11)
	fmt.Println(f.PhraseAdverb())

	// Output: successfully
}

func BenchmarkPhraseAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhraseAdverb()
	}
}

func ExamplePhrasePreposition() {
	Seed(11)
	fmt.Println(PhrasePreposition())

	// Output: with an archipelago
}

func ExampleFaker_PhrasePreposition() {
	f := New(11)
	fmt.Println(f.PhrasePreposition())

	// Output: with an archipelago
}

func BenchmarkPhrasePreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhrasePreposition()
	}
}

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

func ExampleLoremIpsumSentence() {
	Seed(11)
	fmt.Println(LoremIpsumSentence(5))

	// Output: Eveniet vero velit velit non.
}

func ExampleFaker_LoremIpsumSentence() {
	f := New(11)
	fmt.Println(f.LoremIpsumSentence(5))

	// Output: Eveniet vero velit velit non.
}

func TestLoremIpsumSentence(t *testing.T) {
	for _, count := range []int{-100, -1, 0} {
		if LoremIpsumSentence(count) != "" {
			t.Errorf("result should be blank for %d words", count)
		}
	}
}

func BenchmarkLoremIpsumSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsumSentence(10)
	}
}

func ExampleLoremIpsumParagraph() {
	Seed(11)
	fmt.Println(LoremIpsumParagraph(3, 5, 12, "\n"))

	// Output: Eveniet vero velit velit non incidunt est sed tenetur consequatur ut architecto. Quod non non et voluptas et esse quisquam ullam placeat molestiae laboriosam. Consequatur in perferendis molestiae quam voluptatem eaque quia facilis quo hic dignissimos. Architecto ut tenetur aut qui et optio id sint quae rerum labore. Quae deserunt quia mollitia id accusantium itaque nisi aut omnis occaecati repellendus.
	// Ad sunt neque placeat et qui saepe voluptatem blanditiis voluptatum possimus saepe. Velit distinctio quia quas cupiditate sequi dolorum delectus quia quo corrupti labore. Aut mollitia et aut eligendi blanditiis corrupti aut rem voluptatibus veritatis ut. Et ab voluptas delectus repellat eligendi qui eos blanditiis et et eos. Qui ipsum aut ut voluptate nulla officia qui earum repudiandae quidem officia.
	// Fugit esse beatae placeat nihil libero et sunt voluptas velit maiores voluptatem. Reiciendis quia dolor amet blanditiis labore eligendi fugit reiciendis modi est incidunt. Ut ut rerum odit sit voluptatem ipsam consequatur est incidunt aut quis. Non dolorem voluptatum pariatur esse eos blanditiis rem eos commodi accusamus occaecati. Qui voluptatem eos et tempora et id at voluptas suscipit magni cupiditate.
}

func ExampleFaker_LoremIpsumParagraph() {
	f := New(11)
	fmt.Println(f.LoremIpsumParagraph(3, 5, 12, "\n"))

	// Output: Eveniet vero velit velit non incidunt est sed tenetur consequatur ut architecto. Quod non non et voluptas et esse quisquam ullam placeat molestiae laboriosam. Consequatur in perferendis molestiae quam voluptatem eaque quia facilis quo hic dignissimos. Architecto ut tenetur aut qui et optio id sint quae rerum labore. Quae deserunt quia mollitia id accusantium itaque nisi aut omnis occaecati repellendus.
	// Ad sunt neque placeat et qui saepe voluptatem blanditiis voluptatum possimus saepe. Velit distinctio quia quas cupiditate sequi dolorum delectus quia quo corrupti labore. Aut mollitia et aut eligendi blanditiis corrupti aut rem voluptatibus veritatis ut. Et ab voluptas delectus repellat eligendi qui eos blanditiis et et eos. Qui ipsum aut ut voluptate nulla officia qui earum repudiandae quidem officia.
	// Fugit esse beatae placeat nihil libero et sunt voluptas velit maiores voluptatem. Reiciendis quia dolor amet blanditiis labore eligendi fugit reiciendis modi est incidunt. Ut ut rerum odit sit voluptatem ipsam consequatur est incidunt aut quis. Non dolorem voluptatum pariatur esse eos blanditiis rem eos commodi accusamus occaecati. Qui voluptatem eos et tempora et id at voluptas suscipit magni cupiditate.
}

func TestLoremIpsumParagraph(t *testing.T) {
	for _, count := range []struct{ parag, sent, words int }{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 1, -100},
		{1, -100, 1},
		{-100, 1, 1},
		{0, 0, 0},
	} {
		if LoremIpsumParagraph(count.parag, count.sent, count.words, " ") != "" {
			t.Errorf("result should be blank for %v input", count)
		}
	}
}

func BenchmarkLoremIpsumParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsumParagraph(3, 5, 12, "\n")
	}
}
