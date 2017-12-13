package gofakeit

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/data"
)

// CurrencyInfo is a struct of currency information
type CurrencyInfo struct {
	Code string
	Name string
}

// Currency will generate a struct with random currency information
func Currency() *CurrencyInfo {
	return CurrencyFull()
}

// CurrencyShort will generate a random short currency value
func CurrencyShort() string {
	return getRandValue([]string{"currency", "short"})
}

// CurrencyLong will generate a random long currency name
func CurrencyLong() string {
	return getRandValue([]string{"currency", "long"})
}

// CurrencyFull gives back a full currency with code and name
func CurrencyFull() *CurrencyInfo {
	index := rand.Intn(len(data.Data["currency"]["short"]))
	return &CurrencyInfo{
		Code: data.Data["currency"]["short"][index],
		Name: data.Data["currency"]["long"][index],
	}
}
