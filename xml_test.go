package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleXML_single() {
	Seed(11)

	value, err := XML(&XMLOptions{
		Type:          "single",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      2,
		Indent:        true,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: <xml>
	//     <first_name>Sonny</first_name>
	//     <last_name>Stiedemann</last_name>
	//     <password>8nwf0o3sBXcR</password>
	// </xml>
}

func ExampleFaker_XML_single() {
	f := New(11)

	value, err := f.XML(&XMLOptions{
		Type:          "single",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      2,
		Indent:        true,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: <xml>
	//     <first_name>Sonny</first_name>
	//     <last_name>Stiedemann</last_name>
	//     <password>8nwf0o3sBXcR</password>
	// </xml>
}

func ExampleXML_array() {
	Seed(11)

	value, err := XML(&XMLOptions{
		Type:          "array",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      2,
		Indent:        true,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: <xml>
	//     <record>
	//         <first_name>Sonny</first_name>
	//         <last_name>Stiedemann</last_name>
	//         <password>8nwf0o3sBXcR</password>
	//     </record>
	//     <record>
	//         <first_name>Verda</first_name>
	//         <last_name>Brakus</last_name>
	//         <password>3beWLpq75Lua</password>
	//     </record>
	// </xml>
}

func ExampleFaker_XML_array() {
	f := New(11)

	value, err := f.XML(&XMLOptions{
		Type:          "array",
		RootElement:   "xml",
		RecordElement: "record",
		RowCount:      2,
		Indent:        true,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: <xml>
	//     <record>
	//         <first_name>Sonny</first_name>
	//         <last_name>Stiedemann</last_name>
	//         <password>8nwf0o3sBXcR</password>
	//     </record>
	//     <record>
	//         <first_name>Verda</first_name>
	//         <last_name>Brakus</last_name>
	//         <password>3beWLpq75Lua</password>
	//     </record>
	// </xml>
}

func TestXMLSingle(t *testing.T) {
	Seed(11)

	AddFuncLookup("randmap", Info{
		Display:     "Rand Map",
		Category:    "generate",
		Description: "",
		Example:     "",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return map[string]any{
				"string": "string value",
				"int":    123456789,
				"float":  123.456,
				"array":  []string{"one", "two", "three"},
			}, nil
		},
	})

	value, err := XML(&XMLOptions{
		Type:        "single",
		RootElement: "single",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
			{Name: "person", Function: "person"},
			{Name: "ints", Function: "shufflestrings", Params: MapParams{"strs": {"123", "456", "789"}}},
			{Name: "rand_map", Function: "randmap"},
		},
		Indent: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%v", value) == "" {
		t.Fatal("Value is empty")
	}
}

func TestXMLArray(t *testing.T) {
	Seed(11)

	AddFuncLookup("randmap", Info{
		Display:     "Rand Map",
		Category:    "generate",
		Description: "",
		Example:     "",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return map[string]any{
				"string": "string value",
				"int":    123456789,
				"float":  123.456,
				"array":  []string{"one", "two", "three"},
			}, nil
		},
	})

	value, err := XML(&XMLOptions{
		Type:          "array",
		RootElement:   "array",
		RecordElement: "record",
		RowCount:      10,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
			{Name: "person", Function: "person"},
			{Name: "ints", Function: "shufflestrings", Params: MapParams{"strs": {"123", "456", "789"}}},
			{Name: "rand_map", Function: "randmap"},
		},
		Indent: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	if fmt.Sprintf("%v", value) == "" {
		t.Fatal("Value is empty")
	}
}

func TestXMLLookup(t *testing.T) {
	faker := New(0)
	info := GetFuncLookup("xml")

	m := NewMapParams()
	m.Add("type", "array")
	m.Add("rowcount", "10")
	m.Add("indent", "true")
	m.Add("fields", `{"name":"id","function":"autoincrement"}`)
	m.Add("fields", `{"name":"first_name","function":"firstname"}`)
	m.Add("fields", `{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`)

	_, err := info.Generate(faker, m, info)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestXMLNoOptions(t *testing.T) {
	Seed(11)

	// if XMLOptions is nil -> get a random XMLOptions
	_, err := XML(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func BenchmarkXMLLookup100(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("xml")

		m := NewMapParams()
		m.Add("type", "array")
		m.Add("rowcount", "100")
		m.Add("fields", `{"name":"id","function":"autoincrement"}`)
		m.Add("fields", `{"name":"first_name","function":"firstname"}`)
		m.Add("fields", `{"name":"last_name","function":"lastname"}`)
		m.Add("fields", `{"name":"password","function":"password"}`)
		m.Add("fields", `{"name":"description","function":"paragraph"}`)
		m.Add("fields", `{"name":"created_at","function":"date"}`)

		_, err := info.Generate(faker, m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
