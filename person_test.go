package gofakeit

import (
	"fmt"
	"net/mail"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func ExampleName() {
	Seed(11)
	fmt.Println(Name())

	// Output: Priscilla Thornton
}

func ExampleFaker_Name() {
	f := New(11)
	fmt.Println(f.Name())

	// Output: Priscilla Thornton
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name()
	}
}

func ExampleFirstName() {
	Seed(11)
	fmt.Println(FirstName())

	// Output: Priscilla
}

func ExampleFaker_FirstName() {
	f := New(11)
	fmt.Println(f.FirstName())

	// Output: Priscilla
}

func BenchmarkFirstName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstName()
	}
}

func ExampleMiddleName() {
	Seed(11)
	fmt.Println(MiddleName())

	// Output: Jose
}

func ExampleFaker_MiddleName() {
	f := New(11)
	fmt.Println(f.MiddleName())

	// Output: Jose
}

func BenchmarkMiddleName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MiddleName()
	}
}

func ExampleLastName() {
	Seed(11)
	fmt.Println(LastName())

	// Output: Wagner
}

func ExampleFaker_LastName() {
	f := New(11)
	fmt.Println(f.LastName())

	// Output: Wagner
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

	// Output: 806851996
}

func ExampleFaker_SSN() {
	f := New(11)
	fmt.Println(f.SSN())

	// Output: 806851996
}

func TestSSN(t *testing.T) {
	f := New(11)
	format := regexp.MustCompile(`^\d{9}$`)

	for range 10000 {
		ssn := f.SSN()
		if !format.MatchString(ssn) {
			t.Fatalf("SSN %q should be 9 digits", ssn)
		}

		// The SSA does not issue area numbers 000, 666 and 900 to 999, group 00 or serial 0000
		area, group, serial := ssn[:3], ssn[3:5], ssn[5:]
		if area == "000" || area == "666" || area[0] == '9' {
			t.Fatalf("SSN %q has an area number that is never issued", ssn)
		}
		if group == "00" {
			t.Fatalf("SSN %q has group 00", ssn)
		}
		if serial == "0000" {
			t.Fatalf("SSN %q has serial 0000", ssn)
		}
	}
}

func BenchmarkSSN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SSN()
	}
}

func ExampleEIN() {
	Seed(11)
	fmt.Println(EIN())

	// Output: 88-8125275
}

func ExampleFaker_EIN() {
	f := New(11)
	fmt.Println(f.EIN())

	// Output: 88-8125275
}

func TestEIN(t *testing.T) {
	ein := EIN()
	if len(ein) != 10 {
		t.Errorf("EIN should be 10 characters long, got %d", len(ein))
	}
	if ein[2] != '-' {
		t.Errorf("EIN should have a dash at position 2, got %c", ein[2])
	}

	// Check that it's all digits except for the dash
	for i, char := range ein {
		if i == 2 {
			if char != '-' {
				t.Errorf("EIN should have a dash at position 2, got %c", char)
			}
		} else {
			if char < '0' || char > '9' {
				t.Errorf("EIN should contain only digits (except dash), got %c", char)
			}
		}
	}
}

func TestEINPrefix(t *testing.T) {
	f := New(11)

	// The IRS assigns 01 to 06, 10 to 16, 20 to 27, 30 to 48, 50 to 68, 71 to 77, 80 to 88, 90 to 95, 98 and 99
	valid := regexp.MustCompile(`^(0[1-6]|1[0-6]|2[0-7]|3[0-9]|4[0-8]|5[0-9]|6[0-8]|7[1-7]|8[0-8]|9[0-5]|9[89])-\d{7}$`)
	for range 10000 {
		ein := f.EIN()
		if !valid.MatchString(ein) {
			t.Fatalf("EIN %q does not start with a prefix the IRS assigns", ein)
		}
	}
}

func BenchmarkEIN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EIN()
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

func ExampleAge() {
	Seed(11)
	fmt.Println(Age())

	// Output: 90
}

