package gofakeit

import (
	"math/rand"
	"strconv"
	"strings"
)

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	_, checkOk = Data[dataVal[0]]
	if len(dataVal) == 2 && checkOk {
		_, checkOk = Data[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func getRandValue(dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return Data[dataVal[0]][dataVal[1]][rand.Intn(len(Data[dataVal[0]][dataVal[1]]))]
}

// Replace # with numbers
func replaceWithNumbers(str string) string {
	for strings.Count(str, "#") > 0 {
		str = strings.Replace(str, "#", strconv.Itoa(rand.Intn(9)), 1)
	}

	return str
}

// Replace ? with letters
func replaceWithLetters(str string) string {
	for strings.Count(str, "?") > 0 {
		str = strings.Replace(str, "?", randLetter(), 1)
	}

	return str
}

// Generate random letter
func randLetter() string {
	alpha := []byte("abcdefghijklmnopqrstuvwxyz")
	return string(alpha[rand.Intn(len(alpha))])
}

// Generate random integer between min and max
func randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

func randFloatRange(min, max float64) float64 {
	if min == max {
		return min
	}
	return rand.Float64()*(max-min) + min
}
