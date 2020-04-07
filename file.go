package gofakeit

// MimeType will generate a random mime file type
func MimeType() string {
	return getRandValue([]string{"file", "mime_type"})
}

// Extension will generate a random file extension
func Extension() string {
	return getRandValue([]string{"file", "extension"})
}

func addFileLookup() {
	AddLookupData("file.extension", Info{
		Description: "Random file extension",
		Example:     "nes",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Extension(), nil
		},
	})

	AddLookupData("file.mimetype", Info{
		Description: "Random file mime type",
		Example:     "application/json",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return MimeType(), nil
		},
	})
}
