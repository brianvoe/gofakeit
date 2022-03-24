package gofakeit

import (
	"errors"
	"fmt"
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
func Struct(v interface{}) error { return structFunc(globalFaker.Rand, v) }

// Struct fills in exported fields of a struct with random data
// based on the value of `fake` tag of exported fields.
// Use `fake:"skip"` to explicitly skip an element.
// All built-in types are supported, with templating support
// for string types.
func (f *Faker) Struct(v interface{}) error { return structFunc(f.Rand, v) }

type circleCheck []reflect.Type

func (c circleCheck) IsCircle(t reflect.Type) bool {
	for _, ct := range c {
		if ct == t {
			return true
		}
	}
	return false
}

func structFunc(ra *rand.Rand, v interface{}) error {
	return r(nil, ra, reflect.TypeOf(v), reflect.ValueOf(v), "", 0)
}

func r(cc circleCheck, ra *rand.Rand, t reflect.Type, v reflect.Value, tag string, size int) error {
	switch t.Kind() {
	case reflect.Ptr:
		return rPointer(cc, ra, t, v, tag, size)
	case reflect.Struct:
		if cc.IsCircle(t) {
			return nil
		}
		cc = append(cc, t)
		return rStruct(cc, ra, t, v, tag)
	case reflect.String:
		return rString(ra, v, tag)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rUint(ra, t, v, tag)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rInt(ra, t, v, tag)
	case reflect.Float32, reflect.Float64:
		return rFloat(ra, t, v, tag)
	case reflect.Bool:
		return rBool(ra, v, tag)
	case reflect.Array, reflect.Slice:
		return rSlice(cc, ra, t, v, tag, size)
	case reflect.Map:
		return rMap(cc, ra, t, v, tag, size)
	}

	return nil
}

func rStruct(cc circleCheck, ra *rand.Rand, t reflect.Type, v reflect.Value, tag string) error {
	// If tag is set lets try to set the struct values from the tag response
	if tag != "" {
		fName, fParams := parseNameAndParamsFromTag(tag)
		// Check to see if its a replaceable lookup function
		if info := GetFuncLookup(fName); info != nil {
			// Parse map params
			mapParams := parseMapParams(info, fParams)

			// Call function
			fValue, err := info.Generate(ra, mapParams, info)
			if err != nil {
				return err
			}

			// Create new element of expected type
			field := reflect.New(reflect.TypeOf(fValue))
			field.Elem().Set(reflect.ValueOf(fValue))

			// Check if element is pointer if so
			// grab the underlyning value before setting
			fieldElem := field.Elem()
			if fieldElem.Kind() == reflect.Ptr {
				v.Set(fieldElem.Elem())
			} else {
				v.Set(fieldElem)
			}

			// If a function is called to set the struct
			// stop from going through sub fields
			return nil
		}
	}

	n := t.NumField()
	for i := 0; i < n; i++ {
		elementT := t.Field(i)
		elementV := v.Field(i)
		fakeTag, ok := elementT.Tag.Lookup("fake")

		// Check whether or not to skip this field
		if ok && fakeTag == "skip" {
			// Do nothing, skip it
			continue
		}

		// Check to make sure you can set it or that its an embeded(anonymous) field
		if elementV.CanSet() || elementT.Anonymous {
			// Check if reflect type is of values we can specifically set
			switch elementT.Type.String() {
			case "time.Time":
				err := rTime(ra, elementT, elementV, fakeTag)
				if err != nil {
					return err
				}
				continue
			}

			// Check if fakesize is set
			size := -1 // Set to -1 to indicate fakesize was not set
			fs, ok := elementT.Tag.Lookup("fakesize")
			if ok {
				var err error
				size, err = strconv.Atoi(fs)
				if err != nil {
					return err
				}
			}
			err := r(cc, ra, elementT.Type, elementV, fakeTag, size)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func rPointer(cc circleCheck, ra *rand.Rand, t reflect.Type, v reflect.Value, tag string, size int) error {
	elemT := t.Elem()
	if v.IsNil() {
		nv := reflect.New(elemT)
		err := r(cc, ra, elemT, nv.Elem(), tag, size)
		if err != nil {
			return err
		}
		v.Set(nv)
	} else {
		err := r(cc, ra, elemT, v.Elem(), tag, size)
		if err != nil {
			return err
		}
	}

	return nil
}

func rSlice(cc circleCheck, ra *rand.Rand, t reflect.Type, v reflect.Value, tag string, size int) error {
	// If you cant even set it dont even try
	if !v.CanSet() {
		return errors.New("cannot set slice")
	}

	// Grab original size to use if needed for sub arrays
	ogSize := size

	// If the value has a len and is less than the size
	// use that instead of the requested size
	elemLen := v.Len()
	if elemLen == 0 && size == -1 {
		size = number(ra, 1, 10)
	} else if elemLen != 0 && (size == -1 || elemLen < size) {
		size = elemLen
	}

	// Get the element type
	elemT := t.Elem()

	// Loop through the elements length and set based upon the index
	for i := 0; i < size; i++ {
		nv := reflect.New(elemT)
		err := r(cc, ra, elemT, nv.Elem(), tag, ogSize)
		if err != nil {
			return err
		}

		// If values are already set fill them up, otherwise append
		if elemLen != 0 {
			v.Index(i).Set(reflect.Indirect(nv))
		} else {
			v.Set(reflect.Append(reflect.Indirect(v), reflect.Indirect(nv)))
		}
	}

	return nil
}

func rMap(cc circleCheck, ra *rand.Rand, t reflect.Type, v reflect.Value, tag string, size int) error {
	// If you cant even set it dont even try
	if !v.CanSet() {
		return errors.New("cannot set slice")
	}

	if tag != "" {
		// If tag is set lets try to set the struct values from the tag response
		fName, fParams := parseNameAndParamsFromTag(tag)
		// Check to see if its a replaceable lookup function
		if info := GetFuncLookup(fName); info != nil {
			// Parse map params
			mapParams := parseMapParams(info, fParams)

			// Call function
			fValue, err := info.Generate(ra, mapParams, info)
			if err != nil {
				return err
			} else if reflect.TypeOf(fValue) != t {
				return fmt.Errorf("expected value of type: %s but got value of: %s", t, reflect.TypeOf(fValue))
			}

			// Create new element of expected type
			field := reflect.New(reflect.TypeOf(fValue))
			field.Elem().Set(reflect.ValueOf(fValue))

			// Check if element is pointer if so
			// grab the underlyning value before setting
			fieldElem := field.Elem()
			if fieldElem.Kind() == reflect.Ptr {
				v.Set(fieldElem.Elem())
			} else {
				v.Set(fieldElem)
			}

			// If a function is called to set the struct
			// stop from going through sub fields
			return nil
		}
	}

	// Has no tag or func lookup failed generate map with random data

	// Set a size
	newSize := size
	if newSize == -1 {
		newSize = number(ra, 1, 10)
	}

	// Create new map based upon map key value type
	mapType := reflect.MapOf(t.Key(), t.Elem())
	newMap := reflect.MakeMap(mapType)

	for i := 0; i < newSize; i++ {
		// Create new key
		mapIndex := reflect.New(t.Key())
		err := r(cc, ra, t.Key(), mapIndex.Elem(), "", -1)
		if err != nil {
			return err
		}

		// Create new value
		mapValue := reflect.New(t.Elem())
		err = r(cc, ra, t.Elem(), mapValue.Elem(), "", -1)
		if err != nil {
			return err
		}

		newMap.SetMapIndex(mapIndex.Elem(), mapValue.Elem())
	}

	// Set newMap into struct field
	if t.Kind() == reflect.Ptr {
		v.Set(newMap.Elem())
	} else {
		v.Set(newMap)
	}

	return nil
}

func rString(ra *rand.Rand, v reflect.Value, tag string) error {
	if tag != "" {
		v.SetString(generate(ra, tag))
	} else {
		v.SetString(generate(ra, strings.Repeat("?", number(ra, 4, 10))))
	}

	return nil
}

func rInt(ra *rand.Rand, t reflect.Type, v reflect.Value, tag string) error {
	if tag != "" {
		i, err := strconv.ParseInt(generate(ra, tag), 10, 64)
		if err != nil {
			return err
		}

		v.SetInt(i)
		return nil
	}

	// If no tag or error converting to int, set with random value
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

	return nil
}

func rUint(ra *rand.Rand, t reflect.Type, v reflect.Value, tag string) error {
	if tag != "" {
		u, err := strconv.ParseUint(generate(ra, tag), 10, 64)
		if err != nil {
			return err
		}

		v.SetUint(u)
		return nil
	}

	// If no tag or error converting to uint, set with random value
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

	return nil
}

func rFloat(ra *rand.Rand, t reflect.Type, v reflect.Value, tag string) error {
	if tag != "" {
		f, err := strconv.ParseFloat(generate(ra, tag), 64)
		if err != nil {
			return err
		}

		v.SetFloat(f)
		return nil
	}

	// If no tag or error converting to float, set with random value
	switch t.Kind() {
	case reflect.Float64:
		v.SetFloat(float64Func(ra))
	case reflect.Float32:
		v.SetFloat(float64(float32Func(ra)))
	}

	return nil
}

func rBool(ra *rand.Rand, v reflect.Value, tag string) error {
	if tag != "" {
		b, err := strconv.ParseBool(generate(ra, tag))
		if err != nil {
			return err
		}

		v.SetBool(b)
		return nil
	}

	// If no tag or error converting to boolean, set with random value
	v.SetBool(boolFunc(ra))

	return nil
}

// rTime will set a time.Time field the best it can from either the default date tag or from the generate tag
func rTime(ra *rand.Rand, t reflect.StructField, v reflect.Value, tag string) error {
	if tag != "" {
		// Generate time
		timeOutput := generate(ra, tag)

		// Check to see if they are passing in a format	to parse the time
		timeFormat, timeFormatOK := t.Tag.Lookup("format")
		if timeFormatOK {
			timeFormat = javaDateFormatToGolangDateFormat(timeFormat)
		} else {
			// If tag == "{date}" use time.RFC3339
			// They are attempting to use the default date lookup
			if tag == "{date}" {
				timeFormat = time.RFC3339
			} else {
				// Default format of time.Now().String()
				timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"
			}
		}

		// If output is larger than format cut the output
		// This helps us avoid errors from time.Parse
		if len(timeOutput) > len(timeFormat) {
			timeOutput = timeOutput[:len(timeFormat)]
		}

		// Attempt to parse the time
		timeStruct, err := time.Parse(timeFormat, timeOutput)
		if err != nil {
			return err
		}

		v.Set(reflect.ValueOf(timeStruct))
		return nil
	}

	v.Set(reflect.ValueOf(date(ra)))
	return nil
}
