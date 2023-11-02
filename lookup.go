package gofakeit

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// FuncLookups is the primary map array with mapping to all available data
var FuncLookups map[string]Info
var lockFuncLookups sync.Mutex

// internalFuncLookups is the internal map array with mapping to all available data
var internalFuncLookups map[string]Info = map[string]Info{
	"internal_exampleFields": {
		Description: "Example fields for generating csv, json and xml",
		Output:      "gofakeit.Field",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			name, _ := getRandomFuncLookup(r, excludeWithParams,
				validTypes("string", "int", "[]string", "[]int"))
			return Field{
				Name:     name,
				Function: name,
			}, nil
		},
	},
}

// filterFuncLookup returns true when the lookup should be accepted
type filterFuncLookup func(Info) bool

var (
	excludeWithParams filterFuncLookup = func(info Info) bool {
		return len(info.Params) == 0
	}

	validTypes = func(acceptedTypes ...string) filterFuncLookup {
		return func(info Info) bool {
			for _, t := range acceptedTypes {
				if info.Output == t {
					return true
				}
			}
			return false
		}
	}
)

func getRandomFuncLookup(r *rand.Rand, filters ...filterFuncLookup) (string, Info) {
	var keys []string
	for k, v := range FuncLookups {
		isValid := true
		for _, filter := range filters {
			isValid = isValid && filter(v)
		}
		if isValid {
			keys = append(keys, k)
		}
	}

	sort.Stable(sort.StringSlice(keys))

	selected := keys[r.Intn(len(keys))]
	return selected, FuncLookups[selected]
}

// MapParams is the values to pass into a lookup generate
type MapParams map[string]MapParamsValue

type MapParamsValue []string

// Info structures fields to better break down what each one generates
type Info struct {
	Display     string                                                            `json:"display"`
	Category    string                                                            `json:"category"`
	Description string                                                            `json:"description"`
	Example     string                                                            `json:"example"`
	Output      string                                                            `json:"output"`
	ContentType string                                                            `json:"content_type"`
	Params      []Param                                                           `json:"params"`
	Generate    func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) `json:"-"`
}

