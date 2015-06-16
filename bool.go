package gofakeit

// Generate Random Boolean value
func Bool() bool {
	if randIntRange(0, 1) == 1 {
		return true
	}

	return false
}
