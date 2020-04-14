package gofakeit

// Fruit will return a random fruit name
func Fruit() string {
	return getRandValue([]string{"food", "fruit"})
}

// Vegetable will return a random vegetable name
func Vegetable() string {
	return getRandValue([]string{"food", "vegetable"})
}

func addFoodLookup() {
	AddLookupData("fruit", Info{
		Category:    "food",
		Description: "Random fruit",
		Example:     "Dates",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Fruit(), nil
		},
	})

	AddLookupData("vegetable", Info{
		Category:    "food",
		Description: "Random vegetable",
		Example:     "Amaranth Leaves",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Vegetable(), nil
		},
	})
}
