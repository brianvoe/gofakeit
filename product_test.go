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

	// Output: olive copper monitor
	// Backwards caused quarterly without week it hungry thing someone him regularly. Whomever this revolt hence from his timing as quantity us these yours.
	// [clothing tools and hardware]
	// 7.61
	// [ultra-lightweight]
	// navy
	// brass
	// 012780949980
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

	// Output: olive copper monitor
	// Backwards caused quarterly without week it hungry thing someone him regularly. Whomever this revolt hence from his timing as quantity us these yours.
	// [clothing tools and hardware]
	// 7.61
	// [ultra-lightweight]
	// navy
	// brass
	// 012780949980
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
	b.Run("package", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Product()
		}
	})

	b.Run("Faker math", func(b *testing.B) {
		f := New(0)

		for i := 0; i < b.N; i++ {
			f.Product()
		}
	})

	b.Run("Faker crypto", func(b *testing.B) {
		f := NewCrypto()

		for i := 0; i < b.N; i++ {
			f.Product()
		}
	})
}

func ExampleProductName() {
	Seed(11)
	fmt.Println(ProductName())

	// Output: appliance pulse leather
}

func ExampleFaker_ProductName() {
	f := New(11)
	fmt.Println(f.ProductName())

	// Output: appliance pulse leather
}

func BenchmarkProductName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductName()
	}
}

func ExampleProductDescription() {
	Seed(11)
	fmt.Println(ProductDescription())

	// Output: Generally half where everything niche backwards caused quarterly without week it hungry thing someone him. Today whomever this revolt hence from his timing as quantity us these yours.
}

func ExampleFaker_ProductDescription() {
	f := New(11)
	fmt.Println(f.ProductDescription())

	// Output: Generally half where everything niche backwards caused quarterly without week it hungry thing someone him. Today whomever this revolt hence from his timing as quantity us these yours.
}

func BenchmarkProductDescription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductDescription()
	}
}

func ExampleProductCategory() {
	Seed(11)
	fmt.Println(ProductCategory())

	// Output: pet supplies
}

func ExampleFaker_ProductCategory() {
	f := New(11)
	fmt.Println(f.ProductCategory())

	// Output: pet supplies
}

func BenchmarkProductCategory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductCategory()
	}
}

func ExampleProductFeature() {
	Seed(11)
	fmt.Println(ProductFeature())

	// Output: wireless
}

func ExampleFaker_ProductFeature() {
	f := New(11)
	fmt.Println(f.ProductFeature())

	// Output: wireless
}

func BenchmarkProductFeature(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductFeature()
	}
}

func ExampleProductMaterial() {
	Seed(11)
	fmt.Println(ProductMaterial())

	// Output: silicon
}

func ExampleFaker_ProductMaterial() {
	f := New(11)
	fmt.Println(f.ProductMaterial())

	// Output: silicon
}

func BenchmarkProductMaterial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductMaterial()
	}
}

func ExampleProductUPC() {
	Seed(11)
	fmt.Println(ProductUPC())

	// Output: 056990598130
}

func ExampleFaker_ProductUPC() {
	f := New(11)
	fmt.Println(f.ProductUPC())

	// Output: 056990598130
}

func BenchmarkProductUPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductUPC()
	}
}
