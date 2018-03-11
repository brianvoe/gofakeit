package gofakeit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func ExampleJSONString() {
	Seed(42)
	jsonMap, err := JSONString(`{"Name":"{person.first}"}`)
	if err != nil {
		//handle err
	}
	jsonString, err := json.Marshal(jsonMap)
	fmt.Printf("%s", jsonString)
	// Output: {"Name":"Jeromy"}
}

func ExampleJSON() {
	Seed(42)
	m := make(map[string]interface{})
	// generate random keys as well
	m["{person.first}"] = "{person.last}"
	jsonMap, err := JSON(m)
	if err != nil {
		//handle err
	}
	jsonString, err := json.Marshal(jsonMap)
	fmt.Printf("%s", jsonString)
	// Output: {"Jeromy":"Schmeler"}
}

func TestJSONString(t *testing.T) {
	type args struct {
		template string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "PlainJSON",
			args: args{
				template: `{  
					"Hello":"World",
					"Bob":{  
					   "The":"Builder"
					},
					"Tools":[  
					   "Truck",
					   "Trolley"
					],
					"Bool":true,
					"Nil":null
				 }`,
			},
			want: map[string]interface{}{
				"Hello": "World",
				"Bob": map[string]interface{}{
					"The": "Builder",
				},
				"Tools": []interface{}{
					"Truck", "Trolley",
				},
				"Bool": true,
				"Nil":  nil,
			},
			wantErr: false,
		},
		{
			name: "FaultyJSON1",
			args: args{
				template: `{  
					"Hello":"World",
					`,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			json.Compact(&buf, []byte(tt.args.template))
			got, err := JSONString(buf.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("JSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// err returned, should get nil
			if tt.wantErr && got == nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
