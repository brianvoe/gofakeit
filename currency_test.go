package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCurrency() {
	Seed(11)
	fmt.Println(Currency())
}

func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency()
	}
}

func ExampleCurrencyShort() {
	Seed(11)
	fmt.Println(CurrencyShort())
}

func BenchmarkCurrencyShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyShort()
	}
}

func ExampleCurrencyLong() {
	Seed(11)
	fmt.Println(CurrencyLong())
}

func BenchmarkCurrencyLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyLong()
	}
}
