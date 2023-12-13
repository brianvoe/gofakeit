package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleColor() {
	Seed(11)
	fmt.Println(Color())

	// Output: MediumOrchid
}

func ExampleFaker_Color() {
	f := New(11)
	fmt.Println(f.Color())

	// Output: MediumOrchid
}

func BenchmarkColor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Color()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Color()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Color()
		}
	})
}

func ExampleNiceColors() {
	Seed(11)
	fmt.Println(NiceColors())

	// Output: [#f6f6f6 #e8e8e8 #333333 #990100 #b90504]
}

func ExampleFaker_NiceColors() {
	f := New(11)
	fmt.Println(f.NiceColors())

	// Output: [#f6f6f6 #e8e8e8 #333333 #990100 #b90504]
}

func BenchmarkNiceColors(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NiceColors()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.NiceColors()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.NiceColors()
		}
	})
}

func ExampleSafeColor() {
	Seed(11)
	fmt.Println(SafeColor())

	// Output: black
}

func ExampleFaker_SafeColor() {
	f := New(11)
	fmt.Println(f.SafeColor())

	// Output: black
}

func BenchmarkSafeColor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SafeColor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.SafeColor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.SafeColor()
		}
	})
}

func ExampleHexColor() {
	Seed(11)
	fmt.Println(HexColor())

	// Output: #a99fb4
}

func ExampleFaker_HexColor() {
	f := New(11)
	fmt.Println(f.HexColor())

	// Output: #a99fb4
}

func BenchmarkHexColor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HexColor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HexColor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HexColor()
		}
	})
}

func ExampleRGBColor() {
	Seed(11)
	fmt.Println(RGBColor())

	// Output: [89 176 195]
}

func ExampleFaker_RGBColor() {
	f := New(11)
	fmt.Println(f.RGBColor())

	// Output: [89 176 195]
}

func BenchmarkRGBColor(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RGBColor()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.RGBColor()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.RGBColor()
		}
	})
}
