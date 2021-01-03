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
	// 1967
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
	// 1967
}

func BenchmarkCar(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Car()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Car()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Car()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CarType()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CarType()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CarType()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CarFuelType()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CarFuelType()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CarFuelType()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CarTransmissionType()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CarTransmissionType()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CarTransmissionType()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CarMaker()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CarMaker()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CarMaker()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CarModel()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.CarModel()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.CarModel()
		}
	})
}
