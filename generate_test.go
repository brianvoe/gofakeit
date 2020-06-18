package gofakeit

import (
	"fmt"
	"regexp"
	"testing"
)

func TestGenerate(t *testing.T) {
	numTests := 100
	for i := 0; i < numTests; i++ {
		Generate("{firstname} {lastname} {email} #?#?#?")
	}
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

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("{firstname} {lastname} {email} #?#?#?")
	}
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

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Regex(`(hello|world|whats|up)`)
	}
}

func BenchmarkRegexEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Regex("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	}
}

func ExampleMap() {
	Seed(11)

	fmt.Println(Map())
	// Output: map[approach:map[mind:[arrival should resolve outcome hurt]] arrive:Coordinator consult:respond context:9285735 water:5081652]
}

func TestMap(t *testing.T) {
	for i := 0; i < 100; i++ {
		Map()
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Map()
	}
}
