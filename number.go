package gofakeit

import (
	"math"
	"math/bits"
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

// UintN will generate a random uint value between 0 and n
func UintN(n uint) uint { return uintNFunc(GlobalFaker, n) }

// UintN will generate a random uint value between 0 and n
func (f *Faker) UintN(n uint) uint { return uintNFunc(f, n) }

func uintNFunc(f *Faker, n uint) uint {
	if n == 0 {
		return 0
	}
	return uint(uint64NFunc(f, uint64(n)))
}

// Uint8 will generate a random uint8 value
func Uint8() uint8 { return uint8Func(GlobalFaker) }

// Uint8 will generate a random uint8 value
func (f *Faker) Uint8() uint8 { return uint8Func(f) }

func uint8Func(f *Faker) uint8 { return uint8(randIntRange(f, minUint, math.MaxUint8)) }

// Uint16 will generate a random uint16 value
func Uint16() uint16 { return uint16Func(GlobalFaker) }

// Uint16 will generate a random uint16 value
func (f *Faker) Uint16() uint16 { return uint16Func(f) }

func uint16Func(f *Faker) uint16 { return uint16(randIntRange(f, minUint, math.MaxUint16)) }

// Uint32 will generate a random uint32 value
func Uint32() uint32 { return uint32Func(GlobalFaker) }

// Uint32 will generate a random uint32 value
func (f *Faker) Uint32() uint32 { return uint32Func(f) }

func uint32Func(f *Faker) uint32 { return uint32(f.Uint64() >> 32) }

// Uint64 will generate a random uint64 value
func Uint64() uint64 { return GlobalFaker.Uint64() }

// Uint64 will generate a random uint64 value
// This is the primary location in which the random number is generated.
// This will be the only location in which reading from Rand.Uint64() is lockable
func (f *Faker) Uint64() uint64 {
	// Check if the source is locked
	if f.Locked {
		// Lock the source
		f.mu.Lock()
		defer f.mu.Unlock()
	}

	return f.Rand.Uint64()
}

// uint64n is the no-bounds-checks version of Uint64N.
// See https://cs.opensource.google/go/go/+/refs/tags/go1.22.0:src/math/rand/v2/rand.go;l=78
// hidden as to not clutter with additional N functions
func uint64NFunc(f *Faker, n uint64) uint64 {
	if is32bit && uint64(uint32(n)) == n {
		// create reusable function here
		uint32NFunc := func(f *Faker, n uint32) uint32 {
			if n&(n-1) == 0 { // n is power of two, can mask
				return uint32(f.Uint64()) & (n - 1)
			}

			x := f.Uint64()
			lo1a, lo0 := bits.Mul32(uint32(x), n)
			hi, lo1b := bits.Mul32(uint32(x>>32), n)
			lo1, c := bits.Add32(lo1a, lo1b, 0)
			hi += c
			if lo1 == 0 && lo0 < uint32(n) {
				n64 := uint64(n)
				thresh := uint32(-n64 % n64)
				for lo1 == 0 && lo0 < thresh {
					x := f.Uint64()
					lo1a, lo0 = bits.Mul32(uint32(x), n)
					hi, lo1b = bits.Mul32(uint32(x>>32), n)
					lo1, c = bits.Add32(lo1a, lo1b, 0)
					hi += c
				}
			}
			return hi
		}

		return uint64(uint32NFunc(f, uint32(n)))
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return f.Uint64() & (n - 1)
	}

	hi, lo := bits.Mul64(f.Uint64(), n)
	if lo < n {
		thresh := -n % n
		for lo < thresh {
			hi, lo = bits.Mul64(f.Uint64(), n)
		}
	}
	return hi
}

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
		return 0
	}
	return int(uint64NFunc(f, uint64(n)))
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

// int32n is an identical computation to int64n
// hidden as to not clutter with additional N functions
func int32NFunc(f *Faker, n int32) int32 {
	if n <= 0 {
		return 0
	}
	return int32(uint64NFunc(f, uint64(n)))
}

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

