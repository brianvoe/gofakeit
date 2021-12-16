package gofakeit

import rand "math/rand"

// Noun will generate a random noun
func Noun() string { return noun(globalFaker.Rand) }

// Noun will generate a random noun
func (f *Faker) Noun() string { return noun(f.Rand) }

func noun(r *rand.Rand) string { return getRandValue(r, []string{"word", "noun"}) }

// NounCommon will generate a random common noun
func NounCommon() string { return nounCommon(globalFaker.Rand) }

// NounCommon will generate a random common noun
func (f *Faker) NounCommon() string { return nounCommon(f.Rand) }

func nounCommon(r *rand.Rand) string { return getRandValue(r, []string{"word", "noun_common"}) }

// NounProper will generate a random proper noun
func NounProper() string { return nounProper(globalFaker.Rand) }

// NounProper will generate a random proper noun
func (f *Faker) NounProper() string { return nounProper(f.Rand) }

func nounProper(r *rand.Rand) string {
	switch randInt := randIntRange(r, 1, 3); randInt {
	case 1:
		return getRandValue(r, []string{"person", "first"})
	case 2:
		return getRandValue(r, []string{"celebrity", "actor"})
	case 3:
		return generate(r, getRandValue(r, []string{"address", "city"}))
	}

	return getRandValue(r, []string{"person", "first"})
}
