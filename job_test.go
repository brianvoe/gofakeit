package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleJob() {
	fmt.Println(Job())
	// Output: Analyst
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Job()
	}
}

func ExampleJobDescriptor() {
	fmt.Println(JobDescriptor())
	// Output: Internal
}

func BenchmarkJobDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobDescriptor()
	}
}

func ExampleJobLevel() {
	fmt.Println(JobLevel())
	// Output: Assurance
}

func BenchmarkJobLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobLevel()
	}
}
