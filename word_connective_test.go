package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleConnective() {
	Seed(11)
	fmt.Println(Connective())

	// Output: such as
}

func ExampleFaker_Connective() {
	f := New(11)
	fmt.Println(f.Connective())

	// Output: such as
}

func BenchmarkConnective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Connective()
	}
}

func ExampleConnectiveTime() {
	Seed(11)
	fmt.Println(ConnectiveTime())

	// Output: finally
}

func ExampleFaker_ConnectiveTime() {
	f := New(11)
	fmt.Println(f.ConnectiveTime())

	// Output: finally
}

func BenchmarkConnectiveTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveTime()
	}
}

func ExampleConnectiveComparative() {
	Seed(11)
	fmt.Println(ConnectiveComparative())

	// Output: in addition
}

func ExampleFaker_ConnectiveComparative() {
	f := New(11)
	fmt.Println(f.ConnectiveComparative())

	// Output: in addition
}

func BenchmarkConnectiveComparative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveComparative()
	}
}

func ExampleConnectiveComplaint() {
	Seed(11)
	fmt.Println(ConnectiveComplaint())

	// Output: besides
}

func ExampleFaker_ConnectiveComplaint() {
	f := New(11)
	fmt.Println(f.ConnectiveComplaint())

	// Output: besides
}

func BenchmarkConnectiveComplaint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveComplaint()
	}
}

func ExampleConnectiveListing() {
	Seed(11)
	fmt.Println(ConnectiveListing())

	// Output: firstly
}

func ExampleFaker_ConnectiveListing() {
	f := New(11)
	fmt.Println(f.ConnectiveListing())

	// Output: firstly
}

func BenchmarkConnectiveListing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveListing()
	}
}

func ExampleConnectiveCasual() {
	Seed(11)
	fmt.Println(ConnectiveCasual())

	// Output: an outcome of
}

func ExampleFaker_ConnectiveCasual() {
	f := New(11)
	fmt.Println(f.ConnectiveCasual())

	// Output: an outcome of
}

func BenchmarkConnectiveCasual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveCasual()
	}
}

func ExampleConnectiveExamplify() {
	Seed(11)
	fmt.Println(ConnectiveExamplify())

	// Output: then
}

func ExampleFaker_ConnectiveExamplify() {
	f := New(11)
	fmt.Println(f.ConnectiveExamplify())

	// Output: then
}

func BenchmarkConnectiveExamplify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConnectiveExamplify()
	}
}
