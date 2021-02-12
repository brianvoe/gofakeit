package gofakeit

import (
	"errors"
	rand "math/rand"
)

// Weighted will take in an array of options and weights and return a random selection based upon its indexed weight
func Weighted(options []interface{}, weights []float32) (interface{}, error) {
	return weighted(globalFaker.Rand, options, weights)
}

// Weighted will take in an array of options and weights and return a random selection based upon its indexed weight
func (f *Faker) Weighted(options []interface{}, weights []float32) (interface{}, error) {
	return weighted(f.Rand, options, weights)
}

// Weighted will take in an array of options and weights and return a random selection based upon its indexed weight
func weighted(r *rand.Rand, options []interface{}, weights []float32) (interface{}, error) {
	ol := len(options)
	wl := len(weights)

	// If options length is 1 just return it back
	if ol == 1 {
		return options[0], nil
	}

	// Make sure they are passing in options
	if ol == 0 {
		return nil, errors.New("Didnt pass options")
	}

	// Make sure they are passing in weights
	if wl == 0 {
		return nil, errors.New("Didnt pass weights")
	}

	// Make sure they are passing in the same length
	if ol != wl {
		return nil, errors.New("Options and weights need to be the same length")
	}

	// Compute the discrete cumulative density from the sum of the weights
	cdf := make([]float32, wl)
	var sumOfWeights float32 = 0.0
	for i, weight := range weights {
		if i > 0 {
			cdf[i] = cdf[i-1] + weight
			sumOfWeights += weight
			continue
		}

		cdf[i] = weight
		sumOfWeights += weight
	}

	// Get rand value from a multple of sumOfWeights
	randSumOfWeights := r.Float32() * sumOfWeights

	var l int = 0
	var h int = wl - 1
	for l <= h {
		m := l + (h-l)/2
		if randSumOfWeights <= cdf[m] {
			if m == 0 || (m > 0 && randSumOfWeights > cdf[m-1]) {
				return options[m], nil
			}
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return nil, errors.New("End of function")
}

func addWeightedLookup() {
	AddFuncLookup("weighted", Info{
		Display:     "Weighted",
		Category:    "misc",
		Description: "Randomly select ",
		Example:     "g",
		Output:      "any",
		Params: []Param{
			{Field: "options", Display: "Count", Type: "uint", Description: "Number of digits to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			options, err := info.GetStringArray(m, "options")
			if err != nil {
				return nil, err
			}

			weights, err := info.GetFloat32Array(m, "weights")
			if err != nil {
				return nil, err
			}

			return weighted(r, options, weights)
		},
	})
}
