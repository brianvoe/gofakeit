package gofakeit

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
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
	// 1,Priscilla,Thornton,F395YfK9tHs9
	// 2,Emmett,Black,PL0G37o1pY0p
	// 3,Liza,Vincent,q11fR4MwsM4N
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
	// 1,Priscilla,Thornton,F395YfK9tHs9
	// 2,Emmett,Black,PL0G37o1pY0p
	// 3,Liza,Vincent,q11fR4MwsM4N
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

func TestCSVErrorAndTimeValues(t *testing.T) {
	// An error used to be written as {} and a time with the quotes of json around it
	AddFuncLookup("csv_error_value", Info{Generate: func(*Faker, *MapParams, *Info) (any, error) { return errors.New("not found"), nil }})
	AddFuncLookup("csv_time_value", Info{Generate: func(*Faker, *MapParams, *Info) (any, error) {
		return time.Date(2026, 11, 22, 10, 30, 0, 0, time.UTC), nil
	}})
	t.Cleanup(func() {
		RemoveFuncLookup("csv_error_value")
		RemoveFuncLookup("csv_time_value")
	})

	got, err := New(11).CSV(&CSVOptions{RowCount: 1, Fields: []Field{{Name: "e", Function: "csv_error_value"}, {Name: "t", Function: "csv_time_value"}}})
	if err != nil {
		t.Fatal(err)
	}
	if want := "e,t\nnot found,2026-11-22T10:30:00Z\n"; string(got) != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
