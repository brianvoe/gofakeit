package gofakeit

import (
	"errors"
	"strings"
)

// Generate fake information from given string. String should contain {category.subcategory}
//
// Ex: {name.first} - random firstname
//
// Ex: {name.first}###{name.last}@{name.last}.{internet.domain_suffix} - billy834smith@smith.com
//
// Ex: ### - 481 - random numbers
//
// Ex: ??? - fda - random letters
//
// For a complete list possible categories use the Catagories() function.
func Generate(dataVal string) (string, error) {
	// Identify items between brackets: {name.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		catValue := ""
		startIndex := strings.Index(dataVal, "{")
		endIndex := strings.Index(dataVal, "}")
		replace := dataVal[(startIndex + 1):endIndex]
		categories := strings.Split(replace, ".")

		if len(categories) >= 2 && dataCheck([]string{categories[0], categories[1]}) {
			catValue = getRandValue([]string{categories[0], categories[1]})
		} else {
			return "", errors.New("not a valid {catagory.subcatagory} :" + replace)
		}

		dataVal = strings.Replace(dataVal, "{"+replace+"}", catValue, 1)
	}

	// Replace # with numbers
	dataVal = replaceWithNumbers(dataVal)

	// Replace ? with letters
	dataVal = replaceWithLetters(dataVal)

	return dataVal, nil
}
