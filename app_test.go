package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAppName() {
	Seed(11)
	fmt.Println(AppName())
	// Output: Parkrespond
}

func BenchmarkAppName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppName()
	}
}

func ExampleAppVersion() {
	Seed(11)
	fmt.Println(AppVersion())
	// Output: 1.12.14
}

func BenchmarkAppVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppVersion()
	}
}

func ExampleAppAuthor() {
	Seed(11)
	fmt.Println(AppAuthor())
	// Output: Qado Energy, Inc.
}

func BenchmarkAppAuthor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppAuthor()
	}
}
