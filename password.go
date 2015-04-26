package gofakeit

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

	passBytes := []byte(passString)
	finalBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		finalBytes[i] = passBytes[randIntRange(0, len(passBytes))]
	}
	return string(finalBytes)
}
