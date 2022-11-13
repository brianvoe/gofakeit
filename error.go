package gofakeit

import "math/rand"

type fakeError struct {
	msg string
}

func (f *fakeError) Error() string {
	return f.msg
}

// Error will return a random error
func Error() error {
	return err(globalFaker.Rand)
}

// Error will return a random error
func (f *Faker) Error() error {
	return err(f.Rand)
}

func err(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "message"})
	return &fakeError{msg}
}
