package gofakeit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

type wrapper[T any] struct {
	Value T
}

func (o *wrapper[T]) Receptor() reflect.Value {
	return reflect.ValueOf(&o.Value)
}

// Check that default fake generation was not broken
func ExampleRecipient_notags() {
	f := New(11)
	type testStruct struct {
		A struct {
			B string
		}
	}
	var noTags struct {
		String    wrapper[string]
		StringPtr wrapper[*string]
		Int       wrapper[int]
		Slice     wrapper[[]string]
		SlicePtr  wrapper[*[]string]
		Map       wrapper[map[string]int]
		MapPtr    wrapper[*map[string]int]
		Time      wrapper[time.Time]
		TimePtr   wrapper[*time.Time]
		Struct    wrapper[testStruct]
		StructPtr wrapper[*testStruct]
	}
	if err := f.Struct(&noTags); err != nil {
		fmt.Printf("Recipient noTags error: %v", err)
		return
	}
	m, _ := json.MarshalIndent(noTags, "", "\t")
	fmt.Printf("%s\n", m)
	// Output:
	// {
	// 	"String": {
	// 		"Value": "mwwv"
	// 	},
	// 	"StringPtr": {
	// 		"Value": "AeAwVH"
	// 	},
	// 	"Int": {
	// 		"Value": 3790609878067081083
	// 	},
	// 	"Slice": {
	// 		"Value": [
	// 			"CQyzjqs",
	// 			"gaHqIpy",
	// 			"seUxvh",
	// 			"GVEnWigE",
	// 			"VzByQJJXM",
	// 			"GyLfLRuh",
	// 			"RHQq",
	// 			"MyOwqYsV",
	// 			"XYHS",
	// 			"aXQHvLvfsQ"
	// 		]
	// 	},
	// 	"SlicePtr": {
	// 		"Value": [
	// 			"oQVw",
	// 			"jEFvDCErTA",
	// 			"bMGl",
	// 			"soNQiOZ",
	// 			"QJUI",
	// 			"rmnR",
	// 			"fuGl",
	// 			"HAUoyWW",
	// 			"PnBNKbb"
	// 		]
	// 	},
	// 	"Map": {
	// 		"Value": {
	// 			"BWuMS": 7323578435895571780,
	// 			"DNjxfzPeq": 556126185402360932,
	// 			"EIZKArbndP": 6017639345642484830,
	// 			"JRLWHfau": 504481847486751011,
	// 			"ejKE": 2170064023810017206,
	// 			"nqmqrtj": 2085803212000261243,
	// 			"qjTyxj": 4814308557197142590,
	// 			"qlirZp": 366677794067212421
	// 		}
	// 	},
	// 	"MapPtr": {
	// 		"Value": {
	// 			"AdXkaOVB": 6747100307432224147,
	// 			"BZRr": 2090942964323674902,
	// 			"MAfKVViKzL": 6399228998106268675,
	// 			"OYMnl": 9084218315452557071,
	// 			"QMKdaijT": 3319565698096787296,
	// 			"SDUGngYTGP": 8106940092362766713,
	// 			"uOmmoA": 7076386249594780656,
	// 			"uipZzPwyhH": 6533871989415815454
	// 		}
	// 	},
	// 	"Time": {
	// 		"Value": "1920-08-01T14:50:27.045347196Z"
	// 	},
	// 	"TimePtr": {
	// 		"Value": "1933-03-10T05:22:34.252653011Z"
	// 	},
	// 	"Struct": {
	// 		"Value": {
	// 			"A": {
	// 				"B": "cRXwiWI"
	// 			}
	// 		}
	// 	},
	// 	"StructPtr": {
	// 		"Value": {
	// 			"A": {
	// 				"B": "GfgLRoSjK"
	// 			}
	// 		}
	// 	}
	// }
}

// Basic examples with scalar types
func ExampleRecipient_tags() {
	f := New(1)
	var withTags struct {
		Skip     wrapper[string]  `fake:"-"`
		Sentence wrapper[string]  `fake:"{sentence:5}"`
		Email    wrapper[string]  `fake:"{email}"`
		EmailPtr wrapper[*string] `fake:"{email}"`
		Int      wrapper[int]     `fake:"{number:1,3}"`
	}
	if err := f.Struct(&withTags); err != nil {
		fmt.Printf("Recipient withTags error: %v", err)
		return
	}
	m, _ := json.MarshalIndent(withTags, "", "\t")
	fmt.Printf("%s\n", m)
	// Output:
	// {
	// 	"Skip": {
	// 		"Value": ""
	// 	},
	// 	"Sentence": {
	// 		"Value": "Honour while Bismarckian example product."
	// 	},
	// 	"Email": {
	// 		"Value": "lacybogan@schulist.com"
	// 	},
	// 	"EmailPtr": {
	// 		"Value": "devinmohr@ebert.info"
	// 	},
	// 	"Int": {
	// 		"Value": 1
	// 	}
	// }
}

