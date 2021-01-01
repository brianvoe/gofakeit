package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())
	// Output: oil
}

func BenchmarkNoun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Noun()
	}
}

func ExampleVerb() {
	Seed(11)
	fmt.Println(Verb())
	// Output: issue
}

func BenchmarkVerb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Verb()
	}
}

func ExampleAdverb() {
	Seed(11)
	fmt.Println(Adverb())
	// Output: incidentally
}

func BenchmarkAdverb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adverb()
	}
}

func ExamplePreposition() {
	Seed(11)
	fmt.Println(Preposition())
	// Output: over against
}

func BenchmarkPreposition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Preposition()
	}
}

func ExampleAdjective() {
	Seed(11)
	fmt.Println(Adjective())
	// Output: simple
}

func BenchmarkAdjective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Adjective()
	}
}

func ExampleWord() {
	Seed(11)
	fmt.Println(Word())
	// Output: succeed
}

func BenchmarkWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Word()
	}
}

func ExampleSentence() {
	Seed(11)
	fmt.Println(Sentence(5))
	// Output: Enhance resistance form account operate.
}

func TestSentence(t *testing.T) {
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
	// Output: Think copy diet point hate section duty peace receive engage disease language. Rock sense complain advice ride sign west culture doctor issue in purchase. Alter picture neck director set family sit effort change research institution follow. Doubt die suit name victory lie port coal gain help aircraft answer. Lift interpret get sell establish repeat simon noise reflect assist push spirit.
	// Behave do suffer liability award appeal paint search soil reach hotel imagine. Tend hang majority let break article flight remind acquire seem course pray. Spokesman cut request gain apply adult bottom half truth argue get intention. Mark sit content chief holiday mark speak observe project supply imagine help. Cry church investigate acid offer coffee speaker manner theory activity owe cut.
	// Worry secretary recognise treaty cast become succeed observe address miss include spend. Development truth generate remember mike play surface direction john relief measure explanation. Sort persuade sit television performance transport perform sound iron county complete finance. Family prepare organise spot recognise charge pray specialist way victim inform burst. Resolution shop noise pull appeal close assist combination fuel press appeal consider.
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

func ExampleLoremIpsumWord() {
	Seed(11)
	fmt.Println(LoremIpsumWord())
	// Output: est
}

func BenchmarkLoremIpsumWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsumWord()
	}
}

func ExampleLoremIpsumSentence() {
	Seed(11)
	fmt.Println(LoremIpsumSentence(5))
	// Output: Eum dolores fugiat consequatur odit.
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
	// Output: In aut ut omnis eius assumenda eligendi ad consectetur enim occaecati expedita. Et vel incidunt possimus iste et iure expedita illum beatae nihil et. Reprehenderit quam sed quae aliquid velit voluptates nemo in sed aut harum. Rem hic et sit itaque et quibusdam ipsam in laboriosam numquam culpa. Maxime voluptas natus facilis eius voluptatem et error amet dolor labore est.
	// Consectetur eos veritatis optio facilis distinctio occaecati deleniti incidunt autem expedita consectetur. Voluptates magnam excepturi eveniet est nihil consequatur quibusdam incidunt est maxime est. Voluptates placeat qui voluptatem expedita quae et nemo voluptatem mollitia aliquid repudiandae. Alias veritatis odio omnis alias minima ad labore id porro explicabo ut. Ea aut dolorem voluptates eligendi quis iusto culpa non ipsum repellendus rem.
	// Minima laborum qui ut ratione illo itaque velit voluptas perferendis fugit est. Quia fugiat sunt deserunt amet ut et voluptatem earum non ratione non. Sint aut ea autem eum vel aperiam omnis in tenetur est atque. Et est dicta molestiae libero nobis ut earum non repellendus laborum atque. Earum molestiae est sunt in ipsum consequatur est maiores ex doloribus debitis.
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

func ExampleQuestion() {
	Seed(11)
	fmt.Println(Question())
	// Output: Goth bicycle rights twee?
}

func BenchmarkQuestion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Question()
	}
}

func ExampleQuote() {
	Seed(11)
	fmt.Println(Quote())
	// Output: "Banjo next level Thundercats." - Hyman Wyman
}

func BenchmarkQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Quote()
	}
}

func ExamplePhrase() {
	Seed(11)
	fmt.Println(Phrase())
	// Output: what's yours
}

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}
