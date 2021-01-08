package gofakeit

import (
	"fmt"
	"math/rand"
)

// Gamertag will generate a random video game username
func Gamertag() string { return gamertag(globalFaker.Rand) }

// Gamertag will generate a random video game username
func (f *Faker) Gamertag() string { return gamertag(f.Rand) }

func gamertag(r *rand.Rand) string {
	return fmt.Sprintf("%s%s%d", getRandValue(r, []string{"word", "noun"}), getRandValue(r, []string{"word", "verb"}), number(r, 10, 999))
}

func addGameLookup() {
	AddFuncLookup("gamertag", Info{
		Display:     "Gamertag",
		Category:    "game",
		Description: "Random gamertag",
		Example:     "footinterpret63",
		Output:      "string",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return gamertag(r), nil
		},
	})
}
