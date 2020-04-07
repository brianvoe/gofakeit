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

func ExampleMap() {
	Seed(11)
	fmt.Println(Map())
	// Output: map[eaque:784604.3 excepturi:Christy Ratke quo:99536 North Streamville, Rossieview, Hawaii 42591 repellat:776037 voluptatem:57704.402]
}

func TestMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		Map()
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map()
	}
}
