package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleConnective() {
	Seed(11)
	fmt.Println(Connective())

	// Output: through
}

func ExampleFaker_Connective() {
	f := New(11)
	fmt.Println(f.Connective())

	// Output: through
}

func BenchmarkConnective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Connective()
	}
}

func ExampleConnectiveTime() {
	Seed(11)
	fmt.Println(ConnectiveTime())

	// Output: when
}

func ExampleFaker_ConnectiveTime() {
	f := New(11)
	fmt.Println(f.ConnectiveTime())

	// Output: when
}

func BenchmarkConnectiveTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveTime()
	}
}

func ExampleConnectiveComparative() {
	Seed(11)
	fmt.Println(ConnectiveComparative())

	// Output: after all
}

func ExampleFaker_ConnectiveComparative() {
	f := New(11)
	fmt.Println(f.ConnectiveComparative())

	// Output: after all
}

func BenchmarkConnectiveComparative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveComparative()
	}
}

func ExampleConnectiveComplaint() {
	Seed(11)
	fmt.Println(ConnectiveComplaint())

	// Output: i.e.
}

func ExampleFaker_ConnectiveComplaint() {
	f := New(11)
	fmt.Println(f.ConnectiveComplaint())

	// Output: i.e.
}

func BenchmarkConnectiveComplaint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveComplaint()
	}
}

func ExampleConnectiveListing() {
	Seed(11)
	fmt.Println(ConnectiveListing())

	// Output: in summation
}

func ExampleFaker_ConnectiveListing() {
	f := New(11)
	fmt.Println(f.ConnectiveListing())

	// Output: in summation
}

func BenchmarkConnectiveListing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveListing()
	}
}

func ExampleConnectiveCasual() {
	Seed(11)
	fmt.Println(ConnectiveCasual())

	// Output: though
}

func ExampleFaker_ConnectiveCasual() {
	f := New(11)
	fmt.Println(f.ConnectiveCasual())

	// Output: though
}

func BenchmarkConnectiveCasual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveCasual()
	}
}

func ExampleConnectiveExamplify() {
	Seed(11)
	fmt.Println(ConnectiveExamplify())

	// Output: unless
}

func ExampleFaker_ConnectiveExamplify() {
	f := New(11)
	fmt.Println(f.ConnectiveExamplify())

	// Output: unless
}

func BenchmarkConnectiveExamplify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveExamplify()
	}
}
