package gofakeit

import (
	"fmt"
	"strings"
	"testing"
	"text/template"
)

func ExampleTemplate() {
	Seed(11)

	template := `{{range IntRange 1 4}}{{FirstName}} {{LastName}}\n{{end}}`

	value, err := Template(template, &TemplateOptions{Data: 4})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Markus Moen
	// Alayna Wuckert
	// Lura Lockman
	// Sylvan Mraz

}

func ExampleFaker_Template() {
	f := New(11)

	template := `{{range IntRange 1 4}}{{FirstName}} {{LastName}}\n{{end}}`

	value, err := f.Template(template, &TemplateOptions{Data: 4})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Markus Moen
	// Alayna Wuckert
	// Lura Lockman
	// Sylvan Mraz
}

func TestPassedInFunctionMap(t *testing.T) {
	f := New(11)

	funcMap := template.FuncMap{
		"title": func(s string) string {
			words := strings.Fields(s)
			for i, word := range words {
				letters := strings.Split(word, "")
				letters[0] = strings.ToUpper(letters[0])
				words[i] = strings.Join(letters, "")
			}
			return strings.Join(words, " ")
		},
	}

	s := `{{printf "%q" (title .Data)}}`

	value, err := f.Template(s, &TemplateOptions{Funcs: funcMap, Data: "the go programming language"})
	if err != nil {
		fmt.Println(err)
	}

	if value != "\"The Go Programming Language\"" {
		t.Error("Expected \"The Go Programming Language\", got ", value)
	}
}

func TestTemplate_misc(t *testing.T) {
	f := New(11)

	// functions with parameters
	tests := map[string]string{
		"Weighted":            "{{Weighted (SliceAny `hello` 2 6.9) (SliceF32 1 2 3)}}",
		"Dice":                "{{Dice 3 (SliceUInt 1 5 3) }}",
		"IntRange":            "{{IntRange 1 5}}",
		"UintRange":           "{{UintRange 1 5}}",
		"RandomInt":           "{{RandomInt (SliceInt 1 5 3) }}",
		"RandomString":        "{{RandomString (SliceString `key1` `key2` `key3`)}}",
		"RandomUint":          "{{RandomUint (SliceUInt 2 6 9)}}",
		"Teams":               "{{Teams (SliceString `person_a` `person_b` `person_c`) (SliceString `team_a` `team_b` `team_c`)}}",
		"SliceAny":            "{{SliceAny `hello` 2 6.9}}",
		"SliceString":         "{{SliceString `testa` `testb` `testc`}}",
		"SliceF32":            "{{SliceF32 1.2 5.55 7.955}}",
		"SliceInt":            "{{SliceInt 1 5 9}}",
		"SliceUInt":           "{{SliceUInt 1 5 9}}",
		"Password":            "{{Password true false false false false 32}}",
		"ToUpper":             "{{ToUpper `this was lower`}}",
		"ToLower":             "{{ToLower `THIS WAS UPPER`}}",
		"Sentence":            "{{Sentence 3}}",
		"LoremIpsumSentence":  "{{LoremIpsumSentence 3}}",
		"Image":               "{{Image 1 1}}",
		"ImageJpeg":           "{{ImageJpeg 1 1}}",
		"ImagePng":            "{{ImagePng 1 1}}",
		"Numerify":            "{{Numerify `10` }}",
		"LatitudeInRange":     "{{LatitudeInRange 21 42 }}",
		"Number":              "{{Number 1 42 }}",
		"Price":               "{{Price 1 100 }}",
		"Regex":               "{{Regex `[abcdef]{5}` }}",
		"ImageURL":            "{{ImageURL 640 480 }}",
		"DigitN":              "{{DigitN 10 }}",
		"LoremIpsumParagraph": "{{LoremIpsumParagraph 3  5  12 `\n` }}",
		"Paragraph":           "{{Paragraph 3  5  12 `\n` }}",
		"Lexify":              "{{Lexify `hello`}}",
		"LogLevel":            "{{LogLevel `apache`}}",
		"CreditCardNumber":    "{{CreditCardNumber nil}}",
		"Markdown":            "{{Markdown nil}}",
		"EmailText":           "{{EmailText nil}}",
		"LetterN":             "{{LetterN 10}}",
		"FixedWidth":          "{{FixedWidth nil}}",
		"Float32Range":        "{{Float32Range 2.2 55.6}}",
		"Float64Range":        "{{Float64Range 2.2 55.6}}",
		"HipsterParagraph":    "{{HipsterParagraph 3  5  12 `\n`}}",
		"HipsterSentence":     "{{HipsterParagraph 3  5  12 `\n`}}",
		"LongitudeInRange":    "{{LongitudeInRange 21 42 }}",
		"Svg":                 "{{Svg  nil}}",
		"DateRange":           "{{DateRange (ToDate `2006-01-02`) (ToDate `2006-01-02`)}}",
	}

	for k, v := range tests {

		value, err := f.Template(v, nil)

		if err != nil {
			t.Error(k, err)
		}

		if string(value) == "" {
			t.Errorf("expected a value for %s, got nothing", k)
		}
	}
}

