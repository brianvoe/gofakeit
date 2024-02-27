package gofakeit

import (
	"fmt"
	"regexp"
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

	// Output: Sonny Stiedemann
	// Cody Donnelly
	// Julius Farrell
	// Mollie Legros

}

func ExampleFaker_Template() {
	f := New(11)

	template := `{{range IntRange 1 4}}{{FirstName}} {{LastName}}\n{{end}}`

	value, err := f.Template(template, &TemplateOptions{Data: 4})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Sonny Stiedemann
	// Cody Donnelly
	// Julius Farrell
	// Mollie Legros
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

	value, err := info.Generate(f, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Check if value is in the expected format. name[space]name\n
	reg := `^(\w+\s\w+\n){5}$`
	matched, err := regexp.MatchString(reg, value.(string))
	if err != nil {
		t.Fatal(err.Error())
	}

	if !matched {
		t.Errorf("expected %s to match %s", value, reg)
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

	// The above code is declaring a variable named "value" in the Go programming language.
	value, err := Template("{{FirstName}} {{LastName}}", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Make sure value has a space in it
	if !strings.Contains(value, " ") {
		t.Errorf("expected a space in %s", value)
	}

}

func ExampleEmailText() {
	Seed(11)

	value, err := EmailText(&EmailOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Subject: Hi from Russ!
	//
	// Dear Donnelly,
	//
	// Greetings! I hope your day is going well.
	//
	// I trust this email finds you well. Sending good vibes your way.
	//
	// Galaxy they how shall cut lastly that someone additionally ambulance. Could a tennis occur why depend heels out can fire. Anyone sometimes that leap whom troop now scarcely since dance. Huh somebody constantly Italian outfit certain weekly murder who cackle. Because her break write nobody could according whoa bevy we.
	//
	// Them summation to oxygen Greek how previously a what answer. Handle daily ouch upon encourage that after his every she. Wearily still he herself huge failure for yay sparse thoughtfully. Box anyway Jungian regularly world great daily recently whose Elizabethan. Outside to theirs block while French closely sit you to.
	//
	// Out fight bored yours those bale Parisian host which so. Today Ecuadorian hourly decidedly his everybody this play this little. Them on recently eager usually bunch daily whose somebody them. Few tonight dishonesty spell battery less patience whose inside hair. Your answer hand tonight these she team which today Einsteinian.
	//
	// I'm eager to hear your feedback on it. If you have a moment, please feel free to check out the project on GitLab
	//
	// Your insights would be invaluable. Your thoughts matter to me.
	//
	// I appreciate your attention to this matter. Your feedback is greatly appreciated.
	//
	// Sincerely
	// Hyman Rogahn
	// mervingottlieb@goldner.com
	// 1-275-132-4165
}

func ExampleFaker_EmailText() {
	f := New(11)

	value, err := f.EmailText(&EmailOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: Subject: Hi from Russ!
	//
	// Dear Donnelly,
	//
	// Greetings! I hope your day is going well.
	//
	// I trust this email finds you well. Sending good vibes your way.
	//
	// Galaxy they how shall cut lastly that someone additionally ambulance. Could a tennis occur why depend heels out can fire. Anyone sometimes that leap whom troop now scarcely since dance. Huh somebody constantly Italian outfit certain weekly murder who cackle. Because her break write nobody could according whoa bevy we.
	//
	// Them summation to oxygen Greek how previously a what answer. Handle daily ouch upon encourage that after his every she. Wearily still he herself huge failure for yay sparse thoughtfully. Box anyway Jungian regularly world great daily recently whose Elizabethan. Outside to theirs block while French closely sit you to.
	//
	// Out fight bored yours those bale Parisian host which so. Today Ecuadorian hourly decidedly his everybody this play this little. Them on recently eager usually bunch daily whose somebody them. Few tonight dishonesty spell battery less patience whose inside hair. Your answer hand tonight these she team which today Einsteinian.
	//
	// I'm eager to hear your feedback on it. If you have a moment, please feel free to check out the project on GitLab
	//
	// Your insights would be invaluable. Your thoughts matter to me.
	//
	// I appreciate your attention to this matter. Your feedback is greatly appreciated.
	//
	// Sincerely
	// Hyman Rogahn
	// mervingottlieb@goldner.com
	// 1-275-132-4165
}

// TemplateMarkdown examples and tests
func ExampleMarkdown() {
	Seed(11)

	value, err := Markdown(&MarkdownOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: # TurkeyThinker
	//
	// *Author: Tomasa Waters*
	//
	// Cut lastly that someone additionally ambulance could. A tennis occur why depend heels out. Can fire anyone sometimes that leap whom. Troop now scarcely since dance huh somebody. Constantly Italian outfit certain weekly murder who.
	//
	// Cackle because her break write nobody could. According whoa bevy we them summation to. Oxygen Greek how previously a what answer. Handle daily ouch upon encourage that after. His every she wearily still he herself.
	//
	// ## Table of Contents
	// - [Installation](#installation)
	// - [Usage](#usage)
	// - [License](#license)
	//
	// ## Installation
	// '''bash
	// pip install TurkeyThinker
	// '''
	//
	// ## Usage
	// '''python
	// result = turkeythinker.handle("whimsical story")
	// print("turkeythinker result:", "terminated")
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

	// Output: # TurkeyThinker
	//
	// *Author: Tomasa Waters*
	//
	// Cut lastly that someone additionally ambulance could. A tennis occur why depend heels out. Can fire anyone sometimes that leap whom. Troop now scarcely since dance huh somebody. Constantly Italian outfit certain weekly murder who.
	//
	// Cackle because her break write nobody could. According whoa bevy we them summation to. Oxygen Greek how previously a what answer. Handle daily ouch upon encourage that after. His every she wearily still he herself.
	//
	// ## Table of Contents
	// - [Installation](#installation)
	// - [Usage](#usage)
	// - [License](#license)
	//
	// ## Installation
	// '''bash
	// pip install TurkeyThinker
	// '''
	//
	// ## Usage
	// '''python
	// result = turkeythinker.handle("whimsical story")
	// print("turkeythinker result:", "terminated")
	// '''
	//
	// ## License
	// MIT
}

func BenchmarkTemplate(b *testing.B) {
	f := New(11)

	for i := 0; i < b.N; i++ {
		_, err := f.Template("{{range IntRange 1 (ToInt .Data)}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
