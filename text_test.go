package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleComment() {
	Seed(11)
	fmt.Println(Comment())

	// Output: After seeing this fear, maybe drink the trip and adjust the here class
}

func ExampleFaker_Comment() {
	f := New(11)
	fmt.Println(f.Comment())

	// Output: After seeing this fear, maybe drink the trip and adjust the here class
}

func TestComment(t *testing.T) {
	f := New(11)
	for i := 0; i < 10000; i++ {
		comment := f.Comment()
		if comment == "" {
			t.Error("Comment is empty")
		}

		// make sure it doesn't contain { or }
		if strings.Contains(comment, "{") || strings.Contains(comment, "}") {
			t.Error("Comment contains { or }")
		}
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
	fmt.Println(Sentence())

	// Output: Weekends reserve time for Sun Bathing and day.
}

func ExampleFaker_Sentence() {
	f := New(11)
	fmt.Println(f.Sentence())

	// Output: Weekends reserve time for Sun Bathing and day.
}

func TestSentence(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sentence := Sentence()
		if sentence == "" {
			t.Error("sentence should not be empty")
		}

		// make sure it doesn't contain { or }
		if strings.Contains(sentence, "{") || strings.Contains(sentence, "}") {
			t.Error("sentence contains { or }")
		}
	}
}

func BenchmarkSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sentence()
	}
}

func ExampleParagraph() {
	Seed(11)
	fmt.Println(Paragraph())

	// Output: Publish a changelog entry for the day. Steady jittery progress in Denver should visible.
}

func ExampleFaker_Paragraph() {
	f := New(11)
	fmt.Println(f.Paragraph())

	// Output: Publish a changelog entry for the day. Steady jittery progress in Denver should visible.
}

func TestParagraph(t *testing.T) {
	for i := 0; i < 1000; i++ {
		paragraph := Paragraph()
		if paragraph == "" {
			t.Error("paragraph should not be empty")
		}
	}
}

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Paragraph()
	}
}

func ExampleQuestion() {
	Seed(11)
	fmt.Println(Question())

	// Output: What is the marginal gain from fear?
}

func ExampleFaker_Question() {
	f := New(11)
	fmt.Println(f.Question())

	// Output: What is the marginal gain from fear?
}

func BenchmarkQuestion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Question()
	}
}

func ExampleQuote() {
	Seed(11)
	fmt.Println(Quote())

	// Output: "Energy is contagious—so is fear"
}

func ExampleFaker_Quote() {
	f := New(11)
	fmt.Println(f.Quote())

	// Output: "Energy is contagious—so is fear"
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
