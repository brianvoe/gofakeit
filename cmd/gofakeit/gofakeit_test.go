package main

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7/data"
)

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// add function argument
	os.Args = []string{}
	os.Args = append(os.Args, "gofakeit.go") // Fake filename to more accurately test main
	os.Args = append(os.Args, "lastname")
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	// Get output string
	outStr := string(out)

	// Check if lastname is in the data array
	isIn := false
	for _, v := range data.Person["last"] {
		if v == outStr {
			isIn = true
		}
	}

	// If not in array, fail
	if !isIn {
		t.Errorf("lastname %s not in data array", outStr)
	}
}

func TestFunctionLoop(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// add function argument and loop
	os.Args = []string{}
	os.Args = append(os.Args, "gofakeit.go") // Fake filename to more accurately test main
	os.Args = append(os.Args, "generate")
	os.Args = append(os.Args, "hello")
	os.Args = append(os.Args, "-loop=5")
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	// Get output string
	outStr := string(out)

	// Check to make sure hello is in the output 5 times
	if strings.Count(outStr, "hello") != 5 {
		t.Errorf("hello not in output 5 times")
	}
}

// runMain runs the command like it is run from the shell and returns what it printed
func runMain(args ...string) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = append([]string{"gofakeit.go"}, args...) // Fake filename to more accurately test main
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

func TestFunctionLoopWithSpace(t *testing.T) {
	// The value of the loop flag can be its own argument
	outStr := runMain("generate", "hello", "-loop", "3")
	if outStr != "hello\nhello\nhello" {
		t.Errorf("expected hello 3 times, got %q", outStr)
	}

	outStr = runMain("--loop", "2", "generate", "hello")
	if outStr != "hello\nhello" {
		t.Errorf("expected hello 2 times, got %q", outStr)
	}
}

func TestArgumentWithLoopInIt(t *testing.T) {
	// Only a flag is a flag, an argument that has -loop in it is passed to the function
	outStr := runMain("generate", "hello-loop")
	if outStr != "hello-loop" {
		t.Errorf("expected hello-loop, got %q", outStr)
	}
}

func TestNegativeNumberArguments(t *testing.T) {
	// A dash before a number is a sign, not a flag
	outStr := runMain("number", "-5", "-1", "-loop=200")
	lines := strings.Split(outStr, "\n")
	if len(lines) != 200 {
		t.Fatalf("expected 200 lines, got %d", len(lines))
	}
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil || n < -5 || n > -1 {
			t.Fatalf("expected a number from -5 to -1, got %q", line)
		}
	}
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		args    []string
		want    []string
		loop    int
		wantErr bool
	}{
		{args: []string{"name"}, want: []string{"name"}, loop: 1},
		{args: []string{"name", "-loop=3"}, want: []string{"name"}, loop: 3},
		{args: []string{"name", "-loop", "3"}, want: []string{"name"}, loop: 3},
		{args: []string{"name", "--loop=3"}, want: []string{"name"}, loop: 3},
		{args: []string{"--loop", "3", "name"}, want: []string{"name"}, loop: 3},
		{args: []string{"-loop=2", "name"}, want: []string{"name"}, loop: 2},
		{args: []string{"number", "-5", "-1"}, want: []string{"number", "-5", "-1"}, loop: 1},
		{args: []string{"float64range", "-1.5", "2.5", "-loop=4"}, want: []string{"float64range", "-1.5", "2.5"}, loop: 4},
		{args: []string{"randomstring", "a-loop,b"}, want: []string{"randomstring", "a-loop,b"}, loop: 1},
		{args: []string{"name", "-verbose"}, want: []string{"name"}, loop: 1},
		{args: []string{"name", "-loop"}, wantErr: true},
		{args: []string{"name", "-loop=x"}, wantErr: true},
		{args: []string{"name", "-loop", "x"}, wantErr: true},
	}

	for _, tt := range tests {
		got, loop, err := parseArgs(tt.args)
		if tt.wantErr {
			if err == nil {
				t.Errorf("%q: expected an error", tt.args)
			}
			continue
		}
		if err != nil {
			t.Errorf("%q: unexpected error %v", tt.args, err)
			continue
		}
		if loop != tt.loop || strings.Join(got, "\x00") != strings.Join(tt.want, "\x00") {
			t.Errorf("%q: expected %q and loop %d, got %q and loop %d", tt.args, tt.want, tt.loop, got, loop)
		}
	}
}

