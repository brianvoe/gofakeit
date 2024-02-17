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

	// Output: Things Fall Apart
	// Toni Morrison
	// Erotic
}

func ExampleFaker_Book() {
	f := New(11)
	book := f.Book()
	fmt.Println(book.Title)
	fmt.Println(book.Author)
	fmt.Println(book.Genre)

	// Output: Things Fall Apart
	// Toni Morrison
	// Erotic
}

func BenchmarkBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Book()
	}
}

func TestBook(t *testing.T) {
	for i := 0; i < 100; i++ {
		Book()
	}
}

func ExampleBookTitle() {
	Seed(11)
	fmt.Println(BookTitle())

	// Output: Things Fall Apart
}

func ExampleFaker_BookTitle() {
	f := New(11)
	fmt.Println(f.BookTitle())

	// Output: Things Fall Apart
}

func BenchmarkBookTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BookTitle()
	}
}

func ExampleBookAuthor() {
	Seed(11)
	fmt.Println(BookAuthor())

	// Output: Vladimir Nabokov
}

func ExampleFaker_BookAuthor() {
	f := New(11)
	fmt.Println(f.BookAuthor())

	// Output: Vladimir Nabokov
}

func BenchmarkBookAuthor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BookAuthor()
	}
}

func ExampleBookGenre() {
	Seed(11)
	fmt.Println(BookGenre())

	// Output: Thriller
}

func ExampleFaker_BookGenre() {
	f := New(11)
	fmt.Println(f.BookGenre())

	// Output: Thriller
}

func BenchmarkBookGenre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BookGenre()
	}
}
