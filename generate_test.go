package gofakeit

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	var numTests int = 100

	for i := 0; i < numTests; i++ {
		Generate("{name.first} {name.last} {contact.email} #?#?#?")
	}
}

func ExampleGenerate() {
	Seed(11)
	fmt.Println(Generate("{name.first} {name.last} lives at {address.number} {address.street_name} {address.street_suffix}"))
	// Output: Markus Moen lives at 715 Garden mouth
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("{name.first} {name.last} {contact.email} #?#?#?")
	}
}
