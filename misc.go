package gofakeit

import (
	"fmt"
	"reflect"
	"sort"

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

	// A map is read in a random order, sort the keys so the same seed picks the same key
	sort.Slice(keys, func(i, j int) bool { return lessMapKey(keys[i], keys[j]) })

	return keys[f.IntN(len(keys))].Interface()
}

// lessMapKey orders strings, integers and floats by value, and any other key by its printed form
func lessMapKey(a, b reflect.Value) bool {
	// The keys of a map[any]T hold their value in an interface
	if a.Kind() == reflect.Interface {
		a = a.Elem()
	}
	if b.Kind() == reflect.Interface {
		b = b.Elem()
	}

	// Keys of different types, like 1 and "1" in a map[any]T, are ordered by type first
	if a.Type() != b.Type() {
		return a.Type().String() < b.Type().String()
	}

	switch a.Kind() {
	case reflect.String:
		return a.String() < b.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() < b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return a.Uint() < b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() < b.Float()
	}
	return fmt.Sprint(a.Interface()) < fmt.Sprint(b.Interface())
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
