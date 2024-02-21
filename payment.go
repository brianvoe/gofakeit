package gofakeit

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7/data"
)

// CurrencyInfo is a struct of currency information
type CurrencyInfo struct {
	Short string `json:"short" xml:"short"`
	Long  string `json:"long" xml:"long"`
}

// Currency will generate a struct with random currency information
func Currency() *CurrencyInfo { return currency(GlobalFaker) }

// Currency will generate a struct with random currency information
func (f *Faker) Currency() *CurrencyInfo { return currency(f) }

func currency(f *Faker) *CurrencyInfo {
	index := f.IntN(len(data.Data["currency"]["short"]))
	return &CurrencyInfo{
		Short: data.Data["currency"]["short"][index],
		Long:  data.Data["currency"]["long"][index],
	}
}

// CurrencyShort will generate a random short currency value
func CurrencyShort() string { return currencyShort(GlobalFaker) }

// CurrencyShort will generate a random short currency value
func (f *Faker) CurrencyShort() string { return currencyShort(f) }

func currencyShort(f *Faker) string { return getRandValue(f, []string{"currency", "short"}) }

// CurrencyLong will generate a random long currency name
func CurrencyLong() string { return currencyLong(GlobalFaker) }

// CurrencyLong will generate a random long currency name
func (f *Faker) CurrencyLong() string { return currencyLong(f) }

func currencyLong(f *Faker) string { return getRandValue(f, []string{"currency", "long"}) }

// Price will take in a min and max value and return a formatted price
func Price(min, max float64) float64 { return price(GlobalFaker, min, max) }

// Price will take in a min and max value and return a formatted price
func (f *Faker) Price(min, max float64) float64 { return price(f, min, max) }

func price(f *Faker, min, max float64) float64 {
	return math.Floor(float64Range(f, min, max)*100) / 100
}

// CreditCardInfo is a struct containing credit variables
type CreditCardInfo struct {
	Type   string `json:"type" xml:"type"`
	Number string `json:"number" xml:"number"`
	Exp    string `json:"exp" xml:"exp"`
	Cvv    string `json:"cvv" xml:"cvv"`
}

// CreditCard will generate a struct full of credit card information
func CreditCard() *CreditCardInfo { return creditCard(GlobalFaker) }

// CreditCard will generate a struct full of credit card information
func (f *Faker) CreditCard() *CreditCardInfo { return creditCard(f) }

func creditCard(f *Faker) *CreditCardInfo {
	ccType := randomString(f, data.CreditCardTypes)
	ccv, _ := generate(f, strings.Repeat("#", int(data.CreditCards[randomString(f, data.CreditCardTypes)].Code.Size)))

	return &CreditCardInfo{
		Type:   data.CreditCards[randomString(f, data.CreditCardTypes)].Display,
		Number: creditCardNumber(f, &CreditCardOptions{Types: []string{ccType}}),
		Exp:    creditCardExp(f),
		Cvv:    ccv,
	}
}

// CreditCardType will generate a random credit card type string
func CreditCardType() string { return creditCardType(GlobalFaker) }

// CreditCardType will generate a random credit card type string
func (f *Faker) CreditCardType() string { return creditCardType(f) }

func creditCardType(f *Faker) string {
	return data.CreditCards[randomString(f, data.CreditCardTypes)].Display
}

// CreditCardOptions is the options for credit card number
type CreditCardOptions struct {
	Types []string `json:"types"`
	Bins  []string `json:"bins"` // optional parameter of prepended numbers
	Gaps  bool     `json:"gaps"`
}

// CreditCardNumber will generate a random luhn credit card number
func CreditCardNumber(cco *CreditCardOptions) string { return creditCardNumber(GlobalFaker, cco) }

// CreditCardNumber will generate a random luhn credit card number
func (f *Faker) CreditCardNumber(cco *CreditCardOptions) string { return creditCardNumber(f, cco) }

