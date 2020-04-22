package gofakeit

import "strings"

// Fruit will return a random fruit name
func Fruit() string {
	return getRandValue([]string{"food", "fruit"})
}

// Vegetable will return a random vegetable name
func Vegetable() string {
	return getRandValue([]string{"food", "vegetable"})
}

// Breakfast will return a random breakfast name
func Breakfast() string {
	v := getRandValue([]string{"food", "breakfast"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Lunch will return a random lunch name
func Lunch() string {
	v := getRandValue([]string{"food", "lunch"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Dinner will return a random dinner name
func Dinner() string {
	v := getRandValue([]string{"food", "dinner"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Snack will return a random snack name
func Snack() string {
	v := getRandValue([]string{"food", "snack"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Dessert will return a random dessert name
func Dessert() string {
	v := getRandValue([]string{"food", "dessert"})
	return strings.ToUpper(v[:1]) + v[1:]
}

func addFoodLookup() {
	AddFuncLookup("fruit", Info{
		Display:     "Fruit",
		Category:    "food",
		Description: "Random fruit",
		Example:     "Dates",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Fruit(), nil
		},
	})

	AddFuncLookup("vegetable", Info{
		Display:     "Vegetable",
		Category:    "food",
		Description: "Random vegetable",
		Example:     "Amaranth Leaves",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Vegetable(), nil
		},
	})

	AddFuncLookup("breakfast", Info{
		Display:     "Breakfast",
		Category:    "food",
		Description: "Random breakfast",
		Example:     "Blueberry banana happy face pancakes",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Breakfast(), nil
		},
	})

	AddFuncLookup("lunch", Info{
		Display:     "Lunch",
		Category:    "food",
		Description: "Random lunch",
		Example:     "No bake hersheys bar pie",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Lunch(), nil
		},
	})

	AddFuncLookup("dinner", Info{
		Display:     "Dinner",
		Category:    "food",
		Description: "Random dinner",
		Example:     "Wild addicting dip",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Dinner(), nil
		},
	})

	AddFuncLookup("snack", Info{
		Display:     "Snack",
		Category:    "food",
		Description: "Random snack",
		Example:     "Hoisin marinated wing pieces",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Snack(), nil
		},
	})

	AddFuncLookup("dessert", Info{
		Display:     "Dessert",
		Category:    "food",
		Description: "Random dessert",
		Example:     "French napoleons",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Dessert(), nil
		},
	})
}
