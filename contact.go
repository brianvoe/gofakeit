package gofakeit

import (
	"strings"
)

type ContactInfo struct {
	Phone string
	Email string
}

func Contact() *ContactInfo {
	return &ContactInfo{
		Phone: Phone(),
		Email: Email(),
	}
}

// Generate Phone Number
func Phone() string {
	return replaceWithNumbers(getRandValue([]string{"contact", "phone"}))
}

// Generate Email
func Email() string {
	var email string

	email = getRandValue([]string{"name", "first"}) + getRandValue([]string{"name", "last"})
	email += "@"
	email += getRandValue([]string{"name", "last"}) + "." + getRandValue([]string{"internet", "domain_suffix"})

	return strings.ToLower(email)
}
