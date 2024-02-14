package gofakeit

import (
	"errors"
	"math/rand/v2"
	"reflect"
	"sync"
)

// Create global variable to deal with global function call
var GlobalFaker *Faker = New(0)

// Faker struct is the primary struct for using localized
type Faker struct {
	Rand rand.Source

	// Lock to make thread safe
	Locked bool
	mu     sync.Mutex
}

// New creates and returns a new Faker struct seeded with a given seed
// using the PCG algorithm in lock mode for thread safety
func New(seed uint64) *Faker {
	return &Faker{
		Rand:   rand.NewPCG(seed, seed),
		Locked: true,
	}
}

// NewFaker takes in a rand.Source and thread lock state and returns a new Faker struct
func NewFaker(src rand.Source, lock bool) *Faker {
	return &Faker{Rand: src}
}

// Seed attempts to seed the Faker with the given seed
func (f *Faker) Seed(args ...any) error {
	// Lock if locked
	if f.Locked {
		f.mu.Lock()
		defer f.mu.Unlock()
	}

	// Ensure GlobalFaker is not nil
	if GlobalFaker.Rand == nil {
		return errors.New("GlobalFaker.Rand is nil")
	}

	// The standard Seed method on rand.Source expects a single int64 argument.
	// This code checks for custom Seed implementations with different parameters.
	method := reflect.ValueOf(GlobalFaker.Rand).MethodByName("Seed")

	// Check if a Seed method is found
	if !method.IsValid() {
		return errors.New("Seed method not found")
	}

	// Adjust args if method requires exactly 2 args but only 1 was provided
	if method.Type().NumIn() == 2 && len(args) == 1 {
		args = append(args, args[0]) // Double the first value if only one is provided
	}

	// Convert args to the expected type by the Seed method
	convertedArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		// Assume the method expects uint64, convert the argument
		if argInt, ok := arg.(int); ok {
			convertedArgs[i] = reflect.ValueOf(uint64(argInt))
		} else {
			// For simplicity, if it's not an int, just reflect the value directly
			// In a real-world scenario, you might need to handle other types and conversions
			convertedArgs[i] = reflect.ValueOf(arg)
		}
	}

	// Dynamically call the Seed method with converted arguments
	method.Call(convertedArgs)

	return nil
}

// Seed attempts to seed the GlobalFaker with the given seed
func Seed(args ...any) error {
	return GlobalFaker.Seed(args...)
}
