package gofakeit

import (
	"fmt"
	rand "math/rand"
	"strings"
	"testing"
)

func Example_custom() {
	Seed(11)

	AddFuncLookup("friendname", Info{
		Category:    "custom",
		Description: "Random friend name",
		Example:     "bill",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
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
	faker := New(0)

	for field, info := range FuncLookups {
		var mapData MapParams
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
				case "string":
					mapData[p.Field] = []string{Letter()}
				case "uint":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Uint16())}
				case "int":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Int16())}
				case "float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Float32())}
				case "[]string":
					mapData[p.Field] = []string{Letter(), Letter(), Letter(), Letter()}
				case "[]int":
					mapData[p.Field] = []string{fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8())}
				case "[]float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", Float32()), fmt.Sprintf("%v", Float32()), fmt.Sprintf("%v", Float32()), fmt.Sprintf("%v", Float32())}
				case "[]Field":
					mapData[p.Field] = []string{`{"name":"first_name","function":"firstname"}`}
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		_, err := info.Generate(faker.Rand, &mapData, &info)
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
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

func TestLookupCalls(t *testing.T) {
	faker := New(0)

	for _, info := range FuncLookups {
		mapData := make(MapParams)

		// If parameters are required build it
		if info.Params != nil && len(info.Params) != 0 {
			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				if p.Default != "" {
					mapData.Add(p.Field, p.Default)
					continue
				}

				switch p.Type {
				case "bool":
					mapData.Add(p.Field, fmt.Sprintf("%v", Bool()))
				case "string":
					mapData.Add(p.Field, Letter())
				case "uint":
					mapData.Add(p.Field, fmt.Sprintf("%v", Uint16()))
				case "int":
					mapData.Add(p.Field, fmt.Sprintf("%v", Int16()))
				case "float":
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
				case "[]string":
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
				case "[]int":
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
				case "[]float":
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
				case "[]Field":
					mapData.Add(p.Field, `{"name":"first_name","function":"firstname"}`)
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		_, err := info.Generate(faker.Rand, &mapData, &info)
		if err != nil {
			t.Fatal(err)
		}

	}
}

func TestLookupCallsErrorParams(t *testing.T) {
	// Look through lookups and empty defaults values to help with tests
	for funcName, info := range FuncLookups {
		if info.Params != nil || len(info.Params) != 0 {
			params := []Param{}
			for _, p := range info.Params {
				p.Default = ""
				params = append(params, p)
			}
			info.Params = params
			AddFuncLookup(funcName, info)
		}
	}

	// Initiate new faker
	faker := New(0)

	for funcName, info := range FuncLookups {
		// If parameters are not required skip it. We are only testing params
		if info.Params == nil || len(info.Params) == 0 {
			continue
		}

		// Loop through each param and mark each one fail
		for i := 0; i < len(info.Params); i++ {
			mapData := NewMapParams()
			skip := false
			currentEmptyParam := ""

			// Loop through params and add fields to mapdata
			for ii, p := range info.Params {
				// If the length is equal to the param loop
				// purposly not add that field to trigger an error
				if i == ii {
					currentEmptyParam = p.Field

					// If param is optional skip it
					if p.Optional {
						skip = true
					}

					continue
				}

				if p.Default != "" {
					mapData.Add(p.Field, p.Default)
					continue
				}

				switch p.Type {
				case "bool":
					mapData.Add(p.Field, fmt.Sprintf("%v", Bool()))
				case "string":
					mapData.Add(p.Field, Letter())
				case "uint":
					mapData.Add(p.Field, fmt.Sprintf("%v", Uint16()))
				case "int":
					mapData.Add(p.Field, fmt.Sprintf("%v", Int16()))
				case "float":
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
				case "[]string":
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
					mapData.Add(p.Field, Letter())
				case "[]int":
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", Int8()))
				case "[]float":
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", Float32()))
				case "[]Field":
					mapData.Add(p.Field, `{"name":"first_name","function":"firstname"}`)
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}

			if !skip {
				_, err := info.Generate(faker.Rand, mapData, &info)
				if err == nil {
					t.Error(funcName+" should have failed on param", currentEmptyParam)
				}
			}
		}
	}

	// Reset lookup functions back
	initLookup()
}
