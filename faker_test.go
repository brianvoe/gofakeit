package gofakeit

import "fmt"

func Example() {
	Seed(11)
	fmt.Println("Name:", Name())
	fmt.Println("Email:", Email())
	fmt.Println("Phone:", Phone())
	fmt.Println("Address:", Address().Address)
	// Output:
	// Name: Markus Moen
	// Email: alaynawuckert@kozey.biz
	// Phone: (570)245-7485
	// Address: 75776 Lake View land, Rueckerstad, New Hampshire 82250-2868
}
