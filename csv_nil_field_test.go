package gofakeit

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestCSVNilFieldValues(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value any
		want  string
	}{
		{"nil", nil, fmt.Sprint(nil)},
		{"json_null", []byte("null"), fmt.Sprint(nil)},
		{"nil_pointer", (*int)(nil), "null"},
		{"nil_map", map[string]int(nil), "null"},
		{"nil_slice", []int(nil), "null"},
		{"nested", map[string]int{"count": 2}, `{"count":2}`},
		{"text", "hello,world", "hello,world"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			name := "csv_nil_field_values_" + tc.name
			AddFuncLookup(name, Info{Generate: func(*Faker, *MapParams, *Info) (any, error) { return tc.value, nil }})
			t.Cleanup(func() { RemoveFuncLookup(name) })
			for _, delimiter := range []string{",", ";", "tab"} {
				t.Run(delimiter, func(t *testing.T) {
					defer func() {
						if v := recover(); v != nil {
							t.Fatalf("CSV panicked for %s: %v", tc.name, v)
						}
					}()
					output, err := New(1).CSV(&CSVOptions{Delimiter: delimiter, RowCount: 2, Fields: []Field{{Name: "id", Function: "autoincrement"}, {Name: "value", Function: name}}})
					if err != nil {
						t.Fatal(err)
					}
					r := csv.NewReader(strings.NewReader(string(output)))
					if delimiter == "tab" {
						r.Comma = '\t'
					} else {
						r.Comma = []rune(delimiter)[0]
					}
					records, err := r.ReadAll()
					if err != nil {
						t.Fatal(err)
					}
					want := [][]string{{"id", "value"}, {"1", tc.want}, {"2", tc.want}}
					if !reflect.DeepEqual(records, want) {
						t.Fatalf("records = %#v, want %#v", records, want)
					}
				})
			}
		})
	}
}
