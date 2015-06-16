package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())
	// Output: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362
}

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
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
	// Output: Mozilla/5.0 (X11; Linux x86_64; rv:6.0) Gecko/1912-07-01 Firefox/37.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirefoxUserAgent()
	}
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())
	// Output: Mozilla/5.0 (Windows; U; Windows NT 6.1) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/5.2 Safari/531.15.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SafariUserAgent()
	}
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())
	// Output: Opera/9.15 (X11; Linux i686; en-US) Presto/2.12.262 Version/11.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OperaUserAgent()
	}
}
