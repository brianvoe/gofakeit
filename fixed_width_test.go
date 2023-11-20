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

	// Output:
	// Name               Email                          Password         Age
	// Markus Moen        sylvanmraz@murphy.net          6VlvH6qqXc7g     13
	// Alayna Wuckert     santinostanton@carroll.biz     g7sLrS0gEwLO     46
	// Lura Lockman       zacherykuhic@feil.name         S8gV7Z64KlHG     12
}

func ExampleFixedWidth_default() {
	Seed(11)

	value, err := FixedWidth(nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Name             Email                        Password
	// Marcel Pagac     anibalkozey@lockman.name     X3ZoWoYkWA6L
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

	// Output:
	// Markus Moen        sylvanmraz@murphy.net          6VlvH6qqXc7g     13
	// Alayna Wuckert     santinostanton@carroll.biz     g7sLrS0gEwLO     46
	// Lura Lockman       zacherykuhic@feil.name         S8gV7Z64KlHG     12
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

	// Output:
	// Name               Email                          Password         Age
	// Markus Moen        sylvanmraz@murphy.net          6VlvH6qqXc7g     13
	// Alayna Wuckert     santinostanton@carroll.biz     g7sLrS0gEwLO     46
	// Lura Lockman       zacherykuhic@feil.name         S8gV7Z64KlHG     12
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

	output, err := info.Generate(faker.Rand, &m, info)
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
				`{"name":"Name","function":"{firstname} {lastname}"}"}`,
				`{"name":"Email","function":"email"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Age","function":"{number:1,100}"}`,
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
				`{"name":"Name","function":"{firstname} {lastname}"}"}`,
				`{"name":"Email","function":"email"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Age","function":"{number:1,100}"}`,
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
				`{"name":"Name","function":"{firstname} {lastname}"}"}`,
				`{"name":"Email","function":"email"}`,
				`{"name":"Password","function":"password","params":{"special":["false"],"length":["20"]}}`,
				`{"name":"Age","function":"{number:1,100}"}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
