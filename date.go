package gofakeit

import "time"

// Generate Random Date
func Date() time.Time {
	return time.Date(Year(), time.Month(Number(0, 12)), Day(), Hour(), Minute(), Second(), NanoSecond(), time.UTC)
}

// Generate Random Date From Given Range
func DateRange(start, end time.Time) time.Time {
	return time.Unix(0, int64(Number(int(start.UnixNano()), int(end.UnixNano()))))
}

// Generate Random Month
func Month() string {
	return time.Month(Number(1, 12)).String()
}

// Generate Random Day, 1 - 31
func Day() int {
	return Number(1, 31)
}

// Generate Random Year, 1900 - current year
func Year() int {
	return Number(1900, time.Now().Year())
}

// Generate Random Hour
func Hour() int {
	return Number(0, 23)
}

// Generate Random Minute
func Minute() int {
	return Number(0, 59)
}

// Generate Random Second
func Second() int {
	return Number(0, 59)
}

// Generate Random Nano Second
func NanoSecond() int {
	return Number(0, 999999999)
}
