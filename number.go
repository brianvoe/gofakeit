package gofakeit

import (
	"math"
)

// Number will generate a random number between given min and max
func Number(min int, max int) int { return number(GlobalFaker, min, max) }

// Number will generate a random number between given min and max
func (f *Faker) Number(min int, max int) int { return number(f, min, max) }

func number(f *Faker, min int, max int) int { return randIntRange(f, min, max) }

// Uint will generate a random uint value
func Uint() uint { return uintFunc(GlobalFaker) }

// Uint will generate a random uint value
func (f *Faker) Uint() uint { return uintFunc(f) }

func uintFunc(f *Faker) uint { return uint(f.Uint64()) }

// Uint8 will generate a random uint8 value
func Uint8() uint8 { return uint8Func(GlobalFaker) }

// Uint8 will generate a random uint8 value
func (f *Faker) Uint8() uint8 { return uint8Func(f) }

func uint8Func(f *Faker) uint8 { return uint8(randUintRange(f, minUint, math.MaxUint8)) }

// Uint16 will generate a random uint16 value
func Uint16() uint16 { return uint16Func(GlobalFaker) }

// Uint16 will generate a random uint16 value
func (f *Faker) Uint16() uint16 { return uint16Func(f) }

func uint16Func(f *Faker) uint16 { return uint16(randUintRange(f, minUint, math.MaxUint16)) }

// Uint32 will generate a random uint32 value
func Uint32() uint32 { return uint32Func(GlobalFaker) }

// Uint32 will generate a random uint32 value
func (f *Faker) Uint32() uint32 { return uint32Func(f) }

func uint32Func(f *Faker) uint32 { return uint32(f.Uint64() >> 32) }

// Uint32N will generate a random uint32 value between 0 and n
func Uint32N(n uint32) uint32 { return uint32NFunc(GlobalFaker, n) }

// Uint32N will generate a random uint32 value between 0 and n
func (f *Faker) Uint32N(n uint32) uint32 { return uint32NFunc(f, n) }

func uint32NFunc(f *Faker, n uint32) uint32 { return uint32(f.uint64n(uint64(n))) }

// Uint64 will generate a random uint64 value
func Uint64() uint64 { return uint64Func(GlobalFaker) }

// Uint64 will generate a random uint64 value
func (f *Faker) Uint64() uint64 { return uint64Func(f) }

func uint64Func(f *Faker) uint64 { return f.Uint64() }

// UintRange will generate a random uint value between min and max
func UintRange(min, max uint) uint { return uintRangeFunc(GlobalFaker, min, max) }

// UintRange will generate a random uint value between min and max
func (f *Faker) UintRange(min, max uint) uint { return uintRangeFunc(f, min, max) }

func uintRangeFunc(f *Faker, min, max uint) uint { return randUintRange(f, min, max) }

// Int will generate a random int value
func Int() int { return intFunc(GlobalFaker) }

// Int will generate a random int value
func (f *Faker) Int() int { return intFunc(f) }

func intFunc(f *Faker) int { return int(uint(f.Uint64()) << 1 >> 1) }

// IntN will generate a random int value between 0 and n
func IntN(n int) int { return intNFunc(GlobalFaker, n) }

// IntN will generate a random int value between 0 and n
func (f *Faker) IntN(n int) int { return intNFunc(f, n) }

func intNFunc(f *Faker, n int) int {
	if n <= 0 {
		panic("invalid argument to IntN")
	}
	return int(f.uint64n(uint64(n)))
}

// Int8 will generate a random Int8 value
func Int8() int8 { return int8Func(GlobalFaker) }

// Int8 will generate a random Int8 value
func (f *Faker) Int8() int8 { return int8Func(f) }

func int8Func(f *Faker) int8 { return int8(randIntRange(f, math.MinInt8, math.MaxInt8)) }

// Int16 will generate a random int16 value
func Int16() int16 { return int16Func(GlobalFaker) }

// Int16 will generate a random int16 value
func (f *Faker) Int16() int16 { return int16Func(f) }

func int16Func(f *Faker) int16 { return int16(randIntRange(f, math.MinInt16, math.MaxInt16)) }

// Int32 will generate a random int32 value
func Int32() int32 { return int32Func(GlobalFaker) }

// Int32 will generate a random int32 value
func (f *Faker) Int32() int32 { return int32Func(f) }

func int32Func(f *Faker) int32 { return int32(f.Uint64() >> 33) }

// Int64 will generate a random int64 value
func Int64() int64 { return int64Func(GlobalFaker) }

// Int64 will generate a random int64 value
func (f *Faker) Int64() int64 { return int64Func(f) }

func int64Func(f *Faker) int64 { return int64(f.Uint64() &^ (1 << 63)) }

