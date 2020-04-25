package gofakeit

import (
	"math/rand"
)

// Username will genrate a random username based upon picking a random lastname and random numbers at the end
func Username() string {
	return getRandValue([]string{"person", "last"}) + replaceWithNumbers("####")
}

// Password will generate a random password
// Minimum number length of 5 if less than
func Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	// Make sure the num minimun is at least 5
	if num < 5 {
		num = 5
	}
	i := 0
	b := make([]byte, num)
	var passString string

	if lower {
		passString += lowerStr
		b[i] = lowerStr[rand.Int63()%int64(len(lowerStr))]
		i++
	}
	if upper {
		passString += upperStr
		b[i] = upperStr[rand.Int63()%int64(len(upperStr))]
		i++
	}
	if numeric {
		passString += numericStr
		b[i] = numericStr[rand.Int63()%int64(len(numericStr))]
		i++
	}
	if special {
		passString += specialStr
		b[i] = specialStr[rand.Int63()%int64(len(specialStr))]
		i++
	}
	if space {
		passString += spaceStr
		b[i] = spaceStr[rand.Int63()%int64(len(spaceStr))]
		i++
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	// Loop through and add it up
	for i <= num-1 {
		b[i] = passString[rand.Int63()%int64(len(passString))]
		i++
	}

	// Shuffle bytes
	for i := range b {
		j := rand.Intn(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func addAuthLookup() {
	AddFuncLookup("username", Info{
		Display:     "Username",
		Category:    "auth",
		Description: "Generates a random username",
		Example:     "Daniel1364",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Username(), nil
		},
	})

	AddFuncLookup("password", Info{
		Display:     "Password",
		Category:    "auth",
		Description: "Generates a random password",
		Example:     "EEP+wwpk 4lU-eHNXlJZ4n K9%v&TZ9e",
		Output:      "string",
		Params: []Param{
			{Field: "lower", Display: "Lower", Type: "bool", Default: "true", Description: "Whether or not to add lower case characters"},
			{Field: "upper", Display: "Upper", Type: "bool", Default: "true", Description: "Whether or not to add upper case characters"},
			{Field: "numeric", Display: "Numeric", Type: "bool", Default: "true", Description: "Whether or not to add numeric characters"},
			{Field: "special", Display: "Special", Type: "bool", Default: "true", Description: "Whether or not to add special characters"},
			{Field: "space", Display: "Space", Type: "bool", Default: "false", Description: "Whether or not to add spaces"},
			{Field: "length", Display: "Length", Type: "int", Default: "12", Description: "Number of characters in password"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			lower, err := info.GetBool(m, "lower")
			if err != nil {
				return nil, err
			}

			upper, err := info.GetBool(m, "upper")
			if err != nil {
				return nil, err
			}

			numeric, err := info.GetBool(m, "numeric")
			if err != nil {
				return nil, err
			}

			special, err := info.GetBool(m, "special")
			if err != nil {
				return nil, err
			}

			space, err := info.GetBool(m, "space")
			if err != nil {
				return nil, err
			}

			length, err := info.GetInt(m, "length")
			if err != nil {
				return nil, err
			}

			return Password(lower, upper, numeric, special, space, length), nil
		},
	})
}
