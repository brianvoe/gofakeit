//go:build go1.18
// +build go1.18

package gofakeit

import (
	"encoding/binary"
	"math/rand"
	"regexp"
	"strings"
	"testing"
)

func FuzzRegex(f *testing.F) {
	for i, regex := range regexes {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(i*i*1234567))
		// Reuse TestRegex cases to seed corpus.
		f.Add(buffer, regex.test)

		// Cases found and skipped by fuzz test.
		f.Add([]byte("0"), string("$0"))     // can not match any string
		f.Add([]byte("0"), string("0^00"))   // can not match any string
		f.Add([]byte("0"), string("000\\A")) // `\A` gives unexpected results
		f.Add([]byte("0"), string("\\b"))    // `\b` gives unexpected results

		// Fixed by modifying regexGenerate.
		f.Add([]byte("("), string("((5555555555555555){700})")) // OOM after fuzz run for a while. fixed by adding limit to regexGenerate

		// Fixed in notSoRandom.
		f.Add([]byte("\xff\xff\xff\xff\xff\xff\xff\xf1"), string("123[0-2]w{3}")) // choose 1 in 3 loops forever if notSoRandom loops it self instead of using a random tail
	}
	f.Fuzz(func(t *testing.T, rand []byte, regex string) {
		if len(regex) > 20 {
			return // long regexes take longer to test without adding much, increase this limit will make it easer to find OOM and timeouts.
		}

		// case added after each character gives bad result to get other cases
		if strings.ContainsAny(regex, `^$\`) {
			return
		}

		regex = "^" + regex + "$" // make sure we do a full capture

		// Try to compile regexTest
		regCompile, err := regexp.Compile(regex)
		if err != nil {
			return // Ignore bad regex
		}

		// Let fuzz have control over random behavior
		fuzzRand := &notSoRandom{}
		fuzzRand.useBytes(rand)
		faker := NewCustom(fuzzRand)

		// Generate string and test if it matches the regex syntax
		reg := faker.Regex(regex)
		if reg == "Length limit reached when generating output" {
			// ignore LimitReached error
			return
		}
		if !regCompile.MatchString(reg) {
			t.Error("Generated data does not match regex. Regex: ", regex, " output: ", reg)
		}
	})
}

// notSoRandom is a random source that start with a stream of predetermined data.
// This allows fuzz to slowly change sudo random behavior.
type notSoRandom struct {
	data   []uint64
	offset int

	// Make long tail behavior more random once we run out of data.
	// This avoids dead loop in code that expects a statically random source.
	tail rand.Source64
}

func (r *notSoRandom) Int63() int64 {
	if r.tail != nil {
		return r.tail.Int63()
	}
	return int64(r.Uint64() & ^uint64(1<<63))
}

func (r *notSoRandom) Uint64() uint64 {
	if r.tail != nil {
		return r.tail.Uint64()
	}
	out := r.data[r.offset]
	r.offset = (r.offset + 1) % len(r.data)
	if r.offset == 0 {
		r.tail = rand.NewSource(int64(out)).(rand.Source64)
	}
	return out
}

func (r *notSoRandom) Seed(seed int64) {
	panic("unimplemented")
}

func (r *notSoRandom) useBytes(seed []byte) {
	if len(seed) == 0 || len(seed)%8 != 0 {
		r.useBytes(append(seed, 0))
		return
	}
	var data []uint64
	for i := 0; i+7 < len(seed); i += 8 {
		data = append(data, binary.BigEndian.Uint64(seed[i:]))
	}
	r.data = data
}
