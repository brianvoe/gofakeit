package gofakeit

import (
	"fmt"
	"net/http"
)

// Global
var routesSlices []string

// ServerRoutes returns all route combinations available for fake data
func ServerRoutes() *http.ServeMux {
	routesSlices = nil
	mux := http.NewServeMux()

	// Person
	mux.HandleFunc(appendToRoutesSlice("/name"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, Name())
	})
	mux.HandleFunc(appendToRoutesSlice("/firstname"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, FirstName())
	})
	mux.HandleFunc(appendToRoutesSlice("/lastname"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, LastName())
	})

	// Address
	mux.HandleFunc(appendToRoutesSlice("/address/full"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, Address().Address)
	})
	mux.HandleFunc(appendToRoutesSlice("/address/street"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, Street())
	})

	// Misc
	mux.HandleFunc(appendToRoutesSlice("/bs"), func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, BS())
	})

	return mux
}

func appendToRoutesSlice(value string) string {
	routesSlices = append(routesSlices, value)
	return value
}
