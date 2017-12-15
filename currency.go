package gofakeit

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/data"
)

// CurrencyInfo is a struct of currency information
type CurrencyInfo struct {
	Short string
	Long  string
}

// Currency will generate a struct with random currency information
func Currency() *CurrencyInfo {
	index := rand.Intn(len(data.Data["currency"]["short"]))
	return &CurrencyInfo{
		Short: data.Data["currency"]["short"][index],
		Long:  data.Data["currency"]["long"][index],
	}
}

// CurrencyShort will generate a random short currency value
func CurrencyShort() string {
	return getRandValue([]string{"currency", "short"})
}

// CurrencyLong will generate a random long currency name
func CurrencyLong() string {
	return getRandValue([]string{"currency", "long"})
}
