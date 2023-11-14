package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

// Template examples and tests

func ExampleFixedWidth() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	opts := &FixedWidthOptions{
		Header:    []string{"Name", "Email", "Cost", "Account"},
		Row:       []string{"{{FirstName}} {{LastName}}", "{{Email}}", "{{Number 1 100}}", "{{AchAccount}}"},
		Footer:    []string{"", "", "{{.GetTotal}}", ""},
		HeaderPad: []string{" ", " ", " ", " "},
		RowPad:    []string{"", " ", " ", "0"},
		FooterPad: []string{" ", " ", " ", " "},
		Align:     []string{"<", "<", "<", ">"},
		Spacing:   []int{30, 30, 10, 20},
		Count:     4,
	}

	value, err := FixedWidth(opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Name                          Email                         Cost                   Account
	// Markus Moen-------------------alaynawuckert@kozey.biz       80        00000000948995369063
	// Rossie Mosciski---------------fredhickle@bahringer.org      41        00000000914583023202
	// Jalon Rolfson-----------------effiestiedemann@koelpin.biz   28        00000000171480088998
	// Freeda Keebler----------------gabriellehuels@borer.io       69        00000000932452944030
	//                                                             218.00
	//

}

func ExampleFaker_FixedWidth() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)

	opts := &FixedWidthOptions{
		Header:    []string{"Name", "Email", "Spend"},
		Row:       []string{"{{FirstName}} {{LastName}}", "{{Email}}", "{{Price 10 1000}}"},
		Footer:    []string{" ", " ", "{{.GetTotal}}"},
		HeaderPad: []string{"=", "=", "="},
		RowPad:    []string{" ", " ", " "},
		FooterPad: []string{" ", " ", " "},
		Align:     []string{"<", "<", ">"},
		Spacing:   []int{-1, -1, -1},
		Count:     4,
	}

	value, err := f.FixedWidth(opts)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Name                          Email                         Cost                   Account
	// Markus Moen                   alaynawuckert@kozey.biz       80        00000000948995369063
	// Rossie Mosciski               fredhickle@bahringer.org      41        00000000914583023202
	// Jalon Rolfson                 effiestiedemann@koelpin.biz   28        00000000171480088998
	// Freeda Keebler                gabriellehuels@borer.io       69        00000000932452944030
	//                                                             218.00
}

func TestFixedWidthLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("fixed_width")

	m := MapParams{
		"header":     {"Name", "Email", "Cost", "Account"},
		"row":        {"{{FirstName}} {{LastName}}", "{{Email}}", "{{Number 1 100}}", "{{AchAccount}}"},
		"footer":     {"", "", "{{.GetTotal}}", ""},
		"header_pad": {"", "", "", ""},
		"row_pad":    {"", "", "", ""},
		"footer_pad": {"", "", "", ""},
		"align":      {"<", "<", ">", ">"},
		"spacing":    {"30", "30", "10", "20"},
		"count":      {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	test := map[string]string{
		"header": "Name                          Email                         Cost      Account",
		"row1":   "Rossie Mosciski               fredhickle@bahringer.org      41        914583023202",
		"row2":   "Freeda Keebler                gabriellehuels@borer.io       69        932452944030",
		"Footer": "                                                            261.00              ",
	}

	for k, v := range test {
		if !strings.Contains(string(value.(string)), v) {
			t.Errorf("expected a value of %s for the %s, got nothing", v, k)
		}
	}

}

func TestFixedWidthDefaults(t *testing.T) {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)

	opts := &FixedWidthOptions{
		Header:  []string{"Name", "Email", "Cost", "Account"},
		Row:     []string{"{{FirstName}} {{LastName}}", "{{Email}}", "{{Number 1 100}}", "{{AchAccount}}"},
		Footer:  []string{"", "", "{{.GetTotal}}", ""},
		Spacing: []int{30, 30, 10, 20},
		Count:   4,
	}

	value, err := f.FixedWidth(opts)
	if err != nil {
		t.Error(err)
	}

	test := map[string]string{
		"header": "Name                          Email                         Cost      Account             ",
		"row1":   "Markus Moen                   alaynawuckert@kozey.biz       80        948995369063        ",
		"row2":   "Freeda Keebler                gabriellehuels@borer.io       69        932452944030        ",
		"Footer": "                                                            218.00                        ",
	}

	for k, v := range test {
		if !strings.Contains(string(value), v) {
			t.Errorf("expected a value of %s for the %s, got nothing", v, k)
		}
	}
}
func TestFixedWidthDefaultsNoFooter(t *testing.T) {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)

	opts := &FixedWidthOptions{
		Header:  []string{"Name", "Email", "Cost", "Account"},
		Row:     []string{"{{FirstName}} {{LastName}}", "{{Email}}", "{{Number 1 100}}", "{{AchAccount}}"},
		Spacing: []int{30, 30, 10, 20},
		Count:   4,
	}

	test := map[string]string{
		"header": "Name                          Email                         Cost      Account",
		"row1":   "Rossie Mosciski               fredhickle@bahringer.org      41        914583023202",
		"row2":   "Freeda Keebler                gabriellehuels@borer.io       69        932452944030",
	}
	value, err := f.FixedWidth(opts)
	if err != nil {
		t.Error(err)
	}

	for k, v := range test {
		if !strings.Contains(string(value), v) {
			t.Errorf("expected a value of %s for the %s, got nothing", v, k)
		}
	}
}

func TestFixedWidthFormatting(t *testing.T) {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)

	opts := &FixedWidthOptions{
		Header:    []string{"Name", "Email", "Cost", "Account"},
		Row:       []string{"{{FirstName}} {{LastName}}", "{{Email}}", "{{Number 1 100}}", "{{AchAccount}}"},
		Footer:    []string{"", "", "{{.GetTotal}}", ""},
		HeaderPad: []string{"=", "=", "=", "="},
		RowPad:    []string{"", " ", " ", "0"},
		FooterPad: []string{"*", "*", "*", "*"},
		Align:     []string{"<", "<", ">", ">"},
		Spacing:   []int{30, 30, 10, 20},
		Count:     4,
	}

	test := map[string]string{
		"header": "Name==========================Email===============================Cost=============Account",
		"row1":   "Markus Moen                   alaynawuckert@kozey.biz               8000000000948995369063",
		"row2":   "Freeda Keebler                gabriellehuels@borer.io               6900000000932452944030",
		"footer": "****************************************************************218.00********************",
	}
	value, err := f.FixedWidth(opts)
	if err != nil {
		t.Error(err)
	}

	for k, v := range test {
		if !strings.Contains(string(value), v) {
			t.Errorf("expected a value of %s for the %s, got nothing", v, k)
		}
	}
}

func TestFixedWidthNoOptions(t *testing.T) {
	// Make sure we get the same results every time
	f := New(5)
	globalFaker.Rand.Seed(5)

	test := map[string]string{
		"header": "Number_1_100First_Last_Name Price_10_1000First_Last_Name   Currency_ShortPrice_10_1000Price_10_1000Number_1_100",
		"row1":   "61          Anissa Stracke  600.78       Emelia Rodriguez             RON859.35       306.03       35         ",
		"row2":   "33          Demond Davis    672.19       Delpha Bartell               TVD618.74       580.06       74          ",
		"footer": "1199.00                     11163.36                                     10876.01     13253.33     1035.00     ",
	}
	value, err := f.FixedWidth(nil)
	if err != nil {
		t.Error(err)
	}

	for k, v := range test {
		if !strings.Contains(string(value), v) {
			t.Errorf("expected a value of %s for the %s, got nothing", v, k)
		}
	}
}
