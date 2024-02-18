package gofakeit

import (
	"fmt"
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

	// Output: Wave Precision Lamp
	// Since previously was that there a tennis occur why. Heels out can fire anyone sometimes. Leap whom troop now scarcely.
	// [cosmetics outdoor gear]
	// 49.18
	// [touchscreen ultra-lightweight gps-enabled biometric]
	// green
	// brass
	// 082447816155
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

	// Output: Wave Precision Lamp
	// Since previously was that there a tennis occur why. Heels out can fire anyone sometimes. Leap whom troop now scarcely.
	// [cosmetics outdoor gear]
	// 49.18
	// [touchscreen ultra-lightweight gps-enabled biometric]
	// green
	// brass
	// 082447816155
}

func TestProduct(t *testing.T) {
	for i := 0; i < 1000; i++ {
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

	// Output: Green Glass Hair Dryer
}

func ExampleFaker_ProductName() {
	f := New(11)
	fmt.Println(f.ProductName())

	// Output: Green Glass Hair Dryer
}

func BenchmarkProductName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductName()
	}
}

func ExampleProductDescription() {
	Seed(11)
	fmt.Println(ProductDescription())

	// Output: Regularly quiver these sprint fight something am elsewhere since previously was that there a. Occur why depend heels out can fire anyone sometimes that leap whom troop now.
}

func ExampleFaker_ProductDescription() {
	f := New(11)
	fmt.Println(f.ProductDescription())

	// Output: Regularly quiver these sprint fight something am elsewhere since previously was that there a. Occur why depend heels out can fire anyone sometimes that leap whom troop now.
}

func BenchmarkProductDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDescription()
	}
}

func ExampleProductCategory() {
	Seed(11)
	fmt.Println(ProductCategory())

	// Output: pet food
}

func ExampleFaker_ProductCategory() {
	f := New(11)
	fmt.Println(f.ProductCategory())

	// Output: pet food
}

func BenchmarkProductCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductCategory()
	}
}

func ExampleProductFeature() {
	Seed(11)
	fmt.Println(ProductFeature())

	// Output: fast-charging
}

func ExampleFaker_ProductFeature() {
	f := New(11)
	fmt.Println(f.ProductFeature())

	// Output: fast-charging
}

func BenchmarkProductFeature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductFeature()
	}
}

func ExampleProductMaterial() {
	Seed(11)
	fmt.Println(ProductMaterial())

	// Output: quartz
}

func ExampleFaker_ProductMaterial() {
	f := New(11)
	fmt.Println(f.ProductMaterial())

	// Output: quartz
}

func BenchmarkProductMaterial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductMaterial()
	}
}

func ExampleProductUPC() {
	Seed(11)
	fmt.Println(ProductUPC())

	// Output: 088125275989
}

func ExampleFaker_ProductUPC() {
	f := New(11)
	fmt.Println(f.ProductUPC())

	// Output: 088125275989
}

func BenchmarkProductUPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUPC()
	}
}
