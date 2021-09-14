package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

var port string
var faker *gofakeit.Faker

func init() {
	flag.StringVar(&port, "port", "8080", "server port")
	faker = gofakeit.New(0)
}

func main() {
	flag.Parse()

	gofakeit.Seed(0)

	// Set router
	mux := http.NewServeMux()
	routes(mux)

	fmt.Println("Running on port " + port)

	http.ListenAndServe(":"+port, mux)
}

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/favicon.ico", favicon)
	mux.HandleFunc("/list", list)
	mux.HandleFunc("/", lookup)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write(gofakeit.ImagePng(32, 32))
}

func list(w http.ResponseWriter, r *http.Request) {
	ok(w, gofakeit.FuncLookups)
}

func lookup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		lookupGet(w, r)
		return
	case http.MethodPost:
		lookupPost(w, r)
		return
	}

	badrequest(w, "Only Get and Post methods allowed")
}

func lookupGet(w http.ResponseWriter, r *http.Request) {
	// Check url query for field(func) and
	info, err := getInfoFromPath(r)
	if err != nil {
		badrequest(w, err.Error())
		return
	}

	// Get url params
	m := gofakeit.NewMapParams()
	urlParams := r.URL.Query()
	for key, values := range urlParams {

		for _, value := range values {
			m.Add(key, value)
		}
	}

	// Generate requested data
	data, err := info.Generate(faker.Rand, m, info)
	if err != nil {
		badrequest(w, err.Error())
		return
	}

	ok(w, data)
}

func lookupPost(w http.ResponseWriter, r *http.Request) {
	// Expect content type json
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		badrequest(w, `Expects header to be of Content-Type "application/json"`)
		return
	}

	// Check url query for field(func) and
	info, err := getInfoFromPath(r)
	if err != nil {
		badrequest(w, err.Error())
		return
	}

	var mapString *gofakeit.MapParams

	// Try to decode body if params are needed
	if len(info.Params) > 0 {
		mapString = gofakeit.NewMapParams()
		mapInterface := map[string]interface{}{}
		err = json.NewDecoder(r.Body).Decode(&mapInterface)
		if err != nil {
			badrequest(w, "Could not parse post body. Expects key(string) - value or []value")
			return
		}
		defer r.Body.Close()

		for key, value := range mapInterface {
			v := reflect.ValueOf(value)
			switch v.Kind() {
			case reflect.Bool:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.Float32, reflect.Float64:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.String:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.Map:
				mapString.Add(key, fmt.Sprintf("%v", value))
			case reflect.Slice:
				var vals []string
				for i := 0; i < v.Len(); i++ {
					vals = append(vals, fmt.Sprintf("%v", v.Index(i).Interface()))
				}

				for _, val := range vals {
					mapString.Add(key, val)
				}
			}
		}
	}

	// Generate requested data
	data, err := info.Generate(faker.Rand, mapString, info)
	if err != nil {
		badrequest(w, err.Error())
		return
	}

	ok(w, data)
}

func getInfoFromPath(r *http.Request) (*gofakeit.Info, error) {
	paths := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(paths) == 0 {
		return nil, errors.New("No function was called, please pass func parameter")
	}

	// Lookup fake data method
	info := gofakeit.GetFuncLookup(paths[0])
	if info == nil {
		return nil, errors.New("No function was called, please pass func parameter")
	}

	return info, nil
}

func encodeResponse(v interface{}) []byte {
	// Set new bytes buffer
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	enc.Encode(v)
	return buf.Bytes()
}

func ok(w http.ResponseWriter, data interface{}) {
	var resp []byte
	d := reflect.ValueOf(data)
	switch d.Kind() {
	case reflect.Bool:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		resp = []byte(fmt.Sprintf("%v", data))
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		resp = []byte(fmt.Sprintf("%v", data))
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		resp = []byte(fmt.Sprintf("%v", data))
	case reflect.Float32, reflect.Float64:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		resp = []byte(fmt.Sprintf("%v", data))
	case reflect.String:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		resp = []byte((data).(string))
	case reflect.Map:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp = encodeResponse(data)
	case reflect.Slice:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp = encodeResponse(data)
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp = encodeResponse(data)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func badrequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func notfound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}
