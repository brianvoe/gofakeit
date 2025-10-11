package gofakeit

// Emoji will return a random fun emoji
func Emoji() string { return emoji(GlobalFaker) }

// Emoji will return a random fun emoji
func (f *Faker) Emoji() string { return emoji(f) }

func emoji(f *Faker) string {
	// Explicit allow-list of emoji subgroups (sorted for deterministic behavior)
	allowed := []string{
		"animal",
		"clothing",
		"costume",
		"electronics",
		"face",
		"flag",
		"food",
		"game",
		"gesture",
		"hand",
		"job",
		"landmark",
		"music",
		"person",
		"plant",
		"sport",
		"tools",
		"vehicle",
		"weather",
	}

	return getRandValue(f, []string{"emoji", randomString(f, allowed)})
}

// EmojiCategory will return a random fun emoji category
func EmojiCategory() string { return emojiCategory(GlobalFaker) }

// EmojiCategory will return a random fun emoji category
func (f *Faker) EmojiCategory() string { return emojiCategory(f) }

func emojiCategory(f *Faker) string { return getRandValue(f, []string{"emoji", "category"}) }

// EmojiAlias will return a random fun emoji alias
func EmojiAlias() string { return emojiAlias(GlobalFaker) }

// EmojiAlias will return a random fun emoji alias
func (f *Faker) EmojiAlias() string { return emojiAlias(f) }

func emojiAlias(f *Faker) string { return getRandValue(f, []string{"emoji", "alias"}) }

// EmojiTag will return a random fun emoji tag
func EmojiTag() string { return emojiTag(GlobalFaker) }

// EmojiTag will return a random fun emoji tag
func (f *Faker) EmojiTag() string { return emojiTag(f) }

func emojiTag(f *Faker) string { return getRandValue(f, []string{"emoji", "tag"}) }

// EmojiFlag will return a random country flag emoji
func EmojiFlag() string { return emojiFlag(GlobalFaker) }

// EmojiFlag will return a random country flag emoji
func (f *Faker) EmojiFlag() string { return emojiFlag(f) }

func emojiFlag(f *Faker) string { return getRandValue(f, []string{"emoji", "flag"}) }

// EmojiAnimal will return a random animal emoji
func EmojiAnimal() string { return emojiAnimal(GlobalFaker) }

// EmojiAnimal will return a random animal emoji
func (f *Faker) EmojiAnimal() string { return emojiAnimal(f) }

func emojiAnimal(f *Faker) string { return getRandValue(f, []string{"emoji", "animal"}) }

// EmojiFood will return a random food emoji
func EmojiFood() string { return emojiFood(GlobalFaker) }

// EmojiFood will return a random food emoji
func (f *Faker) EmojiFood() string { return emojiFood(f) }

func emojiFood(f *Faker) string { return getRandValue(f, []string{"emoji", "food"}) }

// EmojiPlant will return a random plant emoji
func EmojiPlant() string { return emojiPlant(GlobalFaker) }

// EmojiPlant will return a random plant emoji
func (f *Faker) EmojiPlant() string { return emojiPlant(f) }

func emojiPlant(f *Faker) string { return getRandValue(f, []string{"emoji", "plant"}) }

// EmojiMusic will return a random music-related emoji
func EmojiMusic() string { return emojiMusic(GlobalFaker) }

// EmojiMusic will return a random music-related emoji
func (f *Faker) EmojiMusic() string { return emojiMusic(f) }

func emojiMusic(f *Faker) string { return getRandValue(f, []string{"emoji", "music"}) }

// EmojiVehicle will return a random vehicle/transport emoji
func EmojiVehicle() string { return emojiVehicle(GlobalFaker) }

// EmojiVehicle will return a random vehicle/transport emoji
func (f *Faker) EmojiVehicle() string { return emojiVehicle(f) }

func emojiVehicle(f *Faker) string { return getRandValue(f, []string{"emoji", "vehicle"}) }

// EmojiSport will return a random sports emoji
func EmojiSport() string { return emojiSport(GlobalFaker) }

// EmojiSport will return a random sports emoji
func (f *Faker) EmojiSport() string { return emojiSport(f) }

