package gofakeit

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

// getValue returns the data interface
type FixedWidthOptions struct {
	baseTemplateOptions
	Header        []string    `json:"header" xml:"header" fake:"{{AutoHeader}}" fakesize:"8"`
	Row           []string    `json:"row" xml:"row" fake:"{randomstring:[{{FirstName}} {{LastName}},{{Email}},{{Number 1 100}},{{AchAccount}},{{Company}},{{Price 10 1000}},{{CurrencyShort}}]}" fakesize:"8"`
	Footer        []string    `json:"footer" xml:"footer" fake:"{{.GetTotal}}" fakesize:"8"`
	Spacing       []int       `json:"spacing" xml:"spacing" fake:"-1" fakesize:"8"`               //Required Column Spacing -1 = auto
	Align         []string    `json:"align" xml:"align" fake:"{randomstring:[<,>]}" fakesize:"8"` //Alignment of the columns < left, > right or ^ center
	RowPad        []string    `json:"row_pad" xml:"row_pad" fake:" "  fakesize:"8"`               //What character to use for padding the footer column
	HeaderPad     []string    `json:"header_pad" xml:"header_pad" fake:" "  fakesize:"8"`         //What character to use for padding the footer column
	FooterPad     []string    `json:"footer_pad" xml:"footer_pad" fake:" "  fakesize:"8"`         //What character to use for padding the footer column
	Count         int         `json:"count" xml:"count" fake:"{number:5,30}" `                    //Number of Rows to generate
	runningTotal  []float64   //keep track of the running total
	currentColumn int         //keep track of the current column
	Data          interface{} //Data to be used in template
}

// getValue returns the data interface
func (f FixedWidthOptions) GetData() interface{} {
	return f.Data
}

// get the running total for the current column
func (f FixedWidthOptions) GetTotal() interface{} {
	if f.runningTotal[f.currentColumn] == -999999999999999 {
		return ""
	}
	return fmt.Sprintf("%.2f", f.runningTotal[f.currentColumn])
}

// get the spacing for the current column
func (f FixedWidthOptions) GetSpacing(col int) int {
	return f.Spacing[col]
}

// Template generates an document based on the the supplied template
// A nil TemplateOptions returns a document.
func FixedWidth(co *FixedWidthOptions) ([]byte, error) {
	return fixedWidth(globalFaker, co)
}

// Template generates an document or an array of objects in json format
// A nil TemplateOptions returns a randomly structured CSV.
func (f *Faker) FixedWidth(co *FixedWidthOptions) ([]byte, error) {
	return fixedWidth(f, co)
}

// function to help setup the parameters for the template engine
func defaultValues(name string, Param []string, defaultValue string, size int) ([]string, error) {

	// see if we have alignment for each column if not set it to left
	if len(Param) == 0 {
		Param = make([]string, size)
		for i := range Param {
			//test if Param is empty and does not contain < > ^ then set it to the default value
			if Param[i] != "<" && Param[i] != ">" && Param[i] != "^" {
				Param[i] = defaultValue
			}
		}
	}
	// check if we have alignment that it matches the number of columns
	if len(Param) != size {
		Param = makeStringSliceRequiredSize(Param, defaultValue, size)
	}

	return Param, nil
}

// function  to make sure the sting slice is the correct size
func makeStringSliceRequiredSize(Param []string, defaultValue string, size int) []string {
	// check if we have alignment that it matches the number of columns
	if len(Param) != size {
		if len(Param) > size {
			Param = Param[:size]
		} else if len(Param) == 0 {
			Param = make([]string, size)
			for i := range Param {
				Param[i] = defaultValue
			}
		} else if len(Param) < size {
			//make the Param the same size as the number of columns
			Param = append(Param, make([]string, size-len(Param))...)
			for i := range Param {
				Param[i] = defaultValue
			}
		}
	}
	return Param
}

// function  to make sure the Int slice is the correct size
func makeIntSliceRequiredSize(Param []int, defaultValue int, size int) []int {
	// check if we have alignment that it matches the number of columns
	if len(Param) != size {
		if len(Param) > size {
			Param = Param[:size]
		} else if len(Param) == 0 {
			Param = make([]int, size)
			for i := range Param {
				Param[i] = defaultValue
			}
		} else if len(Param) < size {
			//make the Param the same size as the number of columns
			Param = append(Param, make([]int, size-len(Param))...)
			for i := range Param {
				Param[i] = defaultValue
			}
		}
	}
	return Param
}

