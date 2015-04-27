package gofakeit

// Generate Random Lower Case Letter
func Letter() string {
	return randLetter()
}

// Generate Random Letters Replacing ? With Letters
func Lexify(str string) string {
	return replaceWithLetters(str)
}

// Return Random String From Slice of Strings
func SliceString(slice []string) string {
	return slice[randIntRange(0, len(slice))]
}
