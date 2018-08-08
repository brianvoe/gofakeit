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

	// Test fully empty
	pass = Password(false, false, false, false, false, length)
	if pass == "" {
		t.Error("Password should not be empty")
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

func ExamplePasswordBulider() {
	Seed(11)
	fmt.Println(PasswordBulider(20))
	// Output: fsdafas
}

func BenchmarkPasswordBulider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PasswordBulider(32)
	}
}
