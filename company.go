package gofakeit

// Company will generate a random company name string
func Company() (company string) {
	switch randInt := randIntRange(1, 3); randInt {
	case 1:
		company = LastName() + ", " + LastName() + " and " + LastName()
	case 2:
		company = LastName() + "-" + LastName()
	case 3:
		company = LastName() + " " + CompanySuffix()
	}

	return
}

// CompanySuffix will generate a random company suffix string
func CompanySuffix() string {
	return getRandValue([]string{"company", "suffix"})
}

// BuzzWord will generate a random company buzz word string
func BuzzWord() string {
	return getRandValue([]string{"company", "buzzwords"})
}

// BS will generate a random company bs string
func BS() string {
	return getRandValue([]string{"company", "bs"})
}

func addCompanyLookup() {
	AddLookupData("company", Info{
		Description: "Random company name",
		Example:     "Moen, Pagac and Wuckert",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Company(), nil
		},
	})

	AddLookupData("company.suffix", Info{
		Description: "Random company name suffix",
		Example:     "Inc",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return CompanySuffix(), nil
		},
	})

	AddLookupData("company.bs", Info{
		Description: "Random bs company word",
		Example:     "front-end",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BS(), nil
		},
	})

	AddLookupData("company.buzzword", Info{
		Description: "Random company buzzwords",
		Example:     "disintermediate",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BuzzWord(), nil
		},
	})

	AddLookupData("company.job.title", Info{
		Description: "Random job title",
		Example:     "Director",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return JobTitle(), nil
		},
	})

	AddLookupData("company.job.level", Info{
		Description: "Random job level",
		Example:     "Assurance",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return JobLevel(), nil
		},
	})

	AddLookupData("company.job.descriptor", Info{
		Description: "Random job descriptor",
		Example:     "Central",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return JobDescriptor(), nil
		},
	})
}
