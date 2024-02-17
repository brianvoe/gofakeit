package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleCSV_array() {
	Seed(11)

	value, err := CSV(&CSVOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: id,first_name,last_name,password
	// 1,Sonny,Stiedemann,8nwf0o3sBXcR
	// 2,Verda,Brakus,3beWLpq75Lua
	// 3,Jules,Cremin,Uu38J14Y8W82
}

func ExampleFaker_CSV_array() {
	f := New(11)

	value, err := f.CSV(&CSVOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: id,first_name,last_name,password
	// 1,Sonny,Stiedemann,8nwf0o3sBXcR
	// 2,Verda,Brakus,3beWLpq75Lua
	// 3,Jules,Cremin,Uu38J14Y8W82
}

func TestCSVLookup(t *testing.T) {
	faker := New(0)

	info := GetFuncLookup("csv")

	m := MapParams{
		"rowcount": {"10"},
		"fields": {
			`{"name":"id","function":"autoincrement"}`,
			`{"name":"first_name","function":"firstname"}`,
			`{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`,
			`{"name":"address","function":"address"}`,
			`{
				"name":"json",
				"function":"json",
				"params":{
					"type":"object",
					"fields":[
						{"name":"id","function":"autoincrement"},
						{"name":"first_name","function":"firstname"},
						{"name":"last_name","function":"lastname"},
						{"name":"password","function":"password","params":{"special":"false","length":"20"}}
					]
				}
			}`,
		},
	}

	output, err := info.Generate(faker, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	value := string(output.([]byte))

	// Check that value has the correct number of rows via new line characters plus 1 for the header
	if strings.Count(value, "\n") != 11 {
		t.Error("Expected 10+1(header row) rows, got", strings.Count(value, "\n")+1)
	}

	// t.Fatal(fmt.Sprintf("%s", value.([]byte)))
}

func TestCSVNoOptions(t *testing.T) {
	Seed(11)

	// if CSVOptions is nil -> get a random CSVOptions
	_, err := CSV(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func BenchmarkCSVLookup100(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("csv")
		m := MapParams{
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
		_, err := info.Generate(faker, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
