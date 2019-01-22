package gofakeit

import "fmt"

func ExampleCarModel() {
	Seed(11)
	fmt.Println(CarModel())
	// Output: Aveo
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

func ExampleVehicle() {
	Seed(11)
	vehicle := Vehicle()
	fmt.Println(vehicle.Brand)
	fmt.Println(vehicle.Fuel)
	fmt.Println(vehicle.Model)
	fmt.Println(vehicle.TransmissionGear)
	fmt.Println(vehicle.VehicleType)
	fmt.Println(vehicle.Year)

	// Output: Fiat
	// Gasoline
	// Freestyle Fwd
	// Automatic
	// Passenger car mini
	// 1925
}
