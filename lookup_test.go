package gofakeit

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	SetLookups()
}

func TestLookupChecking(t *testing.T) {
	Seed(time.Now().UnixNano())

	for field, info := range MapLookups.Map {
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
				case "stringarray":
					mapData[p.Field] = []string{Letter(), Letter(), Letter(), Letter()}
					break
				case "intarray":
					mapData[p.Field] = []string{fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8()), fmt.Sprintf("%d", Int8())}
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

// Make sure all lookups have a category
func TestLookupCategories(t *testing.T) {
	for field, info := range MapLookups.Map {
		if info.Category == "" {
			t.Fatalf("%s is missing a category", field)
		}
	}
}
