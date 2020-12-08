package gofakeit

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v5/data"
)

// DomainName will generate a random url domain name
func DomainName() string {
	return strings.Replace(strings.ToLower(JobDescriptor()+BS()), " ", "", -1) + "." + DomainSuffix()
}

// DomainSuffix will generate a random domain suffix
func DomainSuffix() string {
	return getRandValue([]string{"internet", "domain_suffix"})
}

// URL will generate a random url string
func URL() string {
	// Slugs
	num := Number(1, 4)
	slug := make([]string, num)
	for i := 0; i < num; i++ {
		slug[i] = BS()
	}
	url := "http" + RandomString([]string{"s", ""}) + "://www." + DomainName() + "/" + strings.ToLower(strings.Join(slug, "/"))
	url = strings.Replace(url, " ", "", -1)

	return url
}

// HTTPMethod will generate a random http method
func HTTPMethod() string {
	return getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func IPv4Address() string {
	num := func() int { return rand.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func IPv6Address() string {
	num := func() int { return rand.Intn(65536) }
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", num(), num(), num(), num(), num(), num(), num(), num())
}

// MacAddress will generate a random mac address
func MacAddress() string {
	num := 255
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// HTTPStatusCode will generate a random status code
func HTTPStatusCode() int {
	return getRandIntValue([]string{"status_code", "general"})
}

// HTTPStatusCodeSimple will generate a random simple status code
func HTTPStatusCodeSimple() int {
	return getRandIntValue([]string{"status_code", "simple"})
}

// LogLevel will generate a random log level
// See data/LogLevels for list of available levels
func LogLevel(logType string) string {
	if _, ok := data.LogLevels[logType]; ok {
		return getRandValue([]string{"log_level", logType})
	}

	return getRandValue([]string{"log_level", "general"})
}

// UserAgent will generate a random broswer user agent
func UserAgent() string {
	randNum := randIntRange(0, 4)
	switch randNum {
	case 0:
		return ChromeUserAgent()
	case 1:
		return FirefoxUserAgent()
	case 2:
		return SafariUserAgent()
	case 3:
		return OperaUserAgent()
	default:
		return ChromeUserAgent()
	}
}

// ChromeUserAgent will generate a random chrome browser user agent string
func ChromeUserAgent() string {
	randNum1 := strconv.Itoa(randIntRange(531, 536)) + strconv.Itoa(randIntRange(0, 2))
	randNum2 := strconv.Itoa(randIntRange(36, 40))
	randNum3 := strconv.Itoa(randIntRange(800, 899))
	return "Mozilla/5.0 " + "(" + randomPlatform() + ") AppleWebKit/" + randNum1 + " (KHTML, like Gecko) Chrome/" + randNum2 + ".0." + randNum3 + ".0 Mobile Safari/" + randNum1
}

// FirefoxUserAgent will generate a random firefox broswer user agent string
func FirefoxUserAgent() string {
	ver := "Gecko/" + Date().Format("2006-02-01") + " Firefox/" + strconv.Itoa(randIntRange(35, 37)) + ".0"
	platforms := []string{
		"(" + windowsPlatformToken() + "; " + "en-US" + "; rv:1.9." + strconv.Itoa(randIntRange(0, 3)) + ".20) " + ver,
		"(" + linuxPlatformToken() + "; rv:" + strconv.Itoa(randIntRange(5, 8)) + ".0) " + ver,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(randIntRange(2, 7)) + ".0) " + ver,
	}

	return "Mozilla/5.0 " + RandomString(platforms)
}

// SafariUserAgent will generate a random safari browser user agent string
func SafariUserAgent() string {
	randNum := strconv.Itoa(randIntRange(531, 536)) + "." + strconv.Itoa(randIntRange(1, 51)) + "." + strconv.Itoa(randIntRange(1, 8))
	ver := strconv.Itoa(randIntRange(4, 6)) + "." + strconv.Itoa(randIntRange(0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + windowsPlatformToken() + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(randIntRange(4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + RandomString(mobileDevices) + " " + strconv.Itoa(randIntRange(7, 9)) + "_" + strconv.Itoa(randIntRange(0, 3)) + "_" + strconv.Itoa(randIntRange(1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(randIntRange(3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(randIntRange(111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + RandomString(platforms)
}

// OperaUserAgent will generate a random opera browser user agent string
func OperaUserAgent() string {
	platform := "(" + randomPlatform() + "; en-US) Presto/2." + strconv.Itoa(randIntRange(8, 13)) + "." + strconv.Itoa(randIntRange(160, 355)) + " Version/" + strconv.Itoa(randIntRange(10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(randIntRange(8, 10)) + "." + strconv.Itoa(randIntRange(10, 99)) + " " + platform
}

// linuxPlatformToken will generate a random linux platform
func linuxPlatformToken() string {
	return "X11; Linux " + getRandValue([]string{"computer", "linux_processor"})
}

// macPlatformToken will generate a random mac platform
func macPlatformToken() string {
	return "Macintosh; " + getRandValue([]string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(randIntRange(5, 9)) + "_" + strconv.Itoa(randIntRange(0, 10))
}

// windowsPlatformToken will generate a random windows platform
func windowsPlatformToken() string {
	return getRandValue([]string{"computer", "windows_platform"})
}

// randomPlatform will generate a random platform
func randomPlatform() string {
	platforms := []string{
		linuxPlatformToken(),
		macPlatformToken(),
		windowsPlatformToken(),
	}

	return RandomString(platforms)
}

func addInternetLookup() {
	AddFuncLookup("url", Info{
		Display:     "URL",
		Category:    "internet",
		Description: "Random url",
		Example:     "http://www.principalproductize.biz/target",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return URL(), nil
		},
	})

	AddFuncLookup("domainname", Info{
		Display:     "Domain Name",
		Category:    "internet",
		Description: "Random domain name",
		Example:     "centraltarget.biz",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return DomainName(), nil
		},
	})

	AddFuncLookup("domainsuffix", Info{
		Display:     "Domain Suffix",
		Category:    "internet",
		Description: "Random domain suffix",
		Example:     "org",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return DomainSuffix(), nil
		},
	})

	AddFuncLookup("ipv4address", Info{
		Display:     "IPv4 Address",
		Category:    "internet",
		Description: "Random ip address v4",
		Example:     "222.83.191.222",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return IPv4Address(), nil
		},
	})

	AddFuncLookup("ipv6address", Info{
		Display:     "IPv6 Address",
		Category:    "internet",
		Description: "Random ip address v6",
		Example:     "2001:cafe:8898:ee17:bc35:9064:5866:d019",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return IPv6Address(), nil
		},
	})

	AddFuncLookup("httpmethod", Info{
		Display:     "HTTP Method",
		Category:    "internet",
		Description: "Random http method",
		Example:     "HEAD",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HTTPMethod(), nil
		},
	})

	AddFuncLookup("loglevel", Info{
		Display:     "Log Level",
		Category:    "internet",
		Description: "Random log level",
		Example:     "error",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return LogLevel(""), nil
		},
	})

	AddFuncLookup("useragent", Info{
		Display:     "User Agent",
		Category:    "internet",
		Description: "Random browser user agent",
		Example:     "Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return UserAgent(), nil
		},
	})

	AddFuncLookup("chromeuseragent", Info{
		Display:     "Chrome User Agent",
		Category:    "internet",
		Description: "Random chrome user agent",
		Example:     "Mozilla/5.0 (X11; Linux i686) AppleWebKit/5312 (KHTML, like Gecko) Chrome/39.0.836.0 Mobile Safari/5312",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return ChromeUserAgent(), nil
		},
	})

	AddFuncLookup("firefoxuseragent", Info{
		Display:     "Firefox User Agent",
		Category:    "internet",
		Description: "Random browser user agent",
		Example:     "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1900-07-01 Firefox/37.0",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return FirefoxUserAgent(), nil
		},
	})

	AddFuncLookup("operauseragent", Info{
		Display:     "Opera User Agent",
		Category:    "internet",
		Description: "Random browser user agent",
		Example:     "Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return OperaUserAgent(), nil
		},
	})

	AddFuncLookup("safariuseragent", Info{
		Display:     "Safari User Agent",
		Category:    "internet",
		Description: "Random safari user agent",
		Example:     "Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return SafariUserAgent(), nil
		},
	})

	AddFuncLookup("httpstatuscode", Info{
		Display:     "HTTP Status Code",
		Category:    "internet",
		Description: "Random http status code",
		Example:     "200",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HTTPStatusCode(), nil
		},
	})

	AddFuncLookup("httpstatuscodesimple", Info{
		Display:     "HTTP Status Code Simple",
		Category:    "internet",
		Description: "Random http status code within more general usage codes",
		Example:     "404",
		Output:      "int",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return HTTPStatusCodeSimple(), nil
		},
	})
}