func emojiSport(f *Faker) string { return getRandValue(f, []string{"emoji", "sport"}) }

// EmojiFace will return a random face emoji
func EmojiFace() string { return emojiFace(GlobalFaker) }

// EmojiFace will return a random face emoji
func (f *Faker) EmojiFace() string { return emojiFace(f) }

func emojiFace(f *Faker) string { return getRandValue(f, []string{"emoji", "face"}) }

// EmojiHand will return a random hand emoji
func EmojiHand() string { return emojiHand(GlobalFaker) }

// EmojiHand will return a random hand emoji
func (f *Faker) EmojiHand() string { return emojiHand(f) }

func emojiHand(f *Faker) string { return getRandValue(f, []string{"emoji", "hand"}) }

// EmojiClothing will return a random clothing or accessory emoji
func EmojiClothing() string { return emojiClothing(GlobalFaker) }

// EmojiClothing will return a random clothing or accessory emoji
func (f *Faker) EmojiClothing() string { return emojiClothing(f) }

func emojiClothing(f *Faker) string { return getRandValue(f, []string{"emoji", "clothing"}) }

// EmojiLandmark will return a random landmark or place emoji
func EmojiLandmark() string { return emojiLandmark(GlobalFaker) }

// EmojiLandmark will return a random landmark or place emoji
func (f *Faker) EmojiLandmark() string { return emojiLandmark(f) }

func emojiLandmark(f *Faker) string { return getRandValue(f, []string{"emoji", "landmark"}) }

// EmojiElectronics will return a random electronics/media device emoji
func EmojiElectronics() string { return emojiElectronics(GlobalFaker) }

// EmojiElectronics will return a random electronics/media device emoji
func (f *Faker) EmojiElectronics() string { return emojiElectronics(f) }

func emojiElectronics(f *Faker) string { return getRandValue(f, []string{"emoji", "electronics"}) }

// EmojiGame will return a random game/leisure emoji
func EmojiGame() string { return emojiGame(GlobalFaker) }

// EmojiGame will return a random game/leisure emoji
func (f *Faker) EmojiGame() string { return emojiGame(f) }

func emojiGame(f *Faker) string { return getRandValue(f, []string{"emoji", "game"}) }

// EmojiTools will return a random tools/weapons emoji
func EmojiTools() string { return emojiTools(GlobalFaker) }

// EmojiTools will return a random tools/weapons emoji
func (f *Faker) EmojiTools() string { return emojiTools(f) }

func emojiTools(f *Faker) string { return getRandValue(f, []string{"emoji", "tools"}) }

// EmojiWeather will return a random weather/celestial emoji
func EmojiWeather() string { return emojiWeather(GlobalFaker) }

// EmojiWeather will return a random weather/celestial emoji
func (f *Faker) EmojiWeather() string { return emojiWeather(f) }

func emojiWeather(f *Faker) string { return getRandValue(f, []string{"emoji", "weather"}) }

// EmojiJob will return a random job/occupation emoji
func EmojiJob() string { return emojiJob(GlobalFaker) }

// EmojiJob will return a random job/occupation emoji
func (f *Faker) EmojiJob() string { return emojiJob(f) }

func emojiJob(f *Faker) string { return getRandValue(f, []string{"emoji", "job"}) }

// EmojiPerson will return a random person variant emoji
func EmojiPerson() string { return emojiPerson(GlobalFaker) }

// EmojiPerson will return a random person variant emoji
func (f *Faker) EmojiPerson() string { return emojiPerson(f) }

func emojiPerson(f *Faker) string { return getRandValue(f, []string{"emoji", "person"}) }

// EmojiGesture will return a random gesture emoji
func EmojiGesture() string { return emojiGesture(GlobalFaker) }

// EmojiGesture will return a random gesture emoji
func (f *Faker) EmojiGesture() string { return emojiGesture(f) }

func emojiGesture(f *Faker) string { return getRandValue(f, []string{"emoji", "gesture"}) }

// EmojiCostume will return a random costume/fantasy emoji
func EmojiCostume() string { return emojiCostume(GlobalFaker) }

