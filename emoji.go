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
