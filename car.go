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
func Car() *CarInfo { return car(GlobalFaker) }

// Car will generate a struct with car information
func (f *Faker) Car() *CarInfo { return car(f) }

func car(f *Faker) *CarInfo {
	return &CarInfo{
		Type:         carType(f),
		Fuel:         carFuelType(f),
		Transmission: carTransmissionType(f),
		Brand:        carMaker(f),
		Model:        carModel(f),
		Year:         year(f),
	}
}

// CarType will generate a random car type string
func CarType() string { return carType(GlobalFaker) }

// CarType will generate a random car type string
func (f *Faker) CarType() string { return carType(f) }

func carType(f *Faker) string { return getRandValue(f, []string{"car", "type"}) }

// CarFuelType will return a random fuel type
func CarFuelType() string { return carFuelType(GlobalFaker) }

// CarFuelType will return a random fuel type
func (f *Faker) CarFuelType() string { return carFuelType(f) }

func carFuelType(f *Faker) string { return getRandValue(f, []string{"car", "fuel_type"}) }

// CarTransmissionType will return a random transmission type
func CarTransmissionType() string { return carTransmissionType(GlobalFaker) }

// CarTransmissionType will return a random transmission type
func (f *Faker) CarTransmissionType() string { return carTransmissionType(f) }

func carTransmissionType(f *Faker) string {
	return getRandValue(f, []string{"car", "transmission_type"})
}

// CarMaker will return a random car maker
func CarMaker() string { return carMaker(GlobalFaker) }

// CarMaker will return a random car maker
func (f *Faker) CarMaker() string { return carMaker(f) }

func carMaker(f *Faker) string { return getRandValue(f, []string{"car", "maker"}) }

// CarModel will return a random car model
func CarModel() string { return carModel(GlobalFaker) }

// CarModel will return a random car model
func (f *Faker) CarModel() string { return carModel(f) }

func carModel(f *Faker) string { return getRandValue(f, []string{"car", "model"}) }

func addCarLookup() {
	AddFuncLookup("car", Info{
		Display:     "Car",
		Category:    "car",
		Description: "Wheeled motor vehicle used for transportation",
		Example: `{
	"type": "Passenger car mini",
	"fuel": "Gasoline",
	"transmission": "Automatic",
	"brand": "Fiat",
	"model": "Freestyle Fwd",
	"year": 1991
}`,
		Output:      "map[string]any",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return car(f), nil
		},
	})

	AddFuncLookup("cartype", Info{
		Display:     "Car Type",
		Category:    "car",
		Description: "Classification of cars based on size, use, or body style",
		Example:     "Passenger car mini",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return carType(f), nil
		},
	})

	AddFuncLookup("carfueltype", Info{
		Display:     "Car Fuel Type",
		Category:    "car",
		Description: "Type of energy source a car uses",
		Example:     "CNG",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return carFuelType(f), nil
		},
	})

	AddFuncLookup("cartransmissiontype", Info{
		Display:     "Car Transmission Type",
		Category:    "car",
		Description: "Mechanism a car uses to transmit power from the engine to the wheels",
		Example:     "Manual",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return carTransmissionType(f), nil
		},
	})

	AddFuncLookup("carmaker", Info{
		Display:     "Car Maker",
		Category:    "car",
		Description: "Company or brand that manufactures and designs cars",
		Example:     "Nissan",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return carMaker(f), nil
		},
	})

	AddFuncLookup("carmodel", Info{
		Display:     "Car Model",
		Category:    "car",
		Description: "Specific design or version of a car produced by a manufacturer",
		Example:     "Aveo",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return carModel(f), nil
		},
	})
}
