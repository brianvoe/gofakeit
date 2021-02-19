package gofakeit

import (
	"fmt"
)

func ExampleSlice() {
	Seed(11)

	var S []string
	Slice(&S)

	I := make([]int8, 3)
	Slice(&I)

	fmt.Println(S)
	fmt.Println(I)
	// Output: [RMaR]
	// [21 -41 72]
}

func ExampleFaker_Slice() {
	f := New(11)

	var S []string
	f.Slice(&S)

	I := make([]int8, 3)
	f.Slice(&I)

	fmt.Println(S)
	fmt.Println(I)
	// Output: [RMaR]
	// [21 -41 72]
}

func ExampleSlice_struct() {
	Seed(11)

	type Basic struct {
		S string `fake:"{firstname}"`
		I int
		F float32
	}

	var B []Basic
	Slice(&B)

	fmt.Println(B)
	// Output: [{Marcel -1697368647228132669 1.9343967e+38}]
}

func ExampleFaker_Slice_struct() {
	f := New(11)

	type Basic struct {
		S string `fake:"{firstname}"`
		I int
		F float32
	}

	var B []Basic
	f.Slice(&B)

	fmt.Println(B)
	// Output: [{Marcel -1697368647228132669 1.9343967e+38}]
}
