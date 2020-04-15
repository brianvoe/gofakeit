package gofakeit

import (
	"errors"
	"fmt"
	"strings"
)

// Generate fake information from given string.
// Replaceable values should be within {}
//
// Functions
// Ex: {firstname} - billy
// Ex: {sentence:3} - Record river mind.
// Ex: {uuid} - 590c1440-9888-45b0-bd51-a817ee07c3f2
//
// Letters/Numbers
// Ex: ### - 481 - random numbers
// Ex: ??? - fda - random letters
//
// For a complete list of runnable functions use FuncsLookup
func Generate(dataVal string) string {
	// Replace # with numbers and ? with letters
	dataVal = replaceWithNumbers(dataVal)
	dataVal = replaceWithLetters(dataVal)

	// Identify items between brackets: {person.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		fParts := dataVal[(strings.Index(dataVal, "{") + 1):strings.Index(dataVal, "}")]

		// Check if has params seperated by :
		fNameSplit := strings.SplitN(fParts, ":", 2)
		fName := ""
		fParams := ""
		if len(fNameSplit) >= 1 {
			fName = fNameSplit[0]
		}
		if len(fNameSplit) >= 2 {
			fParams = fNameSplit[1]
		}

		// Check to see if its a replacable lookup function
		if info := GetFuncLookup(fName); info != nil {
			// Get parameters, make sure params and the split both have values
			var mapParams map[string][]string
			paramsLen := len(info.Params)
			if paramsLen > 0 && fParams != "" {
				splitVals := funcLookupSplit(fParams)
				for i := 0; i < len(splitVals); i++ {
					if paramsLen-1 >= i {
						if mapParams == nil {
							mapParams = make(map[string][]string)
						}
						if strings.HasPrefix(splitVals[i], "[") {
							mapParams[info.Params[i].Field] = funcLookupSplit(strings.TrimRight(strings.TrimLeft(splitVals[i], "["), "]"))
						} else {
							mapParams[info.Params[i].Field] = []string{splitVals[i]}
						}
					}
				}
			}

			// Call function
			fValue, err := info.Call(&mapParams, info)
			if err != nil {
				// If we came across an error just dont replace value
				dataVal = strings.Replace(dataVal, "{"+fParts+"}", err.Error(), 1)
				continue
			}
			dataVal = strings.Replace(dataVal, "{"+fParts+"}", fmt.Sprintf("%v", fValue), 1)
			continue
		}

		// Couldnt find anything - set to n/a
		dataVal = strings.Replace(dataVal, "{"+fParts+"}", "n/a", 1)
		continue
	}

	return dataVal
}

// Map will generate a random set of map data
func Map() map[string]interface{} {
	m := map[string]interface{}{}

	randWordType := func() string {
		s := RandomString([]string{"lorem", "bs", "job", "name", "address"})
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
		t := RandomString([]string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[Word()] = randWordType()
		case "int":
			m[Word()] = Number(1, 10000000)
		case "float":
			m[Word()] = Float32Range(1, 1000000)
		case "slice":
			m[Word()] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := RandomString([]string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[Word()] = randWordType()
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
	AddFuncLookup("generate", Info{
		Category:    "misc",
		Description: "Random string generated from string value based upon available data sets",
		Example:     "{firstname} {lastname} {email} - Markus Moen markusmoen@pagac.net",
		Output:      "string",
		Params: []Param{
			{Field: "value", Type: "string", Description: "String value to generate from"},
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

	AddFuncLookup("map", Info{
		Category:    "misc",
		Description: "Random map data",
		Example:     `{"aut":{"iste":"qui"},"est":["velit","architecto","doloribus","fugiat"],"id":620384.06,"maiores":"Bertha Weber","quisquam":715240.3,"tempore":7415722}`,
		Output:      "map[string]interface",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Map(), nil
		},
	})
}
