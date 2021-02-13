package gofakeit

import (
	"errors"
	"math"
	"math/rand"
)

// Number will generate a random number between given min And max
func Number(min int, max int) int { return number(globalFaker.Rand, min, max) }

// Number will generate a random number between given min And max
func (f *Faker) Number(min int, max int) int { return number(f.Rand, min, max) }

func number(r *rand.Rand, min int, max int) int { return randIntRange(r, min, max) }

// Uint8 will generate a random uint8 value
func Uint8() uint8 { return uint8Func(globalFaker.Rand) }

// Uint8 will generate a random uint8 value
func (f *Faker) Uint8() uint8 { return uint8Func(f.Rand) }

func uint8Func(r *rand.Rand) uint8 { return uint8(randIntRange(r, 0, math.MaxUint8)) }

// Uint16 will generate a random uint16 value
func Uint16() uint16 { return uint16Func(globalFaker.Rand) }

// Uint16 will generate a random uint16 value
func (f *Faker) Uint16() uint16 { return uint16Func(f.Rand) }

func uint16Func(r *rand.Rand) uint16 { return uint16(randIntRange(r, 0, math.MaxUint16)) }

// Uint32 will generate a random uint32 value
func Uint32() uint32 { return uint32Func(globalFaker.Rand) }

// Uint32 will generate a random uint32 value
func (f *Faker) Uint32() uint32 { return uint32Func(f.Rand) }

func uint32Func(r *rand.Rand) uint32 { return uint32(randIntRange(r, 0, math.MaxInt32)) }

// Uint64 will generate a random uint64 value
func Uint64() uint64 { return uint64Func(globalFaker.Rand) }

// Uint64 will generate a random uint64 value
func (f *Faker) Uint64() uint64 { return uint64Func(f.Rand) }

func uint64Func(r *rand.Rand) uint64 { return uint64(r.Int63n(math.MaxInt64)) }

// Int8 will generate a random Int8 value
func Int8() int8 { return int8Func(globalFaker.Rand) }

// Int8 will generate a random Int8 value
func (f *Faker) Int8() int8 { return int8Func(f.Rand) }

func int8Func(r *rand.Rand) int8 { return int8(randIntRange(r, math.MinInt8, math.MaxInt8)) }

// Int16 will generate a random int16 value
func Int16() int16 { return int16Func(globalFaker.Rand) }

// Int16 will generate a random int16 value
func (f *Faker) Int16() int16 { return int16Func(f.Rand) }

func int16Func(r *rand.Rand) int16 { return int16(randIntRange(r, math.MinInt16, math.MaxInt16)) }

// Int32 will generate a random int32 value
func Int32() int32 { return int32Func(globalFaker.Rand) }

// Int32 will generate a random int32 value
func (f *Faker) Int32() int32 { return int32Func(f.Rand) }

func int32Func(r *rand.Rand) int32 { return int32(randIntRange(r, math.MinInt32, math.MaxInt32)) }

// Int64 will generate a random int64 value
func Int64() int64 { return int64Func(globalFaker.Rand) }

// Int64 will generate a random int64 value
func (f *Faker) Int64() int64 { return int64Func(f.Rand) }

func int64Func(r *rand.Rand) int64 { return r.Int63n(math.MaxInt64) + math.MinInt64 }

// Float32 will generate a random float32 value
func Float32() float32 { return float32Func(globalFaker.Rand) }

// Float32 will generate a random float32 value
func (f *Faker) Float32() float32 { return float32Func(f.Rand) }

func float32Func(r *rand.Rand) float32 {
	return float32Range(r, math.SmallestNonzeroFloat32, math.MaxFloat32)
}

// Float32Range will generate a random float32 value between min and max
func Float32Range(min, max float32) float32 {
	return float32Range(globalFaker.Rand, min, max)
}

// Float32Range will generate a random float32 value between min and max
func (f *Faker) Float32Range(min, max float32) float32 {
	return float32Range(f.Rand, min, max)
}

func float32Range(r *rand.Rand, min, max float32) float32 {
	if min == max {
		return min
	}
	return r.Float32()*(max-min) + min
}

// Float64 will generate a random float64 value
func Float64() float64 {
	return float64Func(globalFaker.Rand)
}

// Float64 will generate a random float64 value
func (f *Faker) Float64() float64 {
	return float64Func(f.Rand)
}

func float64Func(r *rand.Rand) float64 {
	return float64Range(r, math.SmallestNonzeroFloat64, math.MaxFloat64)
}

// Float64Range will generate a random float64 value between min and max
func Float64Range(min, max float64) float64 {
	return float64Range(globalFaker.Rand, min, max)
}

// Float64Range will generate a random float64 value between min and max
func (f *Faker) Float64Range(min, max float64) float64 {
	return float64Range(f.Rand, min, max)
}

