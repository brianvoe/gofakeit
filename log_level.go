package gofakeit

// SysLogLevel will generate a random log level for syslog
func SysLogLevel() string {
	return getRandValue([]string{"log_level", "syslog"})
}

// ApacheLogLevel will generate a random log level for apache log
func ApacheLogLevel() string {
	return getRandValue([]string{"log_level", "apache"})
}
