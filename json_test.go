package gofakeit

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
