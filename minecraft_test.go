package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleMinecraftOre() {
	Seed(11)
	fmt.Println(MinecraftOre())

	// Output: redstone
}

func ExampleFaker_MinecraftOre() {
	f := New(11)
	fmt.Println(f.MinecraftOre())

	// Output: redstone
}

func BenchmarkMinecraftOre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftOre()
	}
}

func ExampleMinecraftWood() {
	Seed(11)
	fmt.Println(MinecraftWood())

	// Output: dark oak
}

func ExampleFaker_MinecraftWood() {
	f := New(11)
	fmt.Println(f.MinecraftWood())

	// Output: dark oak
}

func BenchmarkMinecraftWood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftWood()
	}
}

func ExampleMinecraftArmorTier() {
	Seed(11)
	fmt.Println(MinecraftArmorTier())

	// Output: netherite
}

func ExampleFaker_MinecraftArmorTier() {
	f := New(11)
	fmt.Println(f.MinecraftArmorTier())

	// Output: netherite
}

func BenchmarkMinecraftArmorTier(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftArmorTier()
	}
}

func ExampleMinecraftArmorPart() {
	Seed(11)
	fmt.Println(MinecraftArmorPart())

	// Output:
	// helmet
}

func ExampleFaker_MinecraftArmorPart() {
	f := New(11)
	fmt.Println(f.MinecraftArmorPart())

	// Output:
	// helmet
}

func BenchmarkMinecraftArmorPart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftArmorPart()
	}
}

func ExampleMinecraftWeapon() {
	Seed(11)
	fmt.Println(MinecraftWeapon())

	// Output: shield
}

func ExampleFaker_MinecraftWeapon() {
	f := New(11)
	fmt.Println(f.MinecraftWeapon())

	// Output: shield
}

func BenchmarkMinecraftWeapon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftWeapon()
	}
}

func ExampleMinecraftTool() {
	Seed(11)
	fmt.Println(MinecraftTool())

	// Output: fishing rod
}

func ExampleFaker_MinecraftTool() {
	f := New(11)
	fmt.Println(f.MinecraftTool())

	// Output: fishing rod
}

func BenchmarkMinecraftTool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftTool()
	}
}

func ExampleMinecraftDye() {
	Seed(11)
	fmt.Println(MinecraftDye())

	// Output: yellow
}

func ExampleFaker_MinecraftDye() {
	f := New(11)
	fmt.Println(f.MinecraftDye())

	// Output: yellow
}

func BenchmarkMinecraftDye(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftDye()
	}
}

func ExampleMinecraftFood() {
	Seed(11)
	fmt.Println(MinecraftFood())

	// Output: steak
}

func ExampleFaker_MinecraftFood() {
	f := New(11)
	fmt.Println(f.MinecraftFood())

	// Output: steak
}

func BenchmarkMinecraftFood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftFood()
	}
}

func ExampleMinecraftAnimal() {
	Seed(11)
	fmt.Println(MinecraftAnimal())

	// Output: wolf
}

func ExampleFaker_MinecraftAnimal() {
	f := New(11)
	fmt.Println(f.MinecraftAnimal())

	// Output: wolf
}

func BenchmarkMinecraftAnimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftAnimal()
	}
}

func ExampleMinecraftVillagerJob() {
	Seed(11)
	fmt.Println(MinecraftVillagerJob())

	// Output: toolsmith
}

func ExampleFaker_MinecraftVillagerJob() {
	f := New(11)
	fmt.Println(f.MinecraftVillagerJob())

	// Output:
	// toolsmith
}

func BenchmarkMinecraftVillagerJob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftVillagerJob()
	}
}

func ExampleMinecraftVillagerStation() {
	Seed(11)
	fmt.Println(MinecraftVillagerStation())

	// Output: stonecutter
}

func ExampleFaker_MinecraftVillagerStation() {
	f := New(11)
	fmt.Println(f.MinecraftVillagerStation())

	// Output: stonecutter
}

func BenchmarkMinecraftVillagerStation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftVillagerStation()
	}
}

func ExampleMinecraftVillagerLevel() {
	Seed(11)
	fmt.Println(MinecraftVillagerLevel())

	// Output: master
}

func ExampleFaker_MinecraftVillagerLevel() {
	f := New(11)
	fmt.Println(f.MinecraftVillagerLevel())

	// Output: master
}

func BenchmarkMinecraftVillagerLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftVillagerLevel()
	}
}

func ExampleMinecraftMobPassive() {
	Seed(11)
	fmt.Println(MinecraftMobPassive())

	// Output: turtle
}

func ExampleFaker_MinecraftMobPassive() {
	f := New(11)
	fmt.Println(f.MinecraftMobPassive())

	// Output: turtle
}

func BenchmarkMinecraftMobPassive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftMobPassive()
	}
}

func ExampleMinecraftMobNeutral() {
	Seed(11)
	fmt.Println(MinecraftMobNeutral())

	// Output: wolf
}

func ExampleFaker_MinecraftMobNeutral() {
	f := New(11)
	fmt.Println(f.MinecraftMobNeutral())

	// Output:
	// wolf
}

func BenchmarkMinecraftMobNeutral(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftMobNeutral()
	}
}

func ExampleMinecraftMobHostile() {
	Seed(11)
	fmt.Println(MinecraftMobHostile())

	// Output: wither skeleton
}

func ExampleFaker_MinecraftMobHostile() {
	f := New(11)
	fmt.Println(f.MinecraftMobHostile())

	// Output: wither skeleton
}

func BenchmarkMinecraftMobHostile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftMobHostile()
	}
}

func ExampleMinecraftMobBoss() {
	Seed(11)
	fmt.Println(MinecraftMobBoss())

	// Output:
	// ender dragon
}

func ExampleFaker_MinecraftMobBoss() {
	f := New(11)
	fmt.Println(f.MinecraftMobBoss())

	// Output:
	// ender dragon
}

func BenchmarkMinecraftMobBoss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftMobBoss()
	}
}

func ExampleMinecraftBiome() {
	Seed(11)
	fmt.Println(MinecraftBiome())

	// Output: the nether
}

func ExampleFaker_MinecraftBiome() {
	f := New(11)
	fmt.Println(f.MinecraftBiome())

	// Output: the nether
}

func BenchmarkMinecraftBiome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftBiome()
	}
}

func ExampleMinecraftWeather() {
	Seed(11)
	fmt.Println(MinecraftWeather())

	// Output: thunder
}

func ExampleFaker_MinecraftWeather() {
	f := New(11)
	fmt.Println(f.MinecraftWeather())

	// Output: thunder
}

func BenchmarkMinecraftWeather(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinecraftWeather()
	}
}
