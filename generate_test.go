package gofakeit

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	output := ""
	var err error

	numTests := 1000
	for i := 0; i < numTests; i++ {
		output, err = Generate("{firstname} {lastname} {email} #?#?#?")
		if err != nil {
			t.Error(err)
		}
		if strings.Contains(output, "{") || strings.Contains(output, "}") {
			t.Error("Output should not contain either { or }. Output: ", output)
		}
	}
}

func TestGenerate_Sub(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		Seed(11)

		output, err := Generate("{randomstring:[{firstname},{lastname}]}")
		if err != nil {
			t.Error(err)
		}
		if output == "" {
			t.Error("Output was empty")
		}
	})

	t.Run("Complex", func(t *testing.T) {
		Seed(11)

		output, err := Generate("{randomstring:[{randomstring:[{firstname},{lastname}]},{randomstring:[{firstname},{lastname}]}]}")
		if err != nil {
			t.Error(err)
		}
		if output == "" {
			t.Error("Output was empty")
		}
	})
}

func ExampleGenerate() {
	Seed(11)

	genStr, _ := Generate("{firstname} {lastname} {email} #?#?#?")
	fmt.Println(genStr)
	genStr, _ = Generate("{sentence}")
	fmt.Println(genStr)
	genStr, _ = Generate("{shuffleints:[1,2,3]}")
	fmt.Println(genStr)
	genStr, _ = Generate("{randomint:[1,2,3,-4]}")
	fmt.Println(genStr)
	genStr, _ = Generate("{randomuint:[1,2,3,4]}")
	fmt.Println(genStr)
	genStr, _ = Generate("{number:1,50}")
	fmt.Println(genStr)
	genStr, _ = Generate("{shufflestrings:[key:value,int:string,1:2,a:b]}")
	fmt.Println(genStr)

	// Output: Mollie Legros vanceschroeder@turner.com 8K8b1M
	// Automate problem recovery victoriously.
	// [2 3 1]
	// 2
	// 4
	// 16
	// [int:string key:value a:b 1:2]
}

func ExampleFaker_Generate() {
	f := New(11)

	genStr, _ := f.Generate("{firstname} {lastname} {email} #?#?#?")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{sentence}")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{shuffleints:[1,2,3]}")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{randomint:[1,2,3,-4]}")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{randomuint:[1,2,3,4]}")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{number:1,50}")
	fmt.Println(genStr)
	genStr, _ = f.Generate("{shufflestrings:[key:value,int:string,1:2,a:b]}")
	fmt.Println(genStr)

	// Output: Mollie Legros vanceschroeder@turner.com 8K8b1M
	// Automate problem recovery victoriously.
	// [2 3 1]
	// 2
	// 4
	// 16
	// [int:string key:value a:b 1:2]
}

func BenchmarkGenerate(b *testing.B) {
	f := New(0)

	b.Run("simple_single_function", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{firstname}")
		}
	})

	b.Run("simple_multiple_functions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{firstname} {lastname} {email}")
		}
	})

	b.Run("function_with_params", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{number:1,100}")
		}
	})

	b.Run("multiple_functions_with_params", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{number:1,100} {sentence} {color}")
		}
	})

	b.Run("mixed_letters_numbers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "###-???-###")
		}
	})

	b.Run("complex_mixed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{firstname} {lastname} - {email} - ID: ###-???-### - {number:1,1000}")
		}
	})

	b.Run("no_replacements", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "just a plain string with no replacements")
		}
	})

	b.Run("many_functions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{firstname} {lastname} {email} {phone} {address} {city} {state} {zip} {country}")
		}
	})

	b.Run("nested_params", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{randomstring:[hello,world,foo,bar]}")
		}
	})

	b.Run("nested_complex", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "{randomstring:[{randomstring:[{firstname},{lastname}]},{randomstring:[{firstname},{lastname}]}]}")
		}
	})

	b.Run("bio_template", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "A {jobtitle} who {verb} {noun} and practices {hobby} in {city}.")
		}
	})

	b.Run("sentence_template", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			generate(f, "Choose {adjectivedescriptive} defaults.")
		}
	})

	b.Run("long_string_many_replacements", func(b *testing.B) {
		template := "{firstname} {lastname} lives in {city}, {state} {zip}. Email: {email}, Phone: {phone}. Works as {jobtitle} at {company}. Favorite color: {color}, Animal: {animal}."
		for i := 0; i < b.N; i++ {
			generate(f, template)
		}
	})
}

