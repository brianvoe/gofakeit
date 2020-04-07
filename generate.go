package gofakeit

import (
	"encoding/json"
	"errors"
	"strings"
)

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

func addGenerateLookup() {
	AddLookupData("generate", Info{
		Description: "Random string generated from string value based upon available data sets",
		Example:     "{person.first} {person.last} {person.email} - Markus Moen markusmoen@pagac.net",
		Params: []Param{
			{Field: "value", Required: true, Type: "string", Description: "String value to generate from"},
		},
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
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
		Description: "Random json data",
		Example:     `{"aut":{"iste":"qui"},"est":["velit","architecto","doloribus","fugiat"],"id":620384.06,"maiores":"Bertha Weber","quisquam":715240.3,"tempore":7415722}`,
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			d, _ := json.Marshal(Map())
			return string(d), nil
		},
	})
}
