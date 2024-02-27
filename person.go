package gofakeit

import (
	"math"
	"strconv"
	"strings"
)

// PersonInfo is a struct of person information
type PersonInfo struct {
	FirstName  string          `json:"first_name" xml:"first_name"`
	LastName   string          `json:"last_name" xml:"last_name"`
	Gender     string          `json:"gender" xml:"gender"`
	SSN        string          `json:"ssn" xml:"ssn"`
	Hobby      string          `json:"hobby" xml:"hobby"`
	Job        *JobInfo        `json:"job" xml:"job"`
	Address    *AddressInfo    `json:"address" xml:"address"`
	Contact    *ContactInfo    `json:"contact" xml:"contact"`
	CreditCard *CreditCardInfo `json:"credit_card" xml:"credit_card"`
}

// Person will generate a struct with person information
func Person() *PersonInfo { return person(GlobalFaker) }

// Person will generate a struct with person information
func (f *Faker) Person() *PersonInfo { return person(f) }

func person(f *Faker) *PersonInfo {
	return &PersonInfo{
		FirstName:  firstName(f),
		LastName:   lastName(f),
		Gender:     gender(f),
		SSN:        ssn(f),
		Hobby:      hobby(f),
		Job:        job(f),
		Address:    address(f),
		Contact:    contact(f),
		CreditCard: creditCard(f),
	}
}

// Name will generate a random First and Last Name
func Name() string { return name(GlobalFaker) }

// Name will generate a random First and Last Name
func (f *Faker) Name() string { return name(f) }

func name(f *Faker) string {
	return getRandValue(f, []string{"person", "first"}) + " " + getRandValue(f, []string{"person", "last"})
}

// FirstName will generate a random first name
func FirstName() string { return firstName(GlobalFaker) }

// FirstName will generate a random first name
func (f *Faker) FirstName() string { return firstName(f) }

func firstName(f *Faker) string { return getRandValue(f, []string{"person", "first"}) }

// MiddleName will generate a random middle name
func MiddleName() string { return middleName(GlobalFaker) }

// MiddleName will generate a random middle name
func (f *Faker) MiddleName() string { return middleName(f) }

func middleName(f *Faker) string { return getRandValue(f, []string{"person", "middle"}) }

// LastName will generate a random last name
func LastName() string { return lastName(GlobalFaker) }

// LastName will generate a random last name
func (f *Faker) LastName() string { return lastName(f) }

func lastName(f *Faker) string { return getRandValue(f, []string{"person", "last"}) }

// NamePrefix will generate a random name prefix
func NamePrefix() string { return namePrefix(GlobalFaker) }

// NamePrefix will generate a random name prefix
func (f *Faker) NamePrefix() string { return namePrefix(f) }

func namePrefix(f *Faker) string { return getRandValue(f, []string{"person", "prefix"}) }

// NameSuffix will generate a random name suffix
func NameSuffix() string { return nameSuffix(GlobalFaker) }

// NameSuffix will generate a random name suffix
func (f *Faker) NameSuffix() string { return nameSuffix(f) }

func nameSuffix(f *Faker) string { return getRandValue(f, []string{"person", "suffix"}) }

// SSN will generate a random Social Security Number
func SSN() string { return ssn(GlobalFaker) }

// SSN will generate a random Social Security Number
func (f *Faker) SSN() string { return ssn(f) }

func ssn(f *Faker) string { return strconv.Itoa(randIntRange(f, 100000000, 999999999)) }

// Gender will generate a random gender string
func Gender() string { return gender(GlobalFaker) }

// Gender will generate a random gender string
func (f *Faker) Gender() string { return gender(f) }

func gender(f *Faker) string {
	if boolFunc(f) {
		return "male"
	}

	return "female"
}

// Hobby will generate a random hobby string
func Hobby() string { return hobby(GlobalFaker) }

// Hobby will generate a random hobby string
func (f *Faker) Hobby() string { return hobby(f) }

func hobby(f *Faker) string { return getRandValue(f, []string{"person", "hobby"}) }

// ContactInfo struct full of contact info
type ContactInfo struct {
	Phone string `json:"phone" xml:"phone"`
	Email string `json:"email" xml:"email"`
}

// Contact will generate a struct with information randomly populated contact information
func Contact() *ContactInfo { return contact(GlobalFaker) }

// Contact will generate a struct with information randomly populated contact information
func (f *Faker) Contact() *ContactInfo { return contact(f) }

func contact(f *Faker) *ContactInfo {
	return &ContactInfo{
		Phone: phone(f),
		Email: email(f),
	}
}

// Phone will generate a random phone number string
func Phone() string { return phone(GlobalFaker) }

// Phone will generate a random phone number string
func (f *Faker) Phone() string { return phone(f) }

func phone(f *Faker) string { return replaceWithNumbers(f, "##########") }

// PhoneFormatted will generate a random phone number string
func PhoneFormatted() string { return phoneFormatted(GlobalFaker) }

// PhoneFormatted will generate a random phone number string
func (f *Faker) PhoneFormatted() string { return phoneFormatted(f) }

func phoneFormatted(f *Faker) string {
	return replaceWithNumbers(f, getRandValue(f, []string{"person", "phone"}))
}

// Email will generate a random email string
func Email() string { return email(GlobalFaker) }

// Email will generate a random email string
func (f *Faker) Email() string { return email(f) }

