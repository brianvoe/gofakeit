package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

var noFuncRunMsg = "Could not find function to run\nRun gofakeit help or gofakeit list for available functions"

func main() {
	args := os.Args[1:]

	// Loop through args and remove any that begin with - as an indicator of flags
	cleanArgs := []string{}
	for i := 0; i < len(args); i++ {
		if !strings.HasPrefix(args[i], "-") {
			cleanArgs = append(cleanArgs, args[i])
		}
	}
	args = cleanArgs

	out := mainFunc(0, args)
	fmt.Printf("%s", out)
}

func mainFunc(seed int64, args []string) string {
	faker := gofakeit.New(seed)

	argsLen := len(args)

	// Make sure they passed first argument for function call
	if argsLen < 1 {
		return noFuncRunMsg
	}

	// Get function name
	function := args[0]

	// If function is help, give some information
	if function == "help" {
		// string builder
		var sb strings.Builder
		sb.WriteString("NAME\n")
		sb.WriteString("    gofakeit - Go fake data generator\n")
		sb.WriteString("\n")
		sb.WriteString("SYNOPSIS\n")
		sb.WriteString("    gofakeit list\n")
		sb.WriteString("    gofakeit list [category]\n")
		sb.WriteString("    gofakeit [function] [args]\n")
		sb.WriteString("\n")
		sb.WriteString("DESCRIPTION\n")
		sb.WriteString("    gofakeit is a set of functions that allow you to generate random data.\n")
		return sb.String()
	}

	// If function is list output list
	if function == "list" {
		selectedCat := ""
		if argsLen >= 2 {
			selectedCat = args[1]
		}

		selectedFunc := ""
		if argsLen >= 3 {
			selectedFunc = args[2]
		}

		return listOutput(selectedCat, selectedFunc)
	}

	// Lookup fake data method
	info := gofakeit.GetFuncLookup(function)
	if info == nil {
		return noFuncRunMsg
	}

	// Set function and params
	params := gofakeit.NewMapParams()
	paramsLen := len(info.Params)
	if paramsLen > 0 {
		for i := 0; i < argsLen; i++ {
			if i == 0 {
				continue
			}

			// Map argument to param field
			if paramsLen >= i {
				p := info.Params[i-1]
				if strings.Contains(p.Type, "[]") {
					args := strings.Split(args[i], ",")
					for _, arg := range args {
						params.Add(p.Field, arg)
					}
				} else {
					params.Add(p.Field, args[i])
				}
			}
		}
	}

	value, err := info.Generate(faker.Rand, params, info)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%v", value)
}

func listOutput(selectedCategory string, selectedFunction string) string {
	// Determine space or no space
	functionSpace := ""
	if selectedCategory == "" {
		functionSpace = "    "
	}

	stringInSlice := func(a string, list []string) bool {
		for _, b := range list {
			if b == a {
				return true
			}
		}
		return false
	}

	// Get list of categories
	categories := []string{}
	for _, l := range gofakeit.FuncLookups {
		// If selected category is set only grab of that category
		if selectedCategory != "" && selectedCategory != l.Category {
			continue
		}

		if !stringInSlice(l.Category, categories) {
			categories = append(categories, l.Category)
		}
	}

	// Sort categories
	sort.Strings(categories)

	// Put together an array or strings
	outArray := []string{}

	for i := 0; i < len(categories); i++ {
		// Write category only if category and function are not set
		if selectedCategory == "" && selectedFunction == "" {
			outArray = append(outArray, categories[i])
		}

		// Get all functions for category
		funcNames := []string{}
		for fName, l := range gofakeit.FuncLookups {
			if categories[i] == l.Category && !stringInSlice(fName, funcNames) {
				funcNames = append(funcNames, fName)
			}
		}

		// Sort function names
		sort.Strings(funcNames)

		// Output func info
		for _, fName := range funcNames {
			// If selected function is set only grab that function
			if selectedFunction != "" && selectedFunction != fName {
				continue
			}

			// Get function info
			info := gofakeit.GetFuncLookup(fName)
			outArray = append(outArray, fmt.Sprintf("%s%s - %s", functionSpace, fName, info.Description))
			for _, p := range info.Params {
				// Build string for param using string builder
				var sb strings.Builder
				sb.WriteString(functionSpace)
				sb.WriteString("    ") // Add space for param
				sb.WriteString(fmt.Sprintf("Name: %s", p.Field))
				sb.WriteString(fmt.Sprintf(" Type: %s", p.Type))
				sb.WriteString(fmt.Sprintf(" Default: %s", p.Default))
				sb.WriteString(fmt.Sprintf(" Description: %s", p.Description))
				outArray = append(outArray, sb.String())
			}
		}
	}

	// Loop through array and write to string builder
	var sb strings.Builder
	for i, s := range outArray {
		sb.WriteString(s)

		// Only add new line if not last item
		if i != len(outArray)-1 {
			sb.WriteString("\n")
		}
	}

	// Return the built string
	return sb.String()
}
