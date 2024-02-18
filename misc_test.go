package gofakeit

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/v7/data"
)

func ExampleBool() {
	Seed(11)
	fmt.Println(Bool())

	// Output: false
}

func ExampleFaker_Bool() {
	f := New(11)
	fmt.Println(f.Bool())

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
		t.Error(id)
		t.Error("unique length does not equal requested length")
	}

	// Checking for race conditions, need to run --race
	for i := 0; i < 10000; i++ {
		go func() {
			_ = UUID()
		}()
	}
}

func ExampleUUID() {
	Seed(11)
	fmt.Println(UUID())

	// Output: b412b5fb-33a4-498e-9503-21c6b7e01dcf
}

func ExampleFaker_UUID() {
	f := New(11)
	fmt.Println(f.UUID())

	// Output: b412b5fb-33a4-498e-9503-21c6b7e01dcf
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}

func TestShuffleAnySlice(t *testing.T) {
	ShuffleAnySlice(nil)           // Should do nothing
	ShuffleAnySlice("b")           // Should do nothing
	ShuffleAnySlice([]string{"b"}) // If single value should do nothing

	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := make([]string, len(a))
	copy(b, a)
	ShuffleAnySlice(a)
	if equalSliceString(a, b) {
		t.Errorf("shuffle strings resulted in the same permutation, the odds are slim")
	}

	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	m := make([]int, len(n))
	copy(m, n)
	ShuffleAnySlice(n)
	if equalSliceInt(n, m) {
		t.Errorf("shuffle ints resulted in the same permutation, the odds are slim")
	}

	i := []any{"a", 1, "c", 3, []string{"a", "b", "c"}, -555, []byte{1, 5}, "h"}
	ii := make([]any, len(i))
	copy(ii, i)
	ShuffleAnySlice(i)
	if equalSliceInterface(i, ii) {
		t.Errorf("shuffle interface resulted in the same permutation, the odds are slim")
	}
}

func ExampleShuffleAnySlice() {
	Seed(11)

	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	ShuffleAnySlice(strings)
	fmt.Println(strings)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	ShuffleAnySlice(ints)
	fmt.Println(ints)

	// Output: [for day happy everyone good times a have]
	// [854 52 74125 941 777 8413 841657 89416]
}

func ExampleFaker_ShuffleAnySlice() {
	f := New(11)

	strings := []string{"happy", "times", "for", "everyone", "have", "a", "good", "day"}
	f.ShuffleAnySlice(strings)
	fmt.Println(strings)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	f.ShuffleAnySlice(ints)
	fmt.Println(ints)

	// Output: [for day happy everyone good times a have]
	// [854 52 74125 941 777 8413 841657 89416]
}

func BenchmarkShuffleAnySlice(b *testing.B) {
	a := []any{"a", 1, "c", 3, []string{"a", "b", "c"}, -555, []byte{1, 5}, "h"}
	for i := 0; i < b.N; i++ {
		ShuffleAnySlice(a)
	}
}

func ExampleFlipACoin() {
	Seed(11)
	fmt.Println(FlipACoin())

	// Output: Tails
}

func ExampleFaker_FlipACoin() {
	f := New(11)
	fmt.Println(f.FlipACoin())

	// Output: Tails
}

func TestFlipACoin(t *testing.T) {
	for i := 0; i < 100; i++ {
		FlipACoin()
	}
}

func BenchmarkFlipACoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FlipACoin()
	}
}

func TestRandomMapKey(t *testing.T) {
	mStr := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 100; i++ {
		key := RandomMapKey(mStr)
		if _, ok := mStr[key.(string)]; !ok {
			t.Errorf("key %s not found in map", key)
		}
	}

	mInt := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	for i := 0; i < 100; i++ {
		f := New(11)
		key := f.RandomMapKey(mInt)
		if _, ok := mInt[key.(int)]; !ok {
			t.Errorf("key %d not found in map", key)
		}
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
