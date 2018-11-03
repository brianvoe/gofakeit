package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHipsterWord() {
	Seed(11)
	fmt.Println(HipsterWord())
	// Output: microdosing
}

func BenchmarkHipsterWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterWord()
	}
}

func ExampleHipsterSentence() {
	Seed(11)
	fmt.Println(HipsterSentence(5))
	// Output: Microdosing roof chia echo pickled.
}

func BenchmarkHipsterSentence(b *testing.B) {
	Seed(42)
	for i := 0; i < b.N; i++ {
		HipsterSentence(10)
	}
}

func ExampleHipsterParagraph() {
	Seed(11)
	fmt.Println(HipsterParagraph(3, 5, 12, "\n"))
	// Output: Microdosing roof chia echo pickled meditation cold-pressed raw denim fingerstache normcore sriracha pork belly. Wolf try-hard pop-up blog tilde hashtag health butcher waistcoat paleo portland vinegar. Microdosing sartorial blue bottle slow-carb freegan five dollar toast you probably haven't heard of them asymmetrical chia farm-to-table narwhal banjo. Gluten-free blog authentic literally synth vinyl meh ethical health fixie banh mi Yuccie. Try-hard drinking squid seitan cray VHS echo chillwave hammock kombucha food truck sustainable.
	// Pug bushwick hella tote bag cliche direct trade waistcoat yr waistcoat knausgaard pour-over master. Pitchfork jean shorts franzen flexitarian distillery hella meggings austin knausgaard crucifix wolf heirloom. Crucifix food truck you probably haven't heard of them trust fund fixie gentrify pitchfork stumptown mlkshk umami chambray blue bottle. 3 wolf moon swag +1 biodiesel knausgaard semiotics taxidermy meh artisan hoodie +1 blue bottle. Fashion axe forage mixtape Thundercats pork belly whatever 90's beard selfies chambray cred mlkshk.
	// Shabby chic typewriter VHS readymade lo-fi bitters PBR&B gentrify lomo raw denim freegan put a bird on it. Raw denim cliche dreamcatcher pug fixie park trust fund migas fingerstache sriracha +1 mustache. Tilde shoreditch kickstarter franzen dreamcatcher green juice mustache neutra polaroid stumptown organic schlitz. Flexitarian ramps chicharrones kogi lo-fi mustache tilde forage street church-key williamsburg taxidermy. Chia mustache plaid mumblecore squid slow-carb disrupt Thundercats goth shoreditch master direct trade.
}

func BenchmarkHipsterParagraph(b *testing.B) {
	Seed(42)
	for i := 0; i < b.N; i++ {
		HipsterParagraph(3, 5, 12, "\n")
	}
}
