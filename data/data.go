package data

// StringDataType is the type of value in Data
type StringDataType map[string][]string

// IntStringDataType is the type of value in IntData
type IntStringDataType map[string][]int

// Data consists of the main set of fake information
var Data = map[string]StringDataType{
	"person":    Person,
	"address":   Address,
	"company":   Company,
	"job":       Job,
	"lorem":     Lorem,
	"language":  Languages,
	"internet":  Internet,
	"file":      Files,
	"color":     Colors,
	"computer":  Computer,
	"hipster":   Hipster,
	"beer":      Beer,
	"hacker":    Hacker,
	"animal":    Animal,
	"currency":  Currency,
	"log_level": LogLevels,
	"timezone":  TimeZone,
	"car":       Car,
	"emoji":     Emoji,
	"word":      Word,
	"food":      Food,
}

// IntData consists of the main set of fake information (integer only)
var IntData = map[string]IntStringDataType{
	"status_code": StatusCodes,
}
