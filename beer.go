package gofakeit

import "strconv"

// BeerName will return a random beer name
func BeerName() string {
	return getRandValue([]string{"beer", "name"})
}

// BeerStyle will return a random beer style
func BeerStyle() string {
	return getRandValue([]string{"beer", "style"})
}

// BeerHop will return a random beer hop
func BeerHop() string {
	return getRandValue([]string{"beer", "hop"})
}

// BeerYeast will return a random beer yeast
func BeerYeast() string {
	return getRandValue([]string{"beer", "yeast"})
}

// BeerMalt will return a random beer malt
func BeerMalt() string {
	return getRandValue([]string{"beer", "malt"})
}

// BeerAlcohol will return a random beer alcohol level between 2.0 and 10.0
func BeerAlcohol() string {
	return strconv.FormatFloat(randFloat64Range(2.0, 10.0), 'f', 1, 64) + "%"
}

// BeerIbu will return a random beer ibu value between 10 and 100
func BeerIbu() string {
	return strconv.Itoa(randIntRange(10, 100)) + " IBU"
}

// BeerBlg will return a random beer blg between 5.0 and 20.0
func BeerBlg() string {
	return strconv.FormatFloat(randFloat64Range(5.0, 20.0), 'f', 1, 64) + "°Blg"
}

func addBeerLookup() {
	AddFuncLookup("beername", Info{
		Display:     "Beer Name",
		Category:    "beer",
		Description: "Random beer name",
		Example:     "Duvel",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerName(), nil
		},
	})

	AddFuncLookup("beerstyle", Info{
		Display:     "Beer Style",
		Category:    "beer",
		Description: "Random beer style",
		Example:     "European Amber Lager",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerStyle(), nil
		},
	})

	AddFuncLookup("beerhop", Info{
		Display:     "Beer Hop",
		Category:    "beer",
		Description: "Random beer hop type",
		Example:     "Glacier",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerHop(), nil
		},
	})

	AddFuncLookup("beeryeast", Info{
		Display:     "Beer Yeast",
		Category:    "beer",
		Description: "Random beer yeast value",
		Example:     "1388 - Belgian Strong Ale",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerYeast(), nil
		},
	})

	AddFuncLookup("beermalt", Info{
		Display:     "Beer Malt",
		Category:    "beer",
		Description: "Random beer malt",
		Example:     "Munich",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerMalt(), nil
		},
	})

	AddFuncLookup("beeralcohol", Info{
		Display:     "Beer Alcohol",
		Category:    "beer",
		Description: "Random alcohol percentage",
		Example:     "2.7%",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerAlcohol(), nil
		},
	})

	AddFuncLookup("beeribu", Info{
		Display:     "Beer IBU",
		Category:    "beer",
		Description: "Random beer ibu",
		Example:     "29 IBU",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerIbu(), nil
		},
	})

	AddFuncLookup("beerblg", Info{
		Display:     "Beer BLG",
		Category:    "beer",
		Description: "Random beer blg",
		Example:     "6.4°Blg",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return BeerBlg(), nil
		},
	})
}
