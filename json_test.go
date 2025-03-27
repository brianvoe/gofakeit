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
	//     "first_name": "Sonny",
	//     "last_name": "Stiedemann",
	//     "address": {
	//         "address": "52759 Stationside, San Diego, Oregon 99344",
	//         "street": "52759 Stationside",
	//         "city": "San Diego",
	//         "state": "Oregon",
	//         "zip": "99344",
	//         "country": "Saint Pierre and Miquelon",
	//         "latitude": -30.009814,
	//         "longitude": 154.519771
	//     },
	//     "password": "l4338TebFL55"
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
	//     "first_name": "Sonny",
	//     "last_name": "Stiedemann",
	//     "address": {
	//         "address": "52759 Stationside, San Diego, Oregon 99344",
	//         "street": "52759 Stationside",
	//         "city": "San Diego",
	//         "state": "Oregon",
	//         "zip": "99344",
	//         "country": "Saint Pierre and Miquelon",
	//         "latitude": -30.009814,
	//         "longitude": 154.519771
	//     },
	//     "password": "l4338TebFL55"
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
	//         "first_name": "Sonny",
	//         "last_name": "Stiedemann",
	//         "password": "8nwf0o3sBXcR"
	//     },
	//     {
	//         "id": 2,
	//         "first_name": "Verda",
	//         "last_name": "Brakus",
	//         "password": "3beWLpq75Lua"
	//     },
	//     {
	//         "id": 3,
	//         "first_name": "Jules",
	//         "last_name": "Cremin",
	//         "password": "Uu38J14Y8W82"
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
	//         "first_name": "Sonny",
	//         "last_name": "Stiedemann",
	//         "password": "8nwf0o3sBXcR"
	//     },
	//     {
	//         "id": 2,
	//         "first_name": "Verda",
	//         "last_name": "Brakus",
	//         "password": "3beWLpq75Lua"
	//     },
	//     {
	//         "id": 3,
	//         "first_name": "Jules",
	//         "last_name": "Cremin",
	//         "password": "Uu38J14Y8W82"
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

	_, err := info.Generate(faker, m, info)
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

	output, err := info.Generate(faker, m, info)
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
	if j.JStruct.ID == 0 {
		t.Fatalf("ID is not 1 got: %v", j.JStruct.ID)
	}
	if j.JStruct.FirstName == "" {
		t.Errorf("FirstName is empty")
	}
	if j.JStruct.LastName == "" {
		t.Errorf("LastName is empty")
	}
	if j.JStruct.Password == "" {
		t.Errorf("Password is empty")
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

	output, err := info.Generate(faker, m, info)
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
	if len(j.JStruct) == 0 {
		t.Fatalf("Row count is not 10 got: %v", len(j.JStruct))
	}

	// check that the output values are correct
	if j.JStruct[0].ID == 0 {
		t.Fatalf("ID is empty")
	}
	if j.JStruct[0].FirstName == "" {
		t.Errorf("FirstName is empty")
	}
	if j.JStruct[0].LastName == "" {
		t.Errorf("LastName is empty")
	}
	if j.JStruct[0].Password == "" {
		t.Errorf("Password is empty")
	}
	if j.JStruct[0].Address.City == "" {
		t.Errorf("City is empty")
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

	Seed(11)

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

	Seed(11)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONRawMessageWithCustomFuncTag(t *testing.T) {
	AddFuncLookup("customjsontest", Info{
		Display:     "CustomJSONTest",
		Category:    "file",
		Example:     `{"ErTA":"bale","FQJUIGrmnRBfuGlb":"over","HTJJPnEKGS":"please say that again","HvLvfsQRGbK":"whenever one turns around","KKbMlbxquDmwwvRWVlPmwRAeAw":"Voluptatem eaque quia facilis quo."}`,
		Output:      "[]byte",
		ContentType: "application/json",
		Params:      []Param{},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			length := f.Number(5, 15)
			dataSet := []string{"word", "phrase", "loremipsumsentence", "{hackeradjective}-{hackernoun}", "float64", "bool"}
			resultMap := make(map[string]any)
			for i := 0; i < length; i++ {
				key := f.LetterN(uint(f.Number(3, 30)))
				val, err := f.Generate(RandomString(dataSet))
				if err != nil {
					return nil, err
				}
				resultMap[key] = val
			}
			marshal, err := json.Marshal(resultMap)
			if err != nil {
				panic(err)
			}
			return marshal, nil
		},
	})

	type J struct {
		Field json.RawMessage `json:"field" fake:"customjsontest"`
	}

	Seed(11)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONRawMessageWithInvalidCustomFuncTag(t *testing.T) {
	AddFuncLookup("invalidjsontest", Info{
		Display:  "InvalidJSONTest",
		Category: "file",
		Example:  `[181 251 51 164 185 142 21 3 33]`,
		Output:   "[]byte",
		Params:   []Param{},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			var result []byte
			Slice(&result)
			return result, nil
		},
	})

	type J struct {
		Field json.RawMessage `json:"field" fake:"invalidjsontest"`
	}

	Seed(11)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}

	var obj J
	err = Struct(&obj)
	if err == nil {
		t.Fatal("'Struct(&obj)' was supposed to return an error but didn't")
	}
}

