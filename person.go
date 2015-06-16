package gofakeit

// Generate random Social Security Number
func SSN() string {
	return Numerify("###-###-####")
}
