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
	return replaceWithNumbers(getRandValue([]string{"person", "phone"}))
}

// Email will generate a random email string
func Email() string {
	email := getRandValue([]string{"person", "first"}) + getRandValue([]string{"person", "last"})
	email += "@"
	email += getRandValue([]string{"person", "last"}) + "." + getRandValue([]string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}

func addPersonLookup() {
	AddFuncLookup("person", Info{
		Display:     "Person",
		Category:    "person",
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
		Output: "map[string]interface",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Person(), nil
		},
	})

	AddFuncLookup("name", Info{
		Display:     "Name",
		Category:    "person",
		Description: "Random name",
		Example:     "Markus Moen",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Name(), nil
		},
	})

	AddFuncLookup("nameprefix", Info{
		Display:     "Name Prefix",
		Category:    "person",
		Description: "Random name prefix",
		Example:     "Mr.",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return NamePrefix(), nil
		},
	})

	AddFuncLookup("namesuffix", Info{
		Display:     "Name Suffix",
		Category:    "person",
		Description: "Random name suffix",
		Example:     "Jr.",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return NameSuffix(), nil
		},
	})

	AddFuncLookup("firstname", Info{
		Display:     "First Name",
		Category:    "person",
		Description: "Random first name",
		Example:     "Markus",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FirstName(), nil
		},
	})

	AddFuncLookup("lastname", Info{
		Display:     "Last Name",
		Category:    "person",
		Description: "Random last name",
		Example:     "Daniel",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return LastName(), nil
		},
	})

	AddFuncLookup("gender", Info{
		Display:     "Gender",
		Category:    "person",
		Description: "Random gender",
		Example:     "male",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Gender(), nil
		},
	})

	AddFuncLookup("ssn", Info{
		Display:     "SSN",
		Category:    "person",
		Description: "Random social security number",
		Example:     "296446360",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return SSN(), nil
		},
	})

	AddFuncLookup("email", Info{
		Display:     "Email",
		Category:    "person",
		Description: "Random email",
		Example:     "markusmoen@pagac.net",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Email(), nil
		},
	})

	AddFuncLookup("phone", Info{
		Display:     "Phone",
		Category:    "person",
		Description: "Random phone number",
		Example:     "6136459948",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Phone(), nil
		},
	})

	AddFuncLookup("phoneformatted", Info{
		Display:     "Phone Formatted",
		Category:    "person",
		Description: "Random formatted phone number",
		Example:     "136-459-9489",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return PhoneFormatted(), nil
		},
	})
}
