package gofakeit

import "strconv"
import "time"

var currentYear = time.Now().Year() - 2000

// CreditCardInfo is a struct containing credit variables
type CreditCardInfo struct {
	Type   string
	Number int
	Exp    string
	Cvv    string
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

// CreditCardExp will generate a random credit card expiration date string
func CreditCardExp() string {
	month := strconv.Itoa(randIntRange(1, 12))
	if len(month) == 1 {
		month = "0" + month
	}
	return month + "/" + strconv.Itoa(randIntRange(currentYear, currentYear+5))
}

// CreditCardCvv will generate a random CVV number - Its a string because you could have 017 as an exp date
func CreditCardCvv() string {
	return Numerify("###")
}
