package gofakeit

// Generate random Social Security Number
func SSN() string {
	return Numerify("###-###-####")
}

func Gender() string {
	if Bool() == true {
		return "male"
	}

	return "female"
}

type PersonInfo struct {
	FirstName  string
	LastName   string
	Gender     string
	SSN        string
	Image      string
	Job        *JobInfo
	Address    *AddressInfo
	Contact    *ContactInfo
	CreditCard *CreditCardInfo
}

// Generate struct full of person information
func Person() *PersonInfo {
	return &PersonInfo{
		FirstName:  FirstName(),
		LastName:   LastName(),
		Gender:     Gender(),
		SSN:        SSN(),
		Image:      ImageUrl(300, 300) + "/people",
		Job:        Job(),
		Address:    Address(),
		Contact:    Contact(),
		CreditCard: CreditCard(),
	}
}
