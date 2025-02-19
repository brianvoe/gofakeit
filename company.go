package gofakeit

import "strings"

// Company will generate a random company name string
func Company() string { return company(GlobalFaker) }

// Company will generate a random company name string
func (f *Faker) Company() string { return company(f) }

func company(f *Faker) string { return getRandValue(f, []string{"company", "name"}) }

// CompanySuffix will generate a random company suffix string
func CompanySuffix() string { return companySuffix(GlobalFaker) }

// CompanySuffix will generate a random company suffix string
func (f *Faker) CompanySuffix() string { return companySuffix(f) }

func companySuffix(f *Faker) string { return getRandValue(f, []string{"company", "suffix"}) }

// Blurb will generate a random company blurb string
func Blurb() string { return blurb(GlobalFaker) }

func (f *Faker) Blurb() string { return blurb(f) }

func blurb(f *Faker) string { return getRandValue(f, []string{"company", "blurb"}) }

// BuzzWord will generate a random company buzz word string
func BuzzWord() string { return buzzWord(GlobalFaker) }

// BuzzWord will generate a random company buzz word string
func (f *Faker) BuzzWord() string { return buzzWord(f) }

func buzzWord(f *Faker) string { return getRandValue(f, []string{"company", "buzzwords"}) }

// BS will generate a random company bs string
func BS() string { return bs(GlobalFaker) }

// BS will generate a random company bs string
func (f *Faker) BS() string { return bs(f) }

func bs(f *Faker) string { return getRandValue(f, []string{"company", "bs"}) }

// JobInfo is a struct of job information
type JobInfo struct {
	Company    string `json:"company" xml:"company"`
	Title      string `json:"title" xml:"title"`
	Descriptor string `json:"descriptor" xml:"descriptor"`
	Level      string `json:"level" xml:"level"`
}

// Job will generate a struct with random job information
func Job() *JobInfo { return job(GlobalFaker) }

// Job will generate a struct with random job information
func (f *Faker) Job() *JobInfo { return job(f) }

func job(f *Faker) *JobInfo {
	return &JobInfo{
		Company:    company(f),
		Title:      jobTitle(f),
		Descriptor: jobDescriptor(f),
		Level:      jobLevel(f),
	}
}

// JobTitle will generate a random job title string
func JobTitle() string { return jobTitle(GlobalFaker) }

// JobTitle will generate a random job title string
func (f *Faker) JobTitle() string { return jobTitle(f) }

func jobTitle(f *Faker) string { return getRandValue(f, []string{"job", "title"}) }

// JobDescriptor will generate a random job descriptor string
func JobDescriptor() string { return jobDescriptor(GlobalFaker) }

// JobDescriptor will generate a random job descriptor string
func (f *Faker) JobDescriptor() string { return jobDescriptor(f) }

func jobDescriptor(f *Faker) string { return getRandValue(f, []string{"job", "descriptor"}) }

// JobLevel will generate a random job level string
func JobLevel() string { return jobLevel(GlobalFaker) }

// JobLevel will generate a random job level string
func (f *Faker) JobLevel() string { return jobLevel(f) }

func jobLevel(f *Faker) string { return getRandValue(f, []string{"job", "level"}) }

// Slogan will generate a random company slogan
func Slogan() string { return slogan(GlobalFaker) }

// Slogan will generate a random company slogan
func (f *Faker) Slogan() string { return slogan(f) }

var sloganWords = map[string][]string{
	"NounSingular": {
		"Technology", "Success", "Strategy", "Leadership", "Trust", "Transformation", "Creativity", "Performance", "Progress",
	},
	"VerbSingular": {
		"Empowers", "Creates", "Drives", "Shapes", "Builds", "Transforms", "Inspires", "Leads", "Delivers",
	},
	"Adjective": {
		"Impactful", "Trusted", "Innovative", "Strategic", "Creative", "Visionary",
	},
	"Determiner": {
		"Our", "Your",
	},
	"Interjection": {
		"Wow", "Amazing",
	},
}

