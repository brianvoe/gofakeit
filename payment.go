package gofakeit

import "strconv"

// Generate Random Credit Card Type
func CreditCardType() string {
	return getRandValue([]string{"payment", "card_type"})
}

// Generate Random Credit Card Number
func CreditCardNumber() int {
	integer, _ := strconv.Atoi(replaceWithNumbers(getRandValue([]string{"payment", "number"})))
	return integer
}
