![alt text](https://raw.githubusercontent.com/brianvoe/gofakeit/master/logo.png)

# gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) [![Build Status](https://travis-ci.org/brianvoe/gofakeit.svg?branch=master)](https://travis-ci.org/brianvoe/gofakeit) [![codecov.io](https://codecov.io/github/brianvoe/gofakeit/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/gofakeit) [![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/gofakeit/master/LICENSE.txt)
Random data generator written in go

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

### Features
- Every function has an example and a benchmark,
[see benchmarks](https://github.com/brianvoe/gofakeit/blob/master/BENCHMARKS.md)
- Zero dependencies
- Randomizes user defined structs
- Numerous functions for regular use

### 120+ Functions!!!
If there is something that is generic enough missing from this package [add an issue](https://github.com/brianvoe/gofakeit/issues) and let me know what you need.
Most of the time i'll add it!

## Person
```go
Person()
Name()
NamePrefix()
NameSuffix()
FirstName()
LastName()
Gender()
SSN()
Contact()
Email()
Phone()
PhoneFormatted()
Username()
Password()
```

## Address
```go
Address()
City()
Country()
CountryAbr()
State()
StateAbr()
StatusCode()
Street()
StreetName()
StreetNumber()
StreetPrefix()
StreetSuffix()
Zip()
Latitude()
LatitudeInRange()
Longitude()
LongitudeInRange()
```

## Beer
```go
BeerAlcohol()
BeerBlg()
BeerHop()
BeerIbu()
BeerMalt()
BeerName()
BeerStyle()
BeerYeast()
```

## Cars
```go
CarMaker()
CarModel()
Vehicle()
VehicleType()
FuelType()
TransmissionGearType()
```

## Lorem
```go
Word()
Sentence()
Paragraph()
```

## Shuffle
```go
ShuffleInts()
ShuffleStrings()
```

## Misc
```go
Struct()
Generate()
Letter()
Lexify()
Bool()
BS()
BuzzWord()
UUID()
```

## Colors
```go
Color()
HexColor()
RGBColor()
SafeColor()
```

## Internet
```go
URL()
ImageURL()
DomainName()
DomainSuffix()
IPv4Address()
IPv6Address()
SimpleStatusCode()
LogLevel()
HTTPMethod()
UserAgent()
ChromeUserAgent()
FirefoxUserAgent()
OperaUserAgent()
SafariUserAgent()
```

## Date/Time
```go
Date()
DateRange()
NanoSecond()
Second()
Minute()
Hour()
Month()
Day()
WeekDay()
Year()
TimeZone()
TimeZoneAbv()
TimeZoneFull()
TimeZoneOffset()
```

## Payment
```go
Price()
CreditCard()
CreditCardCvv()
CreditCardExp()
CreditCardNumber()
CreditCardNumberLuhn()
CreditCardType()
Currency()
CurrencyLong()
CurrencyShort()
```

## Company
```go
Company()
CompanySuffix()
Job()
JobDescriptor()
JobLevel()
JobTitle()
```

## Hacker
```go
HackerAbbreviation()
HackerAdjective()
HackerIngverb()
HackerNoun()
HackerPhrase()
HackerVerb()
```

## Hipster
```go
HipsterParagraph()
HipsterSentence()
HipsterWord()
```

## File
```go
Extension()
MimeType()
```

## Numbers
```go
Number()
Numerify()
Int8()
Int16()
Int32()
Int64()
Uint8()
Uint16()
Uint32()
Uint64()
Float32()
Float32Range()
Float64()
Float64Range()
```

## String
```go
Digit()
```

## Documentation
[![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit)

## Example
```go
import "github.com/brianvoe/gofakeit"

gofakeit.Name() // Markus Moen
gofakeit.Email() // alaynawuckert@kozey.biz
gofakeit.Phone() // (570)245-7485
gofakeit.BS() // front-end
gofakeit.BeerName() // Duvel
gofakeit.Color() // MediumOrchid
gofakeit.Company() // Moen, Pagac and Wuckert
gofakeit.CreditCardNumber() // 4287271570245748
gofakeit.HackerPhrase() // Connecting the array won't do anything, we need to generate the haptic COM driver!
gofakeit.JobTitle() // Director
gofakeit.Password(true, true, true, true, true, 32) // WV10MzLxq2DX79w1omH97_0ga59j8!kj
gofakeit.CurrencyShort() // USD
// 120+ more!!!

// Create structs with random injected data
type Foo struct {
	Bar     string
	Baz     string
	Int     int
	Pointer *int
	Skip    *string `fake:"skip"` // Set to "skip" to not generate data for
}
var f Foo
gofakeit.Struct(&f)
fmt.Printf("f.Bar:%s\n", f.Bar) // f.Bar:hrukpttuezptneuvunh
fmt.Printf("f.Baz:%s\n", f.Baz) // f.Baz:uksqvgzadxlgghejkmv
fmt.Printf("f.Int:%d\n", f.Int) // f.Int:-7825289004089916589
fmt.Printf("f.Pointer:%d\n", *f.Pointer) // f.Pointer:-343806609094473732
fmt.Printf("f.Skip:%v\n", f.Skip) // f.Skip:<nil>
```
