package gofakeit

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// XMLOptions defines values needed for json generation
type XMLOptions struct {
	Type          string  `json:"type" xml:"type" fake:"{randomstring:[array,single]}"` // single or array
	RootElement   string  `json:"root_element" xml:"root_element"`
	RecordElement string  `json:"record_element" xml:"record_element"`
	RowCount      int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Indent        bool    `json:"indent" xml:"indent"`
	Fields        []Field `json:"fields" xml:"fields" fake:"{fields}"`
}

type xmlArray struct {
	XMLName xml.Name
	Array   []xmlMap
}

type xmlMap struct {
	XMLName  xml.Name
	KeyOrder []string
	Map      map[string]any `xml:",chardata"`
}

type xmlEntry struct {
	XMLName xml.Name
	Value   any `xml:",chardata"`
}

func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m.Map) == 0 {
		return nil
	}

	start.Name = m.XMLName

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = xmlMapLoop(e, &m)
	if err != nil {
		return err
	}

	return e.EncodeToken(start.End())
}

func xmlMapLoop(e *xml.Encoder, m *xmlMap) error {
	// Check if xmlmap has key order if not create it
	// Get key order by order of fields array
	if m.KeyOrder == nil {
		m.KeyOrder = make([]string, 0, len(m.Map))
		for k := range m.Map {
			m.KeyOrder = append(m.KeyOrder, k)
		}

		// A map is read in a random order, sort the keys so the same seed gives the same xml
		sort.Strings(m.KeyOrder)
	}

	for _, key := range m.KeyOrder {
		if err := xmlEncodeValue(e, key, m.Map[key]); err != nil {
			return err
		}
	}

	return nil
}

// xmlEncodeValue writes value as an element called name. A scalar is its text, an item of a slice is
// a value element of its own and a map or struct has an element for each of its keys, at any depth.
// A nil value or nil pointer writes nothing.
func xmlEncodeValue(e *xml.Encoder, name string, value any) error {
	if v := reflect.ValueOf(value); !v.IsValid() || (v.Kind() == reflect.Pointer && v.IsNil()) {
		return nil
	}

	// An error is its message and a time is its text, they would come out empty below
	value, err := fieldValue(value)
	if err != nil {
		return err
	}

	start := xml.StartElement{Name: xml.Name{Local: name}}

	switch val := value.(type) {
	case map[string]any:
		return e.Encode(xmlMap{XMLName: start.Name, Map: val})
	case []byte:
		// The output of a function like json is written as its data, like the json and csv output do
		if json.Valid(val) {
			var data any
			if err := json.Unmarshal(val, &data); err != nil {
				return err
			}
			return xmlEncodeValue(e, name, data)
		}
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Pointer:
		return xmlEncodeValue(e, name, v.Elem().Interface())
	case reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return e.Encode(xmlEntry{XMLName: start.Name, Value: value})
	case reflect.Slice, reflect.Array:
		if err := e.EncodeToken(start); err != nil {
			return err
		}
		for i := range v.Len() {
			if err := xmlEncodeValue(e, "value", v.Index(i).Interface()); err != nil {
				return err
			}
		}
		return e.EncodeToken(start.End())
	case reflect.Map:
		// Any map, like the map[string][]string of teams, gets an element for each key
		m := make(map[string]any, v.Len())
		iter := v.MapRange()
		for iter.Next() {
			m[fmt.Sprint(iter.Key().Interface())] = iter.Value().Interface()
		}
		return e.Encode(xmlMap{XMLName: start.Name, Map: m})
	case reflect.Struct:
		// Go through json so the elements get the json names of the fields
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		var data any
		if err := json.Unmarshal(b, &data); err != nil {
			return err
		}
		return xmlEncodeValue(e, name, data)
	}

	return e.Encode(value)
}

// XML generates an object or an array of objects in json format
// A nil XMLOptions returns a randomly structured XML.
func XML(xo *XMLOptions) ([]byte, error) { return xmlFunc(GlobalFaker, xo) }

// XML generates an object or an array of objects in json format
// A nil XMLOptions returns a randomly structured XML.
func (f *Faker) XML(xo *XMLOptions) ([]byte, error) { return xmlFunc(f, xo) }