// function to validate the options and set defaults where needed
func validateInputs(co *FixedWidthOptions) error {
	var err error

	//Set a max rows
	if co.Count > 100000 {
		co.Count = 100000
	}
	if co.Count < 0 {
		co.Count = 1
	}

	if len(co.Row) < 1 {
		co.Row = make([]string, 1)
		co.Row[0] = "{{Name}}"
	}

	// Validate the options and fix
	if len(co.Row) != len(co.Spacing) {
		co.Spacing = makeIntSliceRequiredSize(co.Spacing, -1, len(co.Row))
	}

	// make sure we have a row data for each column
	if len(co.Header) > 0 && len(co.Header) != len(co.Row) {
		return errors.New("number of headers does not match number of row columns")
	}

	// see if we have alignment for each column if not set it to left
	co.Align, err = defaultValues("align", co.Align, "<", len(co.Row))
	if err != nil {
		return err
	}

	//see if we have the header padding if not set it to space
	co.HeaderPad, err = defaultValues("header pad", co.HeaderPad, " ", len(co.Row))
	if err != nil {
		return err
	}

	//see if we have the row padding if not set it to space
	co.RowPad, err = defaultValues("row pad", co.RowPad, " ", len(co.Row))
	if err != nil {
		return err
	}

	//see if we have the row padding if not set it to space
	co.FooterPad, err = defaultValues("footer pad", co.FooterPad, " ", len(co.Row))
	if err != nil {
		return err
	}

	//see if we have the footer columns matches the number of header columns
	if len(co.Footer) > 0 && len(co.Header) != len(co.Footer) {
		return errors.New("number of headers does not match number of Footers ")
	}

	return nil
}

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

// Function to remove duplicates from a string slice
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}

	for i := len(strSlice) - 1; i >= 0; i-- {
		if strSlice[i] == "" {
			continue
		}
		if _, value := allKeys[strSlice[i]]; !value {
			allKeys[strSlice[i]] = true
			list = append(list, strSlice[i])
		}
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

	return list
}

// Function to add space to a string
func addSpace(s string) string {
	buf := &bytes.Buffer{}
	for i, rune := range s {
		if unicode.IsUpper(rune) && i > 0 {
			buf.WriteRune(' ')
		}
		buf.WriteRune(rune)
	}
	return buf.String()
}

// Function to make a fixed width document
func fixedWidth(f *Faker, co *FixedWidthOptions) ([]byte, error) {

	if co == nil {
		// We didn't get a CSVOptions, so create a new random one
		err := f.Struct(&co)
		if err != nil {
			return nil, err
		}
	}

	// Validate the options and set defaults where needed
	err := validateInputs(co)
	if err != nil {
		return nil, err
	}

	// Setup the template engine and custom functions
	funcMap := template.FuncMap{}
	funcMap["Pad"] = strPad
	co.SetFuncs(&funcMap)

	// Setup  variables
	column_sizes := make([]int, len(co.Row)) //save the largest column size

	// Setup the running total
	co.runningTotal = make([]float64, len(co.Footer))
	for k := range co.Footer {
		co.runningTotal[k] = -999999999999999
	}

	// Setup the spacing and see if it automatic or specified
	column_spacing := make([]string, len(co.Row)) //tmp store the spacing for each column
	for k := range co.Spacing {
		column_spacing[k] = fmt.Sprintf("%v", co.Spacing[k])
		if co.Spacing[k] < 0 {
			//user the auto format
			column_spacing[k] = fmt.Sprintf("(.GetSpacing %v)", k)
		}
	}

	// Build and Render the header
	header_str := ""
	for k := range co.Header {
		co.currentColumn = k //Set the current column
		//check if the header is auto header or specified
		headerValue := co.Header[k]
		if strings.Contains(headerValue, "AutoHeader") {
			//get the functions from the row column cleanup and use as the title
			newHeader := co.Row[k]
			newHeader = strings.ReplaceAll(newHeader, "{", " ")
			newHeader = strings.ReplaceAll(newHeader, "}", " ")
			newHeader = strings.ReplaceAll(newHeader, "  ", "")
			newHeader = strings.Trim(newHeader, " ")
			newHeader = addSpace(newHeader)
			parts := strings.Split(newHeader, " ")
			parts = removeDuplicateStr(parts)
			headerValue = strings.Join(parts, "_")
		} else {
			// Render the column
			column_data, err := templateFunc(f, co.Header[k], co)
			if err != nil {
				return nil, err
			}
			headerValue = string(column_data)
		}
		//get the width and check if its the larger than the current column size
		if len(headerValue) > column_sizes[k] {
			column_sizes[k] = len(headerValue)
		}
		header_str += fmt.Sprintf("{{Pad `%s` %s  `%s` `%s` true}}", headerValue, column_spacing[k], co.HeaderPad[k], co.Align[k])
	}

	//rendered_template, err = templateFunc(f, header_str, co)
	//if err != nil {
	//	return nil, err
	//}
	//header_str = string(rendered_template)

	// Build the rows and loop through them
	rows := make([]string, co.Count)
	for r := range rows {
		rows[r] = ""

		//Loop though the headers
		for k := range co.Row {
			co.currentColumn = k //Set the current column

			// Render the column
			column_data, err := templateFunc(f, co.Row[k], co)
			if err != nil {
				return nil, err
			}

			// if the data is numeric then add it to the running total
			if len(co.Footer) > 0 {
				if regexp.MustCompile(`\d`).MatchString(string(column_data)) {
					value, err := strconv.ParseFloat(string(column_data), 32)
					if err == nil {
						//the data is numeric so check if value -99999... the set it to zero
						if co.runningTotal[k] == -999999999999999 {
							co.runningTotal[k] = 0
						}
						co.runningTotal[k] += value
					}
				}
			}

			//get the width and check if its the larger than the current column size
			if len(column_data) > column_sizes[k] {
				column_sizes[k] = len(column_data)
			}

			// Add the column to the row
			rows[r] += fmt.Sprintf("{{Pad `%s` %v `%s` `%s` true}}", column_data, column_spacing[k], co.RowPad[k], co.Align[k])
		}

	}

	// Build the footer if footer was configured
	footer_str := ""
	// check to see if we have a footer
	if len(co.Footer) > 0 {

		// Loop through the total columns and render
		for k := range co.Footer {

			co.currentColumn = k //Set the current column

			// Render the column
			column_data, err := templateFunc(f, co.Footer[k], co)
			if err != nil {
				return nil, err
			}

			//get the width and check if its the larger than the current column size
			if len(column_data) > column_sizes[k] {
				column_sizes[k] = len(column_data)
			}

			// build the footer column string
			footer_str += fmt.Sprintf("{{Pad `%s` %v `%s` `%s` true}}", column_data, column_spacing[k], co.FooterPad[k], co.Align[k])
		}

	}

	//need to update the column spacing to the largest column size if set to auto
	for k := range co.Spacing {
		if co.Spacing[k] < 0 {

			co.Spacing[k] = column_sizes[k]
		}
	}

	// Build the final result
	fixed_width_result := ""
	if len(co.Header) > 0 {
		fixed_width_result = fmt.Sprintf("%s\n", header_str)
	}

	//add the rows
	fixed_width_result = fmt.Sprintf("%s%s\n", fixed_width_result, strings.Join(rows, "\n"))

	//add the footer
	if len(co.Footer) > 0 {
		fixed_width_result = fmt.Sprintf("%s%s\n", fixed_width_result, footer_str)
	}

	// Render the column
	final_render, err := templateFunc(f, fixed_width_result, co)
	if err != nil {
		return nil, err
	}

	return final_render, nil
}

