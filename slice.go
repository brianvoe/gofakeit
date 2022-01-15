package gofakeit

import (
	"math/rand"
	"reflect"
)

// Slice fills built-in types and exported fields of a struct with random data.
func Slice(v interface{}) { sliceFunc(globalFaker.Rand, v) }

// Slice fills built-in types and exported fields of a struct with random data.
func (f *Faker) Slice(v interface{}) { sliceFunc(f.Rand, v) }

func sliceFunc(ra *rand.Rand, v interface{}) {
	r(ra, reflect.TypeOf(v), reflect.ValueOf(v), "", -1)
}
