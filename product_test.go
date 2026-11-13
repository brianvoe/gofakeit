package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

func ExampleProduct() {
	Seed(11)
	product := Product()
	fmt.Println(product.Name)
	fmt.Println(product.Description)
	fmt.Println(product.Categories)
	fmt.Println(product.Price)
	fmt.Println(product.Features)
	fmt.Println(product.Color)
	fmt.Println(product.Material)
	fmt.Println(product.UPC)
	fmt.Println(product.Audience)
	fmt.Println(product.Dimension)
	fmt.Println(product.UseCase)
	fmt.Println(product.Benefit)
	fmt.Println(product.Suffix)

	// Output: Upgradeable Side Table
	// This upset product is crafted from gel and includes 4k ready, making it perfect for learning and delivering advanced technology for gamers and DIY enthusiasts.
	// [smart home devices camping gear]
	// 73.93
	// [multi-functional self-adjusting water-resistant underwater-capable]
	// white
	// canvas
	// 041026894051
	// [gamers outdoor enthusiasts]
	// compact
	// office
	// energy savings
	// art
}

func ExampleFaker_Product() {
	f := New(11)
	product := f.Product()
	fmt.Println(product.Name)
	fmt.Println(product.Description)
	fmt.Println(product.Categories)
	fmt.Println(product.Price)
	fmt.Println(product.Features)
	fmt.Println(product.Color)
	fmt.Println(product.Material)
	fmt.Println(product.UPC)
	fmt.Println(product.Audience)
	fmt.Println(product.Dimension)
	fmt.Println(product.UseCase)
	fmt.Println(product.Benefit)
	fmt.Println(product.Suffix)

	// Output: Upgradeable Side Table
	// This upset product is crafted from gel and includes 4k ready, making it perfect for learning and delivering advanced technology for gamers and DIY enthusiasts.
	// [smart home devices camping gear]
	// 73.93
	// [multi-functional self-adjusting water-resistant underwater-capable]
	// white
	// canvas
	// 041026894051
	// [gamers outdoor enthusiasts]
	// compact
	// office
	// energy savings
	// art
}

func TestProduct(t *testing.T) {
	for range 1000 {
		product := Product()
		if product.Name == "" {
			t.Error("Name is empty")
		}

		if product.Description == "" {
			t.Error("Description is empty")
		}

		if len(product.Categories) == 0 {
			t.Error("Categories is empty")
		}

		if product.Price == 0 {
			t.Error("Price is empty")
		}

		if len(product.Features) == 0 {
			t.Error("Features is empty")
		}

		if product.Color == "" {
			t.Error("Color is empty")
		}

		if product.Material == "" {
			t.Error("Material is empty")
		}

		if product.UPC == "" {
			t.Error("UPC is empty")
		}

		if len(product.Audience) == 0 {
			t.Error("Audience is empty")
		}

		if product.Dimension == "" {
			t.Error("Dimension is empty")
		}

		if len(product.UseCase) == 0 {
			t.Error("UseCase is empty")
		}

		if len(product.Benefit) == 0 {
			t.Error("Benefit is empty")
		}

		if product.Suffix == "" {
			t.Error("Suffix is empty")
		}
	}
}

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Product()
	}
}

func ExampleProductName() {
	Seed(11)
	fmt.Println(ProductName())

	// Output: Luxurious Trophy
}

func ExampleFaker_ProductName() {
	f := New(11)
	fmt.Println(f.ProductName())

	// Output: Luxurious Trophy
}

func BenchmarkProductName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductName()
	}
}

func ExampleProductDescription() {
	Seed(11)
	fmt.Println(ProductDescription())

	// Output: This product, ideal for collectors and hobbyists, features puzzled mica and incorporates water-resistant to ensure safety during sports.
}

func ExampleFaker_ProductDescription() {
	f := New(11)
	fmt.Println(f.ProductDescription())

	// Output: This product, ideal for collectors and hobbyists, features puzzled mica and incorporates water-resistant to ensure safety during sports.
}

// Runs 10,000 tests to ensure description doesnt have any { or } in it
func TestProductDescriptionReplacement(t *testing.T) {
	for range 100000 {
		desc := ProductDescription()
		if strings.ContainsAny(desc, "{") || strings.ContainsAny(desc, "}") {
			t.Error("Description contains { or }")
		}
	}
}

func BenchmarkProductDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDescription()
	}
}

func ExampleProductCategory() {
	Seed(11)
	fmt.Println(ProductCategory())

	// Output: tableware
}

func ExampleFaker_ProductCategory() {
	f := New(11)
	fmt.Println(f.ProductCategory())

	// Output: tableware
}

func BenchmarkProductCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductCategory()
	}
}

func ExampleProductFeature() {
	Seed(11)
	fmt.Println(ProductFeature())

	// Output: ultrasonic
}

func ExampleFaker_ProductFeature() {
	f := New(11)
	fmt.Println(f.ProductFeature())

	// Output: ultrasonic
}

func BenchmarkProductFeature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductFeature()
	}
}

func ExampleProductMaterial() {
	Seed(11)
	fmt.Println(ProductMaterial())

	// Output: textile
}

func ExampleFaker_ProductMaterial() {
	f := New(11)
	fmt.Println(f.ProductMaterial())

	// Output: textile
}

func BenchmarkProductMaterial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductMaterial()
	}
}

func ExampleProductUPC() {
	Seed(11)
	fmt.Println(ProductUPC())

	// Output: 088125275985
}

