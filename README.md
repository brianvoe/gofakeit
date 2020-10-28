![alt text](https://raw.githubusercontent.com/brianvoe/gofakeit/master/logo.png)

# Gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) ![Go](https://github.com/brianvoe/gofakeit/workflows/Go/badge.svg) [![codecov.io](https://codecov.io/github/brianvoe/gofakeit/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/gofakeit) [![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/gofakeit/master/LICENSE.txt)
Random data generator written in go

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## Features
- [160+ Functions!!!](#functions)
- [Struct Generator](#example-struct)
- [Custom Functions](#example-custom-functions)
- [Http Server](https://github.com/brianvoe/gofakeit/tree/master/cmd/gofakeitserver)
- [Command Line Tool](https://github.com/brianvoe/gofakeit/tree/master/cmd/gofakeit)
- Zero dependencies
- [Benchmarks](https://github.com/brianvoe/gofakeit/blob/master/BENCHMARKS.md)
- [Issue](https://github.com/brianvoe/gofakeit/issues)

## Installation
```go
go get github.com/brianvoe/gofakeit/v5
```

## Example
```go
import "github.com/brianvoe/gofakeit/v5"

gofakeit.Seed(0)

gofakeit.Name()             // Markus Moen
gofakeit.Email()            // alaynawuckert@kozey.biz
gofakeit.Phone()            // (570)245-7485
gofakeit.BS()               // front-end
gofakeit.BeerName()         // Duvel
gofakeit.Color()            // MediumOrchid
gofakeit.Company()          // Moen, Pagac and Wuckert
gofakeit.CreditCardNumber() // 4287271570245748
gofakeit.HackerPhrase()     // Connecting the array won't do anything, we need to generate the haptic COM driver!
gofakeit.JobTitle()         // Director
gofakeit.CurrencyShort()    // USD
// See full list below
```

## Example Struct
```go
import "github.com/brianvoe/gofakeit/v5"

// Create structs with random injected data
type Foo struct {
	Bar      string
	Int      int
	Pointer  *int
	Name     string  `fake:"{firstname}"`   // Any available function all lowercase
	Sentence string  `fake:"{sentence:3}"`  // Can call with parameters
	RandStr  string  `fake:"{randomstring:[hello,world]}"`
	Number   string  `fake:"{number:1,10}"` // Comma separated for multiple values
	Skip     *string `fake:"skip"`          // Set to "skip" to not generate data for
}

type FooBar struct {
	Bars    []string `fake:"{name}"`              // Array of random size (1-10) with fake function applied
	Foos    []Foo    `fakesize:"3"`               // Array of size specified with faked struct
	FooBars []Foo    `fake:"{name}" fakesize:"3"` // Array of size 3 with fake function applied
}

// Pass your struct as a pointer
var f Foo
gofakeit.Struct(&f)
fmt.Println(f.Bar)      // hrukpttuezptneuvunh
fmt.Println(f.Int)      // -7825289004089916589
fmt.Println(*f.Pointer) // -343806609094473732
fmt.Println(f.Name)     // fred
fmt.Println(f.Sentence) // Record river mind.
fmt.Println(f.RandStr)  // world
fmt.Println(f.Number)   // 4
fmt.Println(f.Skip)     // <nil>

var fb FooBar
gofakeit.Struct(&fb)
fmt.Println(fb.Bars)      // [Charlie Senger]
fmt.Println(fb.Foos)      // [{blmfxy -2585154718894894116 0xc000317bc0 Emmy Attitude demand addition. hello 3 <nil>} {cplbf -1722374676852125164 0xc000317cb0 Viva Addition option link. hello 7 <nil>}]

```

## Example Custom Functions
```go
// Simple
AddFuncLookup("friendname", Info{
	Category:    "custom",
	Description: "Random friend name",
	Example:     "bill",
	Output:      "string",
	Call: func(m *map[string][]string, info *Info) (interface{}, error) {
		return RandomString([]string{"bill", "bob", "sally"}), nil
	},
})

// With Params
AddFuncLookup("jumbleword", Info{
	Category:    "jumbleword",
	Description: "Take a word and jumple it up",
	Example:     "loredlowlh",
	Output:      "string",
	Params: []Param{
		{Field: "word", Type: "int", Description: "Word you want to jumble"},
	},
	Call: func(m *map[string][]string, info *Info) (interface{}, error) {
		word, err := info.GetString(m, "word")
		if err != nil {
			return nil, err
		}

		split := strings.Split(word, "")
		ShuffleStrings(split)
		return strings.Join(split, ""), nil
	},
})

type Foo struct {
	FriendName string `fake:"{friendname}"`
	JumbleWord string `fake:"{jumbleword:helloworld}"`
}

var f Foo
Struct(&f)
fmt.Printf("%s", f.FriendName) // bill
fmt.Printf("%s", f.JumbleWord) // loredlowlh
```

## Functions
### File
```go
JSON(jo *JSONOptions) []byte
XML(xo *XMLOptions) []byte
Extension() string
MimeType() string
```

### Person
```go
Person() *PersonInfo
Name() string
NamePrefix() string
NameSuffix() string
FirstName() string
LastName() string
Gender() string
SSN() string
Contact() *ContactInfo
Email() string
Phone() string
PhoneFormatted() string
Teams(people []string, teams []string) map[string][]string
```

### Generate
```go
Struct(v interface{})
Map() map[string]interface{}
Generate(value string) string
```

### Auth
```go
Username() string
Password(lower bool, upper bool, numeric bool, special bool, space bool, num int) string
```

### Address
```go
Address() *AddressInfo
City() string
Country() string
CountryAbr() string
State() string
StateAbr() string
Street() string
StreetName() string
StreetNumber() string
StreetPrefix() string
StreetSuffix() string
Zip() string
Latitude() float64
LatitudeInRange(min, max float64) (float64, error)
Longitude() float64
LongitudeInRange(min, max float64) (float64, error)
```

### Game
```go
Gamertag() string
```

### Beer
```go
BeerAlcohol() string
BeerBlg() string
BeerHop() string
BeerIbu() string
BeerMalt() string
BeerName() string
BeerStyle() string
BeerYeast() string
```

### Cars
```go
Vehicle() *VehicleInfo
CarMaker() string
CarModel() string
VehicleType() string
FuelType() string
TransmissionGearType() string
```

### Words
```go
Noun() string
Verb() string
Adverb() string
Preposition() string
Adjective() string
Word() string
Sentence(wordCount int) string
Paragraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string
LoremIpsumWord() string
LoremIpsumSentence(wordCount int) string
LoremIpsumParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string
Question() string
Quote() string
Phrase() string
```

### Foods
```go
Fruit() string
Vegetable() string
Breakfast() string
Lunch() string
Dinner() string
Snack() string
Dessert() string
```

### Misc
```go
Bool() bool
UUID() string
```

### Colors
```go
Color() string
HexColor() string
RGBColor() []int
SafeColor() string
```

### Internet
```go
URL() string
ImageURL(width int, height int) string
DomainName() string
DomainSuffix() string
IPv4Address() string
IPv6Address() string
StatusCode() string
SimpleStatusCode() int
LogLevel(logType string) string
HTTPMethod() string
UserAgent() string
ChromeUserAgent() string
FirefoxUserAgent() string
OperaUserAgent() string
SafariUserAgent() string
```

### Date/Time
```go
Date() time.Time
DateRange(start, end time.Time) time.Time
NanoSecond() int
Second() int
Minute() int
Hour() int
Month() string
Day() int
WeekDay() string
Year() int
TimeZone() string
TimeZoneAbv() string
TimeZoneFull() string
TimeZoneOffset() float32
TimeZoneRegion() string
```

### Payment
```go
Price(min, max float64) float64
CreditCard() *CreditCardInfo
CreditCardCvv() string
CreditCardExp() string
CreditCardNumber(*CreditCardOptions) string
CreditCardType() string
Currency() *CurrencyInfo
CurrencyLong() string
CurrencyShort() string
AchRouting() string
AchAccount() string
BitcoinAddress() string
BitcoinPrivateKey() string
```

### Company
```go
BS() string
BuzzWord() string
Company() string
CompanySuffix() string
Job() *JobInfo
JobDescriptor() string
JobLevel() string
JobTitle() string
```

### Hacker
```go
HackerAbbreviation() string
HackerAdjective() string
HackerIngverb() string
HackerNoun() string
HackerPhrase() string
HackerVerb() string
```

### Hipster
```go
HipsterWord() string
HipsterSentence(wordCount int) string
HipsterParagraph(paragraphCount int, sentenceCount int, wordCount int, separator string) string
```

### App
```go
AppName() string
AppVersion() string
AppAuthor() string
```

### Animal
```go
PetName() string
Animal() string
AnimalType() string
FarmAnimal() string
Cat() string
Dog() string
```

### Emoji
```go
Emoji() string // 🤣
EmojiDescription() string // winking face
EmojiCategory() string // Smileys & Emotion
EmojiAlias() string // smiley
EmojiTag() string // happy
```

### Languages
```go
Language() string
LanguageAbbreviation() string
ProgrammingLanguage() string
ProgrammingLanguageBest() string
```

### Numbers
```go
Number(min int, max int) int
Int8() int8
Int16() int16
Int32() int32
Int64() int64
Uint8() uint8
Uint16() uint16
Uint32() uint32
Uint64() uint64
Float32() float32
Float32Range(min, max float32) float32
Float64() float64
Float64Range(min, max float64) float64
ShuffleInts(a []int)
RandomInt(i []int) int
```

### String
```go
Digit() string
DigitN(n int) string
Letter() string
LetterN(n int) string
Lexify(str string) string
Numerify(str string) string
ShuffleStrings(a []string)
RandomString(a []string) string
```
