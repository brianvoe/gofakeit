package gofakeit

import "fmt"

func Example() {
	fmt.Println("Name:", Name())
	fmt.Println("Email:", Email())
	fmt.Println("Phone:", Phone())
	fmt.Println("Address:", Address())
	// Output: Name: Lura Lockman
	// Email: sylvanmraz@murphy.com
	// Phone: (457)485-3675
	// Address: 6037 Port Groves stad, Noeville, North Carolina 50286-8748
}
