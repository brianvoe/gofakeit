package gofakeit

// MinecraftOre will generate a random Minecraft ore
func MinecraftOre() string { return minecraftOre(GlobalFaker) }

// MinecraftOre will generate a random Minecraft ore
func (f *Faker) MinecraftOre() string { return minecraftOre(f) }

func minecraftOre(f *Faker) string { return getRandValue(f, []string{"minecraft", "ore"}) }

// MinecraftWood will generate a random Minecraft wood
func MinecraftWood() string { return minecraftWood(GlobalFaker) }

// MinecraftWood will generate a random Minecraft wood
func (f *Faker) MinecraftWood() string { return minecraftWood(f) }

func minecraftWood(f *Faker) string { return getRandValue(f, []string{"minecraft", "wood"}) }

// MinecraftArmorTier will generate a random Minecraft armor tier
func MinecraftArmorTier() string { return minecraftArmorTier(GlobalFaker) }

// MinecraftArmorTier will generate a random Minecraft armor tier
func (f *Faker) MinecraftArmorTier() string { return minecraftArmorTier(f) }

func minecraftArmorTier(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "armortier"})
}

// MinecraftArmorPart will generate a random Minecraft armor part
func MinecraftArmorPart() string { return minecraftArmorPart(GlobalFaker) }

// MinecraftArmorPart will generate a random Minecraft armor part
func (f *Faker) MinecraftArmorPart() string { return minecraftArmorPart(f) }

func minecraftArmorPart(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "armorpart"})
}

// MinecraftWeapon will generate a random Minecraft weapon
func MinecraftWeapon() string { return minecraftWeapon(GlobalFaker) }

// MinecraftWeapon will generate a random Minecraft weapon
func (f *Faker) MinecraftWeapon() string { return minecraftWeapon(f) }

func minecraftWeapon(f *Faker) string { return getRandValue(f, []string{"minecraft", "weapon"}) }

// MinecraftTool will generate a random Minecraft tool
func MinecraftTool() string { return minecraftTool(GlobalFaker) }

// MinecraftTool will generate a random Minecraft tool
func (f *Faker) MinecraftTool() string { return minecraftTool(f) }

func minecraftTool(f *Faker) string { return getRandValue(f, []string{"minecraft", "tool"}) }

// MinecraftDye will generate a random Minecraft dye
func MinecraftDye() string { return minecraftDye(GlobalFaker) }

// MinecraftDye will generate a random Minecraft dye
func (f *Faker) MinecraftDye() string { return minecraftDye(f) }

func minecraftDye(f *Faker) string { return getRandValue(f, []string{"minecraft", "dye"}) }

// MinecraftFood will generate a random Minecraft food
func MinecraftFood() string { return minecraftFood(GlobalFaker) }

// MinecraftFood will generate a random Minecraft food
func (f *Faker) MinecraftFood() string { return minecraftFood(f) }

func minecraftFood(f *Faker) string { return getRandValue(f, []string{"minecraft", "food"}) }

// MinecraftAnimal will generate a random Minecraft animal
func MinecraftAnimal() string { return minecraftAnimal(GlobalFaker) }

// MinecraftAnimal will generate a random Minecraft animal
func (f *Faker) MinecraftAnimal() string { return minecraftAnimal(f) }

func minecraftAnimal(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "animal"})
}

// MinecraftVillagerJob will generate a random Minecraft villager job
func MinecraftVillagerJob() string { return minecraftVillagerJob(GlobalFaker) }

// MinecraftVillagerJob will generate a random Minecraft villager job
func (f *Faker) MinecraftVillagerJob() string { return minecraftVillagerJob(f) }

func minecraftVillagerJob(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "villagerjob"})
}

// MinecraftVillagerStation will generate a random Minecraft villager station
func MinecraftVillagerStation() string { return minecraftVillagerStation(GlobalFaker) }

// MinecraftVillagerStation will generate a random Minecraft villager station
func (f *Faker) MinecraftVillagerStation() string { return minecraftVillagerStation(f) }

