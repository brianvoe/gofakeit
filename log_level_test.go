package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSysLogLevel() {
	Seed(11)
	fmt.Println(SysLogLevel())
	// Output: EMERG
}

func BenchmarkSysLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SysLogLevel()
	}
}

func ExampleApacheLogLevel() {
	Seed(11)
	fmt.Println(ApacheLogLevel())
	// Output: ERROR
}

func BenchmarkApacheLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ApacheLogLevel()
	}
}
