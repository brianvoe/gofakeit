package gofakeit

// Severity will generate a random severity level
func Severity() string {
	return getRandValue([]string{"severity", "level"})
}
