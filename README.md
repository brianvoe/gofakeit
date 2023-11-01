![alt text](https://raw.githubusercontent.com/brianvoe/gofakeit/master/logo.png)

# Gofakeit [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/gofakeit)](https://goreportcard.com/report/github.com/brianvoe/gofakeit) ![Test](https://github.com/brianvoe/gofakeit/workflows/Test/badge.svg?branch=master) [![codecov.io](https://codecov.io/github/brianvoe/gofakeit/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/gofakeit) [![GoDoc](https://godoc.org/github.com/brianvoe/gofakeit/v6?status.svg)](https://godoc.org/github.com/brianvoe/gofakeit/v6) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/gofakeit/master/LICENSE.txt)

Random data generator written in go

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## Features

- [260+ Functions!!!](#functions)
- [Random Sources](#random-sources)
- [Global Rand](#global-rand-set)
- [Struct Generator](#struct)
- [Custom Functions](#custom-functions)
- [Http Server](https://github.com/brianvoe/gofakeit/tree/master/cmd/gofakeitserver)
- [Command Line Tool](https://github.com/brianvoe/gofakeit/tree/master/cmd/gofakeit)
- Zero dependencies
- [Benchmarks](https://github.com/brianvoe/gofakeit/blob/master/BENCHMARKS.md)
- [Issue](https://github.com/brianvoe/gofakeit/issues)

## Contributors

Thanks to everyone who has contributed to Gofakeit!

<a href="https://github.com/brianvoe/gofakeit/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=brianvoe/gofakeit" />
</a>

## Installation

```go
go get github.com/brianvoe/gofakeit/v6
```

## Simple Usage

```go
import "github.com/brianvoe/gofakeit/v6"

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
```

[See full list of functions](#functions)

## Seed

If you are using the default global usage and dont care about seeding no need to set anything.
Gofakeit will seed it with a cryptographically secure number.

If you need a reproducible outcome you can set it via the Seed function call. Every example in
this repo sets it for testing purposes.

```go
import "github.com/brianvoe/gofakeit/v6"

gofakeit.Seed(0) // If 0 will use crypto/rand to generate a number

// or

gofakeit.Seed(8675309) // Set it to whatever number you want
```

## Random Sources

Gofakeit has a few rand sources, by default it uses math.Rand and uses mutex locking to allow for safe goroutines.

If you want to use a more performant source please use NewUnlocked. Be aware that it is not goroutine safe.

```go
import "github.com/brianvoe/gofakeit/v6"

// Uses math/rand(Pseudo) with mutex locking
faker := gofakeit.New(0)

// Uses math/rand(Pseudo) with NO mutext locking
// More performant but not goroutine safe.
faker := gofakeit.NewUnlocked(0)

// Uses crypto/rand(cryptographically secure) with mutext locking
faker := gofakeit.NewCrypto()

// Pass in your own random source
faker := gofakeit.NewCustom()
```

## Global Rand Set

If you would like to use the simple function calls but need to use something like
crypto/rand you can override the default global with the random source that you want.

```go
import "github.com/brianvoe/gofakeit/v6"

faker := gofakeit.NewCrypto()
gofakeit.SetGlobalFaker(faker)
```

## Struct

Gofakeit can generate random data for struct fields. For the most part it covers all the basic type
as well as some non-basic like time.Time.

Struct fields can also use tags to more specifically generate data for that field type.

```go
import "github.com/brianvoe/gofakeit/v6"

// Create structs with random injected data
type Foo struct {
	Str           string
	Int           int
	Pointer       *int
	Name          string         `fake:"{firstname}"`         // Any available function all lowercase
	Sentence      string         `fake:"{sentence:3}"`        // Can call with parameters
	RandStr       string         `fake:"{randomstring:[hello,world]}"`
	Number        string         `fake:"{number:1,10}"`       // Comma separated for multiple values
	Regex         string         `fake:"{regex:[abcdef]{5}}"` // Generate string from regex
	Map           map[string]int `fakesize:"2"`
	Array         []string       `fakesize:"2"`
	ArrayRange    []string       `fakesize:"2,6"`
    Bar           Bar
	Skip          *string        `fake:"skip"`                // Set to "skip" to not generate data for
	SkipAlt       *string        `fake:"-"`                   // Set to "-" to not generate data for
	Created       time.Time                                   // Can take in a fake tag as well as a format tag
	CreatedFormat time.Time      `fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

type Bar struct {
	Name    string
	Number  int
	Float   float32
}

// Pass your struct as a pointer
var f Foo
gofakeit.Struct(&f)

fmt.Println(f.Str)      		// hrukpttuezptneuvunh
fmt.Println(f.Int)      		// -7825289004089916589
fmt.Println(*f.Pointer) 		// -343806609094473732
fmt.Println(f.Name)     		// fred
fmt.Println(f.Sentence) 		// Record river mind.
fmt.Println(f.RandStr)  		// world
fmt.Println(f.Number)   		// 4
fmt.Println(f.Regex)    		// cbdfc
fmt.Println(f.Map)    			// map[PxLIo:52 lxwnqhqc:846]
fmt.Println(f.Array)    		// cbdfc
fmt.Printf("%+v", f.Bar)    	// {Name:QFpZ Number:-2882647639396178786 Float:1.7636692e+37}
fmt.Println(f.Skip)     		// <nil>
fmt.Println(f.Created.String()) // 1908-12-07 04:14:25.685339029 +0000 UTC

// Supported formats
// int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64,
// float32, float64,
// bool, string,
// array, pointers, map
// time.Time // If setting time you can also set a format tag
// Nested Struct Fields and Embedded Fields
```

## Fakeable types

It is possible to extend a struct by implementing the `Fakeable` interface
in order to control the generation.

For example, this is useful when it is not possible to modify the struct that you want to fake by adding struct tags to a field but you still need to be able to control the generation process.

```go
// Custom string that you want to generate your own data for
// or just return a static value
type CustomString string

func (c *CustomString) Fake(faker *gofakeit.Faker) interface{} {
	return CustomString("my custom string")
}

// Imagine a CustomTime type that is needed to support a custom JSON Marshaler
type CustomTime time.Time

func (c *CustomTime) Fake(faker *gofakeit.Faker) interface{} {
	return CustomTime(time.Now())
}

func (c *CustomTime) MarshalJSON() ([]byte, error) {
	//...
}

// This is the struct that we cannot modify to add struct tags
type NotModifiable struct {
	Token string
	Value CustomString
	Creation *CustomTime
}

var f NotModifiable
gofakeit.Struct(&f)
fmt.Printf("%s", f.Token) // yvqqdH
fmt.Printf("%s", f.Value) // my custom string
fmt.Printf("%s", f.Creation) // 2023-04-02 23:00:00 +0000 UTC m=+0.000000001
```

## Custom Functions

In a lot of situations you may need to use your own random function usage for your specific needs.

If you would like to extend the usage of struct tags, generate function, available usages in the gofakeit server
or gofakeit command sub packages. You can do so via the AddFuncLookup. Each function has their own lookup, if
you need more reference examples you can look at each files lookups.

```go
// Simple
gofakeit.AddFuncLookup("friendname", gofakeit.Info{
	Category:    "custom",
	Description: "Random friend name",
	Example:     "bill",
	Output:      "string",
	Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
		return gofakeit.RandomString([]string{"bill", "bob", "sally"}), nil
	},
})

// With Params
gofakeit.AddFuncLookup("jumbleword", gofakeit.Info{
	Category:    "jumbleword",
	Description: "Take a word and jumble it up",
	Example:     "loredlowlh",
	Output:      "string",
	Params: []gofakeit.Param{
		{Field: "word", Type: "string", Description: "Word you want to jumble"},
	},
	Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
		word, err := info.GetString(m, "word")
		if err != nil {
			return nil, err
		}

		split := strings.Split(word, "")
		gofakeit.ShuffleStrings(split)
		return strings.Join(split, ""), nil
	},
})

type Foo struct {
	FriendName string `fake:"{friendname}"`
	JumbleWord string `fake:"{jumbleword:helloworld}"`
}

var f Foo
gofakeit.Struct(&f)
fmt.Printf("%s", f.FriendName) // bill
fmt.Printf("%s", f.JumbleWord) // loredlowlh
```

## Functions

All functions also exist as methods on the Faker struct

### File

Passing `nil` to `CSV`, `JSON` or `XML` it will auto generate data using a random set of generators.

```go
CSV(co *CSVOptions) ([]byte, error)
JSON(jo *JSONOptions) ([]byte, error)
XML(xo *XMLOptions) ([]byte, error)
FileExtension() string
FileMimeType() string
```

### Person

```go
Person() *PersonInfo
Name() string
NamePrefix() string
NameSuffix() string
FirstName() string
MiddleName() string
LastName() string
Gender() string
SSN() string
Hobby() string
Contact() *ContactInfo
Email() string
Phone() string
PhoneFormatted() string
Teams(peopleArray []string, teamsArray []string) map[string][]string
```

### Generate

```go
Struct(v interface{})
Slice(v interface{})
Map() map[string]interface{}
Generate(value string) string
Regex(value string) string
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
Dice(numDice uint, sides []uint) []uint
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

### Car

```go
Car() *CarInfo
CarMaker() string
CarModel() string
CarType() string
CarFuelType() string
CarTransmissionType() string
```

### Words

#### Noun

```go
Noun() string
NounCommon() string
NounConcrete() string
NounAbstract() string
NounCollectivePeople() string
NounCollectiveAnimal() string
NounCollectiveThing() string
NounCountable() string
NounUncountable() string
```

#### Verb

```go
Verb() string
VerbAction() string
VerbLinking() string
VerbHelping() string
```

#### Adverb

```go
Adverb() string
AdverbManner() string
AdverbDegree() string
AdverbPlace() string
AdverbTimeDefinite() string
AdverbTimeIndefinite() string
AdverbFrequencyDefinite() string
AdverbFrequencyIndefinite() string
```

#### Proposition

```go
Preposition() string
PrepositionSimple() string
PrepositionDouble() string
PrepositionCompound() string
```

#### Adjective

```go
Adjective() string
AdjectiveDescriptive() string
AdjectiveQuantitative() string
AdjectiveProper() string
AdjectiveDemonstrative() string
AdjectivePossessive() string
AdjectiveInterrogative() string
AdjectiveIndefinite() string
```

#### Pronoun

```go
Pronoun() string
PronounPersonal() string
PronounObject() string
PronounPossessive() string
PronounReflective() string
PronounDemonstrative() string
PronounInterrogative() string
PronounRelative() string
```

#### Connective

```go
Connective() string
ConnectiveTime() string
ConnectiveComparative() string
ConnectiveComplaint() string
ConnectiveListing() string
ConnectiveCasual() string
ConnectiveExamplify() string
```

#### Word

```go
Word() string
```

#### Sentences

```go
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
FlipACoin() string
RandomMapKey(mapI interface{}) interface{}
ShuffleAnySlice(v interface{})
```

### Colors

```go
Color() string
HexColor() string
RGBColor() []int
SafeColor() string
NiceColors() string
```

### Images

```go
ImageURL(width int, height int) string
Image(width int, height int) *img.RGBA
ImageJpeg(width int, height int) []byte
ImagePng(width int, height int) []byte
```

### Internet

```go
URL() string
DomainName() string
DomainSuffix() string
IPv4Address() string
IPv6Address() string
MacAddress() string
HTTPStatusCode() string
HTTPStatusCodeSimple() int
LogLevel(logType string) string
HTTPMethod() string
HTTPVersion() string
UserAgent() string
ChromeUserAgent() string
FirefoxUserAgent() string
OperaUserAgent() string
SafariUserAgent() string
```

### HTML

```go
InputName() string
Svg(options *SVGOptions) string
```

### Date/Time

```go
Date() time.Time
FutureDate() time.Time
DateRange(start, end time.Time) time.Time
NanoSecond() int
Second() int
Minute() int
Hour() int
Month() int
MonthString() string
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

### Finance

```go
Cusip() string
Isin() string
```

### Company

```go
BS() string
Blurb() string
BuzzWord() string
Company() string
CompanySuffix() string
Job() *JobInfo
JobDescriptor() string
JobLevel() string
JobTitle() string
Slogan() string
```

### Hacker

```go
HackerAbbreviation() string
HackerAdjective() string
Hackeringverb() string
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
Bird() string
```

### Emoji

```go
Emoji() string
EmojiDescription() string
EmojiCategory() string
EmojiAlias() string
EmojiTag() string
```

### Language

```go
Language() string
LanguageAbbreviation() string
ProgrammingLanguage() string
ProgrammingLanguageBest() string
```

### Number

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
HexUint8() string
HexUint16() string
HexUint32() string
HexUint64() string
HexUint128() string
HexUint256() string
```

### String

```go
Digit() string
DigitN(n uint) string
Letter() string
LetterN(n uint) string
Lexify(str string) string
Numerify(str string) string
ShuffleStrings(a []string)
RandomString(a []string) string
```

### Celebrity

```go
CelebrityActor() string
CelebrityBusiness() string
CelebritySport() string
```

### Minecraft

```go
MinecraftOre() string
MinecraftWood() string
MinecraftArmorTier() string
MinecraftArmorPart() string
MinecraftWeapon() string
MinecraftTool() string
MinecraftDye() string
MinecraftFood() string
MinecraftAnimal() string
MinecraftVillagerJob() string
MinecraftVillagerStation() string
MinecraftVillagerLevel() string
MinecraftMobPassive() string
MinecraftMobNeutral() string
MinecraftMobHostile() string
MinecraftMobBoss() string
MinecraftBiome() string
MinecraftWeather() string
```

### Book

```go
Book() *BookInfo
BookTitle() string
BookAuthor() string
BookGenre() string
```

### Movie

```go
Movie() *MovieInfo
MovieName() string
MovieGenre() string
```

### Error

Unlike most `gofakeit` methods which return a `string`, the error methods return a Go `error`. Access the error message as a string by chaining the `.Error()` method.

```go
Error() error
ErrorDatabase() error
ErrorGRPC() error
ErrorHTTP() error
ErrorHTTPClient() error
ErrorHTTPServer() error
ErrorInput() error
ErrorRuntime() error
```

### School

```go
School() string
```