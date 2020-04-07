package gofakeit

import (
	"strconv"
	"time"
)

// Date will generate a random time.Time struct
func Date() time.Time {
	return time.Date(Year(), time.Month(Number(0, 12)), Day(), Hour(), Minute(), Second(), NanoSecond(), time.UTC)
}

// DateRange will generate a random time.Time struct between a start and end date
func DateRange(start, end time.Time) time.Time {
	return time.Unix(0, int64(Number(int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// Month will generate a random month string
func Month() string {
	return time.Month(Number(1, 12)).String()
}

// Day will generate a random day between 1 - 31
func Day() int {
	return Number(1, 31)
}

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string {
	return time.Weekday(Number(0, 6)).String()
}

// Year will generate a random year between 1900 - current year
func Year() int {
	return Number(1900, time.Now().Year())
}

// Hour will generate a random hour - in military time
func Hour() int {
	return Number(0, 23)
}

// Minute will generate a random minute
func Minute() int {
	return Number(0, 59)
}

// Second will generate a random second
func Second() int {
	return Number(0, 59)
}

// NanoSecond will generate a random nano second
func NanoSecond() int {
	return Number(0, 999999999)
}

// TimeZone will select a random timezone string
func TimeZone() string {
	return getRandValue([]string{"timezone", "text"})
}

// TimeZoneFull will select a random full timezone string
func TimeZoneFull() string {
	return getRandValue([]string{"timezone", "full"})
}

// TimeZoneAbv will select a random timezone abbreviation string
func TimeZoneAbv() string {
	return getRandValue([]string{"timezone", "abr"})
}

// TimeZoneOffset will select a random timezone offset
func TimeZoneOffset() float32 {
	value, _ := strconv.ParseFloat(getRandValue([]string{"timezone", "offset"}), 32)
	return float32(value)
}

func addDateTimeLookup() {
	// TODO: add random datetime output with various options
	// AddLookupData("datetime", Info{
	// 	Description: "Random datetime",
	// 	Example:     "",
	// 	Call: func(m *map[string]string, info *Info) (interface{}, error) {
	// 		return NanoSecond(), nil
	// 	},
	// })

	AddLookupData("datetime.nanosecond", Info{
		Description: "Random nanosecond",
		Example:     "196446360",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return NanoSecond(), nil
		},
	})

	AddLookupData("datetime.second", Info{
		Description: "Random second",
		Example:     "43",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddLookupData("datetime.minute", Info{
		Description: "Random minute",
		Example:     "34",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddLookupData("datetime.hour", Info{
		Description: "Random hour",
		Example:     "8",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddLookupData("datetime.day", Info{
		Description: "Random day",
		Example:     "12",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Day(), nil
		},
	})

	AddLookupData("datetime.weekday", Info{
		Description: "Random week day",
		Example:     "Friday",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return WeekDay(), nil
		},
	})

	AddLookupData("datetime.year", Info{
		Description: "Random year",
		Example:     "1900",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Year(), nil
		},
	})

	AddLookupData("datetime.timezone", Info{
		Description: "Random timezone",
		Example:     "Kaliningrad Standard Time",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return TimeZone(), nil
		},
	})

	AddLookupData("datetime.timezone.abv", Info{
		Description: "Random abreviated timezone",
		Example:     "KST",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return TimeZoneAbv(), nil
		},
	})

	AddLookupData("datetime.timezone.full", Info{
		Description: "Random full timezone",
		Example:     "(UTC+03:00) Kaliningrad, Minsk",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return TimeZoneFull(), nil
		},
	})

	AddLookupData("datetime.timezone.offset", Info{
		Description: "Random timezone offset",
		Example:     "3",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return TimeZoneOffset(), nil
		},
	})

}
