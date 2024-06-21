package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())

	// Output: Treutel8125
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())

	// Output: Treutel8125
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Username()
	}
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

	// Test it doesnt start or end with a space
	for i := 0; i < 1000; i++ {
		pass = Password(true, true, true, true, true, length)
		if pass[0] == ' ' || pass[len(pass)-1] == ' ' {
			t.Error("Password should not start or end with a space")
		}
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

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(true, false, false, false, false, 32))
	fmt.Println(f.Password(false, true, false, false, false, 32))
	fmt.Println(f.Password(false, false, true, false, false, 32))
	fmt.Println(f.Password(false, false, false, true, false, 32))
	fmt.Println(f.Password(true, true, true, true, true, 32))
	fmt.Println(f.Password(true, true, true, true, true, 4))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password(true, true, true, true, true, 50)
	}
}
