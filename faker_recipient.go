package gofakeit

import (
	"errors"
	"fmt"
	"reflect"
)

// Recipient interface can be implemented by wrapper types to define which field
// will be set with faked values.
type Recipient interface {
	// Receptor method implementation MUST have pointer receiver,
	// e.g. func (r *R) Receptor().
	//
	// Returned value MUST have kind reflect.Pointer, or it will not be settable
	Receptor() reflect.Value
}

func isRecipient(t reflect.Type) bool {
	receiverType := reflect.TypeOf((*Recipient)(nil)).Elem()
	return t.Implements(receiverType) || reflect.PointerTo(t).Implements(receiverType)
}

func callReceptor(t reflect.Type, v reflect.Value) (reflect.Value, error) {
	var receptor reflect.Value
	if t.Kind() == reflect.Pointer {
		// Pointer elements need to be initialized first
		ptrValue := reflect.New(t.Elem())
		v.Set(ptrValue)
		receptor = ptrValue.Interface().(Recipient).Receptor()
	} else {
		receptor = v.Addr().Interface().(Recipient).Receptor()
	}
	if receptor.Kind() != reflect.Pointer {
		return reflect.Value{}, errors.New("method Receiver MUST return a pointer when implementing Receiver interface")
	}
	if !receptor.Elem().CanSet() {
		return reflect.Value{}, fmt.Errorf("Receptor for type %s is not settable", t)
	}
	// Set zero to overwrite default generation.
	// This is needed to preserve faked values generation for other potential non-receptor fields.
	receptor.Elem().SetZero()
	return receptor, nil
}
