package gofakeit

import (
	"reflect"

	"github.com/brianvoe/gofakeit/v7/data"
)

// Bool will generate a random boolean value
func Bool() bool { return boolFunc(GlobalFaker) }

// Bool will generate a random boolean value
func (f *Faker) Bool() bool { return boolFunc(f) }

func boolFunc(f *Faker) bool { return randIntRange(f, 0, 1) == 1 }

// ShuffleAnySlice takes in a slice and outputs it in a random order
func ShuffleAnySlice(v any) { shuffleAnySlice(GlobalFaker, v) }

// ShuffleAnySlice takes in a slice and outputs it in a random order
func (f *Faker) ShuffleAnySlice(v any) { shuffleAnySlice(f, v) }

func shuffleAnySlice(f *Faker, v any) {
	if v == nil {
		return
	}

	// Check type of passed in value, if not a slice return with no action taken
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Slice {
		return
	}

	s := reflect.ValueOf(v)
	n := s.Len()

	if n <= 1 {
		return
	}

	swap := func(i, j int) {
		tmp := reflect.ValueOf(s.Index(i).Interface())
		s.Index(i).Set(s.Index(j))
		s.Index(j).Set(tmp)
	}

	//if size is > int32 probably it will never finish, or ran out of entropy
	i := n - 1
	for ; i > 0; i-- {
		j := int(int32NFunc(f, int32(i+1)))
		swap(i, j)
	}
}

// FlipACoin will return a random value of Heads or Tails
func FlipACoin() string { return flipACoin(GlobalFaker) }

// FlipACoin will return a random value of Heads or Tails
func (f *Faker) FlipACoin() string { return flipACoin(f) }

func flipACoin(f *Faker) string {
	if boolFunc(f) {
		return "Heads"
	}

	return "Tails"
}

// RandomMapKey will return a random key from a map
func RandomMapKey(mapI any) any { return randomMapKey(GlobalFaker, mapI) }

// RandomMapKey will return a random key from a map
func (f *Faker) RandomMapKey(mapI any) any { return randomMapKey(f, mapI) }

func randomMapKey(f *Faker, mapI any) any {
	keys := reflect.ValueOf(mapI).MapKeys()
	return keys[f.IntN(len(keys))].Interface()
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
	AddFuncLookup("bool", Info{
		Display:     "Boolean",
		Category:    "misc",
		Description: "Data type that represents one of two possible values, typically true or false",
		Example:     "true",
		Output:      "bool",
		Aliases:     []string{"boolean", "true", "false", "logic", "binary"},
		Keywords:    []string{"bool", "data", "type", "represents", "values", "typically", "two", "possible"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return boolFunc(f), nil
		},
	})

	AddFuncLookup("flipacoin", Info{
		Display:     "Flip A Coin",
		Category:    "misc",
		Description: "Decision-making method involving the tossing of a coin to determine outcomes",
		Example:     "Tails",
		Output:      "string",
		Aliases:     []string{"coin", "flip", "heads", "tails", "decision", "random"},
		Keywords:    []string{"decision-making", "method", "tossing", "determine", "outcomes", "chance", "probability"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return flipACoin(f), nil
		},
	})
}
