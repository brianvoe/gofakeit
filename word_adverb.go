package gofakeit

import rand "math/rand"

// Adverb will generate a random adverb
func Adverb() string { return adverb(globalFaker.Rand) }

// Adverb will generate a random adverb
func (f *Faker) Adverb() string { return adverb(f.Rand) }

func adverb(r *rand.Rand) string { return getRandValue(r, []string{"word", "adverb"}) }

// AdverbManner will generate a random manner adverb
func AdverbManner() string { return adverbManner(globalFaker.Rand) }

// AdverbManner will generate a random manner adverb
func (f *Faker) AdverbManner() string { return adverbManner(f.Rand) }

func adverbManner(r *rand.Rand) string { return getRandValue(r, []string{"word", "adverb_manner"}) }

// AdverbDegree will generate a random degree adverb
func AdverbDegree() string { return adverbDegree(globalFaker.Rand) }

// AdverbDegree will generate a random degree adverb
func (f *Faker) AdverbDegree() string { return adverbDegree(f.Rand) }

func adverbDegree(r *rand.Rand) string { return getRandValue(r, []string{"word", "adverb_degree"}) }

// AdverbPlace will generate a random place adverb
func AdverbPlace() string { return adverbPlace(globalFaker.Rand) }

// AdverbPlace will generate a random place adverb
func (f *Faker) AdverbPlace() string { return adverbPlace(f.Rand) }

func adverbPlace(r *rand.Rand) string { return getRandValue(r, []string{"word", "adverb_place"}) }

// AdverbTimeDefinite will generate a random time definite adverb
func AdverbTimeDefinite() string { return adverbTimeDefinite(globalFaker.Rand) }

// AdverbTimeDefinite will generate a random time definite adverb
func (f *Faker) AdverbTimeDefinite() string { return adverbTimeDefinite(f.Rand) }

func adverbTimeDefinite(r *rand.Rand) string {
	return getRandValue(r, []string{"word", "adverb_time_definite"})
}

// AdverbTimeIndefinite will generate a random time indefinite adverb
func AdverbTimeIndefinite() string { return adverbTimeIndefinite(globalFaker.Rand) }

// AdverbTimeIndefinite will generate a random time indefinite adverb
func (f *Faker) AdverbTimeIndefinite() string { return adverbTimeIndefinite(f.Rand) }

func adverbTimeIndefinite(r *rand.Rand) string {
	return getRandValue(r, []string{"word", "adverb_time_indefinite"})
}

// AdverbFrequencyDefinite will generate a random frequency definite adverb
func AdverbFrequencyDefinite() string { return adverbFrequencyDefinite(globalFaker.Rand) }

// AdverbFrequencyDefinite will generate a random frequency definite adverb
func (f *Faker) AdverbFrequencyDefinite() string { return adverbFrequencyDefinite(f.Rand) }

func adverbFrequencyDefinite(r *rand.Rand) string {
	return getRandValue(r, []string{"word", "adverb_frequency_definite"})
}

// AdverbFrequencyIndefinite will generate a random frequency indefinite adverb
func AdverbFrequencyIndefinite() string { return adverbFrequencyIndefinite(globalFaker.Rand) }

// AdverbFrequencyIndefinite will generate a random frequency indefinite adverb
func (f *Faker) AdverbFrequencyIndefinite() string { return adverbFrequencyIndefinite(f.Rand) }

func adverbFrequencyIndefinite(r *rand.Rand) string {
	return getRandValue(r, []string{"word", "adverb_frequency_indefinite"})
}
