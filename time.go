package gofakeit

import (
	"strconv"
	"strings"
	"time"
)

var currentYear = time.Now().Year()

// Date will generate a random time.Time struct
func Date() time.Time { return date(GlobalFaker) }

// Date will generate a random time.Time struct
func (f *Faker) Date() time.Time { return date(f) }

func date(f *Faker) time.Time {
	return time.Date(year(f), time.Month(month(f)), day(f), hour(f), minute(f), second(f), nanoSecond(f), time.UTC)
}

// FutureDate will generate a random past time.Time struct
func PastDate() time.Time { return pastDate(GlobalFaker) }

// FutureDate will generate a random past time.Time struct
func (f *Faker) PastDate() time.Time { return pastDate(f) }

func pastDate(f *Faker) time.Time {
	return time.Now().Add(time.Hour * -time.Duration(number(f, 1, 12)))
}

// FutureDate will generate a random future time.Time struct
func FutureDate() time.Time { return futureDate(GlobalFaker) }

// FutureDate will generate a random future time.Time struct
func (f *Faker) FutureDate() time.Time { return futureDate(f) }

func futureDate(f *Faker) time.Time {
	return time.Now().Add(time.Hour * time.Duration(number(f, 1, 12)))
}

// DateRange will generate a random time.Time struct between a start and end date
func DateRange(start, end time.Time) time.Time { return dateRange(GlobalFaker, start, end) }

// DateRange will generate a random time.Time struct between a start and end date
func (f *Faker) DateRange(start, end time.Time) time.Time { return dateRange(f, start, end) }

func dateRange(f *Faker, start time.Time, end time.Time) time.Time {
	return time.Unix(0, int64(number(f, int(start.UnixNano()), int(end.UnixNano())))).UTC()
}

// NanoSecond will generate a random nano second
func NanoSecond() int { return nanoSecond(GlobalFaker) }

// NanoSecond will generate a random nano second
func (f *Faker) NanoSecond() int { return nanoSecond(f) }

func nanoSecond(f *Faker) int { return number(f, 0, 999999999) }

// Second will generate a random second
func Second() int { return second(GlobalFaker) }

// Second will generate a random second
func (f *Faker) Second() int { return second(f) }

func second(f *Faker) int { return number(f, 0, 59) }

// Minute will generate a random minute
func Minute() int { return minute(GlobalFaker) }

// Minute will generate a random minute
func (f *Faker) Minute() int { return minute(f) }

func minute(f *Faker) int { return number(f, 0, 59) }

// Hour will generate a random hour - in military time
func Hour() int { return hour(GlobalFaker) }

// Hour will generate a random hour - in military time
func (f *Faker) Hour() int { return hour(f) }

func hour(f *Faker) int { return number(f, 0, 23) }

// Day will generate a random day between 1 - 31
func Day() int { return day(GlobalFaker) }

// Day will generate a random day between 1 - 31
func (f *Faker) Day() int { return day(f) }

func day(f *Faker) int { return number(f, 1, 31) }

// WeekDay will generate a random weekday string (Monday-Sunday)
func WeekDay() string { return weekDay(GlobalFaker) }

// WeekDay will generate a random weekday string (Monday-Sunday)
func (f *Faker) WeekDay() string { return weekDay(f) }

func weekDay(f *Faker) string { return time.Weekday(number(f, 0, 6)).String() }

// Month will generate a random month int
func Month() int { return month(GlobalFaker) }

// Month will generate a random month int
func (f *Faker) Month() int { return month(f) }

func month(f *Faker) int { return number(f, 1, 12) }

// MonthString will generate a random month string
func MonthString() string { return monthString(GlobalFaker) }

// MonthString will generate a random month string
func (f *Faker) MonthString() string { return monthString(f) }

func monthString(f *Faker) string { return time.Month(number(f, 1, 12)).String() }

// Year will generate a random year between 1900 - current year
func Year() int { return year(GlobalFaker) }

// Year will generate a random year between 1900 - current year
func (f *Faker) Year() int { return year(f) }

func year(f *Faker) int { return number(f, 1900, currentYear) }

// TimeZone will select a random timezone string
func TimeZone() string { return timeZone(GlobalFaker) }

// TimeZone will select a random timezone string
func (f *Faker) TimeZone() string { return timeZone(f) }

func timeZone(f *Faker) string { return getRandValue(f, []string{"timezone", "text"}) }

// TimeZoneFull will select a random full timezone string
func TimeZoneFull() string { return timeZoneFull(GlobalFaker) }

// TimeZoneFull will select a random full timezone string
func (f *Faker) TimeZoneFull() string { return timeZoneFull(f) }

func timeZoneFull(f *Faker) string { return getRandValue(f, []string{"timezone", "full"}) }

