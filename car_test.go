package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleCar() {
	Seed(11)
	car := Car()
	fmt.Println(car.Brand)
	fmt.Println(car.Fuel)
	fmt.Println(car.Model)
	fmt.Println(car.Transmission)
	fmt.Println(car.Type)
	fmt.Println(car.Year)

	// Output: Fiat
	// Gasoline
	// Freestyle Fwd
	// Automatic
	// Passenger car mini
	// 1965
}

func ExampleFaker_Car() {
	f := New(11)
	car := f.Car()
	fmt.Println(car.Brand)
	fmt.Println(car.Fuel)
	fmt.Println(car.Model)
	fmt.Println(car.Transmission)
	fmt.Println(car.Type)
	fmt.Println(car.Year)

	// Output: Fiat
	// Gasoline
	// Freestyle Fwd
	// Automatic
	// Passenger car mini
	// 1965
}

func BenchmarkCar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Car()
	}
}

func ExampleCarType() {
	Seed(11)
	fmt.Println(CarType())

	// Output: Passenger car mini
}

func ExampleFaker_CarType() {
	f := New(11)
	fmt.Println(f.CarType())

	// Output: Passenger car mini
}

func BenchmarkCarType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarType()
	}
}

func ExampleCarFuelType() {
	Seed(11)
	fmt.Println(CarFuelType())

	// Output: CNG
}

func ExampleFaker_CarFuelType() {
	f := New(11)
	fmt.Println(f.CarFuelType())

	// Output: CNG
}

func BenchmarkCarFuelType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarFuelType()
	}
}

func ExampleCarTransmissionType() {
	Seed(11)
	fmt.Println(CarTransmissionType())

	// Output: Manual
}

func ExampleFaker_CarTransmissionType() {
	f := New(11)
	fmt.Println(f.CarTransmissionType())

	// Output: Manual
}

func BenchmarkCarTransmissionType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarTransmissionType()
	}
}

func ExampleCarMaker() {
	Seed(11)
	fmt.Println(CarMaker())

	// Output: Nissan
}

func ExampleFaker_CarMaker() {
	f := New(11)
	fmt.Println(f.CarMaker())

	// Output: Nissan
}

func BenchmarkCarMaker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarMaker()
	}
}

func ExampleCarModel() {
	Seed(11)
	fmt.Println(CarModel())

	// Output: Aveo
}

func ExampleFaker_CarModel() {
	f := New(11)
	fmt.Println(f.CarModel())

	// Output: Aveo
}

func BenchmarkCarModel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CarModel()
	}
}
