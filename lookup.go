package gofakeit

import (
	"fmt"
	"strconv"
	"sync"
)

// FuncLookups is the primary map array with mapping to all available data
var FuncLookups map[string]Info
var lockFuncLookups sync.Mutex

// Info structures fields to better break down what each one generates
type Info struct {
	Display     string                                                        `json:"display"`
	Category    string                                                        `json:"category"`
	Description string                                                        `json:"description"`
	Example     string                                                        `json:"example"`
	Output      string                                                        `json:"output"`
	Data        map[string]string                                             `json:"-"`
	Params      []Param                                                       `json:"params"`
	Call        func(m *map[string][]string, info *Info) (interface{}, error) `json:"-"`
}

// Param is a breakdown of param requirements and type definition
type Param struct {
	Field       string   `json:"field"`
	Display     string   `json:"display"`
	Type        string   `json:"type"`
	Default     string   `json:"default"`
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

// Field is used for defining what name and function you to generate for file outuputs
type Field struct {
	Name     string              `json:"name"`
	Function string              `json:"function"`
	Params   map[string][]string `json:"params"`
}

// init will add all the functions to MapLookups
func init() {
	addAuthLookup()
	addAddressLookup()
	addBeerLookup()
	addCarLookup()
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
	addFileJSONLookup()
	addFileXMLLookup()
	addFileCSVLookup()
	addEmojiLookup()
	addImageLookup()
	addNumberLookup()
	addStringLookup()
	addAnimalLookup()
	addGameLookup()
	addFoodLookup()
	addAppLookup()
}

// AddFuncLookup takes a field and adds it to map
func AddFuncLookup(functionName string, info Info) {
	if FuncLookups == nil {
		FuncLookups = make(map[string]Info)
	}

	lockFuncLookups.Lock()
	FuncLookups[functionName] = info
	lockFuncLookups.Unlock()
}

// GetFuncLookup will lookup
func GetFuncLookup(functionName string) *Info {
	info, ok := FuncLookups[functionName]
	if !ok {
		return nil
	}

	return &info
}

// RemoveFuncLookup will remove a function from lookup
func RemoveFuncLookup(functionName string) {
	_, ok := FuncLookups[functionName]
	if !ok {
		return
	}

	lockFuncLookups.Lock()
	delete(FuncLookups, functionName)
	lockFuncLookups.Unlock()
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

	// Get value from map
	if m != nil {
		value, ok := (*m)[field]
		if !ok {
			if p.Default != "" {
				// If default isnt empty use default
				return p, []string{p.Default}, nil
			}

			return nil, nil, fmt.Errorf("Could not find field: %s", field)
		}

		return p, value, nil
	} else if m == nil && p.Default != "" {
		// If default isnt empty use default
		return p, []string{p.Default}, nil
	}

	return nil, nil, fmt.Errorf("Could not find field: %s", field)
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

// GetStringArray will retrieve []string field from data
func (i *Info) GetStringArray(m *map[string][]string, field string) ([]string, error) {
	_, values, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	return values, nil
}

// GetIntArray will retrieve []int field from data
func (i *Info) GetIntArray(m *map[string][]string, field string) ([]int, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	var ints []int
	for i := 0; i < len(value); i++ {
		valueInt, err := strconv.ParseInt(value[i], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%s value could not parse to int", value[i])
		}
		ints = append(ints, int(valueInt))
	}

	return ints, nil
}
