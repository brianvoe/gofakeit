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
