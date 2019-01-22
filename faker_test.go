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
	//Email: alaynawuckert@kozey.biz
	//Phone: 9948995369
	//Address: 35300 South Roadshaven, Hilllville, Montana 30232
	//BS: e-enable
	//Beer Name: Weihenstephaner Hefeweissbier
	//Color: MidnightBlue
	//Company: Heaney-Tromp
	//Credit Card: 2720997148008899
	//Hacker Phrase: Use the wireless SQL sensor, then you can compress the open-source array!
	//Job Title: Producer
	//Password: Hozsav7LgAa -Q$qSu  L*2 ?Jivzjn2
	//Currency: DJF - Djibouti Franc
}

func TestSeed(t *testing.T) {
	Seed(0)
}
