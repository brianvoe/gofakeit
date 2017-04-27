package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func ExampleDate() {
	Seed(11)
	fmt.Println(Date())
	// Output: 1978-01-07 04:14:25.685339029 +0000 UTC
}

func BenchmarkDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Date()
	}
}

func ExampleDateRange() {
	Seed(11)
	fmt.Println(DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-02-04 14:10:37.166933216 +0000 UTC
}

func BenchmarkDateRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
	}
}

func ExampleMonth() {
	Seed(11)
	fmt.Println(Month())
	// Output: January
}

func BenchmarkMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Month()
	}
}

func ExampleDay() {
	Seed(11)
	fmt.Println(Day())
	// Output: 12
}

func BenchmarkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day()
	}
}

func ExampleYear() {
	Seed(11)
	fmt.Println(Year())
	// Output: 1978
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Year()
	}
}

func ExampleHour() {
	Seed(11)
	fmt.Println(Hour())
	// Output: 0
}

func BenchmarkHour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hour()
	}
}

func ExampleMinute() {
	Seed(11)
	fmt.Println(Minute())
	// Output: 0
}

func BenchmarkMinute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Minute()
	}
}

func ExampleSecond() {
	Seed(11)
	fmt.Println(Second())
	// Output: 0
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Second()
	}
}

func ExampleNanoSecond() {
	Seed(11)
	fmt.Println(NanoSecond())
	// Output: 196446360
}

func BenchmarkNanoSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NanoSecond()
	}
}
