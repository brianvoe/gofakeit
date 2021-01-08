package gofakeit

import rand "math/rand"

// Company will generate a random company name string
func Company() string { return company(globalFaker.Rand) }

// Company will generate a random company name string
func (f *Faker) Company() string { return company(f.Rand) }

func company(r *rand.Rand) string { return getRandValue(r, []string{"company", "name"}) }

// CompanySuffix will generate a random company suffix string
func CompanySuffix() string { return companySuffix(globalFaker.Rand) }

// CompanySuffix will generate a random company suffix string
func (f *Faker) CompanySuffix() string { return companySuffix(f.Rand) }

func companySuffix(r *rand.Rand) string { return getRandValue(r, []string{"company", "suffix"}) }

// BuzzWord will generate a random company buzz word string
func BuzzWord() string { return buzzWord(globalFaker.Rand) }

// BuzzWord will generate a random company buzz word string
func (f *Faker) BuzzWord() string { return buzzWord(f.Rand) }

func buzzWord(r *rand.Rand) string { return getRandValue(r, []string{"company", "buzzwords"}) }

// BS will generate a random company bs string
func BS() string { return bs(globalFaker.Rand) }

// BS will generate a random company bs string
func (f *Faker) BS() string { return bs(f.Rand) }

func bs(r *rand.Rand) string { return getRandValue(r, []string{"company", "bs"}) }

// JobInfo is a struct of job information
type JobInfo struct {
	Company    string `json:"company" xml:"company"`
	Title      string `json:"title" xml:"title"`
	Descriptor string `json:"descriptor" xml:"descriptor"`
	Level      string `json:"level" xml:"level"`
}

// Job will generate a struct with random job information
func Job() *JobInfo { return job(globalFaker.Rand) }

// Job will generate a struct with random job information
func (f *Faker) Job() *JobInfo { return job(f.Rand) }

func job(r *rand.Rand) *JobInfo {
	return &JobInfo{
		Company:    company(r),
		Title:      jobTitle(r),
		Descriptor: jobDescriptor(r),
		Level:      jobLevel(r),
	}
}

// JobTitle will generate a random job title string
func JobTitle() string { return jobTitle(globalFaker.Rand) }

// JobTitle will generate a random job title string
func (f *Faker) JobTitle() string { return jobTitle(f.Rand) }

func jobTitle(r *rand.Rand) string { return getRandValue(r, []string{"job", "title"}) }

// JobDescriptor will generate a random job descriptor string
func JobDescriptor() string { return jobDescriptor(globalFaker.Rand) }

// JobDescriptor will generate a random job descriptor string
func (f *Faker) JobDescriptor() string { return jobDescriptor(f.Rand) }

func jobDescriptor(r *rand.Rand) string { return getRandValue(r, []string{"job", "descriptor"}) }

// JobLevel will generate a random job level string
func JobLevel() string { return jobLevel(globalFaker.Rand) }

// JobLevel will generate a random job level string
func (f *Faker) JobLevel() string { return jobLevel(f.Rand) }

func jobLevel(r *rand.Rand) string { return getRandValue(r, []string{"job", "level"}) }

func addCompanyLookup() {
	AddFuncLookup("company", Info{
		Display:     "Company",
		Category:    "company",
		Description: "Random company name",
		Example:     "Moen, Pagac and Wuckert",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return company(r), nil
		},
	})

	AddFuncLookup("companysuffix", Info{
		Display:     "Company Suffix",
		Category:    "company",
		Description: "Random company name suffix",
		Example:     "Inc",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return companySuffix(r), nil
		},
	})

	AddFuncLookup("bs", Info{
		Display:     "BS",
		Category:    "company",
		Description: "Random bs company word",
		Example:     "front-end",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return bs(r), nil
		},
	})

	AddFuncLookup("buzzword", Info{
		Display:     "Buzzword",
		Category:    "company",
		Description: "Random company buzzwords",
		Example:     "disintermediate",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return buzzWord(r), nil
		},
	})

	AddFuncLookup("job", Info{
		Display:     "Job",
		Category:    "company",
		Description: "Random job data set",
		Example:     `{company: "Moen, Pagac and Wuckert", title: "Director", descriptor: "Central", level: "Assurance"}`,
		Output:      "map[string]string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return job(r), nil
		},
	})

	AddFuncLookup("jobtitle", Info{
		Display:     "Job Title",
		Category:    "company",
		Description: "Random job title",
		Example:     "Director",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return jobTitle(r), nil
		},
	})

	AddFuncLookup("jobdescriptor", Info{
		Display:     "Job Descriptor",
		Category:    "company",
		Description: "Random job descriptor",
		Example:     "Central",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return jobDescriptor(r), nil
		},
	})

	AddFuncLookup("joblevel", Info{
		Display:     "Job Level",
		Category:    "company",
		Description: "Random job level",
		Example:     "Assurance",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return jobLevel(r), nil
		},
	})
}
