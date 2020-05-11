package gofakeit

import "fmt"

func ExampleXML_single() {
	Seed(11)

	value, err := XML(&XMLOptions{
		Type: "single",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
			{Name: "password", Function: "password", Params: map[string][]string{"special": {"false"}}},
		},
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: jhiuhuih
}
