package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleEmoji() {
	Seed(11)
	fmt.Println(Emoji())
	// Output: 🧛
}

func ExampleFaker_Emoji() {
	f := New(11)
	fmt.Println(f.Emoji())
	// Output: 🧛
}

func BenchmarkEmoji(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Emoji()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Emoji()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Emoji()
		}
	})
}

func ExampleEmojiDescription() {
	Seed(11)
	fmt.Println(EmojiDescription())
	// Output: confetti ball
}

func ExampleFaker_EmojiDescription() {
	f := New(11)
	fmt.Println(f.EmojiDescription())
	// Output: confetti ball
}

func BenchmarkEmojiDescription(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EmojiDescription()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.EmojiDescription()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.EmojiDescription()
		}
	})
}

func ExampleEmojiCategory() {
	Seed(11)
	fmt.Println(EmojiCategory())
	// Output: Food & Drink
}

func ExampleFaker_EmojiCategory() {
	f := New(11)
	fmt.Println(f.EmojiCategory())
	// Output: Food & Drink
}

func BenchmarkEmojiCategory(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EmojiCategory()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.EmojiCategory()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.EmojiCategory()
		}
	})
}

func ExampleEmojiAlias() {
	Seed(11)
	fmt.Println(EmojiAlias())
	// Output: deaf_person
}

func ExampleFaker_EmojiAlias() {
	f := New(11)
	fmt.Println(f.EmojiAlias())
	// Output: deaf_person
}

func BenchmarkEmojiAlias(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EmojiAlias()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.EmojiAlias()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.EmojiAlias()
		}
	})
}

func ExampleEmojiTag() {
	Seed(11)
	fmt.Println(EmojiTag())
	// Output: strong
}

func ExampleFaker_EmojiTag() {
	f := New(11)
	fmt.Println(f.EmojiTag())
	// Output: strong
}

func BenchmarkEmojiTag(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			EmojiTag()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.EmojiTag()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.EmojiTag()
		}
	})
}
