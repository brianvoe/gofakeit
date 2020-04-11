package main

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v4"
)

func TestList(t *testing.T) {
	var response gofakeit.Lookup
	var statusCode int
	testRequest(&testRequestStruct{
		Testing:    t,
		Method:     "GET",
		Path:       "/list",
		Response:   &response.Map,
		StatusCode: &statusCode,
	})

	if statusCode != 200 {
		t.Fatalf("Was expecting 200 got %d", statusCode)
	}

	if len(response.Map) == 0 {
		t.Fatalf("Was expecting list length to be more than 0 got %d", len(response.Map))
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
		QueryParams: map[string]string{
			"length": "5",
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
		Body: map[string]interface{}{
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
		Body: map[string]interface{}{
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