func email(f *Faker) string {
	email := getRandValue(f, []string{"person", "first"}) + getRandValue(f, []string{"person", "last"})
	email += "@"
	email += getRandValue(f, []string{"person", "last"}) + "." + getRandValue(f, []string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}

// Teams takes in an array of people and team names and randomly places the people into teams as evenly as possible
func Teams(peopleArray []string, teamsArray []string) map[string][]string {
	return teams(GlobalFaker, peopleArray, teamsArray)
}

// Teams takes in an array of people and team names and randomly places the people into teams as evenly as possible
func (f *Faker) Teams(peopleArray []string, teamsArray []string) map[string][]string {
	return teams(f, peopleArray, teamsArray)
}

func teams(f *Faker, people []string, teams []string) map[string][]string {
	// Shuffle the people if more than 1
	if len(people) > 1 {
		shuffleStrings(f, people)
	}

	peopleIndex := 0
	teamsOutput := make(map[string][]string)
	numPer := math.Ceil(float64(len(people)) / float64(len(teams)))
	for _, team := range teams {
		teamsOutput[team] = []string{}
		for i := 0.00; i < numPer; i++ {
			if peopleIndex < len(people) {
				teamsOutput[team] = append(teamsOutput[team], people[peopleIndex])
				peopleIndex++
			}
		}
	}

	return teamsOutput
}

func addPersonLookup() {
	AddFuncLookup("person", Info{
		Display:     "Person",
		Category:    "person",
		Description: "Personal data, like name and contact details, used for identification and communication",
		Example: `{
	"first_name": "Markus",
	"last_name": "Moen",
	"gender": "male",
	"ssn": "275413589",
	"image": "https://picsum.photos/208/500",
	"hobby": "Lacrosse",
	"job": {
		"company": "Intermap Technologies",
		"title": "Developer",
		"descriptor": "Direct",
		"level": "Paradigm"
	},
	"address": {
		"address": "369 North Cornerbury, Miami, North Dakota 24259",
		"street": "369 North Cornerbury",
		"city": "Miami",
		"state": "North Dakota",
		"zip": "24259",
		"country": "Ghana",
		"latitude": -6.662595,
		"longitude": 23.921575
	},
	"contact": {
		"phone": "3023202027",
		"email": "lamarkoelpin@heaney.biz"
	},
	"credit_card": {
		"type": "Maestro",
		"number": "39800889982276",
		"exp": "01/29",
		"cvv": "932"
	}
}`,
		Output:      "map[string]any",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return person(f), nil
		},
	})

	AddFuncLookup("name", Info{
		Display:     "Name",
		Category:    "person",
		Description: "The given and family name of an individual",
		Example:     "Markus Moen",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return name(f), nil
		},
	})

	AddFuncLookup("nameprefix", Info{
		Display:     "Name Prefix",
		Category:    "person",
		Description: "A title or honorific added before a person's name",
		Example:     "Mr.",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return namePrefix(f), nil
		},
	})

	AddFuncLookup("namesuffix", Info{
		Display:     "Name Suffix",
		Category:    "person",
		Description: "A title or designation added after a person's name",
		Example:     "Jr.",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return nameSuffix(f), nil
		},
	})

	AddFuncLookup("firstname", Info{
		Display:     "First Name",
		Category:    "person",
		Description: "The name given to a person at birth",
		Example:     "Markus",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return firstName(f), nil
		},
	})

	AddFuncLookup("middlename", Info{
		Display:     "Middle Name",
		Category:    "person",
		Description: "Name between a person's first name and last name",
		Example:     "Belinda",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return middleName(f), nil
		},
	})

	AddFuncLookup("lastname", Info{
		Display:     "Last Name",
		Category:    "person",
		Description: "The family name or surname of an individual",
		Example:     "Daniel",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return lastName(f), nil
		},
	})

	AddFuncLookup("gender", Info{
		Display:     "Gender",
		Category:    "person",
		Description: "Classification based on social and cultural norms that identifies an individual",
		Example:     "male",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return gender(f), nil
		},
	})

	AddFuncLookup("ssn", Info{
		Display:     "SSN",
		Category:    "person",
		Description: "Unique nine-digit identifier used for government and financial purposes in the United States",
		Example:     "296446360",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return ssn(f), nil
		},
	})

	AddFuncLookup("hobby", Info{
		Display:     "Hobby",
		Category:    "person",
		Description: "An activity pursued for leisure and pleasure",
		Example:     "Swimming",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hobby(f), nil
		},
	})

	AddFuncLookup("email", Info{
		Display:     "Email",
		Category:    "person",
		Description: "Electronic mail used for sending digital messages and communication over the internet",
		Example:     "markusmoen@pagac.net",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return email(f), nil
		},
	})

	AddFuncLookup("phone", Info{
		Display:     "Phone",
		Category:    "person",
		Description: "Numerical sequence used to contact individuals via telephone or mobile devices",
		Example:     "6136459948",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phone(f), nil
		},
	})

	AddFuncLookup("phoneformatted", Info{
		Display:     "Phone Formatted",
		Category:    "person",
		Description: "Formatted phone number of a person",
		Example:     "136-459-9489",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return phoneFormatted(f), nil
		},
	})

	AddFuncLookup("teams", Info{
		Display:     "Teams",
		Category:    "person",
		Description: "Randomly split people into teams",
		Example: `{
	"Team 1": [
		"Justin",
		"Connor",
		"Jeff"
	],
	"Team 2": [
		"Sharon",
		"Fabian",
		"Billy"
	],
	"Team 3": [
		"Steve",
		"Robert"
	]
}`,
		Output:      "map[string][]string",
		ContentType: "application/json",
		Params: []Param{
			{Field: "people", Display: "Strings", Type: "[]string", Description: "Array of people"},
			{Field: "teams", Display: "Strings", Type: "[]string", Description: "Array of teams"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			people, err := info.GetStringArray(m, "people")
			if err != nil {
				return nil, err
			}

			teamsArray, err := info.GetStringArray(m, "teams")
			if err != nil {
				return nil, err
			}

			return teams(f, people, teamsArray), nil
		},
	})
}
