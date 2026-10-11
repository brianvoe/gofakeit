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

// PastDate will generate a random past time.Time struct
func PastDate() time.Time { return pastDate(GlobalFaker) }

// PastDate will generate a random past time.Time struct
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

// javaToGoTimeTokens are Java date tokens mapped to Go layouts, longest first.
// Single letters are matched only when they are not part of a word, so Go text
// such as January, Monday, and PM is left alone.
var javaToGoTimeTokens = []struct{ java, golang string }{
	{"zz:zz", "Z07:00"},
	{"zzzz", "Z0700"},
	{"yyyy", "2006"},
	{"MMMM", "January"},
	{"EEEE", "Monday"},
	{"SSS", "000"},
	{"ddd", "_2"},
	{"MMM", "Jan"},
	{"HH", "15"},
	{"hh", "03"},
	{"mm", "04"},
	{"ss", "05"},
	{"yy", "06"},
	{"aa", "PM"},
	{"ZZ", "-0700"},
	{"dd", "02"},
	{"MM", "01"},
	{"d", "2"},
	{"H", "15"},
	{"h", "3"},
	{"m", "4"},
	{"s", "5"},
	{"y", "06"},
	{"a", "pm"},
	{"M", "1"},
	{"Z", "-07"},
	{"z", "MST"},
	{"E", "Mon"},
}

// javaDateTimeFormatToGolangFormat converts java date/time format into go date/time format
func javaDateTimeFormatToGolangFormat(format string) string {
	var b strings.Builder
	b.Grow(len(format))
	for i := 0; i < len(format); {
		if format[i] == '\'' {
			i = appendJavaQuoted(format, i, &b)
			continue
		}
		if tok, repl, ok := matchJavaTimeToken(format, i); ok {
			b.WriteString(repl)
			i += len(tok)
			continue
		}
		b.WriteByte(format[i])
		i++
	}
	return b.String()
}

// appendJavaQuoted copies a Java quoted literal. ” is one apostrophe.
// i points at the opening quote. The returned index is the first uncopied byte.
func appendJavaQuoted(format string, i int, b *strings.Builder) int {
	if i+1 < len(format) && format[i+1] == '\'' {
		b.WriteByte('\'')
		return i + 2
	}
	i++
	for i < len(format) {
		if format[i] == '\'' {
			if i+1 < len(format) && format[i+1] == '\'' {
				b.WriteByte('\'')
				i += 2
				continue
			}
			return i + 1
		}
		b.WriteByte(format[i])
		i++
	}
	return i
}

func matchJavaTimeToken(format string, i int) (string, string, bool) {
	for _, tok := range javaToGoTimeTokens {
		end := i + len(tok.java)
		if end > len(format) || format[i:end] != tok.java {
			continue
		}
		// Z07:00 and Z0700 are already Go layouts. A Java Z is an offset.
		if tok.java == "Z" && end+2 <= len(format) && format[end:end+2] == "07" {
			continue
		}
		if !javaTokenBounded(format, i, end) {
			continue
		}
		return tok.java, tok.golang, true
	}
	return "", "", false
}

func javaTokenBounded(format string, start, end int) bool {
	if start > 0 && isASCIILetter(format[start-1]) {
		return false
	}
	if end < len(format) && isASCIILetter(format[end]) {
		return false
	}
	return true
}

func isASCIILetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// unpaddedTimeFormat rewrites zero-padded numeric layout tokens (01, 02, 03, 04, 05)
// to the one-or-two-digit forms (1, 2, 3, 4, 5). A token is rewritten only when a
// separator keeps it from touching neighboring digits, so compact layouts such as
// 20060102 stay unchanged.
func unpaddedTimeFormat(format string) (string, bool) {
	var b strings.Builder
	b.Grow(len(format))
	changed := false
	for i := 0; i < len(format); {
		if i+2 <= len(format) && separatedZeroPad(format, i) {
			b.WriteByte(format[i+1])
			changed = true
			i += 2
			continue
		}
		b.WriteByte(format[i])
		i++
	}
	if !changed {
		return "", false
	}
	return b.String(), true
}

