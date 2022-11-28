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
	return fakeErr(globalFaker.Rand)
}

// Error will return a random error
func (f *Faker) Error() error {
	return fakeErr(f.Rand)
}

func fakeErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "generic"})
	return &fakeError{msg}
}

func HTTP() error {
	return httpErr(globalFaker.Rand)
}

func (f *Faker) HTTP() error {
	return httpErr(f.Rand)
}

func httpErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http"})
	return &fakeError{msg}
}

func addErrorLookup() {
	AddFuncLookup("object", Info{
		Display:     "Error object",
		Category:    "error",
		Description: "Random object causing an error",
		Example:     "request",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return getRandValue(r, []string{"error", "object"}), nil
		},
	})
	AddFuncLookup("verb", Info{
		Display:     "Error verb",
		Category:    "error",
		Description: "Random verb describing an error",
		Example:     "parse",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return getRandValue(r, []string{"error", "verb"}), nil
		},
	})
	AddFuncLookup("inputField", Info{
		Display:     "Input field error",
		Category:    "error",
		Description: "Random input field error message",
		Example:     "firstName",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return getRandValue(r, []string{"error", "inputField"}), nil
		},
	})
	AddFuncLookup("httpMethod", Info{
		Display:     "HTTP Method",
		Category:    "error",
		Description: "Random HTTP method",
		Example:     "GET",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return getRandValue(r, []string{"error", "httpMethod"}), nil
		},
	})
}
