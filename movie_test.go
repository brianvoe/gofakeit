package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMovie() {
	Seed(11)
	movie := Movie()
	fmt.Println(movie.Name)
	fmt.Println(movie.Genre)

	// Output: The Terminator
	// Sport
}

func ExampleFaker_Movie() {
	f := New(11)
	movie := f.Movie()
	fmt.Println(movie.Name)
	fmt.Println(movie.Genre)

	// Output: The Terminator
	// Sport
}

func BenchmarkMovie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Movie()
	}
}

func TestMovie(t *testing.T) {
	for i := 0; i < 100; i++ {
		Movie()
	}
}

func ExampleMovieName() {
	Seed(11)
	fmt.Println(MovieName())

	// Output: The Terminator
}

func ExampleFaker_MovieName() {
	f := New(11)
	fmt.Println(f.MovieName())

	// Output: The Terminator
}

func BenchmarkMovieName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MovieName()
	}
}

func ExampleMovieGenre() {
	Seed(11)
	fmt.Println(MovieGenre())

	// Output: Thriller
}

func ExampleFaker_MovieGenre() {
	f := New(11)
	fmt.Println(f.MovieGenre())

	// Output: Thriller
}

func BenchmarkMovieGenre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MovieGenre()
	}
}