func TestJSONNumber(t *testing.T) {
	type J struct {
		Field json.Number `json:"field"`
	}

	Seed(11)

	var objs []J
	Slice(&objs)

	_, err := json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONNumberWithTag(t *testing.T) {
	type J struct {
		Field json.Number `json:"field" fake:"number:3,7"`
	}

	Seed(11)

	var objs []J
	Slice(&objs)

	got, err := objs[0].Field.Int64()
	if err != nil {
		t.Fatal(err)
	}
	if got < 3 || got > 7 {
		t.Errorf("Expected a number between 3 and 7, got %d", got)
	}

	_, err = json.Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}
}

func ExampleJSON_numberWithTag() {
	Seed(11)

	type J struct {
		FieldNumber  json.Number `fake:"number:3,7"`
		FieldInt8    json.Number `fake:"int8"`
		FieldInt16   json.Number `fake:"int16"`
		FieldInt32   json.Number `fake:"int32"`
		FieldInt64   json.Number `fake:"int64"`
		FieldUint8   json.Number `fake:"uint8"`
		FieldUint16  json.Number `fake:"uint16"`
		FieldUint32  json.Number `fake:"uint32"`
		FieldUint64  json.Number `fake:"uint64"`
		FieldFloat32 json.Number `fake:"float32"`
		FieldFloat64 json.Number `fake:"float64range:12,72"`
	}

	var obj J
	Struct(&obj)

	fmt.Printf("obj.FieldNumber = %+v\n", obj.FieldNumber)
	fmt.Printf("obj.FieldInt8 = %+v\n", obj.FieldInt8)
	fmt.Printf("obj.FieldInt16 = %+v\n", obj.FieldInt16)
	fmt.Printf("obj.FieldInt32 = %+v\n", obj.FieldInt32)
	fmt.Printf("obj.FieldInt64 = %+v\n", obj.FieldInt64)
	fmt.Printf("obj.FieldUint8 = %+v\n", obj.FieldUint8)
	fmt.Printf("obj.FieldUint16 = %+v\n", obj.FieldUint16)
	fmt.Printf("obj.FieldUint32 = %+v\n", obj.FieldUint32)
	fmt.Printf("obj.FieldUint64 = %+v\n", obj.FieldUint64)
	fmt.Printf("obj.FieldFloat32 = %+v\n", obj.FieldFloat32)
	fmt.Printf("obj.FieldFloat64 = %+v\n", obj.FieldFloat64)

	// Output: obj.FieldNumber = 7
	// obj.FieldInt8 = -110
	// obj.FieldInt16 = 10933
	// obj.FieldInt32 = 430103905
	// obj.FieldInt64 = 525217394518216243
	// obj.FieldUint8 = 164
	// obj.FieldUint16 = 63417
	// obj.FieldUint32 = 2307233133
	// obj.FieldUint64 = 17560678512042153749
	// obj.FieldFloat32 = 0.11857688426971436
	// obj.FieldFloat64 = 51.03971481390635
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

		_, err := info.Generate(faker, m, info)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
