package data

// LogLevels consists of log levels for several types
var LogLevels = map[string][]string{
	"syslog": {"EMERG", "ALERT", "CRIT", "ERROR", "WARN", "NOTICE", "INFO", "DEBUG"},
	"apache": {"EMERG", "ALERT", "CRIT", "ERROR", "WARN", "NOTICE", "INFO", "DEBUG", "TRACE1-8"},
}
