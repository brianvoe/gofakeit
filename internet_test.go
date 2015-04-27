package gofakeit

import (
	"fmt"
	"testing"
)

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
