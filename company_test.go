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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Company()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Company()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Company()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CompanySuffix()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CompanySuffix()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CompanySuffix()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BuzzWord()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BuzzWord()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BuzzWord()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BS()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.BS()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.BS()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Job()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Job()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Job()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			JobTitle()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.JobTitle()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.JobTitle()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			JobDescriptor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.JobDescriptor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.JobDescriptor()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			JobLevel()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.JobLevel()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.JobLevel()
		}
	})
}
