package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type funcDecl = ast.FuncDecl

// ExportedFuncs returns exported functions from the path which must be a directory.
// Functions, additionally to being exported, must return a string.
// TODO: add support for functions taking one or more params.
func ExportedFuncs(path string) ([]*funcDecl, error) {
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, path, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, 0)
	if err != nil {
		return nil, err
	}
	var funcs []*funcDecl
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				fn, isFn := decl.(*ast.FuncDecl)
				if !isFn || !isValidFunc(fn) {
					continue
				}
				funcs = append(funcs, fn)
			}
		}
	}
	return funcs, nil
}

// isValidFunc validates whether a function is exported,
// takes no parameters and returns a single string value.
func isValidFunc(fn *funcDecl) bool {
	// Only exporter functions can be executed.
	if !fn.Name.IsExported() {
		return false
	}
	// Accept only functions that take no params.
	if fn.Type.Params != nil && fn.Type.Params.NumFields() > 0 {
		return false
	}
	// Accept only functions that return a single string value.
	return funcReturnsString(fn)
}

// funcReturnsString returns true if a function returns
// a single string value as a result.
func funcReturnsString(fn *funcDecl) bool {
	// Accept only functions that return a single result.
	if fn.Type.Results == nil || fn.Type.Results.NumFields() != 1 {
		return false
	}
	// The returned result must be a string.
	ident, ok := fn.Type.Results.List[0].Type.(*ast.Ident)
	return ok && ident.Name == "string"
}
