package gofakeit

import "strings"

// HipsterWord will return a single hipster word
func HipsterWord() string { return hipsterWord(GlobalFaker) }

// HipsterWord will return a single hipster word
func (f *Faker) HipsterWord() string { return hipsterWord(f) }

func hipsterWord(f *Faker) string { return getRandValue(f, []string{"hipster", "word"}) }

// HipsterSentence will generate a random hipster sentence
func HipsterSentence() string { return hipsterSentence(GlobalFaker) }

// HipsterSentence will generate a random hipster sentence
func (f *Faker) HipsterSentence() string { return hipsterSentence(f) }

func hipsterSentence(f *Faker) string {
	sentence, err := generate(f, getRandValue(f, []string{"hipster", "sentence"}))
	if err != nil {
		return ""
	}

	return sentence
}

// HipsterParagraph will generate a random hipster paragraph
func HipsterParagraph() string {
	return hipsterParagraph(GlobalFaker)
}

// HipsterParagraph will generate a random hipster paragraph
func (f *Faker) HipsterParagraph() string {
	return hipsterParagraph(f)
}

func hipsterParagraph(f *Faker) string {
	// generate 2-5 sentences
	sentenceCount := f.Number(2, 5)
	sentences := make([]string, sentenceCount)
	for i := 0; i < sentenceCount; i++ {
		sentences[i] = hipsterSentence(f)
	}

	return strings.Join(sentences, " ")
}

func addHipsterLookup() {
	AddFuncLookup("hipsterword", Info{
		Display:     "Hipster Word",
		Category:    "hipster",
		Description: "Trendy and unconventional vocabulary used by hipsters to express unique cultural preferences",
		Example:     "microdosing",
		Output:      "string",
		Aliases:     []string{"word", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords:    []string{"hipster", "preferences", "microdosing", "artisanal", "craft", "organic", "sustainable", "authentic"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterWord(f), nil
		},
	})

	AddFuncLookup("hipstersentence", Info{
		Display:     "Hipster Sentence",
		Category:    "hipster",
		Description: "Sentence showcasing the use of trendy and unconventional vocabulary associated with hipster culture",
		Example:     "Microdosing roof chia echo pickled.",
		Output:      "string",
		Aliases:     []string{"sentence", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords:    []string{"hipster", "showcasing", "microdosing", "roof", "chia", "echo", "pickled", "artisanal"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterSentence(f), nil
		},
	})

	AddFuncLookup("hipsterparagraph", Info{
		Display:     "Hipster Paragraph",
		Category:    "hipster",
		Description: "Paragraph showcasing the use of trendy and unconventional vocabulary associated with hipster culture",
		Example: `Microdosing roof chia echo pickled meditation cold-pressed raw denim fingerstache normcore sriracha pork belly. Wolf try-hard pop-up blog tilde hashtag health butcher waistcoat paleo portland vinegar. Microdosing sartorial blue bottle slow-carb freegan five dollar toast you probably haven't heard of them asymmetrical chia farm-to-table narwhal banjo. Gluten-free blog authentic literally synth vinyl meh ethical health fixie banh mi Yuccie. Try-hard drinking squid seitan cray VHS echo chillwave hammock kombucha food truck sustainable.

Pug bushwick hella tote bag cliche direct trade waistcoat yr waistcoat knausgaard pour-over master. Pitchfork jean shorts franzen flexitarian distillery hella meggings austin knausgaard crucifix wolf heirloom. Crucifix food truck you probably haven't heard of them trust fund fixie gentrify pitchfork stumptown mlkshk umami chambray blue bottle. 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan hoodie +1 blue bottle. Fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk.

Shabby chic typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it. Raw denim cliche dreamcatcher pug fixie park trust fund migas fingerstache sriracha +1 mustache. Tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid stumptown organic schlitz. Flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy. Chia mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade.`,
		Output:   "string",
		Aliases:  []string{"paragraph", "trendy", "unconventional", "vocabulary", "culture", "modern"},
		Keywords: []string{"hipster", "showcasing", "meditation", "cold-pressed", "raw", "denim", "fingerstache", "normcore", "sriracha"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return hipsterParagraph(f), nil
		},
	})
}