func TestTemplateLookup(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("template")

	m := MapParams{
		"template": {"{{range IntRange 1 (ToInt .Data)}}{{FirstName}} {{LastName}}\n{{end}}"},
		"data":     {"5"},
	}

	value, err := info.Generate(f.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if value != "Markus Moen\nAlayna Wuckert\nLura Lockman\nSylvan Mraz\nPaolo Rutherford\n" {
		t.Error("Expected `Markus Moen Pagac`, got ", value)
	}
}

func TestTemplateInternalFunctions(t *testing.T) {
	Seed(11)

	type test struct {
		template string
		expected string
	}

	tests := []test{
		// ToInt
		{
			template: `{{ ToInt 55 }}`,
			expected: "55",
		},

		// ToFloat
		{
			template: `{{ ToFloat 55.55 }}`,
			expected: "55.55",
		},

		// ToString
		{
			template: `{{ ToString 3.55 }}`,
			expected: "3.55",
		},
	}

	for k, v := range tests {
		value, err := Template(v.template, nil)
		if err != nil {
			t.Fatal(err)
		}
		if value != v.expected {
			t.Errorf("Test: %v expected `%s`, got `%s`", k, v.expected, value)
		}
	}
}

func TestTemplateNoOptions(t *testing.T) {
	Seed(11)

	tests := []string{
		"Markus Moen",
		"Alayna Wuckert",
		"Lura Lockman",
		"Sylvan Mraz",
		"Paolo Rutherford",
	}

	for k, v := range tests {
		// The above code is declaring a variable named "value" in the Go programming language.
		value, err := Template("{{FirstName}} {{LastName}}", nil)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(value, v) {
			t.Errorf("Test: %v could not find `%s`, in result %s", k, v, value)
		}
	}
}

func ExampleEmailText() {
	Seed(11)

	value, err := EmailText(&EmailOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// 	Subject: Greetings from Marcel!
	//
	// Dear Pagac,
	//
	// Hello there! Sending positive vibes your way.
	//
	// I hope you're doing great. May your week be filled with joy.
	//
	// This me far smile where was by army party riches. Theirs instead here mine whichever that those instance growth has. Ouch enough Swiss us since down he she aha us. You to upon how this this furniture way no play. Towel that us to accordingly theirs purse enough so though.
	//
	// Election often until eek weekly yet oops until conclude his. Stay elsewhere such that galaxy clean that last each stack. Reluctantly theirs wisp aid firstly highly butter accordingly should already. Calm shake according fade neither kuban upon this he fortnightly. Occasionally bunch on who elsewhere lastly hourly right there honesty.
	//
	// We is how result out Shakespearean have whom yearly another. Packet are behind late lot finally time themselves goodness quizzical. Our therefore could fact cackle yourselves zebra for whose enormously. All bowl out wandering secondly yellow another your hourly spit. Since tomorrow hers words little think will our by Polynesian.
	//
	// I'm curious to know what you think about it. If you have a moment, please feel free to check out the project on GitLab
	//
	// Your insights would be invaluable. Your thoughts matter to me.
	//
	// I appreciate your attention to this matter. Your feedback is greatly appreciated.
	//
	// Best wishes
	// Daryl Leannon
	// oceaneokuneva@roberts.org
	// 1-816-608-2233
}

func ExampleFaker_EmailText() {
	f := New(11)

	value, err := f.EmailText(&EmailOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// 	Subject: Greetings from Marcel!
	//
	// Dear Pagac,
	//
	// Hello there! Sending positive vibes your way.
	//
	// I hope you're doing great. May your week be filled with joy.
	//
	// This me far smile where was by army party riches. Theirs instead here mine whichever that those instance growth has. Ouch enough Swiss us since down he she aha us. You to upon how this this furniture way no play. Towel that us to accordingly theirs purse enough so though.
	//
	// Election often until eek weekly yet oops until conclude his. Stay elsewhere such that galaxy clean that last each stack. Reluctantly theirs wisp aid firstly highly butter accordingly should already. Calm shake according fade neither kuban upon this he fortnightly. Occasionally bunch on who elsewhere lastly hourly right there honesty.
	//
	// We is how result out Shakespearean have whom yearly another. Packet are behind late lot finally time themselves goodness quizzical. Our therefore could fact cackle yourselves zebra for whose enormously. All bowl out wandering secondly yellow another your hourly spit. Since tomorrow hers words little think will our by Polynesian.
	//
	// I'm curious to know what you think about it. If you have a moment, please feel free to check out the project on GitLab
	//
	// Your insights would be invaluable. Your thoughts matter to me.
	//
	// I appreciate your attention to this matter. Your feedback is greatly appreciated.
	//
	// Best wishes
	// Daryl Leannon
	// oceaneokuneva@roberts.org
	// 1-816-608-2233
}

// TemplateMarkdown examples and tests
func ExampleMarkdown() {
	Seed(11)

	value, err := Markdown(&MarkdownOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// 	# PurpleSheep5
	//
	// *Author: Amie Feil*
	//
	// Was by army party riches theirs instead. Here mine whichever that those instance growth. Has ouch enough Swiss us since down. He she aha us you to upon. How this this furniture way no play.
	//
	// Towel that us to accordingly theirs purse. Enough so though election often until eek. Weekly yet oops until conclude his stay. Elsewhere such that galaxy clean that last. Each stack reluctantly theirs wisp aid firstly.
	//
	// ## Table of Contents
	// - [Installation](#installation)
	// - [Usage](#usage)
	// - [License](#license)
	//
	// ## Installation
	// '''bash
	// pip install PurpleSheep5
	// '''
	//
	// ## Usage
	// '''python
	// result = purplesheep5.process("funny request")
	// print("purplesheep5 result:", "in progress")
	// '''
	//
	// ## License
	// MIT
}

func ExampleFaker_Markdown() {
	f := New(11)

	value, err := f.Markdown(&MarkdownOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// 	# PurpleSheep5
	//
	// *Author: Amie Feil*
	//
	// Was by army party riches theirs instead. Here mine whichever that those instance growth. Has ouch enough Swiss us since down. He she aha us you to upon. How this this furniture way no play.
	//
	// Towel that us to accordingly theirs purse. Enough so though election often until eek. Weekly yet oops until conclude his stay. Elsewhere such that galaxy clean that last. Each stack reluctantly theirs wisp aid firstly.
	//
	// ## Table of Contents
	// - [Installation](#installation)
	// - [Usage](#usage)
	// - [License](#license)
	//
	// ## Installation
	// '''bash
	// pip install PurpleSheep5
	// '''
	//
	// ## Usage
	// '''python
	// result = purplesheep5.process("funny request")
	// print("purplesheep5 result:", "in progress")
	// '''
	//
	// ## License
	// MIT
}

func BenchmarkTemplate100(b *testing.B) {
	f := New(11)

	for i := 0; i < 100; i++ {
		_, err := f.Template("{{range IntRange 1 (ToInt .Data)}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)

	for i := 0; i < 1000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 (ToInt .Data)}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}
