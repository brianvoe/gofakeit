package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func ExampleDate() {
	Seed(11)
	fmt.Println(Date())

	// Output: 2011-11-07 04:31:13.726582492 +0000 UTC
}

func ExampleFaker_Date() {
	f := New(11)
	fmt.Println(f.Date())

	// Output: 2011-11-07 04:31:13.726582492 +0000 UTC
}

func TestDateLookup(t *testing.T) {
	info := GetFuncLookup("date")
	for _, o := range info.Params[0].Options {
		mapParams := NewMapParams()
		mapParams.Add("format", o)
		val, _ := info.Generate(GlobalFaker, mapParams, info)
		if val == "" {
			t.Error("value was empty")
		}
	}
}

func BenchmarkDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Date()
	}
}

func ExamplePastDate() {
	Seed(11)
	fmt.Println(PastDate())
}

func ExampleFaker_PastDate() {
	f := New(11)
	fmt.Println(f.PastDate())
}

func TestPastDate(t *testing.T) {
	now := time.Now()
	pastDate := PastDate()
	if now.Before(pastDate) {
		t.Error("Expected time from past, got: ", pastDate)
	}
}

func BenchmarkPastDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PastDate()
	}
}

func ExampleFutureDate() {
	Seed(11)
	fmt.Println(FutureDate())
}

func ExampleFaker_FutureDate() {
	f := New(11)
	fmt.Println(f.FutureDate())
}

func TestFutureDate(t *testing.T) {
	now := time.Now()
	futureDate := FutureDate()
	if now.After(futureDate) {
		t.Error("Expected time from future, got: ", futureDate)
	}
}

func BenchmarkFutureDate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FutureDate()
	}
}

func ExampleDateRange() {
	Seed(11)
	fmt.Println(DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-03-26 09:20:49.250961474 +0000 UTC
}

func ExampleFaker_DateRange() {
	f := New(11)
	fmt.Println(f.DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-03-26 09:20:49.250961474 +0000 UTC
}

func BenchmarkDateRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
	}
}

func ExampleMonth() {
	Seed(11)
	fmt.Println(Month())

	// Output: 11
}

func ExampleFaker_Month() {
	f := New(11)
	fmt.Println(f.Month())

	// Output: 11
}

func BenchmarkMonth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Month()
	}
}

func ExampleMonthString() {
	Seed(11)
	fmt.Println(MonthString())

	// Output: November
}

func ExampleFaker_MonthString() {
	f := New(11)
	fmt.Println(f.MonthString())

	// Output: November
}

func BenchmarkMonthString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonthString()
	}
}

func ExampleWeekDay() {
	Seed(11)
	fmt.Println(WeekDay())

	// Output: Saturday
}

func ExampleFaker_WeekDay() {
	f := New(11)
	fmt.Println(f.WeekDay())

	// Output: Saturday
}

func BenchmarkWeekDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WeekDay()
	}
}

func ExampleDay() {
	Seed(11)
	fmt.Println(Day())

	// Output: 28
}

func ExampleFaker_Day() {
	f := New(11)
	fmt.Println(f.Day())

	// Output: 28
}

func BenchmarkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day()
	}
}

func ExampleYear() {
	Seed(11)
	fmt.Println(Year())

	// Output: 2011
}

func ExampleFaker_Year() {
	f := New(11)
	fmt.Println(f.Year())

	// Output: 2011
}

func BenchmarkYear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Year()
	}
}

func ExampleHour() {
	Seed(11)
	fmt.Println(Hour())

	// Output: 21
}

func ExampleFaker_Hour() {
	f := New(11)
	fmt.Println(f.Hour())

	// Output: 21
}

func BenchmarkHour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hour()
	}
}

func ExampleMinute() {
	Seed(11)
	fmt.Println(Minute())

	// Output: 53
}

func ExampleFaker_Minute() {
	f := New(11)
	fmt.Println(f.Minute())

	// Output: 53
}

func BenchmarkMinute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Minute()
	}
}

func ExampleSecond() {
	Seed(11)
	fmt.Println(Second())

	// Output: 53
}

func ExampleFaker_Second() {
	f := New(11)
	fmt.Println(f.Second())

	// Output: 53
}

func BenchmarkSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Second()
	}
}

func ExampleNanoSecond() {
	Seed(11)
	fmt.Println(NanoSecond())

	// Output: 895883936
}

func ExampleFaker_NanoSecond() {
	f := New(11)
	fmt.Println(f.NanoSecond())

	// Output: 895883936
}

func BenchmarkNanoSecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NanoSecond()
	}
}

func ExampleTimeZone() {
	Seed(11)
	fmt.Println(TimeZone())

	// Output: West Pacific Standard Time
}

func ExampleFaker_TimeZone() {
	f := New(11)
	fmt.Println(f.TimeZone())

	// Output: West Pacific Standard Time
}

func BenchmarkTimeZone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZone()
	}
}

func ExampleTimeZoneFull() {
	Seed(11)
	fmt.Println(TimeZoneFull())

	// Output: (UTC+10:00) Guam, Port Moresby
}

func ExampleFaker_TimeZoneFull() {
	f := New(11)
	fmt.Println(f.TimeZoneFull())

	// Output: (UTC+10:00) Guam, Port Moresby
}

func BenchmarkTimeZoneFull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneFull()
	}
}

func ExampleTimeZoneAbv() {
	Seed(11)
	fmt.Println(TimeZoneAbv())

	// Output: WPST
}

func ExampleFaker_TimeZoneAbv() {
	f := New(11)
	fmt.Println(f.TimeZoneAbv())

	// Output: WPST
}

func BenchmarkTimeZoneAbv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneAbv()
	}
}

func ExampleTimeZoneOffset() {
	Seed(11)
	fmt.Println(TimeZoneOffset())

	// Output: 10
}

func ExampleFaker_TimeZoneOffset() {
	f := New(11)
	fmt.Println(f.TimeZoneOffset())

	// Output: 10
}

func BenchmarkTimeZoneOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneOffset()
	}
}

func ExampleTimeZoneRegion() {
	Seed(11)
	fmt.Println(TimeZoneRegion())

	// Output: Indian/Chagos
}

func ExampleFaker_TimeZoneRegion() {
	f := New(11)
	fmt.Println(f.TimeZoneRegion())

	// Output: Indian/Chagos
}

func BenchmarkTimeZoneRegion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TimeZoneRegion()
	}
}