// TimeZoneRegion will select a random region style timezone string, e.g. "America/Chicago"
func TimeZoneRegion() string { return timeZoneRegion(GlobalFaker) }

// TimeZoneRegion will select a random region style timezone string, e.g. "America/Chicago"
func (f *Faker) TimeZoneRegion() string { return timeZoneRegion(f) }

func timeZoneRegion(f *Faker) string { return getRandValue(f, []string{"timezone", "region"}) }

// TimeZoneAbv will select a random timezone abbreviation string
func TimeZoneAbv() string { return timeZoneAbv(GlobalFaker) }

// TimeZoneAbv will select a random timezone abbreviation string
func (f *Faker) TimeZoneAbv() string { return timeZoneAbv(f) }

func timeZoneAbv(f *Faker) string { return getRandValue(f, []string{"timezone", "abr"}) }

// TimeZoneOffset will select a random timezone offset
func TimeZoneOffset() float32 { return timeZoneOffset(GlobalFaker) }

// TimeZoneOffset will select a random timezone offset
func (f *Faker) TimeZoneOffset() float32 { return timeZoneOffset(f) }

func timeZoneOffset(f *Faker) float32 {
	value, _ := strconv.ParseFloat(getRandValue(f, []string{"timezone", "offset"}), 32)
	return float32(value)
}

// javaDateFormatToGolangDateFormat converts java date format into go date format
func javaDateFormatToGolangDateFormat(format string) string {
	format = strings.Replace(format, "ddd", "_2", -1)
	format = strings.Replace(format, "dd", "02", -1)
	format = strings.Replace(format, "d", "2", -1)

	format = strings.Replace(format, "HH", "15", -1)

	format = strings.Replace(format, "hh", "03", -1)
	format = strings.Replace(format, "h", "3", -1)

	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "m", "4", -1)

	format = strings.Replace(format, "ss", "05", -1)
	format = strings.Replace(format, "s", "5", -1)

	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "yy", "06", -1)
	format = strings.Replace(format, "y", "06", -1)

	format = strings.Replace(format, "SSS", "000", -1)

	format = strings.Replace(format, "a", "pm", -1)
	format = strings.Replace(format, "aa", "PM", -1)

	format = strings.Replace(format, "MMMM", "January", -1)
	format = strings.Replace(format, "MMM", "Jan", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "M", "1", -1)

	format = strings.Replace(format, "ZZ", "-0700", -1)

	if !strings.Contains(format, "Z07") {
		format = strings.Replace(format, "Z", "-07", -1)
	}

	format = strings.Replace(format, "zz:zz", "Z07:00", -1)
	format = strings.Replace(format, "zzzz", "Z0700", -1)
	format = strings.Replace(format, "z", "MST", -1)

	format = strings.Replace(format, "EEEE", "Monday", -1)
	format = strings.Replace(format, "E", "Mon", -1)

	return format
}

