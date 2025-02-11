package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleBank() {
	Seed(11)
	bank := Bank()
	fmt.Println(bank.Name)
	fmt.Println(bank.Type)

	// Output: Postal Savings Bank of China
	// Retail Bank
}

func ExampleFaker_Bank() {
	f := New(11)
	bank := f.Bank()
	fmt.Println(bank.Name)
	fmt.Println(bank.Type)

	// Output: Postal Savings Bank of China
	// Retail Bank
}

func BenchmarkBank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bank()
	}
}

func TestBank(t *testing.T) {
	for i := 0; i < 100; i++ {
		Bank()
	}
}

func ExampleBankName() {
	Seed(11)
	fmt.Println(BankName())

	// Output: Postal Savings Bank of China
}

func ExampleFaker_BankName() {
	f := New(11)
	fmt.Println(f.BankName())

	// Output: Postal Savings Bank of China
}

func BenchmarkBankName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BankName()
	}
}

func ExampleBankType() {
	Seed(11)
	fmt.Println(BankType())

	// Output: Savings Bank
}

func ExampleFaker_BankType() {
	f := New(11)
	fmt.Println(f.BankType())

	// Output: Savings Bank
}

func BenchmarkBankType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BankType()
	}
}
