package data

import (
	"strings"
	"testing"
)

// Loop through the map keys and loop through the values of each key
// to check for space before or after each string
func TestCheckWordForSpaces(t *testing.T) {
	for key, values := range Word {
		// Loop through the values
		for _, value := range values {
			// Check if value starts with a space
			if strings.HasPrefix(value, " ") {
				t.Errorf("category %s value %s starts with a space", key, value)
			}

			// Check if value ends with a space
			if strings.HasSuffix(value, " ") {
				t.Errorf("category %s value %s starts with a space", key, value)
			}
		}
	}
}
