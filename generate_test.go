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
	fmt.Println(Generate("{sentence:3}"))
	fmt.Println(Generate("{shuffleints:[1,2,3]}"))
	fmt.Println(Generate("{shufflestrings:[key:value,int:string,1:2,a:b]}"))
	// Output: Markus Moen ssn is 952284213 and lives at 599 New Cliffsstad
	// Arrival tour security.
	// [1 3 2]
	// [1:2 key:value int:string a:b]
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("{firstname} {lastname} {email} #?#?#?")
	}
}

func ExampleMap() {
	Seed(11)

	fmt.Println(Map())
	// Output: map[approach:map[mind:[arrival should resolve outcome hurt]] arrive:Coordinator consult:respond context:9285735 water:5081652]
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
