package gofakeit

import (
	"fmt"
	"testing"
)

// Test taking in two random int values and making sure the output is within the range
func TestRandIntRange(t *testing.T) {
	// Create a set of structs to test
	type testStruct struct {
		min, max int
	}

	// Create a set of test values
	tests := []testStruct{
		{0, 0},
		{1000, -1000},
		{minInt, maxInt},
		{minInt, minInt + 100}, // Test min
		{maxInt - 100, maxInt}, // Test max
		{maxInt - 20000, maxInt - 10000},
		{minInt + 10000, maxInt - 10000},
	}

	// Add 10000 random values to the test set
	for i := 0; i < 5000; i++ {
		tests = append(tests, testStruct{
			min: randIntRange(GlobalFaker, 0, maxInt),
			max: randIntRange(GlobalFaker, 0, maxInt),
		})
	}
	for i := 0; i < 5000; i++ {
		tests = append(tests, testStruct{
			min: randIntRange(GlobalFaker, minInt, 0),
			max: randIntRange(GlobalFaker, 0, maxInt),
		})
	}

	// Loop through the tests
	for _, test := range tests {
		// Get the result
		result := randIntRange(GlobalFaker, test.min, test.max)

		// Check the result
		if test.min > test.max {
			// Test if min is greater than max by reversing the check
			if result > test.min && result < test.max {
				t.Fatalf("The result should be between %d and %d. Got: %d", test.min, test.max, result)
			}
		} else if result < test.min || result > test.max {
			t.Fatalf("The result should be between %d and %d. Got: %d", test.min, test.max, result)
		}
	}
}

