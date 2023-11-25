package gofakeit

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

// build a list of all the template functions
func templateFunctionTestList(t *testing.T, r *rand.Rand) map[string]string {

	// functions with parameters
	template_function_tests := map[string]string{
		"Weighted":            "{{Weighted (ListI `hello` 2 6.9) (ListF32 1 2 3)}}",
		"Dice":                "{{Dice 3 (ListUInt 1 5 3) }}",
		"RandomInt":           "{{RandomInt (ListInt 1 5 3) }}",
		"RandomString":        "{{RandomString (ListS `key1` `key2` `key3`)}}",
		"RandomUint":          "{{RandomUint (ListUInt 2 6 9)}}",
		"Teams":               "{{Teams (ListS `person_a` `person_b` `person_c`) (ListS `team_a` `team_b` `team_c`)}}",
		"ListS":               "{{ListS `testa` `testb` `testc`}}",
		"UintRange":           "{{UintRange 1 10}}",
		"IntRange":            "{{IntRange 1 10}}",
		"ListF32":             "{{ListF32 1.2 5.55 7.955}}",
		"ListInt":             "{{ListInt 1 5 9}}",
		"ListUInt":            "{{ListUInt 1 5 9}}",
		"Int":                 "{{Int `5`}}",
		"Password":            "{{Password true false false false false 32}}",
		"Upper":               "{{Upper `this was lower`}}",
		"LCase":               "{{LCase `THIS WAS UPPER`}}",
		"Sentence":            "{{Sentence 3}}",
		"LoremIpsumSentence":  "{{LoremIpsumSentence 3}}",
		"ImagePng":            "{{ImagePng 1 1}}",
		"Image":               "{{Image 1 1}}",
		"ImageJpeg":           "{{ImageJpeg 1 1}}",
		"Numerify":            "{{Numerify `10` }}",
		"CreateListResult":    "{{CreateListResult 1 10}}",
		"ListResult":          "{{ListResult (CreateListResult 1 10) 1 10 true}}",
		"LatitudeInRange":     "{{LatitudeInRange 21 42 }}",
		"Number":              "{{Number 1 42 }}",
		"Price":               "{{Price 1 100 }}",
		"Regex":               "{{Regex `[abcdef]{5}` }}",
		"ImageURL":            "{{ImageURL 640 480 }}",
		"DigitN":              "{{DigitN 10 }}",
		"LoremIpsumParagraph": "{{LoremIpsumParagraph 3  5  12 `\n` }}",
		"ListI":               "{{ListI `hello` 2 6.9}}",
		"Paragraph":           "{{Paragraph 3  5  12 `\n` }}",
		"Lexify":              "{{Lexify `hello`}}",
		"LogLevel":            "{{LogLevel `apache`}}",
		"Replace":             "{{Replace `hello` `l` `z`}}",
		"CreditCardNumber":    "{{CreditCardNumber nil}}",
		"Concat":              "{{Concat `-` `a` `b`}}",
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
		"XML":                 "{{String (XML  nil)}}",
		"JSON":                "{{String (JSON  nil)}}",
		"CSV":                 "{{String (CSV  nil)}}",
		"DateRange":           "{{DateRange (DateS `2006-01-02`) (DateS `2006-01-02`)}}",
	}

	f := &Faker{Rand: r}
	v := reflect.ValueOf(f)

	// A list of all the excluded methods
	templateExclusion := []string{
		"RandomMapKey",
		"Slice",
		"Struct",
		"SQL",
		"Generate",
		"Template",
	}

	// loop through the methods
	for i := 0; i < v.NumMethod(); i++ {

		// check if the method is in the exclusion list
		if stringInSlice(v.Type().Method(i).Name, templateExclusion) {
			continue
		}

		// Check if method has return values
		// If not don't add to function map
		if v.Type().Method(i).Type.NumOut() == 0 {
			continue
		}

		// Check if method has parameters
		if v.Type().Method(i).Type.NumIn() > 1 {
			//check if the method is in the template_function_tests if not error
			if _, ok := template_function_tests[v.Type().Method(i).Name]; !ok {
				t.Errorf("method %s not in template_function_tests", v.Type().Method(i).Name)
			}
			continue
		}

		// add the method to the template_function_tests
		template_function_tests[v.Type().Method(i).Name] = fmt.Sprintf("{{%s}}", v.Type().Method(i).Name)
	}

	return template_function_tests
}

