package gofakeit

// Generate Company Name
func Company() (company string) {
	switch randInt := randIntRange(1, 3); randInt {
	case 1:
		company = LastName() + ", " + LastName() + " and " + LastName()
	case 2:
		company = LastName() + "-" + LastName()
	case 3:
		company = LastName() + " " + CompanySuffix()
	}

	return
}

// Generate Company Suffix
func CompanySuffix() string {
	return getRandValue([]string{"company", "suffix"})
}

// Generate Company Buzzword
func BuzzWord() string {
	return getRandValue([]string{"company", "buzzwords"})
}

// Generate Company BS
func BS() string {
	return getRandValue([]string{"company", "bs"})
}
