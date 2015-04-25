package gofakeit

// Generate unique id
// This should be taken lightly as this function just uses password generator
func UniqueId() string {
	return Password(true, true, true, false, false, 32)
}
