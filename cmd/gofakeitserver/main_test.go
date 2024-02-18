package main

import (
	"fmt"
	"math/rand/v2"
	"net/url"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestList(t *testing.T) {
	var response map[string]gofakeit.Info
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "GET",
		Path:       "/list",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if len(response) == 0 {
		t.Fatalf("Was expecting list length to be more than 0 got %d", len(response))
	}
}

func TestGetAllRequests(t *testing.T) {
	for field, info := range gofakeit.FuncLookups {
		mapData := url.Values{}
		if info.Params != nil && len(info.Params) != 0 {
			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				// If default is empty and has options randomly pick one
				if p.Default == "" && len(p.Options) != 0 {
					mapData.Add(p.Field, p.Options[faker.IntN(len(p.Options))])
					continue
				} else if p.Default != "" {
					// If p.Type is []uint, then we need to convert it to []string
					if strings.HasPrefix(p.Type, "[") {
						// Remove [] from type
						defaultClean := p.Default[1 : len(p.Default)-1]

						// Split on comma
						defaultSplit := strings.Split(defaultClean, ",")

						// Loop through defaultSplit and add to mapData
						for _, s := range defaultSplit {
							mapData.Add(p.Field, s)
						}
					} else {
						mapData.Add(p.Field, p.Default)
					}

					continue
				}

				switch p.Type {
				case "bool":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Bool()))
				case "string":
					mapData.Add(p.Field, gofakeit.Letter())
				case "uint":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Uint16()))
				case "int":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Int16()))
				case "float":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Float32()))
				case "[]string":
					mapData.Add(p.Field, gofakeit.Letter())
					mapData.Add(p.Field, gofakeit.Letter())
					mapData.Add(p.Field, gofakeit.Letter())
					mapData.Add(p.Field, gofakeit.Letter())
				case "[]int":
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Int8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Int8()))
				case "[]uint":
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Uint8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Uint8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Uint8()))
					mapData.Add(p.Field, fmt.Sprintf("%d", gofakeit.Uint8()))
				case "[]float":
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Float32()))
					mapData.Add(p.Field, fmt.Sprintf("%v", gofakeit.Float32()))
				case "[]Field":
					mapData.Add(p.Field, `{"name":"first_name","function":"firstname"}`)
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		var statusCode int
		testRequest(&testRequestStruct{
			Testing:     t,
			Method:      "GET",
			Path:        "/" + field,
			QueryParams: mapData,
			StatusCode:  &statusCode,
		})

		if statusCode != 200 {
			t.Fatalf("Was expecting 200 got %d, function %v, params %+v", statusCode, field, mapData)
		}
	}
}

func TestGetLookupNoParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "GET",
		Path:       "/firstname",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}
}

func TestGetLookupWithParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "GET",
		Path:    "/password",
		QueryParams: url.Values{
			"length": []string{"5"},
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}

	if len(response) != 5 {
		t.Fatalf("Was expecting a string length of 5 got %d", len(response))
	}
}

func TestPostAllRequests(t *testing.T) {
	for field, info := range gofakeit.FuncLookups {
		var mapData map[string][]string
		if info.Params != nil && len(info.Params) != 0 {
			// Make sure mapdata is set
			if mapData == nil {
				mapData = make(map[string][]string)
			}

			// Loop through params and add fields to mapdata
			for _, p := range info.Params {
				// If default is empty and has options randomly pick one
				if p.Default == "" && len(p.Options) != 0 {
					mapData[p.Field] = []string{p.Options[rand.IntN(len(p.Options))]}
					continue
				} else if p.Default != "" {
					// If p.Type is []uint, then we need to convert it to []string
					if strings.HasPrefix(p.Type, "[") {
						// Remove [] from type
						defaultClean := p.Default[1 : len(p.Default)-1]
						mapData[p.Field] = strings.Split(defaultClean, ",")
					} else {
						mapData[p.Field] = []string{p.Default}
					}

					continue
				}

				switch p.Type {
				case "bool":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofakeit.Bool())}
				case "string":
					mapData[p.Field] = []string{gofakeit.Letter()}
				case "uint":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofakeit.Uint16())}
				case "int":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofakeit.Int16())}
				case "float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofakeit.Float32())}
				case "[]string":
					mapData[p.Field] = []string{gofakeit.Letter(), gofakeit.Letter(), gofakeit.Letter(), gofakeit.Letter()}
				case "[]int":
					mapData[p.Field] = []string{fmt.Sprintf("%d", gofakeit.Int8()), fmt.Sprintf("%d", gofakeit.Int8()), fmt.Sprintf("%d", gofakeit.Int8()), fmt.Sprintf("%d", gofakeit.Int8())}
				case "[]uint":
					mapData[p.Field] = []string{fmt.Sprintf("%d", gofakeit.Uint8()), fmt.Sprintf("%d", gofakeit.Uint8()), fmt.Sprintf("%d", gofakeit.Uint8()), fmt.Sprintf("%d", gofakeit.Uint8())}
				case "[]float":
					mapData[p.Field] = []string{fmt.Sprintf("%v", gofakeit.Float32()), fmt.Sprintf("%v", gofakeit.Float32()), fmt.Sprintf("%v", gofakeit.Float32()), fmt.Sprintf("%v", gofakeit.Float32())}
				case "[]Field":
					mapData[p.Field] = []string{`{"name":"first_name","function":"firstname"}`}
				default:
					t.Fatalf("Looking for %s but switch case doesnt have it", p.Type)
				}
			}
		}

		var statusCode int
		testRequest(&testRequestStruct{
			Testing:    t,
			Method:     "POST",
			Path:       "/" + field,
			Body:       mapData,
			StatusCode: &statusCode,
		})

		if statusCode != 200 {
			t.Fatalf("Was expecting 200 got %d", statusCode)
		}
	}
}

func TestPostLookupNoParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "POST",
		Path:       "/firstname",
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}
}

func TestPostLookupWithParams(t *testing.T) {
	var response string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "POST",
		Path:    "/password",
		Body: map[string]any{
			"space":  true,
			"length": "500",
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if response == "" {
		t.Fatalf("Was expecting a string with value got empty")
	}

	if len(response) != 500 {
		t.Fatalf("Was expecting a string length of 500 got %d", len(response))
	}

	if !strings.Contains(response, " ") {
		t.Fatal("Was expecting a space in the password and didnt get it")
	}
}

func TestPostLookupWithParamsArray(t *testing.T) {
	var response []string
	var statusCode int
	testRequest(&testRequestStruct{
		Testing: t,
		Method:  "POST",
		Path:    "/shufflestrings",
		Body: map[string]any{
			"strs": []string{"a", "b", "c", "d", "e", "f"},
		},
		Response:   &response,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if len(response) != 6 {
		t.Fatalf("Was expecting a array length of 6 got %d", len(response))
	}
}
