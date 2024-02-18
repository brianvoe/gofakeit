package data

import (
	"fmt"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	// List data
	list := List()

	// Should get a filled list of map[string][]string
	if len(list) != len(Data) {
		t.Errorf("Expected %d items, got %d", len(Data), len(list))
	}
}

func ExampleGet() {
	// Get data
	data := Get("person")

	// Print the first name
	fmt.Println(data["prefix"])

	// Output:
	// [Mr. Mrs. Ms. Miss Dr.]
}

func ExampleGet_sub() {
	// Get data
	data := GetSubData("person", "prefix")

	// Print the first name
	fmt.Println(data[0])

	// Output:
	// Mr.
}

func TestGetData(t *testing.T) {
	// Get data
	data := Get("things")

	// Should get empty map
	if len(data) != 0 {
		t.Errorf("Expected 0 things, got %d", len(data))
	}
}

func TestGetSub(t *testing.T) {
	// Get data
	data := GetSubData("things", "cake")

	// Should get empty map
	if len(data) != 0 {
		t.Errorf("Expected 0 things, got %d", len(data))
	}
}

func ExampleSet() {
	// Add data
	Set("desserts", map[string][]string{
		"cake":      {"chocolate", "vanilla"},
		"pie":       {"apple", "pecan"},
		"ice cream": {"strawberry", "vanilla"},
	})
}

func ExampleSetSub() {
	// Add data
	SetSub("desserts", "cake", []string{"chocolate", "vanilla"})
}

func TestAddReplaceRemove(t *testing.T) {
	// Add data
	Set("desserts", map[string][]string{
		"cake":      {"chocolate", "vanilla"},
		"pie":       {"apple", "pecan"},
		"ice cream": {"strawberry", "vanilla"},
	})

	if len(Data["desserts"]) != 3 {
		t.Errorf("Expected 3 desserts, got %d", len(Data["desserts"]))
	}
	if len(Data["desserts"]["cake"]) != 2 {
		t.Errorf("Expected 2 cakes, got %d", len(Data["desserts"]["cake"]))
	}

	// Replace data
	Set("desserts", map[string][]string{
		"cake":      {"carrot", "banana"},
		"pie":       {"cherry", "peach"},
		"ice cream": {"caramel", "rainbow"},
	})

	if len(Data["desserts"]) != 3 {
		t.Errorf("Expected 3 desserts, got %d", len(Data["desserts"]))
	}
	if len(Data["desserts"]["cake"]) != 2 {
		t.Errorf("Expected 2 cakes, got %d", len(Data["desserts"]["cake"]))
	}

	// Remove data
	Remove("desserts")

	if len(Data["desserts"]) != 0 {
		t.Errorf("Expected 0 desserts, got %d", len(Data["desserts"]))
	}
}

func TestAddReplaceRemoveSub(t *testing.T) {
	// Add data
	SetSub("desserts", "cake", []string{"chocolate", "vanilla"})

	if len(Data["desserts"]["cake"]) != 2 {
		t.Errorf("Expected 2 cakes, got %d", len(Data["desserts"]["cake"]))
	}

	// Replace data
	SetSub("desserts", "cake", []string{"carrot", "banana"})

	if len(Data["desserts"]["cake"]) != 2 {
		t.Errorf("Expected 2 cakes, got %d", len(Data["desserts"]["cake"]))
	}

	// Remove data
	RemoveSub("desserts", "cake")

	if len(Data["desserts"]["cake"]) != 0 {
		t.Errorf("Expected 0 cakes, got %d", len(Data["desserts"]["cake"]))
	}
}

func TestTrailingSpaces(t *testing.T) {
	assertSpaces := func(selector, str string) {
		t.Helper()
		if str != strings.Trim(str, " ") {
			t.Errorf("In %s dataset: %q has leading or trailing space", selector, str)
		}
	}

	for dataKey, data := range Data {
		for key, values := range data {
			assertSpaces(dataKey, key)
			for _, v := range values {
				assertSpaces(dataKey, v)
			}
		}
	}
}
