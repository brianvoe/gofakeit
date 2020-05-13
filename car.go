package gofakeit

// CarInfo is a struct dataset of all car information
type CarInfo struct {
	Type         string `json:"type" xml:"type"`
	Fuel         string `json:"fuel" xml:"fuel"`
	Transmission string `json:"transmission" xml:"transmission"`
	Brand        string `json:"brand" xml:"brand"`
	Model        string `json:"model" xml:"model"`
	Year         int    `json:"year" xml:"year"`
}

// Car will generate a struct with car information
func Car() *CarInfo {
	return &CarInfo{
		Type:         CarType(),
		Fuel:         CarFuelType(),
		Transmission: CarTransmissionType(),
		Brand:        CarMaker(),
		Model:        CarModel(),
		Year:         Year(),
	}

}

// CarType will generate a random car type string
func CarType() string {
	return getRandValue([]string{"car", "type"})
}

// CarFuelType will return a random fuel type
func CarFuelType() string {
	return getRandValue([]string{"car", "fuel_type"})
}

// CarTransmissionType will return a random transmission type
func CarTransmissionType() string {
	return getRandValue([]string{"car", "transmission_type"})
}

// CarMaker will return a random car maker
func CarMaker() string {
	return getRandValue([]string{"car", "maker"})
}

// CarModel will return a random car model
func CarModel() string {
	return getRandValue([]string{"car", "model"})
}

func addCarLookup() {
	AddFuncLookup("car", Info{
		Display:     "Car",
		Category:    "car",
		Description: "Random car set of data",
		Output:      "map[string]interface",
		Example:     `{type: "Passenger car mini", fuel: "Gasoline", transmission: "Automatic", brand: "Fiat", model: "Freestyle Fwd", year: "1972"}`,
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Car(), nil
		},
	})

	AddFuncLookup("cartype", Info{
		Display:     "Car Type",
		Category:    "car",
		Description: "Random car type",
		Example:     "Passenger car mini",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CarType(), nil
		},
	})

	AddFuncLookup("carfueltype", Info{
		Display:     "Car Fuel Type",
		Category:    "car",
		Description: "Random car fuel type",
		Example:     "CNG",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CarFuelType(), nil
		},
	})

	AddFuncLookup("cartransmissiontype", Info{
		Display:     "Car Transmission Type",
		Category:    "car",
		Description: "Random car transmission type",
		Example:     "Manual",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CarTransmissionType(), nil
		},
	})

	AddFuncLookup("carmaker", Info{
		Display:     "Car Maker",
		Category:    "car",
		Description: "Random car maker",
		Example:     "Nissan",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CarMaker(), nil
		},
	})

	AddFuncLookup("carmodel", Info{
		Display:     "Car Model",
		Category:    "car",
		Description: "Random car model",
		Example:     "Aveo",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return CarModel(), nil
		},
	})
}
