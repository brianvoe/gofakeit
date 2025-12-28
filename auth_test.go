package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())

	// Output: Gregg812
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())

	// Output: Gregg812
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

	// Output: tcvncypbfolpftvlyplgdxiwibpsautg
	// ZXVXDVFQGCECFDBRWEKPWATHKRGKWDIZ
	// 78902501101475351179812748788830
	// !!._-._*--@@.@---@_@.-_!_!-*_!*@
	// 96 rvcB@f0.PNzL!qp 7hP_V 7g!vV 9
	// 6hJDB
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(true, false, false, false, false, 32))
	fmt.Println(f.Password(false, true, false, false, false, 32))
	fmt.Println(f.Password(false, false, true, false, false, 32))
	fmt.Println(f.Password(false, false, false, true, false, 32))
	fmt.Println(f.Password(true, true, true, true, true, 32))
	fmt.Println(f.Password(true, true, true, true, true, 4))

	// Output: tcvncypbfolpftvlyplgdxiwibpsautg
	// ZXVXDVFQGCECFDBRWEKPWATHKRGKWDIZ
	// 78902501101475351179812748788830
	// !!._-._*--@@.@---@_@.-_!_!-*_!*@
	// 96 rvcB@f0.PNzL!qp 7hP_V 7g!vV 9
	// 6hJDB
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password(true, true, true, true, true, 50)
	}
}
