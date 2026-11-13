package gofakeit

import (
	"fmt"
	"reflect"
	"testing"
)

// The same seed has to give the same result. Go reads maps in a random order, so anything
// that picks from a map or writes one out has to put the keys in order first.

func TestRandomMapKeySeed(t *testing.T) {
	mStr := make(map[string]int)
	mInt := make(map[int]string)
	mUint := make(map[uint8]bool)
	mFloat := make(map[float64]bool)
	mArray := make(map[[2]int]bool)
	mAny := make(map[any]bool)
	for i := 0; i < 50; i++ {
		mStr[fmt.Sprintf("key%d", i)] = i
		mInt[i] = fmt.Sprintf("value%d", i)
		mUint[uint8(i)] = true
		mFloat[float64(i)+0.5] = true
		mArray[[2]int{i, i + 1}] = true
		mAny[i] = true
		mAny[fmt.Sprint(i)] = true
		mAny[float64(i)] = true
	}

	for i := 0; i < 20; i++ {
		for _, m := range []any{mStr, mInt, mUint, mFloat, mArray, mAny} {
			if a, b := New(11).RandomMapKey(m), New(11).RandomMapKey(m); a != b {
				t.Fatalf("%T: the same seed gave %v and %v", m, a, b)
			}
		}
	}
}

func TestLessMapKey(t *testing.T) {
	tests := []struct {
		a, b any
		want bool
	}{
		{"a", "b", true},
		{"b", "a", false},
		{"a", "a", false},
		{2, 10, true},
		{10, 2, false},
		{int8(-5), int8(3), true},
		{uint(2), uint(10), true},
		{1.5, 2.5, true},
		{2.5, 1.5, false},
		{1, "1", true}, // different types go by the name of the type, int before string
		{"1", 1, false},
		{[2]int{1, 2}, [2]int{1, 3}, true},
		{false, true, true},
	}
	for _, test := range tests {
		if got := lessMapKey(reflect.ValueOf(test.a), reflect.ValueOf(test.b)); got != test.want {
			t.Errorf("lessMapKey(%#v, %#v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}

	// The keys of a map[any]T come out of MapKeys wrapped in an interface
	keys := reflect.ValueOf(map[any]bool{2: true, 10: true}).MapKeys()
	if lessMapKey(keys[0], keys[1]) == lessMapKey(keys[1], keys[0]) {
		t.Error("wrapped keys 2 and 10 are not ordered")
	}
}

func TestGetRandomSimpleFuncSeed(t *testing.T) {
	for i := 0; i < 20; i++ {
		a, _ := GetRandomSimpleFunc(New(11))
		b, _ := GetRandomSimpleFunc(New(11))
		if a != b {
			t.Fatalf("the same seed gave the functions %s and %s", a, b)
		}
	}
}

// A nil option struct is filled with random fields, and the random fields are picked
// from all the functions without parameters
func TestRandomOptionsSeed(t *testing.T) {
	for i := 0; i < 20; i++ {
		var j1, j2 JSONOptions
		New(11).Struct(&j1)
		New(11).Struct(&j2)
		if !reflect.DeepEqual(j1, j2) {
			t.Fatalf("JSONOptions: the same seed gave %+v and %+v", j1, j2)
		}

		var c1, c2 CSVOptions
		New(11).Struct(&c1)
		New(11).Struct(&c2)
		if !reflect.DeepEqual(c1, c2) {
			t.Fatalf("CSVOptions: the same seed gave %+v and %+v", c1, c2)
		}

		var x1, x2 XMLOptions
		New(11).Struct(&x1)
		New(11).Struct(&x2)
		if !reflect.DeepEqual(x1, x2) {
			t.Fatalf("XMLOptions: the same seed gave %+v and %+v", x1, x2)
		}
	}
}

// An address or a person is written as nested elements
func TestXMLNestedSeed(t *testing.T) {
	gen := func() string {
		value, err := New(11).XML(&XMLOptions{
			Type:          "single",
			RootElement:   "xml",
			RecordElement: "record",
			RowCount:      1,
			Fields: []Field{
				{Name: "home", Function: "address"},
				{Name: "work", Function: "job"},
				{Name: "person", Function: "person"},
			},
		})
		if err != nil {
			t.Fatal(err.Error())
		}
		return string(value)
	}

	first := gen()
	for i := 0; i < 20; i++ {
		if got := gen(); got != first {
			t.Fatalf("the same seed gave different XML:\n%s\n%s", first, got)
		}
	}
}
