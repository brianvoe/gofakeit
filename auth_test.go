package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())
	// Output: Daniel1364
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())
	// Output: Daniel1364
}

func BenchmarkUsername(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Username()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Username()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Username()
		}
	})
}

func TestPassword(t *testing.T) {
	length := 10

	pass := Password(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}

	// Test fully empty
	pass = Password(false, false, false, false, false, length)
	if pass == "" {
		t.Error("Password should not be empty")
	}
}

func ExamplePassword() {
	Seed(11)
	fmt.Println(Password(true, false, false, false, false, 32))
	fmt.Println(Password(false, true, false, false, false, 32))
	fmt.Println(Password(false, false, true, false, false, 32))
	fmt.Println(Password(false, false, false, true, false, 32))
	fmt.Println(Password(true, true, true, true, true, 32))
	fmt.Println(Password(true, true, true, true, true, 4))
	// Output: vodnqxzsuptgehrzylximvylxzoywexw
	// ZSRQWJFJWCSTVGXKYKWMLIAFGFELFJRG
	// 61718615932495608398906260648432
	// @$,@#:,(,).{?:%?)>*..<=};#$(:{==
	// CkF{wwb:?Kb},w?vdz{Zox C&>Prt99:
	// j ;9X
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(true, false, false, false, false, 32))
	fmt.Println(f.Password(false, true, false, false, false, 32))
	fmt.Println(f.Password(false, false, true, false, false, 32))
	fmt.Println(f.Password(false, false, false, true, false, 32))
	fmt.Println(f.Password(true, true, true, true, true, 32))
	fmt.Println(f.Password(true, true, true, true, true, 4))
	// Output: vodnqxzsuptgehrzylximvylxzoywexw
	// ZSRQWJFJWCSTVGXKYKWMLIAFGFELFJRG
	// 61718615932495608398906260648432
	// @$,@#:,(,).{?:%?)>*..<=};#$(:{==
	// CkF{wwb:?Kb},w?vdz{Zox C&>Prt99:
	// j ;9X
}

func BenchmarkPassword(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Password(true, true, true, true, true, 8)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Password(true, true, true, true, true, 8)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Password(true, true, true, true, true, 8)
		}
	})
}
