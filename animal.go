package gofakeit

// PetName will return a random fun pet name
func PetName() string {
	return getRandValue([]string{"animal", "petname"})
}

// Animal will return a random animal
func Animal() string {
	return getRandValue([]string{"animal", "animal"})
}

// AnimalType will return a random animal type
func AnimalType() string {
	return getRandValue([]string{"animal", "type"})
}

// FarmAnimal will return a random animal that usually lives on a farm
func FarmAnimal() string {
	return getRandValue([]string{"animal", "farm"})
}

// Cat will return a random cat breed
func Cat() string {
	return getRandValue([]string{"animal", "cat"})
}

// Dog will return a random dog breed
func Dog() string {
	return getRandValue([]string{"animal", "dog"})
}