func addDateTimeLookup() {
	AddFuncLookup("date", Info{
		Display:     "Date",
		Category:    "time",
		Description: "Representation of a specific day, month, and year, often used for chronological reference",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "string",
		Params: []Param{
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "RFC3339",
				Options:     []string{"ANSIC", "UnixDate", "RubyDate", "RFC822", "RFC822Z", "RFC850", "RFC1123", "RFC1123Z", "RFC3339", "RFC3339Nano"},
				Description: "Date time string format output. You may also use golang time format or java time format",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}

			switch format {
			case "ANSIC":
				return f.Date().Format(time.ANSIC), nil
			case "UnixDate":
				return f.Date().Format(time.UnixDate), nil
			case "RubyDate":
				return f.Date().Format(time.RubyDate), nil
			case "RFC822":
				return f.Date().Format(time.RFC822), nil
			case "RFC822Z":
				return f.Date().Format(time.RFC822Z), nil
			case "RFC850":
				return f.Date().Format(time.RFC850), nil
			case "RFC1123":
				return f.Date().Format(time.RFC1123), nil
			case "RFC1123Z":
				return f.Date().Format(time.RFC1123Z), nil
			case "RFC3339":
				return f.Date().Format(time.RFC3339), nil
			case "RFC3339Nano":
				return f.Date().Format(time.RFC3339Nano), nil
			default:
				if format == "" {
					return f.Date().Format(time.RFC3339), nil
				}

				return f.Date().Format(javaDateFormatToGolangDateFormat(format)), nil
			}
		},
	})

	AddFuncLookup("daterange", Info{
		Display:     "DateRange",
		Category:    "time",
		Description: "Random date between two ranges",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "string",
		Params: []Param{
			{
				Field:       "startdate",
				Display:     "Start Date",
				Type:        "string",
				Default:     "1970-01-01",
				Description: "Start date time string",
			},
			{
				Field:       "enddate",
				Display:     "End Date",
				Type:        "string",
				Default:     time.Now().Format("2006-01-02"),
				Description: "End date time string",
			},
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "yyyy-MM-dd",
				Description: "Date time string format",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}
			format = javaDateFormatToGolangDateFormat(format)

			startdate, err := info.GetString(m, "startdate")
			if err != nil {
				return nil, err
			}
			startDateTime, err := time.Parse(format, startdate)
			if err != nil {
				return nil, err
			}

			enddate, err := info.GetString(m, "enddate")
			if err != nil {
				return nil, err
			}
			endDateTime, err := time.Parse(format, enddate)
			if err != nil {
				return nil, err
			}

			return DateRange(startDateTime, endDateTime).Format(format), nil
		},
	})

	AddFuncLookup("pasttime", Info{
		Display:     "PastTime",
		Category:    "time",
		Description: "Date that has occurred before the current moment in time",
		Example:     "2007-01-24 13:00:35.820738079 +0000 UTC",
		Output:      "time",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return pastDate(f), nil
		},
	})

	AddFuncLookup("futuretime", Info{
		Display:     "FutureTime",
		Category:    "time",
		Description: "Date that has occurred after the current moment in time",
		Example:     "2107-01-24 13:00:35.820738079 +0000 UTC",
		Output:      "time",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return futureDate(f), nil
		},
	})

	AddFuncLookup("nanosecond", Info{
		Display:     "Nanosecond",
		Category:    "time",
		Description: "Unit of time equal to One billionth (10^-9) of a second",
		Example:     "196446360",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return nanoSecond(f), nil
		},
	})

	AddFuncLookup("second", Info{
		Display:     "Second",
		Category:    "time",
		Description: "Unit of time equal to 1/60th of a minute",
		Example:     "43",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return second(f), nil
		},
	})

	AddFuncLookup("minute", Info{
		Display:     "Minute",
		Category:    "time",
		Description: "Unit of time equal to 60 seconds",
		Example:     "34",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minute(f), nil
		},
	})

	AddFuncLookup("hour", Info{
		Display:     "Hour",
		Category:    "time",
		Description: "Unit of time equal to 60 minutes",
		Example:     "8",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hour(f), nil
		},
	})

	AddFuncLookup("day", Info{
		Display:     "Day",
		Category:    "time",
		Description: "24-hour period equivalent to one rotation of Earth on its axis",
		Example:     "12",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return day(f), nil
		},
	})

	AddFuncLookup("weekday", Info{
		Display:     "Weekday",
		Category:    "time",
		Description: "Day of the week excluding the weekend",
		Example:     "Friday",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return weekDay(f), nil
		},
	})

	AddFuncLookup("month", Info{
		Display:     "Month",
		Category:    "time",
		Description: "Division of the year, typically 30 or 31 days long",
		Example:     "1",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return month(f), nil
		},
	})

	AddFuncLookup("monthstring", Info{
		Display:     "Month String",
		Category:    "time",
		Description: "String Representation of a month name",
		Example:     "September",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return monthString(f), nil
		},
	})

	AddFuncLookup("year", Info{
		Display:     "Year",
		Category:    "time",
		Description: "Period of 365 days, the time Earth takes to orbit the Sun",
		Example:     "1900",
		Output:      "int",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return year(f), nil
		},
	})

	AddFuncLookup("timezone", Info{
		Display:     "Timezone",
		Category:    "time",
		Description: "Region where the same standard time is used, based on longitudinal divisions of the Earth",
		Example:     "Kaliningrad Standard Time",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return timeZone(f), nil
		},
	})

	AddFuncLookup("timezoneabv", Info{
		Display:     "Timezone Abbreviation",
		Category:    "time",
		Description: "Abbreviated 3-letter word of a timezone",
		Example:     "KST",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return timeZoneAbv(f), nil
		},
	})

	AddFuncLookup("timezonefull", Info{
		Display:     "Timezone Full",
		Category:    "time",
		Description: "Full name of a timezone",
		Example:     "(UTC+03:00) Kaliningrad, Minsk",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return timeZoneFull(f), nil
		},
	})

	AddFuncLookup("timezoneoffset", Info{
		Display:     "Timezone Offset",
		Category:    "time",
		Description: "The difference in hours from Coordinated Universal Time (UTC) for a specific region",
		Example:     "3",
		Output:      "float32",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return timeZoneOffset(f), nil
		},
	})

	AddFuncLookup("timezoneregion", Info{
		Display:     "Timezone Region",
		Category:    "time",
		Description: "Geographic area sharing the same standard time",
		Example:     "America/Alaska",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return timeZoneRegion(f), nil
		},
	})

}
