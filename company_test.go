package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCompany() {
	Seed(11)
	fmt.Println(Company())

	// Output: TransparaGov
}

func ExampleFaker_Company() {
	f := New(11)
	fmt.Println(f.Company())

	// Output: TransparaGov
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

	// Output: Teamwork
}

func ExampleFaker_Blurb() {
	f := New(11)
	fmt.Println(f.Blurb())

	// Output: Teamwork
}

func BenchmarkBlurb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Blurb()
	}
}
func ExampleBuzzWord() {
	Seed(11)
	fmt.Println(BuzzWord())

	// Output: open system
}

func ExampleFaker_BuzzWord() {
	f := New(11)
	fmt.Println(f.BuzzWord())

	// Output: open system
}

func BenchmarkBuzzWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuzzWord()
	}
}

func ExampleBS() {
	Seed(11)
	fmt.Println(BS())

	// Output: models
}

func ExampleFaker_BS() {
	f := New(11)
	fmt.Println(f.BS())

	// Output: models
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

	// Output: TransparaGov
	// Specialist
	// Direct
	// Configuration
}

func ExampleFaker_Job() {
	f := New(11)
	jobInfo := f.Job()
	fmt.Println(jobInfo.Company)
	fmt.Println(jobInfo.Title)
	fmt.Println(jobInfo.Descriptor)
	fmt.Println(jobInfo.Level)

	// Output: TransparaGov
	// Specialist
	// Direct
	// Configuration
}

func BenchmarkJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Job()
	}
}

func ExampleJobTitle() {
	Seed(11)
	fmt.Println(JobTitle())

	// Output: Strategist
}

func ExampleFaker_JobTitle() {
	f := New(11)
	fmt.Println(f.JobTitle())

	// Output: Strategist
}

func BenchmarkJobTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobTitle()
	}
}

func ExampleJobDescriptor() {
	Seed(11)
	fmt.Println(JobDescriptor())

	// Output: Product
}

func ExampleFaker_JobDescriptor() {
	f := New(11)
	fmt.Println(f.JobDescriptor())

	// Output: Product
}

func BenchmarkJobDescriptor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobDescriptor()
	}
}

func ExampleJobLevel() {
	Seed(11)
	fmt.Println(JobLevel())

	// Output: Solutions
}

func ExampleFaker_JobLevel() {
	f := New(11)
	fmt.Println(f.JobLevel())

	// Output: Solutions
}

func BenchmarkJobLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JobLevel()
	}
}

func ExampleSlogan() {
	Seed(11)
	fmt.Println(Slogan())

	// Output: local area network maximize Drive, mission-critical.
}

func ExampleFaker_Slogan() {
	f := New(11)
	fmt.Println(f.Slogan())

	// Output: local area network maximize Drive, mission-critical.
}

func BenchmarkSlogan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slogan()
	}
}
