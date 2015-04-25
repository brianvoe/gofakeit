package gofakeit

// Generate Job
func Job() string {
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
