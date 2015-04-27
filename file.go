package gofakeit

func MimeType() string {
	return getRandValue([]string{"file", "mime_type"})
}

func Extension() string {
	return getRandValue([]string{"file", "extension"})
}
