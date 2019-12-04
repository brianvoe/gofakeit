package gofakeit

// PetName will return a random fun pet name
func PetName() string {
	return getRandValue([]string{"animal", "petname"})
}