func creditCardNumber(f *Faker, cco *CreditCardOptions) string {
	if cco == nil {
		cco = &CreditCardOptions{}
	}
	if cco.Types == nil || len(cco.Types) == 0 {
		cco.Types = data.CreditCardTypes
	}
	ccType := randomString(f, cco.Types)

	// Get Card info
	var cardInfo data.CreditCardInfo
	if info, ok := data.CreditCards[ccType]; ok {
		cardInfo = info
	} else {
		ccType = randomString(f, data.CreditCardTypes)
		cardInfo = data.CreditCards[ccType]
	}

	// Get length and pattern
	length := randomUint(f, cardInfo.Lengths)
	numStr := ""
	if len(cco.Bins) >= 1 {
		numStr = randomString(f, cco.Bins)
	} else {
		numStr = strconv.FormatUint(uint64(randomUint(f, cardInfo.Patterns)), 10)
	}
	numStr += strings.Repeat("#", int(length)-len(numStr))
	numStr = numerify(f, numStr)
	ui, _ := strconv.ParseUint(numStr, 10, 64)

	// Loop through until its a valid luhn
	for {
		valid := isLuhn(strconv.FormatUint(ui, 10))
		if valid {
			break
		}
		ui++
	}
	numStr = strconv.FormatUint(ui, 10)

	// Add gaps to number
	if cco.Gaps {
		for i, spot := range cardInfo.Gaps {
			numStr = numStr[:(int(spot)+i)] + " " + numStr[(int(spot)+i):]
		}
	}

	return numStr
}

// CreditCardExp will generate a random credit card expiration date string
// Exp date will always be a future date
func CreditCardExp() string { return creditCardExp(GlobalFaker) }

// CreditCardExp will generate a random credit card expiration date string
// Exp date will always be a future date
func (f *Faker) CreditCardExp() string { return creditCardExp(f) }

func creditCardExp(f *Faker) string {
	month := strconv.Itoa(randIntRange(f, 1, 12))
	if len(month) == 1 {
		month = "0" + month
	}

	var currentYear = time.Now().Year() - 2000
	return month + "/" + strconv.Itoa(randIntRange(f, currentYear+1, currentYear+10))
}

// CreditCardCvv will generate a random CVV number
// Its a string because you could have 017 as an exp date
func CreditCardCvv() string { return creditCardCvv(GlobalFaker) }

// CreditCardCvv will generate a random CVV number
// Its a string because you could have 017 as an exp date
func (f *Faker) CreditCardCvv() string { return creditCardCvv(f) }

func creditCardCvv(f *Faker) string { return numerify(f, "###") }