// Test taking in two random uint values and making sure the output is within the range
func TestRandUintRange(t *testing.T) {
	// Create a set of structs to test
	type testStruct struct {
		min, max uint
	}

	// Create a set of test values
	tests := []testStruct{
		{0, 0},
		{100000, 100},
		{minUint, maxUint},
		{minUint + 10000, maxUint - 10000},
	}

	// Add 10000 random values to the test set
	for i := 0; i < 5000; i++ {
		tests = append(tests, testStruct{
			min: randUintRange(GlobalFaker, 0, maxUint),
			max: randUintRange(GlobalFaker, 0, maxUint),
		})
	}
	for i := 0; i < 5000; i++ {
		tests = append(tests, testStruct{
			min: randUintRange(GlobalFaker, 0, maxUint/2),
			max: randUintRange(GlobalFaker, maxUint/2, maxUint),
		})
	}

	// Loop through the tests
	for _, test := range tests {
		// Get the result
		result := randUintRange(GlobalFaker, test.min, test.max)

		// Check the result
		if test.min > test.max {
			// Test if min is greater than max by reversing the check
			if result > test.min && result < test.max {
				t.Fatalf("The result should be between %d and %d. Got: %d", test.min, test.max, result)
			}
		} else if result < test.min || result > test.max {
			t.Fatalf("The result should be between %d and %d. Got: %d", test.min, test.max, result)
		}
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if getRandValue(GlobalFaker, test) != "" {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestReplaceWithNumbers(t *testing.T) {
	if replaceWithNumbers(GlobalFaker, "") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func BenchmarkReplaceWithNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Seed(11)
		replaceWithNumbers(GlobalFaker, "###☺#☻##☹##")
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	for _, test := range []struct{ in string }{
		{"#界#世#"},
		{"☺#☻☹#"},
		{"\x80#¼#語"},
	} {
		Seed(11)
		got := replaceWithNumbers(GlobalFaker, test.in)
		if got == test.in {
			t.Errorf("got %q, want something different", got)
		}

		// Check that # were replaced with numbers
		for _, r := range got {
			if r == '#' {
				t.Errorf("got %q, want something different", got)
			}
		}
	}
}

func TestReplaceWithLetters(t *testing.T) {
	if replaceWithLetters(GlobalFaker, "") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithHexLetters(t *testing.T) {
	if replaceWithHexLetters(GlobalFaker, "") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func TestToFixed(t *testing.T) {
	floats := [][]float64{
		{123.1234567489, 123.123456},
		{987.987654321, 987.987654},
	}

	for _, f := range floats {
		if toFixed(f[0], 6) != f[1] {
			t.Fatalf("%g did not equal %g. Got: %g", f[0], f[1], toFixed(f[0], 6))
		}
	}
}

func TestEqualSlice(t *testing.T) {
	// String Array
	if equalSliceString([]string{"a", "b"}, []string{"a"}) {
		t.Fatalf("Should have returned false because the string array are not the same")
	}
	if equalSliceString([]string{"a", "b"}, []string{"c", "d"}) {
		t.Fatalf("Should have returned false because the string array are not the same")
	}
	if !equalSliceString([]string{"a", "b"}, []string{"a", "b"}) {
		t.Fatalf("Should have returned true because the string array are the same")
	}

	// Int Array
	if equalSliceInt([]int{1, 2}, []int{1}) {
		t.Fatalf("Should have returned false because the int array are not the same")
	}
	if equalSliceInt([]int{1, 2}, []int{3, 4}) {
		t.Fatalf("Should have returned false because the int array are not the same")
	}
	if !equalSliceInt([]int{1, 2}, []int{1, 2}) {
		t.Fatalf("Should have returned true because the int array are the same")
	}

	// Interface Array
	if equalSliceInterface([]any{1, "b"}, []any{1}) {
		t.Fatalf("Should have returned false because the interface array are not the same")
	}
	if equalSliceInterface([]any{1, "b"}, []any{3, "d"}) {
		t.Fatalf("Should have returned false because the interface array are not the same")
	}
	if !equalSliceInterface([]any{1, "b", []int{1, 2}, []string{"a", "b"}}, []any{1, "b", []int{1, 2}, []string{"a", "b"}}) {
		t.Fatalf("Should have returned true because the ints array are the same")
	}
}

func TestStringInSlice(t *testing.T) {
	if stringInSlice("c", []string{"a", "b"}) {
		t.Fatalf("Should have returned true because the string is in the slice")
	}
	if !stringInSlice("a", []string{"a", "b", "c"}) {
		t.Fatalf("Should have returned true because the string is in the slice")
	}
}

func TestAnyToString(t *testing.T) {
	tests := []struct {
		input any
		want  string
	}{
		{input: 42, want: "42"},
		{input: "Hello, Go!", want: "Hello, Go!"},
		{input: 3.14, want: "3.14"},
		{input: true, want: "true"},
		{input: struct{ Name string }{Name: "John"}, want: `{"Name":"John"}`},
		{input: []int{1, 2, 3}, want: "[1,2,3]"},
		{input: map[string]int{"a": 1, "b": 2}, want: `{"a":1,"b":2}`},
		{input: []byte("hello world"), want: "hello world"},
		{input: nil, want: ""},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %v", test.input), func(t *testing.T) {
			got := anyToString(test.input)

			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}

	// Additional tests for edge cases
	t.Run("Edge case: empty byte slice", func(t *testing.T) {
		emptyByteSlice := []byte{}
		got := anyToString(emptyByteSlice)
		want := ""

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestAnyToStringEdgeCases(t *testing.T) {
	// Additional tests for edge cases

	// Test with a struct containing an unexported field
	type unexportedFieldStruct struct {
		unexportedField int
		ExportedField   string
	}
	unexportedStruct := unexportedFieldStruct{unexportedField: 42, ExportedField: "Hello"}
	want := `{"ExportedField":"Hello"}`
	got := anyToString(unexportedStruct)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	// Test with a struct containing a nil interface field
	type nilInterfaceFieldStruct struct {
		Data any
	}
	nilInterfaceStruct := nilInterfaceFieldStruct{}
	want = `{"Data":null}`
	got = anyToString(nilInterfaceStruct)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTitle(t *testing.T) {
	test := map[string]string{
		"":                     "",
		"i have a best friend": "I Have A Best Friend",
		"this is a test":       "This Is A Test",
		"I am 36 years old":    "I Am 36 Years Old",
		"whats_up":             "Whats_up",
	}

	for in, out := range test {
		if title(in) != out {
			t.Fatalf("%s did not equal %s. Got: %s", in, out, title(in))
		}
	}
}

func TestFuncLookupSplit(t *testing.T) {
	tests := map[string][]string{
		"":                  {},
		"a":                 {"a"},
		"a,b,c":             {"a", "b", "c"},
		"a, b, c":           {"a", "b", "c"},
		"[a,b,c]":           {"[a,b,c]"},
		"a,[1,2,3],b":       {"a", "[1,2,3]", "b"},
		"[1,2,3],a,[1,2,3]": {"[1,2,3]", "a", "[1,2,3]"},
	}

	for input, expected := range tests {
		values, err := funcLookupSplit(input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(values) != len(expected) {
			t.Fatalf("%s was not %s", values, expected)
		}
		for i := 0; i < len(values); i++ {
			if values[i] != expected[i] {
				t.Fatalf("expected %s got %s", expected[i], values[i])
			}
		}
	}
}
