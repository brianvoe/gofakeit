package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleDomainName() {
	Seed(11)
	fmt.Println(DomainName())
	// Output: centraltarget.biz
}

func BenchmarkDomainName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainName()
	}
}

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())
	// Output: org
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix()
	}
}

func ExampleURL() {
	Seed(11)
	fmt.Println(URL())
	// Output: http://www.principalproductize.biz/target
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL()
	}
}

func ExampleHTTPMethod() {
	Seed(11)
	fmt.Println(HTTPMethod())
	// Output: HEAD
}

func BenchmarkHTTPMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPMethod()
	}
}

func ExampleIPv4Address() {
	Seed(11)
	fmt.Println(IPv4Address())
	// Output: 222.83.191.222
}

func BenchmarkIPv4Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4Address()
	}
}

func ExampleIPv6Address() {
	Seed(11)
	fmt.Println(IPv6Address())
	// Output: 2001:cafe:8898:ee17:bc35:9064:5866:d019
}

func BenchmarkIPv6Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv6Address()
	}
}

func ExampleMacAddress() {
	Seed(11)
	fmt.Println(MacAddress())
	// Output: e1:74:cb:01:77:91
}

func BenchmarkMacAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MacAddress()
	}
}

func ExampleHTTPStatusCode() {
	Seed(11)
	fmt.Println(HTTPStatusCode())
	// Output: 404
}

func BenchmarkHTTPStatusCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPStatusCode()
	}
}

func ExampleHTTPStatusCodeSimple() {
	Seed(11)
	fmt.Println(HTTPStatusCodeSimple())
	// Output: 200
}

func BenchmarkHTTPStatusCodeSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPStatusCodeSimple()
	}
}

func ExampleLogLevel() {
	Seed(11)
	fmt.Println(LogLevel("")) // This will also use general
	fmt.Println(LogLevel("syslog"))
	fmt.Println(LogLevel("apache"))
	// Output: error
	// debug
	// trace1-8
}

func BenchmarkLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogLevel("general")
	}
}

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())
	// Output: Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362
}

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserAgent()
	}
}

func TestUserAgent(t *testing.T) {
	for i := 0; i < 100; i++ {
		UserAgent()
	}
}

func ExampleChromeUserAgent() {
	Seed(11)
	fmt.Println(ChromeUserAgent())
	// Output: Mozilla/5.0 (X11; Linux i686) AppleWebKit/5312 (KHTML, like Gecko) Chrome/39.0.836.0 Mobile Safari/5312
}

func BenchmarkChromeUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChromeUserAgent()
	}
}

func ExampleFirefoxUserAgent() {
	Seed(11)
	fmt.Println(FirefoxUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1977-07-01 Firefox/37.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirefoxUserAgent()
	}
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())
	// Output: Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafariUserAgent()
	}
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())
	// Output: Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OperaUserAgent()
	}
}
