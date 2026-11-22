package gofakeit

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestXMLNumericFieldsKeepNames(t *testing.T) {
	for _, function := range []string{"int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64"} {
		for _, mode := range []string{"single", "array"} {
			for _, global := range []bool{false, true} {
				t.Run(fmt.Sprintf("%s/%s/global=%v", function, mode, global), func(t *testing.T) {
					count := 1
					if mode == "array" {
						count = 2
					}
					info := GetFuncLookup(function)
					want := make([]string, count)
					expected := New(27)
					for i := range want {
						params := MapParams{}
						value, err := info.Generate(expected, &params, info)
						if err != nil {
							t.Fatal(err)
						}
						want[i] = fmt.Sprint(value)
					}
					options := &XMLOptions{
						Type: mode, RootElement: "sample", RecordElement: "record", RowCount: count,
						Fields: []Field{{Name: "reading", Function: function}},
					}
					var output []byte
					var err error
					if global {
						Seed(27)
						output, err = XML(options)
					} else {
						output, err = New(27).XML(options)
					}
					if err != nil {
						t.Fatal(err)
					}
					type record struct {
						Value string `xml:"reading"`
					}
					var document struct {
						record
						Records []record `xml:"record"`
					}
					if err := xml.Unmarshal(output, &document); err != nil {
						t.Fatalf("invalid XML: %v; %s", err, output)
					}
					got := []string{document.Value}
					if mode == "array" {
						got = nil
						for _, row := range document.Records {
							got = append(got, row.Value)
						}
					}
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("reading values = %q, want %q; XML: %s", got, want, output)
					}
				})
			}
		}
	}
}

func TestXMLShuffleIntsKeepsNumbers(t *testing.T) {
	for _, mode := range []string{"single", "array"} {
		for _, global := range []bool{false, true} {
			t.Run(fmt.Sprintf("%s/global=%v", mode, global), func(t *testing.T) {
				options := &XMLOptions{
					Type: mode, RootElement: "sample", RecordElement: "record", RowCount: 2,
					Fields: []Field{{Name: "numbers", Function: "shuffleints", Params: MapParams{"ints": {"7", "19", "-23"}}}},
				}
				var output []byte
				var err error
				if global {
					Seed(27)
					output, err = XML(options)
				} else {
					output, err = New(27).XML(options)
				}
				if err != nil {
					t.Fatal(err)
				}
				type record struct {
					Values []int `xml:"numbers>value"`
				}
				var document struct {
					record
					Records []record `xml:"record"`
				}
				if err := xml.Unmarshal(output, &document); err != nil {
					t.Fatalf("invalid numeric XML: %v; %s", err, output)
				}
				rows := []record{document.record}
				if mode == "array" {
					rows = document.Records
				}
				wantRows := 1
				if mode == "array" {
					wantRows = 2
				}
				if len(rows) != wantRows {
					t.Fatalf("got %d records, want %d", len(rows), wantRows)
				}
				for _, row := range rows {
					sort.Ints(row.Values)
					if !reflect.DeepEqual(row.Values, []int{-23, 7, 19}) {
						t.Fatalf("numbers = %v, want -23, 7, 19; XML: %s", row.Values, output)
					}
				}
			})
		}
	}
}

func TestXMLSliceValuesAreDataNotReflectionDescriptions(t *testing.T) {
	type integer int32
	integers := []int{-23, 0, 19}
	cases := []struct {
		name  string
		value any
		want  []string
	}{
		{"integers", integers, []string{"-23", "0", "19"}},
		{"unsigned", []uint16{0, 57136, 65535}, []string{"0", "57136", "65535"}},
		{"floats", []float64{1.25, -2.5, 0}, []string{"1.25", "-2.5", "0"}},
		{"booleans", []bool{true, false}, []string{"true", "false"}},
		{"strings", []string{"<&>", "quote\"'", "plain"}, []string{"<&>", "quote\"'", "plain"}},
		{"interfaces", []any{7, "<&>", false}, []string{"7", "<&>", "false"}},
		{"bytes", []byte{0, 1, 255}, []string{"0", "1", "255"}},
		{"named_integer", []integer{-9, 4}, []string{"-9", "4"}},
		{"pointer_to_slice", &integers, []string{"-23", "0", "19"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			const function = "xml_slice_data_test"
			AddFuncLookup(function, Info{Generate: func(*Faker, *MapParams, *Info) (any, error) {
				return tc.value, nil
			}})
			t.Cleanup(func() { RemoveFuncLookup(function) })
			output, err := New(27).XML(&XMLOptions{
				Type: "single", RootElement: "sample",
				Fields: []Field{{Name: "sequence", Function: function}},
			})
			if err != nil {
				t.Fatal(err)
			}
			var document struct {
				Values []string `xml:"sequence>value"`
			}
			if err := xml.Unmarshal(output, &document); err != nil {
				t.Fatalf("invalid XML: %v; %s", err, output)
			}
			if !reflect.DeepEqual(document.Values, tc.want) {
				t.Fatalf("sequence values = %q, want %q; XML: %s", document.Values, tc.want, output)
			}
		})
	}
}
