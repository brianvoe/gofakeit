package gofakeit

import (
	"fmt"
	"math/rand"
	"strings"
)

// Gamertag will generate a random video game username
func Gamertag() string { return gamertag(globalFaker.Rand) }

// Gamertag will generate a random video game username
func (f *Faker) Gamertag() string { return gamertag(f.Rand) }

func gamertag(r *rand.Rand) string {
	str := ""
	num := number(r, 1, 4)
	switch num {
	case 1:
		str = fmt.Sprintf("%s%ser", strings.Title(nounConcrete(r)), strings.Title(verbAction(r)))
	case 2:
		str = fmt.Sprintf("%s%s", strings.Title(adjectiveDescriptive(r)), strings.Title(animal(r)))
	case 3:
		str = fmt.Sprintf("%s%s", strings.Title(adjectiveDescriptive(r)), strings.Title(nounConcrete(r)))
	case 4:
		str = fmt.Sprintf("%s%s", strings.Title(fruit(r)), strings.Title(adjectiveDescriptive(r)))
	}

	// Randomly determine if we should add a number
	if r.Intn(3) == 1 {
		str += digitN(r, uint(number(r, 1, 3)))
	}

	// Remove any spaces
	str = strings.Replace(str, " ", "", -1)

	return str
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
