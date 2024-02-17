package gofakeit

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

type strTyp string

func (t strTyp) Fake(faker *Faker) (any, error) {
	return faker.FirstName(), nil
}

type strTypPtr string

func (t *strTypPtr) Fake(faker *Faker) (any, error) {
	return strTypPtr("hello test ptr"), nil
}

type testStruct1 struct {
	B string `fake:"{firstname}"`
}

type testStruct2 struct {
	B strTyp
}

func TestIsFakeable(t *testing.T) {
	var t1 testStruct2
	var t2 *testStruct2
	var t3 strTyp
	var t4 *strTyp
	var t5 strTypPtr
	var t6 *strTypPtr

	if isFakeable(reflect.ValueOf(t1).Type()) {
		t.Errorf("expected testStruct2 not to be fakeable")
	}

	if isFakeable(reflect.ValueOf(t2).Type()) {
		t.Errorf("expected *testStruct2 not to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t3).Type()) {
		t.Errorf("expected strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t4).Type()) {
		t.Errorf("expected *strTyp to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t5).Type()) {
		t.Errorf("expected strTypPtr to be fakeable")
	}

	if !isFakeable(reflect.ValueOf(t6).Type()) {
		t.Errorf("expected *strTypPtr to be fakeable")
	}
}

func ExampleFakeable() {
	var t1 testStruct1
	var t2 testStruct1
	var t3 testStruct2
	var t4 testStruct2
	New(314).Struct(&t1)
	New(314).Struct(&t2)
	New(314).Struct(&t3)
	New(314).Struct(&t4)

	fmt.Printf("%#v\n", t1)
	fmt.Printf("%#v\n", t2)
	fmt.Printf("%#v\n", t3)
	fmt.Printf("%#v\n", t4)

	// Output: gofakeit.testStruct1{B:"Colton"}
	// gofakeit.testStruct1{B:"Colton"}
	// gofakeit.testStruct2{B:"Colton"}
	// gofakeit.testStruct2{B:"Colton"}
}

type gammaFloat64 float64

func (gammaFloat64) Fake(f *Faker) (any, error) {
	alpha := 2.0

	// Generate a random value from the Gamma distribution
	var r float64
	for r == 0 {
		u := f.Float64Range(0, 1)
		v := f.Float64Range(0, 1)
		w := u * (1 - u)
		y := math.Sqrt(-2 * math.Log(w) / w)
		x := alpha * (y*v + u - 0.5)
		if x > 0 {
			r = x
		}
	}
	return gammaFloat64(r), nil
}

func ExampleFakeable_gammaFloat64() {
	f1 := New(100)

	// Fakes random values from the Gamma distribution
	var A1 gammaFloat64
	var A2 gammaFloat64
	var A3 gammaFloat64
	f1.Struct(&A1)
	f1.Struct(&A2)
	f1.Struct(&A3)

	fmt.Println(A1)
	fmt.Println(A2)
	fmt.Println(A3)

	// Output: 1.9058272589164647
	// 1.951453943304136
	// 4.336093466276675
}

type poissonInt64 int64

func (poissonInt64) Fake(faker *Faker) (any, error) {
	lambda := 15.0

	// Generate a random value from the Poisson distribution
	var k int64
	var p float64 = 1.0
	var L float64 = math.Exp(-lambda)
	for p > L {
		u := faker.Float64Range(0, 1)
		p *= u
		k++
	}
	return poissonInt64(k - 1), nil
}

type employee struct {
	Name             string `fake:"{firstname} {lastname}"`
	CallCountPerHour poissonInt64
}

func ExampleFakeable_employee() {
	f1 := New(100)

	// Fakes random values from the Gamma distribution
	var A1 employee
	var A2 employee
	var A3 employee
	f1.Struct(&A1)
	f1.Struct(&A2)
	f1.Struct(&A3)

	fmt.Printf("%#v\n", A1)
	fmt.Printf("%#v\n", A2)
	fmt.Printf("%#v\n", A3)

	// Output: gofakeit.employee{Name:"Madelynn Hickle", CallCountPerHour:17}
	// gofakeit.employee{Name:"Brooke Berge", CallCountPerHour:8}
	// gofakeit.employee{Name:"Rosalee Roberts", CallCountPerHour:10}
}
