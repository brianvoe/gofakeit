package gofakeit

// Generate Domain Suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}
