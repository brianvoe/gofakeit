package gofakeit

type JobInfo struct {
	Company    string
	Title      string
	Descriptor string
	Level      string
}

// Generate JobInfo Struct
func Job() *JobInfo {
	return &JobInfo{
		Company:    Company(),
		Title:      JobTitle(),
		Descriptor: JobDescriptor(),
		Level:      JobLevel(),
	}
}

// Generate Job Title
func JobTitle() string {
	return getRandValue([]string{"job", "title"})
}

// Generate Job Descriptor
func JobDescriptor() string {
	return getRandValue([]string{"job", "descriptor"})
}

// Generate Job Level
func JobLevel() string {
	return getRandValue([]string{"job", "level"})
}
