package gofakeit

import (
	"errors"
	"reflect"
)

// Transformable interface can be implemented to arbitrarily transform the value
// of a type after it was generated.
type Transformable interface {
	// FakeTransform method implementation must have pointer receiver,
	// so that the Transformable can be mutated.
	FakeTransform(faker *Faker) error
}

func isTransformable(t reflect.Type) bool {
	transformableType := reflect.TypeOf((*Transformable)(nil)).Elem()
	return t.Implements(transformableType) || reflect.PointerTo(t).Implements(transformableType)
}

func callTransform(faker *Faker, v reflect.Value) error {
	var (
		f  Transformable
		ok bool
	)
	if v.Type().Kind() == reflect.Pointer {
		f, ok = v.Interface().(Transformable)
	} else {
		f, ok = v.Addr().Interface().(Transformable)
	}
	if !ok {
		return errors.New("not a Transformable type")
	}
	return f.FakeTransform(faker)
}