// HexUint will generate a random uint hex value with "0x" prefix
func HexUint(bitSize int) string { return hexUint(GlobalFaker, bitSize) }

// HexUint will generate a random uint hex value with "0x" prefix
func (f *Faker) HexUint(bitSize int) string { return hexUint(f, bitSize) }

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
	// integer with min/max
	AddFuncLookup("number", Info{
		Display:     "Number",
		Category:    "number",
		Description: "Integer used for counting or measuring, with optional bounds",
		Example:     "14866",
		Output:      "int",
		Aliases: []string{
			"integer number",
			"random integer",
			"bounded integer",
			"range integer",
		},
		Keywords: []string{
			"number", "integer", "int", "random",
			"min", "max", "range", "bounded", "between", "inclusive",
		},
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

	// unsigned integer
	AddFuncLookup("uint", Info{
		Display:     "Uint",
		Category:    "number",
		Description: "Unsigned integer (nonnegative whole number)",
		Example:     "14866",
		Output:      "uint",
		Aliases: []string{
			"unsigned integer",
			"nonnegative integer",
			"natural number",
		},
		Keywords: []string{
			"uint", "unsigned", "integer", "nonnegative",
			"natural", "zero", "positive", "whole",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return uintFunc(f), nil },
	})

	// unsigned integer in [0,n)
	AddFuncLookup("uintn", Info{
		Display:     "UintN",
		Category:    "number",
		Description: "Unsigned integer between 0 (inclusive) and n (exclusive)",
		Example:     "32783",
		Output:      "uint",
		Aliases: []string{
			"bounded uint",
			"range uint",
			"uint up to n",
		},
		Keywords: []string{
			"uintn", "unsigned", "range", "upper",
			"limit", "bound", "cap", "max", "exclusive",
		},
		Params: []Param{
			{Field: "n", Display: "N", Type: "uint", Default: "4294967295", Description: "Maximum uint value (exclusive)"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			n, err := info.GetUint(m, "n")
			if err != nil {
				return nil, err
			}
			return uintNFunc(f, n), nil
		},
	})

	// 8-bit unsigned
	AddFuncLookup("uint8", Info{
		Display:     "Uint8",
		Category:    "number",
		Description: "Unsigned 8-bit integer, range 0–255",
		Example:     "152",
		Output:      "uint8",
		Aliases: []string{
			"8-bit unsigned integer",
			"byte value",
		},
		Keywords: []string{
			"uint8", "unsigned", "8bit", "0-255",
			"byte", "octet", "range",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return uint8Func(f), nil },
	})

	// 16-bit unsigned
	AddFuncLookup("uint16", Info{
		Display:     "Uint16",
		Category:    "number",
		Description: "Unsigned 16-bit integer, range 0–65,535",
		Example:     "34968",
		Output:      "uint16",
		Aliases: []string{
			"16-bit unsigned integer",
			"word value",
		},
		Keywords: []string{
			"uint16", "unsigned", "16bit", "0-65535",
			"word", "port", "range",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return uint16Func(f), nil },
	})

	// 32-bit unsigned
	AddFuncLookup("uint32", Info{
		Display:     "Uint32",
		Category:    "number",
		Description: "Unsigned 32-bit integer, range 0–4,294,967,295",
		Example:     "1075055705",
		Output:      "uint32",
		Aliases: []string{
			"32-bit unsigned integer",
			"long unsigned",
		},
		Keywords: []string{
			"uint32", "unsigned", "32bit", "0-4294967295",
			"range", "ipv4",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return uint32Func(f), nil },
	})

	// 64-bit unsigned
	AddFuncLookup("uint64", Info{
		Display:     "Uint64",
		Category:    "number",
		Description: "Unsigned 64-bit integer, range 0–18,446,744,073,709,551,615",
		Example:     "843730692693298265",
		Output:      "uint64",
		Aliases: []string{
			"64-bit unsigned integer",
			"big unsigned",
		},
		Keywords: []string{
			"uint64", "unsigned", "64bit",
			"0-18446744073709551615", "range", "bigint",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return f.Uint64(), nil },
	})

	// unsigned in [min,max]
	AddFuncLookup("uintrange", Info{
		Display:     "Uint Range",
		Category:    "number",
		Description: "Unsigned integer value within a given range",
		Example:     "1075055705",
		Output:      "uint",
		Aliases: []string{
			"uint range",
			"bounded uint range",
			"unsigned interval",
		},
		Keywords: []string{
			"uintrange", "unsigned", "range", "min", "max",
			"bounds", "limits", "interval", "span",
		},
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

	// signed integer (any)
	AddFuncLookup("int", Info{
		Display:     "Int",
		Category:    "number",
		Description: "Signed integer",
		Example:     "14866",
		Output:      "int",
		Aliases: []string{
			"signed integer",
			"whole number",
		},
		Keywords: []string{
			"int", "signed", "integer",
			"positive", "negative", "zero", "counting",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return intFunc(f), nil },
	})

	// signed integer in [0,n)
	AddFuncLookup("intn", Info{
		Display:     "IntN",
		Category:    "number",
		Description: "Integer between 0 (inclusive) and n (exclusive)",
		Example:     "32783",
		Output:      "int",
		Aliases: []string{
			"bounded int",
			"range int",
			"int up to n",
		},
		Keywords: []string{
			"intn", "range", "upper", "limit", "bound",
			"cap", "max", "exclusive", "integer",
		},
		Params: []Param{
			{Field: "n", Display: "N", Type: "int", Default: "2147483647", Description: "Maximum int value (exclusive)"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			n, err := info.GetInt(m, "n")
			if err != nil {
				return nil, err
			}
			return intNFunc(f, n), nil
		},
	})

	// 8-bit signed
	AddFuncLookup("int8", Info{
		Display:     "Int8",
		Category:    "number",
		Description: "Signed 8-bit integer, range −128–127",
		Example:     "24",
		Output:      "int8",
		Aliases: []string{
			"8-bit signed integer",
			"small signed",
		},
		Keywords: []string{
			"int8", "signed", "8bit", "-128..127", "range",
			"twoscomplement",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return int8Func(f), nil },
	})

	// 16-bit signed
	AddFuncLookup("int16", Info{
		Display:     "Int16",
		Category:    "number",
		Description: "Signed 16-bit integer, range −32,768–32,767",
		Example:     "2200",
		Output:      "int16",
		Aliases: []string{
			"16-bit signed integer",
			"short signed",
		},
		Keywords: []string{
			"int16", "signed", "16bit", "-32768..32767", "range",
			"word",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return int16Func(f), nil },
	})

	// 32-bit signed
	AddFuncLookup("int32", Info{
		Display:     "Int32",
		Category:    "number",
		Description: "Signed 32-bit integer, range −2,147,483,648–2,147,483,647",
		Example:     "-1072427943",
		Output:      "int32",
		Aliases: []string{
			"32-bit signed integer",
			"standard signed",
		},
		Keywords: []string{
			"int32", "signed", "32bit", "-2147483648..2147483647",
			"range", "ipv4",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return int32Func(f), nil },
	})

	// 64-bit signed
	AddFuncLookup("int64", Info{
		Display:     "Int64",
		Category:    "number",
		Description: "Signed 64-bit integer, range −9,223,372,036,854,775,808–9,223,372,036,854,775,807",
		Example:     "-8379641344161477543",
		Output:      "int64",
		Aliases: []string{
			"64-bit signed integer",
			"long signed",
		},
		Keywords: []string{
			"int64", "signed", "64bit", "bigint", "range",
			"timestamp", "nanosecond",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return int64Func(f), nil },
	})

	// signed in [min,max]
	AddFuncLookup("intrange", Info{
		Display:     "Int Range",
		Category:    "number",
		Description: "Signed integer value within a given range",
		Example:     "-8379477543",
		Output:      "int",
		Aliases: []string{
			"int range",
			"bounded int range",
			"signed interval",
		},
		Keywords: []string{
			"intrange", "int", "range", "min", "max",
			"bounds", "limits", "interval", "span",
		},
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

	// float32
	AddFuncLookup("float32", Info{
		Display:     "Float32",
		Category:    "number",
		Description: "Floating-point number with 32-bit single precision (IEEE 754)",
		Example:     "3.1128167e+37",
		Output:      "float32",
		Aliases: []string{
			"single precision",
			"fp32",
		},
		Keywords: []string{
			"float32", "single-precision", "ieee754",
			"fp32", "mantissa", "exponent", "decimal",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return float32Func(f), nil },
	})

	// float32 in [min,max]
	AddFuncLookup("float32range", Info{
		Display:     "Float32 Range",
		Category:    "number",
		Description: "Float32 value within a given range",
		Example:     "914774.6",
		Output:      "float32",
		Aliases: []string{
			"float32 range",
			"bounded float32",
			"single precision range",
		},
		Keywords: []string{
			"float32range", "float32", "range",
			"min", "max", "bounds", "limits", "interval",
		},
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

	// float64
	AddFuncLookup("float64", Info{
		Display:     "Float64",
		Category:    "number",
		Description: "Floating-point number with 64-bit double precision (IEEE 754)",
		Example:     "1.644484108270445e+307",
		Output:      "float64",
		Aliases: []string{
			"double precision",
			"fp64",
		},
		Keywords: []string{
			"float64", "double-precision", "ieee754",
			"fp64", "mantissa", "exponent", "decimal",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) { return float64Func(f), nil },
	})

	// float64 in [min,max]
	AddFuncLookup("float64range", Info{
		Display:     "Float64 Range",
		Category:    "number",
		Description: "Float64 value within a given range",
		Example:     "914774.5585333086",
		Output:      "float64",
		Aliases: []string{
			"float64 range",
			"bounded float64",
			"double precision range",
		},
		Keywords: []string{
			"float64range", "float64", "range",
			"min", "max", "bounds", "limits", "interval",
		},
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

	// shuffle ints
	AddFuncLookup("shuffleints", Info{
		Display:     "Shuffle Ints",
		Category:    "number",
		Description: "Shuffles an array of ints",
		Example:     "1,2,3,4 => 3,1,4,2",
		Output:      "[]int",
		Aliases: []string{
			"shuffle integers",
			"randomize ints",
			"permute ints",
		},
		Keywords: []string{
			"shuffleints", "shuffle", "permute", "randomize",
			"ints", "slice", "array", "permutation",
		},
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

	// pick random int from slice
	AddFuncLookup("randomint", Info{
		Display:     "Random Int",
		Category:    "number",
		Description: "Randomly selected value from a slice of int",
		Example:     "-1,2,-3,4 => -3",
		Output:      "int",
		Aliases: []string{
			"pick random int",
			"choose int",
			"select int",
		},
		Keywords: []string{
			"randomint", "random", "pick", "choose",
			"select", "ints", "slice", "list",
		},
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

	// pick random uint from slice
	AddFuncLookup("randomuint", Info{
		Display:     "Random Uint",
		Category:    "number",
		Description: "Randomly selected value from a slice of uint",
		Example:     "1,2,3,4 => 4",
		Output:      "uint",
		Aliases: []string{
			"pick random uint",
			"choose uint",
			"select uint",
		},
		Keywords: []string{
			"randomuint", "random", "pick", "choose",
			"select", "uints", "slice", "list", "nonnegative",
		},
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

	// hex representation of unsigned int
	AddFuncLookup("hexuint", Info{
		Display:     "HexUint",
		Category:    "number",
		Description: "Hexadecimal representation of an unsigned integer",
		Example:     "0x87",
		Output:      "string",
		Aliases: []string{
			"hexadecimal uint",
			"unsigned hex",
			"base16 uint",
		},
		Keywords: []string{
			"hexuint", "hex", "base16", "uint", "0x",
			"bits", "width", "format",
		},
		Params: []Param{
			{Field: "bitSize", Display: "Bit Size", Type: "int", Default: "8", Description: "Bit size of the unsigned integer"},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			bitSize, err := info.GetInt(m, "bitSize")
			if err != nil {
				return nil, err
			}
			return hexUint(f, bitSize), nil
		},
	})

}
