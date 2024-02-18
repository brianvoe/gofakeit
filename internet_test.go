package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleDomainName() {
	Seed(11)
	fmt.Println(DomainName())

	// Output: productinfrastructures.biz
}

func ExampleFaker_DomainName() {
	f := New(11)
	fmt.Println(f.DomainName())

	// Output: productinfrastructures.biz
}

func BenchmarkDomainName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainName()
	}
}

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())

	// Output: io
}

func ExampleFaker_DomainSuffix() {
	f := New(11)
	fmt.Println(f.DomainSuffix())

	// Output: io
}

func BenchmarkDomainSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DomainSuffix()
	}
}

func ExampleURL() {
	Seed(11)
	fmt.Println(URL())

	// Output: http://www.directinnovative.biz/infrastructures
}

func ExampleFaker_URL() {
	f := New(11)
	fmt.Println(f.URL())

	// Output: http://www.directinnovative.biz/infrastructures
}

func TestURLValid(t *testing.T) {
	for i := 0; i < 10000; i++ {
		url := URL()

		// Check if url has spaces in it
		if strings.Contains(url, " ") {
			t.Error("URL has spaces")
		}
	}
}

func BenchmarkURL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URL()
	}
}

func ExampleHTTPMethod() {
	Seed(11)
	fmt.Println(HTTPMethod())

	// Output: DELETE
}

func ExampleHTTPVersion() {
	Seed(11)
	fmt.Println(HTTPVersion())

	// Output: HTTP/2.0
}

func ExampleFaker_HTTPMethod() {
	f := New(11)
	fmt.Println(f.HTTPMethod())

	// Output: DELETE
}

func ExampleFaker_HTTPVersion() {
	f := New(11)
	fmt.Println(f.HTTPVersion())

	// Output: HTTP/2.0
}

func BenchmarkHTTPMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPMethod()
	}
}

func ExampleIPv4Address() {
	Seed(11)
	fmt.Println(IPv4Address())

	// Output: 180.18.181.251
}

func ExampleFaker_IPv4Address() {
	f := New(11)
	fmt.Println(f.IPv4Address())

	// Output: 180.18.181.251
}

func BenchmarkIPv4Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv4Address()
	}
}

func ExampleIPv6Address() {
	Seed(11)
	fmt.Println(IPv6Address())

	// Output: ddb4:9212:aab5:87fb:4e33:17a4:f7b9:bf8e
}

func ExampleFaker_IPv6Address() {
	f := New(11)
	fmt.Println(f.IPv6Address())

	// Output: ddb4:9212:aab5:87fb:4e33:17a4:f7b9:bf8e
}

func BenchmarkIPv6Address(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPv6Address()
	}
}

func ExampleMacAddress() {
	Seed(11)
	fmt.Println(MacAddress())

	// Output: e4:da:32:33:86:3b
}

func ExampleFaker_MacAddress() {
	f := New(11)
	fmt.Println(f.MacAddress())

	// Output: e4:da:32:33:86:3b
}

func BenchmarkMacAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MacAddress()
	}
}

func ExampleHTTPStatusCode() {
	Seed(11)
	fmt.Println(HTTPStatusCode())

	// Output: 502
}

func ExampleFaker_HTTPStatusCode() {
	f := New(11)
	fmt.Println(f.HTTPStatusCode())

	// Output: 502
}

func BenchmarkHTTPStatusCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HTTPStatusCode()
	}
}

func ExampleHTTPStatusCodeSimple() {
	Seed(11)
	fmt.Println(HTTPStatusCodeSimple())

	// Output: 500
}

func ExampleFaker_HTTPStatusCodeSimple() {
	f := New(11)
	fmt.Println(f.HTTPStatusCodeSimple())

	// Output: 500
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

	// Output: debug
	// crit
	// alert
}

func ExampleFaker_LogLevel() {
	f := New(11)
	fmt.Println(f.LogLevel("")) // This will also use general
	fmt.Println(f.LogLevel("syslog"))
	fmt.Println(f.LogLevel("apache"))

	// Output: debug
	// crit
	// alert
}

func BenchmarkLogLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogLevel("general")
	}
}

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())

	// Output: Mozilla/5.0 (Windows 98) AppleWebKit/5360 (KHTML, like Gecko) Chrome/37.0.852.0 Mobile Safari/5360
}

func ExampleFaker_UserAgent() {
	f := New(11)
	fmt.Println(f.UserAgent())

	// Output: Mozilla/5.0 (Windows 98) AppleWebKit/5360 (KHTML, like Gecko) Chrome/37.0.852.0 Mobile Safari/5360
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

	// Output: Mozilla/5.0 (Windows CE) AppleWebKit/5362 (KHTML, like Gecko) Chrome/36.0.820.0 Mobile Safari/5362
}

func ExampleFaker_ChromeUserAgent() {
	f := New(11)
	fmt.Println(f.ChromeUserAgent())

	// Output: Mozilla/5.0 (Windows CE) AppleWebKit/5362 (KHTML, like Gecko) Chrome/36.0.820.0 Mobile Safari/5362
}

func BenchmarkChromeUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChromeUserAgent()
	}
}

func ExampleFirefoxUserAgent() {
	Seed(11)
	fmt.Println(FirefoxUserAgent())

	// Output: Mozilla/5.0 (Windows CE; en-US; rv:1.9.3.20) Gecko/2011-11-07 Firefox/36.0
}

func ExampleFaker_FirefoxUserAgent() {
	f := New(11)
	fmt.Println(f.FirefoxUserAgent())

	// Output: Mozilla/5.0 (Windows CE; en-US; rv:1.9.3.20) Gecko/2011-11-07 Firefox/36.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirefoxUserAgent()
	}
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())

	// Output: Mozilla/5.0 (Windows; U; Windows NT 6.0) AppleWebKit/536.44.6 (KHTML, like Gecko) Version/4.1 Safari/536.44.6
}

func ExampleFaker_SafariUserAgent() {
	f := New(11)
	fmt.Println(f.SafariUserAgent())

	// Output: Mozilla/5.0 (Windows; U; Windows NT 6.0) AppleWebKit/536.44.6 (KHTML, like Gecko) Version/4.1 Safari/536.44.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafariUserAgent()
	}
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())

	// Output: Opera/10.91 (X11; Linux i686; en-US) Presto/2.12.265 Version/11.00
}

func ExampleFaker_OperaUserAgent() {
	f := New(11)
	fmt.Println(f.OperaUserAgent())

	// Output: Opera/10.91 (X11; Linux i686; en-US) Presto/2.12.265 Version/11.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OperaUserAgent()
	}
}
