package gofakeit

// Generate Random File Mime Type
func MimeType() string {
	return getRandValue([]string{"file", "mime_type"})
}

// Generate Random File Extension
func Extension() string {
	return getRandValue([]string{"file", "extension"})
}
