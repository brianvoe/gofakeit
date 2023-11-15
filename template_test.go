package gofakeit

import (
	"fmt"
	"strings"
	"testing"
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

		value, err := f.Template(v, nil)
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

	value, err := Template("{{range $y := IntRange 1 .Data}}\n{{FirstName}} {{LastName}}{{end}}", &TemplateOptions{Data: 4})
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
	value, err := f.Template("{{range $y := IntRange 1 .Data}}\n{{FirstName}} {{LastName}}{{end}}", &TemplateOptions{Data: 4})
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
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template")

	m := MapParams{
		"template": {"{{range $y := IntRange 1 (Int (.Data))}}{{FirstName}} {{LastName}}\n{{end}}"},
		"data":     {"5"},
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
	value, err := f.Template("", nil)

	if err == nil {
		t.Error("Was expecting an error ", err)
	}
	if string(value) != "" {
		t.Errorf("expected a nothing, got %s", string(value))
	}
}

func ExampleEmailText() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := EmailText(&EmailOptions{Sections_count: 5})
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

func ExampleFaker_EmailText() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.EmailText(&EmailOptions{Sections_count: 5})
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

func TestEmailTextLookup(t *testing.T) {
	faker := New(6)
	globalFaker.Rand.Seed(6)
	info := GetFuncLookup("email_text")

	m := MapParams{
		"sections_count": {"3"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Marfa deep v kitsch helvetica cold-pressed paleo before they sold out") {
		t.Error("Expected `Marfa deep v kitsch helvetica cold-pressed paleo before they sold out`, got ", value)
	}

}

func TestEmailText(t *testing.T) {
	f := New(5)
	globalFaker.Rand.Seed(5)
	value, err := f.EmailText(&EmailOptions{Sections_count: 6})
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// TemplateMarkdown examples and tests

func ExampleMarkdown() {
	// Make sure we get the same results every time
	Seed(5)
	globalFaker.Rand.Seed(5)

	value, err := Markdown(&MarkdownOptions{Sections_count: 3})
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
	// [https://www.directembrace.biz/deliver](https://www.directembrace.biz/deliver)
	// [http://www.dynamicrevolutionize.net/vortals/action-items/enhance/platforms](http://www.dynamicrevolutionize.net/vortals/action-items/enhance/platforms)
	// [https://www.internationale-tailers.net/viral](https://www.internationale-tailers.net/viral)
	// [http://www.districtrobust.biz/monetize](http://www.districtrobust.biz/monetize)
	// [https://www.districtrepurpose.biz/whiteboard](https://www.districtrepurpose.biz/whiteboard)
	// [https://www.humannetworks.biz/robust/b2b/open-source/reintermediate](https://www.humannetworks.biz/robust/b2b/open-source/reintermediate)
	// [http://www.principalextensible.org/implement/engage](http://www.principalextensible.org/implement/engage)
	// [https://www.chiefdynamic.info/cutting-edge](https://www.chiefdynamic.info/cutting-edge)
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
	// ## Details
	//
	// Spit that cut his whom occasion themselves thankful early archipelago stand the slavery too whichever to sometimes look why pout yourself empty utterly ingeniously itself lastly stack thing which. Deeply whose group later whose you daily totally fully where theater whom that girl ride forest our eyes off effect favor do this each stormy little little school eye. His regularly always yet vanish who my whom flock it daily read besides ours hourly couple those work of much in something whichever nobody cluster entertain repeatedly myself door.
	//
	// Which window behind your outfit limp ask decidedly wisdom normally mine first Turkmen strongly album any theirs mob sparrow ourselves that before for exactly heavily out American week nightly. Himself promptly hers tunnel film student neither mine up whereas i.e. as after will reel for murder those last most afterwards kid secondly anger case rapidly world himself election. Paralyze coffee because daily scold contrast anyone mustering contrary why did whom busily might now you for that totally will meanwhile then reassure begin you why belief yesterday which.
	//
	// From rather annually library generally well pack which quaint fleet together his to how furthermore humour what time since wisdom each anyway factory problem couple but am that am. That several inadequately e.g. that those whomever those behind glamorous failure previously tonight besides away little he gold yet each far nurse up fly some ever has bill account. Earlier there most everyone yesterday abundant fondly staff suspiciously secondly another too whoever will tonight otherwise sleep in win secondly a then any brother then which company one his.
	//
	// Soak since head from her these then yesterday you there day those win whomever kiss moreover finally so regularly this how several orchard laptop defiant arrive could may aside. Could late how backwards but weekly Madagascan in without also ski luxuty never whomever scarcely e.g. above why spit few soon goat many whole throughout just on tenderly grab. Line curios tomorrow decidedly hence lie company e.g. enthusiasm been party already work what finally it sprint weather jump many tomorrow then she joy result flock anger somewhat generally.
	//
	// <details>
	// <summary>Repulsive sedge stupidly ski. </summary>
	//
	// What pack peace stemmed but hungrily weekly hers child sleepy daily. Then bow until her her work army our unless anyway from. Clean rather theirs today turn first hotel one daily why before.
	//
	// Distinct therefore those other American your herself trend father i.e. another. Gauva all motor whichever crew later zebra anyone does whatever fiction. Other government all other addition above its eventually juice boldly after.
	//
	// Fun she pride in for reluctantly pleasure plant since poverty next. These ours knit ear hourly theirs troop float include whose belong. Whomever whose their sometimes to string why choir whose tomorrow of.
	//
	// Of this tonight themselves brave I for i.e. ours openly distinguish. Out hers Putinist being before our quantity before who then right. Something Finnish one drink as mine those stand should these onto.
	//
	// Few straightaway each itself with handle little she somebody off guitar. Seldom board of e.g. even yours wisp down want politely case. Mine poor go herself him mustering it they mine batch group..
	//
	// </details>
	//
	// ---
	//
	//
}

func ExampleFaker_Markdown() {
	// Make sure we get the same results every time
	f := New(5)
	globalFaker.Rand.Seed(5)
	value, err := f.Markdown(&MarkdownOptions{Sections_count: 2})
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
	// [https://www.directembrace.biz/deliver](https://www.directembrace.biz/deliver)
	// [http://www.dynamicrevolutionize.net/vortals/action-items/enhance/platforms](http://www.dynamicrevolutionize.net/vortals/action-items/enhance/platforms)
	// [https://www.internationale-tailers.net/viral](https://www.internationale-tailers.net/viral)
	// [http://www.districtrobust.biz/monetize](http://www.districtrobust.biz/monetize)
	// [https://www.districtrepurpose.biz/whiteboard](https://www.districtrepurpose.biz/whiteboard)
	// [https://www.humannetworks.biz/robust/b2b/open-source/reintermediate](https://www.humannetworks.biz/robust/b2b/open-source/reintermediate)
	// [http://www.principalextensible.org/implement/engage](http://www.principalextensible.org/implement/engage)
	// [https://www.chiefdynamic.info/cutting-edge](https://www.chiefdynamic.info/cutting-edge)
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

}

func TestMarkdownLookup(t *testing.T) {
	faker := New(9)
	globalFaker.Rand.Seed(9)
	info := GetFuncLookup("markdown")

	m := MapParams{
		"sections_count": {"2"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Nevertheless behind am though lead regiment hers where") {
		t.Error("Expected `Nevertheless behind am though lead regiment hers where`, got ", value)
	}

}

func TestMarkdown(t *testing.T) {
	f := New(3)
	globalFaker.Rand.Seed(1)
	value, err := f.Markdown(&MarkdownOptions{Sections_count: 1})
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
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 1000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup10000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 10000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup100000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 100000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Data}}{{FirstName}} {{LastName}}\n{{end}}", &TemplateOptions{Data: 5})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func TestExa(t *testing.T) {
	f := New(3)
	globalFaker.Rand.Seed(1)

	// Accessing the Lines variable from within the template.
	template := `{{range $y := IntRange 1 .Data}}`

	// Example of using a helper functions to build a List of strings and pass it to a gofakeit function, Nested function need to be wrapped in ().
	template += `{{RandomString (ListS "Contact Details" "Customers Details")}}{{$p:=Person}}`

	// Displaying values from a variable, and using Upper helper function to format Company name.
	template += `Name: {{$p.FirstName}} {{$p.LastName}}\nCompany: {{Upper ($p.Job.Company)}}\nContact: {{$p.Contact.Email}}{{end}}`

	value, err := f.Template(template, &TemplateOptions{Data: 5})

	if err != nil {
		t.Error(err)
	}

	println(string(value))
	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}
