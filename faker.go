package gofakeit

import (
	"fmt"
	"math/rand/v2"
	"reflect"

	"github.com/brianvoe/gofakeit/v6/source"
)

// Create global variable to deal with global function call
var GlobalFaker *Faker = New(0)

// Faker struct is the primary struct for using localized
type Faker struct {
	Rand rand.Source
}

func New(seed uint64) *Faker {
	pcg := source.NewPCG(seed, seed, true)

	return &Faker{
		Rand: rand.New(pcg),
	}
}

// NewFaker takes in a rand.Source and returns a new Faker struct
func NewFaker(src rand.Source) *Faker {
	return &Faker{Rand: src}
}

func Seed(args ...any) {
	// Ensure GlobalFaker is not nil
	if GlobalFaker == nil {
		return
	}

	// The standard Seed method on rand.Source expects a single int64 argument.
	// This code checks for custom Seed implementations with different parameters.
	method := reflect.ValueOf(GlobalFaker).MethodByName("Seed")

	fmt.Printf("method: %v\n", method)

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
