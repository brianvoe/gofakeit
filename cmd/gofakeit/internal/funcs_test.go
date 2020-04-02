package internal

import (
	"testing"
)

func TestExportedFuncs(t *testing.T) {
	const projectRoot = "./../../.."
	funcs, err := ExportedFuncs(projectRoot)
	if err != nil {
		t.Fatalf("failed to call ExportedFuncs on the project root: %v", err)
	}
	// BeerName function should be found because it takes no params and returns a single string value.
	if !containsFuncByName(funcs, "BeerName") {
		t.Errorf("failed to find expected 'BeerName' function")
	}
	// RGBColor takes no params but returns a slice of ints and should not be found.
	if containsFuncByName(funcs, "RGBColor") {
		t.Errorf("found 'RGBColor' function which does not meet criteria")
	}
}

func containsFuncByName(funcs []*funcDecl, name string) bool {
	for _, fn := range funcs {
		if fn.Name.Name == name {
			return true
		}
	}
	return false
}
