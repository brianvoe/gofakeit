package gofakeit

import (
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
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
			name: "initial",
			args: args{
				template: `{"Hello" :"World{internet.browser}","Hel" : {	"Hel" : "W"	}}`,
			},
			want: map[string]interface{}{
				"Hello": "Worldchrome",
				"Hel":   map[string]interface{}{"Hel": "W"},
			},
			wantErr: false,
		},
		// {
		// 	name: "array",
		// 	args: args{
		// 		template: `{"Hello" : [ "A","B","C" ]}`,
		// 	},
		// 	want:    map[string]interface{}{"Hello": []string{"A", "B", "C"}},
		// 	wantErr: false,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSON(tt.args.template)
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

func Test_randJSON(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := randJSON(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("randJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
