package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleDomainName() {
	Seed(11)
	fmt.Println(DomainName())
	// Output: centraltarget.org
}

func BenchmarkDomainName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainName()
	}
}

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())
	// Output: com
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix()
	}
}

func ExampleURL() {
	Seed(11)
	fmt.Println(URL())
	// Output: https://www.nationalseamless.net/iterate/streamline/systems
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL()
	}
}

func ExampleHTTPMethod() {
	Seed(11)
	fmt.Println(HTTPMethod())
	// Output: HEAD
}

func BenchmarkHTTPMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPMethod()
	}
}

func ExampleIPv4Address() {
	Seed(11)
	fmt.Println(IPv4Address())
	// Output: 222.83.191.222
}

func BenchmarkIPv4Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4Address()
	}
}

func ExampleIPv6Address() {
	Seed(11)
	fmt.Println(IPv6Address())
	// Output: 2001:cafe:8898:ee17:bc35:9064:5866:d019
}

func BenchmarkIPv6Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv6Address()
	}
}

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())
	// Output: Daniel2872
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