func float64Range(r *rand.Rand, min, max float64) float64 {
	if min == max {
		return min
	}
	return r.Float64()*(max-min) + min
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) { shuffleInts(globalFaker.Rand, a) }

// ShuffleInts will randomize a slice of ints
func (f *Faker) ShuffleInts(a []int) { shuffleInts(f.Rand, a) }

func shuffleInts(r *rand.Rand, a []int) {
	for i := range a {
		j := r.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// RandomInt will take in a slice of int and return a randomly selected value
func RandomInt(i []int) int { return randomInt(globalFaker.Rand, i) }

// RandomInt will take in a slice of int and return a randomly selected value
func (f *Faker) RandomInt(i []int) int { return randomInt(f.Rand, i) }

func randomInt(r *rand.Rand, i []int) int {
	size := len(i)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return i[0]
	}
	return i[r.Intn(size)]
}

// RandomUint will take in a slice of uint and return a randomly selected value
func RandomUint(u []uint) uint { return randomUint(globalFaker.Rand, u) }

// RandomUint will take in a slice of uint and return a randomly selected value
func (f *Faker) RandomUint(u []uint) uint { return randomUint(f.Rand, u) }

func randomUint(r *rand.Rand, u []uint) uint {
	size := len(u)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return u[0]
	}
	return u[r.Intn(size)]
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
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

			return number(r, min, max), nil
		},
	})

	AddFuncLookup("uint8", Info{
		Display:     "Uint8",
		Category:    "number",
		Description: "Random uint8 value",
		Example:     "152",
		Output:      "uint8",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return uint8Func(r), nil
		},
	})

	AddFuncLookup("uint16", Info{
		Display:     "Uint16",
		Category:    "number",
		Description: "Random uint16 value",
		Example:     "34968",
		Output:      "uint16",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return uint16Func(r), nil
		},
	})

	AddFuncLookup("uint32", Info{
		Display:     "Uint32",
		Category:    "number",
		Description: "Random uint32 value",
		Example:     "1075055705",
		Output:      "uint32",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return uint32Func(r), nil
		},
	})

	AddFuncLookup("uint64", Info{
		Display:     "Uint64",
		Category:    "number",
		Description: "Random uint64 value",
		Example:     "843730692693298265",
		Output:      "uint64",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return uint64Func(r), nil
		},
	})

	AddFuncLookup("int8", Info{
		Display:     "Int8",
		Category:    "number",
		Description: "Random int8 value",
		Example:     "24",
		Output:      "int8",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return int8Func(r), nil
		},
	})

	AddFuncLookup("int16", Info{
		Display:     "Int16",
		Category:    "number",
		Description: "Random int16 value",
		Example:     "2200",
		Output:      "int16",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return int16Func(r), nil
		},
	})

	AddFuncLookup("int32", Info{
		Display:     "Int32",
		Category:    "number",
		Description: "Random int32 value",
		Example:     "-1072427943",
		Output:      "int32",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return int32Func(r), nil
		},
	})

	AddFuncLookup("int64", Info{
		Display:     "Int64",
		Category:    "number",
		Description: "Random int64 value",
		Example:     "-8379641344161477543",
		Output:      "int64",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return int64Func(r), nil
		},
	})

	AddFuncLookup("float32", Info{
		Display:     "Float32",
		Category:    "number",
		Description: "Random float32 value",
		Example:     "3.1128167e+37",
		Output:      "float32",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return float32Func(r), nil
		},
	})

	AddFuncLookup("float32range", Info{
		Display:     "Float32 Range",
		Category:    "number",
		Description: "Random float32 between given range",
		Example:     "914774.6",
		Output:      "float32",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Description: "Minimum float32 value"},
			{Field: "max", Display: "Max", Type: "float", Description: "Maximum float32 value"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			min, err := info.GetFloat32(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat32(m, "max")
			if err != nil {
				return nil, err
			}

			return float32Range(r, min, max), nil
		},
	})

	AddFuncLookup("float64", Info{
		Display:     "Float64",
		Category:    "number",
		Description: "Random float64 value",
		Example:     "1.644484108270445e+307",
		Output:      "float64",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return float64Func(r), nil
		},
	})

	AddFuncLookup("float64range", Info{
		Display:     "Float64 Range",
		Category:    "number",
		Description: "Random float64 between given range",
		Example:     "914774.5585333086",
		Output:      "float64",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Description: "Minimum float64 value"},
			{Field: "max", Display: "Max", Type: "float", Description: "Maximum float64 value"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			min, err := info.GetFloat64(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat64(m, "max")
			if err != nil {
				return nil, err
			}

			return float64Range(r, min, max), nil
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			ints, err := info.GetIntArray(m, "ints")
			if err != nil {
				return nil, err
			}

			shuffleInts(r, ints)

			return ints, nil
		},
	})
}