// Param is a breakdown of param requirements and type definition
type Param struct {
	Field       string   `json:"field"`
	Display     string   `json:"display"`
	Type        string   `json:"type"`
	Optional    bool     `json:"optional"`
	Default     string   `json:"default"`
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

// Field is used for defining what name and function you to generate for file outuputs
type Field struct {
	Name     string    `json:"name"`
	Function string    `json:"function"`
	Params   MapParams `json:"params"`
}

func init() { initLookup() }

// init will add all the functions to MapLookups
func initLookup() {
	addAuthLookup()
	addAddressLookup()
	addBeerLookup()
	addCarLookup()
	addPersonLookup()
	addWordGeneralLookup()
	addWordNounLookup()
	addWordVerbLookup()
	addWordAdverbLookup()
	addWordPrepositionLookup()
	addWordAdjectiveLookup()
	addWordPronounLookup()
	addWordConnectiveLookup()
	addWordPhraseLookup()
	addWordSentenceLookup()
	addWordGrammerLookup()
	addLoremLookup()
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
	addWeightedLookup()
	addMinecraftLookup()
	addCelebrityLookup()
	addDatabaseSQLLookup()
	addErrorLookup()
	addHtmlLookup()
	addFinanceLookup()
	addBookLookup()
	addMovieLookup()
	addSchoolLookup()
}

// NewMapParams will create a new MapParams
func NewMapParams() *MapParams {
	return &MapParams{}
}

// Add will take in a field and value and add it to the map params type
func (m *MapParams) Add(field string, value string) {
	_, ok := (*m)[field]
	if !ok {
		(*m)[field] = []string{value}
		return
	}

	(*m)[field] = append((*m)[field], value)
}

// Get will return the array of string from the provided field
func (m *MapParams) Get(field string) []string {
	return (*m)[field]
}

// Size will return the total size of the underlying map
func (m *MapParams) Size() int {
	size := 0
	for range *m {
		size++
	}
	return size
}

// UnmarshalJSON will unmarshal the json into the []string
func (m *MapParamsValue) UnmarshalJSON(data []byte) error {
	// check if the data is an array
	// if so, marshal it into m
	if data[0] == '[' {
		var values []interface{}
		err := json.Unmarshal(data, &values)
		if err != nil {
			return err
		}

		// convert the values to array of strings
		for _, value := range values {
			typeOf := reflect.TypeOf(value).Kind().String()

			if typeOf == "map" {
				v, err := json.Marshal(value)
				if err != nil {
					return err
				}
				*m = append(*m, string(v))
			} else {
				*m = append(*m, fmt.Sprintf("%v", value))
			}
		}
		return nil
	}

	// if not, then convert into a string and add it to m
	var s interface{}
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*m = append(*m, fmt.Sprintf("%v", s))
	return nil
}

// AddFuncLookup takes a field and adds it to map
func AddFuncLookup(functionName string, info Info) {
	if _, ok := internalFuncLookups[functionName]; ok {
		panic(fmt.Sprintf("Function %s is used internally and cannot be overwritten", functionName))
	}

	if FuncLookups == nil {
		FuncLookups = make(map[string]Info)
	}

	// Check content type
	if info.ContentType == "" {
		info.ContentType = "text/plain"
	}

	lockFuncLookups.Lock()
	FuncLookups[functionName] = info
	lockFuncLookups.Unlock()
}

// GetFuncLookup will lookup
func GetFuncLookup(functionName string) *Info {
	var info Info
	var ok bool

	info, ok = internalFuncLookups[functionName]
	if ok {
		return &info
	}

	info, ok = FuncLookups[functionName]
	if ok {
		return &info
	}

	return nil
}

// RemoveFuncLookup will remove a function from lookup
func RemoveFuncLookup(functionName string) {
	if _, ok := internalFuncLookups[functionName]; ok {
		panic(fmt.Sprintf("Function %s is used internally and cannot be overwritten", functionName))
	}

	_, ok := FuncLookups[functionName]
	if !ok {
		return
	}

	lockFuncLookups.Lock()
	delete(FuncLookups, functionName)
	lockFuncLookups.Unlock()
}

// GetField will retrieve field from data
func (i *Info) GetField(m *MapParams, field string) (*Param, []string, error) {
	// Get param
	var p *Param
	for _, param := range i.Params {
		if param.Field == field {
			p = &param
			break
		}
	}
	if p == nil {
		return nil, nil, fmt.Errorf("could not find param field %s", field)
	}

	// Get value from map
	if m != nil {
		value, ok := (*m)[field]
		if !ok {
			// If default isnt empty use default
			if p.Default != "" {
				return p, []string{p.Default}, nil
			}

			return nil, nil, fmt.Errorf("could not find field: %s", field)
		}

		return p, value, nil
	} else if m == nil && p.Default != "" {
		// If p.Type is []uint, then we need to convert it to []string
		if strings.HasPrefix(p.Default, "[") {
			// Remove [] from type
			defaultClean := p.Default[1 : len(p.Default)-1]

			// Split on comma
			defaultSplit := strings.Split(defaultClean, ",")

			return p, defaultSplit, nil
		}

		// If default isnt empty use default
		return p, []string{p.Default}, nil
	}

	return nil, nil, fmt.Errorf("could not find field: %s", field)
}

// GetBool will retrieve boolean field from data
func (i *Info) GetBool(m *MapParams, field string) (bool, error) {
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
func (i *Info) GetInt(m *MapParams, field string) (int, error) {
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
func (i *Info) GetUint(m *MapParams, field string) (uint, error) {
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
func (i *Info) GetFloat32(m *MapParams, field string) (float32, error) {
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
func (i *Info) GetFloat64(m *MapParams, field string) (float64, error) {
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
func (i *Info) GetString(m *MapParams, field string) (string, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return "", err
	}

	return value[0], nil
}

// GetStringArray will retrieve []string field from data
func (i *Info) GetStringArray(m *MapParams, field string) ([]string, error) {
	_, values, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	return values, nil
}

// GetIntArray will retrieve []int field from data
func (i *Info) GetIntArray(m *MapParams, field string) ([]int, error) {
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

// GetUintArray will retrieve []uint field from data
func (i *Info) GetUintArray(m *MapParams, field string) ([]uint, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	var uints []uint
	for i := 0; i < len(value); i++ {
		valueUint, err := strconv.ParseUint(value[i], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%s value could not parse to uint", value[i])
		}
		uints = append(uints, uint(valueUint))
	}

	return uints, nil
}

// GetFloat32Array will retrieve []float field from data
func (i *Info) GetFloat32Array(m *MapParams, field string) ([]float32, error) {
	_, value, err := i.GetField(m, field)
	if err != nil {
		return nil, err
	}

	var floats []float32
	for i := 0; i < len(value); i++ {
		valueFloat, err := strconv.ParseFloat(value[i], 32)
		if err != nil {
			return nil, fmt.Errorf("%s value could not parse to float", value[i])
		}
		floats = append(floats, float32(valueFloat))
	}

	return floats, nil
}
