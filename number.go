package gofakeit

import (
	"errors"
	"math"
	"math/rand"
)

// Number will generate a random number between given min And max
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Uint8 will generate a random uint8 value
func Uint8() uint8 {
	return uint8(randIntRange(0, math.MaxUint8))
}

// Uint16 will generate a random uint16 value
func Uint16() uint16 {
	return uint16(randIntRange(0, math.MaxUint16))
}

// Uint32 will generate a random uint32 value
func Uint32() uint32 {
	return uint32(randIntRange(0, math.MaxInt32))
}

// Uint64 will generate a random uint64 value
func Uint64() uint64 {
	return uint64(rand.Int63n(math.MaxInt64))
}

// Int8 will generate a random Int8 value
func Int8() int8 {
	return int8(randIntRange(math.MinInt8, math.MaxInt8))
}

// Int16 will generate a random int16 value
func Int16() int16 {
	return int16(randIntRange(math.MinInt16, math.MaxInt16))
}

// Int32 will generate a random int32 value
func Int32() int32 {
	return int32(randIntRange(math.MinInt32, math.MaxInt32))
}

// Int64 will generate a random int64 value
func Int64() int64 {
	return rand.Int63n(math.MaxInt64) + math.MinInt64
}

// Float32 will generate a random float32 value
func Float32() float32 {
	return randFloat32Range(math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float32Range will generate a random float32 value between min and max
func Float32Range(min, max float32) float32 {
	return randFloat32Range(min, max)
}

// Float64 will generate a random float64 value
func Float64() float64 {
	return randFloat64Range(math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float64Range will generate a random float64 value between min and max
func Float64Range(min, max float64) float64 {
	return randFloat64Range(min, max)
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// RandomInt will take in a slice of int and return a randomly selected value
func RandomInt(i []int) int {
	size := len(i)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return i[0]
	}
	return i[rand.Intn(size)]
}

// RandomUint will take in a slice of uint and return a randomly selected value
func RandomUint(i []uint) uint {
	size := len(i)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return i[0]
	}
	return i[rand.Intn(size)]
}

func addNumberLookup() {
	AddFuncLookup("number", Info{
		Display:     "Number",
		Category:    "number",
		Description: "Random number between given range",
		Example:     "14866",
		Output:      "int",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "int", Default: "-2147483648", Description: "Minimum integer value"},
			{Field: "max", Display: "Max", Type: "int", Default: "2147483647", Description: "Maximum integer value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetInt(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetInt(m, "max")
			if err != nil {
				return nil, err
			}

			if min > max {
				return nil, errors.New("Max integer must be larger than Min")
			}

			return Number(min, max), nil
		},
	})

	AddFuncLookup("uint8", Info{
		Display:     "Uint8",
		Category:    "number",
		Description: "Random uint8 value",
		Example:     "152",
		Output:      "uint8",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint8(), nil
		},
	})

	AddFuncLookup("uint16", Info{
		Display:     "Uint16",
		Category:    "number",
		Description: "Random uint16 value",
		Example:     "34968",
		Output:      "uint16",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint16(), nil
		},
	})

	AddFuncLookup("uint32", Info{
		Display:     "Uint32",
		Category:    "number",
		Description: "Random uint32 value",
		Example:     "1075055705",
		Output:      "uint32",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint32(), nil
		},
	})

	AddFuncLookup("uint64", Info{
		Display:     "Uint64",
		Category:    "number",
		Description: "Random uint64 value",
		Example:     "843730692693298265",
		Output:      "uint64",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Uint64(), nil
		},
	})

	AddFuncLookup("int8", Info{
		Display:     "Int8",
		Category:    "number",
		Description: "Random int8 value",
		Example:     "24",
		Output:      "int8",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int8(), nil
		},
	})

	AddFuncLookup("int16", Info{
		Display:     "Int16",
		Category:    "number",
		Description: "Random int16 value",
		Example:     "2200",
		Output:      "int16",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int16(), nil
		},
	})

	AddFuncLookup("int32", Info{
		Display:     "Int32",
		Category:    "number",
		Description: "Random int32 value",
		Example:     "-1072427943",
		Output:      "int32",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int32(), nil
		},
	})

	AddFuncLookup("int64", Info{
		Display:     "Int64",
		Category:    "number",
		Description: "Random int64 value",
		Example:     "-8379641344161477543",
		Output:      "int64",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Int64(), nil
		},
	})

	AddFuncLookup("float32", Info{
		Display:     "Float32",
		Category:    "number",
		Description: "Random float32 value",
		Example:     "3.1128167e+37",
		Output:      "float32",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Float32(), nil
		},
	})

	AddFuncLookup("float32range", Info{
		Display:     "Float32 Range",
		Category:    "number",
		Description: "Random float32 between given range",
		Example:     "914774.6",
		Output:      "float32",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "int", Description: "Minimum float32 value"},
			{Field: "max", Display: "Max", Type: "int", Description: "Maximum float32 value"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			min, err := info.GetFloat32(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat32(m, "max")
			if err != nil {
				return nil, err
			}

			return Float32Range(min, max), nil
		},
	})

	AddFuncLookup("float64", Info{
		Display:     "Float64",
		Category:    "number",
		Description: "Random float64 value",
		Example:     "1.644484108270445e+307",
		Output:      "float64",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Float64(), nil
		},
	})

	AddFuncLookup("float64range", Info{
		Display:     "Float64 Range",
		Category:    "number",
		Description: "Random float64 between given range",
		Example:     "914774.5585333086",
		Output:      "float64",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "int", Description: "Minimum float64 value"},
			{Field: "max", Display: "Max", Type: "int", Description: "Maximum float64 value"},
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

			return Float64Range(min, max), nil
		},
	})

	AddFuncLookup("shuffleints", Info{
		Display:     "Shuffle Ints",
		Category:    "number",
		Description: "Shuffle an array of ints",
		Example:     "1,2,3,4 => 3,1,4,2",
		Output:      "[]int",
		Params: []Param{
			{Field: "ints", Display: "Integers", Type: "[]int", Description: "Delimited separated integers"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			ints, err := info.GetIntArray(m, "ints")
			if err != nil {
				return nil, err
			}

			ShuffleInts(ints)

			return ints, nil
		},
	})
}
