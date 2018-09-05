package gofakeit

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleCreditCard() {
	Seed(11)
	ccInfo := CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
	// Output: Visa
	// 6587271570245748
	// 03/25
	// 675
}

func BenchmarkCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCard()
	}
}

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

func ExampleCreditCardNumberLuhn() {
	Seed(11)
	fmt.Println(CreditCardNumberLuhn())
	// Output: 4007208855354357
}

func BenchmarkCreditCardNumberLuhn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardNumberLuhn()
	}
}

func TestCreditCardNumberLuhn(t *testing.T) {
	Seed(0)
	for i := 0; i < 100; i++ {
		cc := strconv.Itoa(CreditCardNumberLuhn())
		if !luhn(cc) {
			t.Errorf("not luhn valid: %s", cc)
		}
	}
}

func TestLuhn(t *testing.T) {
	// Lets make sure this card is invalid
	if luhn("867gfsd5309") {
		t.Error("card should have failed")
	}
}

func ExampleCreditCardExp() {
	Seed(11)
	fmt.Println(CreditCardExp())
	// Output: 01/20
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
