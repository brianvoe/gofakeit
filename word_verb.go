package gofakeit

import rand "math/rand"

// Verb will generate a random verb
func Verb() string { return verb(globalFaker.Rand) }

// Verb will generate a random verb
func (f *Faker) Verb() string { return verb(f.Rand) }

func verb(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb"}) }

// VerbAction will generate a random action verb
func VerbAction() string { return verbAction(globalFaker.Rand) }

// VerbAction will generate a random action verb
func (f *Faker) VerbAction() string { return verbAction(f.Rand) }

func verbAction(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_action"}) }

// VerbLinking will generate a random linking verb
func VerbLinking() string { return verbLinking(globalFaker.Rand) }

// VerbLinking will generate a random linking verb
func (f *Faker) VerbLinking() string { return verbLinking(f.Rand) }

func verbLinking(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_linking"}) }

// VerbHelping will generate a random helping verb
func VerbHelping() string { return verbHelping(globalFaker.Rand) }

// VerbHelping will generate a random helping verb
func (f *Faker) VerbHelping() string { return verbHelping(f.Rand) }

func verbHelping(r *rand.Rand) string { return getRandValue(r, []string{"word", "verb_helping"}) }
