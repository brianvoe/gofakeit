package gofakeit

import (
	"fmt"
	"testing"
)

func ExamplePetName() {
	Seed(11)
	fmt.Println(PetName())
	// Output: Ozzy Pawsborne
}

func BenchmarkPetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PetName()
	}
}

func ExampleAnimal() {
	Seed(11)
	fmt.Println(Animal())
	// Output: elk
}

func BenchmarkAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Animal()
	}
}

func ExampleAnimalType() {
	Seed(11)
	fmt.Println(AnimalType())
	// Output: amphibians
}

func BenchmarkAnimalType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AnimalType()
	}
}

func ExampleFarmAnimal() {
	Seed(11)
	fmt.Println(FarmAnimal())
	// Output: Chicken
}

func BenchmarkFarmAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FarmAnimal()
	}
}

func ExampleCat() {
	Seed(11)
	fmt.Println(Cat())
	// Output: Sokoke
}

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat()
	}
}

func ExampleDog() {
	Seed(11)
	fmt.Println(Dog())
	// Output: Norwich Terrier
}

func BenchmarkDog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dog()
	}
}
