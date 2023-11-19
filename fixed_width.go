package gofakeit

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

// FixedWidthOptions defines values needed for csv generation
type FixedWidthOptions struct {
	baseTemplateOptions
	HideHeader bool    `json:"hide_header" xml:"delimiter" fake:"{bool}"`
	HideFooter bool    `json:"hide_footer" xml:"hide_footer" fake:"{bool}"`
	RowCount   int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Fields     []Field `json:"fields" xml:"fields" fake:"{internal_exampleFields}"`
}

// FixedWidth generates an table of random data in fixed width format
// A nil FixedWidthOptions returns a randomly structured FixedWidth.
func FixedWidth(co *FixedWidthOptions) ([]byte, error) { return fixeWidthFunc(globalFaker, co) }

// FixedWidth generates an table of random data in fixed width format
// A nil FixedWidthOptions returns a randomly structured FixedWidth.
func (f *Faker) FixedWidth(co *FixedWidthOptions) ([]byte, error) { return fixeWidthFunc(f, co) }

// function to truncate a string
func truncateText(s string, max int) string {
	if max < len(s) && max > 0 {
		return s[:max]
	}
	return s
}

// Function to pad a string
func strPad(input string, padLength int, padString string, padType string, trim bool) string {
	var output string
	if padString == "" {
		padString = " "
	}

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		if trim {
			return truncateText(input, padLength)
		}
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "<":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case ">":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "^":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	default:
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	}

	return output
}

// Function to generate a fixed width document
func fixeWidthFunc(f *Faker, co *FixedWidthOptions) ([]byte, error) {
	if co == nil {
		// We didn't get a FixedWidthOptions, so create a new random one
		err := f.Struct(&co)
		if err != nil {
			return nil, err
		}
	}

	// Check fields
	if co.Fields == nil || len(co.Fields) <= 0 {
		return nil, errors.New("must pass fields in order to build json object(s)")
	}

	// Make sure you set a row count
	if co.RowCount <= 0 {
		return nil, errors.New("must have row count")
	}

	// Setup the template engine and custom functions
	funcMap := template.FuncMap{}
	funcMap["Pad"] = strPad
	co.SetFuncs(&funcMap)

	// Setup  variables for storing settings
	var err error
	column_spacing := make([]int, len(co.Fields))  //tmp store the spacing for each column
	column_align := make([]string, len(co.Fields)) //tmp store the align for each column
	var custom_data map[string]interface{}
	var value interface{}
	column_data := ""
	header_data := ""
	//footer_data := ""

	column_sizes := make([]int, len(co.Fields)) //tmp store the spacing for each column
	for k := range column_sizes {
		column_sizes[k] = 0
	}

	// Setup some variables for storing settings
	for i, field := range co.Fields {
		spacing := field.Params.Get("spacing")
		align := field.Params.Get("align")

		// copy parameters to the column variables or initialize them to default values
		if len(spacing) > 0 {
			column_spacing[i], err = strconv.Atoi(spacing[0])
			if err != nil {
				return nil, err
			}
		}

		column_align[i] = "<"
		if len(align) > 0 {
			column_align[i] = align[0]
		}

		// build the header data with padding and spacing

		// get the width and check if its the larger than the current column size
		if len(field.Name) > column_sizes[i] {
			column_sizes[i] = len(field.Name)
		}
		// build the header column string
		header_data += fmt.Sprintf("{{Pad `%s` (%s) `%s` (%s) true}}", field.Name, fmt.Sprintf("index .Data.column_sizes %v", i), " ", fmt.Sprintf("index .Data.column_align %v", i))

	}

	// Build the row data
	rows := make([]string, co.RowCount)
	for r := range rows {
		rows[r] = ""

		// Loop through fields and add to them to map[string]interface{}
		for k, field := range co.Fields {
			if strings.Contains(field.Function, "{{") {
				// Get function info
				value, err = templateFunc(f, field.Function, co)
				if err != nil {
					return nil, err
				}
			} else {
				value, err = runFunction(f, field, field.Function)
				if err != nil {
					return nil, err
				}
			}

			// convert the value to a string
			column_data = fmt.Sprintf("%s", value)

			// update the running total for the column if numeric
			/*if !co.HideFooter {
				updateColumTotals(co, k, column_data)
			}*/

			// get the width and check if its the larger than the current column size
			if len(column_data) > column_sizes[k] {
				column_sizes[k] = len(column_data)
			}

			// build the row data with padding and spacing
			rows[r] += fmt.Sprintf("{{Pad `%s` (%s) `%s` (%s) true}}", column_data, fmt.Sprintf("index .Data.column_sizes %v", k), " ", fmt.Sprintf("index .Data.column_align %v", k))
		}
	}

	// set the current column if the are auto size
	for k := range column_spacing {
		if column_spacing[k] > 1 {
			column_sizes[k] = column_spacing[k]
		}
	}
	// build a map of the custom data to pass to the template engine
	custom_data = map[string]interface{}{
		"column_sizes": column_sizes,
		"column_align": column_align,
	}
	co.Data = custom_data

	// build the file document
	fixed_width_result := ""
	if !co.HideHeader {
		fixed_width_result = fmt.Sprintf("%s\n", header_data)
	}

	// add the rows
	fixed_width_result = fmt.Sprintf("%s%s\n", fixed_width_result, strings.Join(rows, "\n"))

	// Render the column
	final_render, err := templateFunc(f, fixed_width_result, co)
	if err != nil {
		return nil, err
	}

	return final_render, nil

}