// Test generate function with various edge cases
func TestGenerateEdgeCases(t *testing.T) {
	testCases := []struct {
		template string
		validate func(string) bool
	}{
		{"{firstname}", func(s string) bool { return len(s) > 0 && !strings.Contains(s, "{") }},
		{"{firstname} {lastname}", func(s string) bool { return len(strings.Fields(s)) >= 2 && !strings.Contains(s, "{") }},
		{"{firstname} {lastname} {email}", func(s string) bool { return strings.Contains(s, "@") && !strings.Contains(s, "{") }},
		{"{number:1,100}", func(s string) bool { return len(s) > 0 && !strings.Contains(s, "{") }},
		{"###-???-###", func(s string) bool { return len(s) == 11 && s[3] == '-' && s[7] == '-' }},
		{"just a plain string", func(s string) bool { return s == "just a plain string" }},
		{"{randomstring:[hello,world]}", func(s string) bool { return s == "hello" || s == "world" }},
		{"A {jobtitle} in {city}.", func(s string) bool {
			return strings.HasPrefix(s, "A ") && strings.HasSuffix(s, ".") && !strings.Contains(s, "{")
		}},
		// Edge cases
		{"{unknown}", func(s string) bool { return s == "{unknown}" }},
		{"{firstname", func(s string) bool { return s == "{firstname" }},
		{"firstname}", func(s string) bool { return s == "firstname}" }},
		{"{}", func(s string) bool { return s == "{}" }},
	}

	for _, tc := range testCases {
		// Test multiple times with different seeds
		for seed := uint64(0); seed < 10; seed++ {
			f := New(seed)
			result, err := generate(f, tc.template)

			if err != nil {
				t.Errorf("Unexpected error for %q (seed %d): %v", tc.template, seed, err)
				continue
			}

			if !tc.validate(result) {
				t.Errorf("Validation failed for %q (seed %d): got %q", tc.template, seed, result)
			}
		}
	}
}

func ExampleFixedWidth() {
	Seed(11)

	value, err := FixedWidth(&FixedWidthOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "Name", Function: "{firstname} {lastname}"},
			{Name: "Email", Function: "email"},
			{Name: "Password", Function: "password", Params: MapParams{"special": {"false"}}},
			{Name: "Age", Function: "{number:1,100}"},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Name                 Email                          Password         Age
	// Sonny Stiedemann     mollielegros@wilderman.org     3T8l34B5F0eb     62
	// Cody Donnelly        stevebins@robel.io             3fnu2C12baqe     9
	// Julius Farrell       tomasaullrich@hane.name        e8W8SJGZi1Y4     77
}

func ExampleFixedWidth_default() {
	Seed(11)

	value, err := FixedWidth(nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Name                  Email                          Password
	// Russ Donnelly         missourigrimes@rohan.io        N3xh07881HKv
	// Cole Leannon          evanadams@marquardt.com        n2fja08VHMgw
	// Darien Rempel         thorajerde@cormier.com         ZhMPq5v8z6iP
	// Kaitlyn Wilderman     constancenienow@smitham.io     MkoTT5v5Mv6Q
	// Rahul Turner          irvingbins@lemke.io            8Wy0UV016u88
	// April Robel           beaulahmayert@kautzer.net      ykgR5bO83YSa
	// Tomasa Waters         lelandwaelchi@beahan.net       k68Ew58F9a0V
	// Sunny Hane            dawnbaumbach@mcdermott.biz     AbUjm6x5dwd8
	// Hilma Jast            lavernerobel@bailey.io         MyY49BLUWKCh
}

func ExampleFixedWidth_noHeader() {
	Seed(11)

	value, err := FixedWidth(&FixedWidthOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "", Function: "{firstname} {lastname}"},
			{Name: "", Function: "email"},
			{Name: "", Function: "password", Params: MapParams{"special": {"false"}}},
			{Name: "", Function: "{number:1,100}"},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)

	// Output: Sonny Stiedemann     mollielegros@wilderman.org     3T8l34B5F0eb     62
	// Cody Donnelly        stevebins@robel.io             3fnu2C12baqe     9
	// Julius Farrell       tomasaullrich@hane.name        e8W8SJGZi1Y4     77
}

func ExampleFaker_FixedWidth() {
	f := New(11)

	value, err := f.FixedWidth(&FixedWidthOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "Name", Function: "{firstname} {lastname}"},
			{Name: "Email", Function: "email"},
			{Name: "Password", Function: "password", Params: MapParams{"special": {"false"}}},
			{Name: "Age", Function: "{number:1,100}"},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Name                 Email                          Password         Age
	// Sonny Stiedemann     mollielegros@wilderman.org     3T8l34B5F0eb     62
	// Cody Donnelly        stevebins@robel.io             3fnu2C12baqe     9
	// Julius Farrell       tomasaullrich@hane.name        e8W8SJGZi1Y4     77
}

func TestFixedWidthLookup(t *testing.T) {
	faker := New(0)

	info := GetFuncLookup("fixed_width")

	m := MapParams{
		"rowcount": {"10"},
		"fields": {
			`{"name":"Name","function":"{firstname} {lastname}"}`,
			`{"name":"Email","function":"email"}`,
			`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
			`{"name":"Age","function":"{number:1,100}"}`,
		},
	}

	output, err := info.Generate(faker, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	value := string(output.(string))

	// Check the rumber of rows by counting the newlines, end of file has no newline
	if strings.Count(value, "\n") != 10 {
		t.Error("Expected 10+1(header row footer) rows, got", strings.Count(value, "\n"))
	}

}

func TestFixedWidthNoOptions(t *testing.T) {
	Seed(11)

	_, err := FixedWidth(nil)

	if err != nil {
		t.Fatal(err.Error())
	}

}

func BenchmarkFixedWidthLookup100(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("fixed_width")
		m := MapParams{
			"rowcount": {"100"},
			"fields": {
				// `{"name":"Name","function":"{firstname} {lastname}"}`,
				`{"name":"Email","function":"email"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Age","function":"{number:1,100}"}`,
			},
		}
		_, err := info.Generate(faker, &m, info)
		if err != nil {
			// Stack trace

			b.Fatal(err.Error())
		}
	}
}

