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

type ProductInfo struct {
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description"`
	Categories  []string `json:"categories" xml:"categories"`
	Price       float64  `json:"price" xml:"price"`
	Features    []string `json:"features" xml:"features"`
	Color       string   `json:"color" xml:"color"`
	Material    string   `json:"material" xml:"material"`
	UPC         string   `json:"upc" xml:"upc"`
	Audience    []string `json:"audience" xml:"audience"`
	Dimension   string   `json:"dimension" xml:"dimension"`
	UseCase     string   `json:"use_case" xml:"use_case"`
	Benefit     string   `json:"benefit" xml:"benefit"`
	Suffix      string   `json:"suffix" xml:"suffix"`
	ISBN        string   `json:"isbn" xml:"isbn"`
	ISBN10      string   `json:"isbn10" xml:"isbn10"`
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

// Product will generate a random set of product information
func Product() *ProductInfo { return product(GlobalFaker) }

// Product will generate a random set of product information
func (f *Faker) Product() *ProductInfo { return product(f) }

func product(f *Faker) *ProductInfo {
	// Categories
	categories := []string{}
	weightedCategory, _ := weighted(f, []any{1, 2, 3, 4}, []float32{1, 4, 3, 4})

	for i := 0; i < weightedCategory.(int); i++ {
		categories = append(categories, productCategory(f))
	}

	// Features
	features := []string{}
	for i := 0; i < number(f, 1, 5); i++ {
		features = append(features, productFeature(f))
	}

	product := &ProductInfo{
		Name:        productName(f),
		Description: productDescription(f),
		Categories:  categories,
		Price:       price(f, 3.00, 100.00),
		UPC:         productUPC(f),
		Features:    features,
		Color:       safeColor(f),
		Material:    productMaterial(f),
		Audience:    productAudience(f),
		Dimension:   productDimension(f),
		UseCase:     productUseCase(f),
		Benefit:     productBenefit(f),
		Suffix:      productSuffix(f),
		ISBN:        productISBN(f, "-"),
		ISBN10:      productISBN10(f, "-"),
	}

	return product
}

// ProductName will generate a random product name
func ProductName() string { return productName(GlobalFaker) }

// ProductName will generate a random product name
func (f *Faker) ProductName() string { return productName(f) }

func productName(f *Faker) string {
	name := getRandValue(f, []string{"product", "name"})
	switch number(f, 0, 9) {
	case 1:
		// Name + Adjective + Feature
		return title(fmt.Sprintf("%s %s %s", name, getRandValue(f, []string{"product", "adjective"}), productFeature(f)))
	case 2:
		// Adjective + Material + Name
		return title(fmt.Sprintf("%s %s %s", getRandValue(f, []string{"product", "adjective"}), productMaterial(f), name))
	case 3:
		// Color + Name + Suffix
		return title(fmt.Sprintf("%s %s %s", safeColor(f), name, getRandValue(f, []string{"product", "suffix"})))
	case 4:
		// Feature + Name + Adjective
		return title(fmt.Sprintf("%s %s %s", productFeature(f), name, getRandValue(f, []string{"product", "adjective"})))
	case 5:
		// Material + Color + Name
		return title(fmt.Sprintf("%s %s %s", productMaterial(f), safeColor(f), name))
	case 6:
		// Name + Suffix + Material
		return title(fmt.Sprintf("%s %s %s", name, getRandValue(f, []string{"product", "suffix"}), productMaterial(f)))
	case 7:
		// Adjective + Feature + Name
		return title(fmt.Sprintf("%s %s %s", getRandValue(f, []string{"product", "adjective"}), productFeature(f), name))
	case 8:
		// Color + Material + Name
		return title(fmt.Sprintf("%s %s %s", safeColor(f), productMaterial(f), name))
	case 9:
		// Suffix + Adjective + Name
		return title(fmt.Sprintf("%s %s %s", getRandValue(f, []string{"product", "suffix"}), getRandValue(f, []string{"product", "adjective"}), name))
	}

	// case: 0 - Adjective + Name + Suffix
	return title(fmt.Sprintf("%s %s %s", getRandValue(f, []string{"product", "adjective"}), name, getRandValue(f, []string{"product", "suffix"})))
}

// ProductDescription will generate a random product description
func ProductDescription() string { return productDescription(GlobalFaker) }

// ProductDescription will generate a random product description
func (f *Faker) ProductDescription() string { return productDescription(f) }

func productDescription(f *Faker) string {
	prodDesc := getRandValue(f, []string{"product", "description"})

	// Replace all {productaudience} with join "and"
	for strings.Contains(prodDesc, "{productaudience}") {
		prodDesc = strings.Replace(prodDesc, "{productaudience}", strings.Join(productAudience(f), " and "), 1)
	}

	desc, _ := generate(f, prodDesc)
	return desc
}

// ProductCategory will generate a random product category
func ProductCategory() string { return productCategory(GlobalFaker) }

// ProductCategory will generate a random product category
func (f *Faker) ProductCategory() string { return productCategory(f) }

func productCategory(f *Faker) string {
	return getRandValue(f, []string{"product", "category"})
}

// ProductFeature will generate a random product feature
func ProductFeature() string { return productFeature(GlobalFaker) }

// ProductFeature will generate a random product feature
func (f *Faker) ProductFeature() string { return productFeature(f) }

func productFeature(f *Faker) string {
	return getRandValue(f, []string{"product", "feature"})
}

// ProductMaterial will generate a random product material
func ProductMaterial() string { return productMaterial(GlobalFaker) }

// ProductMaterial will generate a random product material
func (f *Faker) ProductMaterial() string { return productMaterial(f) }

func productMaterial(f *Faker) string {
	return getRandValue(f, []string{"product", "material"})
}

// ProductUPC will generate a random product UPC
func ProductUPC() string { return productUPC(GlobalFaker) }

// ProductUPC will generate a random product UPC
func (f *Faker) ProductUPC() string { return productUPC(f) }

func productUPC(f *Faker) string {
	// The first digit of a UPC is a fixed digit (usually 0)
	upc := "0"

	// Generate the remaining 11 digits randomly
	for i := 1; i < 12; i++ {
		digit := number(f, 0, 9)
		upc += fmt.Sprintf("%d", digit)
	}

	return upc
}

// ProductAudience will generate a random target audience
func ProductAudience() []string { return productAudience(GlobalFaker) }

// ProductAudience will generate a random target audience
func (f *Faker) ProductAudience() []string { return productAudience(f) }

func productAudience(f *Faker) []string {
	audiences := []string{}
	for i := 0; i < number(f, 1, 2); i++ {
		// Check if the target audience is already in the list
		// If it is, generate a new target audience
		for {
			audience := getRandValue(f, []string{"product", "target_audience"})
			// Check if in array
			if !stringInSlice(audience, audiences) {
				audiences = append(audiences, audience)
				break
			}
		}
	}
	return audiences
}

// ProductDimension will generate a random product dimension
func ProductDimension() string { return productDimension(GlobalFaker) }

// ProductDimension will generate a random product dimension
func (f *Faker) ProductDimension() string { return productDimension(f) }

func productDimension(f *Faker) string {
	return getRandValue(f, []string{"product", "dimension"})
}

// ProductUseCase will generate a random product use case
func ProductUseCase() string { return productUseCase(GlobalFaker) }

// ProductUseCase will generate a random product use case
func (f *Faker) ProductUseCase() string { return productUseCase(f) }

func productUseCase(f *Faker) string {
	return getRandValue(f, []string{"product", "use_case"})
}

// ProductBenefit will generate a random product benefit
func ProductBenefit() string { return productBenefit(GlobalFaker) }

// ProductBenefit will generate a random product benefit
func (f *Faker) ProductBenefit() string { return productBenefit(f) }

func productBenefit(f *Faker) string {
	return getRandValue(f, []string{"product", "benefit"})
}

// ProductSuffix will generate a random product suffix
func ProductSuffix() string { return productSuffix(GlobalFaker) }

// ProductSuffix will generate a random product suffix
func (f *Faker) ProductSuffix() string { return productSuffix(f) }

func productSuffix(f *Faker) string {
	return getRandValue(f, []string{"product", "suffix"})
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
	// fmt.Println("Digits:", digits)
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

// ProductISBN10 will generate a random ISBN-10 string for the product
func ProductISBN10(sep string) string { return productISBN10(GlobalFaker, sep) }

// ProductISBN10 will generate a random ISBN-10 string for the product
func (f *Faker) ProductISBN10(sep string) string { return productISBN10(f, sep) }

func productISBN10(f *Faker, sep string) string {
	elements, err := prepareElements(f)
	if err != nil {
		log.Fatalf("unable to generate ISBN10 string: %v", err)
		return ""
	}
	elements.Ean.Valid = false
	return createISBN(elements, createISBN10Checksum, sep)
}

// ProductISBN will generate a random ISBN-13 string for the product
func ProductISBN(sep string) string { return productISBN(GlobalFaker, sep) }

// ProductISBN will generate a random ISBN-13 string for the product
func (f *Faker) ProductISBN(sep string) string { return productISBN(f, sep) }

func productISBN(f *Faker, sep string) string {
	elements, err := prepareElements(f)
	if err != nil {
		log.Fatalf("unable to generate ISBN13 string: %v", err)
	}
	return createISBN(elements, createISBN13Checksum, sep)
}

func addProductLookup() {
	AddFuncLookup("product", Info{
		Display:     "Product",
		Category:    "product",
		Description: "An item created for sale or use",
		Example: `{
	"name": "olive copper monitor",
	"description": "Backwards caused quarterly without week it hungry thing someone him regularly. Whomever this revolt hence from his timing as quantity us these yours.",
	"categories": [
		"clothing",
		"tools and hardware"
	],
	"price": 7.61,
	"features": [
		"ultra-lightweight"
	],
	"color": "navy",
	"material": "brass",
	"upc": "012780949980",
	"audience": [
		"adults"
	],
	"dimension": "medium",
	"use_case": "home",
	"benefit": "comfort",
	"suffix": "pro"
	"isbn": "978-1-4028-9462-6"
	"isbn10": "1-4028-9462-7"
}`,
		Output:      "map[string]any",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return product(f), nil
		},
	})

	AddFuncLookup("productname", Info{
		Display:     "Product Name",
		Category:    "product",
		Description: "Distinctive title or label assigned to a product for identification and marketing",
		Example:     "olive copper monitor",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productName(f), nil
		},
	})

	AddFuncLookup("productdescription", Info{
		Display:     "Product Description",
		Category:    "product",
		Description: "Explanation detailing the features and characteristics of a product",
		Example:     "Backwards caused quarterly without week it hungry thing someone him regularly. Whomever this revolt hence from his timing as quantity us these yours.",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productDescription(f), nil
		},
	})

	AddFuncLookup("productcategory", Info{
		Display:     "Product Category",
		Category:    "product",
		Description: "Classification grouping similar products based on shared characteristics or functions",
		Example:     "clothing",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productCategory(f), nil
		},
	})

	AddFuncLookup("productfeature", Info{
		Display:     "Product Feature",
		Category:    "product",
		Description: "Specific characteristic of a product that distinguishes it from others products",
		Example:     "ultra-lightweight",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productFeature(f), nil
		},
	})

	AddFuncLookup("productmaterial", Info{
		Display:     "Product Material",
		Category:    "product",
		Description: "The substance from which a product is made, influencing its appearance, durability, and properties",
		Example:     "brass",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productMaterial(f), nil
		},
	})

	AddFuncLookup("productupc", Info{
		Display:     "Product UPC",
		Category:    "product",
		Description: "Standardized barcode used for product identification and tracking in retail and commerce",
		Example:     "012780949980",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productUPC(f), nil
		},
	})

	AddFuncLookup("productaudience", Info{
		Display:     "Product Audience",
		Category:    "product",
		Description: "The group of people for whom the product is designed or intended",
		Example:     "adults",
		Output:      "[]string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productAudience(f), nil
		},
	})

	AddFuncLookup("productdimension", Info{
		Display:     "Product Dimension",
		Category:    "product",
		Description: "The size or dimension of a product",
		Example:     "medium",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productDimension(f), nil
		},
	})

	AddFuncLookup("productusecase", Info{
		Display:     "Product Use Case",
		Category:    "product",
		Description: "The scenario or purpose for which a product is typically used",
		Example:     "home",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productUseCase(f), nil
		},
	})

	AddFuncLookup("productbenefit", Info{
		Display:     "Product Benefit",
		Category:    "product",
		Description: "The key advantage or value the product provides",
		Example:     "comfort",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productBenefit(f), nil
		},
	})

	AddFuncLookup("productsuffix", Info{
		Display:     "Product Suffix",
		Category:    "product",
		Description: "A suffix used to differentiate product models or versions",
		Example:     "pro",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productSuffix(f), nil
		},
	})

	AddFuncLookup("productisbn", Info{
		Display:     "Product ISBN",
		Category:    "product",
		Description: "ISBN string for the product",
		Example:     "978-1-4028-9462-6",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productISBN(f, "-"), nil
		},
	})

	AddFuncLookup("productisbn10", Info{
		Display:     "Product ISBN-10",
		Category:    "product",
		Description: "ISBN-10 string for the product used before Jan 01, 2007",
		Example:     "1-4028-9462-7",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return productISBN10(f, "-"), nil
		},
	})
}
