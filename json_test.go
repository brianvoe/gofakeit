package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleJSON_object() {
	Seed(11)

	value, err := JSON(&JSONOptions{
		Type: "object",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
		},
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: {
	//     "address": {
	//         "address": "185 Villetown, Aleenside, Illinois 84157",
	//         "street": "185 Villetown",
	//         "city": "Aleenside",
	//         "state": "Illinois",
	//         "zip": "84157",
	//         "country": "Solomon Islands",
	//         "latitude": -53.426363,
	//         "longitude": -50.08629
	//     },
	//     "first_name": "Bart",
	//     "last_name": "Beatty"
	// }
}

func ExampleJSON_array() {
	Seed(11)

	value, err := JSON(&JSONOptions{
		Type: "array",
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
		RowCount: 3,
		Indent:   true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: [
	//     {
	//         "first_name": "Markus",
	//         "id": 1,
	//         "last_name": "Moen"
	//     },
	//     {
	//         "first_name": "Alayna",
	//         "id": 2,
	//         "last_name": "Wuckert"
	//     },
	//     {
	//         "first_name": "Lura",
	//         "id": 3,
	//         "last_name": "Lockman"
	//     }
	// ]
}

func TestJSONLookup(t *testing.T) {
	info := GetFuncLookup("json")

	m := map[string][]string{
		"type":     {"array"},
		"rowcount": {"10"},
		"fields":   {
			`{"name":"id","function":"autoincrement"}`,
			`{"name":"first_name","function":"firstname"}`
		},
	}
	value, err := info.Call(&m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Fatal(fmt.Sprintf("%s", value.([]byte)))
}
