package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleName() {
	Seed(11)
	fmt.Println(Name())

	// Output: Sonny Stiedemann
}

func ExampleFaker_Name() {
	f := New(11)
	fmt.Println(f.Name())

	// Output: Sonny Stiedemann
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}

func ExampleFirstName() {
	Seed(11)
	fmt.Println(FirstName())

	// Output: Sonny
}

func ExampleFaker_FirstName() {
	f := New(11)
	fmt.Println(f.FirstName())

	// Output: Sonny
}

func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

func ExampleMiddleName() {
	Seed(11)
	fmt.Println(MiddleName())

	// Output: Star
}

func ExampleFaker_MiddleName() {
	f := New(11)
	fmt.Println(f.MiddleName())

	// Output: Star
}

func BenchmarkMiddleName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiddleName()
	}
}

func ExampleLastName() {
	Seed(11)
	fmt.Println(LastName())

	// Output: Treutel
}

func ExampleFaker_LastName() {
	f := New(11)
	fmt.Println(f.LastName())

	// Output: Treutel
}

func BenchmarkLastName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LastName()
	}
}

func ExampleNamePrefix() {
	Seed(11)
	fmt.Println(NamePrefix())

	// Output: Dr.
}

func ExampleFaker_NamePrefix() {
	f := New(11)
	fmt.Println(f.NamePrefix())

	// Output: Dr.
}

func BenchmarkNamePrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NamePrefix()
	}
}

func ExampleNameSuffix() {
	Seed(11)
	fmt.Println(NameSuffix())

	// Output: PhD
}

func ExampleFaker_NameSuffix() {
	f := New(11)
	fmt.Println(f.NameSuffix())

	// Output: PhD
}

func BenchmarkNameSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NameSuffix()
	}
}

func ExampleSSN() {
	Seed(11)
	fmt.Println(SSN())

	// Output: 906295542
}

func ExampleFaker_SSN() {
	f := New(11)
	fmt.Println(f.SSN())

	// Output: 906295542
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

func ExampleHobby() {
	Seed(11)
	fmt.Println(Hobby())

	// Output: Marching band
}

func ExampleFaker_Hobby() {
	f := New(11)
	fmt.Println(f.Hobby())

	// Output: Marching band
}

func BenchmarkHobby(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hobby()
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
	fmt.Println(person.Hobby)

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

	// Output: Sonny
	// Stiedemann
	// male
	// 280254464
	// Sailing
	// DataLogix
	// Planner
	// Human
	// Usability
	// 679 Underpassborough, Omaha, Massachusetts 37930
	// 679 Underpassborough
	// Omaha
	// Massachusetts
	// 37930
	// North Macedonia
	// -0.877085
	// 83.264578
	// 4102689405
	// verdabrakus@mayert.name
	// American Express
	// 562570938760
	// 04/25
	// 906
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
	fmt.Println(person.Hobby)

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

	// Output: Sonny
	// Stiedemann
	// male
	// 280254464
	// Sailing
	// DataLogix
	// Planner
	// Human
	// Usability
	// 679 Underpassborough, Omaha, Massachusetts 37930
	// 679 Underpassborough
	// Omaha
	// Massachusetts
	// 37930
	// North Macedonia
	// -0.877085
	// 83.264578
	// 4102689405
	// verdabrakus@mayert.name
	// American Express
	// 562570938760
	// 04/25
	// 906
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

	// Output: 8812527598
	// stevebins@robel.io
}

func ExampleFaker_Contact() {
	f := New(11)
	contact := f.Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)

	// Output: 8812527598
	// stevebins@robel.io
}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contact()
	}
}

func ExamplePhone() {
	Seed(11)
	fmt.Println(Phone())

	// Output: 8812527598
}

func ExampleFaker_Phone() {
	f := New(11)
	fmt.Println(f.Phone())

	// Output: 8812527598
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExamplePhoneFormatted() {
	Seed(11)
	fmt.Println(PhoneFormatted())

	// Output: 812-527-5989
}

func ExampleFaker_PhoneFormatted() {
	f := New(11)
	fmt.Println(f.PhoneFormatted())

	// Output: 812-527-5989
}

func BenchmarkPhoneFormatted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneFormatted()
	}
}

func ExampleEmail() {
	Seed(11)
	fmt.Println(Email())

	// Output: sonnystiedemann@donnelly.biz
}

func ExampleFaker_Email() {
	f := New(11)
	fmt.Println(f.Email())

	// Output: sonnystiedemann@donnelly.biz
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

	// Output: map[Team 1:[Jeff Robert Billy] Team 2:[Connor Fabian Sharon] Team 3:[Justin Steve]]
}

func ExampleFaker_Teams() {
	f := New(11)
	fmt.Println(f.Teams(
		[]string{"Billy", "Sharon", "Jeff", "Connor", "Steve", "Justin", "Fabian", "Robert"},
		[]string{"Team 1", "Team 2", "Team 3"},
	))

	// Output: map[Team 1:[Jeff Robert Billy] Team 2:[Connor Fabian Sharon] Team 3:[Justin Steve]]
}

func BenchmarkTeams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Teams(
			[]string{"Billy", "Sharon", "Jeff", "Connor", "Steve", "Justin", "Fabian", "Robert"},
			[]string{"Team 1", "Team 2", "Team 3"},
		)
	}
}
