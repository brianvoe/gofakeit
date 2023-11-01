package gofakeit

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6/data"
)

func ExampleTemplate() {
	//Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := Template(&TemplateOptions{
		//Template: data.MARKDOWN_TEMPLATE,
		TemplateType: "email",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// From: Markus Moen <luralockman@jakubowski.com>
	// To: Alayna Wuckert <paolorutherford@armstrong.org>
	// Subject: the rest is history
	// Date: 1926-07-07 22:25:40.559683317 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Alayna Wuckert
	//
	// Weight government hourly constantly yesterday someone him regularly today whomever this revolt hence from his timing as quantity us these. Yours live these frantic not may another how this ours his them those whose them batch its Iraqi most that. Few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow. This that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would leap.
	//
	// Plane us first then down them eager would hundred super throughout animal yet themselves been group flock shake part purchase. Up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower in. It everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none boy. Religion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these. Were nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight remove.
	//
	// First including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what. Silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child. Annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever. Contrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally myself. After stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e..
	//
	// Regards Markus Moen

}

func ExampleFaker_Template() {
	//Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template(&TemplateOptions{
		Template: data.GetTemplateData("fixed_format"),
		Lines:    4,
	})

	if err != nil {
		fmt.Println(err)
	}
	v := string(value)
	fmt.Println(v)

	// Output:
	// FIRSTNAME LASTNAME  Gender  Job            Hobby               CreditCard          Address
	// Markus    Moen      male    Developer      Lacrosse            39800889982276      369 North Cornerbury, Miami, North Dakota 24259
	// Matteo    Krajcik   female  Manager        Tour skating        569586614968023     4090 Lake Unionport, Pittsburgh, Iowa 73720
	// Modesta   Hilpert   female  Representative Volleyball          6495603064524543520 976 Forgeport, Aurora, Minnesota 92470
	// Reyes     Murphy    female  Representative Orienteering        4039550319361326    7752 Parkwayville, Virginia Beach, Pennsylvania 76241
}

func TestTemplatePadFunction(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template(&TemplateOptions{
		Template: "{{Pad (Name) 10 `-` `RIGHT` true}}",
	})

	if err != nil {
		fmt.Println(err)
	}

	if string(value) != "Markus Moe" {
		t.Error("Expected `Markus Moe`, got ", value)
	}
}

func TestTemplateLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template")

	m := MapParams{
		"template":      {"{{Name}} {{LastName}}"},
		"template_type": {""},
		"lines":         {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if value != "Markus Moen Pagac" {
		t.Error("Expected `Markus Moen Pagac`, got ", value)
	}

}

func TestTemplateNoOptions(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	_, err := f.Template(nil)

	if err != nil {
		t.Error("Error with template ", err)
	}

}

func TestTemplateEmail(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	_, err := f.Template(&TemplateOptions{
		TemplateType: "email",
	})

	if err != nil {
		t.Error("Error with template ", err)
	}
}

func TestTemplateMD(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	_, err := f.Template(&TemplateOptions{
		TemplateType: "markdown",
	})

	if err != nil {
		t.Error("Error with template ", err)
	}
}

func TestTemplateHTML(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	_, err := f.Template(&TemplateOptions{
		TemplateType: "html",
	})

	if err != nil {
		t.Error("Error with template ", err)
	}
}

func TestTemplateFixedFormat(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	_, err := f.Template(&TemplateOptions{
		TemplateType: "fixed_format",
	})

	if err != nil {
		fmt.Println(err)
	}
}
func BenchmarkTemplate100(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	//Generate the function map once to speed up
	fm := BuildFunctionMap()

	for i := 0; i < 100; i++ {
		_, err := f.Template(&TemplateOptions{
			Template:    data.GetTemplateData("email"),
			Lines:       3,
			FunctionMap: fm,
		})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	//Generate the function map once to speed up
	fm := BuildFunctionMap()

	for i := 0; i < 1000; i++ {
		_, err := f.Template(&TemplateOptions{
			Template:    data.GetTemplateData("email"),
			Lines:       3,
			FunctionMap: fm,
		})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup10000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	//Generate the function map once to speed up
	fm := BuildFunctionMap()

	for i := 0; i < 10000; i++ {
		_, err := f.Template(&TemplateOptions{
			Template:    data.GetTemplateData("email"),
			Lines:       3,
			FunctionMap: fm,
		})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup100000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	//Generate the function map once to speed up
	fm := BuildFunctionMap()

	for i := 0; i < 100000; i++ {
		_, err := f.Template(&TemplateOptions{
			Template:    data.GetTemplateData("email"),
			Lines:       3,
			FunctionMap: fm,
		})
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}
