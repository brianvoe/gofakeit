package gofakeit

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

// MapLookups is the primary map array with mapping to all available data
var MapLookups Lookup

// Lookup is a primary struct used to organize information with mutex locks
type Lookup struct {
	Map  map[string]Info
	lock sync.Mutex
}

// Info structures fields to better break down what each one generates
type Info struct {
	Description string                                                      `json:"description"`
	Example     string                                                      `json:"example"`
	Params      []Param                                                     `json:"params"`
	Call        func(m *map[string]string, info *Info) (interface{}, error) `json:"-"`
}

// Param is a breakdown of param requirements and type definition
type Param struct {
	Field       string `json:"field"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

func init() {
	addAuthLookup()
	addAddressLookup()
	addBeerLookup()
	addVehicleLookup()
	addPersonLookup()
	addWordLookup()
	addGenerateLookup()
	addMiscLookup()
	addColorLookup()
	addInternetLookup()
	addDateTimeLookup()
	addPaymentLookup()
	addCompanyLookup()
	addHackerLookup()
	addHipsterLookup()
	addLanguagesLookup()
	addFileLookup()
}

// AddLookupData takes a field and adds it to map
func AddLookupData(field string, info Info) {
	if MapLookups.Map == nil {
		MapLookups.Map = make(map[string]Info)
	}

	MapLookups.lock.Lock()
	MapLookups.Map[field] = info
	MapLookups.lock.Unlock()
}

// GetLookupData will lookup
func GetLookupData(field string) *Info {
	info, ok := MapLookups.Map[field]
	if !ok {
		return nil
	}

	return &info
}

// GetField will retrieve field from request
func (i *Info) GetField(m *map[string]string, field string) (*Param, string, error) {
	// Get param
	var p *Param
	for _, param := range i.Params {
		if param.Field == field {
			p = &param
			break
		}
	}
	if p == nil {
		return nil, "", fmt.Errorf("Could not find param field %s", field)
	}

	// Get value from
	value, ok := (*m)[field]
	if !ok {
		return nil, "", errors.New("Could not find field")
	}

	// Check requirement
	if p.Required && value == "" {
		return nil, "", fmt.Errorf("%s field is required", p.Field)
	} else if value == "" {
		return p, p.Default, nil
	}

	return p, value, nil
}

// GetBool will retrieve boolean field from request
func (i *Info) GetBool(m *map[string]string, field string) (bool, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return false, err
	}

	// Try to convert to boolean
	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("%s field could not parse to bool value", p.Field)
	}

	return valueBool, nil
}

// GetInt will retrieve int field from request
func (i *Info) GetInt(m *map[string]string, field string) (int, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to int
	valueInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to int value", p.Field)
	}

	return int(valueInt), nil
}

// GetUint will retrieve uint field from request
func (i *Info) GetUint(m *map[string]string, field string) (uint, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to int
	valueUint, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to int value", p.Field)
	}

	return uint(valueUint), nil
}

// GetFloat will retrieve int field from request
func (i *Info) GetFloat(m *map[string]string, field string) (float64, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to int
	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to float value", p.Field)
	}

	return valueFloat, nil
}

// GetString will retrieve string field from request
func (i *Info) GetString(m *map[string]string, field string) (string, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return "", err
	}

	return value, nil
}
