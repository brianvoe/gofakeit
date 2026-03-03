package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())

	// Output: Gregg812
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())

	// Output: Gregg812
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Username()
	}
}

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

	// Test it doesnt start or end with a space
	for i := 0; i < 1000; i++ {
		pass = Password(true, true, true, true, true, length)
		if pass[0] == ' ' || pass[len(pass)-1] == ' ' {
			t.Error("Password should not start or end with a space")
		}
	}
}

func TestPasswordCharacterSetGuarantee(t *testing.T) {
	// Every generated password must contain at least one character from each
	// enabled character set, regardless of length or randomness.
	for i := 0; i < 1000; i++ {
		pass := Password(true, true, true, true, false, 8)

		var hasLower, hasUpper, hasNumeric, hasSpecial bool
		for _, c := range pass {
			if strings.ContainsRune(lowerStr, c) {
				hasLower = true
			}
			if strings.ContainsRune(upperStr, c) {
				hasUpper = true
			}
			if strings.ContainsRune(numericStr, c) {
				hasNumeric = true
			}
			if strings.ContainsRune(specialSafeStr, c) {
				hasSpecial = true
			}
		}

		if !hasLower {
			t.Errorf("iteration %d: password %q missing lowercase character", i, pass)
		}
		if !hasUpper {
			t.Errorf("iteration %d: password %q missing uppercase character", i, pass)
		}
		if !hasNumeric {
			t.Errorf("iteration %d: password %q missing numeric character", i, pass)
		}
		if !hasSpecial {
			t.Errorf("iteration %d: password %q missing special character", i, pass)
		}
	}
}

func ExamplePassword() {
	Seed(11)
	fmt.Println(Password(true, false, false, false, false, 32))
	fmt.Println(Password(false, true, false, false, false, 32))
	fmt.Println(Password(false, false, true, false, false, 32))
	fmt.Println(Password(false, false, false, true, false, 32))
	fmt.Println(Password(true, true, true, true, true, 32))
	fmt.Println(Password(true, true, true, true, true, 4))

	// Output: bakyxpaaiytqtynbqyfssyhweynxkxhf
	// DPETVGIXXHDGKFXKBZEAQRKZRFDCWVWC
	// 98102640783120607214363311930963
	// --*!_---!__!@--.!_.!@*--@.@__*@*
	// 63.0NpLPR Pv9 h! c!V *gzBv97rvf7
	// mZO-4
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(true, false, false, false, false, 32))
	fmt.Println(f.Password(false, true, false, false, false, 32))
	fmt.Println(f.Password(false, false, true, false, false, 32))
	fmt.Println(f.Password(false, false, false, true, false, 32))
	fmt.Println(f.Password(true, true, true, true, true, 32))
	fmt.Println(f.Password(true, true, true, true, true, 4))

	// Output: bakyxpaaiytqtynbqyfssyhweynxkxhf
	// DPETVGIXXHDGKFXKBZEAQRKZRFDCWVWC
	// 98102640783120607214363311930963
	// --*!_---!__!@--.!_.!@*--@.@__*@*
	// 63.0NpLPR Pv9 h! c!V *gzBv97rvf7
	// mZO-4
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password(true, true, true, true, true, 50)
	}
}
