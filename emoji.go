package gofakeit

// Emoji will return a random fun emoji
func Emoji() string {
	return getRandValue([]string{"emoji", "emoji"})
}

// EmojiDescription will return a random fun emoji description
func EmojiDescription() string {
	return getRandValue([]string{"emoji", "description"})
}

// EmojiCategory will return a random fun emoji category
func EmojiCategory() string {
	return getRandValue([]string{"emoji", "category"})
}

// EmojiAlias will return a random fun emoji alias
func EmojiAlias() string {
	return getRandValue([]string{"emoji", "alias"})
}

// EmojiTag will return a random fun emoji tag
func EmojiTag() string {
	return getRandValue([]string{"emoji", "tag"})
}

func addEmojiLookup() {
	AddFuncLookup("emoji", Info{
		Category:    "emoji",
		Description: "Random emoji",
		Example:     "🤣",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return Emoji(), nil
		},
	})

	AddFuncLookup("emojidescription", Info{
		Category:    "emoji",
		Description: "Random emoji description",
		Example:     "face vomiting",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return EmojiDescription(), nil
		},
	})

	AddFuncLookup("emojicategory", Info{
		Category:    "emoji",
		Description: "Random emoji category",
		Example:     "Smileys & Emotion",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return EmojiCategory(), nil
		},
	})

	AddFuncLookup("emojialias", Info{
		Category:    "emoji",
		Description: "Random emoji alias",
		Example:     "smile",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return EmojiAlias(), nil
		},
	})

	AddFuncLookup("emojitag", Info{
		Category:    "emoji",
		Description: "Random emoji tag",
		Example:     "happy",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return EmojiTag(), nil
		},
	})
}
