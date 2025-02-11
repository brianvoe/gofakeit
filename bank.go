package gofakeit

func BankName() string { return bankName(GlobalFaker) }

func (f *Faker) BankName() string { return bankName(f) }

func bankName(f *Faker) string { return getRandValue(f, []string{"bank", "name"}) }

func BankType() string { return bankType(GlobalFaker) }

func (f *Faker) BankType() string { return bankType(f) }

func bankType(f *Faker) string { return getRandValue(f, []string{"bank", "type"}) }

type BankInfo struct {
	Name string `json:"name" xml:"name"`
	Type string `json:"type" xml:"type"`
}

func Bank() *BankInfo { return bank(GlobalFaker) }

func (f *Faker) Bank() *BankInfo { return bank(f) }

func bank(f *Faker) *BankInfo {
	return &BankInfo{
		Name: bankName(f),
		Type: bankType(f),
	}
}

func addBankLookup() {
	AddFuncLookup("bank", Info{
		Display:     "Bank",
		Category:    "finance",
		Description: "A financial institution that provides various banking services",
		Example: `{
	"name": "Bank of America",
	"type": "Commercial Bank"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bank(f), nil
		},
	})

	AddFuncLookup("bankname", Info{
		Display:     "Bank Name",
		Category:    "finance",
		Description: "Name of a financial institution that offers banking services",
		Example:     "Wells Fargo",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bankName(f), nil
		},
	})

	AddFuncLookup("banktype", Info{
		Display:     "Bank Type",
		Category:    "finance",
		Description: "Classification of a bank based on its services and operations",
		Example:     "Investment Bank",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return bankType(f), nil
		},
	})
}
