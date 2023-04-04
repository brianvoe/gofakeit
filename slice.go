package gofakeit

import (
	"reflect"
)

// Slice fills built-in types and exported fields of a struct with random data.
func Slice(v interface{}) { sliceFunc(globalFaker, v) }

// Slice fills built-in types and exported fields of a struct with random data.
func (f *Faker) Slice(v interface{}) { sliceFunc(f, v) }

func sliceFunc(f *Faker, v interface{}) {
	r(f, reflect.TypeOf(v), reflect.ValueOf(v), "", -1)
}
