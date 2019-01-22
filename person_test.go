package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleSSN() {
	Seed(11)
	fmt.Println(SSN())
	// Output: 296446360
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
	// 420776036
	// https://picsum.photos/300/300/people
	// Lockman and Sons
	// Developer
	// Global
	// Brand
	// 5369 Streamville, Rossieview, Hawaii 42591
	// 5369 Streamville
	// Rossieview
	// Hawaii
	// 42591
	// Burkina Faso
	// -6.662594491850811
	// 23.921575244414612
	// 3023202027
	// lamarkoelpin@heaney.biz
	// Discover
	// 4148008899822720
	// 09/29
	// 932
}

func BenchmarkPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Person()
	}
}
