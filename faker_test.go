package gofakeit

import (
	"fmt"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v6/source"
)

func TestSeed(t *testing.T) {
	// Test crypto that has no parameters in Seed
	GlobalFaker = New(0)
	Seed(11)

	// Will panic if Seed couldn't be called
	t.Fatal("Seed failed")
}

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

	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 9948995369
	// Address: 35300 South Roadshaven, Miami, Tennessee 58302
	// BS: streamline
	// Beer Name: Pliny The Elder
	// Color: Gray
	// Company: Center for Responsive Politics
	// Credit Card: 3821714800889989
	// Hacker Phrase: Overriding the capacitor won't do anything, we need to compress the online SMTP protocol!
	// Job Title: Supervisor
	// Password: #8L79W6s4E9jT2Q047??YkyD0xxnC2#u
}

func ExampleNew() {
	// Get new faker with default settings
	fake := New(0)

	// or

	// Create new faker with ChaCha8
	chacha := source.NewChaCha8([32]byte{5, 4, 3, 2, 1, 0}, true)
	fake = NewFaker(chacha)

	// All global functions are also available in the structs methods
	fmt.Println("Name:", fake.Name())
	fmt.Println("Email:", fake.Email())
	fmt.Println("Phone:", fake.Phone())

	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 9948995369
}

func TestSetGlobalFaker(t *testing.T) {
	// Set global to crypto
	cryptoFaker := source.NewCrypto(true)
	GlobalFaker = NewFaker(cryptoFaker)

	// Test a simple function
	name := Name()
	if name == "" {
		t.Error("Name was empty")
	}

	// Set global back to default
	GlobalFaker = New(0)
}

func TestConcurrency(t *testing.T) {
	var setupComplete sync.WaitGroup
	setupComplete.Add(1)

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			setupComplete.Wait()
			Paragraph(1, 5, 20, " ")
			wg.Done()
		}()
	}

	setupComplete.Done()
	wg.Wait()
}
