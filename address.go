package gofakeit

import (
	"strings"
)

// Generate Address
func Address() string {
	return Street() + ", " + City() + ", " + State() + " " + Zip()
}

// Generate Street
func Street() (street string) {
	switch randInt := randIntRange(1, 2); randInt {
	case 1:
		street = AddressNumber() + " " + AddressStreetPrefix() + " " + AddressStreetName() + " " + AddressStreetSuffix()
	case 2:
		street = AddressNumber() + " " + AddressStreetName() + " " + AddressStreetSuffix()
	}

	return
}

// Generate Address Number
func AddressNumber() string {
	return strings.TrimLeft(replaceWithNumbers(getRandValue([]string{"address", "number"})), "0")
}

// Generate Street Prefix
func AddressStreetPrefix() string {
	return getRandValue([]string{"address", "street_prefix"})
}

// Generate Street Name
func AddressStreetName() string {
	return getRandValue([]string{"address", "street_name"})
}

// Generate Street Suffix
func AddressStreetSuffix() string {
	return getRandValue([]string{"address", "street_suffix"})
}

// Generate City
func City() (city string) {
	switch randInt := randIntRange(1, 3); randInt {
	case 1:
		city = FirstName() + AddressStreetSuffix()
	case 2:
		city = LastName() + AddressStreetSuffix()
	case 3:
		city = AddressStreetPrefix() + " " + LastName()
	}

	return
}

// Generate State
func State() string {
	return getRandValue([]string{"address", "state"})
}

// Generate Abreviated State
func StateAbr() string {
	return getRandValue([]string{"address", "state_abr"})
}

// Generate Zip Code
func Zip() string {
	return replaceWithNumbers(getRandValue([]string{"address", "zip"}))
}

// Generate Country
func Country() string {
	return getRandValue([]string{"address", "country"})
}
