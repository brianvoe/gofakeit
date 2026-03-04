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
	username := getRandValue(f, []string{"auth", "username"})
	username, err := generate(f, username)
	if err != nil {
		return username // fallback to raw template if generation fails
	}

	return username
}

// Password will generate a random password.
// Minimum number length of 5 if less than.
func Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return password(GlobalFaker, lower, upper, numeric, special, space, num)
}

// Password will generate a random password.
// Minimum number length of 5 if less than.
func (f *Faker) Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	return password(f, lower, upper, numeric, special, space, num)
}

func password(f *Faker, lower bool, upper bool, numeric bool, special bool, space bool, num int) string {
	// Minimum length
	if num < 5 {
		num = 5
	}

	// Character group types and defaults
	// Each group has a weight for random selection; non-space groups are used for guaranteed slots.
	type charGroup struct {
		chars   string
		weight  int
		isSpace bool
	}

	defaultNonSpace := [...]charGroup{
		{chars: lowerStr, weight: 4},
		{chars: upperStr, weight: 4},
		{chars: numericStr, weight: 3},
	}
	const defaultNonSpaceWeight = 4 + 4 + 3

	var (
		activeBuf      [5]charGroup
		nonSpaceBuf    [5]charGroup
		active         []charGroup // all enabled groups (used for random fill)
		totalWeight    int
		nonSpace       []charGroup // enabled groups excluding space (for guaranteed slots and edge fixes)
		nonSpaceWeight int
	)

	// Build active character sets from enabled options
	appendGroup := func(enabled bool, chars string, weight int, isSpace bool) {
		if !enabled {
			return
		}
		active = append(active, charGroup{chars: chars, weight: weight, isSpace: isSpace})
		totalWeight += weight
		if !isSpace {
			nonSpace = append(nonSpace, charGroup{chars: chars, weight: weight})
			nonSpaceWeight += weight
		}
	}

	active = activeBuf[:0]
	nonSpace = nonSpaceBuf[:0]
	appendGroup(lower, lowerStr, 4, false)
	appendGroup(upper, upperStr, 4, false)
	appendGroup(numeric, numericStr, 3, false)
	appendGroup(special, specialSafeStr, 2, false)
	appendGroup(space, spaceStr, 1, true)

	// Fallbacks when nothing or only space is selected
	if len(active) == 0 {
		active = defaultNonSpace[:]
		totalWeight = defaultNonSpaceWeight
		nonSpace = active
		nonSpaceWeight = totalWeight
	} else if nonSpaceWeight == 0 {
		nonSpace = defaultNonSpace[:]
		nonSpaceWeight = defaultNonSpaceWeight
	}

	// Helper: draw one random character from a weighted set of groups
	draw := func(groups []charGroup, total int) byte {
		if total <= 0 {
			groups = defaultNonSpace[:]
			total = defaultNonSpaceWeight
		}

		r := f.IntN(total)
		for _, g := range groups {
			if r < g.weight {
				return g.chars[f.IntN(len(g.chars))]
			}
			r -= g.weight
		}

		// Should never be reached, but fall back to the last group just in case.
		g := groups[len(groups)-1]
		return g.chars[f.IntN(len(g.chars))]
	}

	// Guaranteed slots: one character from each enabled set
	// First slots: one from each enabled non-space group. Then one space if space is enabled.
	b := make([]byte, num)
	end := len(nonSpace)
	for i, g := range nonSpace {
		b[i] = g.chars[f.IntN(len(g.chars))]
	}
	if space {
		b[end] = ' '
		end++
	}

	// Fill remaining positions randomly from all active groups
	for i := end; i < num; i++ {
		b[i] = draw(active, totalWeight)
	}

	// Shuffle so guaranteed characters are not in fixed positions
	for i := range b {
		j := f.IntN(i + 1)
		b[i], b[j] = b[j], b[i]
	}

	// Ensure password does not start or end with a space
	// Swap with an interior non-space so we keep one space in the password; fallback to overwrite if no interior non-space.
	if b[0] == ' ' {
		for i := 1; i < len(b)-1; i++ {
			if b[i] != ' ' {
				b[0], b[i] = b[i], b[0]
				break
			}
		}
		if b[0] == ' ' {
			b[0] = draw(nonSpace, nonSpaceWeight)
		}
	}
	if b[len(b)-1] == ' ' {
		for i := len(b) - 2; i >= 1; i-- {
			if b[i] != ' ' {
				b[len(b)-1], b[i] = b[i], b[len(b)-1]
				break
			}
		}
		if b[len(b)-1] == ' ' {
			b[len(b)-1] = draw(nonSpace, nonSpaceWeight)
		}
	}

	return string(b)
}

func addAuthLookup() {
	AddFuncLookup("username", Info{
		Display:     "Username",
		Category:    "auth",
		Description: "Unique identifier assigned to a user for accessing an account or system",
		Example:     "Daniel1364",
		Output:      "string",
		Aliases: []string{
			"user name",
			"login name",
			"account username",
			"account login",
			"screen name",
			"user handle",
		},
		Keywords: []string{
			"login", "handle", "userid", "screenname",
			"user", "account", "credential", "signin", "alias", "profile", "uid",
		},
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
		Aliases: []string{
			"user password",
			"account password",
			"login password",
			"secret phrase",
			"auth secret",
		},
		Keywords: []string{
			"passphrase", "pwd", "secret",
			"credential", "authentication", "auth",
			"security", "signin", "login",
			"access", "key", "token", "hash", "encryption",
		},
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
