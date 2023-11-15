package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleFixedWidth() {
	Seed(11)

	value, err := FixedWidth(&FixedWidthOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "name", Function: "{{FirstName}} {{LastName}}",
				Params: MapParams{
					"spacing":    {"15"},
					"header_pad": {"*"}}},
			{Name: "last_name", Function: "lastname",
				Params: MapParams{
					"spacing":    {"15"},
					"header_pad": {"*"}}},
			{Name: "Email", Function: "email",
				Params: MapParams{
					"spacing":    {"30"},
					"header_pad": {"*"}}},
			{Name: "password", Function: "password",
				Params: MapParams{
					"special":    {"false"},
					"spacing":    {"20"},
					"header_pad": {"*"}}},
			{Name: "Account No.", Function: "{{AchAccount}}",
				Params: MapParams{
					"spacing":    {"20"},
					"header_pad": {"*"}}},
			{Name: "Money", Function: "{{Number 1 100}}",
				Params: MapParams{
					"footer":     {"{{.GetTotal}}"},
					"spacing":    {"10"},
					"align":      {">"},
					"footer_pad": {"0"},
					"header_pad": {"*"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// name***********last_name******Email*************************password************Account No.**************Money
	// Markus Moen    Pagac          anibalkozey@lockman.name      X3ZoWoYkWA6L        302320202761                62
	// Estell Fay     Marvin         hardyhintz@crooks.io          wWq8lxcb4woV        389344090143                17
	// Lillie Kuhn    Homenick       hannakassulke@prosacco.name   lJKBT6Pcs9tr        614968021002                87
	//                                                                                                     0000166.00
}

func ExampleFaker_FixedWidth() {

	f := New(11)
	globalFaker.Rand.Seed(11)

	value, err := f.FixedWidth(&FixedWidthOptions{
		RowCount: 3,
		Fields: []Field{
			{Name: "name", Function: "{{FirstName}} {{LastName}}", Params: MapParams{"spacing": {"-1"}, "header_pad": {"*"}}},
			{Name: "last_name", Function: "lastname", Params: MapParams{"spacing": {"-1"}, "header_pad": {"*"}}},
			{Name: "Email", Function: "email", Params: MapParams{"spacing": {"-1"}, "header_pad": {"*"}}},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}, "spacing": {"-1"}, "header_pad": {"*"}}},
			{Name: "Account No.", Function: "{{AchAccount}}", Params: MapParams{"spacing": {"-1"}, "header_pad": {"*"}}},
			{Name: "Money", Function: "{{Number 1 100}}", Params: MapParams{"footer": {"{{.GetTotal}}"}, "spacing": {"-1"}, "align": {">"}, "footer_pad": {"0"}, "header_pad": {"*"}}},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// name************last_nameEmail********************password****Account No.**Money
	// Markus Moen     Daniel   marcelpagac@wuckert.biz  W8DAkpLjYoBW364599489953    35
	// Amie Feil       Kuvalis  trystangislason@pagac.net9G9rOTgd3vDs635300425914    16
	// Trystan GislasonKeebler  gabriellehuels@borer.io  J41S7H76KwYZ323202027613    61
	//                                                                           112.00
}

func TestFixedWidthLookup(t *testing.T) {
	faker := New(0)

	info := GetFuncLookup("fixed_width")

	m := MapParams{
		"rowcount": {"10"},
		"fields": {
			`{"name":"Name","function":"{{FirstName}} {{LastName}}"}`,
			`{"name":"Last Name","function":"lastname"}`,
			`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
			`{"name":"Account No.","function":"{{AchAccount}}"}`,
			`{"name":"Money","function":"{{Number 1 100}}",   
				"params":{
					"align": ">"
				}
			}`,
		},
	}

	output, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	value := string(output.([]byte))

	// Check that value has the correct number of rows via new line characters plus 1 for the header
	if strings.Count(value, "\n") != 11 {
		t.Error("Expected 10+1(header row footer) rows, got", strings.Count(value, "\n"))
	}

}

func TestFixedWidthNoOptions(t *testing.T) {
	Seed(11)

	// if CSVOptions is nil -> get a random CSVOptions
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
				`{"name":"Name","function":"{{FirstName}} {{LastName}}"}`,
				`{"name":"Last Name","function":"lastname"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Account No.","function":"{{AchAccount}}"}`,
				`{"name":"Money","function":"{{Number 1 100}}",   
					"params":{
						"align": ">"
					}
				}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkFixedWidthLookup1000(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("fixed_width")
		m := MapParams{
			"rowcount": {"1000"},
			"fields": {
				`{"name":"Name","function":"{{FirstName}} {{LastName}}"}`,
				`{"name":"Last Name","function":"lastname"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Account No.","function":"{{AchAccount}}"}`,
				`{"name":"Money","function":"{{Number 1 100}}",   
					"params":{
						"align": ">"
					}
				}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkFixedWidthLookup10000(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("fixed_width")
		m := MapParams{
			"rowcount": {"10000"},
			"fields": {
				`{"name":"Name","function":"{{FirstName}} {{LastName}}"}`,
				`{"name":"Last Name","function":"lastname"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Account No.","function":"{{AchAccount}}"}`,
				`{"name":"Money","function":"{{Number 1 100}}",   
					"params":{
						"align": ">"
					}
				}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
