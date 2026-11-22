package gofakeit

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
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
	//     <first_name>Priscilla</first_name>
	//     <last_name>Thornton</last_name>
	//     <password>F395YfK9tHs9</password>
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
	//     <first_name>Priscilla</first_name>
	//     <last_name>Thornton</last_name>
	//     <password>F395YfK9tHs9</password>
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
	//         <first_name>Priscilla</first_name>
	//         <last_name>Thornton</last_name>
	//         <password>F395YfK9tHs9</password>
	//     </record>
	//     <record>
	//         <first_name>Emmett</first_name>
	//         <last_name>Black</last_name>
	//         <password>PL0G37o1pY0p</password>
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
	//         <first_name>Priscilla</first_name>
	//         <last_name>Thornton</last_name>
	//         <password>F395YfK9tHs9</password>
	//     </record>
	//     <record>
	//         <first_name>Emmett</first_name>
	//         <last_name>Black</last_name>
	//         <password>PL0G37o1pY0p</password>
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

// fixedValueLookup registers a lookup that always returns value, for the tests of how a value is written
func fixedValueLookup(t *testing.T, name string, value any) {
	t.Helper()
	AddFuncLookup(name, Info{Generate: func(*Faker, *MapParams, *Info) (any, error) { return value, nil }})
	t.Cleanup(func() { RemoveFuncLookup(name) })
}

func TestXMLValueTypes(t *testing.T) {
	type point struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	for _, tc := range []struct {
		name  string
		value any
		want  string
	}{
		// These used to panic, the map case only took map[string]any
		{"map_of_strings", map[string]string{"b": "2", "a": "1"}, "<r><v><a>1</a><b>2</b></v></r>"},
		{"map_of_slices", map[string][]string{"x": {"a", "b"}, "y": {"c"}}, "<r><v><x><value>a</value><value>b</value></x><y><value>c</value></y></v></r>"},
		// These used to write nothing
		{"error", errors.New("not found"), "<r><v>not found</v></r>"},
		{"time", time.Date(2026, 11, 22, 10, 30, 0, 0, time.UTC), "<r><v>2026-11-22T10:30:00Z</v></r>"},
		// These used to write empty value elements
		{"slice_of_maps", []map[string]any{{"a": 1}, {"a": 2}}, "<r><v><value><a>1</a></value><value><a>2</a></value></v></r>"},
		{"slice_of_structs", []point{{1, 2}}, "<r><v><value><x>1</x><y>2</y></value></v></r>"},
		{"nested_slices", [][]int{{1, 2}, {3}}, "<r><v><value><value>1</value><value>2</value></value><value><value>3</value></value></v></r>"},
		{"slice_of_errors", []error{errors.New("a"), errors.New("b")}, "<r><v><value>a</value><value>b</value></v></r>"},
		// The output of a function like json is written as its data, not as a list of bytes
		{"json_bytes", []byte(`{"id":7,"tags":["a","b"]}`), "<r><v><id>7</id><tags><value>a</value><value>b</value></tags></v></r>"},
		{"struct", point{3, 4}, "<r><v><x>3</x><y>4</y></v></r>"},
		{"pointer_to_struct", &point{3, 4}, "<r><v><x>3</x><y>4</y></v></r>"},
		{"nil_pointer", (*point)(nil), "<r></r>"},
		{"nil", nil, "<r></r>"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixedValueLookup(t, "xml_value_types_"+tc.name, tc.value)

			got, err := New(11).XML(&XMLOptions{
				Type:        "single",
				RootElement: "r",
				Fields:      []Field{{Name: "v", Function: "xml_value_types_" + tc.name}},
			})
			if err != nil {
				t.Fatal(err)
			}
			if string(got) != tc.want {
				t.Fatalf("got %s, want %s", got, tc.want)
			}
		})
	}
}

func TestXMLBuiltinValueTypes(t *testing.T) {
	// teams returns a map[string][]string, it used to panic
	got, err := New(11).XML(&XMLOptions{
		Type:        "single",
		RootElement: "r",
		Fields:      []Field{{Name: "v", Function: "teams", Params: MapParams{"people": {"a", "b", "c"}, "teams": {"x", "y"}}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "<x><value>") || !strings.Contains(string(got), "<y><value>") {
		t.Fatalf("teams: got %s, want an element for each team", got)
	}

	// The errors and dates used to write nothing
	for _, function := range []string{"errorhttp", "errordatabase", "pastdate", "futuredate"} {
		got, err := New(11).XML(&XMLOptions{Type: "single", RootElement: "r", Fields: []Field{{Name: "v", Function: function}}})
		if err != nil {
			t.Fatal(err)
		}
		if !strings.HasPrefix(string(got), "<r><v>") || len(got) <= len("<r><v></v></r>") {
			t.Fatalf("%s: got %s, want a v element with text", function, got)
		}
	}
}

func TestXMLDefaultElements(t *testing.T) {
	fields := []Field{{Name: "v", Function: "autoincrement"}}

	// An empty root element is xml and an empty record element is record, as the docs say.
	// A single document used to fail with "start tag with no name".
	single := &XMLOptions{Type: "single", Fields: []Field{{Name: "v", Function: "xml_default_elements"}}}
	fixedValueLookup(t, "xml_default_elements", "a")
	got, err := New(11).XML(single)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "<xml><v>a</v></xml>" {
		t.Fatalf("single: got %s, want <xml><v>a</v></xml>", got)
	}

	array := &XMLOptions{Type: "array", RowCount: 2, Fields: fields}
	got, err = New(11).XML(array)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "<xml><record><v>1</v></record><record><v>2</v></record></xml>" {
		t.Fatalf("array: got %s, want <xml> with 2 records", got)
	}

	// The options that were passed in are not changed
	if single.RootElement != "" || single.RecordElement != "" || array.RootElement != "" || array.RecordElement != "" {
		t.Fatalf("the options were changed: %+v %+v", single, array)
	}
}
