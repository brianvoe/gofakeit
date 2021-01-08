package gofakeit

import (
	crand "crypto/rand"
	"encoding/binary"
	"math"
	"math/rand"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit/v6/data"
)

const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const upperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numericStr = "0123456789"
const specialStr = "!@#$%&*+-_=?:;,.|(){}<>"
const spaceStr = " "
const allStr = lowerStr + upperStr + numericStr + specialStr + spaceStr
const hashtag = '#'
const questionmark = '?'
const base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Seed will set the global random value. Setting seed to 0 will use crypto/rand
func Seed(seed int64) {
	if seed == 0 {
		binary.Read(crand.Reader, binary.BigEndian, &seed)
		globalFaker.Rand.Seed(seed)
	} else {
		globalFaker.Rand.Seed(seed)
	}
}

// Check if in lib
func dataCheck(dataVal []string) bool {
	var checkOk bool

	if len(dataVal) == 2 {
		_, checkOk = data.Data[dataVal[0]]
		if checkOk {
			_, checkOk = data.Data[dataVal[0]][dataVal[1]]
		}
	}

	return checkOk
}

// Check if in lib
func intDataCheck(dataVal []string) bool {
	if len(dataVal) != 2 {
		return false
	}

	_, checkOk := data.IntData[dataVal[0]]
	if checkOk {
		_, checkOk = data.IntData[dataVal[0]][dataVal[1]]
	}

	return checkOk
}

// Get Random Value
func getRandValue(r *rand.Rand, dataVal []string) string {
	if !dataCheck(dataVal) {
		return ""
	}
	return data.Data[dataVal[0]][dataVal[1]][r.Intn(len(data.Data[dataVal[0]][dataVal[1]]))]
}

// Get Random Integer Value
func getRandIntValue(r *rand.Rand, dataVal []string) int {
	if !intDataCheck(dataVal) {
		return 0
	}
	return data.IntData[dataVal[0]][dataVal[1]][r.Intn(len(data.IntData[dataVal[0]][dataVal[1]]))]
}

// Replace # with numbers
func replaceWithNumbers(r *rand.Rand, str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(randDigit(r))
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(r.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

// Replace ? with ASCII lowercase letters
func replaceWithLetters(r *rand.Rand, str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(randLetter(r))
		}
	}

	return string(bytestr)
}

// Replace ? with ASCII lowercase letters between a and f
func replaceWithHexLetters(r *rand.Rand, str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(randHexLetter(r))
		}
	}

	return string(bytestr)
}

// Generate random lowercase ASCII letter
func randLetter(r *rand.Rand) rune {
	allLetters := upperStr + lowerStr
	return rune(allLetters[r.Intn(len(allLetters))])
}

func randCharacter(r *rand.Rand, s string) string {
	return string(s[r.Int63()%int64(len(s))])
}

// Generate random lowercase ASCII letter between a and f
func randHexLetter(r *rand.Rand) rune {
	return rune(byte(r.Intn(6)) + 'a')
}

// Generate random ASCII digit
func randDigit(r *rand.Rand) rune {
	return rune(byte(r.Intn(10)) + '0')
}

// Generate random integer between min and max
func randIntRange(r *rand.Rand, min, max int) int {
	if min == max {
		return min
	}
	return r.Intn((max+1)-min) + min
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Floor(num*output)) / output
}

func equalSliceString(a, b []string) bool {
	sizeA, sizeB := len(a), len(b)
	if sizeA != sizeB {
		return false
	}

	for i, va := range a {
		vb := b[i]

		if va != vb {
			return false
		}
	}
	return true
}

func equalSliceInt(a, b []int) bool {
	sizeA, sizeB := len(a), len(b)
	if sizeA != sizeB {
		return false
	}

	for i, va := range a {
		vb := b[i]

		if va != vb {
			return false
		}
	}
	return true
}

func equalSliceInterface(a, b []interface{}) bool {
	sizeA, sizeB := len(a), len(b)
	if sizeA != sizeB {
		return false
	}

	for i, va := range a {
		if !reflect.DeepEqual(va, b[i]) {
			return false
		}
	}
	return true
}

func funcLookupSplit(str string) []string {
	out := []string{}
	for str != "" {
		if strings.HasPrefix(str, "[") {
			startIndex := strings.Index(str, "[")
			endIndex := strings.Index(str, "]")
			val := str[(startIndex) : endIndex+1]
			out = append(out, strings.TrimSpace(val))
			str = strings.Replace(str, val, "", 1)

			// Trim off comma if it has it
			if strings.HasPrefix(str, ",") {
				str = strings.Replace(str, ",", "", 1)
			}
		} else {
			strSplit := strings.SplitN(str, ",", 2)
			strSplitLen := len(strSplit)
			if strSplitLen >= 1 {
				out = append(out, strings.TrimSpace(strSplit[0]))
			}
			if strSplitLen >= 2 {
				str = strSplit[1]
			} else {
				str = ""
			}
		}
	}

	return out
}
