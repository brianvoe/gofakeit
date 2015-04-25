package gofakeit

import (
	crand "crypto/rand"
	"math/rand"
	"strconv"
	"strings"
)

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool = false

	_, checkOk = Data[dataVal[0]]
	if len(dataVal) == 2 && checkOk {
		_, checkOk = Data[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func getRandValue(dataVal []string) string {
	return Data[dataVal[0]][dataVal[1]][randIntRange(0, len(Data[dataVal[0]][dataVal[1]]))]
}

// Replace # with numbers
func replaceWithNumbers(str string) string {
	for strings.Count(str, "#") > 0 {
		str = strings.Replace(str, "#", strconv.Itoa(randIntRange(0, 9)), 1)
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

// Generate random integer between min and max
func randIntRange(min, max int) int {
	if min == max {
		return min
	}
	return min + rand.Intn(max-min)
}

// Generate random letter
func randLetter() string {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, 1)
	crand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alpha[b%byte(len(alpha))]
	}
	return string(bytes)
}
