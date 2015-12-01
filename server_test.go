package gofakeit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// View output for full list of possible endpoints
func ExampleServerRoutes() {
	ServerRoutes()

	for _, value := range routesSlices {
		fmt.Println(value)
	}

	// Output:
	// /name
	// /firstname
	// /lastname
	// /address/full
	// /address/street
	// /bs
}

// Take routes and use them to create your own http server
func ExampleServerRoutes_second() {
	// Grab routes from gofakeit.ServerRoutes()
	mux := ServerRoutes()

	// Pass the routes to a server of your choice
	// for example purposes im going to do a simple
	// httptest and pass the routes to it.
	Seed(11)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	// Going to hit firstname endpoint
	res, err := http.Get(testServer.URL + "/firstname")
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
	// Output: Markus
}

var routeValues = map[string]string{
	// Person
	"/name":      "Markus Moen",
	"/firstname": "Markus",
	"/lastname":  "Daniel",
	// Address
	"/address/full":   "872 East Rapids borough, Rutherfordstad, New Jersey 74853-6757",
	"/address/street": "872 East Rapids borough",
	// Misc
	"/bs": "front-end",
}

func TestServerRoutes(t *testing.T) {
	testServer := httptest.NewServer(ServerRoutes())

	// Loop through endpoint and make sure values are right
	for endpoint, value := range routeValues {
		Seed(11) // Set seed to ensure expected value
		err := getRequest(testServer, endpoint, value)
		if err != nil {
			t.Error(err)
		}
	}
}

func getRequest(testServer *httptest.Server, endpoint string, value string) error {
	res, err := http.Get(testServer.URL + endpoint)
	if err != nil {
		return err
	}
	message, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	// Make sure value is equal from return response
	if string(message) != value {
		return errors.New(endpoint + " did not return expected value " + value + " got " + string(message))
	}

	return nil
}
