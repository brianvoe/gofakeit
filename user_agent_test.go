package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())
	// Output: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5321 (KHTML, like Gecko) Chrome/36.0.883.0 Mobile Safari/5321
}

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UserAgent()
	}
}

func ExampleChromeUserAgent() {
	Seed(11)
	fmt.Println(ChromeUserAgent())
	// Output: Mozilla/5.0 (X11; Linux i686) AppleWebKit/5311 (KHTML, like Gecko) Chrome/37.0.807.0 Mobile Safari/5311
}

func BenchmarkChromeUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChromeUserAgent()
	}
}

func ExampleFirefoxUserAgent() {
	Seed(11)
	fmt.Println(FirefoxUserAgent())
	// Output: Mozilla/5.0 (X11; Linux x86_64; rv:6.0) Gecko/1925-24-11 Firefox/36.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirefoxUserAgent()
	}
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())
	// Output: Mozilla/5.0 (Windows; U; Windows NT 6.1) AppleWebKit/531.2.2 (KHTML, like Gecko) Version/4.0 Safari/531.2.2
}

func BenchmarkSafariUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafariUserAgent()
	}
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())
	// Output: Opera/9.57 (X11; Linux i686; en-US) Presto/2.12.298 Version/12.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OperaUserAgent()
	}
}
