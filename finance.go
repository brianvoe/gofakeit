package gofakeit

import (
	"math/rand"
	"strconv"
	"unicode"
)

const cusipRunes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ppnRunes = "*@#"

func replaceWithCusipRunes(r *rand.Rand, str string, ppn bool) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(randCusipRune(r, ppn))
		}
	}

	return string(bytestr)
}

func randCusipRune(r *rand.Rand, ppn bool) rune {
	randRunes := cusipRunes
	if ppn {
		randRunes = randRunes + ppnRunes
	}
	return rune(cusipRunes[r.Intn(len(randRunes))])
}

func cusip(r *rand.Rand, ppn bool) string {
	cusip := replaceWithCusipRunes(r, "????????", ppn)
	chkDigit := CusipCheckDigit(cusip)
	return cusip + chkDigit
}

func CusipCheckDigit(cusip string) string {

	sum := 0
	for i, c := range cusip {
		v := 0
		if unicode.IsDigit(c) {
			v = int(c - '0')
		}
		if unicode.IsLetter(c) {
			v = int(c-'A') + 10
		}
		if c == '*' {
			v = 36
		}
		if c == '@' {
			v = 37
		}
		if c == '#' {
			v = 38
		}

		if (7 - i%2) == 0 {
			v = v * 2
		}

		sum = sum + int(v/10) + v%10
	}

	return strconv.Itoa((10 - (sum % 10)) % 10)
}

func Cusip() string {
	return cusip(globalFaker.Rand, false)
}

func PpnCusip() string {
	return cusip(globalFaker.Rand, true)
}

func (f *Faker) Cusip() string {
	return cusip(f.Rand, false)
}

func (f *Faker) PpnCusip() string {
	return cusip(f.Rand, true)
}

/*
func isin(r *rand.Rand) string {
	return "xxxxxxxxxxxx"
}

func symbol(r *rand.Rand) string {
	return "xxxx"
}
*/

func addFinanceLookup() {
	AddFuncLookup("cusip", Info{
		Display:     "CUSIP",
		Category:    "finance",
		Description: "Random CUSIP",
		Example:     "38259P508",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return cusip(r, false), nil
		},
	})
	/*AddFuncLookup("isin", Info{
		Display:     "ISIN",
		Category:    "finance",
		Description: "Random ISIN",
		Example:     "",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return isin(r), nil
		},
	})
	AddFuncLookup("symbol", Info{
		Display:     "symbol",
		Category:    "finance",
		Description: "Random Symbol",
		Example:     "",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return symbol(r), nil
		},
	})*/
}
