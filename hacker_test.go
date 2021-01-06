package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHackerPhrase() {
	Seed(11)
	fmt.Println(HackerPhrase())
	// Output: If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!
}

func ExampleFaker_HackerPhrase() {
	f := New(11)
	fmt.Println(f.HackerPhrase())
	// Output: If we calculate the program, we can get to the AI pixel through the redundant XSS matrix!
}

func BenchmarkHackerPhrase(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackerPhrase()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackerPhrase()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackerPhrase()
		}
	})
}

func ExampleHackerAbbreviation() {
	Seed(11)
	fmt.Println(HackerAbbreviation())
	// Output: ADP
}

func ExampleFaker_HackerAbbreviation() {
	f := New(11)
	fmt.Println(f.HackerAbbreviation())
	// Output: ADP
}

func BenchmarkHackerAbbreviation(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackerAbbreviation()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackerAbbreviation()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackerAbbreviation()
		}
	})
}

func ExampleHackerAdjective() {
	Seed(11)
	fmt.Println(HackerAdjective())
	// Output: wireless
}

func ExampleFaker_HackerAdjective() {
	f := New(11)
	fmt.Println(f.HackerAdjective())
	// Output: wireless
}

func BenchmarkHackerAdjective(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackerAdjective()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackerAdjective()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackerAdjective()
		}
	})
}

func ExampleHackerNoun() {
	Seed(11)
	fmt.Println(HackerNoun())
	// Output: driver
}

func ExampleFaker_HackerNoun() {
	f := New(11)
	fmt.Println(f.HackerNoun())
	// Output: driver
}

func BenchmarkHackerNoun(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackerNoun()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackerNoun()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackerNoun()
		}
	})
}

func ExampleHackerVerb() {
	Seed(11)
	fmt.Println(HackerVerb())
	// Output: synthesize
}

func ExampleFaker_HackerVerb() {
	f := New(11)
	fmt.Println(f.HackerVerb())
	// Output: synthesize
}

func BenchmarkHackerVerb(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackerVerb()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackerVerb()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackerVerb()
		}
	})
}

func ExampleHackeringVerb() {
	Seed(11)
	fmt.Println(HackeringVerb())
	// Output: connecting
}

func ExampleFaker_HackeringVerb() {
	f := New(11)
	fmt.Println(f.HackeringVerb())
	// Output: connecting
}

func BenchmarkHackeringVerb(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HackeringVerb()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HackeringVerb()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HackeringVerb()
		}
	})
}
