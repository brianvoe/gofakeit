package gofakeit

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
)

// JSONOptions defines values needed for json generation
type JSONOptions struct {
	Type     string  `json:"type" xml:"type"` // array or object
	RowCount int     `json:"row_count" xml:"row_count"`
	Fields   []Field `json:"fields" xml:"fields"`
	Indent   bool    `json:"indent" xml:"indent"`
}

type jsonKeyVal struct {
	Key   string
	Value interface{}
}

type jsonOrderedKeyVal []*jsonKeyVal

func (okv jsonOrderedKeyVal) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	buf.WriteString("{")
	for i, kv := range okv {
		// Add comma to all except last one
		if i != 0 {
			buf.WriteString(",")
		}

		// Marshal key and write
		key, err := json.Marshal(kv.Key)
		if err != nil {
			return nil, err
		}
		buf.Write(key)

		// Write colon separator
		buf.WriteString(":")

		// Marshal value and write
		val, err := json.Marshal(kv.Value)
		if err != nil {
			return nil, err
		}
		buf.Write(val)
	}
	buf.WriteString("}")

	return buf.Bytes(), nil
}

// JSON generates an object or an array of objects in json format
func JSON(jo *JSONOptions) ([]byte, error) { return jsonFunc(globalFaker.Rand, jo) }

// JSON generates an object or an array of objects in json format
func (f *Faker) JSON(jo *JSONOptions) ([]byte, error) { return jsonFunc(f.Rand, jo) }

// JSON generates an object or an array of objects in json format
func jsonFunc(r *rand.Rand, jo *JSONOptions) ([]byte, error) {
	// Check to make sure they passed in a type
	if jo.Type != "array" && jo.Type != "object" {
		return nil, errors.New("invalid type, must be array or object")
	}

	if jo.Fields == nil || len(jo.Fields) <= 0 {
		return nil, errors.New("must pass fields in order to build json object(s)")
	}

	if jo.Type == "object" {
		v := make(jsonOrderedKeyVal, len(jo.Fields))

		// Loop through fields and add to them to map[string]interface{}
		for i, field := range jo.Fields {
			if field.Function == "autoincrement" {
				// Object only has one
				v[i] = &jsonKeyVal{Key: field.Name, Value: 1}
				continue
			}

			// Get function info
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return nil, errors.New("invalid function, " + field.Function + " does not exist")
			}

			// Call function value
			value, err := funcInfo.Generate(r, &field.Params, funcInfo)
			if err != nil {
				return nil, err
			}

			if _, ok := value.([]byte); ok {
				// If it's a slice, unmarshal it into an interface
				var val interface{}
				err := json.Unmarshal(value.([]byte), &val)
				if err != nil {
					return nil, err
				}
				value = val
			}

			v[i] = &jsonKeyVal{Key: field.Name, Value: value}

		}

		// Marshal into bytes
		if jo.Indent {
			j, _ := json.MarshalIndent(v, "", "    ")
			return j, nil
		}

		j, _ := json.Marshal(v)
		return j, nil
	}

	if jo.Type == "array" {
		// Make sure you set a row count
		if jo.RowCount <= 0 {
			return nil, errors.New("must have row count")
		}

		v := make([]jsonOrderedKeyVal, jo.RowCount)

		for i := 0; i < int(jo.RowCount); i++ {
			vr := make(jsonOrderedKeyVal, len(jo.Fields))

			// Loop through fields and add to them to map[string]interface{}
			for ii, field := range jo.Fields {
				if field.Function == "autoincrement" {
					vr[ii] = &jsonKeyVal{Key: field.Name, Value: i + 1} // +1 because index starts with 0
					continue
				}

				// Get function info
				funcInfo := GetFuncLookup(field.Function)
				if funcInfo == nil {
					return nil, errors.New("invalid function, " + field.Function + " does not exist")
				}

				// Call function value
				value, err := funcInfo.Generate(r, &field.Params, funcInfo)
				if err != nil {
					return nil, err
				}

				if _, ok := value.([]byte); ok {
					// If it's a slice, unmarshal it into an interface
					var val interface{}
					err := json.Unmarshal(value.([]byte), &val)
					if err != nil {
						return nil, err
					}
					value = val
				}

				vr[ii] = &jsonKeyVal{Key: field.Name, Value: value}
			}

			v[i] = vr
		}

		// Marshal into bytes
		if jo.Indent {
			j, _ := json.MarshalIndent(v, "", "    ")
			return j, nil
		}

		j, _ := json.Marshal(v)
		return j, nil
	}

	return nil, errors.New("invalid type, must be array or object")
}

func addFileJSONLookup() {
	AddFuncLookup("json", Info{
		Display:     "JSON",
		Category:    "file",
		Description: "Generates an object or an array of objects in json format",
		Example: `[
			{ "id": 1, "first_name": "Markus", "last_name": "Moen" },
			{ "id": 2, "first_name": "Alayna", "last_name": "Wuckert" },
			{ "id": 3, "first_name": "Lura", "last_name": "Lockman" }
		]`,
		Output:      "[]byte",
		ContentType: "application/json",
		Params: []Param{
			{Field: "type", Display: "Type", Type: "string", Default: "object", Options: []string{"object", "array"}, Description: "Type of JSON, object or array"},
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows in JSON array"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function to run in json format"},
			{Field: "indent", Display: "Indent", Type: "bool", Default: "false", Description: "Whether or not to add indents and newlines"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			jo := JSONOptions{}

			typ, err := info.GetString(m, "type")
			if err != nil {
				return nil, err
			}
			jo.Type = typ

			rowcount, err := info.GetInt(m, "rowcount")
			if err != nil {
				return nil, err
			}
			jo.RowCount = rowcount

			fieldsStr, err := info.GetStringArray(m, "fields")
			if err != nil {
				return nil, err
			}

			// Check to make sure fields has length
			if len(fieldsStr) > 0 {
				jo.Fields = make([]Field, len(fieldsStr))

				for i, f := range fieldsStr {
					// Unmarshal fields string into fields array
					err = json.Unmarshal([]byte(f), &jo.Fields[i])
					if err != nil {
						return nil, err
					}
				}
			}

			indent, err := info.GetBool(m, "indent")
			if err != nil {
				return nil, err
			}
			jo.Indent = indent

			return jsonFunc(r, &jo)
		},
	})
}
