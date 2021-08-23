package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleNumber() {
	Seed(11)
	fmt.Println(Number(50, 23456))
	// Output: 14866
}

func ExampleFaker_Number() {
	f := New(11)
	fmt.Println(f.Number(50, 23456))
	// Output: 14866
}

func BenchmarkNumber(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Number(10, 999999)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Number(10, 999999)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Number(10, 999999)
		}
	})
}

func ExampleUint8() {
	Seed(11)
	fmt.Println(Uint8())
	// Output: 152
}

func ExampleFaker_Uint8() {
	f := New(11)
	fmt.Println(f.Uint8())
	// Output: 152
}

func BenchmarkUint8(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Uint8()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Uint8()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Uint8()
		}
	})
}

func ExampleUint16() {
	Seed(11)
	fmt.Println(Uint16())
	// Output: 34968
}

func ExampleFaker_Uint16() {
	f := New(11)
	fmt.Println(f.Uint16())
	// Output: 34968
}

func BenchmarkUint16(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Uint16()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Uint16()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Uint16()
		}
	})
}

func ExampleUint32() {
	Seed(11)
	fmt.Println(Uint32())
	// Output: 1075055705
}

func ExampleFaker_Uint32() {
	f := New(11)
	fmt.Println(f.Uint32())
	// Output: 1075055705
}

func BenchmarkUint32(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Uint32()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Uint32()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Uint32()
		}
	})
}

func ExampleUint64() {
	Seed(11)
	fmt.Println(Uint64())
	// Output: 843730692693298265
}

func ExampleFaker_Uint64() {
	f := New(11)
	fmt.Println(f.Uint64())
	// Output: 843730692693298265
}

func BenchmarkUint64(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Uint64()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Uint64()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Uint64()
		}
	})
}

func ExampleInt8() {
	Seed(11)
	fmt.Println(Int8())
	// Output: 24
}

func ExampleFaker_Int8() {
	f := New(11)
	fmt.Println(f.Int8())
	// Output: 24
}

func BenchmarkInt8(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Int8()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Int8()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Int8()
		}
	})
}

func ExampleInt16() {
	Seed(11)
	fmt.Println(Int16())
	// Output: 2200
}

func ExampleFaker_Int16() {
	f := New(11)
	fmt.Println(f.Int16())
	// Output: 2200
}

func BenchmarkInt16(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Int16()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Int16()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Int16()
		}
	})
}

func ExampleInt32() {
	Seed(11)
	fmt.Println(Int32())
	// Output: -1072427943
}

func ExampleFaker_Int32() {
	f := New(11)
	fmt.Println(f.Int32())
	// Output: -1072427943
}

func BenchmarkInt32(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Int32()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Int32()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Int32()
		}
	})
}

func ExampleInt64() {
	Seed(11)
	fmt.Println(Int64())
	// Output: -8379641344161477543
}

func ExampleFaker_Int64() {
	f := New(11)
	fmt.Println(f.Int64())
	// Output: -8379641344161477543
}

func BenchmarkInt64(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Int64()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Int64()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Int64()
		}
	})
}

func ExampleFloat32() {
	Seed(11)
	fmt.Println(Float32())
	// Output: 3.1128167e+37
}

func ExampleFaker_Float32() {
	f := New(11)
	fmt.Println(f.Float32())
	// Output: 3.1128167e+37
}

func BenchmarkFloat32(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Float32()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Float32()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Float32()
		}
	})
}

func ExampleFloat32Range() {
	Seed(11)
	fmt.Println(Float32Range(0, 9999999))
	// Output: 914774.6
}

func ExampleFaker_Float32Range() {
	f := New(11)
	fmt.Println(f.Float32Range(0, 9999999))
	// Output: 914774.6
}

func BenchmarkFloat32Range(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Float32Range(10, 999999)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Float32Range(10, 999999)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Float32Range(10, 999999)
		}
	})
}

func TestFloat32RangeSame(t *testing.T) {
	if float32Range(globalFaker.Rand, 5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func ExampleFloat64() {
	Seed(11)
	fmt.Println(Float64())
	// Output: 1.644484108270445e+307
}

func ExampleFaker_Float64() {
	f := New(11)
	fmt.Println(f.Float64())
	// Output: 1.644484108270445e+307
}

func BenchmarkFloat64(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Float64()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Float64()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Float64()
		}
	})
}

func ExampleFloat64Range() {
	Seed(11)
	fmt.Println(Float64Range(0, 9999999))
	// Output: 914774.5585333086
}

