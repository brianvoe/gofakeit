package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNoun() {
	Seed(11)
	fmt.Println(Noun())
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

func ExampleLoremIpsumWord() {
	Seed(11)
	fmt.Println(LoremIpsumWord())
	// Output: quia
}

func BenchmarkLoremIpsumWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoremIpsumWord()
	}
}

func ExampleLoremIpsumSentence() {
	Seed(11)
	fmt.Println(LoremIpsumSentence(5))
	// Output: Quia quae repellat consequatur quidem.
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
	// Output: Quia quae repellat consequatur quidem nisi quo qui voluptatum accusantium quisquam amet. Quas et ut non dolorem ipsam aut enim assumenda mollitia harum ut. Dicta similique veniam nulla voluptas at excepturi non ad maxime at non. Eaque hic repellat praesentium voluptatem qui consequuntur dolor iusto autem velit aut. Fugit tempore exercitationem harum consequatur voluptatum modi minima aut eaque et et.
	// Aut ea voluptatem dignissimos expedita odit tempore quod aut beatae ipsam iste. Minus voluptatibus dolorem maiores eius sed nihil vel enim odio voluptatem accusamus. Natus quibusdam temporibus tenetur cumque sint necessitatibus dolorem ex ducimus iusto ex. Voluptatem neque dicta explicabo officiis et ducimus sit ut ut praesentium pariatur. Illum molestias nisi at dolore ut voluptatem accusantium et fugiat et ut.
	// Explicabo incidunt reprehenderit non quia dignissimos recusandae vitae soluta quia et quia. Aut veniam voluptas consequatur placeat sapiente non eveniet voluptatibus magni velit eum. Nobis vel repellendus sed est qui autem laudantium quidem quam ullam consequatur. Aut iusto ut commodi similique quae voluptatem atque qui fugiat eum aut. Quis distinctio consequatur voluptatem vel aliquid aut laborum facere officiis iure tempora.
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

func BenchmarkPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phrase()
	}
}
