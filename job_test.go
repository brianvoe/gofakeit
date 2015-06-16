package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleJob() {
	Seed(11)
	jobInfo := Job()
	fmt.Println(jobInfo.Company)
	fmt.Println(jobInfo.Title)
	fmt.Println(jobInfo.Descriptor)
	fmt.Println(jobInfo.Level)
	// Output: Moen, Pagac and Wuckert
	// Developer
	// National
	// Integration
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Job()
	}
}

func ExampleJobTitle() {
	Seed(11)
	fmt.Println(JobTitle())
	// Output: Director
}

func BenchmarkJobTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobTitle()
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
