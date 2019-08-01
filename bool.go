package gofakeit

// Bool will generate a random boolean value
func Bool() bool {
	return randIntRange(0, 1) == 1
}
