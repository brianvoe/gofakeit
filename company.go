package gofakeit

// Company will generate a random company name string
func Company() (company string) {
	return getRandValue([]string{"company", "name"})
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
	Company    string `json:"company" xml:"company"`
	Title      string `json:"title" xml:"title"`
	Descriptor string `json:"descriptor" xml:"descriptor"`
	Level      string `json:"level" xml:"level"`
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
	AddFuncLookup("company", Info{
		Display:     "Company",
		Category:    "company",
		Description: "Random company name",
		Example:     "Moen, Pagac and Wuckert",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Company(), nil
		},
	})

	AddFuncLookup("companysuffix", Info{
		Display:     "Company Suffix",
		Category:    "company",
		Description: "Random company name suffix",
		Example:     "Inc",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CompanySuffix(), nil
		},
	})

	AddFuncLookup("bs", Info{
		Display:     "BS",
		Category:    "company",
		Description: "Random bs company word",
		Example:     "front-end",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BS(), nil
		},
	})

	AddFuncLookup("buzzword", Info{
		Display:     "Buzzword",
		Category:    "company",
		Description: "Random company buzzwords",
		Example:     "disintermediate",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BuzzWord(), nil
		},
	})

	AddFuncLookup("job", Info{
		Display:     "Job",
		Category:    "company",
		Description: "Random job data set",
		Example:     `{company: "Moen, Pagac and Wuckert", title: "Director", descriptor: "Central", level: "Assurance"}`,
		Output:      "map[string]string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Job(), nil
		},
	})

	AddFuncLookup("jobtitle", Info{
		Display:     "Job Title",
		Category:    "company",
		Description: "Random job title",
		Example:     "Director",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobTitle(), nil
		},
	})

	AddFuncLookup("jobdescriptor", Info{
		Display:     "Job Descriptor",
		Category:    "company",
		Description: "Random job descriptor",
		Example:     "Central",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobDescriptor(), nil
		},
	})

	AddFuncLookup("joblevel", Info{
		Display:     "Job Level",
		Category:    "company",
		Description: "Random job level",
		Example:     "Assurance",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return JobLevel(), nil
		},
	})
}
