package gofakeit

// Generate First and Last Name
func Name() string {
	return getRandValue([]string{"name", "first"}) + " " + getRandValue([]string{"name", "last"})
}

// Generate First Name
func FirstName() string {
	return getRandValue([]string{"name", "first"})
}

// Generate Last Name
func LastName() string {
	return getRandValue([]string{"name", "last"})
}

// Ganerate Name Prefix
func PrefixName() string {
	return getRandValue([]string{"name", "prefix"})
}

// Generate Suffix Name
func SuffixName() string {
	return getRandValue([]string{"name", "suffix"})
}
