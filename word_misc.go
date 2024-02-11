package gofakeit

// Interjection will generate a random word expressing emotion
func Interjection() string { return interjection(GlobalFaker) }

// Interjection will generate a random word expressing emotion
func (f *Faker) Interjection() string { return interjection(f) }

func interjection(f *Faker) string { return getRandValue(f, []string{"word", "interjection"}) }

func addWordMiscLookup() {
	AddFuncLookup("interjection", Info{
		Display:     "Interjection",
		Category:    "word",
		Description: "Word expressing emotion",
		Example:     "wow",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return interjection(f), nil
		},
	})
}
