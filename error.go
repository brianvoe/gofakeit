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
	msg := generate(r, getRandValue(r, []string{"error", "generic"}))
	return &fakeError{msg}
}

func ErrorHTTP() error {
	return errorHTTP(globalFaker.Rand)
}

func (f *Faker) ErrorHTTP() error {
	return errorHTTP(f.Rand)
}

func errorHTTP(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http"})
	return &fakeError{msg}
}

func ErrorHTTPClient() error {
	return errorHTTPClient(globalFaker.Rand)
}

func (f *Faker) ErrorHTTPClient() error {
	return errorHTTPClient(f.Rand)
}

func errorHTTPClient(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http-client"})
	return &fakeError{msg}
}

func ErrorHTTPServer() error {
	return errorHTTPServer(globalFaker.Rand)
}

func (f *Faker) ErrorHTTPServer() error {
	return errorHTTPServer(f.Rand)
}

func errorHTTPServer(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http-server"})
	return &fakeError{msg}
}

func ErrorDatabase() error {
	return errorDatabase(globalFaker.Rand)
}

func (f *Faker) ErrorDatabase() error {
	return errorDatabase(f.Rand)
}

func errorDatabase(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "database"})
	return &fakeError{msg}
}

func ErrorGRPC() error {
	return errorGRPC(globalFaker.Rand)
}

func (f *Faker) ErrorGRPC() error {
	return errorGRPC(f.Rand)
}

func errorGRPC(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "grpc"})
	return &fakeError{msg}
}

func ErrorRuntime() error {
	return errorRuntime(globalFaker.Rand)
}

func (f *Faker) ErrorRuntime() error {
	return errorRuntime(f.Rand)
}

func errorRuntime(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "runtime"})
	return &fakeError{msg}
}
func ErrorInput() error {
	return errorInput(globalFaker.Rand)
}

func (f *Faker) ErrorInput() error {
	return errorInput(f.Rand)
}

func errorInput(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "input"})
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
