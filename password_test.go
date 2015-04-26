package gofakeit

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	length := 10

	pass := Password(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}
}

func ExamplePassword() {
	Seed(11)
	fmt.Println(Password(true, true, true, true, true, 32))
	// Output: WV10MzLxq2DX79w1omH97_0ga59j8 kj
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password(true, true, true, true, true, 32)
	}
}
