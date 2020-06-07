package gofakeit

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Struct fills in exported elements of a struct with random data
// based on the value of `fake` tag of exported elements.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
func Struct(v interface{}) {
	r(reflect.TypeOf(v), reflect.ValueOf(v), "")
}

func r(t reflect.Type, v reflect.Value, template string) {
	switch t.Kind() {
	case reflect.Ptr:
		rPointer(t, v, template)
	case reflect.Struct:
		rStruct(t, v)
	case reflect.String:
		rString(template, v)
	case reflect.Uint8:
		v.SetUint(uint64(Uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(Uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(Uint32()))
	case reflect.Uint64:
		v.SetUint(Uint64())
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
	case reflect.Float64:
		v.SetFloat(Float64())
	case reflect.Float32:
		v.SetFloat(float64(Float32()))
	case reflect.Bool:
		v.SetBool(Bool())
	case reflect.Array, reflect.Slice:
		rSlice(t, v, template)
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
			r(elementT.Type, elementV, t)
		}
	}
}

func rPointer(t reflect.Type, v reflect.Value, template string) {
	elemT := t.Elem()
	if v.IsNil() {
		nv := reflect.New(elemT)
		r(elemT, nv.Elem(), template)
		v.Set(nv)
	} else {
		r(elemT, v.Elem(), template)
	}
}

func rSlice(t reflect.Type, v reflect.Value, template string) {
	elemT := t.Elem()

	params := parseTemplate(template)

	if v.CanSet() {
		for i := 0; i < params.size; i++ {
			nv := reflect.New(elemT)
			r(elemT, nv.Elem(), params.template)
			v.Set(reflect.Append(reflect.Indirect(v), reflect.Indirect(nv)))
		}
	}
}

func rString(template string, v reflect.Value) {
	if template != "" {
		v.SetString(Generate(template))
	} else {
		v.SetString(Generate(strings.Repeat("?", Number(4, 10))))
	}
}

type parameters struct {
	size     int
	template string
}

var sizeRE = regexp.MustCompile(`size=(\d+)`)
var functionRE = regexp.MustCompile(`(\{[a-z]+\})`)

func parseTemplate(template string) parameters {
	p := parameters{
		size: randIntRange(1, 10),
	}

	if sizeRE.MatchString(template) {
		if s, err := strconv.Atoi(sizeRE.FindStringSubmatch(template)[1]); err == nil {
			p.size = s
		}
	}

	if functionRE.MatchString(template) {
		p.template = functionRE.FindStringSubmatch(template)[1]
	}

	return p
}
