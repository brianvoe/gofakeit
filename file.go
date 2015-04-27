package gofakeit

func MimeType() string {
	return getRandValue([]string{"file", "mime_type"})
}
