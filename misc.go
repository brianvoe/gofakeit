package gofakeit

import (
	"encoding/hex"
	"math/rand"
	"reflect"

	"github.com/brianvoe/gofakeit/v5/data"
)

// Bool will generate a random boolean value
func Bool() bool {
	return randIntRange(0, 1) == 1
}

// UUID (version 4) will generate a random unique identifier based upon random numbers
// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func UUID() string {
	return uuid(nil)
}

// UUID (version 4) will generate a random unique identifier based upon random numbers
// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (f Faker) UUID() string {
	return uuid(f.Rand)
}

func uuid(r *rand.Rand) string {
	version := byte(4)
	uuid := make([]byte, 16)
	if r != nil {
		r.Read(uuid)
	} else {
		rand.Read(uuid)
	}

	// Set version
	uuid[6] = (uuid[6] & 0x0f) | (version << 4)

	// Set variant
	uuid[8] = (uuid[8] & 0xbf) | 0x80

	buf := make([]byte, 36)
	var dash byte = '-'
	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = dash
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = dash
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = dash
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = dash
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}

func ShuffleAnySlice(v interface{}) {
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
		j := int(rand.Int31n(int32(i + 1)))
		swap(i, j)
	}
}

// FlipACoin will return a random value of Heads or Tails
func FlipACoin() string {
	if Bool() {
		return "Heads"
	}

	return "Tails"
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
	AddFuncLookup("uuid", Info{
		Display:     "UUID",
		Category:    "misc",
		Description: "Random uuid",
		Example:     "590c1440-9888-45b0-bd51-a817ee07c3f2",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return UUID(), nil
		},
	})

	AddFuncLookup("bool", Info{
		Display:     "Boolean",
		Category:    "misc",
		Description: "Random boolean",
		Example:     "true",
		Output:      "bool",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Bool(), nil
		},
	})

	AddFuncLookup("flipacoin", Info{
		Display:     "Flip A Coin",
		Category:    "misc",
		Description: "Random Heads or Tails outcome",
		Example:     "Tails",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FlipACoin(), nil
		},
	})
}
