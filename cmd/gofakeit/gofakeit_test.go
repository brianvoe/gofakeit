package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6/data"
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
	out, _ := ioutil.ReadAll(r)
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
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	// Get output string
	outStr := string(out)

	// Check to make sure hello is in the output 5 times
	if strings.Count(outStr, "hello") != 5 {
		t.Errorf("hello not in output 5 times")
	}
}

func TestNotEnoughArgs(t *testing.T) {
	seed := int64(11)
	args := []string{}

	outStr := mainFunc(seed, args, 1)

	if outStr != noFuncRunMsg {
		t.Errorf("Expected %s, got %s", noFuncRunMsg, outStr)
	}
}

func TestNoFunction(t *testing.T) {
	seed := int64(11)
	args := []string{"notafunction"}

	outStr := mainFunc(seed, args, 1)

	if outStr != noFuncRunMsg {
		t.Errorf("Expected %s, got %s", noFuncRunMsg, outStr)
	}
}

func TestFunctionSimple(t *testing.T) {
	seed := int64(11)
	args := []string{"firstname"}

	outStr := mainFunc(seed, args, 1)

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

	seed := int64(11)
	args := []string{"shufflestrings", strings.Join(strs, ",")}

	outStr := mainFunc(seed, args, 1)
	if outStr != "[up world whats hello]" {
		t.Errorf("shufflestrings output is not in the right order. Got: %s", outStr)
	}
}

func TestHelp(t *testing.T) {
	seed := int64(11)
	args := []string{"help"}

	outStr := mainFunc(seed, args, 1)

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
	seed := int64(11)
	args := []string{"list"}

	outStr := mainFunc(seed, args, 1)

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
	seed := int64(11)
	args := []string{"list", "person"}

	outStr := mainFunc(seed, args, 1)

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
	seed := int64(11)
	args := []string{"list", "word", "noun"}

	outStr := mainFunc(seed, args, 1)
	outStr = strings.ToLower(outStr)

	// Make sure outStr contains random noun
	if !strings.Contains(outStr, "random noun") {
		t.Errorf("list category function output does not contain random noun")
	}
}
