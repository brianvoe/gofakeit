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
	// Phone: (570)245-7485
	// Address: 75776 Lake Viewland, Sterlingstad, New Hampshire 82250-2868
	// BS: bleeding-edge
	// Beer Name: Founders Kentucky Breakfast
	// Color: Orchid
	// Company: Hintz, Crooks and McKenzie
	// Credit Card: 4071613324482755
	// Hacker Phrase: If we calculate the capacitor, we can get to the THX port through the back-end SSL matrix!
	// Job Title: Manager
	// Password: E1&U_#fm&89QExx$X%83qC-$9Z2qPA2t
	// Currency: INR - India Rupee
}

func TestSeed(t *testing.T) {
	Seed(0)
}
