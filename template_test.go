package gofakeit

import (
	"fmt"
	"strings"
	"testing"
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
	// Subject: Greetings from Marcel!
	//
	// Dear Pagac,
	//
	// Hello there! Sending positive vibes your way.
	//
	// I hope you're doing great. May your week be filled with joy.
	//
	// Virtually woman where team late quarterly without week it hungry. Thing someone him regularly today whomever this revolt hence from. His timing as quantity us these yours live these frantic. Not may another how this ours his them those whose. Them batch its Iraqi most that few abroad cheese this.
	//
	// Whereas next how there gorgeous genetics time choir fiction therefore. Yourselves am those infrequently heap software quarterly rather punctuation yellow. Where several his orchard to frequently hence victorious boxers each. Does auspicious yourselves first soup tomorrow this that must conclude. Anyway some yearly who cough laugh himself both yet rarely.
	//
	// Me dolphin intensely block would leap plane us first then. Down them eager would hundred super throughout animal yet themselves. Been group flock shake part purchase up usually it her. None it hers boat what their there Turkmen moreover one. Lebanese to brace these shower in it everybody should whatever.
	//
	// I'm curious to know what you think about it. If you have a moment, please feel free to check out the project on Bitbucket
	//
	// I'm eager to hear what you think. Looking forward to your feedback!
	//
	// Thank you for your consideration! Thanks in advance for your time.
	//
	// Kind regards
	// Milford Johnston
	// jamelhaag@king.org
	// (507)096-3058
}

func ExampleFaker_EmailText() {
	f := New(11)

	value, err := f.EmailText(&EmailOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Subject: Greetings from Marcel!
	//
	// Dear Pagac,
	//
	// Hello there! Sending positive vibes your way.
	//
	// I hope you're doing great. May your week be filled with joy.
	//
	// Virtually woman where team late quarterly without week it hungry. Thing someone him regularly today whomever this revolt hence from. His timing as quantity us these yours live these frantic. Not may another how this ours his them those whose. Them batch its Iraqi most that few abroad cheese this.
	//
	// Whereas next how there gorgeous genetics time choir fiction therefore. Yourselves am those infrequently heap software quarterly rather punctuation yellow. Where several his orchard to frequently hence victorious boxers each. Does auspicious yourselves first soup tomorrow this that must conclude. Anyway some yearly who cough laugh himself both yet rarely.
	//
	// Me dolphin intensely block would leap plane us first then. Down them eager would hundred super throughout animal yet themselves. Been group flock shake part purchase up usually it her. None it hers boat what their there Turkmen moreover one. Lebanese to brace these shower in it everybody should whatever.
	//
	// I'm curious to know what you think about it. If you have a moment, please feel free to check out the project on Bitbucket
	//
	// I'm eager to hear what you think. Looking forward to your feedback!
	//
	// Thank you for your consideration! Thanks in advance for your time.
	//
	// Kind regards
	// Milford Johnston
	// jamelhaag@king.org
	// (507)096-3058
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
	// # PurpleSheep5
	//
	// *Author: Amie Feil*
	//
	// Quarterly without week it hungry thing someone. Him regularly today whomever this revolt hence. From his timing as quantity us these. Yours live these frantic not may another. How this ours his them those whose.
	//
	// Them batch its Iraqi most that few. Abroad cheese this whereas next how there. Gorgeous genetics time choir fiction therefore yourselves. Am those infrequently heap software quarterly rather. Punctuation yellow where several his orchard to.
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
	// # PurpleSheep5
	//
	// *Author: Amie Feil*
	//
	// Quarterly without week it hungry thing someone. Him regularly today whomever this revolt hence. From his timing as quantity us these. Yours live these frantic not may another. How this ours his them those whose.
	//
	// Them batch its Iraqi most that few. Abroad cheese this whereas next how there. Gorgeous genetics time choir fiction therefore yourselves. Am those infrequently heap software quarterly rather. Punctuation yellow where several his orchard to.
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
