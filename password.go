package gofakeit

import (
	"math/rand"
)

// Password will generate a random password
// Password does not ensure all parameters will be included in final password, if needed use passwordBuilder
func Password(lower bool, upper bool, numeric bool, special bool, space bool, length int) string {
	var passString string
	lowerStr := "abcdefghijklmnopqrstuvwxyz"
	upperStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numericStr := "0123456789"
	specialStr := "!@#$%&?-_"
	spaceStr := " "

	if lower {
		passString += lowerStr
	}
	if upper {
		passString += upperStr
	}
	if numeric {
		passString += numericStr
	}
	if special {
		passString += specialStr
	}
	if space {
		passString += spaceStr
	}

	// Set default if empty
	if passString == "" {
		passString = lowerStr + numericStr
	}

	passBytes := []byte(passString)
	finalBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		finalBytes[i] = passBytes[rand.Intn(len(passBytes))]
	}
	return string(finalBytes)
}

var src = rand.NewSource(11)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
const letterIdxBits = 6                    // 6 bits to represent a letter index
const letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
const letterIdxMax = 63 / letterIdxBits    // # of letter indices fitting in 63 bits

func PasswordBulider(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
