package gofakeit

import (
	"strings"
)

// Fruit will return a random fruit name
func Fruit() string { return fruit(GlobalFaker) }

// Fruit will return a random fruit name
func (f *Faker) Fruit() string { return fruit(f) }

func fruit(f *Faker) string { return getRandValue(f, []string{"food", "fruit"}) }

// Vegetable will return a random vegetable name
func Vegetable() string { return vegetable(GlobalFaker) }

// Vegetable will return a random vegetable name
func (f *Faker) Vegetable() string { return vegetable(f) }

func vegetable(f *Faker) string { return getRandValue(f, []string{"food", "vegetable"}) }

// Breakfast will return a random breakfast name
func Breakfast() string { return breakfast(GlobalFaker) }

// Breakfast will return a random breakfast name
func (f *Faker) Breakfast() string { return breakfast(f) }

func breakfast(f *Faker) string {
	v := getRandValue(f, []string{"food", "breakfast"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Lunch will return a random lunch name
func Lunch() string { return lunch(GlobalFaker) }

// Lunch will return a random lunch name
func (f *Faker) Lunch() string { return lunch(f) }

func lunch(f *Faker) string {
	v := getRandValue(f, []string{"food", "lunch"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Dinner will return a random dinner name
func Dinner() string { return dinner(GlobalFaker) }

// Dinner will return a random dinner name
func (f *Faker) Dinner() string { return dinner(f) }

func dinner(f *Faker) string {
	v := getRandValue(f, []string{"food", "dinner"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Drink will return a random drink name
func Drink() string { return drink(GlobalFaker) }

// Drink will return a random drink name
func (f *Faker) Drink() string { return drink(f) }

func drink(f *Faker) string {
	v := getRandValue(f, []string{"food", "drink"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Snack will return a random snack name
func Snack() string { return snack(GlobalFaker) }

// Snack will return a random snack name
func (f *Faker) Snack() string { return snack(f) }

func snack(f *Faker) string {
	v := getRandValue(f, []string{"food", "snack"})
	return strings.ToUpper(v[:1]) + v[1:]
}

// Dessert will return a random dessert name
func Dessert() string { return dessert(GlobalFaker) }

// Dessert will return a random dessert name
func (f *Faker) Dessert() string { return dessert(f) }

func dessert(f *Faker) string {
	v := getRandValue(f, []string{"food", "dessert"})
	return strings.ToUpper(v[:1]) + v[1:]
}

func addFoodLookup() {
	AddFuncLookup("fruit", Info{
		Display:     "Fruit",
		Category:    "food",
		Description: "Edible plant part, typically sweet, enjoyed as a natural snack or dessert",
		Example:     "Peach",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return fruit(f), nil
		},
	})

	AddFuncLookup("vegetable", Info{
		Display:     "Vegetable",
		Category:    "food",
		Description: "Edible plant or part of a plant, often used in savory cooking or salads",
		Example:     "Amaranth Leaves",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return vegetable(f), nil
		},
	})

	AddFuncLookup("breakfast", Info{
		Display:     "Breakfast",
		Category:    "food",
		Description: "First meal of the day, typically eaten in the morning",
		Example:     "Blueberry banana happy face pancakes",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return breakfast(f), nil
		},
	})

	AddFuncLookup("lunch", Info{
		Display:     "Lunch",
		Category:    "food",
		Description: "Midday meal, often lighter than dinner, eaten around noon",
		Example:     "No bake hersheys bar pie",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return lunch(f), nil
		},
	})

	AddFuncLookup("dinner", Info{
		Display:     "Dinner",
		Category:    "food",
		Description: "Evening meal, typically the day's main and most substantial meal",
		Example:     "Wild addicting dip",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return dinner(f), nil
		},
	})

	AddFuncLookup("drink", Info{
		Display:     "Drink",
		Category:    "food",
		Description: "Liquid consumed for hydration, pleasure, or nutritional benefits",
		Example:     "Soda",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return drink(f), nil
		},
	})

	AddFuncLookup("snack", Info{
		Display:     "Snack",
		Category:    "food",
		Description: "Random snack",
		Example:     "Small, quick food item eaten between meals",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return snack(f), nil
		},
	})

	AddFuncLookup("dessert", Info{
		Display:     "Dessert",
		Category:    "food",
		Description: "Sweet treat often enjoyed after a meal",
		Example:     "French napoleons",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return dessert(f), nil
		},
	})
}
