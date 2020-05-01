package gofakeit

import (
	"encoding/json"
	"errors"
)

type JSONOptions struct {
	Type       string  `json:"type"` // array or object
	RowCount   int     `json:"row_count"`
	Fields     []Field `json:"fields"`
	Whitespace bool    `json:"whitespace"`
}

type Field struct {
	Name     string              `json:"name"`
	Function string              `json:"function"`
	Params   map[string][]string `json:"params"`
}

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
				return nil, errors.New("Invalid function, does not exist")
			}

			value, err := funcInfo.Call(&field.Params, funcInfo)
			if err != nil {
				return nil, err
			}

			v[field.Name] = value
		}

		// Marshal into bytes
		j := []byte{}
		if jo.Whitespace {
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
					return nil, errors.New("Invalid function, does not exist")
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
		if jo.Whitespace {
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
		Description: "Random latitude between given range",
		Example:     "22.921026",
		Output:      "float",
		Params: []Param{
			{Field: "type", Display: "Type", Type: "string", Default: "object", Description: "Type of JSON, object or array"},
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows in JSON array"},
			{Field: "fields", Display: "Fields", Type: "map[string]string", Description: "Fields containing key name and function to run"},
			{Field: "whitespace", Display: "Whitespace", Type: "bool", Default: "false", Description: "Whether or not to add whitespace and newlines"},
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

			// fields, err := info.GetInt(m, "fields")
			// if err != nil {
			// 	return nil, err
			// }
			// jo.Fields = fields

			whitespace, err := info.GetBool(m, "whitespace")
			if err != nil {
				return nil, err
			}
			jo.Whitespace = whitespace

			rangeOut, err := JSON(&jo)
			if err != nil {
				return nil, err
			}

			return rangeOut, nil
		},
	})
}