// IntRange will generate a random int value between min and max
func IntRange(min, max int) int { return intRangeFunc(GlobalFaker, min, max) }

// IntRange will generate a random int value between min and max
func (f *Faker) IntRange(min, max int) int { return intRangeFunc(f, min, max) }

func intRangeFunc(f *Faker, min, max int) int { return randIntRange(f, min, max) }

// Float32 will generate a random float32 value
func Float32() float32 { return float32Func(GlobalFaker) }

// Float32 will generate a random float32 value
func (f *Faker) Float32() float32 { return float32Func(f) }

func float32Func(f *Faker) float32 {
	// There are exactly 1<<24 float32s in [0,1). Use Intn(1<<24) / (1<<24).
	return float32(f.Uint32()<<8>>8) / (1 << 24)
}

// Float32Range will generate a random float32 value between min and max
func Float32Range(min, max float32) float32 {
	return float32Range(GlobalFaker, min, max)
}

// Float32Range will generate a random float32 value between min and max
func (f *Faker) Float32Range(min, max float32) float32 {
	return float32Range(f, min, max)
}

func float32Range(f *Faker, min, max float32) float32 {
	if min == max {
		return min
	}
	return f.Float32()*(max-min) + min
}

// Float64 will generate a random float64 value
func Float64() float64 {
	return float64Func(GlobalFaker)
}

// Float64 will generate a random float64 value
func (f *Faker) Float64() float64 {
	return float64Func(f)
}

func float64Func(f *Faker) float64 {
	// There are exactly 1<<53 float64s in [0,1). Use Intn(1<<53) / (1<<53).
	return float64(f.Uint64()<<11>>11) / (1 << 53)
}

// Float64Range will generate a random float64 value between min and max
func Float64Range(min, max float64) float64 {
	return float64Range(GlobalFaker, min, max)
}

// Float64Range will generate a random float64 value between min and max
func (f *Faker) Float64Range(min, max float64) float64 {
	return float64Range(f, min, max)
}

func float64Range(f *Faker, min, max float64) float64 {
	if min == max {
		return min
	}
	return f.Float64()*(max-min) + min
}

// ShuffleInts will randomize a slice of ints
func ShuffleInts(a []int) { shuffleInts(GlobalFaker, a) }

// ShuffleInts will randomize a slice of ints
func (f *Faker) ShuffleInts(a []int) { shuffleInts(f, a) }

