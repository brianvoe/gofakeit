package gofakeit

import (
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
	Category    string                                                        `json:"category"`
	Description string                                                        `json:"description"`
	Example     string                                                        `json:"example"`
	Params      []Param                                                       `json:"params"`
	Call        func(m *map[string][]string, info *Info) (interface{}, error) `json:"-"`
}

// Param is a breakdown of param requirements and type definition
type Param struct {
	Field       string `json:"field"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Description string `json:"description"`
}

// SetLookups will add all the functions to MapLookups
func SetLookups() {
	addAuthLookup()
	addAddressLookup()
	addBeerLookup()
	addCarLookup()
	addPersonLookup()
	addWordLookup()
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
	addEmojiLookup()
	addImageLookup()
	addNumberLookup()
	addStringLookup()
	addAnimalLookup()
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

// GetField will retrieve field from data
func (i *Info) GetField(m *map[string][]string, field string) (*Param, []string, error) {
	// Get param
	var p *Param
	for _, param := range i.Params {
		if param.Field == field {
			p = &param
			break
		}
	}
	if p == nil {
		return nil, nil, fmt.Errorf("Could not find param field %s", field)
	}

	// If looking for a field and map is nil
	if m == nil {
		return nil, nil, fmt.Errorf("Could not find field: %s", field)
	}

	// Get value from
	value, ok := (*m)[field]
	if !ok {
		return nil, nil, fmt.Errorf("Could not find field: %s", field)
	}

	// Check requirement
	if p.Required && (len(value) == 0 || value[0] == "") {
		return nil, nil, fmt.Errorf("%s field is required", p.Field)
	} else if len(value) == 0 || value[0] == "" {
		return p, []string{p.Default}, nil
	}

	return p, value, nil
}

// GetBool will retrieve boolean field from data
func (i *Info) GetBool(m *map[string][]string, field string) (bool, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return false, err
	}

	// Try to convert to boolean
	valueBool, err := strconv.ParseBool(value[0])
	if err != nil {
		return false, fmt.Errorf("%s field could not parse to bool value", p.Field)
	}

	return valueBool, nil
}

// GetInt will retrieve int field from data
func (i *Info) GetInt(m *map[string][]string, field string) (int, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to int
	valueInt, err := strconv.ParseInt(value[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to int value", p.Field)
	}

	return int(valueInt), nil
}

// GetUint will retrieve uint field from data
func (i *Info) GetUint(m *map[string][]string, field string) (uint, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to int
	valueUint, err := strconv.ParseUint(value[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to int value", p.Field)
	}

	return uint(valueUint), nil
}

// GetFloat32 will retrieve int field from data
func (i *Info) GetFloat32(m *map[string][]string, field string) (float32, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to float
	valueFloat, err := strconv.ParseFloat(value[0], 32)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to float value", p.Field)
	}

	return float32(valueFloat), nil
}

// GetFloat64 will retrieve int field from data
func (i *Info) GetFloat64(m *map[string][]string, field string) (float64, error) {
	p, value, err := i.GetField(m, field)
	if err != nil {
		return 0, err
	}

	// Try to convert to float
	valueFloat, err := strconv.ParseFloat(value[0], 64)
	if err != nil {
		return 0, fmt.Errorf("%s field could not parse to float value", p.Field)
	}

	return valueFloat, nil
}

// GetString will retrieve string field from data
func (i *Info) GetString(m *map[string][]string, field string) (string, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return "", err
	}

	return value[0], nil
}

// GetStringArray will retrieve string field from data
func (i *Info) GetStringArray(m *map[string][]string, field string) ([]string, error) {
	_, values, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	return values, nil
}

// GetIntArray will retrieve string field from data
func (i *Info) GetIntArray(m *map[string][]string, field string) ([]int, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	var ints []int
	for i := 0; i < len(value); i++ {
		valueInt, err := strconv.ParseInt(value[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%s field could not parse to int value", value[i])
		}
		ints = append(ints, int(valueInt))
	}

	return ints, nil
}