// EmojiCostume will return a random costume/fantasy emoji
func (f *Faker) EmojiCostume() string { return emojiCostume(f) }

func emojiCostume(f *Faker) string { return getRandValue(f, []string{"emoji", "costume"}) }

// EmojiSentence will return a random sentence with emojis interspersed
func EmojiSentence() string { return emojiSentence(GlobalFaker) }

// EmojiSentence will return a random sentence with emojis interspersed
func (f *Faker) EmojiSentence() string { return emojiSentence(f) }

func emojiSentence(f *Faker) string {
	sentence, err := generate(f, getRandValue(f, []string{"emoji", "sentence"}))
	if err != nil {
		return ""
	}

	return sentence
}

func addEmojiLookup() {
	AddFuncLookup("emoji", Info{
		Display:     "Emoji",
		Category:    "emoji",
		Description: "Digital symbol expressing feelings or ideas in text messages and online chats",
		Example:     "🤣",
		Output:      "string",
		Aliases: []string{
			"emoticon symbol",
			"chat icon",
			"unicode pictograph",
			"emotional glyph",
			"digital expression",
		},
		Keywords: []string{
			"emoji", "symbol", "text", "message", "online", "chats", "ideas", "feelings", "digital", "reaction",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emoji(f), nil
		},
	})

	AddFuncLookup("emojicategory", Info{
		Display:     "Emoji Category",
		Category:    "emoji",
		Description: "Group or classification of emojis based on their common theme or use, like 'smileys' or 'animals'",
		Example:     "Smileys & Emotion",
		Output:      "string",
		Aliases: []string{
			"emoji group",
			"emoji theme",
			"emoji section",
			"emoji classification",
			"emoji grouping",
		},
		Keywords: []string{
			"emoji", "smileys", "emotion", "animals", "theme", "classification", "set", "category", "collection",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiCategory(f), nil
		},
	})

	AddFuncLookup("emojialias", Info{
		Display:     "Emoji Alias",
		Category:    "emoji",
		Description: "Alternative name or keyword used to represent a specific emoji in text or code",
		Example:     "smile",
		Output:      "string",
		Aliases: []string{
			"emoji nickname",
			"emoji shorthand",
			"emoji label",
			"emoji alt text",
			"emoji identifier",
		},
		Keywords: []string{
			"emoji", "alias", "smile", "code", "specific", "represent", "alternative", "keyword", "mapping",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiAlias(f), nil
		},
	})

	AddFuncLookup("emojitag", Info{
		Display:     "Emoji Tag",
		Category:    "emoji",
		Description: "Label or keyword associated with an emoji to categorize or search for it easily",
		Example:     "happy",
		Output:      "string",
		Aliases: []string{
			"emoji keyword",
			"emoji marker",
			"emoji label",
			"emoji hashtag",
			"emoji reference",
		},
		Keywords: []string{
			"emoji", "tag", "happy", "associated", "categorize", "search", "label", "index", "metadata",
		},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiTag(f), nil
		},
	})

	AddFuncLookup("emojiflag", Info{
		Display:     "Emoji Flag",
		Category:    "emoji",
		Description: "Unicode symbol representing a specific country's flag",
		Example:     "🇺🇸",
		Output:      "string",
		Aliases:     []string{"country flag", "flag emoji", "national flag"},
		Keywords:    []string{"emoji", "flag", "country", "national", "unicode", "symbol", "banner"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiFlag(f), nil
		},
	})

	AddFuncLookup("emojianimal", Info{
		Display:     "Emoji Animal",
		Category:    "emoji",
		Description: "Unicode symbol representing an animal",
		Example:     "🐌",
		Output:      "string",
		Aliases:     []string{"animal emoji", "creature emoji", "wildlife emoji"},
		Keywords:    []string{"emoji", "animal", "creature", "wildlife", "pet", "nature", "zoo", "mammal"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiAnimal(f), nil
		},
	})

	AddFuncLookup("emojifood", Info{
		Display:     "Emoji Food",
		Category:    "emoji",
		Description: "Unicode symbol representing food or drink",
		Example:     "🍾",
		Output:      "string",
		Aliases:     []string{"food emoji", "drink emoji", "meal emoji"},
		Keywords:    []string{"emoji", "food", "drink", "meal", "snack", "beverage", "cuisine", "restaurant"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiFood(f), nil
		},
	})

	AddFuncLookup("emojiplant", Info{
		Display:     "Emoji Plant",
		Category:    "emoji",
		Description: "Unicode symbol representing a plant, flower, or tree",
		Example:     "🌻",
		Output:      "string",
		Aliases:     []string{"plant emoji", "flower emoji", "tree emoji"},
		Keywords:    []string{"emoji", "plant", "flower", "tree", "nature", "botanical", "leaf", "garden"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiPlant(f), nil
		},
	})

	AddFuncLookup("emojimusic", Info{
		Display:     "Emoji Music",
		Category:    "emoji",
		Description: "Unicode symbol representing music or musical instruments",
		Example:     "🎵",
		Output:      "string",
		Aliases:     []string{"music emoji", "instrument emoji", "audio emoji"},
		Keywords:    []string{"emoji", "music", "instrument", "audio", "song", "sound", "melody"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiMusic(f), nil
		},
	})

	AddFuncLookup("emojivehicle", Info{
		Display:     "Emoji Vehicle",
		Category:    "emoji",
		Description: "Unicode symbol representing vehicles or transportation",
		Example:     "🚗",
		Output:      "string",
		Aliases:     []string{"vehicle emoji", "transport emoji", "transportation emoji"},
		Keywords:    []string{"emoji", "vehicle", "transport", "transportation", "car", "train", "plane", "boat"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiVehicle(f), nil
		},
	})

	AddFuncLookup("emojisport", Info{
		Display:     "Emoji Sport",
		Category:    "emoji",
		Description: "Unicode symbol representing sports, activities, or awards",
		Example:     "⚽",
		Output:      "string",
		Aliases:     []string{"sport emoji", "sports emoji", "activity emoji"},
		Keywords:    []string{"emoji", "sport", "activity", "game", "award", "team", "fitness", "exercise"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiSport(f), nil
		},
	})

	AddFuncLookup("emojiface", Info{
		Display:     "Emoji Face",
		Category:    "emoji",
		Description: "Unicode symbol representing faces/smileys (including cat/creature faces)",
		Example:     "😀",
		Output:      "string",
		Aliases:     []string{"face emoji", "smiley emoji", "cat face emoji"},
		Keywords:    []string{"emoji", "face", "smiley", "emotion", "expression", "cat", "creature"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiFace(f), nil
		},
	})

	AddFuncLookup("emojihand", Info{
		Display:     "Emoji Hand",
		Category:    "emoji",
		Description: "Unicode symbol representing hand gestures and hand-related symbols",
		Example:     "👍",
		Output:      "string",
		Aliases:     []string{"hand emoji", "gesture emoji", "hand symbol"},
		Keywords:    []string{"emoji", "hand", "gesture", "thumbs", "fingers", "clap", "pray"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiHand(f), nil
		},
	})

	AddFuncLookup("emojiclothing", Info{
		Display:     "Emoji Clothing",
		Category:    "emoji",
		Description: "Unicode symbol representing clothing and accessories",
		Example:     "👗",
		Output:      "string",
		Aliases:     []string{"clothing emoji", "accessory emoji", "garment emoji", "wardrobe emoji"},
		Keywords:    []string{"emoji", "clothing", "apparel", "accessory", "shoes", "hat", "jewelry"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiClothing(f), nil
		},
	})

	AddFuncLookup("emojilandmark", Info{
		Display:     "Emoji Landmark",
		Category:    "emoji",
		Description: "Unicode symbol representing landmarks and notable places/buildings",
		Example:     "🗽",
		Output:      "string",
		Aliases:     []string{"landmark emoji", "place emoji", "building emoji"},
		Keywords:    []string{"emoji", "landmark", "place", "building", "monument", "site"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiLandmark(f), nil
		},
	})

	AddFuncLookup("emojielectronics", Info{
		Display:     "Emoji Electronics",
		Category:    "emoji",
		Description: "Unicode symbol representing electronic and media devices",
		Example:     "💻",
		Output:      "string",
		Aliases:     []string{"electronics emoji", "device emoji", "media emoji"},
		Keywords:    []string{"emoji", "electronics", "device", "media", "computer", "phone", "camera"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiElectronics(f), nil
		},
	})

	AddFuncLookup("emojigame", Info{
		Display:     "Emoji Game",
		Category:    "emoji",
		Description: "Unicode symbol representing games and leisure",
		Example:     "🎮",
		Output:      "string",
		Aliases:     []string{"game emoji", "leisure emoji", "gaming emoji", "play emoji"},
		Keywords:    []string{"emoji", "game", "leisure", "cards", "dice", "puzzle", "toy"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiGame(f), nil
		},
	})

	AddFuncLookup("emojitools", Info{
		Display:     "Emoji Tools",
		Category:    "emoji",
		Description: "Unicode symbol representing tools or similar equipment",
		Example:     "🔨",
		Output:      "string",
		Aliases:     []string{"tool emoji", "equipment emoji", "hardware emoji", "repair emoji"},
		Keywords:    []string{"emoji", "tools", "equipment", "hardware", "fix", "build"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiTools(f), nil
		},
	})

	AddFuncLookup("emojiweather", Info{
		Display:     "Emoji Weather",
		Category:    "emoji",
		Description: "Unicode symbol representing weather and celestial bodies",
		Example:     "☀️",
		Output:      "string",
		Aliases:     []string{"weather emoji", "sky emoji", "forecast emoji", "climate emoji"},
		Keywords:    []string{"emoji", "weather", "sky", "celestial", "cloud", "rain", "sun"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiWeather(f), nil
		},
	})

	AddFuncLookup("emojijob", Info{
		Display:     "Emoji Job",
		Category:    "emoji",
		Description: "Unicode symbol representing people in occupations/roles",
		Example:     "🧑‍💻",
		Output:      "string",
		Aliases:     []string{"job emoji", "occupation emoji", "career emoji", "profession emoji"},
		Keywords:    []string{"emoji", "job", "occupation", "role", "profession", "worker", "person"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiJob(f), nil
		},
	})

	AddFuncLookup("emojiperson", Info{
		Display:     "Emoji Person",
		Category:    "emoji",
		Description: "Unicode symbol representing human person variants",
		Example:     "👩",
		Output:      "string",
		Aliases:     []string{"person emoji", "human emoji", "adult emoji", "child emoji"},
		Keywords:    []string{"emoji", "person", "human", "man", "woman", "adult", "child"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiPerson(f), nil
		},
	})

	AddFuncLookup("emojigesture", Info{
		Display:     "Emoji Gesture",
		Category:    "emoji",
		Description: "Unicode symbol representing person gestures/poses",
		Example:     "🙋",
		Output:      "string",
		Aliases:     []string{"gesture emoji", "pose emoji", "action emoji"},
		Keywords:    []string{"emoji", "gesture", "pose", "action", "person", "hand", "sign"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiGesture(f), nil
		},
	})

	AddFuncLookup("emojicostume", Info{
		Display:     "Emoji Costume",
		Category:    "emoji",
		Description: "Unicode symbol representing costume/fantasy people and roles",
		Example:     "🦸",
		Output:      "string",
		Aliases:     []string{"costume emoji", "fantasy emoji", "role emoji"},
		Keywords:    []string{"emoji", "costume", "fantasy", "superhero", "prince", "princess"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiCostume(f), nil
		},
	})

	AddFuncLookup("emojisentence", Info{
		Display:     "Emoji Sentence",
		Category:    "emoji",
		Description: "Sentence with random emojis interspersed throughout",
		Example:     "Weekends reserve time for 🖼️ Disc 🏨 golf and day.",
		Output:      "string",
		Aliases:     []string{"sentence with emojis", "emoji text", "emoji message"},
		Keywords:    []string{"emoji", "sentence", "text", "message", "interspersed", "random", "words", "expression"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return emojiSentence(f), nil
		},
	})

}
