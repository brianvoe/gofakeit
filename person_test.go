package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleName() {
	Seed(11)
	fmt.Println(Name())
	// Output: Markus Moen
}

func ExampleFaker_Name() {
	f := New(11)
	fmt.Println(f.Name())
	// Output: Markus Moen
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}

func ExampleFirstName() {
	Seed(11)
	fmt.Println(FirstName())
	// Output: Markus
}

func ExampleFaker_FirstName() {
	f := New(11)
	fmt.Println(f.FirstName())
	// Output: Markus
}

func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

func ExampleLastName() {
	Seed(11)
	fmt.Println(LastName())
	// Output: Daniel
}

func ExampleFaker_LastName() {
	f := New(11)
	fmt.Println(f.LastName())
	// Output: Daniel
}

func BenchmarkLastName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastName()
	}
}

func ExampleNamePrefix() {
	Seed(11)
	fmt.Println(NamePrefix())
	// Output: Mr.
}

func ExampleFaker_NamePrefix() {
	f := New(11)
	fmt.Println(f.NamePrefix())
	// Output: Mr.
}

func BenchmarkNamePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NamePrefix()
	}
}

func ExampleNameSuffix() {
	Seed(11)
	fmt.Println(NameSuffix())
	// Output: Jr.
}

func ExampleFaker_NameSuffix() {
	f := New(11)
	fmt.Println(f.NameSuffix())
	// Output: Jr.
}

func BenchmarkNameSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NameSuffix()
	}
}

func ExampleSSN() {
	Seed(11)
	fmt.Println(SSN())
	// Output: 296446360
}

func ExampleFaker_SSN() {
	f := New(11)
	fmt.Println(f.SSN())
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

func ExampleFaker_Gender() {
	f := New(11)
	fmt.Println(f.Gender())
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

	// Output:
	// Markus
	// Moen
	// male
	// 420776036
	// https://picsum.photos/300/300/people
	// Morgan Stanley
	// Associate
	// Human
	// Usability
	// 99536 North Stream ville, Rossieview, Hawaii 42591
	// 99536 North Stream ville
	// Rossieview
	// Hawaii
	// 42591
	// Ghana
	// -6.662595
	// 23.921575
	// 3023202027
	// lamarkoelpin@heaney.biz
	// Maestro
	// 39800889982276
	// 05/30
	// 932
}

func ExampleFaker_Person() {
	f := New(11)
	person := f.Person()
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

	// Output:
	// Markus
	// Moen
	// male
	// 420776036
	// https://picsum.photos/300/300/people
	// Morgan Stanley
	// Associate
	// Human
	// Usability
	// 99536 North Stream ville, Rossieview, Hawaii 42591
	// 99536 North Stream ville
	// Rossieview
	// Hawaii
	// 42591
	// Ghana
	// -6.662595
	// 23.921575
	// 3023202027
	// lamarkoelpin@heaney.biz
	// Maestro
	// 39800889982276
	// 05/30
	// 932
}

func BenchmarkPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Person()
	}
}

func ExampleContact() {
	Seed(11)
	contact := Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)
	// Output: 6136459948
	// carolecarroll@bosco.com
}

func ExampleFaker_Contact() {
	f := New(11)
	contact := f.Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)
	// Output: 6136459948
	// carolecarroll@bosco.com
}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contact()
	}
}

func ExamplePhone() {
	Seed(11)
	fmt.Println(Phone())
	// Output: 6136459948
}

func ExampleFaker_Phone() {
	f := New(11)
	fmt.Println(f.Phone())
	// Output: 6136459948
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExamplePhoneFormatted() {
	Seed(11)
	fmt.Println(PhoneFormatted())
	// Output: 136-459-9489
}

func ExampleFaker_PhoneFormatted() {
	f := New(11)
	fmt.Println(f.PhoneFormatted())
	// Output: 136-459-9489
}

func BenchmarkPhoneFormatted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneFormatted()
	}
}

func ExampleEmail() {
	Seed(11)
	fmt.Println(Email())
	// Output: markusmoen@pagac.net
}

func ExampleFaker_Email() {
	f := New(11)
	fmt.Println(f.Email())
	// Output: markusmoen@pagac.net
}

func BenchmarkEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Email()
	}
}

func ExampleTeams() {
	Seed(11)
	fmt.Println(Teams(
		[]string{"Billy", "Sharon", "Jeff", "Connor", "Steve", "Justin", "Fabian", "Robert"},
		[]string{"Team 1", "Team 2", "Team 3"},
	))
	// Output: map[Team 1:[Fabian Connor Steve] Team 2:[Jeff Sharon Justin] Team 3:[Robert Billy]]
}

func ExampleFaker_Teams() {
	f := New(11)
	fmt.Println(f.Teams(
		[]string{"Billy", "Sharon", "Jeff", "Connor", "Steve", "Justin", "Fabian", "Robert"},
		[]string{"Team 1", "Team 2", "Team 3"},
	))
	// Output: map[Team 1:[Fabian Connor Steve] Team 2:[Jeff Sharon Justin] Team 3:[Robert Billy]]
}

func BenchmarkTeams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Teams(
			[]string{"Billy", "Sharon", "Jeff", "Connor", "Steve", "Justin", "Fabian", "Robert"},
			[]string{"Team 1", "Team 2", "Team 3"},
		)
	}
}
