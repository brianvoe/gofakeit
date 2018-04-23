package gofakeit

import (
	"fmt"
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
