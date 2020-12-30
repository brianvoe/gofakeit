package gofakeit

import (
	"errors"
	"strconv"
	"time"
)

// Date will generate a random time.Time struct
func Date() time.Time {
	return time.Date(Year(), time.Month(Number(1, 12)), Day(), Hour(), Minute(), Second(), NanoSecond(), time.UTC)
}

// DateRange will generate a random time.Time struct between a start and end date
func DateRange(start, end time.Time) time.Time {
	return time.Unix(0, int64(Number(int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// NanoSecond will generate a random nano second
func NanoSecond() int {
	return Number(0, 999999999)
}

// Second will generate a random second
func Second() int {
	return Number(0, 59)
}

// Minute will generate a random minute
func Minute() int {
	return Number(0, 59)
}

// Hour will generate a random hour - in military time
func Hour() int {
	return Number(0, 23)
}

// Day will generate a random day between 1 - 31
func Day() int {
	return Number(1, 31)
}

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string {
	return time.Weekday(Number(0, 6)).String()
}

// Month will generate a random month string
func Month() string {
	return time.Month(Number(1, 12)).String()
}

// Year will generate a random year between 1900 - current year
func Year() int {
	return Number(1900, time.Now().Year())
}

// TimeZone will select a random timezone string
func TimeZone() string {
	return getRandValue([]string{"timezone", "text"})
}

// TimeZoneFull will select a random full timezone string
func TimeZoneFull() string {
	return getRandValue([]string{"timezone", "full"})
}

// TimeZoneRegion will select a random region style timezone string, e.g. "America/Chicago"
func TimeZoneRegion() string {
	return getRandValue([]string{"timezone", "region"})
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
	AddFuncLookup("date", Info{
		Display:     "Date",
		Category:    "time",
		Description: "Random date",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "string",
		Params: []Param{
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "RFC3339",
				Options:     []string{"ANSIC", "UnixDate", "RubyDate", "RFC822", "RFC822Z", "RFC850", "RFC1123", "RFC1123Z", "RFC3339", "RFC3339Nano"},
				Description: "Date time string format output",
			},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}

			switch format {
			case "ANSIC":
				return Date().Format(time.ANSIC), nil
			case "UnixDate":
				return Date().Format(time.UnixDate), nil
			case "RubyDate":
				return Date().Format(time.RubyDate), nil
			case "RFC822":
				return Date().Format(time.RFC822), nil
			case "RFC822Z":
				return Date().Format(time.RFC822Z), nil
			case "RFC850":
				return Date().Format(time.RFC850), nil
			case "RFC1123":
				return Date().Format(time.RFC1123), nil
			case "RFC1123Z":
				return Date().Format(time.RFC1123Z), nil
			case "RFC3339":
				return Date().Format(time.RFC3339), nil
			case "RFC3339Nano":
				return Date().Format(time.RFC3339Nano), nil
			}

			return "", errors.New("Invalid format")
		},
	})

	AddFuncLookup("nanosecond", Info{
		Display:     "Nanosecond",
		Category:    "time",
		Description: "Random nanosecond",
		Example:     "196446360",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return NanoSecond(), nil
		},
	})

	AddFuncLookup("second", Info{
		Display:     "Second",
		Category:    "time",
		Description: "Random second",
		Example:     "43",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddFuncLookup("minute", Info{
		Display:     "Minute",
		Category:    "time",
		Description: "Random minute",
		Example:     "34",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddFuncLookup("hour", Info{
		Display:     "Hour",
		Category:    "time",
		Description: "Random hour",
		Example:     "8",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Second(), nil
		},
	})

	AddFuncLookup("day", Info{
		Display:     "Day",
		Category:    "time",
		Description: "Random day",
		Example:     "12",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Day(), nil
		},
	})

	AddFuncLookup("weekday", Info{
		Display:     "Weekday",
		Category:    "time",
		Description: "Random week day",
		Example:     "Friday",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return WeekDay(), nil
		},
	})

	AddFuncLookup("year", Info{
		Display:     "Year",
		Category:    "time",
		Description: "Random year",
		Example:     "1900",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Year(), nil
		},
	})

	AddFuncLookup("timezone", Info{
		Display:     "Timezone",
		Category:    "time",
		Description: "Random timezone",
		Example:     "Kaliningrad Standard Time",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return TimeZone(), nil
		},
	})

	AddFuncLookup("timezoneabv", Info{
		Display:     "Timezone Abbreviation",
		Category:    "time",
		Description: "Random abbreviated timezone",
		Example:     "KST",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return TimeZoneAbv(), nil
		},
	})

	AddFuncLookup("timezonefull", Info{
		Display:     "Timezone Full",
		Category:    "time",
		Description: "Random full timezone",
		Example:     "(UTC+03:00) Kaliningrad, Minsk",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return TimeZoneFull(), nil
		},
	})

	AddFuncLookup("timezoneoffset", Info{
		Display:     "Timezone Offset",
		Category:    "time",
		Description: "Random timezone offset",
		Example:     "3",
		Output:      "float32",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return TimeZoneOffset(), nil
		},
	})

	AddFuncLookup("timezoneregion", Info{
		Display:     "Timezone Region",
		Category:    "time",
		Description: "Random region timezone",
		Example:     "America/Alaska",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return TimeZoneRegion(), nil
		},
	})

}
