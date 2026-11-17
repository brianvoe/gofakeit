package gofakeit

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"testing"
)

type failingJSONValue struct {
	err error
}

func (v failingJSONValue) MarshalJSON() ([]byte, error) {
	return nil, v.err
}

func TestJSONMarshalErrors(t *testing.T) {
	sentinel := errors.New("custom JSON marshal failure")
	for _, tc := range []struct {
		name  string
		value any
		check func(*testing.T, error)
	}{
		{"unsupported_type", make(chan int), func(t *testing.T, err error) {
			var target *json.UnsupportedTypeError
			if !errors.As(err, &target) {
				t.Fatalf("error = %v, want UnsupportedTypeError", err)
			}
		}},
		{"unsupported_value", math.Inf(1), func(t *testing.T, err error) {
			var target *json.UnsupportedValueError
			if !errors.As(err, &target) {
				t.Fatalf("error = %v, want UnsupportedValueError", err)
			}
		}},
		{"custom_marshaler", failingJSONValue{err: sentinel}, func(t *testing.T, err error) {
			if !errors.Is(err, sentinel) {
				t.Fatalf("error = %v, want custom marshal failure", err)
			}
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			name := "json_marshal_errors_" + tc.name
			AddFuncLookup(name, Info{Generate: func(*Faker, *MapParams, *Info) (any, error) { return tc.value, nil }})
			t.Cleanup(func() { RemoveFuncLookup(name) })
			for _, typ := range []string{"object", "array"} {
				for _, indent := range []bool{false, true} {
					t.Run(fmt.Sprintf("%s/indent_%t", typ, indent), func(t *testing.T) {
						output, err := New(1).JSON(&JSONOptions{Type: typ, RowCount: 2, Indent: indent, Fields: []Field{{Name: "value", Function: name}}})
						tc.check(t, err)
						if len(output) != 0 {
							t.Fatalf("output = %q, want no serialized output after error", output)
						}
					})
				}
			}
		})
	}
}

func TestJSONMarshalValidValues(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value any
		want  string
	}{
		{"null", nil, `null`},
		{"nested", map[string]any{"message": "ok", "values": []int{1, 2}}, `{"message":"ok","values":[1,2]}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			name := "json_marshal_valid_" + tc.name
			AddFuncLookup(name, Info{Generate: func(*Faker, *MapParams, *Info) (any, error) { return tc.value, nil }})
			t.Cleanup(func() { RemoveFuncLookup(name) })
			for _, typ := range []string{"object", "array"} {
				for _, indent := range []bool{false, true} {
					t.Run(fmt.Sprintf("%s/indent_%t", typ, indent), func(t *testing.T) {
						output, err := New(1).JSON(&JSONOptions{Type: typ, RowCount: 2, Indent: indent, Fields: []Field{{Name: "value", Function: name}}})
						if err != nil {
							t.Fatal(err)
						}
						var compact any
						if err := json.Unmarshal(output, &compact); err != nil {
							t.Fatal(err)
						}
						got, err := json.Marshal(compact)
						if err != nil {
							t.Fatal(err)
						}
						want := `{"value":` + tc.want + `}`
						if typ == "array" {
							want = "[" + want + "," + want + "]"
						}
						if string(got) != want {
							t.Fatalf("output = %s, want %s", got, want)
						}
					})
				}
			}
		})
	}
}