// separatedZeroPad reports whether format[i:i+2] is 01, 02, 03, 04, or 05
// and is not glued to another digit.
func separatedZeroPad(format string, i int) bool {
	if format[i] != '0' {
		return false
	}
	switch format[i+1] {
	case '1', '2', '3', '4', '5':
	default:
		return false
	}
	if i > 0 && isASCIIDigit(format[i-1]) {
		return false
	}
	if i+2 < len(format) && isASCIIDigit(format[i+2]) {
		return false
	}
	return true
}

func isASCIIDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func addDateTimeLookup() {
	AddFuncLookup("date", Info{
		Display:     "Date",
		Category:    "datetime",
		Description: "Representation of a specific day, month, and year, often used for chronological reference",
		Example:     "2006-01-02T15:04:05Z07:00",
		Output:      "string",
		Aliases: []string{
			"date string", "calendar date", "datetime", "timestamp", "chronological reference",
		},
		Keywords: []string{
			"time", "day", "month", "year", "format", "rfc3339", "iso8601", "utc",
		},
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

				return f.Date().Format(javaDateTimeFormatToGolangFormat(format)), nil
			}
		},
	})

	AddFuncLookup("daterange", Info{
		Display:     "Date Range",
		Category:    "datetime",
		Description: "Random date between two ranges",
		Example:     "1995-06-15T14:30:00Z",
		Output:      "string",
		Aliases: []string{
			"date interval", "date span", "date window", "between dates", "bounded period",
		},
		Keywords: []string{
			"daterange", "range", "between", "date", "time", "random", "bounds", "limits", "window",
		},
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
			format = javaDateTimeFormatToGolangFormat(format)

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

	AddFuncLookup("pastdate", Info{
		Display:     "PastDate",
		Category:    "datetime",
		Description: "Date that has occurred before the current moment in time",
		Example:     "2007-01-24 13:00:35.820738079 +0000 UTC",
		Output:      "time",
		Aliases: []string{
			"past date", "historical date", "previous date", "earlier date", "prior time",
		},
		Keywords: []string{
			"date", "time", "occurred", "elapsed", "gone", "expired", "finished", "completed",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return pastDate(f), nil },
	})

	AddFuncLookup("futuredate", Info{
		Display:     "FutureDate",
		Category:    "datetime",
		Description: "Date that has occurred after the current moment in time",
		Example:     "2107-01-24 13:00:35.820738079 +0000 UTC",
		Output:      "time",
		Aliases: []string{
			"future date", "upcoming date", "next date", "scheduled date", "later time",
		},
		Keywords: []string{
			"future", "date", "time", "forthcoming", "prospective", "anticipated", "scheduled",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return futureDate(f), nil },
	})

	AddFuncLookup("nanosecond", Info{
		Display:     "Nanosecond",
		Category:    "datetime",
		Description: "Unit of time equal to one billionth (10^-9) of a second",
		Example:     "196446360",
		Output:      "int",
		Aliases: []string{
			"nano", "ns value", "tiny time", "ultra precision", "fractional second",
		},
		Keywords: []string{
			"time", "unit", "second", "billionth", "ultra", "high", "resolution",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return nanoSecond(f), nil },
	})

	AddFuncLookup("second", Info{
		Display:     "Second",
		Category:    "datetime",
		Description: "Unit of time equal to 1/60th of a minute",
		Example:     "43",
		Output:      "int",
		Aliases: []string{
			"second value", "sec unit", "time second", "sixtieth minute", "time slice",
		},
		Keywords: []string{
			"time", "unit", "minute", "sixtieth", "duration", "interval", "sixty",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return second(f), nil },
	})

	AddFuncLookup("minute", Info{
		Display:     "Minute",
		Category:    "datetime",
		Description: "Unit of time equal to 60 seconds",
		Example:     "34",
		Output:      "int",
		Aliases: []string{
			"minute value", "time minute", "sixty seconds", "short period", "clock minute",
		},
		Keywords: []string{
			"time", "unit", "60", "seconds", "duration", "interval", "sixtieth", "hour",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return minute(f), nil },
	})

	AddFuncLookup("hour", Info{
		Display:     "Hour",
		Category:    "datetime",
		Description: "Unit of time equal to 60 minutes",
		Example:     "8",
		Output:      "int",
		Aliases: []string{
			"hour value", "time hour", "sixty minutes", "clock hour", "time period",
		},
		Keywords: []string{
			"time", "unit", "60", "minutes", "duration", "interval", "day",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return hour(f), nil },
	})

	AddFuncLookup("day", Info{
		Display:     "Day",
		Category:    "datetime",
		Description: "24-hour period equivalent to one rotation of Earth on its axis",
		Example:     "12",
		Output:      "int",
		Aliases: []string{
			"calendar day", "day value", "earth rotation", "daily unit", "full day",
		},
		Keywords: []string{
			"time", "unit", "axis", "24-hour", "calendar", "sunrise", "sunset",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return day(f), nil },
	})

	AddFuncLookup("weekday", Info{
		Display:     "Weekday",
		Category:    "datetime",
		Description: "Day of the week excluding the weekend",
		Example:     "Friday",
		Output:      "string",
		Aliases: []string{
			"weekday name", "business day", "work day", "monday to friday", "weekday label",
		},
		Keywords: []string{
			"day", "week", "workday", "business", "calendar", "monday", "tuesday", "wednesday", "thursday", "friday",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return weekDay(f), nil },
	})

	AddFuncLookup("month", Info{
		Display:     "Month",
		Category:    "datetime",
		Description: "Division of the year, typically 30 or 31 days long",
		Example:     "1",
		Output:      "string",
		Aliases: []string{
			"calendar month", "month value", "monthly unit", "date month", "time month",
		},
		Keywords: []string{
			"year", "time", "30", "31", "days", "calendar", "period",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return month(f), nil },
	})

	AddFuncLookup("monthstring", Info{
		Display:     "Month String",
		Category:    "datetime",
		Description: "String representation of a month name",
		Example:     "September",
		Output:      "string",
		Aliases: []string{
			"month name", "calendar month name", "full month", "month label", "month string",
		},
		Keywords: []string{
			"month", "string", "time", "representation", "january", "september", "december",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return monthString(f), nil },
	})

	AddFuncLookup("year", Info{
		Display:     "Year",
		Category:    "datetime",
		Description: "Period of 365 days, the time Earth takes to orbit the Sun",
		Example:     "1900",
		Output:      "int",
		Aliases: []string{
			"calendar year", "annual period", "orbit year", "year value", "fiscal year",
		},
		Keywords: []string{
			"time", "365", "days", "leap", "calendar", "decade", "century",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return year(f), nil },
	})

	AddFuncLookup("timezone", Info{
		Display:     "Timezone",
		Category:    "datetime",
		Description: "Region where the same standard time is used, based on longitudinal divisions of the Earth",
		Example:     "Kaliningrad Standard Time",
		Output:      "string",
		Aliases: []string{
			"time zone", "tz name", "standard time zone", "geographic zone", "regional time",
		},
		Keywords: []string{
			"time", "earth", "utc", "gmt", "pst", "est", "cst", "mst", "dst",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return timeZone(f), nil },
	})

	AddFuncLookup("timezoneabv", Info{
		Display:     "Timezone Abbreviation",
		Category:    "datetime",
		Description: "Abbreviated 3-letter word of a timezone",
		Example:     "KST",
		Output:      "string",
		Aliases: []string{
			"timezone abbr", "tz short code", "abbreviated zone", "short tz name", "zone abbreviation",
		},
		Keywords: []string{
			"time", "3-letter", "kst", "pst", "est", "gmt", "utc",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return timeZoneAbv(f), nil },
	})

	AddFuncLookup("timezonefull", Info{
		Display:     "Timezone Full",
		Category:    "datetime",
		Description: "Full name of a timezone",
		Example:     "(UTC+03:00) Kaliningrad, Minsk",
		Output:      "string",
		Aliases: []string{
			"timezone full", "full tz name", "complete zone name", "long tz name", "detailed zone",
		},
		Keywords: []string{
			"timezone", "full", "time", "standard", "format", "display", "utc", "gmt",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return timeZoneFull(f), nil },
	})

	AddFuncLookup("timezoneoffset", Info{
		Display:     "Timezone Offset",
		Category:    "datetime",
		Description: "The difference in hours from Coordinated Universal Time (UTC) for a specific region",
		Example:     "-5",
		Output:      "float32",
		Aliases: []string{
			"utc offset", "gmt offset", "tz shift", "time difference", "offset value",
		},
		Keywords: []string{
			"timezone", "offset", "utc", "gmt", "plus", "minus", "east", "west",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return timeZoneOffset(f), nil },
	})

	AddFuncLookup("timezoneregion", Info{
		Display:     "Timezone Region",
		Category:    "datetime",
		Description: "Geographic area sharing the same standard time",
		Example:     "America/Alaska",
		Output:      "string",
		Aliases: []string{
			"region zone", "geo time region", "tz area", "regional timezone", "country zone",
		},
		Keywords: []string{
			"timezone", "time", "america", "europe", "asia", "africa", "australia", "continent", "city",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return timeZoneRegion(f), nil },
	})

	AddFuncLookup("time", Info{
		Display:     "Time",
		Category:    "datetime",
		Description: "Random time string in the specified format",
		Example:     "14:30:25",
		Output:      "string",
		Aliases: []string{
			"time string", "clock time", "time format", "time value", "hour minute second",
		},
		Keywords: []string{
			"clock", "hour", "minute", "second", "format", "24-hour", "12-hour", "am", "pm",
		},
		Params: []Param{
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "HH:mm:ss",
				Options:     []string{"HH:mm:ss", "HH:mm", "hh:mm:ss a", "hh:mm a", "H:mm", "h:mm a"},
				Description: "Time format string. Supports Java time format patterns or Go time format patterns",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}

			// Convert java format to golang format
			golangFormat := javaDateTimeFormatToGolangFormat(format)

			// Create a time with today's date but random time
			t := time.Date(2000, 1, 1, hour(f), minute(f), second(f), nanoSecond(f), time.UTC)

			return t.Format(golangFormat), nil
		},
	})

	AddFuncLookup("timerange", Info{
		Display:     "Time Range",
		Category:    "datetime",
		Description: "Random time string between start and end times",
		Example:     "10:15:30",
		Output:      "string",
		Aliases: []string{
			"time interval", "time span", "time window", "between times", "bounded time",
		},
		Keywords: []string{
			"timerange", "range", "between", "time", "start", "end", "bounds", "limits", "window",
		},
		Params: []Param{
			{
				Field:       "starttime",
				Display:     "Start Time",
				Type:        "string",
				Default:     "00:00:00",
				Description: "Start time string in the specified format",
			},
			{
				Field:       "endtime",
				Display:     "End Time",
				Type:        "string",
				Default:     "23:59:59",
				Description: "End time string in the specified format",
			},
			{
				Field:       "format",
				Display:     "Format",
				Type:        "string",
				Default:     "HH:mm:ss",
				Options:     []string{"HH:mm:ss", "HH:mm", "hh:mm:ss a", "hh:mm a", "H:mm", "h:mm a"},
				Description: "Time format string. Supports Java time format patterns or Go time format patterns",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			format, err := info.GetString(m, "format")
			if err != nil {
				return nil, err
			}

			startTime, err := info.GetString(m, "starttime")
			if err != nil {
				return nil, err
			}

			endTime, err := info.GetString(m, "endtime")
			if err != nil {
				return nil, err
			}

			// Convert java format to golang format
			golangFormat := javaDateTimeFormatToGolangFormat(format)

			// Parse start and end times
			start, err := time.Parse(golangFormat, startTime)
			if err != nil {
				// If parsing fails, use a default start time
				start = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			}

			end, err := time.Parse(golangFormat, endTime)
			if err != nil {
				// If parsing fails, use a default end time
				end = time.Date(2000, 1, 1, 23, 59, 59, 999999999, time.UTC)
			}

			// Generate random time between start and end
			startNano := start.UnixNano()
			endNano := end.UnixNano()

			if startNano > endNano {
				startNano, endNano = endNano, startNano
			}

			randomNano := int64(number(f, int(startNano), int(endNano)))
			randomTime := time.Unix(0, randomNano).UTC()

			return randomTime.Format(golangFormat), nil
		},
	})

}
