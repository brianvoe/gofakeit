package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCreditCardType() {
	Seed(11)
	fmt.Println(CreditCardType())
	// Output: Visa
}

func BenchmarkCreditCardType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardType()
	}
}

func ExampleCreditCardNumber() {
	Seed(11)
	fmt.Println(CreditCardNumber())
	// Output: 4287271570245748
}

func BenchmarkCreditCardNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardNumber()
	}
}

func ExampleCreditCardExp() {
	Seed(11)
	fmt.Println(CreditCardExp())
	// Output: 01/16
}

func BenchmarkCreditCardExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardExp()
	}
}

func ExampleCreditCardCvv() {
	Seed(11)
	fmt.Println(CreditCardCvv())
	// Output: 328
}

func BenchmarkCreditCardCvv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardCvv()
	}
}
