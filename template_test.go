package gofakeit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6/data"
)

// Misc tests
func TestTemplateFunctionsWithSlices(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	test := map[string]string{
		"Weighted":     "{{Weighted (ListI `hello` 2 6.9) (ListF32 1 2 3)}}",
		"Dice":         "{{ Dice 3 (ListUInt 1 5 3) }}",
		"RandomInt":    "{{ RandomInt (ListInt 1 5 3) }}",
		"RandomString": "{{RandomString (ListS `key1` `key2` `key3`)}}",
		"RandomUint":   "{{RandomUint (ListUInt 2 6 9)}}",
		"Teams":        "{{Teams (ListS `person_a` `person_b` `person_c`) (ListS `team_a` `team_b` `team_c`)}}",
	}

	for k, v := range test {

		value, err := f.Template(v, 1)
		if err != nil {
			t.Error(k, err)
		}

		if string(value) == "" {
			t.Errorf("expected a value for %s, got nothing", k)
		}
	}
}

// Template examples and tests

func ExampleTemplate() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := Template("{{range $y := IntRange 1 .Lines}}\n{{FirstName}} {{LastName}}{{end}}", 4)
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
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template("{{range $y := IntRange 1 .Lines}}\n{{FirstName}} {{LastName}}{{end}}", 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Markus Moen
	// Alayna Wuckert
	// Lura Lockman

}

func TestTemplateLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template")

	m := MapParams{
		"template": {"{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}"},
		"lines":    {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if value != "Markus Moen\nAlayna Wuckert\nLura Lockman\nSylvan Mraz\nPaolo Rutherford\n" {
		t.Error("Expected `Markus Moen Pagac`, got ", value)
	}

}

func TestTemplateNoTemplateParam(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template("", 3)

	if err != nil {
		t.Error("Error with template ", err)
	}
	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

func ExampleTemplateDocument() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateDocument(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Good evening, Alayna
	//
	// you have look now that regularly Happens tonight
	//
	// I appreciate your input
	// Freida
	//
	// Phone: 4185959586
	// Email: rosannachristiansen@klocko.io

}

func ExampleFaker_TemplateDocument() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateDocument(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Hi, Marcel
	//
	// you have look now that regularly Happens tonight
	//
	// I appreciate your input
	// Freida
	//
	// Phone: 4185959586
	// Email: rosannachristiansen@klocko.io

}

func TestTemplateDocumentLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template_document")

	m := MapParams{
		"sections": {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Power to any from its few luxury none boy religion") {
		t.Error("Expected `Power to any from its few luxury none boy religion`, got ", value)
	}

}

func ExampleTemplateEmailText() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateEmailText(5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Good evening, Alayna
	//
	// you have look now that regularly Happens tonight
	// Gentrify pitchfork stumptown mlkshk umami chambray blue bottle 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan. Hoodie +1 blue bottle fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk shabby chic. Typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it raw denim cliche dreamcatcher pug fixie.
	//
	// Park trust fund migas fingerstache sriracha +1 mustache tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid. Stumptown organic schlitz flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy chia. Mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade deep v green juice keffiyeh carry poutine.
	//
	// Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat cornhole everyday taxidermy humblebrag bespoke kitsch. Hella venmo umami retro five dollar toast twee bitters microdosing cold-pressed street chambray distillery kale chips disrupt actually typewriter. Lomo kinfolk tilde heirloom normcore farm-to-table master readymade schlitz shoreditch bitters kitsch marfa beard roof butcher.
	// Pinterest semiotics photo booth taxidermy offal tofu whatever schlitz meh?
	// Themselves been group flock shake part purchase up usually it her none it hers boat what their there Turkmen. Moreover one Lebanese to brace these shower in it everybody should whatever those uninterested several nobody yours thankful innocent. Power to any from its few luxury none boy religion themselves it there might must near theirs mine thing. Tonight outside rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly American there time. Many French a each anybody rather paint ours did tonight remove first including chair your at in by cackle.
	//
	// Whose they yearly wisdom nightly nightly when finally without oxygen scold what silence go time behind example me soon. Grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother. Instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary from belief this upon. At most homeless elsewhere has yearly under since where Californian all in today generally myself after stupid highly heavy. Here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me tonight place.
	//
	// Few why from somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple. Consequently in those daily had anywhere anything what in bouquet which annually as Cypriot this our sand tightly we. First their staff invention however whoever itself over this pair smoke yourself so circumstances despite could project did embarrassed. My philosophy imagination did stemmed mob furthermore read myself above had for successful for move here point less to. Myself sparse ours yours because fiercely jump place exist board already highly either monthly pose that our punctuation when.
	// Heirloom chillwave disrupt trust fund iPhone plaid selvage yr quinoa tousled bespoke Yuccie.
	// Cheers
	// Paul Wolff
	//
	// Company: Revaluate
	// Job Role: Administrator
	//
	// Phone: 4496711988
	// Email: matildajohnston@von.info

}

func ExampleFaker_TemplateEmail() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateEmailText(6)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Hi, Marcel
	//
	// you have look now that regularly Happens tonight
	// Gentrify pitchfork stumptown mlkshk umami chambray blue bottle 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan. Hoodie +1 blue bottle fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk shabby chic. Typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it raw denim cliche dreamcatcher pug fixie.
	//
	// Park trust fund migas fingerstache sriracha +1 mustache tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid. Stumptown organic schlitz flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy chia. Mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade deep v green juice keffiyeh carry poutine.
	//
	// Flannel direct trade keffiyeh cold-pressed tilde plaid keytar tilde bespoke waistcoat cornhole everyday taxidermy humblebrag bespoke kitsch. Hella venmo umami retro five dollar toast twee bitters microdosing cold-pressed street chambray distillery kale chips disrupt actually typewriter. Lomo kinfolk tilde heirloom normcore farm-to-table master readymade schlitz shoreditch bitters kitsch marfa beard roof butcher.
	// Pinterest semiotics photo booth taxidermy offal tofu whatever schlitz meh?
	// Themselves been group flock shake part purchase up usually it her none it hers boat what their there Turkmen. Moreover one Lebanese to brace these shower in it everybody should whatever those uninterested several nobody yours thankful innocent. Power to any from its few luxury none boy religion themselves it there might must near theirs mine thing. Tonight outside rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly American there time. Many French a each anybody rather paint ours did tonight remove first including chair your at in by cackle.
	//
	// Whose they yearly wisdom nightly nightly when finally without oxygen scold what silence go time behind example me soon. Grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother. Instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary from belief this upon. At most homeless elsewhere has yearly under since where Californian all in today generally myself after stupid highly heavy. Here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me tonight place.
	//
	// Few why from somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple. Consequently in those daily had anywhere anything what in bouquet which annually as Cypriot this our sand tightly we. First their staff invention however whoever itself over this pair smoke yourself so circumstances despite could project did embarrassed. My philosophy imagination did stemmed mob furthermore read myself above had for successful for move here point less to. Myself sparse ours yours because fiercely jump place exist board already highly either monthly pose that our punctuation when.
	// Heirloom chillwave disrupt trust fund iPhone plaid selvage yr quinoa tousled bespoke Yuccie."Letterpress meditation gentrify." - Marcelo Mertz
	//
	// Thank you
	// Myah
	//
	// Address: 81575 North Daleton, Chicago, New Jersey 44967
	// City: Chicago
	// State: New Jersey
	// Zip: 44967Phone: 8821512123
	// Email: josianebartell@crooks.biz
}

func TestTemplateEmailLookup(t *testing.T) {
	faker := New(6)
	globalFaker.Rand.Seed(6)
	info := GetFuncLookup("template_email")

	m := MapParams{
		"sections": {"3"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Marfa deep v kitsch helvetica cold-pressed paleo before they sold out") {
		t.Error("Expected `Marfa deep v kitsch helvetica cold-pressed paleo before they sold out`, got ", value)
	}

}

func TestTemplateEmail(t *testing.T) {
	f := New(5)
	globalFaker.Rand.Seed(5)
	value, err := f.TemplateEmailText(6)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// TemplateMarkdown examples and tests

func ExampleTemplateMarkdown() {
	// Make sure we get the same results every time
	Seed(5)
	globalFaker.Rand.Seed(5)

	value, err := TemplateMarkdown(3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	//
	// # Paragraph
	//
	// Everyone. These. Kindness. Irritate. The.
	//
	// Most. We. That. Tighten. Burkinese.
	//
	// Regularly. World. Hitlerian. Nearby. For.
	//
	// Nobody. What. Calm. Is. Be.
	//
	// ---
	//
	// # Block Quote
	//
	// My powerless of turn above nest. Generally Honduran left you his from. Notebook none who when myself too.
	//
	// Just even were there others equally. There has munch none Mexican annually. Behind pain within heat bread why.
	//
	// ---
	//
	// ## Details
	//
	// Above progress to enthusiastically next wisp his whose. Why yet sleep those those hang positively tomorrow. Awfully appear e.g. ours band yesterday here did. In his number I in as to what.
	//
	// <details>
	// <summary>A forest wearily read a trench coat. </summary>
	//
	// ```
	// go install https://www.leaddeploy.info/generate/bricks-and-clicks/aggregate/visionary
	// ```
	//
	// </details>
	//
	// ---
	//
	//
}

func ExampleFaker_TemplateMarkdown() {
	// Make sure we get the same results every time
	f := New(5)
	globalFaker.Rand.Seed(5)
	value, err := f.TemplateMarkdown(2)
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
	// [https://www.globalevolve.org/intuitive](https://www.globalevolve.org/intuitive)
	// [https://www.principalsynergies.info/enhance/platforms/cross-platform](https://www.principalsynergies.info/enhance/platforms/cross-platform)
	// [http://www.seniorstreamline.io/functionalities/incubate/e-tailers/utilize](http://www.seniorstreamline.io/functionalities/incubate/e-tailers/utilize)
	// [http://www.productvirtual.org/interfaces/visionary/whiteboard/viral](http://www.productvirtual.org/interfaces/visionary/whiteboard/viral)
	// [https://www.corporateleverage.org/b2b/open-source](https://www.corporateleverage.org/b2b/open-source)
	// [https://www.customermethodologies.net/iterate](https://www.customermethodologies.net/iterate)
	// [http://www.legacyincubate.info/interactive](http://www.legacyincubate.info/interactive)
	//
	// ---
	//
	// ## LISTS
	//
	// Then turn above nest generally Honduran left you his from notebook none who when myself too just even were there others equally there. Has munch none Mexican annually behind pain within heat bread why before house has Eastern wave this how enthusiastically next wisp his whose. Why yet sleep those those hang positively tomorrow awfully appear e.g. ours band yesterday here did in his number I in as to.
	//
	// What without with whose mine accordingly thoughtfully rubbish nightly whereas she himself switch hundreds sufficient work indeed from over school so when part. How of desktop unless do each those occasionally hat lips out single bunch ourselves easily us mirror rarely often much additionally will there. Its envy what army she boy decidedly knit him group barely abundant does massage did Cambodian keep annually elsewhere them her cackle quarterly.
	//
	// Been decidedly today utterly been paint firstly remain faithful Brazilian example though still Belgian trench I these carry example them chaos for elsewhere. Annually simply thing almost off Ecuadorian loudly everybody annually now late out whose Mayan confusion yours it stand snore bevy then in my. Gate what choir be a of how conclude solemnly herself dig anyway for growth wait at you instance to besides by whose elsewhere..
	//
	// 1. fight joyously
	// 1. dream decidedly perfectly
	// 1. fully equally climb awfully regularly
	// 1. hug fully beautifully by uptight batch
	// 1. talk boldly from bored group
	// 1. snore practically sensibly into a literature
	// 1. drink noisily with a healthy nest
	// 1. listen a clear gang
	// 1. solemnly dance
	//
	// ---
	//
	//

}

func TestTemplateMarkdownLookup(t *testing.T) {
	faker := New(9)
	globalFaker.Rand.Seed(9)
	info := GetFuncLookup("template_markdown")

	m := MapParams{
		"sections": {"2"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Then teacher. Our sneeze. Offend you. Of usually. Your now.") {
		t.Error("Expected `Then teacher. Our sneeze. Offend you. Of usually. Your now.`, got ", value)
	}

}

func TestTemplateMarkdown(t *testing.T) {
	f := New(3)
	globalFaker.Rand.Seed(1)
	value, err := f.TemplateMarkdown(1)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// Template Test all Templates make sure no syntax errors
func TestTemplateAll(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	// combine all templates

	build_str := ""
	for k := range data.Template {
		for _, template := range data.Template[k] {
			value := strings.ReplaceAll(template, "'", "`")
			value = strings.ReplaceAll(value, "\\n", "\n")
			value = strings.ReplaceAll(value, "|n", "\\n")
			build_str = build_str + value
		}
	}

	value, err := f.Template(build_str, 1)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// Template Benchmarks

func BenchmarkTemplate100(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 100; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 1000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup10000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 10000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup100000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 100000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}
