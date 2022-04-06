package gofakeit

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type SQLOptions struct {
	Table      string  `json:"table" xml:"table"`
	EntryCount int     `json:"entry_count" xml:"entry_count"`
	Fields     []Field `json:"fields" xml:"fields"`
}

func SQLInsert(so *SQLOptions) ([]byte, error) {
	return sqlInsertFunc(globalFaker.Rand, so)
}

func sqlInsertFunc(r *rand.Rand, so *SQLOptions) ([]byte, error) {
	if so.Table == "" {
		return nil, errors.New("must provide table name to generate SQL")
	}
	if so.Fields == nil || len(so.Fields) <= 0 {
		return nil, errors.New(("must pass fields in order to generate SQL queries"))
	}
	if so.EntryCount <= 0 {
		return nil, errors.New("must have entry count")
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO " + so.Table + " VALUES")

	for i := 0; i < so.EntryCount; i++ {
		sb.WriteString(" (")
		// Now, we need to add all of our fields
		for i, field := range so.Fields {
			if field.Function == "autoincrement" { // One
				// TODO: We need to do something here still...
				continue
			}

			// Get the function info for the field
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return nil, errors.New("invalid function, " + field.Function + " does not exist")
			}

			// Generate the value
			val, err := funcInfo.Generate(r, &field.Params, funcInfo)
			if err != nil {
				return nil, err
			}

			convertType := ConvertType(funcInfo.Output, val)

			if i == len(so.Fields)-1 { // Last field
				sb.WriteString(convertType)
			} else {
				sb.WriteString(convertType + ", ")
			}
		}
		if i == so.EntryCount-1 { // Last tuple
			sb.WriteString(");")
		} else {
			sb.WriteString("),")
		}
	}

	return []byte(sb.String()), nil
}

func ConvertType(t string, val interface{}) string {
	switch t {
	case "string":
		return `'` + fmt.Sprintf("%v", val) + `'`
	default:
		return fmt.Sprintf("%v", val)
	}
}
