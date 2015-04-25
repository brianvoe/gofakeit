package gofakeit

import (
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
func New(dataVal string) string {
	// Identify items between brackets: {name.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		cat_value := ""
		start_index := strings.Index(dataVal, "{")
		end_index := strings.Index(dataVal, "}")
		replace := dataVal[(start_index + 1):end_index]
		categories := strings.Split(replace, ".")

		if len(categories) >= 2 && dataCheck([]string{categories[0], categories[1]}) {
			cat_value = getRandValue([]string{categories[0], categories[1]})
		}

		dataVal = strings.Replace(dataVal, "{"+replace+"}", cat_value, 1)
	}

	// Replace # with numbers
	dataVal = replaceWithNumbers(dataVal)

	// Replace ? with letters
	dataVal = replaceWithLetters(dataVal)

	return dataVal
}
