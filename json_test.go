package gofakeit

import (
	"encoding/json"
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
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: {
	//     "first_name": "Markus",
	//     "last_name": "Moen",
	//     "address": {
	//         "address": "4599 Dale ton, Norfolk, New Jersey 36906",
	//         "street": "4599 Dale ton",
	//         "city": "Norfolk",
	//         "state": "New Jersey",
	//         "zip": "36906",
	//         "country": "Tokelau",
	//         "latitude": 23.058758,
	//         "longitude": 89.022594
	//     },
	//     "password": "qjXy56JHcVlZ"
	// }
}

func ExampleFaker_JSON_object() {
	f := New(11)

	value, err := f.JSON(&JSONOptions{
		Type: "object",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "address", Function: "address"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
		Indent: true,
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output: {
	//     "first_name": "Markus",
	//     "last_name": "Moen",
	//     "address": {
	//         "address": "4599 Dale ton, Norfolk, New Jersey 36906",
	//         "street": "4599 Dale ton",
	//         "city": "Norfolk",
	//         "state": "New Jersey",
	//         "zip": "36906",
	//         "country": "Tokelau",
	//         "latitude": 23.058758,
	//         "longitude": 89.022594
	//     },
	//     "password": "qjXy56JHcVlZ"
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
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
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
	//         "id": 1,
	//         "first_name": "Markus",
	//         "last_name": "Moen",
	//         "password": "Dc0VYXjkWABx"
	//     },
	//     {
	//         "id": 2,
	//         "first_name": "Osborne",
	//         "last_name": "Hilll",
	//         "password": "XPJ9OVNbs5lm"
	//     },
	//     {
	//         "id": 3,
	//         "first_name": "Mertie",
	//         "last_name": "Halvorson",
	//         "password": "eyl3bhwfV8wA"
	//     }
	// ]
}

func ExampleFaker_JSON_array() {
	f := New(11)

	value, err := f.JSON(&JSONOptions{
		Type: "array",
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
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
	//         "id": 1,
	//         "first_name": "Markus",
	//         "last_name": "Moen",
	//         "password": "Dc0VYXjkWABx"
	//     },
	//     {
	//         "id": 2,
	//         "first_name": "Osborne",
	//         "last_name": "Hilll",
	//         "password": "XPJ9OVNbs5lm"
	//     },
	//     {
	//         "id": 3,
	//         "first_name": "Mertie",
	//         "last_name": "Halvorson",
	//         "password": "eyl3bhwfV8wA"
	//     }
	// ]
}

func TestJSONLookup(t *testing.T) {
	faker := New(0)
	info := GetFuncLookup("json")

	m := NewMapParams()
	m.Add("type", "array")
	m.Add("rowcount", "10")
	m.Add("fields", `{"name":"id","function":"autoincrement"}`)
	m.Add("fields", `{"name":"first_name","function":"firstname"}`)
	m.Add("fields", `{"name":"password","function":"password","params":{"special":["false"],"length":["20"]}}`)

	_, err := info.Generate(faker.Rand, m, info)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestJSONObjectLookupWithSubJSON(t *testing.T) {
	faker := New(11)
	info := GetFuncLookup("json")

	m := NewMapParams()
	m.Add("type", "object")
	m.Add("fields", `{
		"name":"json",
		"function":"json",
		"params":{
			"type":"object",
			"fields":[
				{"name":"id","function":"autoincrement"},
				{"name":"first_name","function":"firstname"},
				{"name":"last_name","function":"lastname"},
				{"name":"password","function":"password","params":{"special":"false","length":"20"}}
			]
		}
	}`)

	output, err := info.Generate(faker.Rand, m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// put together a struct to unmarshal the output json into
	type jsonStruct struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	type jsonParent struct {
		JStruct jsonStruct `json:"json"`
	}

	var j jsonParent
	err = json.Unmarshal(output.([]byte), &j)
	if err != nil {
		t.Fatal(err.Error())
	}

	// check that the output values are correct
	if j.JStruct.ID != 1 {
		t.Fatalf("ID is not 1 got: %v", j.JStruct.ID)
	}
	if j.JStruct.FirstName != "Markus" {
		t.Errorf("FirstName is incorrect got: %s", j.JStruct.FirstName)
	}
	if j.JStruct.LastName != "Moen" {
		t.Errorf("LastName is incorrect got: %s", j.JStruct.LastName)
	}
	if j.JStruct.Password != "WWXYVxbjXckoID06qBLA" {
		t.Errorf("Password is incorrect got: %s", j.JStruct.Password)
	}
}

func TestJSONArrayLookupWithSubJSON(t *testing.T) {
	faker := New(11)
	info := GetFuncLookup("json")

	m := NewMapParams()
	m.Add("type", "object")
	m.Add("fields", `{
		"name":"json",
		"function":"json",
		"params":{
			"type":"array",
			"rowcount": 10,
			"fields":[
				{"name":"id","function":"autoincrement"},
				{"name":"first_name","function":"firstname"},
				{"name":"last_name","function":"lastname"},
				{"name":"password","function":"password","params":{"special":"false","length":"20"}},
				{"name":"address","function":"address"}
			]
		}
	}`)

	output, err := info.Generate(faker.Rand, m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	// put together a struct to unmarshal the output json into
	type jsonStruct struct {
		ID        int         `json:"id"`
		FirstName string      `json:"first_name"`
		LastName  string      `json:"last_name"`
		Password  string      `json:"password"`
		Address   AddressInfo `json:"address"`
	}

	type jsonParent struct {
		JStruct []jsonStruct `json:"json"`
	}

	var j jsonParent
	err = json.Unmarshal(output.([]byte), &j)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Check row count
	if len(j.JStruct) != 10 {
		t.Fatalf("Row count is not 10 got: %v", len(j.JStruct))
	}

	// check that the output values are correct
	if j.JStruct[0].ID != 1 {
		t.Fatalf("ID is incorrect should be 1 got: %v", j.JStruct[0].ID)
	}
	if j.JStruct[0].FirstName != "Markus" {
		t.Errorf("FirstName is incorrect got: %s", j.JStruct[0].FirstName)
	}
	if j.JStruct[0].LastName != "Moen" {
		t.Errorf("LastName is incorrect got: %s", j.JStruct[0].LastName)
	}
	if j.JStruct[0].Password != "WWXYVxbjXckoID06qBLA" {
		t.Errorf("Password is incorrect got: %s", j.JStruct[0].Password)
	}
	if j.JStruct[0].Address.City != "San Antonio" {
		t.Errorf("City is incorrect got: %s", j.JStruct[0].Address.City)
	}
}

func TestJSONNoType(t *testing.T) {
	Seed(11)

	_, err := JSON(&JSONOptions{
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
		RowCount: 3,
	})
	if err == nil {
		t.Fatal("Expected error, no type specified")
	}
}

func TestJSONNoFields(t *testing.T) {
	Seed(11)

	_, err := JSON(&JSONOptions{
		Type:     "object",
		RowCount: 3,
	})
	if err == nil {
		t.Fatal("Expected error, no fields set")
	}
}

func TestJSONNoCount(t *testing.T) {
	Seed(11)

	_, err := JSON(&JSONOptions{
		Type: "array",
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "password", Function: "password", Params: MapParams{"special": {"false"}}},
		},
	})
	if err == nil {
		t.Fatal("Expected error, no count set for array type")
	}
}

func TestJSONNoOptions(t *testing.T) {
	Seed(11)

	// if JSONOptions is nil -> get a random JSONOptions
	_, err := JSON(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestJSONRawMessage(t *testing.T) {
	type J struct {
		Field json.RawMessage `json:"field"`
	}

	Seed(100)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONRawMessageWithTag(t *testing.T) {
	type J struct {
		Field json.RawMessage `json:"field" faker:"json"`
	}

	Seed(100)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONNumber(t *testing.T) {
	type J struct {
		Field json.Number `json:"field"`
	}

	Seed(100)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkJSONLookup100(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")

		m := NewMapParams()
		m.Add("type", "array")
		m.Add("rowcount", "100")
		m.Add("fields", `{"name":"id","function":"autoincrement"}`)
		m.Add("fields", `{"name":"first_name","function":"firstname"}`)
		m.Add("fields", `{"name":"last_name","function":"lastname"}`)
		m.Add("fields", `{"name":"password","function":"password"}`)
		m.Add("fields", `{"name":"description","function":"paragraph"}`)
		m.Add("fields", `{"name":"created_at","function":"date"}`)

		_, err := info.Generate(faker.Rand, m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup1000(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := MapParams{
			"type":     {"array"},
			"rowcount": {"1000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup10000(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := MapParams{
			"type":     {"array"},
			"rowcount": {"10000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONLookup100000(b *testing.B) {
	faker := New(0)

	for i := 0; i < b.N; i++ {
		info := GetFuncLookup("json")
		m := MapParams{
			"type":     {"array"},
			"rowcount": {"100000"},
			"fields": {
				`{"name":"id","function":"autoincrement"}`,
				`{"name":"first_name","function":"firstname"}`,
				`{"name":"last_name","function":"lastname"}`,
				`{"name":"password","function":"password"}`,
				`{"name":"description","function":"paragraph"}`,
				`{"name":"created_at","function":"date"}`,
			},
		}
		_, err := info.Generate(faker.Rand, &m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