func shuffleInts(f *Faker, a []int) {
	for i := range a {
		j := f.IntN(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// RandomInt will take in a slice of int and return a randomly selected value
func RandomInt(i []int) int { return randomInt(GlobalFaker, i) }

// RandomInt will take in a slice of int and return a randomly selected value
func (f *Faker) RandomInt(i []int) int { return randomInt(f, i) }

func randomInt(f *Faker, i []int) int {
	size := len(i)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return i[0]
	}
	return i[f.IntN(size)]
}

// RandomUint will take in a slice of uint and return a randomly selected value
func RandomUint(u []uint) uint { return randomUint(GlobalFaker, u) }

// RandomUint will take in a slice of uint and return a randomly selected value
func (f *Faker) RandomUint(u []uint) uint { return randomUint(f, u) }

func randomUint(f *Faker, u []uint) uint {
	size := len(u)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return u[0]
	}
	return u[f.IntN(size)]
}

// HexUint8 will generate a random uint8 hex value with "0x" prefix
func HexUint8() string { return hexUint(GlobalFaker, 8) }

// HexUint8 will generate a random uint8 hex value with "0x" prefix
func (f *Faker) HexUint8() string { return hexUint(f, 8) }

// HexUint16 will generate a random uint16 hex value with "0x" prefix
func HexUint16() string { return hexUint(GlobalFaker, 16) }

// HexUint16 will generate a random uint16 hex value with "0x" prefix
func (f *Faker) HexUint16() string { return hexUint(f, 16) }

// HexUint32 will generate a random uint32 hex value with "0x" prefix
func HexUint32() string { return hexUint(GlobalFaker, 32) }

// HexUint32 will generate a random uint32 hex value with "0x" prefix
func (f *Faker) HexUint32() string { return hexUint(f, 32) }

// HexUint64 will generate a random uint64 hex value with "0x" prefix
func HexUint64() string { return hexUint(GlobalFaker, 64) }

// HexUint64 will generate a random uint64 hex value with "0x" prefix
func (f *Faker) HexUint64() string { return hexUint(f, 64) }

// HexUint128 will generate a random uint128 hex value with "0x" prefix
func HexUint128() string { return hexUint(GlobalFaker, 128) }

// HexUint128 will generate a random uint128 hex value with "0x" prefix
func (f *Faker) HexUint128() string { return hexUint(f, 128) }

// HexUint256 will generate a random uint256 hex value with "0x" prefix
func HexUint256() string { return hexUint(GlobalFaker, 256) }

// HexUint256 will generate a random uint256 hex value with "0x" prefix
func (f *Faker) HexUint256() string { return hexUint(f, 256) }

func hexUint(f *Faker, bitSize int) string {
	digits := []byte("0123456789abcdef")
	hexLen := (bitSize >> 2) + 2
	if hexLen <= 2 {
		return "0x"
	}

	s := make([]byte, hexLen)
	s[0], s[1] = '0', 'x'
	for i := 2; i < hexLen; i++ {
		s[i] = digits[f.IntN(16)]
	}
	return string(s)
}

func addNumberLookup() {
	AddFuncLookup("number", Info{
		Display:     "Number",
		Category:    "number",
		Description: "Mathematical concept used for counting, measuring, and expressing quantities or values",
		Example:     "14866",
		Output:      "int",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "int", Default: "-2147483648", Description: "Minimum integer value"},
			{Field: "max", Display: "Max", Type: "int", Default: "2147483647", Description: "Maximum integer value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetInt(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetInt(m, "max")
			if err != nil {
				return nil, err
			}

			return number(f, min, max), nil
		},
	})

	AddFuncLookup("uint8", Info{
		Display:     "Uint8",
		Category:    "number",
		Description: "Unsigned 8-bit integer, capable of representing values from 0 to 255",
		Example:     "152",
		Output:      "uint8",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return uint8Func(f), nil
		},
	})

	AddFuncLookup("uint16", Info{
		Display:     "Uint16",
		Category:    "number",
		Description: "Unsigned 16-bit integer, capable of representing values from 0 to 65,535",
		Example:     "34968",
		Output:      "uint16",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return uint16Func(f), nil
		},
	})

	AddFuncLookup("uint32", Info{
		Display:     "Uint32",
		Category:    "number",
		Description: "Unsigned 32-bit integer, capable of representing values from 0 to 4,294,967,295",
		Example:     "1075055705",
		Output:      "uint32",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return uint32Func(f), nil
		},
	})

	AddFuncLookup("uint64", Info{
		Display:     "Uint64",
		Category:    "number",
		Description: "Unsigned 64-bit integer, capable of representing values from 0 to 18,446,744,073,709,551,615",
		Example:     "843730692693298265",
		Output:      "uint64",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return uint64Func(f), nil
		},
	})

	AddFuncLookup("uintrange", Info{
		Display:     "UintRange",
		Category:    "number",
		Description: "Non-negative integer value between given range",
		Example:     "1075055705",
		Output:      "uint",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "uint", Default: "0", Description: "Minimum uint value"},
			{Field: "max", Display: "Max", Type: "uint", Default: "4294967295", Description: "Maximum uint value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetUint(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetUint(m, "max")
			if err != nil {
				return nil, err
			}

			return uintRangeFunc(f, min, max), nil
		},
	})

	AddFuncLookup("int8", Info{
		Display:     "Int8",
		Category:    "number",
		Description: "Signed 8-bit integer, capable of representing values from -128 to 127",
		Example:     "24",
		Output:      "int8",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return int8Func(f), nil
		},
	})

	AddFuncLookup("int16", Info{
		Display:     "Int16",
		Category:    "number",
		Description: "Signed 16-bit integer, capable of representing values from 32,768 to 32,767",
		Example:     "2200",
		Output:      "int16",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return int16Func(f), nil
		},
	})

	AddFuncLookup("int32", Info{
		Display:     "Int32",
		Category:    "number",
		Description: "Signed 32-bit integer, capable of representing values from -2,147,483,648 to 2,147,483,647",
		Example:     "-1072427943",
		Output:      "int32",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return int32Func(f), nil
		},
	})

	AddFuncLookup("int64", Info{
		Display:     "Int64",
		Category:    "number",
		Description: "Signed 64-bit integer, capable of representing values from -9,223,372,036,854,775,808 to -9,223,372,036,854,775,807",
		Example:     "-8379641344161477543",
		Output:      "int64",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return int64Func(f), nil
		},
	})

	AddFuncLookup("intrange", Info{
		Display:     "IntRange",
		Category:    "number",
		Description: "Integer value between given range",
		Example:     "-8379477543",
		Output:      "int",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "int", Description: "Minimum int value"},
			{Field: "max", Display: "Max", Type: "int", Description: "Maximum int value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetInt(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetInt(m, "max")
			if err != nil {
				return nil, err
			}

			return intRangeFunc(f, min, max), nil
		},
	})

	AddFuncLookup("float32", Info{
		Display:     "Float32",
		Category:    "number",
		Description: "Data type representing floating-point numbers with 32 bits of precision in computing",
		Example:     "3.1128167e+37",
		Output:      "float32",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return float32Func(f), nil
		},
	})

	AddFuncLookup("float32range", Info{
		Display:     "Float32 Range",
		Category:    "number",
		Description: "Float32 value between given range",
		Example:     "914774.6",
		Output:      "float32",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Description: "Minimum float32 value"},
			{Field: "max", Display: "Max", Type: "float", Description: "Maximum float32 value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetFloat32(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat32(m, "max")
			if err != nil {
				return nil, err
			}

			return float32Range(f, min, max), nil
		},
	})

	AddFuncLookup("float64", Info{
		Display:     "Float64",
		Category:    "number",
		Description: "Data type representing floating-point numbers with 64 bits of precision in computing",
		Example:     "1.644484108270445e+307",
		Output:      "float64",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return float64Func(f), nil
		},
	})

	AddFuncLookup("float64range", Info{
		Display:     "Float64 Range",
		Category:    "number",
		Description: "Float64 value between given range",
		Example:     "914774.5585333086",
		Output:      "float64",
		Params: []Param{
			{Field: "min", Display: "Min", Type: "float", Description: "Minimum float64 value"},
			{Field: "max", Display: "Max", Type: "float", Description: "Maximum float64 value"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			min, err := info.GetFloat64(m, "min")
			if err != nil {
				return nil, err
			}

			max, err := info.GetFloat64(m, "max")
			if err != nil {
				return nil, err
			}

			return float64Range(f, min, max), nil
		},
	})

	AddFuncLookup("shuffleints", Info{
		Display:     "Shuffle Ints",
		Category:    "number",
		Description: "Shuffles an array of ints",
		Example:     "1,2,3,4 => 3,1,4,2",
		Output:      "[]int",
		Params: []Param{
			{Field: "ints", Display: "Integers", Type: "[]int", Description: "Delimited separated integers"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			ints, err := info.GetIntArray(m, "ints")
			if err != nil {
				return nil, err
			}

			shuffleInts(f, ints)

			return ints, nil
		},
	})

	AddFuncLookup("randomint", Info{
		Display:     "Random Int",
		Category:    "number",
		Description: "Randomly selected value from a slice of int",
		Example:     "-1,2,-3,4 => -3",
		Output:      "int",
		Params: []Param{
			{Field: "ints", Display: "Integers", Type: "[]int", Description: "Delimited separated integers"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			ints, err := info.GetIntArray(m, "ints")
			if err != nil {
				return nil, err
			}

			return randomInt(f, ints), nil
		},
	})

	AddFuncLookup("randomuint", Info{
		Display:     "Random Uint",
		Category:    "number",
		Description: "Randomly selected value from a slice of uint",
		Example:     "1,2,3,4 => 4",
		Output:      "uint",
		Params: []Param{
			{Field: "uints", Display: "Unsigned Integers", Type: "[]uint", Description: "Delimited separated unsigned integers"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			uints, err := info.GetUintArray(m, "uints")
			if err != nil {
				return nil, err
			}

			return randomUint(f, uints), nil
		},
	})

	AddFuncLookup("hexuint8", Info{
		Display:     "HexUint8",
		Category:    "number",
		Description: "Hexadecimal representation of an 8-bit unsigned integer",
		Example:     "0x87",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 8), nil
		},
	})

	AddFuncLookup("hexuint16", Info{
		Display:     "HexUint16",
		Category:    "number",
		Description: "Hexadecimal representation of an 16-bit unsigned integer",
		Example:     "0x8754",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 16), nil
		},
	})

	AddFuncLookup("hexuint32", Info{
		Display:     "HexUint32",
		Category:    "number",
		Description: "Hexadecimal representation of an 32-bit unsigned integer",
		Example:     "0x87546957",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 32), nil
		},
	})

	AddFuncLookup("hexuint64", Info{
		Display:     "HexUint64",
		Category:    "number",
		Description: "Hexadecimal representation of an 64-bit unsigned integer",
		Example:     "0x875469578e51b5e5",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 64), nil
		},
	})

	AddFuncLookup("hexuint128", Info{
		Display:     "HexUint128",
		Category:    "number",
		Description: "Hexadecimal representation of an 128-bit unsigned integer",
		Example:     "0x875469578e51b5e56c95b64681d147a1",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 128), nil
		},
	})

	AddFuncLookup("hexuint256", Info{
		Display:     "HexUint256",
		Category:    "number",
		Description: "Hexadecimal representation of an 256-bit unsigned integer",
		Example:     "0x875469578e51b5e56c95b64681d147a12cde48a4f417231b0c486abbc263e48d",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hexUint(f, 256), nil
		},
	})
}
