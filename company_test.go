package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCompany() {
	Seed(11)
	fmt.Println(Company())

	// Output: ClearHealthCosts
}

func ExampleFaker_Company() {
	f := New(11)
	fmt.Println(f.Company())

	// Output: ClearHealthCosts
}

func BenchmarkCompany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Company()
	}
}

func TestCompany(t *testing.T) {
	for i := 0; i < 100; i++ {
		Company()
	}
}

func ExampleCompanySuffix() {
	Seed(11)
	fmt.Println(CompanySuffix())

	// Output: Inc
}

func ExampleFaker_CompanySuffix() {
	f := New(11)
	fmt.Println(f.CompanySuffix())

	// Output: Inc
}

func BenchmarkCompanySuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompanySuffix()
	}
}

func ExampleBlurb() {
	Seed(11)
	fmt.Println(Blurb())

	// Output: Motivation
}

func ExampleFaker_Blurb() {
	f := New(11)
	fmt.Println(f.Blurb())

	// Output: Motivation
}

func BenchmarkBlurb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Blurb()
	}
}
func ExampleBuzzWord() {
	Seed(11)
	fmt.Println(BuzzWord())

	// Output: disintermediate
}

func ExampleFaker_BuzzWord() {
	f := New(11)
	fmt.Println(f.BuzzWord())

	// Output: disintermediate
}

func BenchmarkBuzzWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuzzWord()
	}
}

func ExampleBS() {
	Seed(11)
	fmt.Println(BS())

	// Output: front-end
}

func ExampleFaker_BS() {
	f := New(11)
	fmt.Println(f.BS())

	// Output: front-end
}

func BenchmarkBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BS()
	}
}

func ExampleJob() {
	Seed(11)
	jobInfo := Job()
	fmt.Println(jobInfo.Company)
	fmt.Println(jobInfo.Title)
	fmt.Println(jobInfo.Descriptor)
	fmt.Println(jobInfo.Level)

	// Output: ClearHealthCosts
	// Agent
	// Future
	// Tactics
}

func ExampleFaker_Job() {
	f := New(11)
	jobInfo := f.Job()
	fmt.Println(jobInfo.Company)
	fmt.Println(jobInfo.Title)
	fmt.Println(jobInfo.Descriptor)
	fmt.Println(jobInfo.Level)

	// Output: ClearHealthCosts
	// Agent
	// Future
	// Tactics
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

func ExampleFaker_JobTitle() {
	f := New(11)
	fmt.Println(f.JobTitle())

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

func ExampleFaker_JobDescriptor() {
	f := New(11)
	fmt.Println(f.JobDescriptor())

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

func ExampleFaker_JobLevel() {
	f := New(11)
	fmt.Println(f.JobLevel())

	// Output: Assurance
}

func BenchmarkJobLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobLevel()
	}
}

func ExampleSlogan() {
	Seed(11)
	fmt.Println(Slogan())

	// Output: Universal seamless Focus, interactive.
}

func ExampleFaker_Slogan() {
	f := New(11)
	fmt.Println(f.Slogan())

	// Output: Universal seamless Focus, interactive.
}

func BenchmarkSlogan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slogan()
	}
}