func ExampleFaker_Float64Range() {
	f := New(11)
	fmt.Println(f.Float64Range(0, 9999999))
	// Output: 914774.5585333086
}

func BenchmarkFloat64Range(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Float64Range(0, 999999)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Float64Range(0, 999999)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Float64Range(0, 999999)
		}
	})
}

func TestRandFloat64RangeSame(t *testing.T) {
	if float64Range(globalFaker.Rand, 5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func ExampleShuffleInts() {
	Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	ShuffleInts(ints)
	fmt.Println(ints)
	// Output: [74125 777 941 89416 8413 854 52 841657]
}

func ExampleFaker_ShuffleInts() {
	f := New(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	f.ShuffleInts(ints)
	fmt.Println(ints)
	// Output: [74125 777 941 89416 8413 854 52 841657]
}

func BenchmarkShuffleInts(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ShuffleInts([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})
}

func ExampleRandomInt() {
	Seed(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(RandomInt(ints))
	// Output: 52
}

func ExampleFaker_RandomInt() {
	f := New(11)

	ints := []int{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(f.RandomInt(ints))
	// Output: 52
}

func TestRandomInt(t *testing.T) {
	ints := []int{}
	RandomInt(ints)

	ints = []int{1}
	RandomInt(ints)
}

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInt([]int{52, 854, 941, 74125, 8413, 777, 89416, 841657})
	}
}

func ExampleRandomUint() {
	Seed(11)

	ints := []uint{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(RandomUint(ints))
	// Output: 52
}

func ExampleFaker_RandomUint() {
	f := New(11)

	ints := []uint{52, 854, 941, 74125, 8413, 777, 89416, 841657}
	fmt.Println(f.RandomUint(ints))
	// Output: 52
}

func TestRandomUint(t *testing.T) {
	ints := []uint{}
	RandomUint(ints)

	ints = []uint{1}
	RandomUint(ints)
}

func BenchmarkRandomUint(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RandomUint([]uint{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.RandomUint([]uint{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.RandomUint([]uint{52, 854, 941, 74125, 8413, 777, 89416, 841657})
		}
	})
}

func ExampleHexUint8() {
	Seed(11)
	fmt.Println(HexUint8())
	// Output: 0x87
}

func ExampleFaker_HexUint8() {
	f := New(11)
	fmt.Println(f.HexUint8())
	// Output: 0x87
}

func BenchmarkHexUint8(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint8()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint8()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint8()
		}
	})
}

func ExampleHexUint16() {
	Seed(11)
	fmt.Println(HexUint16())
	// Output: 0x8754
}

func ExampleFaker_HexUint16() {
	f := New(11)
	fmt.Println(f.HexUint16())
	// Output: 0x8754
}

func BenchmarkHexUint16(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint16()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint16()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint16()
		}
	})
}

func ExampleHexUint32() {
	Seed(11)
	fmt.Println(HexUint32())
	// Output: 0x87546957
}

func ExampleFaker_HexUint32() {
	f := New(11)
	fmt.Println(f.HexUint32())
	// Output: 0x87546957
}

func BenchmarkHexUint32(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint32()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint32()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint32()
		}
	})
}

func ExampleHexUint64() {
	Seed(11)
	fmt.Println(HexUint64())
	// Output: 0x875469578e51b5e5
}

func ExampleFaker_HexUint64() {
	f := New(11)
	fmt.Println(f.HexUint64())
	// Output: 0x875469578e51b5e5
}

func BenchmarkHexUint64(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint64()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint64()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint64()
		}
	})
}

func ExampleHexUint128() {
	Seed(11)
	fmt.Println(HexUint128())
	// Output: 0x875469578e51b5e56c95b64681d147a1
}

func ExampleFaker_HexUint128() {
	f := New(11)
	fmt.Println(f.HexUint128())
	// Output: 0x875469578e51b5e56c95b64681d147a1
}

func BenchmarkHexUint128(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint128()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint128()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint128()
		}
	})
}

func ExampleHexUint256() {
	Seed(11)
	fmt.Println(HexUint256())
	// Output: 0x875469578e51b5e56c95b64681d147a12cde48a4f417231b0c486abbc263e48d
}

func ExampleFaker_HexUint256() {
	f := New(11)
	fmt.Println(f.HexUint256())
	// Output: 0x875469578e51b5e56c95b64681d147a12cde48a4f417231b0c486abbc263e48d
}

func BenchmarkHexUint256(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexUint256()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexUint256()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexUint256()
		}
	})
}
