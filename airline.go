package gofakeit

import (
	"fmt"
	"strings"
)

// AirlineAircraftType will generate a random aircraft type
func AirlineAircraftType() string { return airlineAircraftType(GlobalFaker) }

// AirlineAircraftType will generate a random aircraft type
func (f *Faker) AirlineAircraftType() string { return airlineAircraftType(f) }

func airlineAircraftType(f *Faker) string {
	return getRandValue(f, []string{"airline", "aircraft_type"})
}

// AirlineAirplane will generate a random airplane model
func AirlineAirplane() string { return airlineAirplane(GlobalFaker) }

// AirlineAirplane will generate a random airplane model
func (f *Faker) AirlineAirplane() string { return airlineAirplane(f) }

func airlineAirplane(f *Faker) string {
	return getRandValue(f, []string{"airline", "airplane"})
}

// AirlineAirport will generate a random airport name
func AirlineAirport() string { return airlineAirport(GlobalFaker) }

// AirlineAirport will generate a random airport name
func (f *Faker) AirlineAirport() string { return airlineAirport(f) }

func airlineAirport(f *Faker) string {
	return getRandValue(f, []string{"airline", "airport"})
}

// AirlineAirportIATA will generate a random airport IATA code
func AirlineAirportIATA() string { return airlineAirportIATA(GlobalFaker) }

// AirlineAirportIATA will generate a random airport IATA code
func (f *Faker) AirlineAirportIATA() string { return airlineAirportIATA(f) }

func airlineAirportIATA(f *Faker) string {
	return getRandValue(f, []string{"airline", "iata"})
}

// AirlineFlightNumber will generate a random flight number
func AirlineFlightNumber() string { return airlineFlightNumber(GlobalFaker) }

// AirlineFlightNumber will generate a random flight number
func (f *Faker) AirlineFlightNumber() string { return airlineFlightNumber(f) }

func airlineFlightNumber(f *Faker) string {
	// Generate a 2-letter airline code followed by 1-4 digit flight number
	return fmt.Sprintf("%s%d", strings.ToUpper(letterN(f, 2)), f.Number(1, 9999))
}

// AirlineRecordLocator will generate a random record locator (booking reference)
func AirlineRecordLocator() string { return airlineRecordLocator(GlobalFaker) }

// AirlineRecordLocator will generate a random record locator (booking reference)
func (f *Faker) AirlineRecordLocator() string { return airlineRecordLocator(f) }

func airlineRecordLocator(f *Faker) string {
	// Generate a 6-character uppercase alphanumeric record locator
	return strings.ToUpper(letterN(f, 6))
}

// AirlineSeat will generate a random seat assignment
func AirlineSeat() string { return airlineSeat(GlobalFaker) }

// AirlineSeat will generate a random seat assignment
func (f *Faker) AirlineSeat() string { return airlineSeat(f) }

func airlineSeat(f *Faker) string {
	// Generate seat like "12A", "23F", etc.
	// Row: 1-60, Seat: A-K (excluding I)
	row := f.Number(1, 60)
	seats := []string{"A", "B", "C", "D", "E", "F", "G", "H", "J", "K"}
	seat := seats[f.Number(0, len(seats)-1)]
	return fmt.Sprintf("%d%s", row, seat)
}

func addAirlineLookup() {
	AddFuncLookup("airlineaircrafttype", Info{
		Display:     "Airline Aircraft Type",
		Category:    "airline",
		Description: "Distinct category that defines the particular model or series of an aircraft",
		Example:     "narrowbody",
		Output:      "string",
		Aliases:     []string{"aircraft category", "plane type", "airplane classification"},
		Keywords:    []string{"airline", "aircraft", "type", "category", "narrowbody", "widebody", "regional", "plane", "airplane"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineAircraftType(f), nil
		},
	})

	AddFuncLookup("airlineairplane", Info{
		Display:     "Airline Airplane",
		Category:    "airline",
		Description: "Specific model and manufacturer of an aircraft used for air travel",
		Example:     "Airbus A320",
		Output:      "string",
		Aliases:     []string{"aircraft model", "plane model", "airplane name"},
		Keywords:    []string{"airline", "airplane", "aircraft", "model", "airbus", "boeing", "embraer", "bombardier", "manufacturer"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineAirplane(f), nil
		},
	})

	AddFuncLookup("airlineairport", Info{
		Display:     "Airline Airport",
		Category:    "airline",
		Description: "Facility where aircraft take off and land, including terminals and runways",
		Example:     "Hartsfield-Jackson Atlanta International Airport",
		Output:      "string",
		Aliases:     []string{"airport name", "aerodrome", "airfield"},
		Keywords:    []string{"airline", "airport", "facility", "terminal", "runway", "international", "travel", "aviation"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineAirport(f), nil
		},
	})

	AddFuncLookup("airlineairportiata", Info{
		Display:     "Airline Airport IATA",
		Category:    "airline",
		Description: "Three-letter code assigned to airports by the International Air Transport Association",
		Example:     "ATL",
		Output:      "string",
		Aliases:     []string{"airport code", "iata code", "airport abbreviation"},
		Keywords:    []string{"airline", "airport", "iata", "code", "three-letter", "abbreviation", "international", "aviation"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineAirportIATA(f), nil
		},
	})

	AddFuncLookup("airlineflightnumber", Info{
		Display:     "Airline Flight Number",
		Category:    "airline",
		Description: "Unique identifier for a specific flight operated by an airline",
		Example:     "AA1234",
		Output:      "string",
		Aliases:     []string{"flight code", "flight identifier", "flight designation"},
		Keywords:    []string{"airline", "flight", "number", "identifier", "code", "designation", "aviation", "travel"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineFlightNumber(f), nil
		},
	})

	AddFuncLookup("airlinerecordlocator", Info{
		Display:     "Airline Record Locator",
		Category:    "airline",
		Description: "Unique alphanumeric code used to identify and retrieve a flight booking",
		Example:     "ABCDEF",
		Output:      "string",
		Aliases:     []string{"booking reference", "confirmation code", "reservation code", "pnr"},
		Keywords:    []string{"airline", "record", "locator", "booking", "reference", "confirmation", "reservation", "code", "alphanumeric"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineRecordLocator(f), nil
		},
	})

	AddFuncLookup("airlineseat", Info{
		Display:     "Airline Seat",
		Category:    "airline",
		Description: "Designated location within an aircraft assigned to a passenger",
		Example:     "12A",
		Output:      "string",
		Aliases:     []string{"seat assignment", "seat number", "seat location"},
		Keywords:    []string{"airline", "seat", "assignment", "location", "passenger", "aircraft", "row", "position"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return airlineSeat(f), nil
		},
	})
}
