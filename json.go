package gofakeit

import (
	"encoding/json"
	"errors"
)

// JSONOptions defines values needed for json generation
type JSONOptions struct {
	Type     string  `json:"type" xml:"type"` // array or object
	RowCount int     `json:"row_count" xml:"row_count"`
	Fields   []Field `json:"fields" xml:"fields"`
	Indent   bool    `json:"indent" xml:"indent"`
}

// JSON generates an object or an array of objects in json format
func JSON(jo *JSONOptions) ([]byte, error) {
	// Check to make sure they passed in a type
	if jo.Type != "array" && jo.Type != "object" {
		return nil, errors.New("Invalid type, must be array or object")
	}

	if jo.Fields == nil || len(jo.Fields) <= 0 {
		return nil, errors.New("Must pass fields in order to build json object(s)")
	}

	if jo.Type == "object" {
		v := map[string]interface{}{}

		// Loop through fields and add to them to map[string]interface{}
		for _, field := range jo.Fields {
			// Get function info
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return nil, errors.New("Invalid function, " + field.Function + " does not exist")
			}

			value, err := funcInfo.Call(&field.Params, funcInfo)
			if err != nil {
				return nil, err
			}

			v[field.Name] = value
		}

		// Marshal into bytes
		j := []byte{}
		if jo.Indent {
			j, _ = json.MarshalIndent(v, "", "    ")
		} else {
			j, _ = json.Marshal(v)
		}
		return j, nil
	}

	if jo.Type == "array" {
		// Make sure you set a row count
		if jo.RowCount <= 0 {
			return nil, errors.New("Must have row count")
		}

		v := []map[string]interface{}{}

		for i := 1; i <= int(jo.RowCount); i++ {
			vr := map[string]interface{}{}

			// Loop through fields and add to them to map[string]interface{}
			for _, field := range jo.Fields {
				if field.Function == "autoincrement" {
					vr[field.Name] = i
					continue
				}

				// Get function info
				funcInfo := GetFuncLookup(field.Function)
				if funcInfo == nil {
					return nil, errors.New("Invalid function, " + field.Function + " does not exist")
				}

				value, err := funcInfo.Call(&field.Params, funcInfo)
				if err != nil {
					return nil, err
				}

				vr[field.Name] = value
			}

			v = append(v, vr)
		}

		// Marshal into bytes
		j := []byte{}
		if jo.Indent {
			j, _ = json.MarshalIndent(v, "", "    ")
		} else {
			j, _ = json.Marshal(v)
		}
		return j, nil
	}

	return nil, errors.New("Invalid type, must be array or object")
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
		Output: "[]byte",
		Params: []Param{
			{Field: "type", Display: "Type", Type: "string", Default: "object", Description: "Type of JSON, object or array"},
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows in JSON array"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function to run in json format"},
			{Field: "indent", Display: "Indent", Type: "bool", Default: "false", Description: "Whether or not to add indents and newlines"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
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
						return nil, errors.New("Unable to decode json string")
					}
				}
			}

			indent, err := info.GetBool(m, "indent")
			if err != nil {
				return nil, err
			}
			jo.Indent = indent

			return JSON(&jo)
		},
	})
}
