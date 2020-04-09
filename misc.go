package gofakeit

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
	"strings"

	"github.com/brianvoe/gofakeit/v4/data"
)

// Bool will generate a random boolean value
func Bool() bool {
	return randIntRange(0, 1) == 1
}

// UUID (version 4) will generate a random unique identifier based upon random nunbers
// Format: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func UUID() string {
	version := byte(4)
	uuid := make([]byte, 16)
	rand.Read(uuid)

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

// Generate fake information from given string. String should contain {category.subcategory}
//
// Ex: {person.first} - random firstname
//
// Ex: {person.first}###{person.last}@{person.last}.{internet.domain_suffix} - billy834smith@smith.com
//
// Ex: ### - 481 - random numbers
//
// Ex: ??? - fda - random letters
//
// For a complete list possible categories use the Categories() function.
func Generate(dataVal string) string {
	// Identify items between brackets: {person.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		catValue := ""
		startIndex := strings.Index(dataVal, "{")
		endIndex := strings.Index(dataVal, "}")
		replace := dataVal[(startIndex + 1):endIndex]
		categories := strings.Split(replace, ".")

		if len(categories) >= 2 && dataCheck([]string{categories[0], categories[1]}) {
			catValue = getRandValue([]string{categories[0], categories[1]})
		}

		dataVal = strings.Replace(dataVal, "{"+replace+"}", catValue, 1)
	}

	// Replace # with numbers
	dataVal = replaceWithNumbers(dataVal)

	// Replace ? with letters
	dataVal = replaceWithLetters(dataVal)

	return dataVal
}

// Map will generate a random set of map data
func Map() map[string]interface{} {
	m := map[string]interface{}{}

	randString := func() string {
		s := RandString([]string{"lorem", "bs", "job", "name", "address"})
		switch s {
		case "bs":
			return BS()
		case "job":
			return JobTitle()
		case "name":
			return Name()
		case "address":
			return Street() + ", " + City() + ", " + State() + " " + Zip()
		}
		return Word()
	}

	randSlice := func() []string {
		var sl []string
		for ii := 0; ii < Number(3, 10); ii++ {
			sl = append(sl, Word())
		}
		return sl
	}

	for i := 0; i < Number(3, 10); i++ {
		t := RandString([]string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[Word()] = randString()
		case "int":
			m[Word()] = Number(1, 10000000)
		case "float":
			m[Word()] = Float32Range(1, 1000000)
		case "slice":
			m[Word()] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := RandString([]string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[Word()] = randString()
			case "int":
				mm[Word()] = Number(1, 10000000)
			case "float":
				mm[Word()] = Float32Range(1, 1000000)
			case "slice":
				mm[Word()] = randSlice()
			}
			m[Word()] = mm
		}
	}

	return m
}

func addMiscLookup() {
	AddLookupData("uuid", Info{
		Category:    "misc",
		Description: "Random uuid",
		Example:     "590c1440-9888-45b0-bd51-a817ee07c3f2",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return UUID(), nil
		},
	})

	AddLookupData("bool", Info{
		Category:    "misc",
		Description: "Random boolean",
		Example:     "true",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Bool(), nil
		},
	})

	AddLookupData("generate", Info{
		Category:    "misc",
		Description: "Random string generated from string value based upon available data sets",
		Example:     "{person.first} {person.last} {person.email} - Markus Moen markusmoen@pagac.net",
		Params: []Param{
			{Field: "value", Required: true, Type: "string", Description: "String value to generate from"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			value, err := info.GetString(m, "value")
			if err != nil {
				return nil, err
			}

			// Limit the length of the string passed
			if len(value) >= 1000 {
				return nil, errors.New("Value length is too large. Limit to 1000 characters")
			}

			return Generate(value), nil
		},
	})

	AddLookupData("json", Info{
		Category:    "misc",
		Description: "Random json data",
		Example:     `{"aut":{"iste":"qui"},"est":["velit","architecto","doloribus","fugiat"],"id":620384.06,"maiores":"Bertha Weber","quisquam":715240.3,"tempore":7415722}`,
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			d, _ := json.Marshal(Map())
			return string(d), nil
		},
	})
}
