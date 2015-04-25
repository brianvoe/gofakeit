package gofakeit

import (
	"testing"
)

func TestUniqueId(t *testing.T) {
	id := UniqueId()

	if len(id) != 32 {
		t.Error("unique length does not equal requested length")
	}
}
