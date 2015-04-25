package gofakeit

import "crypto/rand"

// Password generator
func Password(lower bool, upper bool, numeric bool, special bool, space bool, length int) string {
	var passString string = ""
	var lowerStr string = "abcdefghijklmnopqrstuvwxyz"
	var upperStr string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numericStr string = "0123456789"
	var specialStr string = "!@#$%&?-_"
	var spaceStr string = " "

	if lower {
		passString += lowerStr
	}
	if upper {
		passString += upperStr
	}
	if numeric {
		passString += numericStr
	}
	if special {
		passString += specialStr
	}
	if space {
		passString += spaceStr
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	var passBytes = make([]byte, length)
	rand.Read(passBytes)
	for i, b := range passBytes {
		passBytes[i] = passString[b%byte(len(passString))]
	}
	return string(passBytes)
}