func ExampleFaker_Age() {
	f := New(11)
	fmt.Println(f.Age())

	// Output: 90
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Age()
	}
}

func ExampleEthnicity() {
	Seed(11)
	fmt.Println(Ethnicity())

	// Output: Swiss
}

func ExampleFaker_Ethnicity() {
	f := New(11)
	fmt.Println(f.Ethnicity())

	// Output: Swiss
}

func BenchmarkEthnicity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ethnicity()
	}
}

func ExampleHobby() {
	Seed(11)
	fmt.Println(Hobby())

	// Output: Travel
}

func ExampleFaker_Hobby() {
	f := New(11)
	fmt.Println(f.Hobby())

	// Output: Travel
}

func BenchmarkHobby(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hobby()
	}
}

func ExampleSocialMedia() {
	Seed(11)
	fmt.Println(SocialMedia())

	// Output: https://linkedin.com/in/CuriosCasino7
}

func ExampleFaker_SocialMedia() {
	f := New(11)
	fmt.Println(f.SocialMedia())

	// Output: https://linkedin.com/in/CuriosCasino7
}

func BenchmarkSocialMedia(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SocialMedia()
	}
}

func ExampleBio() {
	Seed(11)
	fmt.Println(Bio())
	fmt.Println(Bio())
	fmt.Println(Bio())

	// Output: Brews Fruit Beer while streaming Country, day job in Memphis as Musician
	// Singer with a soft spot for Diesel builds and hands-on Weaving in Laredo
	// Known for a hoodie look and a shelf of PaleGreen weasel art
}

func ExampleFaker_Bio() {
	f := New(11)
	fmt.Println(f.Bio())
	fmt.Println(f.Bio())
	fmt.Println(f.Bio())

	// Output: Brews Fruit Beer while streaming Country, day job in Memphis as Musician
	// Singer with a soft spot for Diesel builds and hands-on Weaving in Laredo
	// Known for a hoodie look and a shelf of PaleGreen weasel art
}

func TestBio(t *testing.T) {
	Seed(11)

	bio1 := Bio()
	bio2 := Bio()

	if bio1 == "" {
		t.Error("Bio() returned empty string")
	}

	if bio2 == "" {
		t.Error("Bio() returned empty string")
	}

	if bio1 == bio2 {
		t.Error("Bio() returned same result for different calls")
	}

	// Test that bio contains some expected elements from templates
	if len(bio1) < 10 {
		t.Error("Bio() returned too short result")
	}
}

// Test to make sure all curly braces are replaced
func TestBioReplace(t *testing.T) {
	for range 100000 {
		bio := Bio()
		if strings.Contains(bio, "{") || strings.Contains(bio, "}") {
			t.Error("Bio contains { or }")
		}
	}
}

func BenchmarkBio(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bio()
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
	fmt.Println(person.Age)
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

	// Output: Priscilla
	// Thornton
	// male
	// 20
	// 475237266
	// Lacrosse
	// Weather Channel
	// Substitute Teacher
	// Strategic
	// Applications
	// 993 Lanebury, Lincoln, Kansas 93050
	// 993 Lanebury
	// Lincoln
	// Kansas
	// 93050
	// India
	// -26.936948
	// -28.374174
	// 5643954059
	// april_muller@joseph.net
	// Hipercard
	// 6062825257093882
	// 10/29
	// 162
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
	fmt.Println(person.Age)
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

	// Output: Priscilla
	// Thornton
	// male
	// 20
	// 475237266
	// Lacrosse
	// Weather Channel
	// Substitute Teacher
	// Strategic
	// Applications
	// 993 Lanebury, Lincoln, Kansas 93050
	// 993 Lanebury
	// Lincoln
	// Kansas
	// 93050
	// India
	// -26.936948
	// -28.374174
	// 5643954059
	// april_muller@joseph.net
	// Hipercard
	// 6062825257093882
	// 10/29
	// 162
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

	// Output: 6857205275
	// sethsoto45@wallace.com
}

