package gofakeit

import (
	"fmt"
	"math/rand"
	"strings"
)

// Generate Domain Name
func DomainName() string {
	return strings.ToLower(JobDescriptor()+BS()) + "." + DomainSuffix()
}

// Generate Domain Suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

// Generate Url
func Url() string {
	url := "http" + SliceString([]string{"s", ""}) + "://www."
	url += DomainName()

	// Slugs
	num := Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = BS()
	}
	url += "/" + strings.ToLower(strings.Join(slug, "/"))

	return url
}

// Generate IP Version 4
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// Generate IP Version 6
func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// Generate Username
func Username() string {
	return getRandValue([]string{"name", "last"}) + replaceWithNumbers("####")
}

// Generator Password
func Password(lower bool, upper bool, numeric bool, special bool, space bool, length int) string {
	var passString string = ""
	var lowerStr string = "abcdefghijklmnopqrstuvwxyz"
	var upperStr string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numericStr string = "0123456789"
	var specialStr string = "!@#$%&?-_"
	var spaceStr string = " "

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
		finalBytes[i] = passBytes[randIntRange(0, len(passBytes))]
	}
	return string(finalBytes)
}
