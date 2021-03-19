package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func ExampleDate() {
	Seed(11)
	fmt.Println(Date())
	// Output: 1908-12-07 04:14:25.685339029 +0000 UTC
}

func ExampleFaker_Date() {
	f := New(11)
	fmt.Println(f.Date())
	// Output: 1908-12-07 04:14:25.685339029 +0000 UTC
}

func TestDateLookup(t *testing.T) {
	info := GetFuncLookup("date")
	for _, o := range info.Params[0].Options {
		mapParams := NewMapParams()
		mapParams.Add("format", o)
		val, _ := info.Generate(globalFaker.Rand, mapParams, info)
		if val == "" {
			t.Error("value was empty")
		}
	}
}

func BenchmarkDate(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Date()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Date()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Date()
		}
	})
}

func ExampleDateRange() {
	Seed(11)
	fmt.Println(DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-02-04 14:10:37.166933216 +0000 UTC
}

func ExampleFaker_DateRange() {
	f := New(11)
	fmt.Println(f.DateRange(time.Unix(0, 484633944473634951), time.Unix(0, 1431318744473668209))) // May 10, 1985 years to May 10, 2015
	// Output: 2012-02-04 14:10:37.166933216 +0000 UTC
}

func BenchmarkDateRange(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.DateRange(time.Now().AddDate(-30, 0, 0), time.Now())
		}
	})
}

func ExampleMonth() {
	Seed(11)
	fmt.Println(Month())
	// Output: 1
}

func ExampleFaker_Month() {
	f := New(11)
	fmt.Println(f.Month())
	// Output: 1
}

func BenchmarkMonth(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Month()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Month()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Month()
		}
	})
}

func ExampleMonthString() {
	Seed(11)
	fmt.Println(MonthString())
	// Output: January
}

func ExampleFaker_MonthString() {
	f := New(11)
	fmt.Println(f.MonthString())
	// Output: January
}

func BenchmarkMonthString(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MonthString()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.MonthString()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.MonthString()
		}
	})
}

func ExampleWeekDay() {
	Seed(11)
	fmt.Println(WeekDay())
	// Output: Friday
}

func ExampleFaker_WeekDay() {
	f := New(11)
	fmt.Println(f.WeekDay())
	// Output: Friday
}

func BenchmarkWeekDay(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WeekDay()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.WeekDay()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.WeekDay()
		}
	})
}

func ExampleDay() {
	Seed(11)
	fmt.Println(Day())
	// Output: 12
}

func ExampleFaker_Day() {
	f := New(11)
	fmt.Println(f.Day())
	// Output: 12
}

func BenchmarkDay(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Day()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Day()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Day()
		}
	})
}

func ExampleYear() {
	Seed(11)
	fmt.Println(Year())
	// Output: 1908
}

func ExampleFaker_Year() {
	f := New(11)
	fmt.Println(f.Year())
	// Output: 1908
}

func BenchmarkYear(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Year()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Year()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Year()
		}
	})
}

func ExampleHour() {
	Seed(11)
	fmt.Println(Hour())
	// Output: 0
}

func ExampleFaker_Hour() {
	f := New(11)
	fmt.Println(f.Hour())
	// Output: 0
}

func BenchmarkHour(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hour()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Hour()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Hour()
		}
	})
}

func ExampleMinute() {
	Seed(11)
	fmt.Println(Minute())
	// Output: 0
}

func ExampleFaker_Minute() {
	f := New(11)
	fmt.Println(f.Minute())
	// Output: 0
}

func BenchmarkMinute(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Minute()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Minute()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Minute()
		}
	})
}

func ExampleSecond() {
	Seed(11)
	fmt.Println(Second())
	// Output: 0
}

func ExampleFaker_Second() {
	f := New(11)
	fmt.Println(f.Second())
	// Output: 0
}

func BenchmarkSecond(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Second()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Second()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Second()
		}
	})
}

func ExampleNanoSecond() {
	Seed(11)
	fmt.Println(NanoSecond())
	// Output: 196446360
}

func ExampleFaker_NanoSecond() {
	f := New(11)
	fmt.Println(f.NanoSecond())
	// Output: 196446360
}

func BenchmarkNanoSecond(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NanoSecond()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.NanoSecond()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.NanoSecond()
		}
	})
}

func ExampleTimeZone() {
	Seed(11)
	fmt.Println(TimeZone())
	// Output: Kaliningrad Standard Time
}

func ExampleFaker_TimeZone() {
	f := New(11)
	fmt.Println(f.TimeZone())
	// Output: Kaliningrad Standard Time
}

func BenchmarkTimeZone(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TimeZone()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.TimeZone()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.TimeZone()
		}
	})
}

func ExampleTimeZoneFull() {
	Seed(11)
	fmt.Println(TimeZoneFull())
	// Output: (UTC+03:00) Kaliningrad, Minsk
}

func ExampleFaker_TimeZoneFull() {
	f := New(11)
	fmt.Println(f.TimeZoneFull())
	// Output: (UTC+03:00) Kaliningrad, Minsk
}

func BenchmarkTimeZoneFull(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TimeZoneFull()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.TimeZoneFull()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.TimeZoneFull()
		}
	})
}

func ExampleTimeZoneAbv() {
	Seed(11)
	fmt.Println(TimeZoneAbv())
	// Output: KST
}

func ExampleFaker_TimeZoneAbv() {
	f := New(11)
	fmt.Println(f.TimeZoneAbv())
	// Output: KST
}

func BenchmarkTimeZoneAbv(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TimeZoneAbv()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.TimeZoneAbv()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.TimeZoneAbv()
		}
	})
}

func ExampleTimeZoneOffset() {
	Seed(11)
	fmt.Println(TimeZoneOffset())
	// Output: 3
}

func ExampleFaker_TimeZoneOffset() {
	f := New(11)
	fmt.Println(f.TimeZoneOffset())
	// Output: 3
}

func BenchmarkTimeZoneOffset(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TimeZoneOffset()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.TimeZoneOffset()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.TimeZoneOffset()
		}
	})
}

func ExampleTimeZoneRegion() {
	Seed(11)
	fmt.Println(TimeZoneRegion())
	// Output: America/Vancouver
}

func ExampleFaker_TimeZoneRegion() {
	f := New(11)
	fmt.Println(f.TimeZoneRegion())
	// Output: America/Vancouver
}

func BenchmarkTimeZoneRegion(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TimeZoneRegion()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.TimeZoneRegion()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.TimeZoneRegion()
		}
	})
}
