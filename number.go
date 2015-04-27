package gofakeit

// Generate Random Number Between Min And Max Number
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Replace # With Random Numerical Values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}

// Return Random Int From Slice of Ints
func SliceInt(slice []int) int {
	return slice[randIntRange(0, len(slice))]
}
