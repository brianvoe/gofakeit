package gofakeit

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	numTests := 100
	for i := 0; i < numTests; i++ {
		Generate("{firstname} {lastname} {email} #?#?#?")
	}
}

func ExampleGenerate() {
	Seed(11)
	fmt.Println(Generate("{firstname} {lastname} ssn is {ssn} and lives at {street}"))
	// Output: Markus Moen ssn is 952284213 and lives at 599 New Cliffsstad
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("{firstname} {lastname} {email} #?#?#?")
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
