package source

import "math/bits"

// The JSF(Jenkins Small Fast) pseudo-random number generator.
// Developed by Bob Jenkins, JSF is known for its speed and efficiency, making it suitable
// for applications requiring fast, non-cryptographic quality random numbers. This implementation
// offers seamless integration with Go's math/rand package and includes an improved seeding mechanism.
// It is the 64 bit version from https://burtleburtle.net/bob/rand/smallprng.html

// Pros:
// - Fast and efficient, ideal for high-performance requirements.
// - Good randomness quality for non-cryptographic applications.
// - Small state size and simple operations, ensuring a minimal memory footprint.

// Cons:
// - Not suitable for cryptographic purposes due to its non-cryptographic security level.
// - Quality of randomness may not match that of more complex algorithms.

type JSF struct {
	a, b, c, d uint64
}

// NewJSF creates and returns a new JSF pseudo-random number generator.
func NewJSF(seed uint64) *JSF {
	jsf := &JSF{}
	jsf.Seed(seed)
	return jsf
}

// Seed sets the seed of the JSF like raninit does: a is a constant, b, c and d are the seed
// and the first 20 outputs are thrown away, so that neighbouring seeds do not give neighbouring sequences.
func (jsf *JSF) Seed(seed uint64) {
	jsf.a = 0xf1ea5eed
	jsf.b = seed
	jsf.c = seed
	jsf.d = seed
	for range 20 {
		jsf.Uint64()
	}
}

// Uint64 generates a pseudo-random 64-bit value using the 64 bit ranval of Bob Jenkins.
func (jsf *JSF) Uint64() uint64 {
	e := jsf.a - bits.RotateLeft64(jsf.b, 7)
	jsf.a = jsf.b ^ bits.RotateLeft64(jsf.c, 13)
	jsf.b = jsf.c + bits.RotateLeft64(jsf.d, 37)
	jsf.c = jsf.d + e
	jsf.d = e + jsf.a
	return jsf.d
}
