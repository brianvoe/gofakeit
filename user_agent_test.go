package gofakeit

import (
	"fmt"
	"testing"
)

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
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1900-07-01 Firefox/37.0
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
