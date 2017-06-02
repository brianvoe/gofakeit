package gofakeit

// Data consists of the main set of fake information
var Data = map[string]map[string][]string{
	"person":   DataPerson,
	"contact":  DataContact,
	"address":  DataAddress,
	"company":  DataCompany,
	"job":      DataJob,
	"lorem":    DataLorem,
	"internet": DataInternet,
	"file":     DataFiles,
	"color":    DataColors,
	"computer": DataComputer,
	"payment":  DataPayment,
	"hipster":  DataHipster,
	"beer":     DataBeer,
	"hacker":   DataHacker,
}
