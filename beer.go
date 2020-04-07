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

// BeerIbu will return a random beer ibu value between 10 and 100
func BeerIbu() string {
	return strconv.Itoa(randIntRange(10, 100)) + " IBU"
}

// BeerAlcohol will return a random beer alcohol level between 2.0 and 10.0
func BeerAlcohol() string {
	return strconv.FormatFloat(randFloat64Range(2.0, 10.0), 'f', 1, 64) + "%"
}

// BeerBlg will return a random beer blg between 5.0 and 20.0
func BeerBlg() string {
	return strconv.FormatFloat(randFloat64Range(5.0, 20.0), 'f', 1, 64) + "°Blg"
}

func addBeerLookup() {
	AddLookupData("beer.alcohol", Info{
		Description: "Random alcohol percentage",
		Example:     "2.7%",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerAlcohol(), nil
		},
	})

	AddLookupData("beer.blg", Info{
		Description: "Random beer blg",
		Example:     "6.4°Blg",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerBlg(), nil
		},
	})

	AddLookupData("beer.hop", Info{
		Description: "Random beer hop type",
		Example:     "Glacier",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerHop(), nil
		},
	})

	AddLookupData("beer.ibu", Info{
		Description: "Random beer ibu",
		Example:     "29 IBU",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerIbu(), nil
		},
	})

	AddLookupData("beer.malt", Info{
		Description: "Random beer malt",
		Example:     "Munich",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerMalt(), nil
		},
	})

	AddLookupData("beer.name", Info{
		Description: "Random beer name",
		Example:     "Duvel",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerName(), nil
		},
	})

	AddLookupData("beer.style", Info{
		Description: "Random beer style",
		Example:     "European Amber Lager",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerStyle(), nil
		},
	})

	AddLookupData("beer.yeast", Info{
		Description: "Random beer yeast value",
		Example:     "1388 - Belgian Strong Ale",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return BeerYeast(), nil
		},
	})
}
