package gofakeit

import (
	"testing"
)

func TestPassword(t *testing.T) {
	length := 100

	pass := Password(true, true, true, true, true, length)

	if len(pass) != length {
		t.Error("Password length does not equal requested length")
	}
}
