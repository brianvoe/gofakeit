package gofakeit

import (
	"strings"
	"testing"
)

func TestISBN13(t *testing.T) {
	Seed(11)
	isbn := ISBN13("-")
	if len(strings.ReplaceAll(isbn, "-", "")) != 13 {
		t.Errorf("\"%s\" is not a valid ISBN string with length 13", isbn)
	}
}

func TestISBN10(t *testing.T) {
	Seed(11)
	isbn := ISBN10("-")
	if len(strings.ReplaceAll(isbn, "-", "")) != 10 {
		t.Errorf("\"%s\" is not a valid ISBN string with length 10", isbn)
	}
}

// Benchmarks
func BenchmarkISBN13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ISBN13("-")
	}
}

func BenchmarkISBN10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ISBN10("-")
	}
}
