package gofakeit

import (
	"errors"
	"math/rand"
	"strings"
)

// AddressInfo is a struct full of address information
type AddressInfo struct {
	Address   string  `json:"address" xml:"address"`
	Street    string  `json:"street" xml:"street"`
	City      string  `json:"city" xml:"city"`
	State     string  `json:"state" xml:"state"`
	Zip       string  `json:"zip" xml:"zip"`
	Country   string  `json:"country" xml:"country"`
	Latitude  float64 `json:"latitude" xml:"latitude"`
	Longitude float64 `json:"longitude" xml:"longitude"`
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
		street = StreetNumber() + " " + StreetPrefix() + " " + StreetName() + " " + StreetSuffix()
	case 2:
		street = StreetNumber() + " " + StreetName() + " " + StreetSuffix()
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
func Latitude() float64 { return toFixed((rand.Float64()*180)-90, 6) }

// LatitudeInRange will generate a random latitude within the input range
func LatitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -90 || min > 90 || max < -90 || max > 90 {
		return 0, errors.New("Invalid min or max range, must be valid floats and between -90 and 90")
	}
	return toFixed(randFloat64Range(min, max), 6), nil
}

// Longitude will generate a random longitude float64
func Longitude() float64 { return toFixed((rand.Float64()*360)-180, 6) }

// LongitudeInRange will generate a random longitude within the input range
func LongitudeInRange(min, max float64) (float64, error) {
	if min > max || min < -180 || min > 180 || max < -180 || max > 180 {
		return 0, errors.New("Invalid min or max range, must be valid floats and between -180 and 180")
	}
	return toFixed(randFloat64Range(min, max), 6), nil
}

func addAddressLookup() {
	AddFuncLookup("address", Info{
		Display:     "Address",
		Category:    "address",
		Description: "Random set of address info",
		Example: `{
			address: "364 East Rapidsborough, Rutherfordstad, New Jersey 36906",
			street: "364 East Rapidsborough",
			city: "Rutherfordstad",
			state: "New Jersey",
			zip: "36906",
			country: "South Africa",
			latitude: "23.058758",
			longitude: "89.022594"
		}`,
		Output: "map[string]interface",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Address(), nil
		},
	})

	AddFuncLookup("city", Info{
		Display:     "City",
		Category:    "address",
		Description: "Random city",
		Example:     "Marcelside",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return City(), nil
		},
	})

	AddFuncLookup("country", Info{
		Display:     "Country",
		Category:    "address",
		Description: "Random country",
		Example:     "United States of America",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Country(), nil
		},
	})

	AddFuncLookup("countryabr", Info{
		Display:     "Country Abbreviation",
		Category:    "address",
		Description: "Random 2 digit country abbreviation",
		Example:     "US",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CountryAbr(), nil
		},
	})

	AddFuncLookup("state", Info{
		Display:     "State",
		Category:    "address",
		Description: "Random state",
		Example:     "Illinois",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return State(), nil
		},
	})

	AddFuncLookup("stateabr", Info{
		Display:     "State Abbreviation",
		Category:    "address",
		Description: "Random 2 digit state abbreviation",
		Example:     "IL",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return StateAbr(), nil
		},
	})

	AddFuncLookup("street", Info{
		Display:     "Street",
		Category:    "address",
		Description: "Random full street",
		Example:     "364 East Rapidsborough",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Street(), nil
		},
	})

	AddFuncLookup("streetname", Info{
		Display:     "Street Name",
		Category:    "address",
		Description: "Random street name",
		Example:     "View",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return StreetName(), nil
		},
	})

	AddFuncLookup("streetnumber", Info{
		Display:     "Street Number",
		Category:    "address",
		Description: "Random street number",
		Example:     "13645",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return StreetNumber(), nil
		},
	})

	AddFuncLookup("streetprefix", Info{
		Display:     "Street Prefix",
		Category:    "address",
		Description: "Random street prefix",
		Example:     "Lake",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return StreetPrefix(), nil
		},
	})

	AddFuncLookup("streetsuffix", Info{
		Display:     "Street Suffix",
		Category:    "address",
		Description: "Random street suffix",
		Example:     "land",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return StreetSuffix(), nil
		},
	})

	AddFuncLookup("zip", Info{
		Display:     "Zip",
		Category:    "address",
		Description: "Random street zip",
		Example:     "13645",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Zip(), nil
		},
	})

	AddFuncLookup("latitude", Info{
		Display:     "Latitude",
		Category:    "address",
		Description: "Random latitude",
		Example:     "-73.534056",
		Output:      "float",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Latitude(), nil
		},
	})

	AddFuncLookup("latituderange", Info{
		Display:     "Latitude Range",
		Category:    "address",
		Description: "Random latitude between given range",
		Example:     "22.921026",
		Output:      "float",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Default: "0", Description: "Minimum range"},
			{Field: "max", Display: "Max", Type: "float", Default: "90", Description: "Maximum range"},
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

			rangeOut, err := LatitudeInRange(min, max)
			if err != nil {
				return nil, err
			}

			return rangeOut, nil
		},
	})

	AddFuncLookup("longitude", Info{
		Display:     "Longitude",
		Category:    "address",
		Description: "Random longitude",
		Example:     "-147.068112",
		Output:      "float",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Longitude(), nil
		},
	})

	AddFuncLookup("longituderange", Info{
		Display:     "Longitude Range",
		Category:    "address",
		Description: "Random longitude between given range",
		Example:     "-8.170450",
		Output:      "float",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Default: "0", Description: "Minimum range"},
			{Field: "max", Display: "Max", Type: "float", Default: "180", Description: "Maximum range"},
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

			rangeOut, err := LongitudeInRange(min, max)
			if err != nil {
				return nil, err
			}

			return rangeOut, nil
		},
	})
}
