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

// JobInfo is a struct of job information
type JobInfo struct {
	Company    string `json:"company"`
	Title      string `json:"title"`
	Descriptor string `json:"descriptor"`
	Level      string `json:"level"`
}

// Job will generate a struct with random job information
func Job() *JobInfo {
	return &JobInfo{
		Company:    Company(),
		Title:      JobTitle(),
		Descriptor: JobDescriptor(),
		Level:      JobLevel(),
	}
}

// JobTitle will generate a random job title string
func JobTitle() string {
	return getRandValue([]string{"job", "title"})
}

// JobDescriptor will generate a random job descriptor string
func JobDescriptor() string {
	return getRandValue([]string{"job", "descriptor"})
}

// JobLevel will generate a random job level string
func JobLevel() string {
	return getRandValue([]string{"job", "level"})
}

func addCompanyLookup() {
	AddLookupData("company", Info{
		Category:    "company",
		Description: "Random company name",
		Example:     "Moen, Pagac and Wuckert",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Company(), nil
		},
	})

	AddLookupData("companysuffix", Info{
		Category:    "company",
		Description: "Random company name suffix",
		Example:     "Inc",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CompanySuffix(), nil
		},
	})

	AddLookupData("bs", Info{
		Category:    "company",
		Description: "Random bs company word",
		Example:     "front-end",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BS(), nil
		},
	})

	AddLookupData("buzzword", Info{
		Category:    "company",
		Description: "Random company buzzwords",
		Example:     "disintermediate",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BuzzWord(), nil
		},
	})

	AddLookupData("job", Info{
		Category:    "company",
		Description: "Random job data set",
		Example:     `{company: "Moen, Pagac and Wuckert", title: "Director", descriptor: "Central", level: "Assurance"}`,
		Output:      "map[string]string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Job(), nil
		},
	})

	AddLookupData("jobtitle", Info{
		Category:    "company",
		Description: "Random job title",
		Example:     "Director",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobTitle(), nil
		},
	})

	AddLookupData("jobdescriptor", Info{
		Category:    "company",
		Description: "Random job descriptor",
		Example:     "Central",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobDescriptor(), nil
		},
	})

	AddLookupData("joblevel", Info{
		Category:    "company",
		Description: "Random job level",
		Example:     "Assurance",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobLevel(), nil
		},
	})
}
