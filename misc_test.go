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