func minecraftVillagerStation(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "villagerstation"})
}

// MinecraftVillagerLevel will generate a random Minecraft villager level
func MinecraftVillagerLevel() string { return minecraftVillagerLevel(GlobalFaker) }

// MinecraftVillagerLevel will generate a random Minecraft villager level
func (f *Faker) MinecraftVillagerLevel() string { return minecraftVillagerLevel(f) }

func minecraftVillagerLevel(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "villagerlevel"})
}

// MinecraftMobPassive will generate a random Minecraft mob passive
func MinecraftMobPassive() string { return minecraftMobPassive(GlobalFaker) }

// MinecraftMobPassive will generate a random Minecraft mob passive
func (f *Faker) MinecraftMobPassive() string { return minecraftMobPassive(f) }

func minecraftMobPassive(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "mobpassive"})
}

// MinecraftMobNeutral will generate a random Minecraft mob neutral
func MinecraftMobNeutral() string { return minecraftMobNeutral(GlobalFaker) }

// MinecraftMobNeutral will generate a random Minecraft mob neutral
func (f *Faker) MinecraftMobNeutral() string { return minecraftMobNeutral(f) }

func minecraftMobNeutral(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "mobneutral"})
}

// MinecraftMobHostile will generate a random Minecraft mob hostile
func MinecraftMobHostile() string { return minecraftMobHostile(GlobalFaker) }

// MinecraftMobHostile will generate a random Minecraft mob hostile
func (f *Faker) MinecraftMobHostile() string { return minecraftMobHostile(f) }

func minecraftMobHostile(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "mobhostile"})
}

// MinecraftMobBoss will generate a random Minecraft mob boss
func MinecraftMobBoss() string { return minecraftMobBoss(GlobalFaker) }

// MinecraftMobBoss will generate a random Minecraft mob boss
func (f *Faker) MinecraftMobBoss() string { return minecraftMobBoss(f) }

func minecraftMobBoss(f *Faker) string {
	return getRandValue(f, []string{"minecraft", "mobboss"})
}

// MinecraftBiome will generate a random Minecraft biome
func MinecraftBiome() string { return minecraftBiome(GlobalFaker) }

// MinecraftBiome will generate a random Minecraft biome
func (f *Faker) MinecraftBiome() string { return minecraftBiome(f) }

func minecraftBiome(f *Faker) string { return getRandValue(f, []string{"minecraft", "biome"}) }

// MinecraftWeather will generate a random Minecraft weather
func MinecraftWeather() string { return minecraftWeather(GlobalFaker) }

// MinecraftWeather will generate a random Minecraft weather
func (f *Faker) MinecraftWeather() string { return minecraftWeather(f) }

func minecraftWeather(f *Faker) string { return getRandValue(f, []string{"minecraft", "weather"}) }

