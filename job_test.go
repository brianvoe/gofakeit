package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleJob() {
	Seed(11)
	fmt.Println(Job())
	// Output: Director
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Job()
	}
}

func ExampleJobDescriptor() {
	Seed(11)
	fmt.Println(JobDescriptor())
	// Output: Central
}

func BenchmarkJobDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobDescriptor()
	}
}

func ExampleJobLevel() {
	Seed(11)
	fmt.Println(JobLevel())
	// Output: Assurance
}

func BenchmarkJobLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobLevel()
	}
}
