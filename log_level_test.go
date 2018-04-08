package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLogLevel() {
	Seed(11)
	fmt.Println(LogLevel("")) // This will also use general
	fmt.Println(LogLevel("syslog"))
	fmt.Println(LogLevel("apache"))
	// Output: error
	// debug
	// trace1-8
}

func BenchmarkLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogLevel("general")
	}
}
