package gofakeit

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

// CSVOptions defines values needed for csv generation
type CSVOptions struct {
	Delimiter string  `json:"delimiter" xml:"delimiter"`
	RowCount  int     `json:"row_count" xml:"row_count"`
	Fields    []Field `json:"fields" xml:"fields"`
}

// CSV generates an object or an array of objects in json format
func CSV(co *CSVOptions) ([]byte, error) { return csvFunc(globalFaker.Rand, co) }

// CSV generates an object or an array of objects in json format
func (f *Faker) CSV(co *CSVOptions) ([]byte, error) { return csvFunc(f.Rand, co) }

func csvFunc(r *rand.Rand, co *CSVOptions) ([]byte, error) {
	// Check delimiter
	if co.Delimiter == "" {
		co.Delimiter = ","
	}
	if strings.ToLower(co.Delimiter) == "tab" {
		co.Delimiter = "\t"
	}
	if co.Delimiter != "," && co.Delimiter != "\t" {
		return nil, errors.New("invalid delimiter type")
	}

	// Check fields
	if co.Fields == nil || len(co.Fields) <= 0 {
		return nil, errors.New("must pass fields in order to build json object(s)")
	}

	// Make sure you set a row count
	if co.RowCount <= 0 {
		return nil, errors.New("must have row count")
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	w.Comma = []rune(co.Delimiter)[0]

	// Add header row
	header := make([]string, len(co.Fields))
	for i, field := range co.Fields {
		header[i] = field.Name
	}
	w.Write(header)

	// Loop through row count +1(for header) and add fields
	for i := 1; i < co.RowCount+1; i++ {
		vr := make([]string, len(co.Fields))

		// Loop through fields and add to them to map[string]interface{}
		for ii, field := range co.Fields {
			if field.Function == "autoincrement" {
				vr[ii] = fmt.Sprintf("%d", i)
				continue
			}

			// Get function info
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return nil, errors.New("invalid function, " + field.Function + " does not exist")
			}

			value, err := funcInfo.Generate(r, &field.Params, funcInfo)
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
					return nil, err
				}
				value = string(b)
			}

			vr[ii] = fmt.Sprintf("%v", value)
		}

		w.Write(vr)
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func addFileCSVLookup() {
	AddFuncLookup("csv", Info{
		Display:     "CSV",
		Category:    "file",
		Description: "Generates array of rows in csv format",
		Example: `
			id,first_name,last_name,password
			1,Markus,Moen,Dc0VYXjkWABx
			2,Osborne,Hilll,XPJ9OVNbs5lm
		`,
		Output:      "[]byte",
		ContentType: "text/csv",
		Params: []Param{
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function"},
			{Field: "delimiter", Display: "Delimiter", Type: "string", Default: ",", Description: "Separator in between row values"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := CSVOptions{}

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

			delimiter, err := info.GetString(m, "delimiter")
			if err != nil {
				return nil, err
			}
			co.Delimiter = delimiter

			csvOut, err := csvFunc(r, &co)
			if err != nil {
				return nil, err
			}

			return csvOut, nil
		},
	})
}