// isLuhn check is used for checking if credit card is a valid luhn card
func isLuhn(s string) bool {
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

// AchRouting will generate a 9 digit routing number
func AchRouting() string { return achRouting(GlobalFaker) }

// AchRouting will generate a 9 digit routing number
func (f *Faker) AchRouting() string { return achRouting(f) }

func achRouting(f *Faker) string { return numerify(f, "#########") }

// AchAccount will generate a 12 digit account number
func AchAccount() string { return achAccount(GlobalFaker) }

// AchAccount will generate a 12 digit account number
func (f *Faker) AchAccount() string { return achAccount(f) }

func achAccount(f *Faker) string { return numerify(f, "############") }

// BitcoinAddress will generate a random bitcoin address consisting of numbers, upper and lower characters
func BitcoinAddress() string { return bitcoinAddress(GlobalFaker) }

// BitcoinAddress will generate a random bitcoin address consisting of numbers, upper and lower characters
func (f *Faker) BitcoinAddress() string { return bitcoinAddress(f) }

func bitcoinAddress(f *Faker) string {
	return randomString(f, []string{"1", "3"}) + password(f, true, true, true, false, false, number(f, 25, 34))
}

// BitcoinPrivateKey will generate a random bitcoin private key base58 consisting of numbers, upper and lower characters
func BitcoinPrivateKey() string { return bitcoinPrivateKey(GlobalFaker) }

// BitcoinPrivateKey will generate a random bitcoin private key base58 consisting of numbers, upper and lower characters
func (f *Faker) BitcoinPrivateKey() string { return bitcoinPrivateKey(f) }

func bitcoinPrivateKey(f *Faker) string {
	var b strings.Builder
	for i := 0; i < 49; i++ {
		b.WriteString(randCharacter(f, base58))
	}
	return "5" + randomString(f, []string{"H", "J", "K"}) + b.String()
}

func addPaymentLookup() {
	AddFuncLookup("currency", Info{
		Display:     "Currency",
		Category:    "payment",
		Description: "Medium of exchange, often in the form of paper money or coins, used for trade and transactions",
		Example: `{
	"short": "IQD",
	"long": "Iraq Dinar"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return currency(f), nil
		},
	})

	AddFuncLookup("currencyshort", Info{
		Display:     "Currency Short",
		Category:    "payment",
		Description: "Short 3-letter word used to represent a specific currency",
		Example:     "USD",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return currencyShort(f), nil
		},
	})

	AddFuncLookup("currencylong", Info{
		Display:     "Currency Long",
		Category:    "payment",
		Description: "Complete name of a specific currency used for official identification in financial transactions",
		Example:     "United States Dollar",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return currencyLong(f), nil
		},
	})

	AddFuncLookup("price", Info{
		Display:     "Price",
		Category:    "payment",
		Description: "The amount of money or value assigned to a product, service, or asset in a transaction",
		Example:     "92.26",
		Output:      "float64",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Default: "0", Description: "Minimum price value"},
			{Field: "max", Display: "Max", Type: "float", Default: "1000", Description: "Maximum price value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetFloat64(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat64(m, "max")
			if err != nil {
				return nil, err
			}

			return price(f, min, max), nil
		},
	})

	AddFuncLookup("creditcard", Info{
		Display:     "Credit Card",
		Category:    "payment",
		Description: "Plastic card allowing users to make purchases on credit, with payment due at a later date",
		Example: `{
	"type": "UnionPay",
	"number": "4364599489953698",
	"exp": "02/24",
	"cvv": "300"
}`,
		Output:      "map[string]any",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return creditCard(f), nil
		},
	})

	AddFuncLookup("creditcardtype", Info{
		Display:     "Credit Card Type",
		Category:    "payment",
		Description: "Classification of credit cards based on the issuing company",
		Example:     "Visa",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return creditCardType(f), nil
		},
	})

	AddFuncLookup("creditcardnumber", Info{
		Display:     "Credit Card Number",
		Category:    "payment",
		Description: "Unique numerical identifier on a credit card used for making electronic payments and transactions",
		Example:     "4136459948995369",
		Output:      "string",
		Params: []Param{
			{
				Field: "types", Display: "Types", Type: "[]string", Default: "all",
				Options:     []string{"visa", "mastercard", "american-express", "diners-club", "discover", "jcb", "unionpay", "maestro", "elo", "hiper", "hipercard"},
				Description: "A select number of types you want to use when generating a credit card number",
			},
			{Field: "bins", Display: "Bins", Type: "[]string", Optional: true, Description: "Optional list of prepended bin numbers to pick from"},
			{Field: "gaps", Display: "Gaps", Type: "bool", Default: "false", Description: "Whether or not to have gaps in number"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			types, err := info.GetStringArray(m, "types")
			if err != nil {
				return nil, err
			}
			if len(types) == 1 && types[0] == "all" {
				types = []string{}
			}

			bins, _ := info.GetStringArray(m, "bins")

			gaps, err := info.GetBool(m, "gaps")
			if err != nil {
				return nil, err
			}

			options := CreditCardOptions{
				Types: types,
				Gaps:  gaps,
			}

			if len(bins) >= 1 {
				options.Bins = bins
			}

			return creditCardNumber(f, &options), nil
		},
	})

	AddFuncLookup("creditcardexp", Info{
		Display:     "Credit Card Exp",
		Category:    "payment",
		Description: "Date when a credit card becomes invalid and cannot be used for transactions",
		Example:     "01/21",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return creditCardExp(f), nil
		},
	})

	AddFuncLookup("creditcardcvv", Info{
		Display:     "Credit Card CVV",
		Category:    "payment",
		Description: "Three or four-digit security code on a credit card used for online and remote transactions",
		Example:     "513",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return creditCardCvv(f), nil
		},
	})

	AddFuncLookup("achrouting", Info{
		Display:     "ACH Routing Number",
		Category:    "payment",
		Description: "Unique nine-digit code used in the U.S. for identifying the bank and processing electronic transactions",
		Example:     "513715684",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return achRouting(f), nil
		},
	})

	AddFuncLookup("achaccount", Info{
		Display:     "ACH Account Number",
		Category:    "payment",
		Description: "A bank account number used for Automated Clearing House transactions and electronic transfers",
		Example:     "491527954328",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return achAccount(f), nil
		},
	})

	AddFuncLookup("bitcoinaddress", Info{
		Display:     "Bitcoin Address",
		Category:    "payment",
		Description: "Cryptographic identifier used to receive, store, and send Bitcoin cryptocurrency in a peer-to-peer network",
		Example:     "1lWLbxojXq6BqWX7X60VkcDIvYA",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bitcoinAddress(f), nil
		},
	})

	AddFuncLookup("bitcoinprivatekey", Info{
		Display:     "Bitcoin Private Key",
		Category:    "payment",
		Description: "Secret, secure code that allows the owner to access and control their Bitcoin holdings",
		Example:     "5vrbXTADWJ6sQBSYd6lLkG97jljNc0X9VPBvbVqsIH9lWOLcoqg",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bitcoinPrivateKey(f), nil
		},
	})
}
