package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleHipsterWord() {
	Seed(11)
	fmt.Println(HipsterWord())

	// Output: semiotics
}

func ExampleFaker_HipsterWord() {
	f := New(11)
	fmt.Println(f.HipsterWord())

	// Output: semiotics
}

func BenchmarkHipsterWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterWord()
	}
}

func ExampleHipsterSentence() {
	Seed(11)
	fmt.Println(HipsterSentence(5))

	// Output: Semiotics everyday you probably haven't heard of them you probably haven't heard of them portland.
}

func ExampleFaker_HipsterSentence() {
	f := New(11)
	fmt.Println(f.HipsterSentence(5))

	// Output: Semiotics everyday you probably haven't heard of them you probably haven't heard of them portland.
}

func BenchmarkHipsterSentence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterSentence(10)
	}
}

func ExampleHipsterParagraph() {
	Seed(11)
	fmt.Println(HipsterParagraph(3, 5, 12, "\n"))

	// Output: Semiotics everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro cardigan Yuccie cred hoodie. Vice pug pug cred intelligentsia roof helvetica squid chambray literally ennui ugh. Chicharrones messenger bag narwhal ennui flannel twee art party mustache sartorial gluten-free cardigan cronut. Hoodie kickstarter cardigan Thundercats heirloom hashtag bitters salvia cleanse forage chartreuse keffiyeh. Sustainable tofu mustache bespoke vice aesthetic iPhone ugh lo-fi health put a bird on it blue bottle.
	// Pop-up pabst pitchfork literally roof tattooed tilde shoreditch green juice ethical celiac tilde. Kombucha kinfolk occupy tacos ramps 90's echo meditation kale chips gluten-free humblebrag keffiyeh. Lo-fi bespoke wayfarers chicharrones crucifix green juice humblebrag organic viral shabby chic locavore cred. Roof forage farm-to-table YOLO williamsburg crucifix blog everyday green juice listicle wayfarers post-ironic. Single-origin coffee cray organic YOLO disrupt venmo tofu meggings fanny pack master craft beer tofu.
	// Hella helvetica microdosing literally meh etsy echo pabst goth readymade +1 marfa. Pork belly hammock kale chips yr green juice stumptown crucifix hella pork belly franzen wolf austin. Offal cred chartreuse freegan intelligentsia twee trust fund paleo pinterest austin typewriter kogi. Pug single-origin coffee ethical irony helvetica beard green juice viral post-ironic Godard slow-carb put a bird on it. Schlitz hammock beard chia kitsch cred salvia irony farm-to-table loko truffaut ramps.
}

func ExampleFaker_HipsterParagraph() {
	f := New(11)
	fmt.Println(f.HipsterParagraph(3, 5, 12, "\n"))

	// Output: Semiotics everyday you probably haven't heard of them you probably haven't heard of them portland austin tattooed retro cardigan Yuccie cred hoodie. Vice pug pug cred intelligentsia roof helvetica squid chambray literally ennui ugh. Chicharrones messenger bag narwhal ennui flannel twee art party mustache sartorial gluten-free cardigan cronut. Hoodie kickstarter cardigan Thundercats heirloom hashtag bitters salvia cleanse forage chartreuse keffiyeh. Sustainable tofu mustache bespoke vice aesthetic iPhone ugh lo-fi health put a bird on it blue bottle.
	// Pop-up pabst pitchfork literally roof tattooed tilde shoreditch green juice ethical celiac tilde. Kombucha kinfolk occupy tacos ramps 90's echo meditation kale chips gluten-free humblebrag keffiyeh. Lo-fi bespoke wayfarers chicharrones crucifix green juice humblebrag organic viral shabby chic locavore cred. Roof forage farm-to-table YOLO williamsburg crucifix blog everyday green juice listicle wayfarers post-ironic. Single-origin coffee cray organic YOLO disrupt venmo tofu meggings fanny pack master craft beer tofu.
	// Hella helvetica microdosing literally meh etsy echo pabst goth readymade +1 marfa. Pork belly hammock kale chips yr green juice stumptown crucifix hella pork belly franzen wolf austin. Offal cred chartreuse freegan intelligentsia twee trust fund paleo pinterest austin typewriter kogi. Pug single-origin coffee ethical irony helvetica beard green juice viral post-ironic Godard slow-carb put a bird on it. Schlitz hammock beard chia kitsch cred salvia irony farm-to-table loko truffaut ramps.
}

func BenchmarkHipsterParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HipsterParagraph(3, 5, 12, "\n")
	}
}
