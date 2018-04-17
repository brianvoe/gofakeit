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
	fmt.Println("Password:", Password(true, true, true, true, true, 32))
	currency := Currency()
	fmt.Printf("Currency: %s - %s", currency.Short, currency.Long)
	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 1570245748
	// Address: 67577 South Roadshaven, Hilllville, Montana 58225
	// BS: sexy
	// Beer Name: Hennepin
	// Color: PaleGoldenRod
	// Company: Tromp Group
	// Credit Card: 4615440716133244
	// Hacker Phrase: If we calculate the protocol, we can get to the JSON pixel through the open-source JSON capacitor!
	// Job Title: Assistant
	// Password: Ec_ YE1&U_#fm&89QExx$X%83qC-$9Z2
	// Currency: ZMW - Zambia Kwacha
}

func TestSeed(t *testing.T) {
	Seed(0)
}
