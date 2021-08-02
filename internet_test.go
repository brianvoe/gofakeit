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

func ExampleFaker_DomainName() {
	f := New(11)
	fmt.Println(f.DomainName())
	// Output: centraltarget.biz
}

func BenchmarkDomainName(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DomainName()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.DomainName()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.DomainName()
		}
	})
}

func ExampleDomainSuffix() {
	Seed(11)
	fmt.Println(DomainSuffix())
	// Output: org
}

func ExampleFaker_DomainSuffix() {
	f := New(11)
	fmt.Println(f.DomainSuffix())
	// Output: org
}

func BenchmarkDomainSuffix(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			DomainSuffix()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.DomainSuffix()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.DomainSuffix()
		}
	})
}

func ExampleURL() {
	Seed(11)
	fmt.Println(URL())
	// Output: http://www.principalproductize.biz/target
}

func ExampleFaker_URL() {
	f := New(11)
	fmt.Println(f.URL())
	// Output: http://www.principalproductize.biz/target
}

func BenchmarkURL(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			URL()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.URL()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.URL()
		}
	})
}

func ExampleHTTPMethod() {
	Seed(11)
	fmt.Println(HTTPMethod())
	// Output: HEAD
}

func ExampleHTTPVersion() {
	Seed(11)
	fmt.Println(HTTPVersion())
	// Output: HTTP/1.0
}

func ExampleFaker_HTTPMethod() {
	f := New(11)
	fmt.Println(f.HTTPMethod())
	// Output: HEAD
}

func ExampleFaker_HTTPVersion() {
	f := New(11)
	fmt.Println(f.HTTPVersion())
	// Output: HTTP/1.0
}

func BenchmarkHTTPMethod(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPMethod()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPMethod()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPMethod()
		}
	})
}

func ExampleIPv4Address() {
	Seed(11)
	fmt.Println(IPv4Address())
	// Output: 152.23.53.100
}

func ExampleFaker_IPv4Address() {
	f := New(11)
	fmt.Println(f.IPv4Address())
	// Output: 152.23.53.100
}

func BenchmarkIPv4Address(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IPv4Address()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.IPv4Address()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.IPv4Address()
		}
	})
}

func ExampleIPv6Address() {
	Seed(11)
	fmt.Println(IPv6Address())
	// Output: 8898:ee17:bc35:9064:5866:d019:3b95:7857
}

func ExampleFaker_IPv6Address() {
	f := New(11)
	fmt.Println(f.IPv6Address())
	// Output: 8898:ee17:bc35:9064:5866:d019:3b95:7857
}

func BenchmarkIPv6Address(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IPv6Address()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.IPv6Address()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.IPv6Address()
		}
	})
}

func ExampleMacAddress() {
	Seed(11)
	fmt.Println(MacAddress())
	// Output: e1:74:cb:01:77:91
}

func ExampleFaker_MacAddress() {
	f := New(11)
	fmt.Println(f.MacAddress())
	// Output: e1:74:cb:01:77:91
}

func BenchmarkMacAddress(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MacAddress()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.MacAddress()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.MacAddress()
		}
	})
}

func ExampleHTTPStatusCode() {
	Seed(11)
	fmt.Println(HTTPStatusCode())
	// Output: 404
}

func ExampleFaker_HTTPStatusCode() {
	f := New(11)
	fmt.Println(f.HTTPStatusCode())
	// Output: 404
}

func BenchmarkHTTPStatusCode(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPStatusCode()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPStatusCode()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPStatusCode()
		}
	})
}

func ExampleHTTPStatusCodeSimple() {
	Seed(11)
	fmt.Println(HTTPStatusCodeSimple())
	// Output: 200
}

func ExampleFaker_HTTPStatusCodeSimple() {
	f := New(11)
	fmt.Println(f.HTTPStatusCodeSimple())
	// Output: 200
}

func BenchmarkHTTPStatusCodeSimple(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HTTPStatusCodeSimple()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.HTTPStatusCodeSimple()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.HTTPStatusCodeSimple()
		}
	})
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

func ExampleFaker_LogLevel() {
	f := New(11)
	fmt.Println(f.LogLevel("")) // This will also use general
	fmt.Println(f.LogLevel("syslog"))
	fmt.Println(f.LogLevel("apache"))
	// Output: error
	// debug
	// trace1-8
}

func BenchmarkLogLevel(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			LogLevel("general")
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.LogLevel("general")
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.LogLevel("general")
		}
	})
}

func ExampleUserAgent() {
	Seed(11)
	fmt.Println(UserAgent())
	// Output: Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362
}

func ExampleFaker_UserAgent() {
	f := New(11)
	fmt.Println(f.UserAgent())
	// Output: Mozilla/5.0 (Windows NT 5.0) AppleWebKit/5362 (KHTML, like Gecko) Chrome/37.0.834.0 Mobile Safari/5362
}

func BenchmarkUserAgent(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			UserAgent()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.UserAgent()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.UserAgent()
		}
	})
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

func ExampleFaker_ChromeUserAgent() {
	f := New(11)
	fmt.Println(f.ChromeUserAgent())
	// Output: Mozilla/5.0 (X11; Linux i686) AppleWebKit/5312 (KHTML, like Gecko) Chrome/39.0.836.0 Mobile Safari/5312
}

func BenchmarkChromeUserAgent(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ChromeUserAgent()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.ChromeUserAgent()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.ChromeUserAgent()
		}
	})
}

func ExampleFirefoxUserAgent() {
	Seed(11)
	fmt.Println(FirefoxUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1908-07-12 Firefox/37.0
}

func ExampleFaker_FirefoxUserAgent() {
	f := New(11)
	fmt.Println(f.FirefoxUserAgent())
	// Output: Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_8_3 rv:7.0) Gecko/1908-07-12 Firefox/37.0
}

func BenchmarkFirefoxUserAgent(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FirefoxUserAgent()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.FirefoxUserAgent()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.FirefoxUserAgent()
		}
	})
}

func ExampleSafariUserAgent() {
	Seed(11)
	fmt.Println(SafariUserAgent())
	// Output: Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6
}

func ExampleFaker_SafariUserAgent() {
	f := New(11)
	fmt.Println(f.SafariUserAgent())
	// Output: Mozilla/5.0 (iPad; CPU OS 8_3_2 like Mac OS X; en-US) AppleWebKit/531.15.6 (KHTML, like Gecko) Version/4.0.5 Mobile/8B120 Safari/6531.15.6
}

func BenchmarkSafariUserAgent(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SafariUserAgent()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.SafariUserAgent()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.SafariUserAgent()
		}
	})
}

func ExampleOperaUserAgent() {
	Seed(11)
	fmt.Println(OperaUserAgent())
	// Output: Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00
}

func ExampleFaker_OperaUserAgent() {
	f := New(11)
	fmt.Println(f.OperaUserAgent())
	// Output: Opera/8.39 (Macintosh; U; PPC Mac OS X 10_8_7; en-US) Presto/2.9.335 Version/10.00
}

func BenchmarkOperaUserAgent(b *testing.B) {
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			OperaUserAgent()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.OperaUserAgent()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.OperaUserAgent()
		}
	})
}
