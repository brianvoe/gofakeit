package gofakeit

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	output := ""

	numTests := 1000
	for i := 0; i < numTests; i++ {
		output = Generate("{firstname} {lastname} {email} #?#?#?")
		if strings.Contains(output, "{") || strings.Contains(output, "}") {
			t.Error("Output should not contain either { or }. Output: ", output)
		}
	}
}

func TestGenerate_Sub(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		Seed(11)

		output := Generate("{randomstring:[{firstname},{lastname}]}")
		if output != "Moen" {
			t.Error("Did not generate what was expected. Got: ", output)
		}
	})

	t.Run("Complex", func(t *testing.T) {
		Seed(11)

		output := Generate("{randomstring:[{randomstring:[{firstname},{lastname}]},{randomstring:[{firstname},{lastname}]}]}")
		if output != "Kozey" {
			t.Error("Did not generate what was expected. Got: ", output)
		}
	})
}

func ExampleGenerate() {
	Seed(11)

	fmt.Println(Generate("{firstname} {lastname} ssn is {ssn} and lives at {street}"))
	fmt.Println(Generate("{sentence:3}"))
	fmt.Println(Generate("{shuffleints:[1,2,3]}"))
	fmt.Println(Generate("{number:1,50}"))
	fmt.Println(Generate("{shufflestrings:[key:value,int:string,1:2,a:b]}"))
	// Output: Markus Moen ssn is 952284213 and lives at 599 New Cliffs stad
	// Arrival tour security.
	// [1 3 2]
	// 34
	// [a:b key:value int:string 1:2]
}

func ExampleFaker_Generate() {
	f := New(11)

	fmt.Println(f.Generate("{firstname} {lastname} ssn is {ssn} and lives at {street}"))
	fmt.Println(f.Generate("{sentence:3}"))
	fmt.Println(f.Generate("{shuffleints:[1,2,3]}"))
	fmt.Println(f.Generate("{number:1,50}"))
	fmt.Println(f.Generate("{shufflestrings:[key:value,int:string,1:2,a:b]}"))
	// Output: Markus Moen ssn is 952284213 and lives at 599 New Cliffs stad
	// Arrival tour security.
	// [1 3 2]
	// 34
	// [a:b key:value int:string 1:2]
}

func BenchmarkGenerate(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Generate("{firstname} {lastname} {email} #?#?#?")
		}
	})

	b.Run("Complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Generate("{randomstring:[{randomstring:[{firstname},{lastname}]},{randomstring:[{firstname},{lastname}]}]}")
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Generate("{firstname} {lastname} {email} #?#?#?")
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Generate("{firstname} {lastname} {email} #?#?#?")
		}
	})
}

func ExampleRegex() {
	Seed(11)

	fmt.Println(Regex("[abcdef]{5}"))
	fmt.Println(Regex("[[:upper:]]{5}"))
	fmt.Println(Regex("(hello|world|whats|up)"))
	fmt.Println(Regex(`^[a-z]{5,10}@[a-z]{5,10}\.(com|net|org)$`))
	// Output: affec
	// RXHKI
	// up
	// ptapwy@dnsmkgtl.com
}

func ExampleFaker_Regex() {
	f := New(11)

	fmt.Println(f.Regex("[abcdef]{5}"))
	fmt.Println(f.Regex("[[:upper:]]{5}"))
	fmt.Println(f.Regex("(hello|world|whats|up)"))
	fmt.Println(f.Regex(`^[a-z]{5,10}@[a-z]{5,10}\.(com|net|org)$`))
	// Output: affec
	// RXHKI
	// up
	// ptapwy@dnsmkgtl.com
}

var regexes = []struct{ test string }{
	{`^\d+$`},
	{`\D{3}`},
	{`Z{2,5}`},
	{`[^1]{3,5}`},
	{`(ab|bc)def`},
	{`((123)?){3}`},
	{`[^abcdef]{5}`},
	{`[a-zA-Z]{100}`},
	{`[[:upper:]]{5}`},
	{`[^0-5a-z\s]{5}`},
	{`123[0-2]+.*\w{3}`},
	{`(hello|world|whats|up)`},
	{`^\d{1,2}[/](1[0-2]|[1-9])[/]((19|20)\d{2})$`},
	{`^((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])$`},
	{"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"},
}

