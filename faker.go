package gofakeit

import (
	"math/rand/v2"
	"reflect"

	"github.com/brianvoe/gofakeit/v6/source"
)

// Seed will set the global random value if Seed exists
func Seed(args ...any) {
	// Ensure GlobalFaker.Rand is not nil
	if GlobalFaker.Rand == nil {
		return
	}

	// The standard Seed method on rand.Source expects a single int64 argument.
	// This code checks for custom Seed implementations with different parameters.
	method := reflect.ValueOf(GlobalFaker.Rand).MethodByName("Seed")

	// Check if a Seed method is found that is not the standard one (assuming custom implementations)
	if method.IsValid() && len(args) > 0 {
		// Prepare arguments for the reflection call
		callArgs := make([]reflect.Value, len(args))
		for i, arg := range args {
			callArgs[i] = reflect.ValueOf(arg)
		}

		// Dynamically call the Seed method
		method.Call(callArgs)
		return
	}
}

// Create global variable to deal with global function call
var GlobalFaker *Faker = New(0)

// Faker struct is the primary struct for using localized
type Faker struct {
	Rand *rand.Rand
}

// New will set the default pcg locked rand for concurrent safe random usage
func New(seed uint64) *Faker {
	return &Faker{Rand: rand.New(source.NewPCG(seed, seed, true))}
}

// NewCustom will utilize a custom rand.Source
func NewSource(source rand.Source) *Faker {
	return &Faker{Rand: rand.New(source)}
}