func xmlFunc(f *Faker, xo *XMLOptions) ([]byte, error) {
	if xo == nil {
		// We didn't get a XMLOptions, so create a new random one
		err := f.Struct(&xo)
		if err != nil {
			return nil, err
		}
	}

	// Check to make sure they passed in a type
	if xo.Type != "single" && xo.Type != "array" {
		return nil, errors.New("invalid type, must be array or object")
	}

	// Check fields length
	if len(xo.Fields) <= 0 {
		return nil, errors.New("must pass fields in order to build json object(s)")
	}

	// Default the root and record element names, without changing the options that were passed in
	rootElement := xo.RootElement
	if rootElement == "" {
		rootElement = "xml"
	}
	recordElement := xo.RecordElement
	if recordElement == "" {
		recordElement = "record"
	}

	// Get key order by order of fields array
	keyOrder := make([]string, 0, len(xo.Fields))
	for _, f := range xo.Fields {
		keyOrder = append(keyOrder, f.Name)
	}

	if xo.Type == "single" {
		v := xmlMap{
			XMLName:  xml.Name{Local: rootElement},
			KeyOrder: keyOrder,
			Map:      make(map[string]any),
		}

		// Loop through fields and add to them to map[string]any
		for _, field := range xo.Fields {
			// Get function info
			funcInfo := GetFuncLookup(field.Function)
			if funcInfo == nil {
				return nil, errors.New("invalid function, " + field.Function + " does not exist")
			}

			value, err := funcInfo.Generate(f, &field.Params, funcInfo)
			if err != nil {
				return nil, err
			}

			v.Map[field.Name] = value
		}

		// Marshal into bytes
		var b bytes.Buffer
		x := xml.NewEncoder(&b)
		if xo.Indent {
			x.Indent("", "    ")
		}
		err := x.Encode(v)
		if err != nil {
			return nil, err
		}

		return b.Bytes(), nil
	}

	if xo.Type == "array" {
		// Make sure you set a row count
		if xo.RowCount <= 0 {
			return nil, errors.New("must have row count")
		}

		xa := xmlArray{
			XMLName: xml.Name{Local: rootElement},
			Array:   make([]xmlMap, 0, xo.RowCount),
		}

		for i := 1; i <= int(xo.RowCount); i++ {
			v := xmlMap{
				XMLName:  xml.Name{Local: recordElement},
				KeyOrder: keyOrder,
				Map:      make(map[string]any),
			}

			// Loop through fields and add to them to map[string]any
			for _, field := range xo.Fields {
				if field.Function == "autoincrement" {
					v.Map[field.Name] = i
					continue
				}

				// Get function info
				funcInfo := GetFuncLookup(field.Function)
				if funcInfo == nil {
					return nil, errors.New("invalid function, " + field.Function + " does not exist")
				}

				value, err := funcInfo.Generate(f, &field.Params, funcInfo)
				if err != nil {
					return nil, err
				}

				v.Map[field.Name] = value
			}

			xa.Array = append(xa.Array, v)
		}

		// Marshal into bytes
		var b bytes.Buffer
		x := xml.NewEncoder(&b)
		if xo.Indent {
			x.Indent("", "    ")
		}
		err := x.Encode(xa)
		if err != nil {
			return nil, err
		}

		return b.Bytes(), nil
	}

	return nil, errors.New("invalid type, must be array or object")
}

func addFileXMLLookup() {
	AddFuncLookup("xml", Info{
		Display:     "XML",
		Category:    "file",
		Description: "Generates an single or an array of elements in xml format",
		Example: `<xml>
	<record>
		<first_name>Markus</first_name>
		<last_name>Moen</last_name>
		<password>Dc0VYXjkWABx</password>
	</record>
	<record>
		<first_name>Osborne</first_name>
		<last_name>Hilll</last_name>
		<password>XPJ9OVNbs5lm</password>
	</record>
</xml>`,
		Output:      "[]byte",
		ContentType: "application/xml",
		Aliases: []string{
			"xml document",
			"extensible markup",
			"tagged data",
			"hierarchical structure",
			"serialized tree",
		},
		Keywords: []string{
			"extensible", "markup", "language", "elements", "format",
			"structured", "generates", "tags", "attributes", "nested",
		},
		Params: []Param{
			{Field: "type", Display: "Type", Type: "string", Default: "single", Options: []string{"single", "array"}, Description: "Type of XML, single or array"},
			{Field: "rootelement", Display: "Root Element", Type: "string", Default: "xml", Description: "Root element wrapper name"},
			{Field: "recordelement", Display: "Record Element", Type: "string", Default: "record", Description: "Record element for each record row"},
			{Field: "rowcount", Display: "Row Count", Type: "int", Default: "100", Description: "Number of rows in JSON array"},
			{Field: "indent", Display: "Indent", Type: "bool", Default: "false", Description: "Whether or not to add indents and newlines"},
			{Field: "fields", Display: "Fields", Type: "[]Field", Description: "Fields containing key name and function to run in json format"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			xo := XMLOptions{}

			typ, err := info.GetString(m, "type")
			if err != nil {
				return nil, err
			}
			xo.Type = typ

			rootElement, err := info.GetString(m, "rootelement")
			if err != nil {
				return nil, err
			}
			xo.RootElement = rootElement

			recordElement, err := info.GetString(m, "recordelement")
			if err != nil {
				return nil, err
			}
			xo.RecordElement = recordElement

			rowcount, err := info.GetInt(m, "rowcount")
			if err != nil {
				return nil, err
			}
			xo.RowCount = rowcount

			fieldsStr, err := info.GetStringArray(m, "fields")
			if err != nil {
				return nil, err
			}

			indent, err := info.GetBool(m, "indent")
			if err != nil {
				return nil, err
			}
			xo.Indent = indent

			// Check to make sure fields has length
			if len(fieldsStr) > 0 {
				xo.Fields = make([]Field, len(fieldsStr))

				for i, f := range fieldsStr {
					// Unmarshal fields string into fields array
					err = json.Unmarshal([]byte(f), &xo.Fields[i])
					if err != nil {
						return nil, errors.New("unable to decode json string")
					}
				}
			}

			return xmlFunc(f, &xo)
		},
	})
}
