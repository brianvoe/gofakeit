package gofakeit

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v4/data"
)

func ExampleBool() {
	Seed(11)
	fmt.Println(Bool())
	// Output: false
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool()
	}
}

func TestUUID(t *testing.T) {
	id := UUID()

	if len(id) != 36 {
		t.Error("unique length does not equal requested length")
	}
}

func ExampleUUID() {
	Seed(11)
	fmt.Println(UUID())
	// Output: 590c1440-9888-45b0-bd51-a817ee07c3f2
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func TestCategories(t *testing.T) {
	var got, expected []string
	for k := range Categories() {
		got = append(got, k)
	}
	for k := range data.Data {
		expected = append(expected, k)
	}
	sort.Strings(got)
	sort.Strings(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Error("Type arrays are not the same.\nExpected: ", expected, "\nGot: ", got)
	}
}

func BenchmarkReplaceWithNumbers(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		Seed(42)

		b.StartTimer()
		replaceWithNumbers("###☺#☻##☹##")
		b.StopTimer()
	}
}

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
