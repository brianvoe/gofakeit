package gofakeit

import "fmt"

func ExampleJSON_object() {
	Seed(11)

	value, err := JSON(&JSONOptions{
		Type: "object",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
		},
		Whitespace: true,
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
		RowCount:   3,
		Whitespace: true,
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