func TestNotEnoughArgs(t *testing.T) {
	seed := uint64(11)
	args := []string{}

	_, err := mainFunc(seed, args, 1)
	if !errors.Is(err, errNoFuncRunMsg) {
		t.Error("Expected error, got no error")
	}
}

func TestNoFunction(t *testing.T) {
	seed := uint64(11)
	args := []string{"notafunction"}

	_, err := mainFunc(seed, args, 1)
	if !errors.Is(err, errNoFuncRunMsg) {
		t.Error("Expected error, got no error")
	}
}

func TestFunctionSimple(t *testing.T) {
	seed := uint64(11)
	args := []string{"firstname"}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if lastname is in the data array
	isIn := false
	for _, v := range data.Person["first"] {
		if v == outStr {
			isIn = true
		}
	}

	// If not in array, fail
	if !isIn {
		t.Errorf("firstname %s not in data array", outStr)
	}
}

func TestFunctionWithParams(t *testing.T) {
	strs := []string{"hello", "world", "whats", "up"}

	seed := uint64(11)
	args := []string{"shufflestrings", strings.Join(strs, ",")}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if outStr != "[up world whats hello]" {
		t.Errorf("shufflestrings output is not in the right order. Got: %s", outStr)
	}
}

func TestHelp(t *testing.T) {
	seed := uint64(11)
	args := []string{"help"}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Make sure outStr contains NAME, SYNOPSIS, DESCRIPTION
	if !strings.Contains(outStr, "NAME") {
		t.Errorf("help output does not contain NAME")
	}
	if !strings.Contains(outStr, "SYNOPSIS") {
		t.Errorf("help output does not contain SYNOPSIS")
	}
	if !strings.Contains(outStr, "DESCRIPTION") {
		t.Errorf("help output does not contain DESCRIPTION")
	}
}

func TestList(t *testing.T) {
	seed := uint64(11)
	args := []string{"list"}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Make sure outStr contains person, word, address, string, and number
	if !strings.Contains(outStr, "person") {
		t.Errorf("list output does not contain person")
	}
	if !strings.Contains(outStr, "word") {
		t.Errorf("list output does not contain word")
	}
	if !strings.Contains(outStr, "address") {
		t.Errorf("list output does not contain address")
	}
	if !strings.Contains(outStr, "string") {
		t.Errorf("list output does not contain string")
	}
	if !strings.Contains(outStr, "number") {
		t.Errorf("list output does not contain number")
	}
}

func TestListCategory(t *testing.T) {
	seed := uint64(11)
	args := []string{"list", "person"}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Make sure outStr contains email, firstname, lastname, and gender
	if !strings.Contains(outStr, "email") {
		t.Errorf("list category output does not contain email")
	}
	if !strings.Contains(outStr, "firstname") {
		t.Errorf("list category output does not contain firstname")
	}
	if !strings.Contains(outStr, "lastname") {
		t.Errorf("list category output does not contain lastname")
	}
	if !strings.Contains(outStr, "gender") {
		t.Errorf("list category output does not contain gender")
	}
}

func TestListCategoryFunction(t *testing.T) {
	seed := uint64(11)
	args := []string{"list", "word", "noun"}

	outStr, err := mainFunc(seed, args, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	outStr = strings.ToLower(outStr)

	// Make sure outStr contains random noun
	if !strings.Contains(outStr, "noun") {
		t.Errorf("list category function output does not contain random noun")
	}
}
