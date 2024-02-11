package gofakeit

import "math/bits"

// Int64N returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (f *Faker) Int64N(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int64N")
	}
	return int64(f.uint64n(uint64(n)))
}

// Uint64N returns, as a uint64, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
func (f *Faker) Uint64N(n uint64) uint64 {
	if n == 0 {
		panic("invalid argument to Uint64N")
	}
	return f.uint64n(n)
}

// uint64n is the no-bounds-checks version of Uint64N.
func (f *Faker) uint64n(n uint64) uint64 {
	if is32bit && uint64(uint32(n)) == n {
		return uint64(f.uint32n(uint32(n)))
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return f.Uint64() & (n - 1)
	}

	// Suppose we have a uint64 x uniform in the range [0,2⁶⁴)
	// and want to reduce it to the range [0,n) preserving exact uniformity.
	// We can simulate a scaling arbitrary precision x * (n/2⁶⁴) by
	// the high bits of a double-width multiply of x*n, meaning (x*n)/2⁶⁴.
	// Since there are 2⁶⁴ possible inputs x and only n possible outputs,
	// the output is necessarily biased if n does not divide 2⁶⁴.
	// In general (x*n)/2⁶⁴ = k for x*n in [k*2⁶⁴,(k+1)*2⁶⁴).
	// There are either floor(2⁶⁴/n) or ceil(2⁶⁴/n) possible products
	// in that range, depending on k.
	// But suppose we reject the sample and try again when
	// x*n is in [k*2⁶⁴, k*2⁶⁴+(2⁶⁴%n)), meaning rejecting fewer than n possible
	// outcomes out of the 2⁶⁴.
	// Now there are exactly floor(2⁶⁴/n) possible ways to produce
	// each output value k, so we've restored uniformity.
	// To get valid uint64 math, 2⁶⁴ % n = (2⁶⁴ - n) % n = -n % n,
	// so the direct implementation of this algorithm would be:
	//
	//	hi, lo := bits.Mul64(f.Uint64(), n)
	//	thresh := -n % n
	//	for lo < thresh {
	//		hi, lo = bits.Mul64(f.Uint64(), n)
	//	}
	//
	// That still leaves an expensive 64-bit division that we would rather avoid.
	// We know that thresh < n, and n is usually much less than 2⁶⁴, so we can
	// avoid the last four lines unless lo < n.
	//
	// See also:
	// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
	// https://lemire.me/blog/2016/06/30/fast-random-shuffling
	hi, lo := bits.Mul64(f.Uint64(), n)
	if lo < n {
		thresh := -n % n
		for lo < thresh {
			hi, lo = bits.Mul64(f.Uint64(), n)
		}
	}
	return hi
}

// uint32n is an identical computation to uint64n
// but optimized for 32-bit systems.
func (f *Faker) uint32n(n uint32) uint32 {
	if n&(n-1) == 0 { // n is power of two, can mask
		return uint32(f.Uint64()) & (n - 1)
	}
	// On 64-bit systems we still use the uint64 code below because
	// the probability of a random uint64 lo being < a uint32 n is near zero,
	// meaning the unbiasing loop almost never runs.
	// On 32-bit systems, here we need to implement that same logic in 32-bit math,
	// both to preserve the exact output sequence observed on 64-bit machines
	// and to preserve the optimization that the unbiasing loop almost never runs.
	//
	// We want to compute
	// 	hi, lo := bits.Mul64(f.Uint64(), n)
	// In terms of 32-bit halves, this is:
	// 	x1:x0 := f.Uint64()
	// 	0:hi, lo1:lo0 := bits.Mul64(x1:x0, 0:n)
	// Writing out the multiplication in terms of bits.Mul32 allows
	// using direct hardware instructions and avoiding
	// the computations involving these zeros.
	x := f.Uint64()
	lo1a, lo0 := bits.Mul32(uint32(x), n)
	hi, lo1b := bits.Mul32(uint32(x>>32), n)
	lo1, c := bits.Add32(lo1a, lo1b, 0)
	hi += c
	if lo1 == 0 && lo0 < uint32(n) {
		n64 := uint64(n)
		thresh := uint32(-n64 % n64)
		for lo1 == 0 && lo0 < thresh {
			x := f.Uint64()
			lo1a, lo0 = bits.Mul32(uint32(x), n)
			hi, lo1b = bits.Mul32(uint32(x>>32), n)
			lo1, c = bits.Add32(lo1a, lo1b, 0)
			hi += c
		}
	}
	return hi
}

// Int32N returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n <= 0.
func (f *Faker) Int32N(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int32N")
	}
	return int32(f.uint64n(uint64(n)))
}

// UintN returns, as a uint, a non-negative pseudo-random number in the half-open interval [0,n).
// It panics if n == 0.
func (f *Faker) UintN(n uint) uint {
	if n == 0 {
		panic("invalid argument to UintN")
	}
	return uint(f.uint64n(uint64(n)))
}

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func (f *Faker) Shuffle(n int, swap func(i, j int)) {
	if n < 0 {
		panic("invalid argument to Shuffle")
	}

	// Fisher-Yates shuffle: https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
	// Shuffle really ought not be called with n that doesn't fit in 32 bits.
	// Not only will it take a very long time, but with 2³¹! possible permutations,
	// there's no way that any PRNG can have a big enough internal state to
	// generate even a minuscule percentage of the possible permutations.
	// Nevertheless, the right API signature accepts an int n, so handle it as best we can.
	for i := n - 1; i > 0; i-- {
		j := int(f.uint64n(uint64(i + 1)))
		swap(i, j)
	}
}
