package gofakeit

type VehicleInfo struct {
	// Vehicle type
	VehicleType string
	// Fuel type
	Fuel string
	// Transmission type
	TransmissionGear string
	// Brand name
	Brand string
	// Vehicle model
	Model string
	// Vehicle model year
	Year int
}

// Vehicle will generate a struct with vehicle information
func Vehicle() *VehicleInfo {
	return &VehicleInfo{
		VehicleType:      VehicleType(),
		Fuel:             FuelType(),
		TransmissionGear: TransmissionGearType(),
		Brand:            CarMaker(),
		Model:            CarModel(),
		Year:             Year(),
	}

}

// VehicleType will generate a random vehicle type string
func VehicleType() string {
	return getRandValue([]string{"vehicle", "vehicle_type"})
}

func FuelType() string {
	return getRandValue([]string{"vehicle", "fuel_type"})
}

func TransmissionGearType() string {
	return getRandValue([]string{"vehicle", "transmission_type"})
}

func CarMaker() string {
	return getRandValue([]string{"vehicle", "maker"})
}

func CarModel() string {
	return getRandValue([]string{"vehicle", "model"})
}