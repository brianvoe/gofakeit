package gofakeit

import (
	"fmt"
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
	// Address: 35300 South Roads haven, Hilllville, Montana 30232
	// BS: e-enable
	// Beer Name: Weihenstephaner Hefeweissbier
	// Color: MidnightBlue
	// Company: Epsilon
	// Credit Card: 6375991714800889
	// Hacker Phrase: You can't quantify the application without navigating the bluetooth SMS microchip!
	// Job Title: Architect
	// Password: SHylpDU2nf9(,U*RaJKM1cOL47VY.BQw
}

func ExampleNew() {
	// Create new pseudo random faker struct and set initial seed
	fake := New(11)

	// All global functions are also available in the structs methods
	fmt.Println("Name:", fake.Name())
	fmt.Println("Email:", fake.Email())
	fmt.Println("Phone:", fake.Phone())
	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: 9948995369
}

func ExampleNewCrypto() {
	// Create new crypto faker struct
	fake := NewCrypto()

	// All global functions are also available in the structs methods
	fmt.Println("Name:", fake.Name())
	fmt.Println("Email:", fake.Email())
	fmt.Println("Phone:", fake.Phone())

	// Cannot output example as crypto/rand cant be predicted
}

func TestNewCrypto(t *testing.T) {
	// Create new crypto faker struct
	fake := NewCrypto()

	// All global functions are also available in the structs methods
	name := fake.Name()
	email := fake.Email()
	phone := fake.Phone()

	if name == "" || email == "" || phone == "" {
		t.Error("One of the values was empty")
	}
}

type customRand struct{}

func (c *customRand) Seed(seed int64) {}
func (c *customRand) Uint64() uint64  { return 8675309 }
func (c *customRand) Int63() int64    { return int64(c.Uint64() & ^uint64(1<<63)) }

func ExampleNewCustom() {
	// Setup stuct and methods required to meet interface needs
	// type customRand struct {}
	// func (c *customRand) Seed(seed int64) {}
	// func (c *customRand) Uint64() uint64 { return 8675309 }
	// func (c *customRand) Int63() int64 { return int64(c.Uint64() & ^uint64(1<<63)) }

	// Create new custom faker struct
	fake := NewCustom(&customRand{})

	// All global functions are also available in the structs methods
	fmt.Println("Name:", fake.Name())
	fmt.Println("Email:", fake.Email())
	fmt.Println("Phone:", fake.Phone())
	// Output:
	// Name: Aaliyah Abbott
	// Email: aaliyahabbott@abbott.com
	// Phone: 1000000000
}

func ExampleSetGlobalFaker() {
	cryptoFaker := NewCrypto()
	SetGlobalFaker(cryptoFaker)
}

func TestSetGlobalFaker(t *testing.T) {
	cryptoFaker := NewCrypto()
	SetGlobalFaker(cryptoFaker)

	name := Name()
	if name == "" {
		t.Error("Name was empty")
	}

	// Set global back to psuedo
	SetGlobalFaker(New(0))
}

func TestConcurrency(t *testing.T) {
	var setupComplete sync.WaitGroup
	setupComplete.Add(1)

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
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
