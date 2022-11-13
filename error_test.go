package gofakeit

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := Error()

	// return value satisfies the error interface
	if fmt.Sprintf("%T", err.Error()) != "string" {
		t.Error("return type must satisfy the error interface")
	}
}
