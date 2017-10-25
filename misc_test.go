package gofakeit

import "testing"

func TestRandIntRange(t *testing.T) {
	if randIntRange(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestGetRandValueFail(t *testing.T) {
	if getRandValue([]string{"not", "found"}) != "" {
		t.Error("You should have gotten no value back")
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	if getRandIntValue([]string{"not", "found"}) != 0 {
		t.Error("You should have gotten no value back")
	}
}

func TestRandFloatRangeSame(t *testing.T) {
	if randFloatRange(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}
