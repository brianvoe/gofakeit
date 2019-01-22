package gofakeit

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	numTests := 100

	for i := 0; i < numTests; i++ {
		Generate("{person.first} {person.last} {contact.email} #?#?#?")
	}
}

func ExampleGenerate() {
	Seed(11)
	fmt.Println(Generate("{person.first} {person.last} lives at {address.number} {address.street_name} {address.street_suffix}"))
	// Output: Markus Moen lives at 599 Garden mouth
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("{person.first} {person.last} {contact.email} #?#?#?")
	}
}
