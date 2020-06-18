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
	fmt.Println("Credit Card:", CreditCardNumber(nil))
	fmt.Println("Hacker Phrase:", HackerPhrase())
	fmt.Println("Job Title:", JobTitle())
	fmt.Println("Password:", Password(true, true, true, true, false, 32))
	currency := Currency()
	fmt.Printf("Currency: %s - %s", currency.Short, currency.Long)
	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 9948995369
	// Address: 35300 South Roads haven, Hilllville, Montana 30232
	// BS: e-enable
	// Beer Name: Weihenstephaner Hefeweissbier
	// Color: MidnightBlue
	// Company: Epsilon
	// Credit Card: 6375991714800889
	// Hacker Phrase: You can't quantify the application without navigating the bluetooth SMS microchip!
	// Job Title: Architect
	// Password: SHylpDU2nf9(,U*RaJKM1cOL47VY.BQw
	// Currency: KES - Kenya Shilling
}

func TestSeed(t *testing.T) {
	Seed(0)
}
