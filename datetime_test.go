package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func ExampleDate() {
	Seed(11)
	fmt.Println(Date())
	// Output: 1900-01-07 04:14:25.685339029 +0000 UTC
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

func ExampleWeekDay() {
	Seed(11)
	fmt.Println(WeekDay())
	// Output: Friday
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

func BenchmarkWeekDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WeekDay()
	}
}

func ExampleYear() {
	Seed(11)
	fmt.Println(Year())
	// Output: 1900
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

func ExampleTimeZone() {
	Seed(11)
	fmt.Println(TimeZone())
	// Output: Kaliningrad Standard Time
}

func BenchmarkTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZone()
	}
}

func ExampleTimeZoneFull() {
	Seed(11)
	fmt.Println(TimeZoneFull())
	// Output: (UTC+03:00) Kaliningrad, Minsk
}

func BenchmarkTimeZoneFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneFull()
	}
}

func ExampleTimeZoneAbv() {
	Seed(11)
	fmt.Println(TimeZoneAbv())
	// Output: KST
}

func BenchmarkTimeZoneAbv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneAbv()
	}
}

func ExampleTimeZoneOffset() {
	Seed(11)
	fmt.Println(TimeZoneOffset())
	// Output: 3
}

func BenchmarkTimeZoneOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneOffset()
	}
}
