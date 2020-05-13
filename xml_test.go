package gofakeit

import "fmt"

func ExampleXML_single() {
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
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: jhiuhuih
}

func ExampleXML_array() {
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
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: jhiuhuih
}