func ExampleFaker_ProductUPC() {
	f := New(11)
	fmt.Println(f.ProductUPC())

	// Output: 088125275985
}

// A UPC-A has 12 digits and the weighted sum of them, with weights 3 and 1, is a multiple of ten
func isUPC(upc string) bool {
	if len(upc) != 12 {
		return false
	}

	sum := 0
	for i, c := range upc {
		if c < '0' || c > '9' {
			return false
		}
		weight := 1
		if i%2 == 0 {
			weight = 3
		}
		sum += weight * int(c-'0')
	}
	return sum%10 == 0
}

func TestProductUPC(t *testing.T) {
	for _, upc := range []string{"036000291452", "012345678905", "042100005264"} {
		if !isUPC(upc) {
			t.Errorf("Known UPC %s should be valid", upc)
		}
	}
	if isUPC("036000291453") {
		t.Error("A UPC with a wrong check digit should be invalid")
	}

	f := New(11)
	for range 1000 {
		upc := f.ProductUPC()
		if !isUPC(upc) {
			t.Fatalf("Generated UPC %s has an invalid check digit", upc)
		}
	}
}

func BenchmarkProductUPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUPC()
	}
}

func ExampleProductAudience() {
	Seed(11)
	fmt.Println(ProductAudience())

	// Output: [students collectors]
}

func ExampleFaker_ProductAudience() {
	f := New(11)
	fmt.Println(f.ProductAudience())

	// Output: [students collectors]
}

func BenchmarkProductAudience(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductAudience()
	}
}

func ExampleProductDimension() {
	Seed(11)
	fmt.Println(ProductDimension())

	// Output: x-small
}

func ExampleFaker_ProductDimension() {
	f := New(11)
	fmt.Println(f.ProductDimension())

	// Output: x-small
}

func BenchmarkProductDimension(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDimension()
	}
}

func ExampleProductUseCase() {
	Seed(11)
	fmt.Println(ProductUseCase())

	// Output: sports
}

func ExampleFaker_ProductUseCase() {
	f := New(11)
	fmt.Println(f.ProductUseCase())

	// Output: sports
}

func BenchmarkProductUseCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUseCase()
	}
}

func ExampleProductBenefit() {
	Seed(11)
	fmt.Println(ProductBenefit())

	// Output: time-saving
}

func ExampleFaker_ProductBenefit() {
	f := New(11)
	fmt.Println(f.ProductBenefit())

	// Output: time-saving
}

func BenchmarkProductBenefit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductBenefit()
	}
}

func ExampleProductSuffix() {
	Seed(11)
	fmt.Println(ProductSuffix())

	// Output: turbo
}

func ExampleFaker_ProductSuffix() {
	f := New(11)
	fmt.Println(f.ProductSuffix())

	// Output: turbo
}

func BenchmarkProductSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductSuffix()
	}
}

func ExampleProductISBN() {
	Seed(11)
	fmt.Println(ProductISBN(&ISBNOptions{Version: "13", Separator: "-"}))

	// Output: 978-8-8125-2759-5
}

func ExampleFaker_ProductISBN() {
	f := New(11)
	fmt.Println(f.ProductISBN(&ISBNOptions{Version: "13", Separator: "-"}))

	// Output: 978-8-8125-2759-5
}

func TestProductISBN(t *testing.T) {
	f := New(11)

	// ISBN-10: nine digits and a check digit (0-9 or X), the weighted sum with weights 10 to 1 is a multiple of 11
	validISBN10 := func(isbn string) bool {
		if len(isbn) != 10 {
			return false
		}

		sum := 0
		for i, c := range isbn {
			value := int(c - '0')
			if i == 9 && c == 'X' {
				value = 10
			} else if c < '0' || c > '9' {
				return false
			}

			sum += value * (10 - i)
		}

		return sum%11 == 0
	}

	// ISBN-13: thirteen digits, the sum with weights 1 and 3 is a multiple of 10
	validISBN13 := func(isbn string) bool {
		if len(isbn) != 13 {
			return false
		}

		sum := 0
		for i, c := range isbn {
			if c < '0' || c > '9' {
				return false
			}

			if i%2 == 0 {
				sum += int(c - '0')
			} else {
				sum += 3 * int(c-'0')
			}
		}

		return sum%10 == 0
	}

	for range 10000 {
		isbn10 := f.ProductISBN(&ISBNOptions{Version: "10", Separator: "-"})
		if !validISBN10(strings.ReplaceAll(isbn10, "-", "")) {
			t.Fatalf("ISBN-10 %q is not valid", isbn10)
		}

		isbn13 := f.ProductISBN(&ISBNOptions{Version: "13", Separator: "-"})
		if !validISBN13(strings.ReplaceAll(isbn13, "-", "")) {
			t.Fatalf("ISBN-13 %q is not valid", isbn13)
		}
	}
}

func BenchmarkProductISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductISBN(&ISBNOptions{Version: "13", Separator: "-"})
	}
}

func ExampleProductISBN_isbn10() {
	Seed(11)
	fmt.Println(ProductISBN(&ISBNOptions{Version: "10", Separator: "-"}))

	// Output: 8-8125-2759-0
}

func ExampleFaker_ProductISBN_isbn10() {
	f := New(11)
	fmt.Println(f.ProductISBN(&ISBNOptions{Version: "10", Separator: "-"}))

	// Output: 8-8125-2759-0
}
