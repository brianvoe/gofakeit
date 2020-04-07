package gofakeit

// VehicleInfo is a struct dataset of all vehicle information
type VehicleInfo struct {
	Type         string `json:"type"`
	Fuel         string `json:"fuel"`
	Transmission string `json:"transmission"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
}

// Vehicle will generate a struct with vehicle information
func Vehicle() *VehicleInfo {
	return &VehicleInfo{
		Type:         VehicleType(),
		Fuel:         FuelType(),
		Transmission: TransmissionGearType(),
		Brand:        CarMaker(),
		Model:        CarModel(),
		Year:         Year(),
	}

}

// VehicleType will generate a random vehicle type string
func VehicleType() string {
	return getRandValue([]string{"vehicle", "vehicle_type"})
}

// FuelType will return a random fuel type
func FuelType() string {
	return getRandValue([]string{"vehicle", "fuel_type"})
}

// TransmissionGearType will return a random transmission gear type
func TransmissionGearType() string {
	return getRandValue([]string{"vehicle", "transmission_type"})
}

// CarMaker will return a random car maker
func CarMaker() string {
	return getRandValue([]string{"vehicle", "maker"})
}

// CarModel will return a random car model
func CarModel() string {
	return getRandValue([]string{"vehicle", "model"})
}

func addVehicleLookup() {
	AddLookupData("vehicle", Info{
		Description: "Random vehicle set of data",
		Example:     `{type: "Passenger car mini", fuel: "Gasoline", transmission: "Automatic", brand: "Fiat", model: "Freestyle Fwd", year: "1972"}`,
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return CarMaker(), nil
		},
	})

	AddLookupData("vehicle.maker", Info{
		Description: "Random vehicle maker",
		Example:     "Nissan",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return CarMaker(), nil
		},
	})

	AddLookupData("vehicle.model", Info{
		Description: "Random vehicle model",
		Example:     "Aveo",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return CarModel(), nil
		},
	})

	AddLookupData("vehicle.type", Info{
		Description: "Random vehicle type",
		Example:     "Passenger car mini",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return VehicleType(), nil
		},
	})

	AddLookupData("vehicle.fuel", Info{
		Description: "Random vehicle fuel type",
		Example:     "CNG",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return FuelType(), nil
		},
	})

	AddLookupData("vehicle.transmission", Info{
		Description: "Random vehicle transmission type",
		Example:     "Manual",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return TransmissionGearType(), nil
		},
	})
}