// Test all the function work with he template engine
func TestAllTemplateFunctions(t *testing.T) {
	f := New(11)

	test := templateFunctionTestList(t, f.Rand)

	for k, v := range test {

		value, err := f.Template(v, nil)

		if err != nil {
			t.Error(k, err)
		}

		if string(value) == "" {
			t.Errorf("expected a value for %s, got nothing", k)
		}
	}

}

func ExampleTemplate() {
	Seed(11)

	template := `{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}`

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

	template := `{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}`

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

func TestTemplateLookup(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("template")

	m := MapParams{
		"template": {"{{range $y := IntRange 1 (Int (.Data))}}{{FirstName}} {{LastName}}\n{{end}}"},
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

func TestFakertTemplateNoOption(t *testing.T) {
	f := New(11)

	tests := []string{
		"Markus Moen",
		"Alayna Wuckert",
		"Lura Lockman",
		"Sylvan Mraz",
		"Paolo Rutherford",
	}
	//println(value.(string))
	for i, v := range tests {
		// The above code is declaring a variable named "value" in the Go programming language.
		value, err := f.Template("{{FirstName}} {{LastName}}", nil)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(value, v) {
			t.Errorf("Test: %v could not find `%s`, in result %s", i, v, value)
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
	//println(value.(string))
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

func TestTemplateLookupNoOptions(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("template")

	// The above code is declaring a variable named "value" in the Go programming language.
	_, err := info.Generate(f.Rand, nil, info)
	if err == nil {
		t.Error("Was expecting an error ", err)
	}

}
func TestTemplateNoTemplateParam(t *testing.T) {
	Seed(11)

	value, err := Template("", nil)

	if err == nil {
		t.Error("Was expecting an error ", err)
	}
	if string(value) != "" {
		t.Errorf("expected a nothing, got %s", string(value))
	}
}
func TestFakerTemplateNoTemplateParam(t *testing.T) {
	f := New(11)

	value, err := f.Template("", nil)

	if err == nil {
		t.Error("Was expecting an error ", err)
	}
	if string(value) != "" {
		t.Errorf("expected a nothing, got %s", string(value))
	}
}

func ExampleEmailText() {
	Seed(11)

	value, err := EmailText(&EmailOptions{SectionsCount: 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Hi, Marcel
	//
	// you have look now that regularly Happens tonight
	// Gentrify pitchfork stumptown mlkshk umami chambray blue bottle 3 wolf moon swag &#43;1 biodiesel knausgaard semiotics taxidermy meh artisan. Hoodie &#43;1 blue bottle fashion axe forage mixtape Thundercats pork belly whatever 90&#39;s beard selfies chambray cred mlkshk shabby chic. Typewriter VHS readymade lo-fi bitters PBR&amp;B gentrify lomo raw denim freegan put a bird on it raw denim cliche dreamcatcher pug fixie.
	//
	// Park trust fund migas fingerstache sriracha &#43;1 mustache tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid. Stumptown organic schlitz flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy chia. Mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade deep v green juice keffiyeh carry poutine.
	//
	// Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat cornhole everyday taxidermy humblebrag bespoke kitsch. Hella venmo umami retro five dollar toast twee bitters microdosing cold-pressed street chambray distillery kale chips disrupt actually typewriter. Lomo kinfolk tilde heirloom normcore farm-to-table master readymade schlitz shoreditch bitters kitsch marfa beard roof butcher.
	// Pinterest semiotics photo booth taxidermy offal tofu whatever schlitz meh?
	//
	// I appreciate your feedback
	// Pinkie
	//
	// Company: Equal Pay for Women
	// Job Role: Director
	//
	// Address: 36132 Lake Portsview, Chula Vista, Ohio 99001
	// City: Chula Vista
	// State: Ohio
	// Zip: 99001Phone: 7871591604
	// Email: winstonhartmann@sauer.biz

}

func ExampleFaker_EmailText() {
	f := New(11)

	value, err := f.EmailText(&EmailOptions{SectionsCount: 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Hi, Marcel
	//
	// you have look now that regularly Happens tonight
	// Gentrify pitchfork stumptown mlkshk umami chambray blue bottle 3 wolf moon swag &#43;1 biodiesel knausgaard semiotics taxidermy meh artisan. Hoodie &#43;1 blue bottle fashion axe forage mixtape Thundercats pork belly whatever 90&#39;s beard selfies chambray cred mlkshk shabby chic. Typewriter VHS readymade lo-fi bitters PBR&amp;B gentrify lomo raw denim freegan put a bird on it raw denim cliche dreamcatcher pug fixie.
	//
	// Park trust fund migas fingerstache sriracha &#43;1 mustache tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid. Stumptown organic schlitz flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy chia. Mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade deep v green juice keffiyeh carry poutine.
	//
	// Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat cornhole everyday taxidermy humblebrag bespoke kitsch. Hella venmo umami retro five dollar toast twee bitters microdosing cold-pressed street chambray distillery kale chips disrupt actually typewriter. Lomo kinfolk tilde heirloom normcore farm-to-table master readymade schlitz shoreditch bitters kitsch marfa beard roof butcher.
	// Pinterest semiotics photo booth taxidermy offal tofu whatever schlitz meh?
	//
	// I appreciate your feedback
	// Pinkie
	//
	// Company: Equal Pay for Women
	// Job Role: Director
	//
	// Address: 36132 Lake Portsview, Chula Vista, Ohio 99001
	// City: Chula Vista
	// State: Ohio
	// Zip: 99001Phone: 7871591604
	// Email: winstonhartmann@sauer.biz

}

func TestEmailTextLookup(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("email_text")

	m := MapParams{
		"sections_count": {"3"},
	}

	value, err := info.Generate(f.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}
	tests := map[string]string{
		"Test 1": "you have look now that regularly Happens tonight",
		"Test 2": "mustache tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid",
		"Test 3": "cold-pressed tilde plaid keytar tilde bespoke waistcoat",
		"Test 4": "I appreciate your feedback",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value.(string), v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value.(string))
		}
	}

}

func TestEmailText(t *testing.T) {
	f := New(11)

	// The above code is declaring a variable named "value" in the Go programming language.
	value, err := f.EmailText(&EmailOptions{SectionsCount: 6})
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat",
		"Test 2": "oxygen scold what silence go time behind example me soon. Grade purely that",
		"Test 3": "however whoever itself over this pair smoke yourself so circumstances despite",
		"Test 4": "Thank you",
		"Test 5": "Address: 81575 North Daleton, Chicago, New Jersey 44967",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestFakertEmailNoOption(t *testing.T) {
	f := New(11)

	// The above code is declaring a variable named "value" in the Go programming language.
	value, err := f.EmailText(nil)
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat",
		"Test 2": "oxygen scold what silence go time behind example me soon. Grade purely that",
		"Test 3": "however whoever itself over this pair smoke yourself so circumstances despite",
		"Test 4": "Thank you",
		"Test 5": "Address: 81575 North Daleton, Chicago, New Jersey 44967",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestEmailNoOptions(t *testing.T) {
	Seed(11)

	value, err := EmailText(nil)
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat",
		"Test 2": "oxygen scold what silence go time behind example me soon. Grade purely that",
		"Test 3": "however whoever itself over this pair smoke yourself so circumstances despite",
		"Test 4": "Thank you",
		"Test 5": "Address: 81575 North Daleton, Chicago, New Jersey 44967",
	}
	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestEmailLookupNoOptions(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("email_text")

	value, err := info.Generate(f.Rand, nil, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	tests := map[string]string{
		"Test 1": "Email: rosannachristiansen@klocko.io",
		"Test 2": "Phone: 4185959586",
		"Test 3": "Freida",
		"Test 4": "I appreciate your input",
		"Test 5": "you have look now that regularly Happens tonight",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value.(string), v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value.(string))
		}
	}

}

// TemplateMarkdown examples and tests

func ExampleMarkdown() {
	Seed(11)

	value, err := Markdown(&MarkdownOptions{SectionsCount: 3})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	//
	// ## Url
	//
	// Golang you will need to install
	//
	// [http://www.humannext-generation.com/systems/distributed/seize/users](http://www.humannext-generation.com/systems/distributed/seize/users)
	// [http://www.dynamicincentivize.net/holistic/syndicate](http://www.dynamicincentivize.net/holistic/syndicate)
	// [http://www.futuredistributed.io/morph/sticky](http://www.futuredistributed.io/morph/sticky)
	// [http://www.internalarchitectures.com/transition/cross-platform/virtual/robust](http://www.internalarchitectures.com/transition/cross-platform/virtual/robust)
	// [https://www.productwebservices.biz/e-enable/sexy](https://www.productwebservices.biz/e-enable/sexy)
	// [http://www.legacyinfrastructures.io/plug-and-play](http://www.legacyinfrastructures.io/plug-and-play)
	//
	// ---
	//
	// ## Details
	//
	// You. Some. Delay. Sudanese.
	//
	// Of. Table. Class. May.
	//
	// Another. How. This. Ours.
	//
	// His. Them. Those. Whose.
	//
	// Them. Batch. Its. Iraqi.
	//
	// <details>
	// <summary>A calm street climb a ship. </summary>
	//
	// How there gorgeous genetics time choir fiction. Therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his. Orchard to frequently hence victorious boxers each. Does auspicious yourselves first soup tomorrow this.
	//
	// That must conclude anyway some yearly who. Cough laugh himself both yet rarely me. Dolphin intensely block would leap plane us. First then down them eager would hundred. Super throughout animal yet themselves been group.
	//
	// Flock shake part purchase up usually it. Her none it hers boat what their. There Turkmen moreover one Lebanese to brace. These shower in it everybody should whatever. Those uninterested several nobody yours thankful innocent.
	//
	// Power to any from its few luxury. None boy religion themselves it there might. Must near theirs mine thing tonight outside. Rapidly spotted solemnly finally been did confusion. These were nobody early silently company quarterly.
	//
	// American there time many French a each. Anybody rather paint ours did tonight remove. First including chair your at in by. Cackle whose they yearly wisdom nightly nightly. When finally without oxygen scold what silence..
	//
	// </details>
	//
	// ---
	//
	// ## LISTS
	//
	// Example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually ours problem slavery. Someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary from belief this. Upon at most homeless elsewhere has yearly under since where Californian all in today generally myself after stupid highly heavy. Here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me tonight place few.
	//
	// Why from somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple consequently in. Those daily had anywhere anything what in bouquet which annually as Cypriot this our sand tightly we first their staff. Invention however whoever itself over this pair smoke yourself so circumstances despite could project did embarrassed my philosophy imagination did. Stemmed mob furthermore read myself above had for successful for move here point less to myself sparse ours yours because..
	//
	// 1. ingeniously dive
	// 1. carelessly cry
	// 1. mysteriously shake a bored envy blindly straight
	// 1. dance courageously
	// 1. elegantly give
	// 1. rudely clap
	// 1. think a body well patiently weakly
	// 1. turn a photographer
	//
	// ---
	//
}

func ExampleFaker_Markdown() {
	f := New(11)

	value, err := f.Markdown(&MarkdownOptions{SectionsCount: 2})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	//
	// ## Url
	//
	// Golang you will need to install
	//
	// [http://www.humannext-generation.com/systems/distributed/seize/users](http://www.humannext-generation.com/systems/distributed/seize/users)
	// [http://www.dynamicincentivize.net/holistic/syndicate](http://www.dynamicincentivize.net/holistic/syndicate)
	// [http://www.futuredistributed.io/morph/sticky](http://www.futuredistributed.io/morph/sticky)
	// [http://www.internalarchitectures.com/transition/cross-platform/virtual/robust](http://www.internalarchitectures.com/transition/cross-platform/virtual/robust)
	// [https://www.productwebservices.biz/e-enable/sexy](https://www.productwebservices.biz/e-enable/sexy)
	// [http://www.legacyinfrastructures.io/plug-and-play](http://www.legacyinfrastructures.io/plug-and-play)
	//
	// ---
	//
	// ## Details
	//
	// You. Some. Delay. Sudanese.
	//
	// Of. Table. Class. May.
	//
	// Another. How. This. Ours.
	//
	// His. Them. Those. Whose.
	//
	// Them. Batch. Its. Iraqi.
	//
	// <details>
	// <summary>A calm street climb a ship. </summary>
	//
	// How there gorgeous genetics time choir fiction. Therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his. Orchard to frequently hence victorious boxers each. Does auspicious yourselves first soup tomorrow this.
	//
	// That must conclude anyway some yearly who. Cough laugh himself both yet rarely me. Dolphin intensely block would leap plane us. First then down them eager would hundred. Super throughout animal yet themselves been group.
	//
	// Flock shake part purchase up usually it. Her none it hers boat what their. There Turkmen moreover one Lebanese to brace. These shower in it everybody should whatever. Those uninterested several nobody yours thankful innocent.
	//
	// Power to any from its few luxury. None boy religion themselves it there might. Must near theirs mine thing tonight outside. Rapidly spotted solemnly finally been did confusion. These were nobody early silently company quarterly.
	//
	// American there time many French a each. Anybody rather paint ours did tonight remove. First including chair your at in by. Cackle whose they yearly wisdom nightly nightly. When finally without oxygen scold what silence..
	//
	// </details>
	//
	// ---
	//
	//

}

func TestMarkdownLookup(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("markdown")

	m := MapParams{
		"sections_count": {"2"},
	}

	value, err := info.Generate(f.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	tests := map[string]string{
		"Test 1": "[http://www.humannext-generation.com/systems/distributed/seize/users](http://www.humannext-generation.com/systems/distributed/seize/users)",
		"Test 2": "Them. Batch. Its. Iraqi.",
		"Test 3": "How there gorgeous genetics time choir fiction. Therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his.",
		"Test 4": "Power to any from its few luxury. None boy religion themselves it there might. Must near theirs mine thing",
		"Test 5": "hat silence..",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value.(string), v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value.(string))
		}
	}

}

func TestMarkdown(t *testing.T) {
	f := New(11)

	value, err := f.Markdown(&MarkdownOptions{SectionsCount: 3})
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "[http://www.humannext-generation.com/systems/distributed/seize/users](http://www.humannext-generation.com/systems/distributed/seize/users)",
		"Test 2": "His. Them. Those. Whose.",
		"Test 3": "How there gorgeous genetics time choir fiction. Therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his.",
		"Test 4": "1. ingeniously dive",
		"Test 5": "Why from somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple consequently",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestFakerMarkdownNoOptions(t *testing.T) {
	f := New(11)

	value, err := f.Markdown(nil)
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "[http://www.productextensible.info/leverage/integrated](http://www.productextensible.info/leverage/integrated)",
		"Test 2": "Otherwise the as our of school anthology trend host this whatever were their watch because",
		"Test 3": "Couple host much staff some whose reel themselves which. Can vanish rather these each previously these who few. Onto next he frighten",
		"Test 4": "1. silently fight the fact",
		"Test 5": "Us we it engine on hurt everyone cardigan who us above each longue bravery childhood host",
	}

	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestMarkdownNoOptions(t *testing.T) {
	Seed(11)

	value, err := Markdown(nil)
	if err != nil {
		t.Error(err)
	}

	tests := map[string]string{
		"Test 1": "[http://www.productextensible.info/leverage/integrated](http://www.productextensible.info/leverage/integrated)",
		"Test 2": "Otherwise the as our of school anthology trend host this whatever were their watch because",
		"Test 3": "Couple host much staff some whose reel themselves which. Can vanish rather these each previously these who few. Onto next he frighten",
		"Test 4": "1. silently fight the fact",
		"Test 5": "Us we it engine on hurt everyone cardigan who us above each longue bravery childhood host",
	}

	for k, v := range tests {
		if !strings.Contains(value, v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value)
		}
	}
}

func TestMarkdownLookupNoOptions(t *testing.T) {
	f := New(11)

	info := GetFuncLookup("markdown")

	value, err := info.Generate(f.Rand, nil, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	tests := map[string]string{
		"Test 1": "[http://www.humannext-generation.com/systems/distributed/seize/users](http://www.humannext-generation.com/systems/distributed/seize/users)",
		"Test 2": "Golang you will need to install",
		"Test 3": "[http://www.futuredistributed.io/morph/sticky](http://www.futuredistributed.io/morph/sticky)",
		"Test 4": "[http://www.legacyinfrastructures.io/plug-and-play](http://www.legacyinfrastructures.io/plug-and-play)",
	}
	//println(value.(string))
	for k, v := range tests {
		if !strings.Contains(value.(string), v) {
			t.Errorf("Test: %s could not find `%s`, in result %s", k, v, value.(string))
		}
	}

}

// Template Benchmarks

func BenchmarkTemplate100(b *testing.B) {
	f := New(11)

	for i := 0; i < 100; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)

	for i := 0; i < 1000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup10000(b *testing.B) {
	f := New(11)

	for i := 0; i < 10000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup100000(b *testing.B) {
	f := New(11)

	for i := 0; i < 100000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}
