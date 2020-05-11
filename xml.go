package gofakeit

import (
	"bytes"
	"encoding/xml"
	"errors"
)

// XMLOptions defines values needed for json generation
type XMLOptions struct {
	Type          string  `json:"type"` // single or multiple
	RootElement   string  `json:"root_element"`
	RecordElement string  `json:"record_element"`
	RowCount      int     `json:"row_count"`
	Fields        []Field `json:"fields"`
	Indent        bool    `json:"indent"`
}

type xmlMap map[string]interface{}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   interface{} `xml:",chardata"`
}

func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for key, value := range m {
		// fmt.Println(fmt.Sprintf("%T", value))
		// x, err := xml.MarshalIndent(value, "", "  ")
		// if err != nil {
		// 	return err
		// }
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: key}, Value: value})
	}

	return e.EncodeToken(start.End())

	// tokens := []xml.Token{start}

	// for key, value := range m {
	// 	// fmt.Println(fmt.Sprintf("%T", value))
	// 	// x, err := xml.MarshalIndent(value, "", "    ")
	// 	// if err != nil {
	// 	// 	return err
	// 	// }

	// 	var b bytes.Buffer
	// 	x := xml.NewEncoder(&b)
	// 	// if jo.Indent {
	// 	// 	x.Indent("", "    ")
	// 	// }
	// 	err := x.Encode(value)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	t := xml.StartElement{Name: xml.Name{Space: "", Local: key}}
	// 	tokens = append(tokens, t, b, xml.EndElement{Name: t.Name})
	// }

	// tokens = append(tokens, xml.EndElement{Name: start.Name})

	// for _, t := range tokens {
	// 	err := e.EncodeToken(t)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	// // flush to ensure tokens are written
	// return e.Flush()
}

// XML generates an object or an array of objects in json format
func XML(jo *XMLOptions) ([]byte, error) {
	// Check to make sure they passed in a type
	if jo.Type != "multiple" && jo.Type != "single" {
		return nil, errors.New("Invalid type, must be array or object")
	}

	if jo.Fields == nil || len(jo.Fields) <= 0 {
		return nil, errors.New("Must pass fields in order to build json object(s)")
	}

	if jo.Type == "single" {
		v := xmlMap{}
		// v.XMLName = xml.Name{Local: jo.RecordElement}

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
		var b bytes.Buffer
		x := xml.NewEncoder(&b)
		if jo.Indent {
			x.Indent("", "    ")
		}
		err := x.Encode(v)
		if err != nil {
			return nil, err
		}

		return b.Bytes(), nil
	}

	// if jo.Type == "array" {
	// 	// Make sure you set a row count
	// 	if jo.RowCount <= 0 {
	// 		return nil, errors.New("Must have row count")
	// 	}

	// 	v := []map[string]interface{}{}

	// 	for i := 1; i <= int(jo.RowCount); i++ {
	// 		vr := map[string]interface{}{}

	// 		// Loop through fields and add to them to map[string]interface{}
	// 		for _, field := range jo.Fields {
	// 			if field.Function == "autoincrement" {
	// 				vr[field.Name] = i
	// 				continue
	// 			}

	// 			// Get function info
	// 			funcInfo := GetFuncLookup(field.Function)
	// 			if funcInfo == nil {
	// 				return nil, errors.New("Invalid function, " + field.Function + " does not exist")
	// 			}

	// 			value, err := funcInfo.Call(&field.Params, funcInfo)
	// 			if err != nil {
	// 				return nil, err
	// 			}

	// 			vr[field.Name] = value
	// 		}

	// 		v = append(v, vr)
	// 	}

	// 	// Marshal into bytes
	// 	j := []byte{}
	// 	if jo.Indent {
	// 		j, _ = xml.MarshalIndent(v, "", "    ")
	// 	} else {
	// 		j, _ = xml.Marshal(v)
	// 	}
	// 	return j, nil
	// }

	return nil, errors.New("Invalid type, must be array or object")
}
