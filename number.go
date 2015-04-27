package gofakeit

// Generate Random Number Between Min And Max Number
func Number(min int, max int) int {
	return randIntRange(min, max)
}

// Replace # With Random Numerical Values
func Numerify(str string) string {
	return replaceWithNumbers(str)
}
