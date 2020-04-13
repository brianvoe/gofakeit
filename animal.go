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

func addAnimalLookup() {
	AddLookupData("petname", Info{
		Category:    "animal",
		Description: "Random pet name",
		Example:     "Ozzy Pawsborne",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return PetName(), nil
		},
	})

	AddLookupData("animal", Info{
		Category:    "animal",
		Description: "Random animal",
		Example:     "elk",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Animal(), nil
		},
	})

	AddLookupData("animaltype", Info{
		Category:    "animal",
		Description: "Random animal type",
		Example:     "amphibians",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return AnimalType(), nil
		},
	})

	AddLookupData("farmanimal", Info{
		Category:    "animal",
		Description: "Random farm animal",
		Example:     "Chicken",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FarmAnimal(), nil
		},
	})

	AddLookupData("cat", Info{
		Category:    "animal",
		Description: "Random cat type",
		Example:     "Chausie",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Cat(), nil
		},
	})

	AddLookupData("dog", Info{
		Category:    "animal",
		Description: "Random dog type",
		Example:     "Norwich Terrier",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Dog(), nil
		},
	})
}
