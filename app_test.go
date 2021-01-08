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

func ExampleFaker_AppName() {
	f := New(11)
	fmt.Println(f.AppName())
	// Output: Parkrespond
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AppName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.AppName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.AppName()
		}
	})
}

func ExampleAppVersion() {
	Seed(11)
	fmt.Println(AppVersion())
	// Output: 1.12.14
}

func ExampleFaker_AppVersion() {
	f := New(11)
	fmt.Println(f.AppVersion())
	// Output: 1.12.14
}

func BenchmarkAppVersion(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AppVersion()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.AppVersion()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.AppVersion()
		}
	})
}

func ExampleAppAuthor() {
	Seed(11)
	fmt.Println(AppAuthor())
	// Output: Qado Energy, Inc.
}

func ExampleFaker_AppAuthor() {
	f := New(11)
	fmt.Println(f.AppAuthor())
	// Output: Qado Energy, Inc.
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AppAuthor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.AppAuthor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.AppAuthor()
		}
	})
}
