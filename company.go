package gofakeit

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

// Slogan will generate a random company slogan
func slogan(f *Faker) string {
	slogan := ""
	var sloganStyle = number(f, 0, 2)
	switch sloganStyle {
	// Noun. Buzzword!
	case 0:
		slogan = getRandValue(f, []string{"company", "blurb"}) + ". " + getRandValue(f, []string{"company", "buzzwords"}) + "!"
	// Buzzword Noun, Buzzword Noun.
	case 1:
		slogan = getRandValue(f, []string{"company", "buzzwords"}) + " " + getRandValue(f, []string{"company", "blurb"}) + ", " + getRandValue(f, []string{"company", "buzzwords"}) + " " + getRandValue(f, []string{"company", "blurb"}) + "."
	// Buzzword bs Noun, Buzzword.
	case 2:
		slogan = getRandValue(f, []string{"company", "buzzwords"}) + " " + getRandValue(f, []string{"company", "bs"}) + " " + getRandValue(f, []string{"company", "blurb"}) + ", " + getRandValue(f, []string{"company", "buzzwords"}) + "."
	}
	return slogan
}

func addCompanyLookup() {
	AddFuncLookup("company", Info{
		Display:     "Company",
		Category:    "company",
		Description: "Designated official name of a business or organization",
		Example:     "Moen, Pagac and Wuckert",
		Output:      "string",
		Aliases:     []string{"organization", "corporation", "enterprise", "firm", "official", "name"},
		Keywords:    []string{"company", "business", "organization", "corporation", "enterprise", "firm", "official", "name"},
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
		Aliases:     []string{"structure", "inc", "llc", "corporation", "ending", "legal"},
		Keywords:    []string{"company", "suffix", "business", "structure", "inc", "llc", "corporation", "end"},
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
		Aliases:     []string{"company", "word", "corporate", "random", "nonsense"},
		Keywords:    []string{"bs", "company", "word", "business", "corporate", "random"},
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
		Aliases:     []string{"summary", "company", "purpose", "brief", "overview"},
		Keywords:    []string{"blurb", "description", "summary", "company", "purpose", "products", "services"},
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
		Aliases:     []string{"term", "corporate", "impressive", "jargon", "fashionable"},
		Keywords:    []string{"buzzword", "trendy", "business", "term", "corporate", "impressive"},
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
		Aliases:     []string{"position", "role", "employment", "work", "career"},
		Keywords:    []string{"job", "position", "role", "employment", "work", "tasks", "responsibilities"},
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
		Aliases:     []string{"title", "position", "role", "employment", "occupation"},
		Keywords:    []string{"job", "title", "position", "role", "employment", "company", "organization"},
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
		Aliases:     []string{"descriptor", "duties", "requirements", "nature", "responsibilities"},
		Keywords:    []string{"job", "descriptor", "duties", "requirements", "nature", "description"},
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
		Aliases:     []string{"position", "rank", "seniority", "tier", "grade"},
		Keywords:    []string{"job", "level", "position", "rank", "seniority", "random"},
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
		Aliases:     []string{"catchphrase", "motto", "brand", "values", "tagline"},
		Keywords:    []string{"slogan", "catchphrase", "motto", "brand", "values", "company"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return slogan(f), nil
		},
	})
}
