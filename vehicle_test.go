package gofakeit

import "fmt"

func ExampleCarModel() {
	Seed(11)
	fmt.Println(CarModel())
	// Output: XK8 CONVERTIBLE
}

func ExampleCarMaker() {
	Seed(11)
	fmt.Println(CarMaker())
	// Output: Nissan
}

func ExampleTransmissionGearType() {
	Seed(11)
	fmt.Println(TransmissionGearType())
	// Output: Manual
}

func ExampleFuelType() {
	Seed(11)
	fmt.Println(FuelType())
	// Output: CNG
}

func ExampleVehicleType() {
	Seed(11)
	fmt.Println(VehicleType())
	// Output: Passenger car mini
}