func ExampleRegex() {
	Seed(11)

	fmt.Println(Regex("[abcdef]{5}"))
	fmt.Println(Regex("[[:upper:]]{5}"))
	fmt.Println(Regex("(hello|world|whats|up)"))
	fmt.Println(Regex(`^[a-z]{5,10}@[a-z]{5,10}\.(com|net|org)$`))

	// Output: ffbbd
	// GSNYV
	// up
	// tyyxi@kshtyia.net
}

func ExampleFaker_Regex() {
	f := New(11)

	fmt.Println(f.Regex("[abcdef]{5}"))
	fmt.Println(f.Regex("[[:upper:]]{5}"))
	fmt.Println(f.Regex("(hello|world|whats|up)"))
	fmt.Println(f.Regex(`^[a-z]{5,10}@[a-z]{5,10}\.(com|net|org)$`))

	// Output: ffbbd
	// GSNYV
	// up
	// tyyxi@kshtyia.net
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

	// Test Str1 matches regex
	if !regexp.MustCompile(`^\d+$`).MatchString(rg.Str1) {
		t.Errorf("Str1 should match regex: ^\\d+$ got: %s", rg.Str1)
	}
	if !regexp.MustCompile(`\D{3}`).MatchString(rg.Str2) {
		t.Errorf("Str2 should match regex: \\D{3} got: %s", rg.Str2)
	}
	if !regexp.MustCompile(`Z{2,5}`).MatchString(rg.Str3) {
		t.Errorf("Str3 should match regex: Z{2,5} got: %s", rg.Str3)
	}
	if !regexp.MustCompile(`[^1]{3,5}`).MatchString(rg.Str4) {
		t.Errorf("Str4 should match regex: [^1]{3,5} got: %s", rg.Str4)
	}
	if !regexp.MustCompile(`(ab|bc)def`).MatchString(rg.Str5) {
		t.Errorf("Str5 should match regex: (ab|bc)def got: %s", rg.Str5)
	}
	if !regexp.MustCompile(`((123)?){3}`).MatchString(rg.Str6) {
		t.Errorf("Str6 should match regex: ((123)?){3} got: %s", rg.Str6)
	}
	if !regexp.MustCompile(`[^abcdef]{5}`).MatchString(rg.Str7) {
		t.Errorf("Str7 should match regex: [^abcdef]{5} got: %s", rg.Str7)
	}
	if !regexp.MustCompile(`[a-zA-Z]{10}`).MatchString(rg.Str8) {
		t.Errorf("Str8 should match regex: [a-zA-Z]{10} got: %s", rg.Str8)
	}
	if !regexp.MustCompile(`[[:upper:]]{5}`).MatchString(rg.Str9) {
		t.Errorf("Str9 should match regex: [[:upper:]]{5} got: %s", rg.Str9)
	}
	if !regexp.MustCompile(`[^0-5a-z\s]{5}`).MatchString(rg.Str10) {
		t.Errorf("Str10 should match regex: [^0-5a-z\\s]{5} got: %s", rg.Str10)
	}
	if !regexp.MustCompile(`123[0-2]+.*\w{3}`).MatchString(rg.Str11) {
		t.Errorf("Str11 should match regex: 123[0-2]+.*\\w{3} got: %s", rg.Str11)
	}
	if !regexp.MustCompile(`(hello|world|whats|up)`).MatchString(rg.Str12) {
		t.Errorf("Str12 should match regex: (hello|world|whats|up) got: %s", rg.Str12)
	}
	if !regexp.MustCompile(`^\d{1,2}[/](1[0-2]|[1-9])[/]((19|20)\d{2})$`).MatchString(rg.Str13) {
		t.Errorf("Str13 should match regex: ^\\d{1,2}[/](1[0-2]|[1-9])[/]((19|20)\\d{2})$ got: %s", rg.Str13)
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Regex(`(hello|world|whats|up`)
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

	// Output: map[a:map[that:could] above:[sometimes vast whom us scarcely these huh] am:map[how:9310653] depend:map[hey:5402877] out:381794.78 since:916659.94 these:map[frequently:2314010]]
}

func ExampleFaker_Map() {
	f := New(11)
	fmt.Println(f.Map())

	// Output: map[a:map[that:could] above:[sometimes vast whom us scarcely these huh] am:map[how:9310653] depend:map[hey:5402877] out:381794.78 since:916659.94 these:map[frequently:2314010]]
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
