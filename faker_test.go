package gofakeit

import (
	"fmt"
	"testing"
)

func Example() {
	Seed(11)
	fmt.Println("Name:", Name())
	fmt.Println("Email:", Email())
	fmt.Println("Phone:", Phone())
	fmt.Println("Address:", Address().Address)
	fmt.Println("BS:", BS())
	fmt.Println("Beer Name:", BeerName())
	fmt.Println("Color:", Color())
	fmt.Println("Company:", Company())
	fmt.Println("Credit Card:", CreditCardNumber())
	fmt.Println("Hacker Phrase:", HackerPhrase())
	fmt.Println("Job Title:", JobTitle())
	fmt.Println("Password:", Password(true, true, true, true, false, 32))
	currency := Currency()
	fmt.Printf("Currency: %s - %s", currency.Short, currency.Long)
	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 9948995369
	// Address: 35300 South Roadshaven, Hilllville, Montana 30232
	// BS: e-enable
	// Beer Name: Weihenstephaner Hefeweissbier
	// Color: MidnightBlue
	// Company: Epsilon
	// Credit Card: 6521714800889982
	// Hacker Phrase: N/A the sensor won't do anything, we need to compress the open-source EXE program!
	// Job Title: Producer
	// Password: gzDGWlrJgsM?npd7*Yyd9FmCcE%SKX7y
	// Currency: MYR - Malaysia Ringgit
}

func TestSeed(t *testing.T) {
	Seed(0)
}
