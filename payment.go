package gofakeit

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v5/data"
)

// CurrencyInfo is a struct of currency information
type CurrencyInfo struct {
	Short string `json:"short"`
	Long  string `json:"long"`
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

// Price will take in a min and max value and return a formatted price
func Price(min, max float64) float64 {
	return math.Floor(randFloat64Range(min, max)*100) / 100
}

// CreditCardInfo is a struct containing credit variables
type CreditCardInfo struct {
	Type   string `json:"type"`
	Number int    `json:"number"`
	Exp    string `json:"exp"`
	Cvv    string `json:"cvv"`
}

// CreditCard will generate a struct full of credit card information
func CreditCard() *CreditCardInfo {
	return &CreditCardInfo{
		Type:   CreditCardType(),
		Number: CreditCardNumber(),
		Exp:    CreditCardExp(),
		Cvv:    CreditCardCvv(),
	}
}

// CreditCardType will generate a random credit card type string
func CreditCardType() string {
	return getRandValue([]string{"payment", "card_type"})
}

// CreditCardNumber will generate a random credit card number int
func CreditCardNumber() int {
	integer, _ := strconv.Atoi(replaceWithNumbers(getRandValue([]string{"payment", "number"})))
	return integer
}

// CreditCardNumberLuhn will generate a random credit card number int that passes luhn test
func CreditCardNumberLuhn() int {
	cc := ""
	for i := 0; i < 100000; i++ {
		cc = replaceWithNumbers(getRandValue([]string{"payment", "number"}))
		if luhn(cc) {
			break
		}
	}
	integer, _ := strconv.Atoi(cc)
	return integer
}

// CreditCardExp will generate a random credit card expiration date string
// Exp date will always be a future date
func CreditCardExp() string {
	month := strconv.Itoa(randIntRange(1, 12))
	if len(month) == 1 {
		month = "0" + month
	}

	var currentYear = time.Now().Year() - 2000
	return month + "/" + strconv.Itoa(randIntRange(currentYear+1, currentYear+10))
}

// CreditCardCvv will generate a random CVV number - Its a string because you could have 017 as an exp date
func CreditCardCvv() string {
	return Numerify("###")
}

// luhn check is used for checking if credit card is valid
func luhn(s string) bool {
	var t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

func addPaymentLookup() {
	AddFuncLookup("currency", Info{
		Display:     "Currency",
		Category:    "payment",
		Description: "Random currency data set",
		Example:     `{short: "USD", long: "United States Dollar"}`,
		Output:      "map[string]string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CurrencyShort(), nil
		},
	})

	AddFuncLookup("currencyshort", Info{
		Display:     "Currency Short",
		Category:    "payment",
		Description: "Random currency abreviated",
		Example:     "USD",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CurrencyShort(), nil
		},
	})

	AddFuncLookup("currencylong", Info{
		Display:     "Currency Long",
		Category:    "payment",
		Description: "Random currency",
		Example:     "United States Dollar",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CurrencyLong(), nil
		},
	})

	AddFuncLookup("price", Info{
		Display:     "Price",
		Category:    "payment",
		Description: "Random monitary price",
		Example:     "92.26",
		Output:      "float64",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Default: "0", Description: "Minumum price value"},
			{Field: "max", Display: "Max", Type: "float", Default: "1000", Description: "Maximum price value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat64(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat64(m, "max")
			if err != nil {
				return nil, err
			}

			return Price(min, max), nil
		},
	})

	AddFuncLookup("creditcard", Info{
		Display:     "Credit Card",
		Category:    "payment",
		Description: "Random credit card data set",
		Example:     `{type: "Visa", number: "4136459948995369", exp: "01/21", cvv: "513"}`,
		Output:      "map[string]interface",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCard(), nil
		},
	})

	AddFuncLookup("creditcardtype", Info{
		Display:     "Credit Card Type",
		Category:    "payment",
		Description: "Random credit card type",
		Example:     "Visa",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCardType(), nil
		},
	})

	AddFuncLookup("creditcardnumber", Info{
		Display:     "Credit Card Number",
		Category:    "payment",
		Description: "Random credit card number",
		Example:     "4136459948995369",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCardNumber(), nil
		},
	})

	AddFuncLookup("creditcardnumberluhn", Info{
		Display:     "Credit Card Number Luhn",
		Category:    "payment",
		Description: "Random credit card number that passes luhn test",
		Example:     "2720996615546177",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCardNumberLuhn(), nil
		},
	})

	AddFuncLookup("creditcardexp", Info{
		Display:     "Credit Card Exp",
		Category:    "payment",
		Description: "Random credit card expiraction date",
		Example:     "01/21",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCardExp(), nil
		},
	})

	AddFuncLookup("creditcardcvv", Info{
		Display:     "Credit Card CVV",
		Category:    "payment",
		Description: "Random credit card number",
		Example:     "513",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CreditCardCvv(), nil
		},
	})
}
