package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleDomainName() {
	Seed(11)
	fmt.Println(DomainName())
	// Output: centraltarget.biz
}

func BenchmarkDomainName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainName()
	}
}

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())
	// Output: org
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

func ExampleMacAddress() {
	Seed(11)
	fmt.Println(MacAddress())
	// Output: e1:74:cb:01:77:91
}

func BenchmarkMacAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MacAddress()
	}
}

func ExampleUsername() {
	Seed(11)
	fmt.Println(Username())
	// Output: Daniel1364
}

func BenchmarkUsername(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Username()
	}
}
