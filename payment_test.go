package gofakeit

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleCurrency() {
	Seed(11)
	currency := Currency()
	fmt.Printf("%s - %s", currency.Short, currency.Long)
	// Output: IQD - Iraq Dinar
}

func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency()
	}
}

func ExampleCurrencyShort() {
	Seed(11)
	fmt.Println(CurrencyShort())
	// Output: IQD
}

func BenchmarkCurrencyShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyShort()
	}
}

func ExampleCurrencyLong() {
	Seed(11)
	fmt.Println(CurrencyLong())
	// Output: Iraq Dinar
}

func BenchmarkCurrencyLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyLong()
	}
}

func ExamplePrice() {
	Seed(11)
	fmt.Printf("%.2f", Price(0.8618, 1000))
	// Output: 92.26
}

func BenchmarkPrice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Price(0, 1000)
	}
}

func ExampleCreditCard() {
	Seed(11)
	ccInfo := CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
	// Output: Visa
	// 6536459948995369
	// 03/27
	// 353
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
	// Output: 4136459948995369
}

func BenchmarkCreditCardNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardNumber()
	}
}

func ExampleCreditCardNumberLuhn() {
	Seed(11)
	fmt.Println(CreditCardNumberLuhn())
	// Output: 2720996615546177
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
	// Output: 01/22
}

func BenchmarkCreditCardExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardExp()
	}
}

func ExampleCreditCardCvv() {
	Seed(11)
	fmt.Println(CreditCardCvv())
	// Output: 513
}

func BenchmarkCreditCardCvv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardCvv()
	}
}
