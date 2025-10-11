package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleEmoji() {
	Seed(11)
	fmt.Println(Emoji())

	// Output: 🛫
}

func ExampleFaker_Emoji() {
	f := New(11)
	fmt.Println(f.Emoji())

	// Output: 🛫
}

func BenchmarkEmoji(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Emoji()
	}
}

func ExampleEmojiCategory() {
	Seed(11)
	fmt.Println(EmojiCategory())

	// Output: Flags
}

func ExampleFaker_EmojiCategory() {
	f := New(11)
	fmt.Println(f.EmojiCategory())

	// Output: Flags
}

func BenchmarkEmojiCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiCategory()
	}
}

func ExampleEmojiAlias() {
	Seed(11)
	fmt.Println(EmojiAlias())

	// Output: key
}

func ExampleFaker_EmojiAlias() {
	f := New(11)
	fmt.Println(f.EmojiAlias())

	// Output: key
}

func BenchmarkEmojiAlias(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiAlias()
	}
}

func ExampleEmojiTag() {
	Seed(11)
	fmt.Println(EmojiTag())

	// Output: password
}

func ExampleFaker_EmojiTag() {
	f := New(11)
	fmt.Println(f.EmojiTag())

	// Output: password
}

func BenchmarkEmojiTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiTag()
	}
}

func ExampleEmojiFlag() {
	Seed(11)
	fmt.Println(EmojiFlag())

	// Output: 🇹🇴
}

func ExampleFaker_EmojiFlag() {
	f := New(11)
	fmt.Println(f.EmojiFlag())

	// Output: 🇹🇴
}

func BenchmarkEmojiFlag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiFlag()
	}
}

func ExampleEmojiAnimal() {
	Seed(11)
	fmt.Println(EmojiAnimal())

	// Output: 🐌
}

func ExampleFaker_EmojiAnimal() {
	f := New(11)
	fmt.Println(f.EmojiAnimal())

	// Output: 🐌
}

func BenchmarkEmojiAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiAnimal()
	}
}

func ExampleEmojiFood() {
	Seed(11)
	fmt.Println(EmojiFood())

	// Output: 🍾
}

func ExampleFaker_EmojiFood() {
	f := New(11)
	fmt.Println(f.EmojiFood())

	// Output: 🍾
}

func BenchmarkEmojiFood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiFood()
	}
}

func ExampleEmojiPlant() {
	Seed(11)
	fmt.Println(EmojiPlant())

	// Output: 🍁
}

func ExampleFaker_EmojiPlant() {
	f := New(11)
	fmt.Println(f.EmojiPlant())

	// Output: 🍁
}

func BenchmarkEmojiPlant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiPlant()
	}
}

func ExampleEmojiMusic() {
	Seed(11)
	fmt.Println(EmojiMusic())

	// Output: 🎚️
}

func ExampleFaker_EmojiMusic() {
	f := New(11)
	fmt.Println(f.EmojiMusic())

	// Output: 🎚️
}

func BenchmarkEmojiMusic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiMusic()
	}
}

func ExampleEmojiSentence() {
	Seed(11)
	fmt.Println(EmojiSentence())

	// Output: Category: Symbols 📠.
}

func ExampleFaker_EmojiSentence() {
	f := New(11)
	fmt.Println(f.EmojiSentence())

	// Output: Category: Symbols 📠.
}

func TestEmojiSentence(t *testing.T) {
	for i := 0; i < 100; i++ {
		sentence := EmojiSentence()

		// Should not be empty
		if sentence == "" {
			t.Error("Emoji sentence should not be empty")
		}

		// Should not contain unreplaced braces
		if strings.Contains(sentence, "{") || strings.Contains(sentence, "}") {
			t.Errorf("Emoji sentence should not contain braces: %s", sentence)
		}
	}

	// Test with 0 (should default to 1-3)
	sentence := EmojiSentence()
	if sentence == "" {
		t.Error("Emoji sentence with 0 count should not be empty")
	}
}

func BenchmarkEmojiSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiSentence()
	}
}
