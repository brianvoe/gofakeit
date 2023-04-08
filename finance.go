package gofakeit

import (
	"math/rand"
	"strconv"
	"unicode"
)

// Reference Constants
const cusipRunes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const ppnRunes = "*@#"

// CUSIP
func Cusip() string {
	return cusip(globalFaker.Rand)
}

func (f *Faker) Cusip() string {
	return cusip(f.Rand)
}

func cusip(r *rand.Rand) string {
	return buildCusip(r, false)
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

// PPN CUSIP (Insurance Industry Specific)
func PpnCusip() string {
	return ppnCusip(globalFaker.Rand)
}

func (f *Faker) PpnCusip() string {
	return ppnCusip(f.Rand)
}

func ppnCusip(r *rand.Rand) string {
	return buildCusip(r, true)
}

// ISIN
func Isin() string {
	return isin(globalFaker.Rand)
}

func (f *Faker) Isin() string {
	return isin(f.Rand)
}

func isin(r *rand.Rand) string {
	countryCode := CountryAbr()
	nsin := cusip(r)
	isinChkDig := IsinCheckDigit(countryCode + nsin)
	return countryCode + nsin + isinChkDig
}

func IsinCheckDigit(isin string) string {
	flattened := []int{}
	for _, c := range isin {
		if unicode.IsLetter(c) {
			convVal := int(c) - 55
			// Each digit is added as a separate value
			flattened = append(flattened, convVal/10)
			flattened = append(flattened, convVal%10)
		}
		if unicode.IsDigit(c) {
			flattened = append(flattened, int(c-'0'))
		}
	}

	oddSum := 0
	evenSum := 0

	// Per digit summation of each side.
	for i, d := range flattened {
		if i%2 == 0 {
			elem := 2 * d
			if elem > 9 {
				lastDigit := elem % 10
				firstDigit := elem / 10
				elem = firstDigit + lastDigit
			}
			evenSum += elem
		} else {
			oddSum += d
		}
	}

	return strconv.Itoa((10 - (oddSum+evenSum)%10) % 10)
}

// Helper Functions
func buildCusip(r *rand.Rand, ppn bool) string {
	cusip := replaceWithCusipRunes(r, "????????", ppn)
	chkDigit := CusipCheckDigit(cusip)
	return cusip + chkDigit
}

func replaceWithCusipRunes(r *rand.Rand, str string, ppn bool) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(randCusipRune(r))
		}
	}
	if ppn {
		// PPN Identifiers occur in the 6-8th place
		bytestr[5+r.Intn(3)] = byte(randPpnRune(r))
	}

	return string(bytestr)
}

func randCusipRune(r *rand.Rand) rune {
	return rune(cusipRunes[r.Intn(len(cusipRunes))])
}

func randPpnRune(r *rand.Rand) rune {
	return rune(ppnRunes[r.Intn(len(ppnRunes))])
}

// Lookup Adds
func addFinanceLookup() {
	AddFuncLookup("cusip", Info{
		Display:     "CUSIP",
		Category:    "finance",
		Description: "Random CUSIP",
		Example:     "38259P508",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return cusip(r), nil
		},
	})
	AddFuncLookup("ppncusip", Info{
		Display:     "PPN CUSIP",
		Category:    "finance",
		Description: "Random PPN CUSIP",
		Example:     "38259P508",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return ppnCusip(r), nil
		},
	})
	AddFuncLookup("isin", Info{
		Display:     "ISIN",
		Category:    "finance",
		Description: "Random ISIN",
		Example:     "",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return isin(r), nil
		},
	})
}
