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

	// Output: [RMaRxHkiJ PtapWYJdn MKgtlxwnq qclaYkWw oRLOPxLIok qanPAKaXS]
	// [-88 -101 60]
}

func ExampleFaker_Slice() {
	f := New(11)

	var S []string
	f.Slice(&S)

	I := make([]int8, 3)
	f.Slice(&I)

	fmt.Println(S)
	fmt.Println(I)

	// Output: [RMaRxHkiJ PtapWYJdn MKgtlxwnq qclaYkWw oRLOPxLIok qanPAKaXS]
	// [-88 -101 60]
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
	// Output: [{Marcel -652938557 1.9343967e+38} {Lura -101175613 2.670526e+38} {Lucinda 1412367049 1.0427679e+38} {Santino 1631045574 2.2573786e+38} {Dawn -924758850 4.249854e+37} {Alice -937677379 1.0573345e+38}]
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
	// Output: [{Marcel -652938557 1.9343967e+38} {Lura -101175613 2.670526e+38} {Lucinda 1412367049 1.0427679e+38} {Santino 1631045574 2.2573786e+38} {Dawn -924758850 4.249854e+37} {Alice -937677379 1.0573345e+38}]
}
