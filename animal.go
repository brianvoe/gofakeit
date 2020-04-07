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
	AddLookupData("animal", Info{
		Description: "Random animal",
		Example:     "elk",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Animal(), nil
		},
	})

	AddLookupData("animal.type", Info{
		Description: "Random animal type",
		Example:     "amphibians",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return AnimalType(), nil
		},
	})

	AddLookupData("animal.farm", Info{
		Description: "Random farm animal",
		Example:     "Chicken",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return FarmAnimal(), nil
		},
	})

	AddLookupData("animal.petname", Info{
		Description: "Random pet name",
		Example:     "Ozzy Pawsborne",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return PetName(), nil
		},
	})

	AddLookupData("animal.cat", Info{
		Description: "Random cat type",
		Example:     "Chausie",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Cat(), nil
		},
	})

	AddLookupData("animal.dog", Info{
		Description: "Random dog type",
		Example:     "Norwich Terrier",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Dog(), nil
		},
	})
}
