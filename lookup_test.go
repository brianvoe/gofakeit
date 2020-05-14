package gofakeit

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func Example_custom() {
	Seed(11)

	AddFuncLookup("friendname", Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})

	type Foo struct {
		FriendName string `fake:"{friendname}"`
	}

	var f Foo
	Struct(&f)

	fmt.Printf("%s", f.FriendName)
	// Output: bill
}

func Example_custom_with_params() {
	Seed(11)

	AddFuncLookup("jumbleword", Info{
		Category:    "jumbleword",
		Description: "Take a word and jumple it up",
		Example:     "loredlowlh",
		Output:      "string",
		Params: []Param{
			{Field: "word", Type: "int", Description: "Word you want to jumble"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			split := strings.Split(word, "")
			ShuffleStrings(split)
			return strings.Join(split, ""), nil
		},
	})

	type Foo struct {
		JumbleWord string `fake:"{jumbleword:helloworld}"`
	}

	var f Foo
	Struct(&f)

	fmt.Printf("%s", f.JumbleWord)
	// Output: loredlowlh
}

func TestLookupChecking(t *testing.T) {
	Seed(time.Now().UnixNano())

	for field, info := range FuncLookups {
		var mapData map[string][]string
		if info.Params != nil && len(info.Params) != 0 {
			// Make sure mapdata is set
			if mapData == nil {
				mapData = make(map[string][]string)
			}

			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				if p.Default != "" {
					mapData[p.Field] = []string{p.Default}
					continue
				}

				switch p.Type {
				case "bool":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Bool())}
					break
				case "string":
					mapData[p.Field] = []string{Letter()}
					break
				case "uint":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Uint16())}
				case "int":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Int16())}
				case "float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Float32())}
					break
				case "[]string":
					mapData[p.Field] = []string{Letter(), Letter(), Letter(), Letter()}
					break
				case "[]int":
					mapData[p.Field] = []string{fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8())}
					break
				case "[]Field":
					mapData[p.Field] = []string{`{"name":"first_name","function":"firstname"}`}
					break
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		_, err := info.Call(&mapData, &info)
		if err != nil {
			t.Fatalf("%s failed - Err: %s - Data: %v", field, err, mapData)
		}
	}
}

// Make sure all lookups have specific fields
func TestLookupCheckFields(t *testing.T) {
	for field, info := range FuncLookups {
		if info.Display == "" {
			t.Fatalf("%s is missing a display", field)
		}
		if info.Category == "" {
			t.Fatalf("%s is missing a category", field)
		}
		if info.Output == "" {
			t.Fatalf("%s is misssing output", field)
		}

		// Check params
		if info.Params != nil {
			for _, p := range info.Params {
				if p.Field == "" {
					t.Fatalf("Field %s param %s is missing a field", field, p.Field)
				}
				if p.Display == "" {
					t.Fatalf("Field %s param %s is missing a display", field, p.Field)
				}
				if p.Type == "" {
					t.Fatalf("Field %s param %s is missing a type", field, p.Field)
				}
				if p.Description == "" {
					t.Fatalf("Field %s param %s is missing a description", field, p.Field)
				}
			}
		}
	}
}

func TestLookupRemove(t *testing.T) {
	funcName := "friendname"

	AddFuncLookup(funcName, Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return RandomString([]string{"bill", "bob", "sally"}), nil
		},
	})

	info := GetFuncLookup(funcName)
	if info == nil {
		t.Fatal("Could not find lookup")
	}

	RemoveFuncLookup(funcName)

	info = GetFuncLookup(funcName)
	if info != nil {
		t.Fatal("Got info when I shouldn't have")
	}
}