// Examples with maps/slices, check that fakesize works
func ExampleRecipient_map_slice_tags() {
	f := New(1)
	var withTagMapSlice struct {
		SliceStr    wrapper[[]string]          `fake:"{firstname}" fakesize:"2"`
		SliceStrPtr wrapper[*[]string]         `fake:"{sentence:3}" fakesize:"1"`
		Map         wrapper[map[string]string] `fakesize:"2"`
	}
	if err := f.Struct(&withTagMapSlice); err != nil {
		fmt.Printf("Recipient withTagMapSlice error: %v", err)
		return
	}
	m, _ := json.MarshalIndent(withTagMapSlice, "", "\t")
	fmt.Printf("%s\n", m)
	// Output:
	// {
	// 	"SliceStr": {
	// 		"Value": [
	// 			"Kailyn",
	// 			"Donny"
	// 		]
	// 	},
	// 	"SliceStrPtr": {
	// 		"Value": [
	// 			"Alternatively phew trousers."
	// 		]
	// 	},
	// 	"Map": {
	// 		"Value": {
	// 			"AZspPNGDKv": "CJmo",
	// 			"usUDZFHd": "wNCmJy"
	// 		}
	// 	}
	// }
}

func ExampleRecipient_time() {
	f := New(18)
	type recipient struct {
		Time       wrapper[time.Time]
		TimeMonth  wrapper[time.Time] `fake:"{year}-{month}" format:"2006-1"`
		TimeFormat wrapper[time.Time] `fake:"{year}-{month}-{day}" format:"2006-1-1"`
	}
	var r recipient
	err := f.Struct(&r)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	m, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", m)

	// Output:
	// {
	// 	"Time": {
	// 		"Value": "1909-11-11T21:39:33.574884884Z"
	// 	},
	// 	"TimeMonth": {
	// 		"Value": "1941-09-01T00:00:00Z"
	// 	},
	// 	"TimeFormat": {
	// 		"Value": "1996-08-01T00:00:00Z"
	// 	}
	// }
}

func TestRecipient_customTag(t *testing.T) {
	AddFuncLookup("customTag", Info{
		Category:    "custom",
		Description: "Random int array",
		Example:     "[1]",
		Output:      "CustomType",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			data := make([]int, 1)
			data[0] = 42
			return data, nil
		},
	})
	defer RemoveFuncLookup("customTag")

	type customType []int
	var withCustomTag struct {
		A     wrapper[customType]   `fake:"{customTag}"`
		Array wrapper[[]customType] `fake:"{customTag}"`
	}
	if err := Struct(&withCustomTag); err != nil {
		t.Errorf("Recipient withCustomTag error: %v", err)
	}
	if len(withCustomTag.A.Value) != 1 || withCustomTag.A.Value[0] != 42 {
		t.Errorf("custom tag expected to generate [42], got %v", withCustomTag.A.Value)
	}
	if withCustomTag.Array.Value[0][0] != 42 {
		t.Error("withCustomTag.Array is not populated from custom function")
	}
}

type fakeableRecipient[T any] struct {
	Faked string // populated by Fake()
	Value T
}

func (fr *fakeableRecipient[T]) Fake(faker *Faker) (any, error) {
	return fakeableRecipient[T]{
		Faked: "faked",
	}, nil
}

func (fr *fakeableRecipient[T]) Receptor() reflect.Value {
	return reflect.ValueOf(&fr.Value)
}

func TestRecipient_fakeable(t *testing.T) {

	var withFakeable struct {
		R fakeableRecipient[string] `fake:"{email}"`
	}
	if err := Struct(&withFakeable); err != nil {
		t.Errorf("Recipient withFakeable error: %v\n", err)
	}
	if withFakeable.R.Faked != "faked" {
		t.Errorf("withFakeable: expected string 'faked', got: '%s'", withFakeable.R.Faked)
	}
	if !strings.Contains(withFakeable.R.Value, "@") {
		t.Errorf("withFakeable: expected generated value to be email-like, got: '%s'", withFakeable.R.Value)
	}
}

type Wrapper[T any] wrapper[T]

func (o *Wrapper[T]) Receptor() reflect.Value {
	return reflect.ValueOf(&o.Value)
}
func TestRecipient_embed(t *testing.T) {

	var withEmbed struct {
		Wrapper[string] `fake:"{email}"`
	}
	if err := Struct(&withEmbed); err != nil {
		t.Errorf("Recipient withEmbed error: %v\n", err)
	}
	if !strings.Contains(withEmbed.Value, "@") {
		t.Errorf("Recipient withEmbed: expected email-like value, got '%s'", withEmbed.Value)
	}
}
