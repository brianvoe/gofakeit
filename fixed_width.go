package gofakeit

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

// FixedWidthOptions defines values needed for csv generation
type FixedWidthOptions struct {
	baseTemplateOptions
	HideHeader    bool      `json:"hide_header" xml:"delimiter" fake:"{bool}"`
	HideFooter    bool      `json:"hide_footer" xml:"hide_footer" fake:"{bool}"`
	RowCount      int       `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Fields        []Field   `json:"fields" xml:"fields" fake:"{internal_exampleFields}"`
	runningTotal  []float64 //keep track of the running total
	currentColumn int       //keep track of the current column
}

// get the running total for the current column
func (f FixedWidthOptions) GetTotal() interface{} {
	if f.runningTotal[f.currentColumn] == -999999999999999 {
		return ""
	}
	return fmt.Sprintf("%.2f", f.runningTotal[f.currentColumn])
}

// FixedWidth generates an object or an array of objects in json format
// A nil FixedWidthOptions returns a randomly structured FixedWidth.
func FixedWidth(co *FixedWidthOptions) ([]byte, error) { return fixeWidthFunc(globalFaker, co) }

// FixedWidth generates an object or an array of objects in json format
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

	// Setup  variables
	var err error
	column_spacing := make([]int, len(co.Fields))       //tmp store the spacing for each column
	column_align := make([]string, len(co.Fields))      //tmp store the spacing for each column
	column_header_pad := make([]string, len(co.Fields)) //tmp store the spacing for each column
	column_row_pad := make([]string, len(co.Fields))    //tmp store the spacing for each column
	column_footer_pad := make([]string, len(co.Fields)) //tmp store the spacing for each column
	column_footer := make([]string, len(co.Fields))     //save the largest column size
	var custom_data map[string]interface{}              //custom data for the template
	var value interface{}
	column_data := ""
	header_data := ""
	footer_data := ""

	column_sizes := make([]int, len(co.Fields)) //tmp store the spacing for each column
	for k := range column_sizes {
		column_sizes[k] = 0
	}

	// Setup the running total
	co.runningTotal = make([]float64, len(co.Fields))
	for k := range co.runningTotal {
		co.runningTotal[k] = -999999999999999
	}

	// Setup some variables for storing settings
	for i, field := range co.Fields {
		co.currentColumn = i
		spacing := field.Params.Get("spacing")
		align := field.Params.Get("align")
		footer := field.Params.Get("footer")
		header_pad := field.Params.Get("header_pad")
		row_pad := field.Params.Get("row_pad")
		footer_pad := field.Params.Get("footer_pad")

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

		column_footer[i] = ""
		if len(footer) > 0 {
			column_footer[i] = footer[0]
		}

		column_header_pad[i] = " "
		if len(header_pad) > 0 {
			column_header_pad[i] = header_pad[0]
		}

		column_row_pad[i] = " "
		if len(row_pad) > 0 {
			column_row_pad[i] = row_pad[0]
		}

		column_footer_pad[i] = " "
		if len(footer_pad) > 0 {
			column_footer_pad[i] = footer_pad[0]
		}

		//build the header data with padding and spacing
		if !co.HideHeader {
			//get the width and check if its the larger than the current column size
			if len(field.Name) > column_sizes[i] {
				column_sizes[i] = len(field.Name)
			}
			header_data += fmt.Sprintf("{{Pad `%s` (%s)  (%s) (%s) true}}", field.Name, fmt.Sprintf("index .Data.column_sizes %v", i), fmt.Sprintf("index .Data.column_header_pad %v", i), fmt.Sprintf("index .Data.column_align %v", i))
		}
	}

	// Build the row data
	rows := make([]string, co.RowCount)
	for r := range rows {
		rows[r] = ""

		// Loop through fields and add to them to map[string]interface{}
		for k, field := range co.Fields {
			co.currentColumn = k
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
			if !co.HideFooter {
				updateColumTotals(co, k, column_data)
			}

			// get the width and check if its the larger than the current column size
			if len(column_data) > column_sizes[k] {
				column_sizes[k] = len(column_data)
			}

			// build the row data with padding and spacing
			rows[r] += fmt.Sprintf("{{Pad `%s` (%s)  (%s) (%s) true}}", column_data, fmt.Sprintf("index .Data.column_sizes %v", k), fmt.Sprintf("index .Data.column_row_pad %v", k), fmt.Sprintf("index .Data.column_align %v", k))
		}
	}

	// build the footer data
	if !co.HideFooter {
		for i, field := range co.Fields {
			co.currentColumn = i
			// parse the footer function
			if strings.Contains(column_footer[i], "{{") {
				// Get function info
				value, err = templateFunc(f, column_footer[i], co)
				if err != nil {
					return nil, err
				}
			} else {
				value, err = runFunction(f, field, column_footer[i])
				if err != nil {
					return nil, err
				}
			}

			// convert the value to a string
			column_data = fmt.Sprintf("%s", value)

			// build the footer data with padding and spacing
			if len(column_data) > column_sizes[i] {
				column_sizes[i] = len(column_data)
			}

			// build the footer column string
			footer_data += fmt.Sprintf("{{Pad `%s` (%s)  (%s) (%s) true}}", column_data, fmt.Sprintf("index .Data.column_sizes %v", i), fmt.Sprintf("index .Data.column_footer_pad %v", i), fmt.Sprintf("index .Data.column_align %v", i))
		}
	}

	// set the current column if the are auto size
	for k := range column_spacing {
		if column_spacing[k] > 1 {
			column_sizes[k] = column_spacing[k]
		}
	}

	custom_data = map[string]interface{}{
		"column_sizes":      column_sizes,
		"column_header_pad": column_header_pad,
		"column_row_pad":    column_row_pad,
		"column_footer_pad": column_footer_pad,
		"column_align":      column_align,
	}
	co.Data = custom_data

	// build the file document
	fixed_width_result := ""
	if !co.HideHeader {
		fixed_width_result = fmt.Sprintf("%s\n", header_data)
	}

	// add the rows
	fixed_width_result = fmt.Sprintf("%s%s\n", fixed_width_result, strings.Join(rows, "\n"))

	// add the footer
	if !co.HideFooter {
		fixed_width_result = fmt.Sprintf("%s%s", fixed_width_result, footer_data)
	}

	// Render the column
	final_render, err := templateFunc(f, fixed_width_result, co)
	if err != nil {
		return nil, err
	}

	return final_render, nil

}

// function to update the running total for the column
func updateColumTotals(co *FixedWidthOptions, column int, value interface{}) {
	column_data := fmt.Sprintf("%s", value)
	if regexp.MustCompile(`\d`).MatchString(column_data) {
		value, err := strconv.ParseFloat(column_data, 32)
		if err == nil {
			//the data is numeric so check if value -99999... the set it to zero
			if co.runningTotal[column] == -999999999999999 {
				co.runningTotal[column] = 0
			}
			co.runningTotal[column] += value
		}
	}
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
			"Name==========================Email===============================Cost=============Account"
			"Markus Moen                   alaynawuckert@kozey.biz               8000000000948995369063"
			"Freeda Keebler                gabriellehuels@borer.io               6900000000932452944030"
			"****************************************************************218.00********************"
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function"},
			{Field: "hide_header", Display: "HideHeader", Type: "bool", Default: "false", Description: "Hide the header"},
			{Field: "hide_footer", Display: "HideFooter", Type: "bool", Default: "false", Description: "Hide the footer"},
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
