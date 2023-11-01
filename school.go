package gofakeit

import "math/rand"

// SchoolGen will generate a random School type
func SchoolGen() string { return schoolGen(globalFaker.Rand, globalFaker.Rand, globalFaker.Rand) }

func (f *Faker) SchoolGen() string { return schoolGen(f.Rand, f.Rand, f.Rand) }

func schoolGen(r *rand.Rand, r1 *rand.Rand, r2 *rand.Rand) string {
	return getRandValue(r, []string{"school", "name"}) + " " + getRandValue(r1, []string{"school", "isPrivate"}) + " " + getRandValue(r2, []string{"school", "type"})
}
