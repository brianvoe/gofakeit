package gofakeit

import (
	"fmt"
	"strings"
	"testing"
)

type prefixer string

func (s *prefixer) FakeTransform(f *Faker) error {
	*s = prefixer(fmt.Sprintf("PREFIX.%s", *s))
	return nil
}

func TestTransformable(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		var p prefixer
		if err := Struct(&p); err != nil {
			t.Fatalf("transformable error: %v", err)
		}
		if !strings.HasPrefix(string(p), "PREFIX.") {
			t.Errorf("Prefix transformable: missing prefix, got %s", string(p))
		}
	})

	t.Run("struct", func(t *testing.T) {
		var transformFields struct {
			P1      prefixer
			P2      *prefixer
			Wrapped wrapper[prefixer] `fake:"{email}"`
		}
		if err := Struct(&transformFields); err != nil {
			t.Fatalf("transformable error: %v", err)
		}
		if !strings.HasPrefix(string(transformFields.P1), "PREFIX.") {
			t.Errorf(
				"P1 transformable: missing prefix, got %s",
				string(transformFields.P1),
			)
		}
		if !strings.HasPrefix(string(*transformFields.P2), "PREFIX.") {
			t.Errorf(
				"P2 transformable: missing prefix, got %s",
				string(*transformFields.P2),
			)
		}
		if strings.HasPrefix(string(*transformFields.P2), "PREFIX.PREFIX.") {
			t.Errorf("FakeTransform called multiple times")
		}
		if !strings.HasPrefix(string(transformFields.Wrapped.Value), "PREFIX.") {
			t.Errorf(
				"Wrapped transformable: missing prefix, got %s",
				string(transformFields.Wrapped.Value),
			)
		}
		if !strings.Contains(string(transformFields.Wrapped.Value), "@") {
			t.Errorf(
				"Wrapped transformable: expected email-like string, got '%s'",
				string(transformFields.Wrapped.Value),
			)
		}
	})

}
