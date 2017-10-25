package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSSN() {
	Seed(11)
	fmt.Println(SSN())
	// Output: 328-727-1570
}

func BenchmarkSSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SSN()
	}
}

func ExampleGender() {
	Seed(11)
	fmt.Println(Gender())
	// Output: female
}

func BenchmarkGender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gender()
	}
}

func ExamplePerson() {
	Seed(11)
	person := Person()
	job := person.Job
	address := person.Address
	contact := person.Contact
	creditCard := person.CreditCard

	fmt.Println(person.FirstName)
	fmt.Println(person.LastName)
	fmt.Println(person.Gender)
	fmt.Println(person.SSN)
	fmt.Println(person.Image)

	fmt.Println(job.Company)
	fmt.Println(job.Title)
	fmt.Println(job.Descriptor)
	fmt.Println(job.Level)

	fmt.Println(address.Address)
	fmt.Println(address.Street)
	fmt.Println(address.City)
	fmt.Println(address.State)
	fmt.Println(address.Zip)
	fmt.Println(address.Country)
	fmt.Println(address.Latitude)
	fmt.Println(address.Longitude)

	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)

	fmt.Println(creditCard.Type)
	fmt.Println(creditCard.Number)
	fmt.Println(creditCard.Exp)
	fmt.Println(creditCard.Cvv)

	// Output: Markus
	// Moen
	// male
	// 727-157-0245
	// http://lorempixel.com/300/300/people
	// Jacobi-Kuhic
	// Director
	// Forward
	// Accounts
	// 776 Roadshaven, Hilllville, Montana 58225
	// 776 Roadshaven
	// Hilllville
	// Montana
	// 58225
	// Rwanda
	// -24.329611640922778
	// -133.75991475181377
	// 748.033.6154
	// berniecebernier@halvorson.biz
	// Discover
	// 343324482755424
	// 04/15
	// 278
}

func BenchmarkPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Person()
	}
}
