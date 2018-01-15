goos: darwin  
goarch: amd64  
pkg: github.com/brianvoe/gofakeit  

||||||
|-|-|-|-|-|
|BenchmarkAddress-8              | 1000000|      2088 ns/op|     255 B/op|       7 allocs/op|
|BenchmarkStreet-8               | 2000000|       781 ns/op|      62 B/op|       3 allocs/op|
|BenchmarkStreetNumber-8         | 5000000|       395 ns/op|      36 B/op|       2 allocs/op|
|BenchmarkStreetPrefix-8         |20000000|       115 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkStreetName-8           |10000000|       124 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkStreetSuffix-8         |10000000|       128 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkCity-8                 | 5000000|       356 ns/op|      15 B/op|       1 allocs/op|
|BenchmarkState-8                |10000000|       118 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkStateAbr-8             |10000000|       116 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkZip-8                  | 5000000|       401 ns/op|       9 B/op|       1 allocs/op|
|BenchmarkCountry-8              |10000000|       140 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLatitude-8             |50000000|        25.4 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLongitude-8            |50000000|        25.0 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLatitudeInRange-8      |50000000|        30.6 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLongitudeInRange-8     |50000000|        29.7 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerName-8             |20000000|       107 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerStyle-8            |10000000|       128 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerHop-8              |20000000|       112 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerYeast-8            |20000000|       114 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerMalt-8             |10000000|       122 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBeerIbu-8              |20000000|        78.7 ns/op|       8 B/op|       1 allocs/op|
|BenchmarkBeerAlcohol-8          | 5000000|       388 ns/op|      40 B/op|       3 allocs/op|
|BenchmarkBeerBlg-8              | 3000000|       396 ns/op|      48 B/op|       3 allocs/op|
|BenchmarkBool-8                 |50000000|        35.8 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkColor-8                |20000000|       116 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSafeColor-8            |20000000|       109 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHexColor-8             | 3000000|       562 ns/op|      24 B/op|       3 allocs/op|
|BenchmarkRGBColor-8             |10000000|       114 ns/op|      32 B/op|       1 allocs/op|
|BenchmarkCompany-8              | 3000000|       397 ns/op|      22 B/op|       1 allocs/op|
|BenchmarkCompanySuffix-8        |20000000|       101 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBuzzWord-8             |20000000|       107 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkBS-8                   |10000000|       120 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkContact-8              | 1000000|      1411 ns/op|     179 B/op|       7 allocs/op|
|BenchmarkPhone-8                | 3000000|       477 ns/op|      16 B/op|       1 allocs/op|
|BenchmarkEmail-8                | 2000000|       790 ns/op|     131 B/op|       5 allocs/op|
|BenchmarkCurrency-8             |10000000|       133 ns/op|      32 B/op|       1 allocs/op|
|BenchmarkCurrencyShort-8        |20000000|       109 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkCurrencyLong-8         |20000000|       111 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkDate-8                 | 1000000|      1119 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkDateRange-8            | 1000000|      1608 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkMonth-8                |30000000|        47.5 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkDay-8                  |50000000|        41.6 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkWeekDay-8              |30000000|        44.0 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkYear-8                 | 2000000|       816 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHour-8                 |30000000|        42.5 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkMinute-8               |30000000|        41.8 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSecond-8               |30000000|        41.7 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkNanoSecond-8           |30000000|        43.1 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkMimeType-8             |20000000|       100 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkExtension-8            |10000000|       113 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkGenerate-8             | 1000000|      1689 ns/op|     414 B/op|      11 allocs/op|
|BenchmarkHackerPhrase-8         |  300000|      4850 ns/op|    2296 B/op|      26 allocs/op|
|BenchmarkHackerAbbreviation-8   |20000000|       112 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHackerAdjective-8      |20000000|       115 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHackerNoun-8           |10000000|       115 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHackerVerb-8           |10000000|       123 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHackerIngverb-8        |10000000|       111 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHipsterWord-8          |20000000|       108 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkHipsterSentence-8      | 1000000|      2018 ns/op|     787 B/op|       9 allocs/op|
|BenchmarkHipsterParagraph-8     |   50000|     34627 ns/op|   18582 B/op|     170 allocs/op|
|BenchmarkImageURL-8             |10000000|       113 ns/op|      38 B/op|       3 allocs/op|
|BenchmarkDomainName-8           | 3000000|       542 ns/op|      83 B/op|       3 allocs/op|
|BenchmarkDomainSuffix-8         |20000000|        98.7 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkURL-8                  | 1000000|      1314 ns/op|     285 B/op|       8 allocs/op|
|BenchmarkHTTPMethod-8           |20000000|        99.9 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkIPv4Address-8          | 3000000|       388 ns/op|      48 B/op|       5 allocs/op|
|BenchmarkIPv6Address-8          | 3000000|       546 ns/op|      96 B/op|       7 allocs/op|
|BenchmarkUsername-8             | 5000000|       372 ns/op|      16 B/op|       2 allocs/op|
|BenchmarkPassword-8             | 1000000|      1267 ns/op|     432 B/op|       7 allocs/op|
|BenchmarkJob-8                  | 2000000|       751 ns/op|      86 B/op|       2 allocs/op|
|BenchmarkJobTitle-8             |20000000|       104 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkJobDescriptor-8        |20000000|       106 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkJobLevel-8             |10000000|       126 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkName-8                 | 5000000|       287 ns/op|      17 B/op|       1 allocs/op|
|BenchmarkFirstName-8            |20000000|       100 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLastName-8             |10000000|       104 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkNamePrefix-8           |20000000|        95.4 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkNameSuffix-8           |10000000|       109 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkNumber-8               |50000000|        35.7 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkNumerify-8             | 5000000|       340 ns/op|      16 B/op|       1 allocs/op|
|BenchmarkShuffleInts-8          | 5000000|       215 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkCreditCard-8           | 1000000|      1019 ns/op|      88 B/op|       4 allocs/op|
|BenchmarkCreditCardType-8       |20000000|        93.0 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkCreditCardNumber-8     | 2000000|       610 ns/op|      16 B/op|       1 allocs/op|
|BenchmarkCreditCardExp-8        |10000000|       133 ns/op|       5 B/op|       1 allocs/op|
|BenchmarkCreditCardCvv-8        |10000000|       141 ns/op|       3 B/op|       1 allocs/op|
|BenchmarkSSN-8                  | 3000000|       375 ns/op|      16 B/op|       1 allocs/op|
|BenchmarkGender-8               |30000000|        45.7 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkPerson-8               |  200000|      6481 ns/op|     821 B/op|      26 allocs/op|
|BenchmarkSimpleStatusCode-8     |20000000|        72.3 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkStatusCode-8           |20000000|        72.3 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkLetter-8               |30000000|        57.9 ns/op|       4 B/op|       1 allocs/op|
|BenchmarkLexify-8               |10000000|       213 ns/op|       8 B/op|       1 allocs/op|
|BenchmarkShuffleStrings-8       |10000000|       223 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkUUID-8                 |20000000|       105 ns/op|      48 B/op|       1 allocs/op|
|BenchmarkUserAgent-8            | 1000000|      1428 ns/op|     305 B/op|       5 allocs/op|
|BenchmarkChromeUserAgent-8      | 2000000|       907 ns/op|     188 B/op|       5 allocs/op|
|BenchmarkFirefoxUserAgent-8     |  500000|      2531 ns/op|     386 B/op|       7 allocs/op|
|BenchmarkSafariUserAgent-8      | 1000000|      1695 ns/op|     551 B/op|       7 allocs/op|
|BenchmarkOperaUserAgent-8       | 1000000|      1105 ns/op|     216 B/op|       5 allocs/op|
|BenchmarkWord-8                 |20000000|       108 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSentence-8             | 1000000|      2019 ns/op|     727 B/op|      10 allocs/op|
|BenchmarkParagraph-8            |   30000|     38478 ns/op|   16199 B/op|     172 allocs/op|
