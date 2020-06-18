package gofakeit

import (
	"reflect"
	"strconv"
	"strings"
)

// Struct fills in exported elements of a struct with random data
// based on the value of `fake` tag of exported elements.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
func Struct(v interface{}) {
	r(reflect.TypeOf(v), reflect.ValueOf(v), "", 0)
}

func r(t reflect.Type, v reflect.Value, template string, size int) {
	switch t.Kind() {
	case reflect.Ptr:
		rPointer(t, v, template)
	case reflect.Struct:
		rStruct(t, v)
	case reflect.String:
		rString(t, v, template)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rUint(t, v, template)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rInt(t, v, template)
	case reflect.Float32, reflect.Float64:
		rFloat(t, v, template)
	case reflect.Bool:
		rBool(t, v, template)
	case reflect.Array, reflect.Slice:
		rSlice(t, v, template, size)
	}
}

func rStruct(t reflect.Type, v reflect.Value) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		elementT := t.Field(i)
		elementV := v.Field(i)
		t, ok := elementT.Tag.Lookup("fake")
		if ok && t == "skip" {
			// Do nothing, skip it
		} else if elementV.CanSet() {
			// Check if fakesize is set
			size := Number(1, 10)
			fs, ok := elementT.Tag.Lookup("fakesize")
			if ok {
				var err error
				size, err = strconv.Atoi(fs)
				if err != nil {
					size = Number(1, 10)
				}
			}
			r(elementT.Type, elementV, t, size)
		}
	}
}

func rPointer(t reflect.Type, v reflect.Value, template string) {
	elemT := t.Elem()
	if v.IsNil() {
		nv := reflect.New(elemT)
		r(elemT, nv.Elem(), template, 0)
		v.Set(nv)
	} else {
		r(elemT, v.Elem(), template, 0)
	}
}

func rSlice(t reflect.Type, v reflect.Value, template string, size int) {
	elemT := t.Elem()

	if v.CanSet() {
		for i := 0; i < size; i++ {
			nv := reflect.New(elemT)
			r(elemT, nv.Elem(), template, size)
			v.Set(reflect.Append(reflect.Indirect(v), reflect.Indirect(nv)))
		}
	}
}

func rString(t reflect.Type, v reflect.Value, template string) {
	if template != "" {
		v.SetString(Generate(template))
	} else {
		v.SetString(Generate(strings.Repeat("?", Number(4, 10))))
	}
}

func rInt(t reflect.Type, v reflect.Value, template string) {
	if template != "" {
		i, err := strconv.ParseInt(Generate(template), 10, 64)
		if err == nil {
			v.SetInt(i)
			return
		}
	}

	// If no template or error converting to int, set with random value
	switch t.Kind() {
	case reflect.Int:
		v.SetInt(Int64())
	case reflect.Int8:
		v.SetInt(int64(Int8()))
	case reflect.Int16:
		v.SetInt(int64(Int16()))
	case reflect.Int32:
		v.SetInt(int64(Int32()))
	case reflect.Int64:
		v.SetInt(Int64())
	}
}

func rUint(t reflect.Type, v reflect.Value, template string) {
	if template != "" {
		u, err := strconv.ParseUint(Generate(template), 10, 64)
		if err == nil {
			v.SetUint(u)
			return
		}
	}

	// If no template or error converting to uint, set with random value
	switch t.Kind() {
	case reflect.Uint:
		v.SetUint(Uint64())
	case reflect.Uint8:
		v.SetUint(uint64(Uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(Uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(Uint32()))
	case reflect.Uint64:
		v.SetUint(Uint64())
	}
}

func rFloat(t reflect.Type, v reflect.Value, template string) {
	if template != "" {
		f, err := strconv.ParseFloat(Generate(template), 64)
		if err == nil {
			v.SetFloat(f)
			return
		}
	}

	// If no template or error converting to float, set with random value
	switch t.Kind() {
	case reflect.Float64:
		v.SetFloat(Float64())
	case reflect.Float32:
		v.SetFloat(float64(Float32()))
	}
}

func rBool(t reflect.Type, v reflect.Value, template string) {
	if template != "" {
		b, err := strconv.ParseBool(Generate(template))
		if err == nil {
			v.SetBool(b)
			return
		}
	}

	// If no template or error converting to boolean, set with random value
	v.SetBool(Bool())
}
