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

func ExampleFaker_Currency() {
	f := New(11)
	currency := f.Currency()
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

func ExampleFaker_CurrencyShort() {
	f := New(11)
	fmt.Println(f.CurrencyShort())
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

func ExampleFaker_CurrencyLong() {
	f := New(11)
	fmt.Println(f.CurrencyLong())
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

func ExampleFaker_Price() {
	f := New(11)
	fmt.Printf("%.2f", f.Price(0.8618, 1000))
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
	// Output:
	// UnionPay
	// 4364599489953698
	// 01/25
	// 300
}

func ExampleFaker_CreditCard() {
	f := New(11)
	ccInfo := f.CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)
	// Output:
	// UnionPay
	// 4364599489953698
	// 01/25
	// 300
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

func ExampleFaker_CreditCardType() {
	f := New(11)
	fmt.Println(f.CreditCardType())
	// Output: Visa
}

func BenchmarkCreditCardType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardType()
	}
}

func ExampleCreditCardNumber() {
	Seed(11)
	fmt.Println(CreditCardNumber(nil))
	fmt.Println(CreditCardNumber(&CreditCardOptions{Types: []string{"visa", "discover"}}))
	fmt.Println(CreditCardNumber(&CreditCardOptions{Bins: []string{"4111"}}))
	fmt.Println(CreditCardNumber(&CreditCardOptions{Gaps: true}))
	// Output:
	// 4136459948995375
	// 4635300425914586
	// 4111232020276132
	// 3054 800889 9827
}

func ExampleFaker_CreditCardNumber() {
	f := New(11)
	fmt.Println(f.CreditCardNumber(nil))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Types: []string{"visa", "discover"}}))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Bins: []string{"4111"}}))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Gaps: true}))
	// Output:
	// 4136459948995375
	// 4635300425914586
	// 4111232020276132
	// 3054 800889 9827
}

func TestCreditCardNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		if !isLuhn(CreditCardNumber(nil)) {
			t.Error("Number was not luhn")
		}
	}
}

func TestCreditCardNumberLookup(t *testing.T) {
	faker := New(0)
	info := GetFuncLookup("creditcardnumber")

	m := NewMapParams()
	m.Add("gaps", "true")

	_, err := info.Generate(faker.Rand, m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// t.Fatal(fmt.Sprintf("%s", value.(string)))
}

func BenchmarkCreditCardNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardNumber(nil)
	}
}

func TestIsLuhn(t *testing.T) {
	// Lets make sure this card is invalid
	if isLuhn("867gfsd5309") {
		t.Error("Card should have failed")
	}

	// Lets make sure this card is valid
	if !isLuhn("4716685826369360") {
		t.Error("Card should not have failed")
	}
}

func ExampleCreditCardExp() {
	Seed(11)
	fmt.Println(CreditCardExp())
	// Output: 01/23
}

func ExampleFaker_CreditCardExp() {
	f := New(11)
	fmt.Println(f.CreditCardExp())
	// Output: 01/23
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

func ExampleFaker_CreditCardCvv() {
	f := New(11)
	fmt.Println(f.CreditCardCvv())
	// Output: 513
}

func BenchmarkCreditCardCvv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardCvv()
	}
}

func ExampleAchRouting() {
	Seed(11)
	fmt.Println(AchRouting())
	// Output: 713645994
}

func ExampleFaker_AchRouting() {
	f := New(11)
	fmt.Println(f.AchRouting())
	// Output: 713645994
}

func BenchmarkAchRouting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AchRouting()
	}
}

func ExampleAchAccount() {
	Seed(11)
	fmt.Println(AchAccount())
	// Output: 413645994899
}

func ExampleFaker_AchAccount() {
	f := New(11)
	fmt.Println(f.AchAccount())
	// Output: 413645994899
}

func BenchmarkAchAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AchAccount()
	}
}

func ExampleBitcoinAddress() {
	Seed(11)
	fmt.Println(BitcoinAddress())
	// Output: 1lWLbxojXq6BqWX7X60VkcDIvYA
}

func ExampleFaker_BitcoinAddress() {
	f := New(11)
	fmt.Println(f.BitcoinAddress())
	// Output: 1lWLbxojXq6BqWX7X60VkcDIvYA
}

func BenchmarkBitcoinAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitcoinAddress()
	}
}

func ExampleBitcoinPrivateKey() {
	Seed(11)
	fmt.Println(BitcoinPrivateKey())
	// Output: 5KWjEJ7SnBNJyDjdPUjLuYByYzM9rG1trax8c2NTSBtv7YtR57v
}

func ExampleFaker_BitcoinPrivateKey() {
	f := New(11)
	fmt.Println(f.BitcoinPrivateKey())
	// Output: 5KWjEJ7SnBNJyDjdPUjLuYByYzM9rG1trax8c2NTSBtv7YtR57v
}

func BenchmarkBitcoinPrivateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitcoinPrivateKey()
	}
}
