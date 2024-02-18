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

	// Output: [KKbMlbxqu mwwv WVlPmw AeAwVH Khrx DcxFeWk vChMCeKf BwRtnboOE mWluN]
	// [102 -7 -125]
}

func ExampleFaker_Slice() {
	f := New(11)

	var S []string
	f.Slice(&S)

	I := make([]int8, 3)
	f.Slice(&I)

	fmt.Println(S)
	fmt.Println(I)

	// Output: [KKbMlbxqu mwwv WVlPmw AeAwVH Khrx DcxFeWk vChMCeKf BwRtnboOE mWluN]
	// [102 -7 -125]
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

	// Output: [{Russ 3680786209731553973 0.27238095} {Julius 4268594234476337060 0.0051180124} {Kaitlyn 8337306475187377941 0.118576884} {Steve 1365845625386394310 0.27625358} {Tomasa 7952567920265354269 0.648698} {Ernest 7933890822314871011 0.37052673} {Missouri 5542429450337529393 0.36615264} {Tiana 6292602578870227868 0.9382272} {Koby 229639691709918065 0.5914113}]
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

	// Output: [{Russ 3680786209731553973 0.27238095} {Julius 4268594234476337060 0.0051180124} {Kaitlyn 8337306475187377941 0.118576884} {Steve 1365845625386394310 0.27625358} {Tomasa 7952567920265354269 0.648698} {Ernest 7933890822314871011 0.37052673} {Missouri 5542429450337529393 0.36615264} {Tiana 6292602578870227868 0.9382272} {Koby 229639691709918065 0.5914113}]
}
