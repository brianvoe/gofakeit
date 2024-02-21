package gofakeit

import (
	"errors"
)

// Error will return a random generic error
func Error() error {
	return err(GlobalFaker)
}

// Error will return a random generic error
func (f *Faker) Error() error {
	return err(f)
}

func err(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "generic"}))
	return errors.New(genStr)
}

// ErrorObject will return a random error object word
func ErrorObject() error {
	return errorObject(GlobalFaker)
}

// ErrorObject will return a random error object word
func (f *Faker) ErrorObject() error {
	return errorObject(f)
}

func errorObject(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "object"}))
	return errors.New(genStr)
}

// ErrorDatabase will return a random database error
func ErrorDatabase() error {
	return errorDatabase(GlobalFaker)
}

// ErrorDatabase will return a random database error
func (f *Faker) ErrorDatabase() error {
	return errorDatabase(f)
}

func errorDatabase(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "database"}))
	return errors.New(genStr)
}

// ErrorGRPC will return a random gRPC error
func ErrorGRPC() error {
	return errorGRPC(GlobalFaker)
}

// ErrorGRPC will return a random gRPC error
func (f *Faker) ErrorGRPC() error {
	return errorGRPC(f)
}

func errorGRPC(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "grpc"}))
	return errors.New(genStr)
}

// ErrorHTTP will return a random HTTP error
func ErrorHTTP() error {
	return errorHTTP(GlobalFaker)
}

// ErrorHTTP will return a random HTTP error
func (f *Faker) ErrorHTTP() error {
	return errorHTTP(f)
}

func errorHTTP(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "http"}))
	return errors.New(genStr)
}

// ErrorHTTPClient will return a random HTTP client error response (400-418)
func ErrorHTTPClient() error {
	return errorHTTPClient(GlobalFaker)
}

// ErrorHTTPClient will return a random HTTP client error response (400-418)
func (f *Faker) ErrorHTTPClient() error {
	return errorHTTPClient(f)
}

func errorHTTPClient(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "http_client"}))
	return errors.New(genStr)
}

// ErrorHTTPServer will return a random HTTP server error response (500-511)
func ErrorHTTPServer() error {
	return errorHTTPServer(GlobalFaker)
}

// ErrorHTTPServer will return a random HTTP server error response (500-511)
func (f *Faker) ErrorHTTPServer() error {
	return errorHTTPServer(f)
}

func errorHTTPServer(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "http_server"}))
	return errors.New(genStr)
}

// ErrorRuntime will return a random runtime error
func ErrorRuntime() error {
	return errorRuntime(GlobalFaker)
}

// ErrorRuntime will return a random runtime error
func (f *Faker) ErrorRuntime() error {
	return errorRuntime(f)
}

func errorRuntime(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "runtime"}))
	return errors.New(genStr)
}

// ErrorValidation will return a random validation error
func ErrorValidation() error {
	return errorValidation(GlobalFaker)
}

// ErrorValidation will return a random validation error
func (f *Faker) ErrorValidation() error {
	return errorValidation(f)
}

func errorValidation(f *Faker) error {
	genStr, _ := generate(f, getRandValue(f, []string{"error", "validation"}))
	return errors.New(genStr)
}

func addErrorLookup() {
	AddFuncLookup("error", Info{
		Display:     "Error",
		Category:    "error",
		Description: "Message displayed by a computer or software when a problem or mistake is encountered",
		Example:     "syntax error",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return err(f), nil
		},
	})

	AddFuncLookup("errorobject", Info{
		Display:     "Error object word",
		Category:    "error",
		Description: "Various categories conveying details about encountered errors",
		Example:     "protocol",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorObject(f), nil
		},
	})

	AddFuncLookup("errordatabase", Info{
		Display:     "Database error",
		Category:    "error",
		Description: "A problem or issue encountered while accessing or managing a database",
		Example:     "sql error",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorDatabase(f), nil
		},
	})

	AddFuncLookup("errorgrpc", Info{
		Display:     "gRPC error",
		Category:    "error",
		Description: "Communication failure in the high-performance, open-source universal RPC framework",
		Example:     "client protocol error",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorGRPC(f), nil
		},
	})

	AddFuncLookup("errorhttp", Info{
		Display:     "HTTP error",
		Category:    "error",
		Description: "A problem with a web http request",
		Example:     "invalid method",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorHTTP(f), nil
		},
	})

	AddFuncLookup("errorhttpclient", Info{
		Display:     "HTTP client error",
		Category:    "error",
		Description: "Failure or issue occurring within a client software that sends requests to web servers",
		Example:     "request timeout",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorHTTPClient(f), nil
		},
	})

	AddFuncLookup("errorhttpserver", Info{
		Display:     "HTTP server error",
		Category:    "error",
		Description: "Failure or issue occurring within a server software that recieves requests from clients",
		Example:     "internal server error",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorHTTPServer(f), nil
		},
	})

	AddFuncLookup("errorruntime", Info{
		Display:     "Runtime error",
		Category:    "error",
		Description: "Malfunction occuring during program execution, often causing abrupt termination or unexpected behavior",
		Example:     "address out of bounds",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorRuntime(f), nil
		},
	})

	AddFuncLookup("errorvalidation", Info{
		Display:     "Validation error",
		Category:    "error",
		Description: "Occurs when input data fails to meet required criteria or format specifications",
		Example:     "missing required field",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return errorValidation(f), nil
		},
	})
}
