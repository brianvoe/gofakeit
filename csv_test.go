package gofakeit

import (
	"fmt"
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
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// id,first_name,last_name,password
	// 1,Markus,Moen,Dc0VYXjkWABx
	// 2,Osborne,Hilll,XPJ9OVNbs5lm
}

func TestCSVLookup(t *testing.T) {
	info := GetFuncLookup("csv")

	m := map[string][]string{
		"rowcount": {"10"},
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

func BenchmarkCSVLookup100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("csv")
		m := map[string][]string{
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

func BenchmarkCSVLookup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("csv")
		m := map[string][]string{
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

func BenchmarkCSVLookup10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("csv")
		m := map[string][]string{
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

func BenchmarkCSVLookup100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("csv")
		m := map[string][]string{
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
