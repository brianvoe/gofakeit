package gofakeit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

// FixedWidthOptions defines values needed for csv generation
type FixedWidthOptions struct {
	RowCount int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Fields   []Field `json:"fields" xml:"fields" fake:"{internal_exampleFields}"`
}

// FixedWidth generates an table of random data in fixed width format
// A nil FixedWidthOptions returns a randomly structured FixedWidth.
func FixedWidth(co *FixedWidthOptions) (string, error) { return fixeWidthFunc(globalFaker.Rand, co) }

// FixedWidth generates an table of random data in fixed width format
// A nil FixedWidthOptions returns a randomly structured FixedWidth.
func (f *Faker) FixedWidth(co *FixedWidthOptions) (string, error) { return fixeWidthFunc(f.Rand, co) }

// Function to generate a fixed width document
func fixeWidthFunc(r *rand.Rand, co *FixedWidthOptions) (string, error) {
	// If we didn't get FixedWidthOptions, create a new random one
	if co == nil {
		co = &FixedWidthOptions{}
	}

	// Make sure you set a row count
	if co.RowCount <= 0 {
		co.RowCount = r.Intn(10) + 1
	}

	// Check fields
	if len(co.Fields) <= 0 {
		// Create random fields
		co.Fields = []Field{
			{Name: "Name", Function: "{firstname} {lastname}"},
			{Name: "Email", Function: "email"},
			{Name: "Password", Function: "password", Params: MapParams{"special": {"false"}, "space": {"false"}}},
		}
	}

	data := [][]string{}
	hasHeader := false

	// Loop through fields, generate data and add to data array
	for _, field := range co.Fields {
		// Start new row
		row := []string{}

		// Add name to first value
		if field.Name != "" {
			hasHeader = true
		}
		row = append(row, field.Name)

		// Get function
		funcInfo := GetFuncLookup(field.Function)
		var value any
		if funcInfo == nil {
			// Try to run the function through generate
			for i := 0; i < co.RowCount; i++ {
				row = append(row, generate(r, field.Function))
			}
		} else {
			// Generate function value
			var err error
			for i := 0; i < co.RowCount; i++ {
				value, err = funcInfo.Generate(r, &field.Params, funcInfo)
				if err != nil {
					value = ""
				}

				// Add value to row
				row = append(row, anyToString(value))
			}
		}

		// Add row to data
		data = append(data, row)
	}

	var result strings.Builder

	// Calculate column widths
	colWidths := make([]int, len(data))
	for i, row := range data {
		for _, value := range row {
			width := len(value) + 5
			if width > colWidths[i] {
				colWidths[i] = width
			}
		}
	}

	// Append table rows to the string, excluding the entire row if the first value is empty
	for i := 0; i < len(data[0]); i++ {
		if !hasHeader && i == 0 {
			continue // Skip the entire column if the first value is empty
		}

		var resultRow strings.Builder
		for j, row := range data {
			resultRow.WriteString(fmt.Sprintf("%-*s", colWidths[j], row[i]))
		}

		// Trim trailing spaces
		result.WriteString(strings.TrimRight(resultRow.String(), " "))

		// Only add new line if not the last row
		if i != len(data[0])-1 {
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func addFixedWidthLookup() {
	AddFuncLookup("fixed_width", Info{
		Display:     "FixedWidth",
		Category:    "template",
		Description: "Generates FixedWidth document",
		Example: `
			 Name                          Email                                 Account 
			 Markus Moen                   alaynawuckert@kozey.biz               8000000000948995369063 
			 Freeda Keebler                gabriellehuels@borer.io               6900000000932452944030 
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "10", Description: "Number of rows"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			co := FixedWidthOptions{}

			rowcount, err := info.GetInt(m, "rowcount")
			if err == nil {
				co.RowCount = rowcount
			}

			fields, _ := info.GetStringArray(m, "fields")

			// Check to make sure fields has length
			if len(fields) > 0 {
				co.Fields = make([]Field, len(fields))

				for i, f := range fields {
					// Unmarshal fields string into fields array
					err = json.Unmarshal([]byte(f), &co.Fields[i])
					if err != nil {
						return nil, err
					}
				}
			}

			out, err := fixeWidthFunc(r, &co)
			if err != nil {
				return nil, err
			}

			return out, nil
		},
	})
}
