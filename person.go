package gofakeit

import (
	"strconv"
	"strings"
)

// PersonInfo is a struct of person information
type PersonInfo struct {
	FirstName  string          `json:"first_name"`
	LastName   string          `json:"last_name"`
	Gender     string          `json:"gender"`
	SSN        string          `json:"ssn"`
	Image      string          `json:"image"`
	Job        *JobInfo        `json:"job"`
	Address    *AddressInfo    `json:"address"`
	Contact    *ContactInfo    `json:"contact"`
	CreditCard *CreditCardInfo `json:"credit_card"`
}

// Person will generate a struct with person information
func Person() *PersonInfo {
	return &PersonInfo{
		FirstName:  FirstName(),
		LastName:   LastName(),
		Gender:     Gender(),
		SSN:        SSN(),
		Image:      ImageURL(300, 300) + "/people",
		Job:        Job(),
		Address:    Address(),
		Contact:    Contact(),
		CreditCard: CreditCard(),
	}
}

// Name will generate a random First and Last Name
func Name() string {
	return getRandValue([]string{"person", "first"}) + " " + getRandValue([]string{"person", "last"})
}

// FirstName will generate a random first name
func FirstName() string {
	return getRandValue([]string{"person", "first"})
}

// LastName will generate a random last name
func LastName() string {
	return getRandValue([]string{"person", "last"})
}

// NamePrefix will generate a random name prefix
func NamePrefix() string {
	return getRandValue([]string{"person", "prefix"})
}

// NameSuffix will generate a random name suffix
func NameSuffix() string {
	return getRandValue([]string{"person", "suffix"})
}

// SSN will generate a random Social Security Number
func SSN() string {
	return strconv.Itoa(randIntRange(100000000, 999999999))
}

// Gender will generate a random gender string
func Gender() string {
	if Bool() {
		return "male"
	}

	return "female"
}

// ContactInfo struct full of contact info
type ContactInfo struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Contact will generate a struct with information randomly populated contact information
func Contact() *ContactInfo {
	return &ContactInfo{
		Phone: Phone(),
		Email: Email(),
	}
}

// Phone will generate a random phone number string
func Phone() string {
	return replaceWithNumbers("##########")
}

// PhoneFormatted will generate a random phone number string
func PhoneFormatted() string {
	return replaceWithNumbers(getRandValue([]string{"contact", "phone"}))
}

// Email will generate a random email string
func Email() string {
	email := getRandValue([]string{"person", "first"}) + getRandValue([]string{"person", "last"})
	email += "@"
	email += getRandValue([]string{"person", "last"}) + "." + getRandValue([]string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}

func addPersonLookup() {
	AddLookupData("person", Info{
		Description: "Random set of person info",
		Example: `{
			first_name: "Markus",
			last_name: "Moen",
			gender: "male",
			ssn: "420776036",
			image: "https://picsum.photos/300/300/people",
			job: {
				company: "Lockman and Sons",
				title: "Developer",
				descriptor: "Global",
				level: "Brand"
			}, 
			address: {
				address: "5369 Streamville, Rossieview, Hawaii 42591",
				street: "5369 Streamville",
				city: "Rossieview",
				state: "Hawaii",
				zip: "42591",
				country: "Burkina Faso",
				latitude: "-6.662594491850811",
				longitude: "23.921575244414612"
			},
			contact: {
				phone: "6136459948",
				email: "carolecarroll@bosco.com"
			},
			credit_card: {
				type: "Visa",
				number: "6536459948995369",
				exp: "03/27",
				cvv: "353"
			}
		}`,
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Person(), nil
		},
	})

	AddLookupData("person.name", Info{
		Description: "Random name",
		Example:     "Markus Moen",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Name(), nil
		},
	})

	AddLookupData("person.name.prefix", Info{
		Description: "Random name prefix",
		Example:     "Mr.",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return NamePrefix(), nil
		},
	})

	AddLookupData("person.name.suffix", Info{
		Description: "Random name suffix",
		Example:     "Jr.",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return NameSuffix(), nil
		},
	})

	AddLookupData("person.name.first", Info{
		Description: "Random first name",
		Example:     "Markus",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return FirstName(), nil
		},
	})

	AddLookupData("person.name.last", Info{
		Description: "Random last name",
		Example:     "Daniel",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return LastName(), nil
		},
	})

	AddLookupData("person.gender", Info{
		Description: "Random gender",
		Example:     "male",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Gender(), nil
		},
	})

	AddLookupData("person.ssn", Info{
		Description: "Random social security number",
		Example:     "296446360",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return SSN(), nil
		},
	})

	AddLookupData("person.email", Info{
		Description: "Random email",
		Example:     "markusmoen@pagac.net",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Email(), nil
		},
	})

	AddLookupData("person.phone", Info{
		Description: "Random phone number",
		Example:     "6136459948",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Phone(), nil
		},
	})

	AddLookupData("person.phone.formatted", Info{
		Description: "Random formatted phone number",
		Example:     "136-459-9489",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return PhoneFormatted(), nil
		},
	})
}
