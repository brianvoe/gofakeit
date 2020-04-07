package gofakeit

import (
	"fmt"
	"math/rand"
	"strings"
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
	url := "http" + RandString([]string{"s", ""}) + "://www." + DomainName() + "/" + strings.ToLower(strings.Join(slug, "/"))
	url = strings.Replace(url, " ", "", -1)

	return url
}

// HTTPMethod will generate a random http method
func HTTPMethod() string {
	return getRandValue([]string{"internet", "http_method"})
}

// IPv4Address will generate a random version 4 ip address
func IPv4Address() string {
	num := func() int { return 2 + rand.Intn(254) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num())
}

// IPv6Address will generate a random version 6 ip address
func IPv6Address() string {
	num := 65536
	return fmt.Sprintf("2001:cafe:%x:%x:%x:%x:%x:%x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

// MacAddress will generate a random mac address
func MacAddress() string {
	num := 255
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num), rand.Intn(num))
}

func addInternetLookup() {
	AddLookupData("internet.url", Info{
		Description: "Random url",
		Example:     "http://www.principalproductize.biz/target",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return URL(), nil
		},
	})

	AddLookupData("internet.domain", Info{
		Description: "Random domain name",
		Example:     "centraltarget.biz",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return DomainName(), nil
		},
	})

	AddLookupData("internet.domain.suffix", Info{
		Description: "Random domain suffix",
		Example:     "org",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return DomainSuffix(), nil
		},
	})

	AddLookupData("internet.imageurl", Info{
		Description: "Random image url",
		Example:     "https://picsum.photos/640/480",
		Params: []Param{
			{Field: "width", Required: true, Type: "uint", Description: "Image width"},
			{Field: "height", Required: true, Type: "uint", Description: "Image height"},
		},
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			width, err := info.GetUint(m, "width")
			if err != nil {
				return nil, err
			}

			height, err := info.GetUint(m, "height")
			if err != nil {
				return nil, err
			}

			return ImageURL(int(width), int(height)), nil
		},
	})

	AddLookupData("internet.ip.v4", Info{
		Description: "Random ip address v4",
		Example:     "222.83.191.222",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return IPv4Address(), nil
		},
	})

	AddLookupData("internet.ip.v6", Info{
		Description: "Random ip address v6",
		Example:     "2001:cafe:8898:ee17:bc35:9064:5866:d019",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return IPv6Address(), nil
		},
	})

	AddLookupData("internet.http.statuscode", Info{
		Description: "Random http status code",
		Example:     "404",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return StatusCode(), nil
		},
	})

	AddLookupData("internet.http.method", Info{
		Description: "Random http method",
		Example:     "HEAD",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return HTTPMethod(), nil
		},
	})

	AddLookupData("internet.loglevel", Info{
		Description: "Random log level",
		Example:     "error",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return LogLevel(""), nil
		},
	})

	AddLookupData("internet.useragent", Info{
		Description: "Random browser user agent",
		Example:     "Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return UserAgent(), nil
		},
	})

	AddLookupData("internet.useragent.chrome", Info{
		Description: "Random chrome user agent",
		Example:     "Mozilla/5.0 (X11; Linux i686) AppleWebKit/5312 (KHTML, like Gecko) Chrome/39.0.836.0 Mobile Safari/5312",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return ChromeUserAgent(), nil
		},
	})

	AddLookupData("internet.useragent.firefox", Info{
		Description: "Random browser user agent",
		Example:     "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1900-07-01 Firefox/37.0",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return FirefoxUserAgent(), nil
		},
	})

	AddLookupData("internet.useragent.opera", Info{
		Description: "Random browser user agent",
		Example:     "Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return OperaUserAgent(), nil
		},
	})

	AddLookupData("internet.useragent.safari", Info{
		Description: "Random safari user agent",
		Example:     "Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6",
		Call: func(m *map[string]string, info *Info) (interface{}, error) {
			return SafariUserAgent(), nil
		},
	})
}
