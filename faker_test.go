package gofakeit

import (
	"fmt"
	"math/rand/v2"
	"sync"
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
	fake := New(11)

	// All global functions are also available in the structs methods
	fmt.Println("Name:", fake.Name())
	fmt.Println("Email:", fake.Email())
	fmt.Println("Phone:", fake.Phone())

	// Output:
	// Name: Sonny Stiedemann
	// Email: codydonnelly@leannon.biz
	// Phone: 7598907999
}

func ExampleNewFaker() {
	// Create new faker with ChaCha8, cryptographically secure
	chacha := rand.NewChaCha8([32]byte{5, 4, 3, 2, 1, 0})
	fake := NewFaker(chacha, true)

	// or

	// Create new faker with PCG, pseudo-random
	pcg := rand.NewPCG(0, 0)
	fake = NewFaker(pcg, false)

	fmt.Println("Name:", fake.Name())

	// Output:
	// Name: Damian Pagac
}

func TestSeed(t *testing.T) {
	// Test crypto that has no parameters in Seed
	GlobalFaker = New(11)

	// Test a simple function
	name := Name()

	// Seed
	err := Seed(11)
	if err != nil {
		t.Error(err)
	}

	// Make sure name is the same
	if name != Name() {
		t.Error("Name was different after seed")
	}
}

func TestSetGlobalFaker(t *testing.T) {
	// Set global to crypto
	crypto := rand.NewPCG(11, 11)
	GlobalFaker = NewFaker(crypto, true)

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
