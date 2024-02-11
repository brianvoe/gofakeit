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

func ExampleFaker_PetName() {
	f := New(11)
	fmt.Println(f.PetName())

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

func ExampleFaker_Animal() {
	f := New(11)
	fmt.Println(f.Animal())

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

func ExampleFaker_AnimalType() {
	f := New(11)
	fmt.Println(f.AnimalType())

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

func ExampleFaker_FarmAnimal() {
	f := New(11)
	fmt.Println(f.FarmAnimal())

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

func ExampleFaker_Cat() {
	f := New(11)
	fmt.Println(f.Cat())

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

func ExampleFaker_Dog() {
	f := New(11)
	fmt.Println(f.Dog())

	// Output: Norwich Terrier
}

func BenchmarkDog(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dog()
	}
}

func ExampleBird() {
	Seed(11)
	fmt.Println(Bird())

	// Output: goose
}

func ExampleFaker_Bird() {
	f := New(11)
	fmt.Println(f.Bird())

	// Output: goose
}

func BenchmarkBird(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bird()
	}
}