// function to add the fixed width lookup
func addFixedWidthLookup() {
	AddFuncLookup("fixed_width", Info{
		Display:     "FixedWidth",
		Category:    "template",
		Description: "Generates FixedWidth document",
		Example: `
		"Name==========================Email===============================Cost=============Account"
		"Markus Moen                   alaynawuckert@kozey.biz               8000000000948995369063"
		"Freeda Keebler                gabriellehuels@borer.io               6900000000932452944030"
		"****************************************************************218.00********************"`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "header", Display: "Header data", Type: "[]string", Default: "[Name,Email]", Description: "Header data for each column"},
			{Field: "row", Display: "Row data", Type: "[]string", Default: "[{{Name}},{{Email}}]", Description: "Row data for each column"},
			{Field: "footer", Display: "Footer Data", Type: "[]string", Optional: true, Default: "[,]", Description: "Footer data for each column"},
			{Field: "header_pad", Display: "Header Padding", Type: "[]string", Optional: true, Default: "[,]", Description: "What character to use for padding the footer column"},
			{Field: "row_pad", Display: "Row Padding", Type: "[]string", Optional: true, Default: "[,]", Description: "What character to use for padding the footer column"},
			{Field: "row_pad", Display: "Row Padding", Type: "[]string", Optional: true, Default: "[,]", Description: "What character to use for padding the footer column"},
			{Field: "footer_pad", Display: "Footer Padding", Type: "[]string", Optional: true, Default: "[,]", Description: "What character to use for padding the footer column"},
			{Field: "align", Display: "Column Alignment", Type: "[]string", Optional: true, Default: "[=,=]", Description: "Alignment of the columns < left, > right or ^ center"},
			{Field: "spacing", Display: "Column Spacing", Type: "[]int", Default: "[10,10]", Description: "Required Column Spacing"},
			{Field: "count", Display: "Number of Rows", Type: "int", Default: "1", Description: "Number of Rows to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {

			//template to use
			header, err := info.GetStringArray(m, "header")
			if err != nil {
				return nil, err
			}
			row, err := info.GetStringArray(m, "row")
			if err != nil {
				return nil, err
			}
			footer, err := info.GetStringArray(m, "footer")
			if err != nil {
				return nil, err
			}
			header_pad, _ := info.GetStringArray(m, "header_pad")
			/*if err != nil {
				return nil, err
			}*/
			row_pad, _ := info.GetStringArray(m, "row_pad")
			//if err != nil {
			//	return nil, err
			//}
			footer_pad, _ := info.GetStringArray(m, "footer_pad")
			/*if err != nil {
				return nil, err
			}*/
			spacing, err := info.GetIntArray(m, "spacing")
			if err != nil {
				return nil, err
			}
			count, err := info.GetInt(m, "count")
			if err != nil {
				return nil, err
			}

			f := &Faker{Rand: r}
			opt := FixedWidthOptions{
				Header:    header,
				Row:       row,
				Footer:    footer,
				HeaderPad: header_pad,
				RowPad:    row_pad,
				FooterPad: footer_pad,
				Spacing:   spacing,
				Count:     count,
			}
			templateOut, err := f.FixedWidth(&opt)
			if err != nil {
				return nil, err
			}

			return fixString(string(templateOut)), nil
		},
	})
}
