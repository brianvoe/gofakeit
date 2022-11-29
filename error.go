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

func HTTPError() error {
	return httpErr(globalFaker.Rand)
}

func (f *Faker) HTTPError() error {
	return httpErr(f.Rand)
}

func httpErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http"})
	return &fakeError{msg}
}

func HTTPClientError() error {
	return httpClientErr(globalFaker.Rand)
}

func (f *Faker) HTTPClientError() error {
	return httpClientErr(f.Rand)
}

func httpClientErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http-client"})
	return &fakeError{msg}
}

func HTTPServerError() error {
	return httpServerErr(globalFaker.Rand)
}

func (f *Faker) HTTPServerError() error {
	return httpServerErr(f.Rand)
}

func httpServerErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "http-server"})
	return &fakeError{msg}
}

func DatabaseError() error {
	return databaseErr(globalFaker.Rand)
}

func (f *Faker) DatabaseError() error {
	return databaseErr(f.Rand)
}

func databaseErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "database"})
	return &fakeError{msg}
}

func GRPCError() error {
	return grpcErr(globalFaker.Rand)
}

func (f *Faker) GRPCError() error {
	return grpcErr(f.Rand)
}

func grpcErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "grpc"})
	return &fakeError{msg}
}

func RuntimeError() error {
	return runtimeErr(globalFaker.Rand)
}

func (f *Faker) RuntimeError() error {
	return runtimeErr(f.Rand)
}

func runtimeErr(r *rand.Rand) error {
	msg := getRandValue(r, []string{"error", "runtime"})
	return &fakeError{msg}
}
func InputError() error {
	return inputErr(globalFaker.Rand)
}

func (f *Faker) InputError() error {
	return inputErr(f.Rand)
}

func inputErr(r *rand.Rand) error {
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
