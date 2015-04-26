package gofakeit

import (
	"math/rand"
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
		street = StreetNumber() + " " + StreetPrefix() + " " + StreetName() + " " + StreetSuffix()
	case 2:
		street = StreetNumber() + " " + StreetName() + " " + StreetSuffix()
	}

	return
}

// Generate Street Number
func StreetNumber() string {
	return strings.TrimLeft(replaceWithNumbers(getRandValue([]string{"address", "number"})), "0")
}

// Generate Street Prefix
func StreetPrefix() string {
	return getRandValue([]string{"address", "street_prefix"})
}

// Generate Street Name
func StreetName() string {
	return getRandValue([]string{"address", "street_name"})
}

// Generate Street Suffix
func StreetSuffix() string {
	return getRandValue([]string{"address", "street_suffix"})
}

// Generate City
func City() (city string) {
	switch randInt := randIntRange(1, 3); randInt {
	case 1:
		city = FirstName() + StreetSuffix()
	case 2:
		city = LastName() + StreetSuffix()
	case 3:
		city = StreetPrefix() + " " + LastName()
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

// Generate Latitude
func Latitude() float64 { return (rand.Float64() * 180) - 90 }

// Generate Latitude
func Longitude() float64 { return (rand.Float64() * 360) - 180 }
