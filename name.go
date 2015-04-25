package gofakeit

// Ex: Billy Smith, Kim Barker, James Volk
func Name() string {
	return getRandValue([]string{"name", "first"}) + " " + getRandValue([]string{"name", "last"})
}

// Ex: Billy, Kim, James
func FirstName() string {
	return getRandValue([]string{"name", "first"})
}

// Ex: Smith, Barker, Volk
func LastName() string {
	return getRandValue([]string{"name", "last"})
}

// Ex: Mr., Mrs., Ms., Miss, Dr.
func PrefixName() string {
	return getRandValue([]string{"name", "prefix"})
}

// Ex: Jr., Sr., MD, PhD
func SuffixName() string {
	return getRandValue([]string{"name", "suffix"})
}
