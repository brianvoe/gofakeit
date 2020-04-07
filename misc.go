package gofakeit

import (
	"github.com/brianvoe/gofakeit/v4/data"
)

// Bool will generate a random boolean value
func Bool() bool {
	return randIntRange(0, 1) == 1
}

// Categories will return a map string array of available data categories and sub categories
func Categories() map[string][]string {
	types := make(map[string][]string)
	for category, subCategoriesMap := range data.Data {
		subCategories := make([]string, 0)
		for subType := range subCategoriesMap {
			subCategories = append(subCategories, subType)
		}
		types[category] = subCategories
	}
	return types
}

func addMiscLookup() {
	AddLookupData("bool", Info{
		Description: "Random boolean",
		Example:     "true",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Bool(), nil
		},
	})
}
