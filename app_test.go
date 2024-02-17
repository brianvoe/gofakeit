package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAppName() {
	Seed(11)
	fmt.Println(AppName())

	// Output: Swanthink
}

func ExampleFaker_AppName() {
	f := New(11)
	fmt.Println(f.AppName())

	// Output: Swanthink
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

	// Output: 5.18.4
}

func ExampleFaker_AppVersion() {
	f := New(11)
	fmt.Println(f.AppVersion())

	// Output: 5.18.4
}

func BenchmarkAppVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppVersion()
	}
}

func ExampleAppAuthor() {
	Seed(11)
	fmt.Println(AppAuthor())

	// Output: StreetEasy
}

func ExampleFaker_AppAuthor() {
	f := New(11)
	fmt.Println(f.AppAuthor())

	// Output: StreetEasy
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