// Slogan will generate a random company slogan
func slogan(f *Faker) string {
	// Predefined slogan structures (Mad Libs style)
	var structures = [][]string{
		{"Adjective", "NounSingular", "Empowers"},
		{"NounSingular", "Transforms", "Adjective", "NounSingular"},
		{"NounSingular", "Drives", "Success"},
		{"Adjective", "NounSingular", "Inspires"},
		{"NounSingular", "Shapes", "Adjective", "Progress"},

		{"Interjection", "!", "Adjective", "NounSingular"},
		{"Interjection", "!", "NounSingular", "Leads"},

		{"VerbSingular", "Adjective", "NounSingular"},
		{"Adjective", "NounSingular", "Leads"},
		{"Determiner", "Adjective", "NounSingular", "Delivers"},
		{"Determiner", "NounSingular", "Inspires", "Adjective", "NounSingular"},

		{"NounSingular", "Inspires"},
		{"NounSingular", "Empowers"},
		{"VerbSingular", "NounSingular"},
		{"Adjective", "NounSingular"},
	}

	// Select a random structure
	structure := structures[number(f, 0, len(structures)-1)]

	// Build the slogan
	var sloganParts []string
	for _, wordType := range structure {
		switch wordType {
		case "NounSingular":
			sloganParts = append(sloganParts, sloganWords["NounSingular"][number(f, 0, len(sloganWords["NounSingular"])-1)])
		case "VerbSingular":
			sloganParts = append(sloganParts, sloganWords["VerbSingular"][number(f, 0, len(sloganWords["VerbSingular"])-1)])
		case "Adjective":
			sloganParts = append(sloganParts, sloganWords["Adjective"][number(f, 0, len(sloganWords["Adjective"])-1)])
		case "Determiner":
			sloganParts = append(sloganParts, sloganWords["Determiner"][number(f, 0, len(sloganWords["Determiner"])-1)])
		case "Interjection":
			sloganParts = append(sloganParts, sloganWords["Interjection"][number(f, 0, len(sloganWords["Interjection"])-1)])
		default:
			// Directly append non-placeholder words
			sloganParts = append(sloganParts, wordType)
		}
	}

	// Combine words into a sentence
	slogan := strings.Join(sloganParts, " ")

	// Capitalize the first letter
	slogan = title(slogan)

	// Add a period to the end
	slogan = slogan + "."

	return slogan
}

func addCompanyLookup() {
	AddFuncLookup("company", Info{
		Display:     "Company",
		Category:    "company",
		Description: "Designated official name of a business or organization",
		Example:     "Moen, Pagac and Wuckert",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return company(f), nil
		},
	})

	AddFuncLookup("companysuffix", Info{
		Display:     "Company Suffix",
		Category:    "company",
		Description: "Suffix at the end of a company name, indicating business structure, like 'Inc.' or 'LLC'",
		Example:     "Inc",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return companySuffix(f), nil
		},
	})

	AddFuncLookup("bs", Info{
		Display:     "BS",
		Category:    "company",
		Description: "Random bs company word",
		Example:     "front-end",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bs(f), nil
		},
	})

	AddFuncLookup("blurb", Info{
		Display:     "Blurb",
		Category:    "company",
		Description: "Brief description or summary of a company's purpose, products, or services",
		Example:     "word",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return blurb(f), nil
		},
	})

	AddFuncLookup("buzzword", Info{
		Display:     "Buzzword",
		Category:    "company",
		Description: "Trendy or overused term often used in business to sound impressive",
		Example:     "disintermediate",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return buzzWord(f), nil
		},
	})

	AddFuncLookup("job", Info{
		Display:     "Job",
		Category:    "company",
		Description: "Position or role in employment, involving specific tasks and responsibilities",
		Example: `{
	"company": "ClearHealthCosts",
	"title": "Agent",
	"descriptor": "Future",
	"level": "Tactics"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return job(f), nil
		},
	})

	AddFuncLookup("jobtitle", Info{
		Display:     "Job Title",
		Category:    "company",
		Description: "Specific title for a position or role within a company or organization",
		Example:     "Director",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return jobTitle(f), nil
		},
	})

	AddFuncLookup("jobdescriptor", Info{
		Display:     "Job Descriptor",
		Category:    "company",
		Description: "Word used to describe the duties, requirements, and nature of a job",
		Example:     "Central",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return jobDescriptor(f), nil
		},
	})

	AddFuncLookup("joblevel", Info{
		Display:     "Job Level",
		Category:    "company",
		Description: "Random job level",
		Example:     "Assurance",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return jobLevel(f), nil
		},
	})

	AddFuncLookup("slogan", Info{
		Display:     "Slogan",
		Category:    "company",
		Description: "Catchphrase or motto used by a company to represent its brand or values",
		Example:     "Universal seamless Focus, interactive.",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return slogan(f), nil
		},
	})
}
