package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMusic() {
	Seed(11)
	music := Music()
	fmt.Println(music.Name)
	fmt.Println(music.Artist)
	fmt.Println(music.Genre)

	// Output: The Terminator
	// Sport
}

func ExampleFaker_Music() {
	f := New(11)
	music := f.Music()
	fmt.Println(music.Name)
	fmt.Println(music.Artist)
	fmt.Println(music.Genre)

	// Output: The Terminator
	// Sport
}

func BenchmarkMusic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Music()
	}
}

func TestMusic(t *testing.T) {
	for i := 0; i < 100; i++ {
		Music()
	}
}

func ExampleMusicName() {
	Seed(11)
	fmt.Println(MusicName())

	// Output: The Terminator
}

func ExampleFaker_MusicName() {
	f := New(11)
	fmt.Println(f.MusicName())

	// Output: The Terminator
}

func BenchmarkMusicName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MusicName()
	}
}

func ExampleMusicArtist() {
	Seed(11)
	fmt.Println(MusicArtist())

	// Output: Thriller
}

func ExampleFaker_MusicArtist() {
	f := New(11)
	fmt.Println(f.MusicArtist())

	// Output: Thriller
}

func BenchmarkMusicArtist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MusicArtist()
	}
}

func ExampleMusicGenre() {
	Seed(11)
	fmt.Println(MusicGenre())

	// Output: Thriller
}

func ExampleFaker_MusicGenre() {
	f := New(11)
	fmt.Println(f.MusicGenre())

	// Output: Thriller
}

func BenchmarkMusicGenre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MusicGenre()
	}
}
