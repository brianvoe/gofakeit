package gofakeit

import (
	"testing"
)

func TestRandIntRange(t *testing.T) {
	if randIntRange(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if getRandValue(test) != "" {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"status_code", "notfound"}} {
		if getRandIntValue(test) != 0 {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestRandFloat32RangeSame(t *testing.T) {
	if randFloat32Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestRandFloat64RangeSame(t *testing.T) {
	if randFloat64Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestReplaceWithNumbers(t *testing.T) {
	if replaceWithNumbers("") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	for _, test := range []struct{ in, should string }{
		{"#界#世#", "5界7世8"},
		{"☺#☻☹#", "☺5☻☹7"},
		{"\x80#¼#語", "\x805¼7語"},
	} {
		Seed(42)
		got := replaceWithNumbers(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%s' got '%s' should '%s'",
			test.in, got, test.should)
	}
}

func TestReplaceWithLetters(t *testing.T) {
	if replaceWithLetters("") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithHexLetters(t *testing.T) {
	if "" != replaceWithHexLetters("") {
		t.Error("You should have gotten an empty string")
	}
}
