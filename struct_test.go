package gofakeit

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Basic struct {
	s string
	S string
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
}

type BuiltIn struct {
	Int     *int
	Float32 *float32
	Float64 *float32
	Bool    *bool
}

type Template struct {
	Number *string `fake:"#"`
	Name   *string `fake:"{person.first}"`
}

func TestStructBasic(t *testing.T) {
	t.Parallel()
	var basic Basic
	Struct(&basic)
	assert.Empty(t, basic.s, "unexported field is not populated")
	assert.NotEmpty(t, basic.S, "exported field is populated")
}

func TestStructNested(t *testing.T) {
	t.Parallel()
	var nested Nested
	Struct(&nested)
	assert.NotEmpty(t, nested.A, "exported string field is populated")
	assert.NotNil(t, nested.B, "exported struct field is populated")
	assert.NotNil(t, nested.B.S, "nested struct fields are populated")
	assert.Nil(t, nested.bar, "unexported nested struct is not populated")
}

func TestStructBuiltInTypes(t *testing.T) {
	t.Parallel()
	var builtIn BuiltIn
	Struct(&builtIn)
	assert.NotNil(t, builtIn.Int, "populated basic int type via pointer")
	assert.NotNil(t, builtIn.Float32, "populated basic float32 type via pointer")
	assert.NotNil(t, builtIn.Float64, "populated basic float64 type via pointer")
	assert.NotNil(t, builtIn.Bool, "populated basic bool type via pointer")
}

func TestStructWithTemplate(t *testing.T) {
	t.Parallel()
	var template Template
	Struct(&template)
	assert.NotNil(t, template.Number, "populated Number based on template")
	assert.NotNil(t, template.Name, "populated Name based on template")
	assert.NotEmpty(t, template.Name, "populated Name based on template")
	_, err := strconv.ParseInt(*template.Number, 10, 64)
	assert.Nil(t, err, "number is populated")
}

func ExampleStruct() {
	Seed(11)
	type Foo struct {
		Bar     string
		Baz     string
		Int     int
		Pointer *int
		Skip    *string `fake:"skip"`
	}
	var f Foo
	Seed(42)
	Struct(&f)
	fmt.Printf("f.Bar:%s\n", f.Bar)
	fmt.Printf("f.Baz:%s\n", f.Baz)
	fmt.Printf("f.Int:%d\n", f.Int)
	fmt.Printf("f.Pointer:%d\n", *f.Pointer)
	fmt.Printf("f.Skip:%v\n", f.Skip)
	// Output: f.Bar:hrukpttuezptneuvunh
	// f.Baz:uksqvgzadxlgghejkmv
	// f.Int:-7825289004089916589
	// f.Pointer:-3438066090944737321
	// f.Skip:<nil>
}
