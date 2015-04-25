package gofakeit

import (
	"testing"
)

func TestNew(t *testing.T) {
	var numTests int = 100

	for i := 0; i < numTests; i++ {
		New("{name.first} {name.last} {contact.email} #?#?#?")
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("{name.first} {name.last} {contact.email} #?#?#?")
	}
}