func TestRegex(t *testing.T) {
	for _, r := range regexes {
		for i := 0; i < 100; i++ {
			// Try to compile regexTest
			regCompile, err := regexp.Compile(r.test)
			if err != nil {
				t.Fatal("Error compiling. regex: ", r.test, " failed to compile:", err)
			}

			// Generate string and test it matches the regex syntax
			reg := Regex(r.test)
			if !regCompile.MatchString(reg) {
				t.Error("Generated data does not match regex. Regex: ", r.test, " output: ", reg)
			}
		}
	}
}

func TestRegex_Struct(t *testing.T) {
	Seed(11)

	type Reggy struct {
		Str1  string `fake:"{regex:^\\d+$}"`
		Str2  string `fake:"{regex:\\D{3}}"`
		Str3  string `fake:"{regex:Z{2,5}}"`
		Str4  string `fake:"{regex:[^1]{3,5}}"`
		Str5  string `fake:"{regex:(ab|bc)def}"`
		Str6  string `fake:"{regex:((123)?){3}}"`
		Str7  string `fake:"{regex:[^abcdef]{5}}"`
		Str8  string `fake:"{regex:[a-zA-Z]{10}}"`
		Str9  string `fake:"{regex:[[:upper:]]{5}}"`
		Str10 string `fake:"{regex:[^0-5a-z\\s]{5}}"`
		Str11 string `fake:"{regex:123[0-2]+.*\\w{3}}"`
		Str12 string `fake:"{regex:(hello|world|whats|up)}"`
		Str13 string `fake:"{regex:^\\d{1,2}[/](1[0-2]|[1-9])[/]((19|20)\\d{2})$}"`
	}

	rg := Reggy{}
	Struct(&rg)

	if rg.Str1 != "165989" {
		t.Errorf("Str1 should be 165989 got: %s", rg.Str1)
	}
	if rg.Str2 != "X,P" {
		t.Errorf("Str2 should be X,P got: %s", rg.Str2)
	}
	if rg.Str3 != "ZZZZ" {
		t.Errorf("Str3 should be ZZZZ got: %s", rg.Str3)
	}
	if rg.Str4 != "+EW" {
		t.Errorf("Str4 should be +EW got: %s", rg.Str4)
	}
	if rg.Str5 != "abdef" {
		t.Errorf("Str5 should be abdef got: %s", rg.Str5)
	}
	if rg.Str6 != "123M123M123M" {
		t.Errorf("Str6 should be 123M123M123M got: %s", rg.Str6)
	}
	if rg.Str7 != "8|_Jh" {
		t.Errorf("Str7 should be 8|_Jh got: %s", rg.Str7)
	}
	if rg.Str8 != "wnqhqclaYk" {
		t.Errorf("Str8 should be wnqhqclaYk got: %s", rg.Str8)
	}
	if rg.Str9 != "WWFOR" {
		t.Errorf("Str9 should be WWFOR got: %s", rg.Str9)
	}
	if rg.Str10 != "G)#AA" {
		t.Errorf("Str10 should be G)#AA got: %s", rg.Str10)
	}
	if rg.Str11 != "12320kxu2d" {
		t.Errorf("Str11 should be 12320kxu2d got: %s", rg.Str11)
	}
	if rg.Str12 != "world" {
		t.Errorf("Str12 should be world got: %s", rg.Str12)
	}
	if rg.Str13 != "93/11/2029" {
		t.Errorf("Str13 should be 93/11/2029 got: %s", rg.Str13)
	}
}

func BenchmarkRegex(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Regex(`(hello|world|whats|up`)
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Regex(`(hello|world|whats|up`)
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Regex(`(hello|world|whats|up`)
		}
	})
}

func BenchmarkRegexEmail(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Regex("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Regex("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Regex("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		}
	})
}

func ExampleMap() {
	Seed(11)
	fmt.Println(Map())
	// Output: map[approach:map[mind:[arrival should resolve outcome hurt]] arrive:Coordinator consult:respond context:9285735 water:5081652]
}

func ExampleFaker_Map() {
	f := New(11)
	fmt.Println(f.Map())
	// Output: map[approach:map[mind:[arrival should resolve outcome hurt]] arrive:Coordinator consult:respond context:9285735 water:5081652]
}

func TestMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		Map()
	}
}

func BenchmarkMap(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Map()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Map()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Map()
		}
	})
}