func addMinecraftLookup() {
	AddFuncLookup("minecraftore", Info{
		Display:     "Minecraft ore",
		Category:    "minecraft",
		Description: "Naturally occurring minerals found in the game Minecraft, used for crafting purposes",
		Example:     "coal",
		Output:      "string",
		Aliases:     []string{"ore", "mineral", "minecraft", "game", "crafting"},
		Keywords:    []string{"minecraft", "game", "ore", "mineral", "crafting", "resource", "mining"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftOre(f), nil
		},
	})

	AddFuncLookup("minecraftwood", Info{
		Display:     "Minecraft wood",
		Category:    "minecraft",
		Description: "Natural resource in Minecraft, used for crafting various items and building structures",
		Example:     "oak",
		Output:      "string",
		Aliases:     []string{"wood", "tree", "minecraft", "game", "crafting"},
		Keywords:    []string{"minecraft", "game", "wood", "tree", "crafting", "resource", "building"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftWood(f), nil
		},
	})

	AddFuncLookup("minecraftarmortier", Info{
		Display:     "Minecraft armor tier",
		Category:    "minecraft",
		Description: "Classification system for armor sets in Minecraft, indicating their effectiveness and protection level",
		Example:     "iron",
		Output:      "string",
		Aliases:     []string{"armor", "tier", "minecraft", "game", "protection"},
		Keywords:    []string{"minecraft", "game", "armor", "tier", "protection", "defense", "combat"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftArmorTier(f), nil
		},
	})

	AddFuncLookup("minecraftarmorpart", Info{
		Display:     "Minecraft armor part",
		Category:    "minecraft",
		Description: "Component of an armor set in Minecraft, such as a helmet, chestplate, leggings, or boots",
		Example:     "helmet",
		Output:      "string",
		Aliases:     []string{"armor", "part", "game", "equipment", "protection"},
		Keywords:    []string{"minecraft", "game", "armor", "part", "equipment", "helmet", "chestplate", "leggings", "boots"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftArmorPart(f), nil
		},
	})

	AddFuncLookup("minecraftweapon", Info{
		Display:     "Minecraft weapon",
		Category:    "minecraft",
		Description: "Tools and items used in Minecraft for combat and defeating hostile mobs",
		Example:     "bow",
		Output:      "string",
		Aliases:     []string{"weapon", "game", "combat", "fight", "attack"},
		Keywords:    []string{"minecraft", "game", "weapon", "combat", "fight", "attack", "damage"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftWeapon(f), nil
		},
	})

	AddFuncLookup("minecrafttool", Info{
		Display:     "Minecraft tool",
		Category:    "minecraft",
		Description: "Items in Minecraft designed for specific tasks, including mining, digging, and building",
		Example:     "shovel",
		Output:      "string",
		Aliases:     []string{"tool", "game", "mining", "digging", "building"},
		Keywords:    []string{"minecraft", "game", "tool", "mining", "digging", "building", "crafting"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftTool(f), nil
		},
	})

	AddFuncLookup("minecraftdye", Info{
		Display:     "Minecraft dye",
		Category:    "minecraft",
		Description: "Items used to change the color of various in-game objects",
		Example:     "white",
		Output:      "string",
		Aliases:     []string{"color", "game", "coloring", "pigment", "customization"},
		Keywords:    []string{"minecraft", "game", "dye", "color", "coloring", "customization"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftDye(f), nil
		},
	})

	AddFuncLookup("minecraftfood", Info{
		Display:     "Minecraft food",
		Category:    "minecraft",
		Description: "Consumable items in Minecraft that provide nourishment to the player character",
		Example:     "apple",
		Output:      "string",
		Aliases:     []string{"game", "eating", "nourishment", "consumable", "hunger"},
		Keywords:    []string{"minecraft", "game", "food", "eating", "nourishment", "hunger", "survival"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftFood(f), nil
		},
	})

	AddFuncLookup("minecraftanimal", Info{
		Display:     "Minecraft animal",
		Category:    "minecraft",
		Description: "Non-hostile creatures in Minecraft, often used for resources and farming",
		Example:     "chicken",
		Output:      "string",
		Aliases:     []string{"creature", "farming", "resource", "passive", "mob"},
		Keywords:    []string{"minecraft", "game", "animal", "creature", "farming", "resource", "passive"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftAnimal(f), nil
		},
	})

	AddFuncLookup("minecraftvillagerjob", Info{
		Display:     "Minecraft villager job",
		Category:    "minecraft",
		Description: "The profession or occupation assigned to a villager character in the game",
		Example:     "farmer",
		Output:      "string",
		Aliases:     []string{"villager", "job", "profession", "game", "occupation"},
		Keywords:    []string{"minecraft", "game", "villager", "job", "profession", "occupation", "npc"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftVillagerJob(f), nil
		},
	})

	AddFuncLookup("minecraftvillagerstation", Info{
		Display:     "Minecraft villager station",
		Category:    "minecraft",
		Description: "Designated area or structure in Minecraft where villagers perform their job-related tasks and trading",
		Example:     "furnace",
		Output:      "string",
		Aliases:     []string{"villager", "station", "workstation", "minecraft", "game"},
		Keywords:    []string{"minecraft", "game", "villager", "station", "workstation", "trading", "work"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftVillagerStation(f), nil
		},
	})

	AddFuncLookup("minecraftvillagerlevel", Info{
		Display:     "Minecraft villager level",
		Category:    "minecraft",
		Description: "Measure of a villager's experience and proficiency in their assigned job or profession",
		Example:     "master",
		Output:      "string",
		Aliases:     []string{"villager", "level", "experience", "minecraft", "game"},
		Keywords:    []string{"minecraft", "game", "villager", "level", "experience", "proficiency", "skill"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftVillagerLevel(f), nil
		},
	})

	AddFuncLookup("minecraftmobpassive", Info{
		Display:     "Minecraft mob passive",
		Category:    "minecraft",
		Description: "Non-aggressive creatures in the game that do not attack players",
		Example:     "cow",
		Output:      "string",
		Aliases:     []string{"mob", "passive", "creature", "game", "peaceful"},
		Keywords:    []string{"minecraft", "game", "mob", "passive", "creature", "non-aggressive", "friendly"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftMobPassive(f), nil
		},
	})

	AddFuncLookup("minecraftmobneutral", Info{
		Display:     "Minecraft mob neutral",
		Category:    "minecraft",
		Description: "Creature in the game that only becomes hostile if provoked, typically defending itself when attacked",
		Example:     "bee",
		Output:      "string",
		Aliases:     []string{"mob", "neutral", "creature", "minecraft", "game"},
		Keywords:    []string{"minecraft", "game", "mob", "neutral", "creature", "defensive", "provoked"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftMobNeutral(f), nil
		},
	})

	AddFuncLookup("minecraftmobhostile", Info{
		Display:     "Minecraft mob hostile",
		Category:    "minecraft",
		Description: "Aggressive creatures in the game that actively attack players when encountered",
		Example:     "spider",
		Output:      "string",
		Aliases:     []string{"hostile", "creature", "game", "aggressive", "enemy"},
		Keywords:    []string{"minecraft", "game", "mob", "hostile", "creature", "aggressive", "dangerous"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftMobHostile(f), nil
		},
	})

	AddFuncLookup("minecraftmobboss", Info{
		Display:     "Minecraft mob boss",
		Category:    "minecraft",
		Description: "Powerful hostile creature in the game, often found in challenging dungeons or structures",
		Example:     "ender dragon",
		Output:      "string",
		Aliases:     []string{"mob", "boss", "creature", "minecraft", "game"},
		Keywords:    []string{"minecraft", "game", "mob", "boss", "creature", "powerful", "challenging", "dungeon"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftMobBoss(f), nil
		},
	})

	AddFuncLookup("minecraftbiome", Info{
		Display:     "Minecraft biome",
		Category:    "minecraft",
		Description: "Distinctive environmental regions in the game, characterized by unique terrain, vegetation, and weather",
		Example:     "forest",
		Output:      "string",
		Aliases:     []string{"biome", "environment", "minecraft", "game", "terrain"},
		Keywords:    []string{"minecraft", "game", "biome", "environment", "terrain", "vegetation", "weather"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftBiome(f), nil
		},
	})

	AddFuncLookup("minecraftweather", Info{
		Display:     "Minecraft weather",
		Category:    "minecraft",
		Description: "Atmospheric conditions in the game that include rain, thunderstorms, and clear skies, affecting gameplay and ambiance",
		Example:     "rain",
		Output:      "string",
		Aliases:     []string{"weather", "atmosphere", "minecraft", "game", "climate"},
		Keywords:    []string{"minecraft", "game", "weather", "atmosphere", "climate", "rain", "thunderstorm"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return minecraftWeather(f), nil
		},
	})
}