func ExampleFaker_Contact() {
	f := New(11)
	contact := f.Contact()
	fmt.Println(contact.Phone)
	fmt.Println(contact.Email)

	// Output: 6857205275
	// sethsoto45@wallace.com
}

func BenchmarkContact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Contact()
	}
}

func ExamplePhone() {
	Seed(11)
	fmt.Println(Phone())

	// Output: 6857205275
}

func ExampleFaker_Phone() {
	f := New(11)
	fmt.Println(f.Phone())

	// Output: 6857205275
}

// checkNANPDigits checks the ten digits of a phone number against the North American numbering plan:
// the area code and the exchange start with 2 to 9 and neither of them is N11
func checkNANPDigits(t *testing.T, phone, digits string) {
	t.Helper()

	if len(digits) != 10 {
		t.Fatalf("Phone %q should have 10 digits", phone)
	}
	for _, part := range []string{digits[:3], digits[3:6]} {
		if part[0] < '2' || part[1:] == "11" {
			t.Fatalf("Phone %q has the area code or exchange %s, which is not used", phone, part)
		}
	}
}

func TestPhone(t *testing.T) {
	f := New(11)

	for range 10000 {
		phone := f.Phone()
		checkNANPDigits(t, phone, phone)
	}
}

func BenchmarkPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Phone()
	}
}

func ExamplePhoneFormatted() {
	Seed(11)
	fmt.Println(PhoneFormatted())

	// Output: 420-553-2759
}

func ExampleFaker_PhoneFormatted() {
	f := New(11)
	fmt.Println(f.PhoneFormatted())

	// Output: 420-553-2759
}

func TestPhoneFormatted(t *testing.T) {
	f := New(11)
	formats := []string{"###-###-####", "(###)###-####", "###.###.####"}
	digit := regexp.MustCompile(`\d`)
	nonDigit := regexp.MustCompile(`\D`)

	for range 10000 {
		phone := f.PhoneFormatted()

		// 1-###-###-#### is one of the formats above with the country code in front
		number := strings.TrimPrefix(phone, "1-")
		if shape := digit.ReplaceAllString(number, "#"); !slices.Contains(formats, shape) {
			t.Fatalf("Phone %q is not in one of the formats %v", phone, formats)
		}

		checkNANPDigits(t, phone, nonDigit.ReplaceAllString(number, ""))
	}
}

func BenchmarkPhoneFormatted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PhoneFormatted()
	}
}

func TestEmail(t *testing.T) {
	for i := range 100000 {
		e := Email()
		if e == "" {
			t.Fatalf("empty email at iteration %d", i)
		}

		addr, err := mail.ParseAddress(e)
		if err != nil || addr.Address != e {
			t.Fatalf("invalid email at iteration %d: %q err=%v", i, e, err)
		}

		at := strings.LastIndex(e, "@")
		if at <= 0 || at >= len(e)-1 {
			t.Fatalf("invalid email at iteration %d: %q", i, e)
		}

		local := e[:at]
		domain := e[at+1:]
		if len(local) > 64 || len(e) > 254 {
			t.Fatalf("invalid email at iteration %d: %q", i, e)
		}
		if strings.Contains(local, "..") || strings.Contains(domain, "..") {
			t.Fatalf("invalid email at iteration %d: %q", i, e)
		}
		if strings.HasPrefix(domain, ".") || strings.HasSuffix(domain, ".") {
			t.Fatalf("invalid email at iteration %d: %q", i, e)
		}
		if !strings.Contains(domain, ".") {
			t.Fatalf("invalid email at iteration %d: %q", i, e)
		}
		for j := 0; j < len(e); j++ {
			if e[j] <= ' ' {
				t.Fatalf("invalid email at iteration %d: %q", i, e)
			}
		}
	}
}

func ExampleEmail() {
	Seed(11)
	fmt.Println(Email())

	// Output: priscilla.thornton@duncan.biz
}

func ExampleFaker_Email() {
	f := New(11)
	fmt.Println(f.Email())

	// Output: priscilla.thornton@duncan.biz
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
