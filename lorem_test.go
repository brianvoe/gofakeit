package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLoremIpsumWord() {
	Seed(11)
	fmt.Println(LoremIpsumWord())

	// Output: eveniet
}

func ExampleFaker_LoremIpsumWord() {
	f := New(11)
	fmt.Println(f.LoremIpsumWord())

	// Output: eveniet
}

func BenchmarkLoremIpsumWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsumWord()
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
