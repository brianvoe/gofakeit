package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSong() {
	Seed(11)
	song := Song()
	fmt.Println(song.Name)
	fmt.Println(song.Artist)
	fmt.Println(song.Genre)

	// Output: What Was I Made For?
	// Taylor Swift
	// Country
}

func ExampleFaker_Song() {
	f := New(11)
	song := f.Song()
	fmt.Println(song.Name)
	fmt.Println(song.Artist)
	fmt.Println(song.Genre)

	// Output: What Was I Made For?
	// Taylor Swift
	// Country
}

func BenchmarkSong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Song()
	}
}

func TestSong(t *testing.T) {
	for i := 0; i < 100; i++ {
		Song()
	}
}

func ExampleSongName() {
	Seed(11)
	fmt.Println(SongName())

	// Output: What Was I Made For?
}

func ExampleFaker_SongName() {
	f := New(11)
	fmt.Println(f.SongName())

	// Output: What Was I Made For?
}

func BenchmarkSongName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SongName()
	}
}

func ExampleSongArtist() {
	Seed(11)
	fmt.Println(SongArtist())

	// Output: The Jacksons
}

func ExampleFaker_SongArtist() {
	f := New(11)
	fmt.Println(f.SongArtist())

	// Output: The Jacksons
}

func BenchmarkSongArtist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SongArtist()
	}
}

func ExampleSongGenre() {
	Seed(11)
	fmt.Println(SongGenre())

	// Output: Synthwave
}

func ExampleFaker_SongGenre() {
	f := New(11)
	fmt.Println(f.SongGenre())

	// Output: Synthwave
}

func BenchmarkSongGenre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SongGenre()
	}
}
