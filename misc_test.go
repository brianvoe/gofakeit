package gofakeit

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v5/data"
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
