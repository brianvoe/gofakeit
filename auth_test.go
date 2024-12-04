package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())

	// Output: Treutel8125
}

func ExampleFaker_Username() {
	f := New(11)
	fmt.Println(f.Username())

	// Output: Treutel8125
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

func TestStrongPasswordConstraints(t *testing.T) {
	length := 3

	pass := StrongPassword(true, true, true, true, true, length)
	if len(pass) < 5 {
		t.Error("Password length cannot be less than 5")
	}

	length = 10
	pass = StrongPassword(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}

	// Test fully empty
	pass = StrongPassword(false, false, false, false, false, length)
	if pass == "" {
		t.Error("Password should not be empty")
	}

	// Test it doesn't start or end with a space
	for i := 0; i < 1000; i++ {
		pass = StrongPassword(true, true, true, true, true, length)
		if pass[0] == ' ' || pass[len(pass)-1] == ' ' {
			t.Error("Password should not start or end with a space")
		}
	}
}

func TestStrongPassword(t *testing.T) {
	type args struct {
		lower   bool
		upper   bool
		numeric bool
		special bool
		space   bool
		num     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Only lowercases allowed",
			args: args{
				lower:   true,
				upper:   false,
				numeric: false,
				special: false,
				space:   false,
				num:     10,
			},
		},
		{
			name: "Only uppercases allowed",
			args: args{
				lower:   false,
				upper:   true,
				numeric: false,
				special: false,
				space:   false,
				num:     10,
			},
		},
		{
			name: "Only numerics allowed",
			args: args{
				lower:   false,
				upper:   false,
				numeric: true,
				special: false,
				space:   false,
				num:     10,
			},
		},
		{
			name: "Only special characters allowed",
			args: args{
				lower:   false,
				upper:   false,
				numeric: false,
				special: true,
				space:   false,
				num:     10,
			},
		},
		{
			name: "All options allowed",
			args: args{
				lower:   true,
				upper:   true,
				numeric: true,
				special: true,
				space:   true,
				num:     10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StrongPassword(tt.args.lower, tt.args.upper, tt.args.numeric, tt.args.special, tt.args.space, tt.args.num)

			var lowercasePresented bool
			var uppercasePresented bool
			var numericPresented bool
			var specialPresented bool
			var spacePresented bool

			for i := 0; i < len(got); i++ {
				char := got[i]

				for j := 0; j < len(lowerStr); j++ {
					if char == lowerStr[j] {
						lowercasePresented = true
					}
				}

				for j := 0; j < len(upperStr); j++ {
					if char == upperStr[j] {
						uppercasePresented = true
					}
				}

				for j := 0; j < len(numericStr); j++ {
					if char == numericStr[j] {
						numericPresented = true
					}
				}

				for j := 0; j < len(specialStr); j++ {
					if char == specialStr[j] {
						specialPresented = true
					}
				}

				for j := 0; j < len(spaceStr); j++ {
					if char == spaceStr[j] {
						spacePresented = true
					}
				}
			}

			if tt.args.lower != lowercasePresented {
				sh := "should"
				if !tt.args.lower {
					sh += " not"
				}
				t.Errorf("StrongPassword %s contain lowercase character when lower parameter is set to %t, got this password: %s", sh, tt.args.lower, got)
			}

			if tt.args.upper != uppercasePresented {
				sh := "should"
				if !tt.args.upper {
					sh += " not"
				}
				t.Errorf("StrongPassword %s contain uppercase character when upper parameter is set to %t, got this password: %s", sh, tt.args.lower, got)
			}

			if tt.args.numeric != numericPresented {
				sh := "should"
				if !tt.args.numeric {
					sh += " not"
				}
				t.Errorf("StrongPassword %s contain numeric character when numeric parameter is set to %t, got this password: %s", sh, tt.args.numeric, got)
			}

			if tt.args.special != specialPresented {
				sh := "should"
				if !tt.args.special {
					sh += " not"
				}
				t.Errorf("StrongPassword %s contain special character when special parameter is set to %t, got this password: %s", sh, tt.args.special, got)
			}

			if tt.args.space != spacePresented {
				sh := "should"
				if !tt.args.space {
					sh += " not"
				}
				t.Errorf("StrongPassword %s contain space when space parameter is set to %t, got this password: %s", sh, tt.args.space, got)
			}
		})
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

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func ExampleFaker_Password() {
	f := New(11)
	fmt.Println(f.Password(true, false, false, false, false, 32))
	fmt.Println(f.Password(false, true, false, false, false, 32))
	fmt.Println(f.Password(false, false, true, false, false, 32))
	fmt.Println(f.Password(false, false, false, true, false, 32))
	fmt.Println(f.Password(true, true, true, true, true, 32))
	fmt.Println(f.Password(true, true, true, true, true, 4))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
	// foZnB
}

func ExampleStrongPassword() {
	Seed(11)
	fmt.Println(StrongPassword(true, false, false, false, false, 32))
	fmt.Println(StrongPassword(false, true, false, false, false, 32))
	fmt.Println(StrongPassword(false, false, true, false, false, 32))
	fmt.Println(StrongPassword(false, false, false, true, false, 32))
	fmt.Println(StrongPassword(true, true, true, true, true, 32))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
}

func ExampleFaker_StrongPassword() {
	f := New(11)
	fmt.Println(f.StrongPassword(true, false, false, false, false, 32))
	fmt.Println(f.StrongPassword(false, true, false, false, false, 32))
	fmt.Println(f.StrongPassword(false, false, true, false, false, 32))
	fmt.Println(f.StrongPassword(false, false, false, true, false, 32))
	fmt.Println(f.StrongPassword(true, true, true, true, true, 32))

	// Output: cfelntbponnbbzrhswobuwlxajeeclrx
	// KYEKNGUUNKUYSFBUFFTGDKUVCVYKPONP
	// 43622637275953627791234759581343
	// @.__-._-!-!_..!-_*_*__-@*.__.__!
	// -DTHJ@.oF@d@L5F65 N-.@U5xWX F0DI
}

func BenchmarkPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Password(true, true, true, true, true, 50)
	}
}
