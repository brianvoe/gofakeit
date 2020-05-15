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
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <xml>
	//     <first_name>Markus</first_name>
	//     <last_name>Moen</last_name>
	//     <password>Dc0VYXjkWABx</password>
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
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <xml>
	//     <record>
	//         <first_name>Markus</first_name>
	//         <last_name>Moen</last_name>
	//         <password>Dc0VYXjkWABx</password>
	//     </record>
	//     <record>
	//         <first_name>Osborne</first_name>
	//         <last_name>Hilll</last_name>
	//         <password>XPJ9OVNbs5lm</password>
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
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return map[string]interface{}{
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
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
			{Name: "person", Function: "person"},
			{Name: "ints", Function: "shufflestrings", Params: map[string][]string{"strs": {"123", "456", "789"}}},
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
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return map[string]interface{}{
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
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
			{Name: "person", Function: "person"},
			{Name: "ints", Function: "shufflestrings", Params: map[string][]string{"strs": {"123", "456", "789"}}},
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
	info := GetFuncLookup("xml")

	m := map[string][]string{
		"type":     {"array"},
		"rowcount": {"10"},
		"indent":   {"true"},
		"fields": {
			`{"name":"id","function":"autoincrement"}`,
			`{"name":"first_name","function":"firstname"}`,
			`{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`,
		},
	}
	_, err := info.Call(&m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// t.Fatal(fmt.Sprintf("%s", value.([]byte)))
}

func BenchmarkXMLLookup100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("xml")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"100"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkXMLLookup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("xml")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"1000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkXMLLookup10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("xml")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"10000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
func BenchmarkXMLLookup100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("xml")
		m := map[string][]string{
			"type":     {"array"},
			"rowcount": {"100000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Call(&m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
