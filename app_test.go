package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAppName() {
	Seed(11)
	fmt.Println(AppName())

	// Output: Oxbeing
}

func ExampleFaker_AppName() {
	f := New(11)
	fmt.Println(f.AppName())

	// Output: Oxbeing
}

func TestAppName(t *testing.T) {
	for i := 0; i < 100; i++ {
		name := AppName()
		if name == "" {
			t.Error("app name should not be empty")
		}
	}
}

func BenchmarkAppName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppName()
	}
}

func ExampleAppVersion() {
	Seed(11)
	fmt.Println(AppVersion())

	// Output: 1.17.20
}

func ExampleFaker_AppVersion() {
	f := New(11)
	fmt.Println(f.AppVersion())

	// Output: 1.17.20
}

func BenchmarkAppVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppVersion()
	}
}

func ExampleAppAuthor() {
	Seed(11)
	fmt.Println(AppAuthor())

	// Output: Marcel Pagac
}

func ExampleFaker_AppAuthor() {
	f := New(11)
	fmt.Println(f.AppAuthor())

	// Output: Marcel Pagac
}

func TestAuthor(t *testing.T) {
	for i := 0; i < 100; i++ {
		author := AppAuthor()
		if author == "" {
			t.Error("app author should not be empty")
		}
	}
}

func BenchmarkAppAuthor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppAuthor()
	}
}
