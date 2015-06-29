package gofakeit

// Name will generate a random First and Last Name
func Name() string {
	return getRandValue([]string{"name", "first"}) + " " + getRandValue([]string{"name", "last"})
}

// FirstName will generate a random first name
func FirstName() string {
	return getRandValue([]string{"name", "first"})
}

// LastName will generate a random last name
func LastName() string {
	return getRandValue([]string{"name", "last"})
}

// NamePrefix will generate a random name prefix
func NamePrefix() string {
	return getRandValue([]string{"name", "prefix"})
}

// NameSuffix will generate a random name suffix
func NameSuffix() string {
	return getRandValue([]string{"name", "suffix"})
}
