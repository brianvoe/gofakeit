package gofakeit

// FileExtension will generate a random file extension
func FileExtension() string {
	return getRandValue([]string{"file", "extension"})
}

// FileMimeType will generate a random mime file type
func FileMimeType() string {
	return getRandValue([]string{"file", "mime_type"})
}

func addFileLookup() {
	AddLookupData("fileextension", Info{
		Category:    "file",
		Description: "Random file extension",
		Example:     "nes",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FileExtension(), nil
		},
	})

	AddLookupData("filemimetype", Info{
		Category:    "file",
		Description: "Random file mime type",
		Example:     "application/json",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FileMimeType(), nil
		},
	})
}
