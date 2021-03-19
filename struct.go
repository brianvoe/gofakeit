package gofakeit

import (
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Struct fills in exported fields of a struct with random data
// based on the value of `fake` tag of exported fields.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
func Struct(v interface{}) { structFunc(globalFaker.Rand, v) }

// Struct fills in exported fields of a struct with random data
// based on the value of `fake` tag of exported fields.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
func (f *Faker) Struct(v interface{}) { structFunc(f.Rand, v) }

func structFunc(ra *rand.Rand, v interface{}) {
	r(ra, reflect.TypeOf(v), reflect.ValueOf(v), "", 0)
}

func r(ra *rand.Rand, t reflect.Type, v reflect.Value, function string, size int) {
	switch t.Kind() {
	case reflect.Ptr:
		rPointer(ra, t, v, function, size)
	case reflect.Struct:
		rStruct(ra, t, v, function)
	case reflect.String:
		rString(ra, t, v, function)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		rUint(ra, t, v, function)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rInt(ra, t, v, function)
	case reflect.Float32, reflect.Float64:
		rFloat(ra, t, v, function)
	case reflect.Bool:
		rBool(ra, t, v, function)
	case reflect.Array, reflect.Slice:
		rSlice(ra, t, v, function, size)
	}
}

func rStruct(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	n := t.NumField()
	for i := 0; i < n; i++ {
		elementT := t.Field(i)
		elementV := v.Field(i)
		fakeTag, ok := elementT.Tag.Lookup("fake")
		if ok && fakeTag == "skip" {
			// Do nothing, skip it
		} else if elementV.CanSet() || elementT.Anonymous {
			// Check if reflect type is of values we can specifically set
			switch elementT.Type.String() {
			case "time.Time":
				rTime(ra, elementT, elementV, fakeTag)
				continue
			}

			// Check if fakesize is set
			size := -1 // Set to -1 to indicate fakesize was not set
			fs, ok := elementT.Tag.Lookup("fakesize")
			if ok {
				var err error
				size, err = strconv.Atoi(fs)
				if err != nil {
					size = number(ra, 1, 10)
				}
			}
			r(ra, elementT.Type, elementV, fakeTag, size)
		}
	}
}

func rPointer(ra *rand.Rand, t reflect.Type, v reflect.Value, function string, size int) {
	elemT := t.Elem()
	if v.IsNil() {
		nv := reflect.New(elemT)
		r(ra, elemT, nv.Elem(), function, size)
		v.Set(nv)
	} else {
		r(ra, elemT, v.Elem(), function, size)
	}
}

func rSlice(ra *rand.Rand, t reflect.Type, v reflect.Value, function string, size int) {
	// If you cant even set it dont even try
	if !v.CanSet() {
		return
	}

	// Grab original size to use if needed for sub arrays
	ogSize := size

	// If the value has a cap and is less than the size
	// use that instead of the requested size
	elemCap := v.Cap()
	if elemCap == 0 && size == -1 {
		size = number(ra, 1, 10)
	} else if elemCap != 0 && (size == -1 || elemCap < size) {
		size = elemCap
	}

	// Get the element type
	elemT := t.Elem()

	// If values are already set fill them up, otherwise append
	if v.Len() != 0 {
		// Loop through the elements length and set based upon the index
		for i := 0; i < size; i++ {
			nv := reflect.New(elemT)
			r(ra, elemT, nv.Elem(), function, ogSize)
			v.Index(i).Set(reflect.Indirect(nv))
		}
	} else {
		// Loop through the size and append and set
		for i := 0; i < size; i++ {
			nv := reflect.New(elemT)
			r(ra, elemT, nv.Elem(), function, ogSize)
			v.Set(reflect.Append(reflect.Indirect(v), reflect.Indirect(nv)))
		}
	}
}

func rString(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		v.SetString(generate(ra, function))
	} else {
		v.SetString(generate(ra, strings.Repeat("?", number(ra, 4, 10))))
	}
}

func rInt(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		i, err := strconv.ParseInt(generate(ra, function), 10, 64)
		if err == nil {
			v.SetInt(i)
			return
		}
	}

	// If no function or error converting to int, set with random value
	switch t.Kind() {
	case reflect.Int:
		v.SetInt(int64Func(ra))
	case reflect.Int8:
		v.SetInt(int64(int8Func(ra)))
	case reflect.Int16:
		v.SetInt(int64(int16Func(ra)))
	case reflect.Int32:
		v.SetInt(int64(int32Func(ra)))
	case reflect.Int64:
		v.SetInt(int64Func(ra))
	}
}

func rUint(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		u, err := strconv.ParseUint(generate(ra, function), 10, 64)
		if err == nil {
			v.SetUint(u)
			return
		}
	}

	// If no function or error converting to uint, set with random value
	switch t.Kind() {
	case reflect.Uint:
		v.SetUint(uint64Func(ra))
	case reflect.Uint8:
		v.SetUint(uint64(uint8Func(ra)))
	case reflect.Uint16:
		v.SetUint(uint64(uint16Func(ra)))
	case reflect.Uint32:
		v.SetUint(uint64(uint32Func(ra)))
	case reflect.Uint64:
		v.SetUint(uint64Func(ra))
	}
}

func rFloat(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		f, err := strconv.ParseFloat(generate(ra, function), 64)
		if err == nil {
			v.SetFloat(f)
			return
		}
	}

	// If no function or error converting to float, set with random value
	switch t.Kind() {
	case reflect.Float64:
		v.SetFloat(float64Func(ra))
	case reflect.Float32:
		v.SetFloat(float64(float32Func(ra)))
	}
}

func rBool(ra *rand.Rand, t reflect.Type, v reflect.Value, function string) {
	if function != "" {
		b, err := strconv.ParseBool(generate(ra, function))
		if err == nil {
			v.SetBool(b)
			return
		}
	}

	// If no function or error converting to boolean, set with random value
	v.SetBool(boolFunc(ra))
}

// rTime will set a time.Time field the best it can from either the default date function or from the generate function
func rTime(ra *rand.Rand, t reflect.StructField, v reflect.Value, tag string) {
	if tag != "" {
		timeFormat, timeFormatOK := t.Tag.Lookup("format")
		// If no format passed in use the most popular RFC3339
		if !timeFormatOK {
			timeFormat = time.RFC3339
		}
		timeOutput := generate(ra, tag)
		timeStruct, err := time.Parse(timeFormat, timeOutput)
		if err == nil {
			v.Set(reflect.ValueOf(timeStruct))
		}
	} else {
		v.Set(reflect.ValueOf(date(ra)))
	}
}
