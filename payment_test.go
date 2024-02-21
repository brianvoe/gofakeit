package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCurrency() {
	Seed(11)
	currency := Currency()
	fmt.Println(currency.Short)
	fmt.Println(currency.Long)

	// Output: UGX
	// Uganda Shilling
}

func ExampleFaker_Currency() {
	f := New(11)
	currency := f.Currency()
	fmt.Println(currency.Short)
	fmt.Println(currency.Long)

	// Output: UGX
	// Uganda Shilling
}

func BenchmarkCurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Currency()
	}
}

func ExampleCurrencyShort() {
	Seed(11)
	fmt.Println(CurrencyShort())

	// Output: UGX
}

func ExampleFaker_CurrencyShort() {
	f := New(11)
	fmt.Println(f.CurrencyShort())

	// Output: UGX
}

func BenchmarkCurrencyShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyShort()
	}
}

func ExampleCurrencyLong() {
	Seed(11)
	fmt.Println(CurrencyLong())

	// Output: Uganda Shilling
}

func ExampleFaker_CurrencyLong() {
	f := New(11)
	fmt.Println(f.CurrencyLong())

	// Output: Uganda Shilling
}

func BenchmarkCurrencyLong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrencyLong()
	}
}

func ExamplePrice() {
	Seed(11)
	fmt.Printf("%.2f", Price(0.8618, 1000))

	// Output: 770.49
}

func ExampleFaker_Price() {
	f := New(11)
	fmt.Printf("%.2f", f.Price(0.8618, 1000))

	// Output: 770.49
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

	// Output: American Express
	// 6376095989079994
	// 06/29
	// 125
}

func ExampleFaker_CreditCard() {
	f := New(11)
	ccInfo := f.CreditCard()
	fmt.Println(ccInfo.Type)
	fmt.Println(ccInfo.Number)
	fmt.Println(ccInfo.Exp)
	fmt.Println(ccInfo.Cvv)

	// Output: American Express
	// 6376095989079994
	// 06/29
	// 125
}

func BenchmarkCreditCard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCard()
	}
}

func ExampleCreditCardType() {
	Seed(11)
	fmt.Println(CreditCardType())

	// Output: Hiper
}

func ExampleFaker_CreditCardType() {
	f := New(11)
	fmt.Println(f.CreditCardType())

	// Output: Hiper
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

	// Output: 6376121252759896
	// 6449344737930519
	// 4111026894059156
	// 6706 2052 5709 6
}

func ExampleFaker_CreditCardNumber() {
	f := New(11)
	fmt.Println(f.CreditCardNumber(nil))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Types: []string{"visa", "discover"}}))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Bins: []string{"4111"}}))
	fmt.Println(f.CreditCardNumber(&CreditCardOptions{Gaps: true}))

	// Output: 6376121252759896
	// 6449344737930519
	// 4111026894059156
	// 6706 2052 5709 6
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

	_, err := info.Generate(faker, m, info)
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

	// Output: 11/33
}

func ExampleFaker_CreditCardExp() {
	f := New(11)
	fmt.Println(f.CreditCardExp())

	// Output: 11/33
}

func BenchmarkCreditCardExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardExp()
	}
}

func ExampleCreditCardCvv() {
	Seed(11)
	fmt.Println(CreditCardCvv())

	// Output: 881
}

func ExampleFaker_CreditCardCvv() {
	f := New(11)
	fmt.Println(f.CreditCardCvv())

	// Output: 881
}

func BenchmarkCreditCardCvv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreditCardCvv()
	}
}

func ExampleAchRouting() {
	Seed(11)
	fmt.Println(AchRouting())

	// Output: 881252759
}

func ExampleFaker_AchRouting() {
	f := New(11)
	fmt.Println(f.AchRouting())

	// Output: 881252759
}

func BenchmarkAchRouting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AchRouting()
	}
}

func ExampleAchAccount() {
	Seed(11)
	fmt.Println(AchAccount())

	// Output: 881252759890
}

func ExampleFaker_AchAccount() {
	f := New(11)
	fmt.Println(f.AchAccount())

	// Output: 881252759890
}

func BenchmarkAchAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AchAccount()
	}
}

func ExampleBitcoinAddress() {
	Seed(11)
	fmt.Println(BitcoinAddress())

	// Output: 13blsBo8bffq7a35c5nwLT4eXWu0pReLF1
}

func ExampleFaker_BitcoinAddress() {
	f := New(11)
	fmt.Println(f.BitcoinAddress())

	// Output: 13blsBo8bffq7a35c5nwLT4eXWu0pReLF1
}

func BenchmarkBitcoinAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitcoinAddress()
	}
}

func ExampleBitcoinPrivateKey() {
	Seed(11)
	fmt.Println(BitcoinPrivateKey())

	// Output: 5JMZxkQX2PgaasaHc8wnWLNdMu7rxeU7xS64ev7RWNinacicPfm
}

func ExampleFaker_BitcoinPrivateKey() {
	f := New(11)
	fmt.Println(f.BitcoinPrivateKey())

	// Output: 5JMZxkQX2PgaasaHc8wnWLNdMu7rxeU7xS64ev7RWNinacicPfm
}

func BenchmarkBitcoinPrivateKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitcoinPrivateKey()
	}
}
