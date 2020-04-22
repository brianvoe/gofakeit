package gofakeit

import "fmt"

// Gamertag will generate a random video game username
func Gamertag() string {
	return getRandValue([]string{"word", "noun"}) + getRandValue([]string{"word", "verb"}) + fmt.Sprintf("%d", Number(10, 999))
}

func addGameLookup() {
	AddFuncLookup("gamertag", Info{
		Display:     "Gamertag",
		Category:    "game",
		Description: "Random gamertag",
		Example:     "footinterpret63",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Gamertag(), nil
		},
	})
}
