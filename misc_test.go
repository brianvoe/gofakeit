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
