package gofakeit

import (
	"reflect"
	"sort"
	"testing"

	"github.com/brianvoe/gofakeit/data"
)

func TestRandIntRange(t *testing.T) {
	if randIntRange(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestGetRandValueFail(t *testing.T) {
	if getRandValue([]string{"not", "found"}) != "" {
		t.Error("You should have gotten no value back")
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	if getRandIntValue([]string{"not", "found"}) != 0 {
		t.Error("You should have gotten no value back")
	}
}

func TestRandFloatRangeSame(t *testing.T) {
	if randFloatRange(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestReplaceWithNumbers(t *testing.T) {
	if "" != replaceWithNumbers("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithLetters(t *testing.T) {
	if "" != replaceWithLetters("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestCatagories(t *testing.T) {
	var got, expected []string
	for k := range Catagories() {
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
