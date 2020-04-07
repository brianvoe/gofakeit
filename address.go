package gofakeit

import (
	"errors"
	"math/rand"
	"strings"
)

// AddressInfo is a struct full of address information
type AddressInfo struct {
	Address   string  `json:"address"`
	Street    string  `json:"street"`
	City      string  `json:"city"`
	State     string  `json:"state"`
	Zip       string  `json:"zip"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Address will generate a struct of address information
func Address() *AddressInfo {
	street := Street()
	city := City()
	state := State()
	zip := Zip()

	return &AddressInfo{
		Address:   street + ", " + city + ", " + state + " " + zip,
		Street:    street,
		City:      city,
		State:     state,
		Zip:       zip,
		Country:   Country(),
		Latitude:  Latitude(),
		Longitude: Longitude(),
	}
}

// Street will generate a random address street string
func Street() (street string) {
	switch randInt := randIntRange(1, 2); randInt {
	case 1:
		street = StreetNumber() + " " + StreetPrefix() + " " + StreetName() + StreetSuffix()
	case 2:
		street = StreetNumber() + " " + StreetName() + StreetSuffix()
	}

	return
}

// StreetNumber will generate a random address street number string
func StreetNumber() string {
	return strings.TrimLeft(replaceWithNumbers(getRandValue([]string{"address", "number"})), "0")
}

// StreetPrefix will generate a random address street prefix string
func StreetPrefix() string {
	return getRandValue([]string{"address", "street_prefix"})
}

// StreetName will generate a random address street name string
func StreetName() string {
	return getRandValue([]string{"address", "street_name"})
}

// StreetSuffix will generate a random address street suffix string
func StreetSuffix() string {
	return getRandValue([]string{"address", "street_suffix"})
}

// City will generate a random city string
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

// State will generate a random state string
func State() string {
	return getRandValue([]string{"address", "state"})
}

// StateAbr will generate a random abbreviated state string
func StateAbr() string {
	return getRandValue([]string{"address", "state_abr"})
}

// Zip will generate a random Zip code string
func Zip() string {
	return replaceWithNumbers(getRandValue([]string{"address", "zip"}))
}

// Country will generate a random country string
func Country() string {
	return getRandValue([]string{"address", "country"})
}

// CountryAbr will generate a random abbreviated country string
func CountryAbr() string {
	return getRandValue([]string{"address", "country_abr"})
}

// Latitude will generate a random latitude float64
func Latitude() float64 { return (rand.Float64() * 180) - 90 }

// LatitudeInRange will generate a random latitude within the input range
func LatitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -90 || min > 90 || max < -90 || max > 90 {
		return 0, errors.New("input range is invalid")
	}
	return randFloat64Range(min, max), nil
}

// Longitude will generate a random longitude float64
func Longitude() float64 { return (rand.Float64() * 360) - 180 }

// LongitudeInRange will generate a random longitude within the input range
func LongitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -180 || min > 180 || max < -180 || max > 180 {
		return 0, errors.New("input range is invalid")
	}
	return randFloat64Range(min, max), nil
}

func addAddressLookup() {
	AddLookupData("address.city", Info{
		Description: "Random city",
		Example:     "Marcelside",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return City(), nil
		},
	})

	AddLookupData("address.country", Info{
		Description: "Random country",
		Example:     "United States of America",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Country(), nil
		},
	})

	AddLookupData("address.country.abr", Info{
		Description: "Random 2 digit country abbreviation",
		Example:     "US",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return CountryAbr(), nil
		},
	})

	AddLookupData("address.state", Info{
		Description: "Random state",
		Example:     "Illinois",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return State(), nil
		},
	})

	AddLookupData("address.state.abr", Info{
		Description: "Random 2 digit state abbreviation",
		Example:     "IL",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StateAbr(), nil
		},
	})

	AddLookupData("address.street", Info{
		Description: "Random full street",
		Example:     "364 East Rapidsborough",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Street(), nil
		},
	})

	AddLookupData("address.street.name", Info{
		Description: "Random street name",
		Example:     "View",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StreetName(), nil
		},
	})

	AddLookupData("address.street.number", Info{
		Description: "Random street number",
		Example:     "13645",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StreetNumber(), nil
		},
	})

	AddLookupData("address.street.prefix", Info{
		Description: "Random street prefix",
		Example:     "Lake",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StreetPrefix(), nil
		},
	})

	AddLookupData("address.street.suffix", Info{
		Description: "Random street suffix",
		Example:     "land",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StreetSuffix(), nil
		},
	})

	AddLookupData("address.zip", Info{
		Description: "Random street zip",
		Example:     "13645",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Zip(), nil
		},
	})

	AddLookupData("address.latitude", Info{
		Description: "Random latitude",
		Example:     "-73.53405629980608",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Latitude(), nil
		},
	})

	AddLookupData("address.latitude.range", Info{
		Description: "Random latitude between given range",
		Example:     "22.921026765022624",
		Params: []Param{
			{Field: "min", Required: true, Type: "float", Description: "Minimum range"},
			{Field: "max", Required: true, Type: "float", Description: "Maximum range"},
		},
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat(m, "max")
			if err != nil {
				return nil, err
			}

			rangeOut, err := LatitudeInRange(min, max)
			if err != nil {
				return nil, errors.New("Invalid min or max range, must be valid floats and between -90 and 90")
			}

			return rangeOut, nil
		},
	})

	AddLookupData("address.longitude", Info{
		Description: "Random longitude",
		Example:     "-147.06811259961216",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return Longitude(), nil
		},
	})

	AddLookupData("address.longitude.range", Info{
		Description: "Random longitude between given range",
		Example:     "-8.170450699978453",
		Params: []Param{
			{Field: "min", Required: true, Type: "float", Description: "Minimum range"},
			{Field: "max", Required: true, Type: "float", Description: "Maximum range"},
		},
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat(m, "max")
			if err != nil {
				return nil, err
			}

			rangeOut, err := LongitudeInRange(min, max)
			if err != nil {
				return nil, errors.New("Invalid min or max range, must be valid floats and between -90 and 90")
			}

			return rangeOut, nil
		},
	})
}
