package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

var ts *httptest.Server

func init() {
	// Set random seed for app
	gofakeit.Seed(0)

	// Set up test server
	mux := http.NewServeMux()
	routes(mux)
	ts = httptest.NewServer(mux)
}

type testRequestStruct struct {
	Testing     *testing.T
	Method      string
	Path        string
	ContentType string
	QueryParams url.Values
	Headers     map[string]string
	Body        any

	Response   any
	StatusCode *int
}

func testRequest(tr *testRequestStruct) {
	// Make sure testing is set
	if tr.Testing == nil {
		panic("Must set testing")
	}

	if tr.ContentType == "" {
		tr.ContentType = "application/json"
	}

	// Check body and convert to io.Reader if body is passed
	var body io.Reader
	if tr.Body != nil {
		if tr.ContentType == "application/json" {
			typeOf := reflect.TypeOf(tr.Body).String()
			if strings.Contains(typeOf, "map") {
				jsonValue, err := json.Marshal(tr.Body)
				if err != nil {
					tr.Testing.Fatal("Unable to marshal Body")
					return
				}
				body = bytes.NewBuffer(jsonValue)
			} else if typeOf == "string" {
				body = bytes.NewBuffer([]byte(tr.Body.(string)))
			} else if typeOf == "[]uint8" {
				body = bytes.NewBuffer(tr.Body.([]byte))
			} else {
				jsonValue, err := json.Marshal(tr.Body)
				if err != nil {
					tr.Testing.Fatal("Unable to marshal Body")
					return
				}
				body = bytes.NewBuffer(jsonValue)
			}
		} else {
			body = bytes.NewBuffer([]byte(tr.Body.(string)))
		}
	}

	// Put together request
	req, err := http.NewRequest(tr.Method, ts.URL+"/"+strings.TrimLeft(tr.Path, "/"), body)
	if err != nil {
		tr.Testing.Fatal(err)
	}

	// Set Content Type
	req.Header.Set("Content-Type", tr.ContentType)

	// Add headers
	for key, value := range tr.Headers {
		req.Header.Set(key, value)
	}

	// Add query parameters
	if tr.QueryParams != nil {
		req.URL.RawQuery = tr.QueryParams.Encode()
	}

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		tr.Testing.Fatal(err)
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")

	// Read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tr.Testing.Fatal(err)
	}

	// Get http status code
	if tr.StatusCode != nil {
		*tr.StatusCode = resp.StatusCode
	}

	// Unmarshal to response
	if tr.Response != nil {
		// Check content type
		switch contentType {
		case "application/json; charset=utf-8":
			if err := json.Unmarshal(respBody, tr.Response); err != nil {
				if ute, ok := err.(*json.UnmarshalTypeError); ok {
					tr.Testing.Fatalf("UnmarshalTypeError %v: %v - %v - %v\n", ute.Field, ute.Value, ute.Type, ute.Offset)
				} else {
					tr.Testing.Fatal("Unmarshal error: ", err)
				}
			}
		case "text/plain; charset=utf-8":
			*tr.Response.(*string) = string(respBody)
		}
	}
}