// Run a faker function and return the value
func runFunction(f *Faker, field Field, function string) (interface{}, error) {

	funcInfo := GetFuncLookup(function)
	if funcInfo == nil {
		// in this case the function is not a faker function so just return the value
		return function, nil
	}

	value, err := funcInfo.Generate(f.Rand, &field.Params, funcInfo)
	if err != nil {
		return nil, err
	}

	if _, ok := value.([]byte); ok {
		// If it's a slice of bytes or struct, unmarshal it into an interface
		var v interface{}
		if err := json.Unmarshal(value.([]byte), &v); err != nil {
			return nil, err
		}
		value = v
	}

	// If the value is a list of possible values, marsha it into a string
	if reflect.TypeOf(value).Kind() == reflect.Struct ||
		reflect.TypeOf(value).Kind() == reflect.Ptr ||
		reflect.TypeOf(value).Kind() == reflect.Map ||
		reflect.TypeOf(value).Kind() == reflect.Slice {
		b, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		value = string(b)
	}

	return value, nil
}

func addFixedWidthLookup() {
	AddFuncLookup("fixed_width", Info{
		Display:     "FixedWidth",
		Category:    "template",
		Description: "Generates FixedWidth document",
		Example: `
			 Name                          Email                               Cost             Account 
			 Markus Moen                   alaynawuckert@kozey.biz               8000000000948995369063 
			 Freeda Keebler                gabriellehuels@borer.io               6900000000932452944030 
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := FixedWidthOptions{}

			rowcount, err := info.GetInt(m, "rowcount")
			if err != nil {
				return nil, err
			}
			co.RowCount = rowcount

			fieldsStr, err := info.GetStringArray(m, "fields")
			if err != nil {
				return nil, err
			}

			// Check to make sure fields has length
			if len(fieldsStr) > 0 {
				co.Fields = make([]Field, len(fieldsStr))

				for i, f := range fieldsStr {
					// Unmarshal fields string into fields array
					err = json.Unmarshal([]byte(f), &co.Fields[i])
					if err != nil {
						return nil, err
					}
				}
			}

			f := &Faker{Rand: r}
			csvOut, err := fixeWidthFunc(f, &co)
			if err != nil {
				return nil, err
			}

			return csvOut, nil
		},
	})
}
