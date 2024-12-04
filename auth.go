package gofakeit

// Username will generate a random username based upon picking a random lastname and random numbers at the end
func Username() string {
	return username(GlobalFaker)
}

// Username will generate a random username based upon picking a random lastname and random numbers at the end
func (f *Faker) Username() string {
	return username(f)
}

func username(f *Faker) string {
	return getRandValue(f, []string{"person", "last"}) + replaceWithNumbers(f, "####")
}

// Password will generate a random password.
// Minimum number length of 5 if less than.
func Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return password(GlobalFaker, lower, upper, numeric, special, space, num)
}

// StrongPassword will generate a random password as the original function Password,
// but with a strong inclusion of the required symbols depending on options.
// Minimum number length of 5 if less than.
func StrongPassword(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return strongPassword(GlobalFaker, lower, upper, numeric, special, space, num)
}

// Password will generate a random password.
// Minimum number length of 5 if less than.
func (f *Faker) Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return password(f, lower, upper, numeric, special, space, num)
}

// StrongPassword will generate a random password as the original function Password,
// but with a strong inclusion of the required symbols depending on options.
// Minimum number length of 5 if less than.
func (f *Faker) StrongPassword(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return strongPassword(f, lower, upper, numeric, special, space, num)
}

func password(f *Faker, lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	// Make sure the num minimum is at least 5
	if num < 5 {
		num = 5
	}

	items, weights := itemsAndWeights(lower, upper, numeric, special, space)

	// Create byte slice
	b := make([]byte, num)

	for i := 0; i <= num-1; i++ {
		// Run weighted
		weight, _ := weighted(f, items, weights)

		b[i] = selectCharacter(f, weight.(string))
	}

	// Shuffle bytes
	for i := range b {
		j := f.IntN(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	// Replace first or last character if it's a space, and other options are available
	if b[0] == ' ' {
		b[0] = password(f, lower, upper, numeric, special, false, 1)[0]
	}
	if b[len(b)-1] == ' ' {
		b[len(b)-1] = password(f, lower, upper, numeric, special, false, 1)[0]
	}

	return string(b)
}

func strongPassword(f *Faker, lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	items, weights := itemsAndWeights(lower, upper, numeric, special, space)

	// Make sure the num minimum is at least 5
	if num < 5 {
		num = 5
	}

	// Create byte slice
	b := make([]byte, num)

	// Depending on the provided options, need to track what rules to apply
	strongRules := generateRules(lower, upper, numeric, special, space)

	for i := 0; i < num; i++ {
		var currentWeight string

		// When len of rules equal to number of left characters,
		// need to apply option from the rule,
		// otherwise it is not possible to use all the required options
		if len(strongRules) == num-i {
			for w := range strongRules {
				currentWeight = w
				break
			}
		} else {
			// Run weighted
			weight, _ := weighted(f, items, weights)
			currentWeight = weight.(string)
		}

		b[i] = selectCharacter(f, currentWeight)

		// Remove the key from rules, because it has been applied in the current iteration
		delete(strongRules, currentWeight)
	}

	// Shuffle bytes
	for i := range b {
		j := f.IntN(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	// Replace first or last character if it's a space, and other options are available
	for b[0] == ' ' {
		i := 0
		j := f.IntN(num)
		// Find index j, so that b[j] is not a space and j is not the last character
		for b[j] == ' ' || j == num-1 {
			j = f.IntN(num)
		}
		b[i], b[j] = b[j], b[i]
	}
	if b[len(b)-1] == ' ' {
		i := num - 1
		j := f.IntN(num)
		// Find index j, so that b[j] is not a space and j is not the first character
		for b[j] == ' ' || j == 0 {
			j = f.IntN(num)
		}
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

func itemsAndWeights(lower bool, upper bool, numeric bool, special bool, space bool) ([]any, []float32) {
	// Setup weights
	items := make([]any, 0)
	weights := make([]float32, 0)
	if lower {
		items = append(items, "l")
		weights = append(weights, 4)
	}
	if upper {
		items = append(items, "u")
		weights = append(weights, 4)
	}
	if numeric {
		items = append(items, "n")
		weights = append(weights, 3)
	}
	if special {
		items = append(items, "e")
		weights = append(weights, 2)
	}
	if space {
		items = append(items, "a")
		weights = append(weights, 1)
	}

	// If no items are selected then default to lower, upper, numeric
	if len(items) == 0 {
		items = append(items, "l", "u", "n")
		weights = append(weights, 4, 4, 3)
	}

	return items, weights
}

func generateRules(lower bool, upper bool, numeric bool, special bool, space bool) map[string]bool {
	rules := make(map[string]bool)

	if lower {
		rules["l"] = true
	}
	if upper {
		rules["u"] = true
	}
	if numeric {
		rules["n"] = true
	}
	if special {
		rules["e"] = true
	}
	if space {
		rules["a"] = true
	}
	return rules
}

func selectCharacter(f *Faker, key string) byte {
	switch key {
	case "l":
		return lowerStr[f.Int64()%int64(len(lowerStr))]
	case "u":
		return upperStr[f.Int64()%int64(len(upperStr))]
	case "n":
		return numericStr[f.Int64()%int64(len(numericStr))]
	case "e":
		return specialSafeStr[f.Int64()%int64(len(specialSafeStr))]
	case "a":
		return spaceStr[f.Int64()%int64(len(spaceStr))]
	}
	return byte(0)
}

func addAuthLookup() {
	AddFuncLookup("username", Info{
		Display:     "Username",
		Category:    "auth",
		Description: "Unique identifier assigned to a user for accessing an account or system",
		Example:     "Daniel1364",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return username(f), nil
		},
	})

	AddFuncLookup("password", Info{
		Display:     "Password",
		Category:    "auth",
		Description: "Secret word or phrase used to authenticate access to a system or account",
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
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
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

			return password(f, lower, upper, numeric, special, space, length), nil
		},
	})
}
