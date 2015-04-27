package gofakeit

import (
	"fmt"
	"math/rand"
)

// Generate Domain Suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

// Generate
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}
