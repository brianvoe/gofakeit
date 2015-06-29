package gofakeit

// UniqueID will generate a random unique identifier
// This should be taken lightly as this function just uses password generator
func UniqueID() string {
	return Password(true, true, true, false, false, 32)
}
