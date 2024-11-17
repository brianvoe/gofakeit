package gofakeit

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v7/data"
)

type createChecksumFn func(digits string) string

// This pattern is inspired from: https://pkg.go.dev/database/sql#NullInt64
type NullString struct {
	String string
	Valid  bool // Valid is true if String is not NULL
}

type ISBN struct {
	Ean         NullString `json:"ean"`
	Group       string     `json:"group"`
	Registrant  string     `json:"registrant"`
	Publication string     `json:"publication"`
	Check       string     `json:"checksum"`
	separator   string
}

func (ele *ISBN) String() string {
	elements := []string{ele.Ean.String, ele.Group, ele.Registrant, ele.Publication, ele.Check}
	if !ele.Ean.Valid {
		elements = elements[1:]
	}
	return strings.Join(elements, ele.separator)
}

func keys[T any](v map[string]T) *[]string {
	var keys []string
	for k := range v {
		keys = append(keys, k)
	}
	return &keys
}

// Registrant and Publication are two separate elements in ISBN string, and can have variable lengths
// depending on rules defined for each ISBN agencies. Currently our logic is focused into breaking
// the combined Reg/Pub string for US Registration Groups only (i.e. 0 & 1).
// This may change in future depending on how many regions we want to support and providing rules
// for each of them.
func getRegistrantPublication(regPub string, rules []data.RegistrantElements) (string, string, error) {
	regLength := len(regPub)
	ruleFound := false
	for _, rule := range rules {
		iRule0, _ := strconv.Atoi(rule.Min)
		iRule1, _ := strconv.Atoi(rule.Max)
		iRegPub, _ := strconv.Atoi(regPub[:regLength-1])
		if iRule0 <= iRegPub && iRegPub <= iRule1 {
			regLength = rule.Length
			ruleFound = true
			break
		}
	}

	if !ruleFound {
		return "", "", fmt.Errorf("registrant/publication %s not found in any Registrant Rules", regPub)
	}

	return regPub[:regLength], regPub[regLength:], nil
}

// Since ISBN is divided into 4/5 string segments of variable lengths, we first need to
// prepare each segment separately and forward it to the consumer.
// Details on each ISBN element can be read here: https://www.isbn-international.org/content/what-isbn/10
func prepareElements(f *Faker) (*ISBN, error) {
	rules := data.ISBNRules
	ean := f.RandomString(*keys(rules))
	regGroup := f.RandomString(*keys(rules[ean]))

	// Based on the lengths of EAN/Registration group, we need to evaluate the length of registrant & publication
	// length which would be length of EAN, Registration Group, Check Digit (checksum), subtracted from
	// ISBN13 string length
	regPubLength := 13 - len(ean) - len(regGroup) - 1
	regPub := f.Numerify(strings.Repeat("#", regPubLength))

	regPubRules := rules[ean][regGroup]
	if registrant, publication, err := getRegistrantPublication(regPub, regPubRules); err != nil {
		return nil, err
	} else {
		return &ISBN{
			Ean:         NullString{ean, true},
			Group:       regGroup,
			Registrant:  registrant,
			Publication: publication,
		}, nil
	}
}

func createISBN(elements *ISBN, fn createChecksumFn, sep string) string {
	digits := elements.String()
	elements.Check = fn(digits)
	elements.separator = sep
	return fmt.Sprint(elements)
}

// Find the checksum/check digit to complete the ISBN string
// Ref: https://en.wikipedia.org/wiki/ISBN#ISBN-10_check_digit_calculation
func createISBN10Checksum(digits string) string {
	sum := 0
	for i, c := range digits {
		digit := int(c - '0')
		digit *= 10 - i
		sum += digit
	}

	var result string
	remainder := (11 - (sum % 11)) % 11
	if remainder == 10 {
		result = "X"
	} else {
		result = strconv.Itoa(remainder)
	}
	return result
}

func createISBN10(f *Faker, sep string) string {
	elements, err := prepareElements(f)
	if err != nil {
		log.Fatalf("unable to generate ISBN10 string: %v", err)
		return ""
	}
	elements.Ean.Valid = false
	return createISBN(elements, createISBN10Checksum, sep)
}

// Find the checksum/check digit to complete the ISBN string
// Ref: https://en.wikipedia.org/wiki/ISBN#ISBN-13_check_digit_calculation
func createISBN13Checksum(digits string) string {
	sum := 0
	for i, c := range digits {
		digit := int(c - '0') // Converts rune to it's ASCII integer value
		if i%2 != 0 {
			digit *= 3
		}
		sum += digit
	}

	remainder := (10 - (sum % 10)) % 10
	return strconv.Itoa(remainder)
}

func createISBN13(f *Faker, sep string) string {
	elements, err := prepareElements(f)
	if err != nil {
		log.Fatalf("unable to generate ISBN13 string: %v", err)
	}
	return createISBN(elements, createISBN13Checksum, sep)
}

func ISBN13(sep string) string            { return createISBN13(GlobalFaker, sep) }
func (f *Faker) ISBN13(sep string) string { return createISBN13(f, sep) }
func isbn13(f *Faker, sep string) string  { return createISBN13(f, sep) }

func ISBN10(sep string) string            { return createISBN10(GlobalFaker, sep) }
func (f *Faker) ISBN10(sep string) string { return createISBN10(f, sep) }
func isbn10(f *Faker, sep string) string  { return createISBN10(f, sep) }

func addISBNLookup() {
	AddFuncLookup("isbn13", Info{
		Display:     "ISBN-13",
		Category:    "isbn",
		Description: "Represents ISBN-13 string, mostly used as book identifiers",
		Example:     "978-0-8125-2759-9",
		Output:      "string",
		Params: []Param{
			{
				Field:       "sep",
				Display:     "Separator",
				Type:        "string",
				Default:     "-",
				Description: "Seperator used to separate individual ISBN elements",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			sep, err := info.GetString(m, "sep")
			if err != nil {
				return nil, err
			}
			return isbn13(f, sep), nil
		},
	})

	AddFuncLookup("isbn10", Info{
		Display:     "ISBN-10",
		Category:    "isbn",
		Description: "Represents ISBN-10 string, mostly used as book identifiers, before 2007",
		Example:     "0-8125-2759-3",
		Output:      "string",
		Params: []Param{
			{
				Field:       "sep",
				Display:     "Separator",
				Type:        "string",
				Default:     "-",
				Description: "Seperator used to separate individual ISBN elements",
			},
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			sep, err := info.GetString(m, "sep")
			if err != nil {
				return nil, err
			}
			return isbn10(f, sep), nil
		},
	})
}
