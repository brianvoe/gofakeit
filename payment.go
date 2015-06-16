package gofakeit

import "strconv"

type CreditCardInfo struct {
	Type   string
	Number int
	Exp    string
	Cvv    string
}

func CreditCard() *CreditCardInfo {
	return &CreditCardInfo{
		Type:   CreditCardType(),
		Number: CreditCardNumber(),
		Exp:    CreditCardExp(),
		Cvv:    CreditCardCvv(),
	}
}

// Generate Random Credit Card Type
func CreditCardType() string {
	return getRandValue([]string{"payment", "card_type"})
}

// Generate Random Credit Card Number
func CreditCardNumber() int {
	integer, _ := strconv.Atoi(replaceWithNumbers(getRandValue([]string{"payment", "number"})))
	return integer
}

// Generate Random Credit Card Expiration Date
func CreditCardExp() string {
	month := strconv.Itoa(randIntRange(1, 12))
	if len(month) == 1 {
		month = "0" + month
	}
	return month + "/" + strconv.Itoa(randIntRange(15, 20))
}

// Generate Random CVV - Its a string because you could have 017 as an exp date
func CreditCardCvv() string {
	return Numerify("###")
}
