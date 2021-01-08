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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			PetName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.PetName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.PetName()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Animal()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Animal()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Animal()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AnimalType()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.AnimalType()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.AnimalType()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FarmAnimal()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.FarmAnimal()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.FarmAnimal()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Cat()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Cat()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Cat()
		}
	})
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Dog()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Dog()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Dog()
		}
	})
}
