package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleEmoji() {
	Seed(11)
	fmt.Println(Emoji())

	// Output: 🇫🇴
}

func ExampleFaker_Emoji() {
	f := New(11)
	fmt.Println(f.Emoji())

	// Output: 🇫🇴
}

func BenchmarkEmoji(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Emoji()
	}
}

func ExampleEmojiDescription() {
	Seed(11)
	fmt.Println(EmojiDescription())

	// Output: flag: European Union
}

func ExampleFaker_EmojiDescription() {
	f := New(11)
	fmt.Println(f.EmojiDescription())

	// Output: flag: European Union
}

func BenchmarkEmojiDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiDescription()
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

	// Output: eritrea
}

func ExampleFaker_EmojiAlias() {
	f := New(11)
	fmt.Println(f.EmojiAlias())

	// Output: eritrea
}

func BenchmarkEmojiAlias(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiAlias()
	}
}

func ExampleEmojiTag() {
	Seed(11)
	fmt.Println(EmojiTag())

	// Output: toilet
}

func ExampleFaker_EmojiTag() {
	f := New(11)
	fmt.Println(f.EmojiTag())

	// Output: toilet
}

func BenchmarkEmojiTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EmojiTag()
	}
}
