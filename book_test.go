package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBook() {
	Seed(11)
	book := Book()
	fmt.Println(book.Title)
	fmt.Println(book.Author)
	fmt.Println(book.Genre)

	// Output: Anna Karenina
	// Toni Morrison
	// Thriller
}

func ExampleFaker_Book() {
	f := New(11)
	book := f.Book()
	fmt.Println(book.Title)
	fmt.Println(book.Author)
	fmt.Println(book.Genre)

	// Output: Anna Karenina
	// Toni Morrison
	// Thriller
}

func BenchmarkBook(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Book()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Book()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Book()
		}
	})
}

func TestBook(t *testing.T) {
	for i := 0; i < 100; i++ {
		Book()
	}
}

func ExampleBookTitle() {
	Seed(11)
	fmt.Println(BookTitle())

	// Output: Anna Karenina
}

func ExampleFaker_BookTitle() {
	f := New(11)
	fmt.Println(f.BookTitle())

	// Output: Anna Karenina
}

func BenchmarkBookTitle(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BookTitle()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BookTitle()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BookTitle()
		}
	})
}

func ExampleBookAuthor() {
	Seed(11)
	fmt.Println(BookAuthor())

	// Output: James Joyce
}

func ExampleFaker_BookAuthor() {
	f := New(11)
	fmt.Println(f.BookAuthor())

	// Output: James Joyce
}

func BenchmarkBookAuthor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BookAuthor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BookAuthor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BookAuthor()
		}
	})
}

func ExampleBookGenre() {
	Seed(11)
	fmt.Println(BookGenre())

	// Output: Crime
}

func ExampleFaker_BookGenre() {
	f := New(11)
	fmt.Println(f.BookGenre())

	// Output: Crime
}

func BenchmarkBookGenre(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BookGenre()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BookGenre()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BookGenre()
		}
	})
}
